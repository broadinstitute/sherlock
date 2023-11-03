package slack

import (
	"context"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/rs/zerolog/log"
	"github.com/slack-go/slack"
)

type Attachment interface {
	toSlackAttachment() slack.Attachment
}

type GreenBlock struct {
	Text string
}

func (g GreenBlock) toSlackAttachment() slack.Attachment {
	return slack.Attachment{
		Color: config.Config.String("slack.colors.green"),
		Text:  g.Text,
	}
}

type RedBlock struct {
	Text string
}

func (g RedBlock) toSlackAttachment() slack.Attachment {
	return slack.Attachment{
		Color: config.Config.String("slack.colors.red"),
		Text:  g.Text,
	}
}

func SendMessage(ctx context.Context, channel string, text string, attachments ...Attachment) {
	if err := SendMessageReturnError(ctx, channel, text, attachments...); err != nil {
		log.Warn().Err(err).Msgf("SLCK | unable to send message to %s: %v", channel, err)
	}
}

func SendMessageReturnError(ctx context.Context, channel string, text string, attachments ...Attachment) error {
	if isEnabled() && (text != "" || len(attachments) > 0) {
		var options []slack.MsgOption
		if text != "" {
			options = append(options, slack.MsgOptionText(text, false))
		}
		if len(attachments) > 0 {
			options = append(options, slack.MsgOptionAttachments(
				utils.Map(attachments, func(a Attachment) slack.Attachment { return a.toSlackAttachment() })...,
			))
		}
		_, _, _, err := client.SendMessageContext(ctx, channel, options...)
		return err
	}
	return nil
}
