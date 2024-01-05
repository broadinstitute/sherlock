package slack

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/rs/zerolog/log"
	"github.com/slack-go/slack"
	"math/rand"
)

type DeploymentNotificationInputs struct {
	Title      string
	EntryLines []string
	FooterText []string
}

func makeDeploymentNotificationBlocks(inputs DeploymentNotificationInputs) []slack.Block {
	blocks := make([]slack.Block, 0)
	if inputs.Title != "" {
		blocks = append(blocks, slack.NewSectionBlock(slack.NewTextBlockObject("mrkdwn", inputs.Title, false, true), nil, nil))
	}
	blocks = append(blocks, utils.Map(inputs.EntryLines, func(text string) slack.Block {
		return slack.NewSectionBlock(slack.NewTextBlockObject("mrkdwn", text, false, true), nil, nil)
	})...)
	if len(inputs.FooterText) > 0 {
		blocks = append(blocks, slack.NewContextBlock("",
			utils.Map(inputs.FooterText, func(text string) slack.MixedElement {
				return slack.NewTextBlockObject("mrkdwn", text, false, true)
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
	if err != nil {
		if bytes, jsonErr := json.Marshal(blocks); jsonErr != nil {
			err = fmt.Errorf("(also failed to marshal blocks to JSON: %v) %v", jsonErr, err)
		} else {
			identifier := rand.Int()
			log.Warn().Bytes("blocks", bytes).Int("identifier", identifier).Msg("failed to send deployment notification, embedding blocks in log")
			err = fmt.Errorf("(embedded blocks in log, seek identifier %d) %v", identifier, err)
		}
	}
	return
}

func SendDeploymentChangelogNotification(ctx context.Context, channel, timestamp, title string, sections [][]string) error {
	blocks := []slack.Block{
		slack.NewSectionBlock(slack.NewTextBlockObject("mrkdwn", title, false, true), nil, nil),
	}
	for sectionIdx, section := range sections {
		blocks = append(blocks, chunkLinesToSectionMrkdwnBlocks(section)...)
		if sectionIdx < len(sections)-1 {
			blocks = append(blocks, slack.NewDividerBlock())
		}
	}
	if isEnabled() && len(blocks) > 1 {
		var chunks [][]slack.Block
		for 50 < len(blocks) {
			blocks, chunks = blocks[50:], append(chunks, blocks[0:50:50])
		}
		for _, chunk := range append(chunks, blocks) {
			_, _, _, err := client.SendMessageContext(ctx, channel,
				slack.MsgOptionTS(timestamp),
				slack.MsgOptionBlocks(chunk...))
			if err != nil {
				if bytes, jsonErr := json.Marshal(blocks); jsonErr != nil {
					err = fmt.Errorf("(also failed to marshal blocks to JSON: %v) %v", jsonErr, err)
				} else {
					identifier := rand.Int()
					log.Warn().Bytes("blocks", bytes).Int("identifier", identifier).Msg("failed to send deployment changelog notification, embedding blocks in log")
					err = fmt.Errorf("(embedded blocks in log, seek identifier %d) %v", identifier, err)
				}
			}
			return err
		}
	}
	return nil
}

func SendDeploymentFailureNotification(ctx context.Context, channel, timestamp, text string) error {
	if isEnabled() && timestamp != "" && text != "" {
		_, _, _, err := client.SendMessageContext(ctx, channel,
			slack.MsgOptionTS(timestamp),
			slack.MsgOptionBroadcast(),
			slack.MsgOptionBlocks(slack.NewSectionBlock(slack.NewTextBlockObject("mrkdwn", text, false, true), nil, nil)))
		return err
	}
	return nil
}
