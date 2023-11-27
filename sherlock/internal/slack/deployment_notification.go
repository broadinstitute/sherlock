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
	blocks = append(blocks, chunkLinesToMrkdwnBlocks(inputs.EntryLines)...)
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
