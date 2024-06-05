package slack

import (
	"context"
	"fmt"
	"github.com/broadinstitute/sherlock/sherlock/internal/clients/slack/slack_mocks"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestReportError(t *testing.T) {
	config.LoadTestConfig()
	zerolog.SetGlobalLevel(zerolog.Disabled)
	ctx := context.Background()
	type args struct {
		description string
		errs        []error
	}
	tests := []struct {
		name       string
		args       args
		mockConfig func(c *slack_mocks.MockMockableClient)
	}{
		{
			name: "normal case",
			args: args{errs: []error{fmt.Errorf("some error")}},
			mockConfig: func(c *slack_mocks.MockMockableClient) {
				c.On("SendMessageContext", ctx, "channel 1",
					mock.AnythingOfType("slack.MsgOption"),
					mock.AnythingOfType("slack.MsgOption")).Return("", "", "", nil)
				c.On("SendMessageContext", ctx, "channel 2",
					mock.AnythingOfType("slack.MsgOption"),
					mock.AnythingOfType("slack.MsgOption")).Return("", "", "", nil)
			},
		},
		{
			name: "sends no errors",
			args: args{errs: []error{}},
			mockConfig: func(c *slack_mocks.MockMockableClient) {
				c.On("SendMessageContext", ctx, "channel 1",
					mock.AnythingOfType("slack.MsgOption")).Return("", "", "", nil)
				c.On("SendMessageContext", ctx, "channel 2",
					mock.AnythingOfType("slack.MsgOption")).Return("", "", "", nil)
			},
		},
		{
			name: "sends multiple errors",
			args: args{errs: []error{fmt.Errorf("some error"), fmt.Errorf("some second error")}},
			mockConfig: func(c *slack_mocks.MockMockableClient) {
				c.On("SendMessageContext", ctx, "channel 1",
					mock.AnythingOfType("slack.MsgOption"),
					mock.AnythingOfType("slack.MsgOption")).Return("", "", "", nil)
				c.On("SendMessageContext", ctx, "channel 2",
					mock.AnythingOfType("slack.MsgOption"),
					mock.AnythingOfType("slack.MsgOption")).Return("", "", "", nil)
				c.On("SendMessageContext", ctx, "channel 1",
					mock.AnythingOfType("slack.MsgOption"),
					mock.AnythingOfType("slack.MsgOption")).Return("", "", "", nil)
				c.On("SendMessageContext", ctx, "channel 2",
					mock.AnythingOfType("slack.MsgOption"),
					mock.AnythingOfType("slack.MsgOption")).Return("", "", "", nil)
			},
		},
		{
			name: "sending on channel 1 errors",
			args: args{errs: []error{fmt.Errorf("some error")}},
			mockConfig: func(c *slack_mocks.MockMockableClient) {
				c.On("SendMessageContext", ctx, "channel 1",
					mock.AnythingOfType("slack.MsgOption"),
					mock.AnythingOfType("slack.MsgOption")).Return("", "", "", fmt.Errorf("some send error"))
				c.On("SendMessageContext", ctx, "channel 2",
					mock.AnythingOfType("slack.MsgOption"),
					mock.AnythingOfType("slack.MsgOption")).Return("", "", "", nil)
			},
		},
		{
			name: "sending on channel 2 errors",
			args: args{errs: []error{fmt.Errorf("some error")}},
			mockConfig: func(c *slack_mocks.MockMockableClient) {
				c.On("SendMessageContext", ctx, "channel 1",
					mock.AnythingOfType("slack.MsgOption"),
					mock.AnythingOfType("slack.MsgOption")).Return("", "", "", nil)
				c.On("SendMessageContext", ctx, "channel 2",
					mock.AnythingOfType("slack.MsgOption"),
					mock.AnythingOfType("slack.MsgOption")).Return("", "", "", fmt.Errorf("some send error"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UseMockedClient(t, tt.mockConfig, func() {
				ReportError(ctx, tt.args.description, tt.args.errs...)
			})
		})
	}
}
