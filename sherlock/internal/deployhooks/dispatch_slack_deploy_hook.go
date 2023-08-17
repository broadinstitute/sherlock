package deployhooks

import (
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/broadinstitute/sherlock/sherlock/internal/slack"
	"gorm.io/gorm"
	"strings"
)

func dispatchSlackDeployHook(db *gorm.DB, hook models.SlackDeployHook, ciRun models.CiRun) error {
	if hook.SlackChannel == nil {
		return fmt.Errorf("slack channel was nil on SlackDeployHook %d, shouldn't be possible", hook.ID)
	} else if attachment, err := generateSlackAttachment(db, hook, ciRun); err != nil {
		return fmt.Errorf("failed to generate summary of CiRun %d for SlackDeployHook %d: %v", ciRun.ID, hook.ID, err)
	} else {
		slack.SendMessage(db.Statement.Context, *hook.SlackChannel, "", attachment)
		return nil
	}
}

func generateSlackAttachment(_ *gorm.DB, hook models.SlackDeployHook, ciRun models.CiRun) (slack.Attachment, error) {
	var sb strings.Builder

	sb.WriteString("Deployment to ")

	if hook.Trigger.OnEnvironment != nil {
		sb.WriteString(hook.Trigger.OnEnvironment.Name)
	} else if hook.Trigger.OnChartRelease != nil {
		sb.WriteString(hook.Trigger.OnChartRelease.Name)
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
