package slack

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"

	"github.com/rs/zerolog/log"
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
		// We don't expect these to be multiple blocks, but better safe than sorry
		blocks = append(blocks, chunkLinesToSectionMrkdwnBlocks([]string{inputs.Title})...)
	}
	for _, line := range inputs.EntryLines {
		// We put each line through manually because we want each line as its own section for spacing
		blocks = append(blocks, chunkLinesToSectionMrkdwnBlocks([]string{line})...)
	}
	if len(inputs.FooterText) > 0 {
		contextBlocks := make([]slack.MixedElement, 0, len(inputs.FooterText))
		for _, text := range inputs.FooterText {
			if len(text) > slackTextBlockLengthLimit {
				text = text[:slackTextBlockLengthLimit-3] + "..."
			}
			contextBlocks = append(contextBlocks, slack.NewTextBlockObject("mrkdwn", text, false, true))
		}
		if len(contextBlocks) > 0 {
			blocks = append(blocks, slack.NewContextBlock("", contextBlocks...))
		}
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
	} else {
		// Make sure we don't return emptier channel/timestamp than we were given, even if we did nothing in this case
		_channel = channel
		_timestamp = timestamp
	}
	if err != nil {
		if bytes, jsonErr := json.Marshal(blocks); jsonErr != nil {
			err = fmt.Errorf("(also failed to marshal blocks to JSON: %v) %v", jsonErr, err)
		} else {
			identifier := rand.Int()
			log.Warn().Bytes("blocks", bytes).Int("identifier", identifier).Msgf("failed to send deployment notification, embedding blocks in log (identifier %d)", identifier)
			err = fmt.Errorf("(embedded blocks in log, seek identifier %d) %v", identifier, err)
		}
	}
	return
}

func SendDeploymentChangelogNotification(ctx context.Context, channel, timestamp, title string, sections [][]string) error {
	blocks := chunkLinesToSectionMrkdwnBlocks([]string{title})
	for sectionIdx, section := range sections {
		blocks = append(blocks, chunkLinesToSectionMrkdwnBlocks(section)...)
		if sectionIdx < len(sections)-1 {
			blocks = append(blocks, slack.NewDividerBlock())
		}
	}
	if isEnabled() && timestamp != "" && len(sections) > 0 && len(blocks) > 0 {
		var chunks [][]slack.Block
		for len(blocks) > 50 {
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
					log.Warn().Bytes("blocks", bytes).Int("identifier", identifier).Msgf("failed to send deployment changelog notification, embedding blocks in log (identifier %d)", identifier)
					err = fmt.Errorf("(embedded blocks in log, seek identifier %d) %v", identifier, err)
				}
				return err
			}
		}
	}
	return nil
}

func SendDeploymentFailureNotification(ctx context.Context, channel, timestamp, text string) error {
	if isEnabled() && timestamp != "" && text != "" {
		_, _, _, err := client.SendMessageContext(ctx, channel,
			slack.MsgOptionTS(timestamp),
			slack.MsgOptionBroadcast(),
			slack.MsgOptionBlocks(chunkLinesToSectionMrkdwnBlocks([]string{text})...))
		return err
	}
	return nil
}
