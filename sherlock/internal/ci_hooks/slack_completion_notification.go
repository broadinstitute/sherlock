package ci_hooks

import (
	"context"

	"github.com/broadinstitute/sherlock/sherlock/internal/clients/slack"
)

// DispatchSlackCompletionNotification is pretty much a re-export of slack.SendMessageReturnError.
// It encapsulates the call to the slack package so we don't need to reach into it when mocking
// callers of this package.
func (_ *dispatcherImpl) DispatchSlackCompletionNotification(ctx context.Context, channel string, text string, succeeded bool, icon *string) error {
	var attachment slack.Attachment
	if succeeded {
		attachment = slack.GreenBlock{Text: text}
	} else {
		attachment = slack.RedBlock{Text: text}
	}
	return slack.SendMessageReturnError(ctx, channel, "", icon, attachment)
}
