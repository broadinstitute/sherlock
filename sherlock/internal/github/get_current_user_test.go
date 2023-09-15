package github

import (
	"context"
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/testutils"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/google/go-github/v50/github"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetCurrentUser(t *testing.T) {
	config.LoadTestConfig()
	ctx := context.Background()
	type args struct {
		authTokenOverride string
	}
	tests := []struct {
		name         string
		args         args
		mockConfig   func(c *MockClient)
		wantGithubID string
		wantUsername string
		wantName     string
		wantErr      assert.ErrorAssertionFunc
	}{
		{
			name: "normal case",
			args: args{authTokenOverride: "some token"},
			mockConfig: func(c *MockClient) {
				c.Users.EXPECT().Get(ctx, "").Return(&github.User{
					ID:    testutils.PointerTo[int64](123),
					Login: testutils.PointerTo("username"),
					Name:  testutils.PointerTo("name"),
				}, nil, nil)
			},
			wantGithubID: "123",
			wantUsername: "username",
			wantName:     "name",
			wantErr:      assert.NoError,
		},
		{
			name: "fine without an auth token",
			args: args{},
			mockConfig: func(c *MockClient) {
				c.Users.EXPECT().Get(ctx, "").Return(&github.User{
					ID:    testutils.PointerTo[int64](123),
					Login: testutils.PointerTo("username"),
					Name:  testutils.PointerTo("name"),
				}, nil, nil)
			},
			wantGithubID: "123",
			wantUsername: "username",
			wantName:     "name",
			wantErr:      assert.NoError,
		},
		{
			name: "handles empty fields",
			args: args{authTokenOverride: "some token"},
			mockConfig: func(c *MockClient) {
				c.Users.EXPECT().Get(ctx, "").Return(&github.User{}, nil, nil)
			},
			wantGithubID: "",
			wantUsername: "",
			wantName:     "",
			wantErr:      assert.NoError,
		},
		{
			name: "handles nil",
			args: args{authTokenOverride: "some token"},
			mockConfig: func(c *MockClient) {
				c.Users.EXPECT().Get(ctx, "").Return(nil, nil, nil)
			},
			wantGithubID: "",
			wantUsername: "",
			wantName:     "",
			wantErr:      assert.NoError,
		},
		{
			name: "handles errors",
			args: args{authTokenOverride: "some token"},
			mockConfig: func(c *MockClient) {
				c.Users.EXPECT().Get(ctx, "").Return(nil, nil, fmt.Errorf("error"))
			},
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UseMockedClient(t, tt.mockConfig, func() {
				var variadicAuthTokenOverride []string
				if tt.args.authTokenOverride != "" {
					variadicAuthTokenOverride = []string{tt.args.authTokenOverride}
				}
				gotGithubID, gotUsername, gotName, err := GetCurrentUser(ctx, variadicAuthTokenOverride...)
				if !tt.wantErr(t, err, fmt.Sprintf("GetCurrentUser(%v, %v)", ctx, tt.args.authTokenOverride)) {
					return
				}
				assert.Equalf(t, tt.wantGithubID, gotGithubID, "GetCurrentUser(%v, %v)", ctx, tt.args.authTokenOverride)
				assert.Equalf(t, tt.wantUsername, gotUsername, "GetCurrentUser(%v, %v)", ctx, tt.args.authTokenOverride)
				assert.Equalf(t, tt.wantName, gotName, "GetCurrentUser(%v, %v)", ctx, tt.args.authTokenOverride)
			})
		})
	}
}
