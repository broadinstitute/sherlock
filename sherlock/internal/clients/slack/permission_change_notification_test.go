package slack

import (
	"context"
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/clients/slack/slack_mocks"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestSetContextToSquelchPermissionChangeNotifications(t *testing.T) {
	config.LoadTestConfig()
	zerolog.SetGlobalLevel(zerolog.Disabled)
	ctx := context.Background()
	ctx = SetContextToSquelchPermissionChangeNotifications(ctx)
	if ctx.Value(permissionChangeSquelchContextKey) == nil {
		t.Errorf("SetContextToSquelchPermissionChangeNotifications() did not set the correct context key")
	}

	t.Run("shouldn't contact Slack if no errors", func(t *testing.T) {
		UseMockedClient(t, func(c *slack_mocks.MockMockableClient) {
			// Shouldn't touch Slack
		}, func() {
			SendPermissionChangeNotificationReturnError(ctx, "test", PermissionChangeNotificationInputs{
				Summary: "summary",
				Results: []string{"result"},
				Errors:  []error{},
			})
		})
	})

	t.Run("should contact Slack if errors", func(t *testing.T) {
		UseMockedClient(t, func(c *slack_mocks.MockMockableClient) {
			c.EXPECT().SendMessageContext(ctx, "#notification-channel", mock.Anything).Return("", "", "", nil).Once()
			c.EXPECT().SendMessageContext(ctx, "#permission-change-channel", mock.Anything).Return("", "", "", nil).Once()
		}, func() {
			SendPermissionChangeNotificationReturnError(ctx, "test", PermissionChangeNotificationInputs{
				Summary: "summary",
				Results: []string{"result"},
				Errors:  []error{fmt.Errorf("error")},
			})
		})
	})
}

func TestSendPermissionChangeNotification(t *testing.T) {
	config.LoadTestConfig()
	zerolog.SetGlobalLevel(zerolog.Disabled)
	ctx := context.Background()

	var fortyStrings []string
	for i := 0; i < 40; i++ {
		fortyStrings = append(fortyStrings, "string")
	}

	type args struct {
		actor  string
		inputs PermissionChangeNotificationInputs
	}
	tests := []struct {
		name       string
		args       args
		mockConfig func(c *slack_mocks.MockMockableClient)
	}{
		{
			name: "normal",
			args: args{
				actor: "test",
				inputs: PermissionChangeNotificationInputs{
					Summary: "summary",
					Results: []string{"result"},
					Errors:  []error{fmt.Errorf("error")},
				},
			},
			mockConfig: func(c *slack_mocks.MockMockableClient) {
				c.EXPECT().SendMessageContext(ctx, "#notification-channel", mock.Anything).Return("", "", "", nil).Once()
				c.EXPECT().SendMessageContext(ctx, "#permission-change-channel", mock.Anything).Return("", "", "", nil).Once()
			},
		},
		{
			name: "normal with no errors",
			args: args{
				actor: "test",
				inputs: PermissionChangeNotificationInputs{
					Summary: "summary",
					Results: []string{"result"},
				},
			},
			mockConfig: func(c *slack_mocks.MockMockableClient) {
				c.EXPECT().SendMessageContext(ctx, "#notification-channel", mock.Anything).Return("", "", "", nil).Once()
				c.EXPECT().SendMessageContext(ctx, "#permission-change-channel", mock.Anything).Return("", "", "", nil).Once()
			},
		},
		{
			name: "normal but chunked",
			args: args{
				actor: "test",
				inputs: PermissionChangeNotificationInputs{
					Summary: "summary",
					Results: fortyStrings,
					Errors:  utils.Map(fortyStrings, func(_ string) error { return fmt.Errorf("error") }),
				},
			},
			mockConfig: func(c *slack_mocks.MockMockableClient) {
				// Two messages to each channel, the second one to each is threaded
				c.EXPECT().SendMessageContext(ctx, "#notification-channel", mock.Anything).Return("", "123", "", nil).Once()
				c.EXPECT().SendMessageContext(ctx, "#notification-channel", mock.Anything, mock.Anything).Return("", "", "", nil).Once()

				c.EXPECT().SendMessageContext(ctx, "#permission-change-channel", mock.Anything).Return("", "456", "", nil).Once()
				c.EXPECT().SendMessageContext(ctx, "#permission-change-channel", mock.Anything, mock.Anything).Return("", "", "", nil).Once()
			},
		},
		{
			name: "report one error",
			args: args{
				actor: "test",
				inputs: PermissionChangeNotificationInputs{
					Summary: "summary",
					Results: []string{"result"},
				},
			},
			mockConfig: func(c *slack_mocks.MockMockableClient) {
				c.EXPECT().SendMessageContext(ctx, "#notification-channel", mock.Anything).Return("", "", "", fmt.Errorf("error")).Once()
				c.EXPECT().SendMessageContext(ctx, "#permission-change-channel", mock.Anything).Return("", "", "", nil).Once()
				// error should be reported
				c.EXPECT().SendMessageContext(ctx, "#error-channel", mock.Anything, mock.Anything).Return("", "", "", nil).Once()
				c.EXPECT().SendMessageContext(ctx, "#notification-channel", mock.Anything, mock.Anything).Return("", "", "", nil).Once()
			},
		},
		{
			name: "report two errors",
			args: args{
				actor: "test",
				inputs: PermissionChangeNotificationInputs{
					Summary: "summary",
					Results: []string{"result"},
				},
			},
			mockConfig: func(c *slack_mocks.MockMockableClient) {
				c.EXPECT().SendMessageContext(ctx, "#notification-channel", mock.Anything).Return("", "", "", fmt.Errorf("error 1")).Once()
				c.EXPECT().SendMessageContext(ctx, "#permission-change-channel", mock.Anything).Return("", "", "", fmt.Errorf("error 2")).Once()
				// error should be reported
				c.EXPECT().SendMessageContext(ctx, "#error-channel", mock.Anything, mock.Anything).Return("", "", "", nil).Once()
				c.EXPECT().SendMessageContext(ctx, "#notification-channel", mock.Anything, mock.Anything).Return("", "", "", nil).Once()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UseMockedClient(t, tt.mockConfig, func() {
				SendPermissionChangeNotification(ctx, tt.args.actor, tt.args.inputs)
			})
		})
	}
}

