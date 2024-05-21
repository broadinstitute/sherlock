package slack

import (
	"context"
	"fmt"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/broadinstitute/sherlock/sherlock/internal/slack/slack_mocks"
	"github.com/rs/zerolog"
	"github.com/slack-go/slack"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetUser(t *testing.T) {
	config.LoadTestConfig()
	zerolog.SetGlobalLevel(zerolog.Disabled)
	ctx := context.Background()
	type args struct {
		email string
	}
	tests := []struct {
		name         string
		args         args
		mockConfig   func(c *slack_mocks.MockMockableClient)
		wantSlackID  string
		wantUsername string
		wantName     string
		wantErr      assert.ErrorAssertionFunc
	}{
		{
			name:         "no email",
			args:         args{email: ""},
			mockConfig:   func(_ *slack_mocks.MockMockableClient) {},
			wantSlackID:  "",
			wantUsername: "",
			wantName:     "",
			wantErr:      assert.NoError,
		},
		{
			name: "normal case",
			args: args{email: "blah"},
			mockConfig: func(c *slack_mocks.MockMockableClient) {
				c.EXPECT().GetUserByEmailContext(ctx, "blah").Return(&slack.User{
					ID:       "id",
					Name:     "username",
					RealName: "real name",
				}, nil)
			},
			wantSlackID:  "id",
			wantUsername: "username",
			wantName:     "real name",
			wantErr:      assert.NoError,
		},
		{
			name: "not found",
			args: args{email: "blah"},
			mockConfig: func(c *slack_mocks.MockMockableClient) {
				c.EXPECT().GetUserByEmailContext(ctx, "blah").Return(nil, slack.SlackErrorResponse{
					Err: "users_not_found",
				})
			},
			wantSlackID:  "",
			wantUsername: "",
			wantName:     "",
			wantErr:      assert.Error,
		},
		{
			name: "missing fields in response",
			args: args{email: "blah"},
			mockConfig: func(c *slack_mocks.MockMockableClient) {
				c.EXPECT().GetUserByEmailContext(ctx, "blah").Return(&slack.User{}, nil)
			},
			wantSlackID:  "",
			wantUsername: "",
			wantName:     "",
			wantErr:      assert.NoError,
		},
		{
			name: "nil response",
			args: args{email: "blah"},
			mockConfig: func(c *slack_mocks.MockMockableClient) {
				c.EXPECT().GetUserByEmailContext(ctx, "blah").Return(nil, nil)
			},
			wantSlackID:  "",
			wantUsername: "",
			wantName:     "",
			wantErr:      assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UseMockedClient(t, tt.mockConfig, func() {
				gotSlackID, gotUsername, gotName, err := GetUser(ctx, tt.args.email)
				if !tt.wantErr(t, err, fmt.Sprintf("GetUser(%v, %v)", ctx, tt.args.email)) {
					return
				}
				assert.Equalf(t, tt.wantSlackID, gotSlackID, "GetUser(%v, %v)", ctx, tt.args.email)
				assert.Equalf(t, tt.wantUsername, gotUsername, "GetUser(%v, %v)", ctx, tt.args.email)
				assert.Equalf(t, tt.wantName, gotName, "GetUser(%v, %v)", ctx, tt.args.email)
			})
		})
	}
}
