package slack

import (
	"context"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/slack-go/slack"
)

type DeploymentNotificationInputs struct {
	Title      string
	EntryLines []string
	FooterText []string
}

func makeDeploymentNotificationBlocks(inputs DeploymentNotificationInputs) []slack.Block {
	blocks := make([]slack.Block, 0)
	if inputs.Title != "" {
		blocks = append(blocks, slack.NewTextBlockObject("mrkdwn", inputs.Title, true, true))
	}
	blocks = append(blocks, utils.Map(inputs.EntryLines, func(text string) slack.Block {
		return slack.NewTextBlockObject("mrkdwn", text, true, true)
	})...)
	if len(inputs.FooterText) > 0 {
		blocks = append(blocks, slack.NewContextBlock("",
			utils.Map(inputs.FooterText, func(text string) slack.MixedElement {
				return slack.NewTextBlockObject("mrkdwn", text, true, true)
			})...))
	}
	return blocks
}

func SendDeploymentNotification(ctx context.Context, channel, timestamp string, inputs DeploymentNotificationInputs) (_channel, _timestamp string, err error) {
	blocks := makeDeploymentNotificationBlocks(inputs)
	if isEnabled() && len(blocks) > 0 {
		opts := []slack.MsgOption{slack.MsgOptionBlocks(blocks...)}
		if timestamp != "" {
			_channel, _timestamp, _, err = client.UpdateMessageContext(ctx, channel, timestamp, opts...)
		} else {
			_channel, _timestamp, _, err = client.SendMessageContext(ctx, channel, opts...)
		}
	}
	return
}

func SendDeploymentChangelogNotification(ctx context.Context, channel, timestamp, title string, sections [][]string) error {
	blocks := []slack.Block{
		slack.NewTextBlockObject("mrkdwn", title, true, true),
	}
	for sectionIdx, section := range sections {
		for _, textBlob := range section {
			blocks = append(blocks, slack.NewTextBlockObject("mrkdwn", textBlob, true, true))
		}
		if sectionIdx < len(sections)-1 {
			blocks = append(blocks, slack.NewDividerBlock())
		}
	}
	if isEnabled() && len(blocks) > 1 {
		_, _, _, err := client.SendMessageContext(ctx, channel,
			slack.MsgOptionTS(timestamp),
			slack.MsgOptionBlocks(blocks...))
		return err
	}
	return nil
}

func SendDeploymentFailureNotification(ctx context.Context, channel, timestamp, text string) error {
	if isEnabled() && timestamp != "" && text != "" {
		_, _, _, err := client.SendMessageContext(ctx, channel,
			slack.MsgOptionTS(timestamp),
			slack.MsgOptionBroadcast(),
			slack.MsgOptionBlocks(slack.NewTextBlockObject("mrkdwn", text, true, true)))
		return err
	}
	return nil
}
