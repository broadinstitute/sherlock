package hooks

import (
	"cmp"
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/broadinstitute/sherlock/sherlock/internal/slack"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"regexp"
	"slices"
	"sort"
	"strings"
)

var numericProgressRegex = regexp.MustCompile(`(\d+)/(\d+)`)

func (_ *dispatcherImpl) DispatchSlackDeployHook(db *gorm.DB, hook models.SlackDeployHook, ciRun models.CiRun) error {
	// Bail out to the old behavior by default
	if hook.Beta == nil || !*hook.Beta {
		if ciRun.TerminalAt != nil {
			return deprecatedSlackDeployHook(db, hook, ciRun)
		} else {
			return nil
		}
	}

	if hook.SlackChannel == nil {
		// Shouldn't ever hit this case, but better to error out than panic
		return fmt.Errorf("slack channel was nil on SlackDeployHook %d, shouldn't be possible", hook.ID)
	}

	changesetStatuses := make(map[uint]string)
	var changesetIDs []uint
	for _, ciIdentifier := range ciRun.RelatedResources {
		if ciIdentifier.ResourceType == "changeset" {
			changesetIDs = append(changesetIDs, ciIdentifier.ResourceID)
			if ciIdentifier.ResourceStatus != nil {
				changesetStatuses[ciIdentifier.ResourceID] = *ciIdentifier.ResourceStatus
			}
		}
	}
	var changesets []models.Changeset
	if err := db.
		Model(&models.Changeset{}).
		Where("id IN ?", changesetIDs).
		Preload(clause.Associations).
		Find(&changesets).
		Error; err != nil {
		return fmt.Errorf("failed to query Changesets for CiRun %d: %w", ciRun.ID, err)
	}
	slices.SortFunc(changesets, func(a, b models.Changeset) int {
		if a.ChartRelease == nil && b.ChartRelease == nil {
			return 0
		} else if a.ChartRelease == nil {
			return -1
		} else if b.ChartRelease == nil {
			return 1
		} else {
			return cmp.Compare(a.ChartRelease.Name, b.ChartRelease.Name)
		}
	})
	var deploymentByUsers []models.User
	for _, changeset := range changesets {
		if changeset.AppliedBy != nil {
			var exists bool
			for _, existing := range deploymentByUsers {
				if existing.ID == changeset.AppliedBy.ID {
					exists = true
					break
				}
			}
			if !exists {
				deploymentByUsers = append(deploymentByUsers, *changeset.AppliedBy)
			}
		}
		if changeset.PlannedBy != nil && (changeset.AppliedBy == nil || changeset.PlannedBy.ID != changeset.AppliedBy.ID) {
			var exists bool
			for _, existing := range deploymentByUsers {
				if existing.ID == changeset.PlannedBy.ID {
					exists = true
					break
				}
			}
			if !exists {
				deploymentByUsers = append(deploymentByUsers, *changeset.PlannedBy)
			}
		}
	}

	// Has failures starts out true if and only if the overall workflow has finished and failed.
	// As we assemble the message, we'll also set it to true if any of the individual changesets have failed.
	hasFailure := ciRun.TerminalAt != nil && !ciRun.Succeeded()
	var mainMessage slack.DeploymentNotificationInputs
	mainMessage.Title = fmt.Sprintf("Deployment to *%s* %s:", hook.Trigger.SlackBeehiveLink(), ciRun.DoneOrUnderway())
	for _, changeset := range changesets {
		var rawStatus string
		if changesetStatus, ok := changesetStatuses[changeset.ID]; ok {
			rawStatus = changesetStatus
		} else if ciRun.Status != nil {
			rawStatus = *ciRun.Status
		}

		emoji := config.Config.String("slack.emoji.unknown")
		status := rawStatus
		if strings.HasPrefix(rawStatus, "queued") {
			// "queued" is both a GHA status and a Thelma phase
			emoji = config.Config.String("slack.emoji.beehiveWaiting")
			status = "Waiting..."
		} else if strings.HasPrefix(rawStatus, "running") || rawStatus == "in_progress" {
			// "running" is a Thelma phase, "in_progress" is a GHA status
			emoji = config.Config.String("slack.emoji.beehiveLoading")
			status = "Progressing..."
			if numericProgress := numericProgressRegex.FindStringSubmatch(rawStatus); len(numericProgress) >= 3 {
				// If there's something like "1/3" in the rawStatus, set status to "Progressing... 33%"
				numerator, numeratorErr := utils.ParseInt(numericProgress[1])
				denominator, denominatorErr := utils.ParseInt(numericProgress[2])
				if numeratorErr == nil && denominatorErr == nil && denominator != 0 {
					status = fmt.Sprintf("Progressing... %d%%", int(float64(numerator)/float64(denominator)*100))
				}
			}
		} else if strings.HasPrefix(rawStatus, "success") {
			// "success" is both a GHA status and a Thelma phase
			emoji = config.Config.String("slack.emoji.succeeded")
			status = "Success"
		} else if strings.HasPrefix(rawStatus, "error") || rawStatus == "failure" {
			// "error" is a Thelma phase, "failure" is a GHA status
			emoji = config.Config.String("slack.emoji.failed")
			status = "Failed"
			hasFailure = true
		}
		mainMessage.EntryLines = append(mainMessage.EntryLines,
			fmt.Sprintf("*:%s: %s* [%s]: %s",
				emoji,
				changeset.ChartRelease.Name,
				changeset.Summarize(false),
				status))
	}
	if len(deploymentByUsers) > 0 {
		mainMessage.FooterText = append(mainMessage.FooterText, fmt.Sprintf("By %s", strings.Join(
			utils.Map(deploymentByUsers, func(u models.User) string {
				if hook.MentionPeople != nil && *hook.MentionPeople {
					return u.SlackReference()
				} else {
					return u.NameOrEmailHandle()
				}
			}),
			", ")))
	}
	var beehiveUrl string
	if len(changesets) > 0 {
		beehiveUrl = fmt.Sprintf("%s?%s",
			config.Config.String("beehive.reviewChangesetsUrl"),
			strings.Join(
				utils.Map(changesets, func(c models.Changeset) string { return fmt.Sprintf("changeset=%d", c.ID) }),
				"&"))
		mainMessage.FooterText = append(mainMessage.FooterText, slack.LinkHelper(
			beehiveUrl,
			"Beehive"))
	}
	if ciRun.Platform == "github-actions" {
		mainMessage.FooterText = append(mainMessage.FooterText, slack.LinkHelper(
			ciRun.WebURL(),
			"GitHub Actions"))
	} else {
		mainMessage.FooterText = append(mainMessage.FooterText, slack.LinkHelper(
			ciRun.WebURL(),
			"Workflow"))
	}
	if argoCdUrl, ok := hook.Trigger.ArgoCdUrl(); ok {
		mainMessage.FooterText = append(mainMessage.FooterText, slack.LinkHelper(
			argoCdUrl,
			"Argo CD"))
	}

	var messageState models.SlackDeployHookState
	if err := db.
		Where(&models.SlackDeployHookState{
			SlackDeployHookID: hook.ID,
			CiRunID:           ciRun.ID,
		}).
		Attrs(&models.SlackDeployHookState{
			MessageChannel: *hook.SlackChannel,
		}).
		FirstOrInit(&messageState).
		Error; err != nil {
		return fmt.Errorf("failed to query SlackDeployHookState for SlackDeployHook %d and CiRun %d: %w", hook.ID, ciRun.ID, err)
	}
	var err error
	messageState.MessageChannel, messageState.MessageTimestamp, err = slack.SendDeploymentNotification(
		db.Statement.Context, messageState.MessageChannel, messageState.MessageTimestamp, mainMessage)
	if err != nil {
		return fmt.Errorf("failed to send deployment notification for CiRun %d: %w", ciRun.ID, err)
	}
	if recoverableErr := db.Save(&messageState).Error; recoverableErr != nil {
		log.Error().Err(err).Msgf("failed to save in-progress SlackDeployHookState for SlackDeployHook %d and CiRun %d but continuing anyway", hook.ID, ciRun.ID)
	}

	// If the run is complete or there's already a failure, send a changelog to @ everyone who had changes go out
	if !messageState.ChangelogSent && (ciRun.TerminalAt != nil || hasFailure) {
		sectionsPerChart := make([][]string, 0)
		for _, changeset := range changesets {
			sections := make([]string, 0)
			sections = append(sections, fmt.Sprintf("*%s* [%s]", changeset.ChartRelease.SlackBeehiveLink(), changeset.Summarize(true)))
			for _, version := range models.InterleaveVersions(changeset.NewAppVersions, changeset.NewChartVersions) {
				sections = append(sections, version.SlackChangelogEntry(hook.MentionPeople != nil && *hook.MentionPeople))
			}
			if len(sections) == 1 {
				sections = append(sections, fmt.Sprintf("- *No changelog entries found;* %s", slack.LinkHelper(beehiveUrl, "Beehive might have more information")))
			}
			sectionsPerChart = append(sectionsPerChart, sections)
		}
		var title string
		if hasFailure {
			title = fmt.Sprintf("Failures deploying to *%s*, changelog:", hook.Trigger.SlackBeehiveLink())
		} else {
			title = fmt.Sprintf("Successfully deployed to *%s*, changelog:", hook.Trigger.SlackBeehiveLink())
		}
		if err = slack.SendDeploymentChangelogNotification(
			db.Statement.Context, messageState.MessageChannel, messageState.MessageTimestamp,
			title, sectionsPerChart); err != nil {
			return fmt.Errorf("failed to send deployment changelog notification for CiRun %d: %w", ciRun.ID, err)
		} else {
			messageState.ChangelogSent = true
			if recoverableErr := db.Save(&messageState).Error; recoverableErr != nil {
				log.Error().Err(err).Msgf("failed to save in-progress SlackDeployHookState for SlackDeployHook %d and CiRun %d but continuing anyway", hook.ID, ciRun.ID)
			}
		}
	}

	// If there's a failure, send an alert in the thread/channel
	if !messageState.FailureAlertSent && hasFailure {
		if err = slack.SendDeploymentFailureNotification(
			db.Statement.Context, messageState.MessageChannel, messageState.MessageTimestamp,
			fmt.Sprintf(":%s: Failures deploying to *%s*, please %s",
				config.Config.String("slack.emoji.alert"),
				hook.Trigger.SlackBeehiveLink(),
				slack.LinkHelper(ciRun.WebURL(), "take a look"))); err != nil {
			return fmt.Errorf("failed to send deployment failure notification for CiRun %d: %w", ciRun.ID, err)
		} else {
			messageState.FailureAlertSent = true
			if recoverableErr := db.Save(&messageState).Error; recoverableErr != nil {
				log.Error().Err(err).Msgf("failed to save in-progress SlackDeployHookState for SlackDeployHook %d and CiRun %d but continuing anyway", hook.ID, ciRun.ID)
			}
		}
	}
	return nil
}

