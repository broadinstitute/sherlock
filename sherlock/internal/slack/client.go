package slack

import (
	"context"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/broadinstitute/sherlock/sherlock/internal/slack/slack_mocks"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/socketmode"
	"testing"
)

// If you're using GoLand you can use the gutter annotation to the left of the line below to regenerate the mock.
// Otherwise, `make generate-mocks` from the root of the repo. In either case, you'll need to `brew install mockery`.
//
//go:generate mockery
type mockableClient interface {
	SendMessageContext(ctx context.Context, channelID string, options ...slack.MsgOption) (_channel string, _timestamp string, _text string, err error)
	AddReaction(name string, item slack.ItemRef) error
}

var (
	// client is what functions in this package should use whenever possible.
	client mockableClient

	// rawClient is used to pass the socketmode.Client between Init and Start.
	// During development, you may use it instead of client, since it has full
	// access to Slack's entire API surface. Once you know what methods you
	// need, you can add them to mockableClient and switch your new code from
	// rawClient to client.
	rawClient *socketmode.Client
)

// isEnabled is true when we have a client and either slack.enable is true or slack.enableMocked is true and we have
// no rawClient (indicating that whatever client is, it isn't an actual connection to Slack).
// This doesn't check if any individual behaviors are enabled, just the top-level of if Slack code can run at all.
// Not all functions need to worry about calling this, really just exported functions so that the caller doesn't
// need to do the check on their end.
func isEnabled() bool {
	return client != nil && (config.Config.Bool("slack.enable") || (config.Config.Bool("slack.enableMocked") && rawClient == nil))
}

func UseMockedClient(t *testing.T, config func(c *slack_mocks.MockMockableClient), callback func()) {
	c := slack_mocks.NewMockMockableClient(t)
	config(c)
	temp := client
	client = c
	callback()
	c.AssertExpectations(t)
	client = temp
}
