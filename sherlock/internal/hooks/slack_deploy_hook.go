package hooks

import (
	"errors"
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/broadinstitute/sherlock/sherlock/internal/slack"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"sort"
	"strings"
)

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

	// TODO: Send the main message
	log.Debug().Msg("TODO: Send the main message")

	if ciRun.TerminalAt != nil {
		// TODO: Send a message with the changelog
		log.Debug().Msg("TODO: Send a message with the changelog")

		if !ciRun.Succeeded() {
			// TODO: Send a message into the channel alerting to the failure
			log.Debug().Msg("TODO: Send a message into the channel alerting to the failure")
		}
	}

	// If the CiRun is ongoing, update the state, otherwise delete it if it exists
	if ciRun.TerminalAt == nil {
		if err := db.Save(&messageState).Error; err != nil {
			return fmt.Errorf("failed to save SlackDeployHookState for SlackDeployHook %d and CiRun %d: %w", hook.ID, ciRun.ID, err)
		}
	} else if messageState.CiRunID != 0 && messageState.SlackDeployHookID != 0 {
		if err := db.Delete(&messageState).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("failed to delete SlackDeployHookState for SlackDeployHook %d and CiRun %d: %w", hook.ID, ciRun.ID, err)
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