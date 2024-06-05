package slack

import (
	"fmt"
	"github.com/broadinstitute/sherlock/sherlock/internal/clients/slack/slack_mocks"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/rs/zerolog"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_logOnlyHandlers(t *testing.T) {
	config.LoadTestConfig()
	assert.NotPanics(t, func() {
		handleConnecting(nil, nil)
	})
	assert.NotPanics(t, func() {
		handleConnectionError(nil, nil)
	})
	assert.NotPanics(t, func() {
		handleConnected(nil, nil)
	})
	assert.NotPanics(t, func() {
		handleHello(nil, nil)
	})
}

func Test_handleAppMentionEvent(t *testing.T) {
	config.LoadTestConfig()
	zerolog.SetGlobalLevel(zerolog.Disabled)
	type args struct {
		event *slackevents.AppMentionEvent
	}
	tests := []struct {
		name       string
		args       args
		mockConfig func(c *slack_mocks.MockMockableClient)
	}{
		{
			name: "normal case",
			args: args{event: &slackevents.AppMentionEvent{
				Channel:        "channel",
				EventTimeStamp: "123",
			}},
			mockConfig: func(c *slack_mocks.MockMockableClient) {
				c.On("AddReaction", config.Config.String("slack.behaviors.reactToMentionsWithEmoji.emoji"),
					slack.ItemRef{
						Channel:   "channel",
						Timestamp: "123",
					}).Return(nil)
			},
		},
		{
			name: "error",
			args: args{event: &slackevents.AppMentionEvent{
				Channel:        "channel",
				EventTimeStamp: "123",
			}},
			mockConfig: func(c *slack_mocks.MockMockableClient) {
				c.On("AddReaction", config.Config.String("slack.behaviors.reactToMentionsWithEmoji.emoji"),
					slack.ItemRef{
						Channel:   "channel",
						Timestamp: "123",
					}).Return(fmt.Errorf("some error"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UseMockedClient(t, tt.mockConfig, func() {
				handleAppMentionEvent(client, tt.args.event)
			})
		})
	}
}
