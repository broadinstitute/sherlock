package ci_hooks

import (
	"fmt"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/broadinstitute/sherlock/sherlock/internal/slack"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"slices"
)

func (_ *dispatcherImpl) DispatchSlackDeployHook(db *gorm.DB, hook models.SlackDeployHook, ciRun models.CiRun) error {
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
		Preload("NewAppVersions.AuthoredBy").
		Preload("NewChartVersions.AuthoredBy").
		Find(&changesets).
		Error; err != nil {
		return fmt.Errorf("failed to query Changesets for CiRun %d: %w", ciRun.ID, err)
	}

	slices.SortFunc(changesets, models.CompareChangesetsByName)
	deploymentByUsers := models.UsersFromChangesets(changesets)

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
		status, emoji, hasChangesetFailure := slackDeployHookParseStatus(rawStatus)
		hasFailure = hasFailure || hasChangesetFailure
		mainMessage.EntryLines = append(mainMessage.EntryLines,
			fmt.Sprintf("*:%s: %s* [%s]: %s",
				emoji,
				changeset.ChartRelease.Name,
				changeset.Summarize(false),
				status))
	}
	if len(mainMessage.EntryLines) == 0 {
		mainMessage.EntryLines = append(mainMessage.EntryLines, "(No changes listed)")
	}
	var beehiveUrl string
	if len(changesets) > 0 {
		beehiveUrl = slackDeployHookBeehiveLink(changesets)
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
	if len(deploymentByUsers) > 0 {
		mainMessage.FooterText = append(mainMessage.FooterText,
			slackDeployHookSummarizeUsers(deploymentByUsers, hook.MentionPeople != nil && *hook.MentionPeople))
	}

	var messageState models.SlackDeployHookState
	if tx := db.
		Where(&models.SlackDeployHookState{
			SlackDeployHookID: hook.ID,
			CiRunID:           ciRun.ID,
		}).
		Attrs(&models.SlackDeployHookState{
			MessageChannel: *hook.SlackChannel,
		}).
		FirstOrCreate(&messageState); tx.Error != nil {
		return fmt.Errorf("failed to query SlackDeployHookState for SlackDeployHook %d and CiRun %d: %w", hook.ID, ciRun.ID, tx.Error)
	} else if tx.RowsAffected == 0 && messageState.MessageTimestamp == "" {
		// If rows affected is 0, we found a record rather than creating one.
		// If message timestamp is empty, there's not been a message fully sent for this run yet.
		// That means we should actually not send a message at all, because some other instance is in between initializing
		// the state and sending the message. If we were to send one too, we'd be a duplicate send.
		log.Info().
			Uint("CiRun", ciRun.ID).
			Uint("SlackDeployHook", hook.ID).
			Msg("Skipping SlackDeployHook dispatch because another instance is already sending the first message")
		return nil
	}

	if channelFromResponse, timestampFromResponse, err := slack.SendDeploymentNotification(
		db.Statement.Context, messageState.MessageChannel, messageState.MessageTimestamp, mainMessage); err != nil {
		// If we errored sending the message, and we were sending rather than updating, we should actually delete the state record so the next instance can try again.
		if messageState.MessageTimestamp == "" {
			if recoverableErr := db.Delete(&messageState).Error; recoverableErr != nil {
				log.Error().Err(err).Msgf("failed to delete in-progress SlackDeployHookState for SlackDeployHook %d and CiRun %d (was erroring out on initial send)", hook.ID, ciRun.ID)
			}
		}
		return fmt.Errorf("failed to send deployment notification for CiRun %d: %w", ciRun.ID, err)
	} else {
		messageState.MessageChannel = channelFromResponse
		messageState.MessageTimestamp = timestampFromResponse
		if recoverableErr := db.Save(&messageState).Error; recoverableErr != nil {
			log.Error().Err(err).Msgf("failed to save in-progress SlackDeployHookState for SlackDeployHook %d and CiRun %d but continuing anyway", hook.ID, ciRun.ID)
		}
	}

	// If the run is complete or there's already a failure, send a changelog to @ everyone who had changes go out
	if !messageState.ChangelogSent && (ciRun.TerminalAt != nil || hasFailure) {
		sectionsPerChart := slackDeployHookChangesetsToChangelogSections(changesets, hook.MentionPeople != nil && *hook.MentionPeople, beehiveUrl)
		title := slackDeployHookChangelogTitle(hasFailure, hook.Trigger.SlackBeehiveLink())
		if err := slack.SendDeploymentChangelogNotification(
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
		if err := slack.SendDeploymentFailureNotification(
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
