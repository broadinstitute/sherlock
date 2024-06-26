package slack

import (
	"fmt"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/slack-go/slack"
	"golang.org/x/net/context"
)

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
	blocks = append(blocks, chunkLinesToSectionMrkdwnBlocks([]string{headline})...)
	for _, err := range inputs.Errors {
		blocks = append(blocks, chunkLinesToSectionMrkdwnBlocks([]string{fmt.Sprintf("- *Error:* %v", err)})...)
	}
	for _, result := range inputs.Results {
		blocks = append(blocks, chunkLinesToSectionMrkdwnBlocks([]string{fmt.Sprintf("- %s", result)})...)
	}
	errs := make([]error, 0)
	if isEnabled() && config.Config.Bool("slack.behaviors.permissionChanges.enable") && len(blocks) > 0 {
		for _, channel := range config.Config.Strings("slack.behaviors.permissionChanges.channels") {
			if _, _, _, err := client.SendMessageContext(ctx, channel, slack.MsgOptionBlocks(blocks...)); err != nil {
				errs = append(errs, err)
			}
		}
	}
	return errs
}