// Deprecated
func deprecatedSlackDeployHook(db *gorm.DB, hook models.SlackDeployHook, ciRun models.CiRun) error {
	if hook.SlackChannel == nil {
		return fmt.Errorf("slack channel was nil on SlackDeployHook %d, shouldn't be possible", hook.ID)
	} else if attachment, err := generateSlackAttachment(db, hook, ciRun); err != nil {
		return fmt.Errorf("failed to generate summary of CiRun %d for SlackDeployHook %d: %w", ciRun.ID, hook.ID, err)
	} else if config.Config.Bool("slack.behaviors.deployHooks.enable") {
		slack.SendMessage(db.Statement.Context, *hook.SlackChannel, "", attachment)
	}
	return nil
}

func generateSlackAttachment(db *gorm.DB, hook models.SlackDeployHook, ciRun models.CiRun) (slack.Attachment, error) {
	var sb strings.Builder

	sb.WriteString("Deployment to ")

	if hook.Trigger.OnEnvironment != nil {
		sb.WriteString(slack.LinkHelper(
			fmt.Sprintf(config.Config.String("beehive.environmentUrlFormatString"), hook.Trigger.OnEnvironment.Name),
			hook.Trigger.OnEnvironment.Name))

		var chartReleaseNames []string
		for _, ciIdentifier := range ciRun.RelatedResources {
			if ciIdentifier.ResourceType == "chart-release" {
				var chartRelease models.ChartRelease
				if err := db.First(&chartRelease, ciIdentifier.ResourceID).Error; err != nil {
					return nil, err
				} else {
					chartReleaseNames = append(chartReleaseNames, chartRelease.Name)
				}
			}
		}

		if len(chartReleaseNames) > 0 {
			sort.Strings(chartReleaseNames)
			sb.WriteString(" (")
			sb.WriteString(strings.Join(utils.Map(chartReleaseNames, func(name string) string {
				return slack.LinkHelper(
					fmt.Sprintf(config.Config.String("beehive.chartReleaseUrlFormatString"), name),
					name)
			}), ", "))
			sb.WriteString(")")
		}

	} else if hook.Trigger.OnChartRelease != nil {
		sb.WriteString(slack.LinkHelper(
			fmt.Sprintf(config.Config.String("beehive.chartReleaseUrlFormatString"), hook.Trigger.OnChartRelease.Name),
			hook.Trigger.OnChartRelease.Name))
	} else {
		return nil, fmt.Errorf("SlackDeployHook %d didn't have Trigger fully loaded", hook.ID)
	}

	sb.WriteString(": ")

	if ciRun.Status != nil {
		sb.WriteString(slack.LinkHelper(ciRun.WebURL(), *ciRun.Status))
	} else {
		return nil, fmt.Errorf("CiRun %d didn't have status present", ciRun.ID)
	}

	sb.WriteString(".")

	var changesetIDs []uint
	for _, ciIdentifier := range ciRun.RelatedResources {
		if ciIdentifier.ResourceType == "changeset" {
			changesetIDs = append(changesetIDs, ciIdentifier.ResourceID)
		}
	}

	if len(changesetIDs) > 0 {
		sb.WriteString(" Review all changes made by this deployment ")

		var changesetLinkSB strings.Builder
		changesetLinkSB.WriteString(config.Config.String("beehive.reviewChangesetsUrl"))
		changesetLinkSB.WriteString("?")
		changesetLinkSB.WriteString(strings.Join(utils.Map(changesetIDs, func(id uint) string {
			return fmt.Sprintf("changeset=%d", id)
		}), "&"))

		sb.WriteString(slack.LinkHelper(changesetLinkSB.String(), "here"))
		sb.WriteString(".")
	}

	if ciRun.Succeeded() {
		return slack.GreenBlock{Text: sb.String()}, nil
	} else {
		return slack.RedBlock{Text: sb.String()}, nil
	}
}
