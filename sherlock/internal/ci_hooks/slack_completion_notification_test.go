package ci_hooks

import (
	"context"
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/broadinstitute/sherlock/sherlock/internal/slack"
	"github.com/broadinstitute/sherlock/sherlock/internal/slack/slack_mocks"
	"github.com/stretchr/testify/mock"
	"testing"
)

func Test_dispatcherImpl_DispatchSlackCompletionNotification(t *testing.T) {
	config.LoadTestConfig()
	ctx := context.Background()
	type args struct {
		channel   string
		text      string
		succeeded bool
		icon      *string
	}
	tests := []struct {
		name       string
		args       args
		mockConfig func(c *slack_mocks.MockMockableClient)
		wantErr    bool
	}{
		{
			name: "success",
			args: args{
				channel:   "channel",
				text:      "text",
				succeeded: true,
			},
			mockConfig: func(c *slack_mocks.MockMockableClient) {
				// We can't mock the type of MsgOption sent, unfortunately
				c.EXPECT().SendMessageContext(ctx, "channel", mock.AnythingOfType("slack.MsgOption")).Return("", "", "", nil)
			},
			wantErr: false,
		},
		{
			name: "failure",
			args: args{
				channel:   "channel",
				text:      "text",
				succeeded: false,
			},
			mockConfig: func(c *slack_mocks.MockMockableClient) {
				// We can't mock the type of MsgOption sent, unfortunately
				c.EXPECT().SendMessageContext(ctx, "channel", mock.AnythingOfType("slack.MsgOption")).Return("", "", "", nil)
			},
			wantErr: false,
		},
		{
			name: "error",
			args: args{
				channel:   "channel",
				text:      "text",
				succeeded: true,
			},
			mockConfig: func(c *slack_mocks.MockMockableClient) {
				// We can't mock the type of MsgOption sent, unfortunately
				c.EXPECT().SendMessageContext(ctx, "channel", mock.AnythingOfType("slack.MsgOption")).Return("", "", "", fmt.Errorf("some error"))
			},
			wantErr: true,
		},
		{
			name: "icon",
			args: args{
				channel:   "channel",
				text:      "text",
				succeeded: true,
				icon:      utils.PointerTo(":smiley:"),
			},
			mockConfig: func(c *slack_mocks.MockMockableClient) {
				// We can't mock the type of MsgOption sent, unfortunately
				c.EXPECT().SendMessageContext(ctx, "channel", mock.AnythingOfType("slack.MsgOption"), mock.AnythingOfType("slack.MsgOption")).Return("", "", "", nil)
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			di := &dispatcherImpl{}
			slack.UseMockedClient(t, tt.mockConfig, func() {
				if err := di.DispatchSlackCompletionNotification(ctx, tt.args.channel, tt.args.text, tt.args.succeeded, tt.args.icon); (err != nil) != tt.wantErr {
					t.Errorf("DispatchSlackCompletionNotification() error = %v, wantErr %v", err, tt.wantErr)
				}
			})
		})
	}
}
