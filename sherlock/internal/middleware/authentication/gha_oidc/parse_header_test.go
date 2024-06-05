package gha_oidc

import (
	"fmt"
	"github.com/broadinstitute/sherlock/sherlock/internal/clients/slack"
	"github.com/broadinstitute/sherlock/sherlock/internal/clients/slack/slack_mocks"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/broadinstitute/sherlock/sherlock/internal/middleware/authentication/gha_oidc/gha_oidc_claims"
	"github.com/broadinstitute/sherlock/sherlock/internal/middleware/authentication/gha_oidc/gha_oidc_mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestParseHeader(t *testing.T) {
	gin.SetMode(gin.TestMode)
	config.LoadTestConfig()
	engine := gin.New()
	var gotClaims *gha_oidc_claims.Claims
	var gotErr error
	engine.Use(func(ctx *gin.Context) {
		gotClaims, gotErr = ParseHeader(ctx)
		ctx.JSON(http.StatusOK, struct{}{})
	})
	tests := []struct {
		name            string
		headerContents  string
		mockConfig      func(v *gha_oidc_mocks.MockMockableVerifier)
		slackMockConfig func(c *slack_mocks.MockMockableClient)
		wantClaims      *gha_oidc_claims.Claims
		wantErr         assert.ErrorAssertionFunc
	}{
		{
			name:           "empty",
			headerContents: "",
			mockConfig:     func(v *gha_oidc_mocks.MockMockableVerifier) {},
			wantClaims:     nil,
			wantErr:        assert.NoError,
		},
		{
			name:           "no verifier",
			headerContents: "some header",
			mockConfig:     nil,
			wantClaims:     nil,
			wantErr:        assert.NoError,
		},
		{
			name:           "fails validation",
			headerContents: "some header",
			mockConfig: func(v *gha_oidc_mocks.MockMockableVerifier) {
				v.EXPECT().VerifyAndParseClaims(mock.Anything, "some header").Return(
					gha_oidc_claims.Claims{}, fmt.Errorf("some error"))
			},
			wantClaims: nil,
			wantErr:    assert.Error,
		},
		{
			name:           "repository owner not accepted",
			headerContents: "some header",
			mockConfig: func(v *gha_oidc_mocks.MockMockableVerifier) {
				v.EXPECT().VerifyAndParseClaims(mock.Anything, "some header").Return(
					gha_oidc_claims.Claims{
						RepositoryOwner: "some unknown owner",
					}, nil)
			},
			slackMockConfig: func(c *slack_mocks.MockMockableClient) {
				c.EXPECT().SendMessageContext(mock.Anything, mock.AnythingOfType("string"), mock.Anything, mock.Anything).
					Return("", "", "", nil)
			},
			wantClaims: nil,
			wantErr:    assert.NoError,
		},
		{
			name:           "repository owner not accepted",
			headerContents: "some header",
			mockConfig: func(v *gha_oidc_mocks.MockMockableVerifier) {
				v.EXPECT().VerifyAndParseClaims(mock.Anything, "some header").Return(
					gha_oidc_claims.Claims{
						RepositoryOwner: config.Config.Strings("auth.githubActionsOIDC.allowedOrganizations")[0],
						Repository:      "owner/name",
					}, nil)
			},
			wantClaims: &gha_oidc_claims.Claims{
				RepositoryOwner: config.Config.Strings("auth.githubActionsOIDC.allowedOrganizations")[0],
				Repository:      "owner/name",
			},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			slack.UseMockedClient(t, tt.slackMockConfig, func() {
				UseMockedVerifier(t, tt.mockConfig, func() {
					request, err := http.NewRequest(http.MethodGet, "/", nil)
					assert.NoError(t, err)
					request.Header.Set(Header, tt.headerContents)
					recorder := httptest.NewRecorder()
					engine.ServeHTTP(recorder, request)
					if !tt.wantErr(t, gotErr, fmt.Sprintf("verifier.Verify(ctx, %s)", tt.headerContents)) {
						return
					}
					assert.Equal(t, tt.wantClaims, gotClaims)
				})
			})
		})
	}
}
