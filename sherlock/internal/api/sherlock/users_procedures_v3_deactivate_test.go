package sherlock

import (
	"fmt"
	"github.com/broadinstitute/sherlock/sherlock/internal/clients/google_workspace/google_workspace_mocks"
	"github.com/broadinstitute/sherlock/sherlock/internal/clients/slack"
	"github.com/broadinstitute/sherlock/sherlock/internal/clients/slack/slack_mocks"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"testing"
)

func (s *handlerSuite) TestUsersProceduresV3Deactivate_badBody() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("POST", "/api/users/procedures/v3/deactivate", gin.H{
			"userEmails": 3,
		}),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
	s.Contains(got.Message, "userEmails")
}

func (s *handlerSuite) TestUsersProceduresV3Deactivate_notSuperAdmin() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewSuitableRequest("POST", "/api/users/procedures/v3/deactivate", UserV3DeactivateRequest{
			UserEmails: []string{"user@broadinstitute.org"},
		}),
		&got)
	s.Equal(http.StatusForbidden, code)
	s.Equal(errors.Forbidden, got.Type)
}

func (s *handlerSuite) TestUsersProceduresV3Deactivate_emptyEmails() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewSuperAdminRequest("POST", "/api/users/procedures/v3/deactivate", UserV3DeactivateRequest{
			UserEmails: []string{},
		}),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
	s.Contains(got.Message, "no user emails provided")
}

func (s *handlerSuite) TestUsersProceduresV3Deactivate_emptyEmailString() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewSuperAdminRequest("POST", "/api/users/procedures/v3/deactivate", UserV3DeactivateRequest{
			UserEmails: []string{""},
		}),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
	s.Contains(got.Message, "empty email provided")
}

func (s *handlerSuite) TestUsersProceduresV3Deactivate_deactivateSelf() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewSuperAdminRequest("POST", "/api/users/procedures/v3/deactivate", UserV3DeactivateRequest{
			UserEmails: []string{s.TestData.User_SuperAdmin().Email},
		}),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
	s.Contains(got.Message, "you cannot deactivate yourself")
}

func (s *handlerSuite) TestUsersProceduresV3Deactivate_emptyGoogleWorkspaceDomain() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewSuperAdminRequest("POST", "/api/users/procedures/v3/deactivate", UserV3DeactivateRequest{
			UserEmails: []string{s.TestData.User_Suitable().Email},
			SuspendEmailHandlesAcrossGoogleWorkspaceDomains: []string{""},
		}),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
	s.Contains(got.Message, "empty Google Workspace domain provided")
}

func (s *handlerSuite) TestUsersProceduresV3Deactivate() {
	testDataDeactivatedUserTime := s.TestData.User_Deactivated().DeactivatedAt

	var got UserV3DeactivateResponse
	var code int
	slack.UseMockedClient(s.T(), func(c *slack_mocks.MockMockableClient) {
		c.EXPECT().SendMessageContext(mock.Anything, "#notification-channel", mock.Anything).Return("", "", "", nil).Once()
		c.EXPECT().SendMessageContext(mock.Anything, "#permission-change-channel", mock.Anything).Return("", "", "", nil).Once()
	}, func() {
		code = s.HandleRequest(
			s.NewSuperAdminRequest("POST", "/api/users/procedures/v3/deactivate", UserV3DeactivateRequest{
				UserEmails: []string{
					s.TestData.User_Suitable().Email,
					s.TestData.User_NonSuitable().Email,
					s.TestData.User_Deactivated().Email,
					"does-not-exist@broadinstitute.org",
					"another-domain@example.org",
				},
			}),
			&got)
	})
	s.Equal(http.StatusOK, code)
	s.ElementsMatch([]string{
		s.TestData.User_Suitable().Email,
		s.TestData.User_NonSuitable().Email,
	}, got.NewlyDeactivatedEmails)
	s.ElementsMatch([]string{
		s.TestData.User_Deactivated().Email,
	}, got.AlreadyDeactivatedEmails)
	s.ElementsMatch([]string{
		"does-not-exist@broadinstitute.org",
		"another-domain@example.org",
	}, got.NotFoundEmails)

	s.Run("suitable user deactivated", func() {
		var user models.User
		s.NoError(s.DB.First(&user, s.TestData.User_Suitable().ID).Error)
		s.NotNil(user.DeactivatedAt)
	})
	s.Run("non-suitable user deactivated", func() {
		var user models.User
		s.NoError(s.DB.First(&user, s.TestData.User_NonSuitable().ID).Error)
		s.NotNil(user.DeactivatedAt)
	})
	s.Run("deactivated user not changed", func() {
		var user models.User
		s.NoError(s.DB.First(&user, s.TestData.User_Deactivated().ID).Error)
		s.Equal(testDataDeactivatedUserTime, user.DeactivatedAt)
	})
}

func Test_processGoogleWorkspaceSuspensions(t *testing.T) {
	config.LoadTestConfig()
	type args struct {
		actor      string
		domain     string
		emails     []string
		homeDomain string
	}
	tests := []struct {
		name         string
		clientConfig func(c *google_workspace_mocks.MockWorkspaceClient)
		args         args
	}{
		{
			name: "no emails",
			args: args{
				actor:      "actor",
				domain:     "domain",
				emails:     []string{},
				homeDomain: "",
			},
		},
		{
			name: "suspend",
			args: args{
				actor:      "actor",
				domain:     "target-domain",
				emails:     []string{"email@home-domain"},
				homeDomain: "home-domain",
			},
			clientConfig: func(c *google_workspace_mocks.MockWorkspaceClient) {
				c.EXPECT().SuspendUser(mock.Anything, "email@target-domain").Return(nil).Once()
			},
		},
		{
			name: "suspend error",
			args: args{
				actor:      "actor",
				domain:     "target-domain",
				emails:     []string{"email@home-domain"},
				homeDomain: "home-domain",
			},
			clientConfig: func(c *google_workspace_mocks.MockWorkspaceClient) {
				c.EXPECT().SuspendUser(mock.Anything, "email@target-domain").Return(assert.AnError).Once()
			},
		},
		{
			name: "domain miss",
			args: args{
				actor:      "actor",
				domain:     "target-domain",
				emails:     []string{"email@home-domain"},
				homeDomain: "other-domain",
			},
		},
		{
			name: "multiple",
			args: args{
				actor:      "actor",
				domain:     "target-domain",
				emails:     []string{"email-1@home-domain", "email-2@home-domain", "email-3@other-domain"},
				homeDomain: "home-domain",
			},
			clientConfig: func(c *google_workspace_mocks.MockWorkspaceClient) {
				c.EXPECT().SuspendUser(mock.Anything, "email-1@target-domain").Return(nil).Once()
				c.EXPECT().SuspendUser(mock.Anything, "email-2@target-domain").Return(fmt.Errorf("not found")).Once()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := google_workspace_mocks.NewMockWorkspaceClient(t)
			if tt.clientConfig != nil {
				tt.clientConfig(client)
			}
			slack.UseMockedClient(t, func(c *slack_mocks.MockMockableClient) {
				c.EXPECT().SendMessageContext(mock.Anything, "#notification-channel", mock.Anything).Return("", "", "", nil).Once()
				c.EXPECT().SendMessageContext(mock.Anything, "#permission-change-channel", mock.Anything).Return("", "", "", nil).Once()
			}, func() {
				processGoogleWorkspaceSuspensions(tt.args.actor, tt.args.domain, client, tt.args.emails, tt.args.homeDomain)
			})
		})
	}
}
