package slack

import (
	"encoding/json"
	"fmt"
	"math/rand"

	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/rs/zerolog/log"
	"github.com/slack-go/slack"
	"golang.org/x/net/context"
)

const permissionChangeSquelchContextKey = "sherlock-slack-permission-change-squelch"

// SetContextToSquelchPermissionChangeNotifications should be used very very carefully. It creates a new
// context that, if passed to SendPermissionChangeNotification or SendPermissionChangeNotificationReturnError,
// will prevent a notification from actually being sent if there's no errors.
//
// This is useful when regular routines have multiple notifications (e.g. they send their own but the changes
// to Role or RoleAssignment tables also trigger their own notifications). This can be used to stop inner
// notifications from being sent if they don't have any errors to communicate.
func SetContextToSquelchPermissionChangeNotifications(ctx context.Context) context.Context {
	return context.WithValue(ctx, permissionChangeSquelchContextKey, true)
}

type PermissionChangeNotificationInputs struct {
	Summary string
	Results []string
	Errors  []error
}

func SendPermissionChangeNotification(ctx context.Context, actor string, inputs PermissionChangeNotificationInputs) {
	toRun := func() {
		if errs := SendPermissionChangeNotificationReturnError(ctx, actor, inputs); len(errs) > 0 {
			ReportError(ctx, "unable to send permission change notification", errs...)
		}
	}
	// Permission changes happen very, *very* frequently during tests as we add test users and roles. Rather than
	// making all manner of tests worry about asynchronous behavior, we just run the notification synchronously in
	// debug mode. Note that "run" is relative, because the client library won't be called unless it is mocked.
	if config.Config.String("mode") == "debug" {
		toRun()
	} else {
		go toRun()
	}
}

func SendPermissionChangeNotificationReturnError(ctx context.Context, actor string, inputs PermissionChangeNotificationInputs) []error {
	blocks := make([]slack.Block, 0)
	var headline string
	if len(inputs.Errors) > 0 {
		headline = fmt.Sprintf(":%s:", config.Config.String("slack.emoji.failed"))
	} else {
		headline = fmt.Sprintf(":%s:", config.Config.String("slack.emoji.succeeded"))
	}
	if actor != "" {
		headline += fmt.Sprintf(" %s", actor)
	}
	if actor != "" && inputs.Summary != "" {
		headline += ":"
	}
	if inputs.Summary != "" {
		headline += fmt.Sprintf(" %s", inputs.Summary)
	}

	if squelch, ok := ctx.Value(permissionChangeSquelchContextKey).(bool); ok && squelch && len(inputs.Errors) == 0 {
		log.Info().Strs("results", inputs.Results).Errs("errors", inputs.Errors).Msgf("SLCK | Squelching permission change notification: `%s`", headline)
		return nil
	}

	log.Info().Strs("results", inputs.Results).Errs("errors", inputs.Errors).Msgf("SLCK | Sending permission change notification: `%s`", headline)

	blocks = append(blocks, chunkLinesToSectionMrkdwnBlocks([]string{headline})...)
	for _, err := range inputs.Errors {
		blocks = append(blocks, chunkLinesToSectionMrkdwnBlocks([]string{fmt.Sprintf("- *Error:* %v", err)})...)
	}
	for _, result := range inputs.Results {
		blocks = append(blocks, chunkLinesToSectionMrkdwnBlocks([]string{fmt.Sprintf("- %s", result)})...)
	}
	errs := make([]error, 0)
	if isEnabled() && config.Config.Bool("slack.behaviors.permissionChanges.enable") && len(blocks) > 0 {
		// We can't send more than 50 blocks in one message, so we do some madness to chunk them and send
		// consecutive messages into a thread. Note that deployment_notification.go does something very similar
		// but we don't abstract between them here because there's a few subtle differences in quantity of
		// channels being handled and how the messages are threaded. An abstraction would just be needlessly
		// complex.

		// First we assemble chunks of 50. This fancy slice syntax and multiple slice assignment does that in
		// chunks of 50 and then we append anything remaining after that.
		var chunks [][]slack.Block
		for len(blocks) > 50 {
			blocks, chunks = blocks[50:], append(chunks, blocks[0:50:50])
		}
		chunks = append(chunks, blocks)

		// The timestamp we'll use to thread is different per channel, so we have to send to each channel
		// sequentially.
		for _, channel := range config.Config.Strings("slack.behaviors.permissionChanges.channels") {
			var timestamp string
			for _, chunk := range chunks {

				// For each chunk, if we already have a timestamp (meaning that we've already successfully sent a
				// message), we use that timestamp to thread the message. If we don't have a timestamp, we send a new
				// message.
				var err error
				if timestamp == "" {
					_, timestamp, _, err = client.SendMessageContext(ctx, channel,
						slack.MsgOptionBlocks(chunk...))
				} else {
					_, _, _, err = client.SendMessageContext(ctx, channel,
						slack.MsgOptionTS(timestamp), slack.MsgOptionBlocks(chunk...))
				}

				// If we got an error, we do some legwork to make debugging easier. Blocks are bytes and should be able
				// to be marshalled to JSON. If we can do that, then we can safely include the blocks in a log message to
				// help unpack whatever went wrong. We generate a random identifier to help correlate the error that gets
				// reported to the log with the blocks that were sent.
				if err != nil {
					if bytes, jsonErr := json.Marshal(chunk); jsonErr != nil {
						err = fmt.Errorf("(also failed to marshal chunk of blocks to JSON: %v) %v", jsonErr, err)
					} else {
						identifier := rand.Int()
						log.Warn().Bytes("blocks", bytes).Int("identifier", identifier).Msgf("failed to send permission change notification chunk, embedding blocks in log (identifier %d)", identifier)
						err = fmt.Errorf("(embedded chunk of blocks in log, seek identifier %d) %v", identifier, err)
					}
					errs = append(errs, err)
				}
			}
		}
	}
	return errs
}
