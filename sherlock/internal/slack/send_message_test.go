package slack

import (
	"context"
	"fmt"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/broadinstitute/sherlock/sherlock/internal/slack/slack_mocks"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestSendMessage(t *testing.T) {
	config.LoadTestConfig()
	zerolog.SetGlobalLevel(zerolog.Disabled)
	ctx := context.Background()
	type args struct {
		channel     string
		text        string
		attachments []Attachment
	}
	tests := []struct {
		name       string
		args       args
		mockConfig func(c *slack_mocks.MockMockableClient)
	}{
		{
			name:       "doesn't do anything when empty",
			args:       args{channel: "foo"},
			mockConfig: func(client *slack_mocks.MockMockableClient) {},
		},
		{
			name: "sends text when provided",
			args: args{channel: "foo", text: "text"},
			mockConfig: func(c *slack_mocks.MockMockableClient) {
				c.On("SendMessageContext", ctx, "foo",
					mock.AnythingOfType("slack.MsgOption")).Return("", "", "", nil)
			},
		},
		{
			name: "sends attachments when provided",
			args: args{channel: "foo", attachments: []Attachment{
				GreenBlock{Text: "blah"}, RedBlock{"bleh"},
			}},
			mockConfig: func(c *slack_mocks.MockMockableClient) {
				c.On("SendMessageContext", ctx, "foo",
					mock.AnythingOfType("slack.MsgOption")).Return("", "", "", nil)
			},
		},
		{
			name: "sends both when provided",
			args: args{channel: "foo", text: "text", attachments: []Attachment{
				GreenBlock{Text: "blah"}, RedBlock{"bleh"},
			}},
			mockConfig: func(c *slack_mocks.MockMockableClient) {
				c.On("SendMessageContext", ctx, "foo",
					mock.AnythingOfType("slack.MsgOption"),
					mock.AnythingOfType("slack.MsgOption")).Return("", "", "", nil)
			},
		},
		{
			name: "swallows errors",
			args: args{channel: "foo", text: "text", attachments: []Attachment{
				GreenBlock{Text: "blah"}, RedBlock{"bleh"},
			}},
			mockConfig: func(c *slack_mocks.MockMockableClient) {
				c.On("SendMessageContext", ctx, "foo",
					mock.AnythingOfType("slack.MsgOption"),
					mock.AnythingOfType("slack.MsgOption")).Return("", "", "", fmt.Errorf("some error"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UseMockedClient(t, tt.mockConfig, func() {
				SendMessage(ctx, tt.args.channel, tt.args.text, tt.args.attachments...)
			})
		})
	}
}