func TestSendPermissionChangeNotificationReturnError(t *testing.T) {
	config.LoadTestConfig()
	zerolog.SetGlobalLevel(zerolog.Disabled)
	ctx := context.Background()
	type args struct {
		actor  string
		inputs PermissionChangeNotificationInputs
	}
	tests := []struct {
		name       string
		args       args
		mockConfig func(c *slack_mocks.MockMockableClient)
		wantErrs   int
	}{
		{
			name: "normal",
			args: args{
				actor: "test",
				inputs: PermissionChangeNotificationInputs{
					Summary: "summary",
					Results: []string{"result"},
					Errors:  []error{fmt.Errorf("error")},
				},
			},
			mockConfig: func(c *slack_mocks.MockMockableClient) {
				c.EXPECT().SendMessageContext(ctx, "#notification-channel", mock.Anything).Return("", "", "", nil).Once()
				c.EXPECT().SendMessageContext(ctx, "#permission-change-channel", mock.Anything).Return("", "", "", nil).Once()
			},
			wantErrs: 0,
		},
		{
			name: "normal with no errors",
			args: args{
				actor: "test",
				inputs: PermissionChangeNotificationInputs{
					Summary: "summary",
					Results: []string{"result"},
				},
			},
			mockConfig: func(c *slack_mocks.MockMockableClient) {
				c.EXPECT().SendMessageContext(ctx, "#notification-channel", mock.Anything).Return("", "", "", nil).Once()
				c.EXPECT().SendMessageContext(ctx, "#permission-change-channel", mock.Anything).Return("", "", "", nil).Once()
			},
			wantErrs: 0,
		},
		{
			name: "return one error",
			args: args{
				actor: "test",
				inputs: PermissionChangeNotificationInputs{
					Summary: "summary",
					Results: []string{"result"},
				},
			},
			mockConfig: func(c *slack_mocks.MockMockableClient) {
				c.EXPECT().SendMessageContext(ctx, "#notification-channel", mock.Anything).Return("", "", "", fmt.Errorf("error")).Once()
				c.EXPECT().SendMessageContext(ctx, "#permission-change-channel", mock.Anything).Return("", "", "", nil).Once()
			},
			wantErrs: 1,
		},
		{
			name: "return two errors",
			args: args{
				actor: "test",
				inputs: PermissionChangeNotificationInputs{
					Summary: "summary",
					Results: []string{"result"},
				},
			},
			mockConfig: func(c *slack_mocks.MockMockableClient) {
				c.EXPECT().SendMessageContext(ctx, "#notification-channel", mock.Anything).Return("", "", "", fmt.Errorf("error 1")).Once()
				c.EXPECT().SendMessageContext(ctx, "#permission-change-channel", mock.Anything).Return("", "", "", fmt.Errorf("error 2")).Once()
			},
			wantErrs: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UseMockedClient(t, tt.mockConfig, func() {
				gotErrs := SendPermissionChangeNotificationReturnError(ctx, tt.args.actor, tt.args.inputs)
				if len(gotErrs) != tt.wantErrs {
					t.Errorf("SendPermissionChangeNotificationReturnError() gotErrs = %v, want %v", gotErrs, tt.wantErrs)
				}
			})
		})
	}
}
