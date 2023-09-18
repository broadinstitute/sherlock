package sherlock

import (
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/authentication/test_users"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/github"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/broadinstitute/sherlock/sherlock/internal/slack"
	"github.com/broadinstitute/sherlock/sherlock/internal/slack/slack_mocks"
	"github.com/gin-gonic/gin"
	github2 "github.com/google/go-github/v50/github"
	"github.com/rs/zerolog"
	slack2 "github.com/slack-go/slack"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"testing"
)

func (s *handlerSuite) TestUserV3Upsert_error() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("PUT", "/api/users/v3", gin.H{"name": true}),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
}

func (s *handlerSuite) TestUserV3Upsert_empty() {
	s.Run("{}", func() {
		var got UserV3
		code := s.HandleRequest(
			s.NewRequest("PUT", "/api/users/v3", UserV3Upsert{}),
			&got)
		s.Equal(http.StatusOK, code)
		s.Equal(test_users.SuitableTestUserEmail, got.Email)
	})
	s.Run("no body", func() {
		var got UserV3
		code := s.HandleRequest(
			s.NewRequest("PUT", "/api/users/v3", nil),
			&got)
		s.Equal(http.StatusOK, code)
		s.Equal(test_users.SuitableTestUserEmail, got.Email)
	})
}

func (s *handlerSuite) TestUserV3Upsert_name_minimal() {

	var got UserV3
	code := s.HandleRequest(
		s.NewRequest("PUT", "/api/users/v3", UserV3Upsert{
			userDirectlyEditableFields: userDirectlyEditableFields{
				Name: utils.PointerTo("a name"),
			},
		}),
		&got)
	s.Equal(http.StatusCreated, code)
	s.Equal(test_users.SuitableTestUserEmail, got.Email)
	s.Equal("a name", *got.Name)
	s.Equal("sherlock", *got.NameFrom)
	s.False(*got.NameInferredFromGithub)

	s.Run("update name", func() {
		code = s.HandleRequest(
			s.NewRequest("PUT", "/api/users/v3", UserV3Upsert{
				userDirectlyEditableFields: userDirectlyEditableFields{
					Name: utils.PointerTo("a different name"),
				},
			}),
			&got)
		s.Equal(http.StatusCreated, code)
		s.Equal(test_users.SuitableTestUserEmail, got.Email)
		s.Equal("a different name", *got.Name)
		s.Equal("sherlock", *got.NameFrom)
		s.False(*got.NameInferredFromGithub)
	})
}

func (s *handlerSuite) TestUserV3Upsert_nameInferredFromGithub() {
	var got UserV3
	code := s.HandleRequest(
		s.NewRequest("PUT", "/api/users/v3", UserV3Upsert{
			userDirectlyEditableFields: userDirectlyEditableFields{
				NameInferredFromGithub: utils.PointerTo(true),
			},
		}),
		&got)
	s.Equal(http.StatusCreated, code)
	s.Equal(test_users.SuitableTestUserEmail, got.Email)
	s.True(*got.NameInferredFromGithub)

	s.Run("then doesn't update name", func() {
		code = s.HandleRequest(
			s.NewRequest("PUT", "/api/users/v3", UserV3Upsert{
				userDirectlyEditableFields: userDirectlyEditableFields{
					Name: utils.PointerTo("a different name"),
				},
			}),
			&got)
		s.Equal(http.StatusOK, code)
		s.Equal(test_users.SuitableTestUserEmail, got.Email)
		s.Nil(got.Name)
		s.True(*got.NameInferredFromGithub)
	})

	s.Run("can set to false", func() {
		code = s.HandleRequest(
			s.NewRequest("PUT", "/api/users/v3", UserV3Upsert{
				userDirectlyEditableFields: userDirectlyEditableFields{
					NameInferredFromGithub: utils.PointerTo(false),
				},
			}),
			&got)
		s.Equal(http.StatusCreated, code)
		s.Equal(test_users.SuitableTestUserEmail, got.Email)
		s.False(*got.NameInferredFromGithub)

		s.Run("then updates name", func() {
			code = s.HandleRequest(
				s.NewRequest("PUT", "/api/users/v3", UserV3Upsert{
					userDirectlyEditableFields: userDirectlyEditableFields{
						Name: utils.PointerTo("a different name"),
					},
				}),
				&got)
			s.Equal(http.StatusCreated, code)
			s.Equal(test_users.SuitableTestUserEmail, got.Email)
			s.Equal("a different name", *got.Name)
			s.False(*got.NameInferredFromGithub)
		})
	})
}

func (s *handlerSuite) TestUserV3Upsert_maximal_sherlockName() {
	slack.UseMockedClient(s.T(), func(c *slack_mocks.MockMockableClient) {
		c.EXPECT().GetUserByEmailContext(mock.Anything, test_users.SuitableTestUserEmail).Return(&slack2.User{
			ID:       "slack ID",
			Name:     "slack username",
			RealName: "name from slack",
		}, nil)
	}, func() {
		github.UseMockedClient(s.T(), func(c *github.MockClient) {
			c.Users.EXPECT().Get(mock.Anything, "").Return(&github2.User{
				ID:    utils.PointerTo[int64](123),
				Login: utils.PointerTo("github username"),
				Name:  utils.PointerTo("name from github"),
			}, nil, nil)
		}, func() {
			var got UserV3
			code := s.HandleRequest(
				s.NewRequest("PUT", "/api/users/v3", UserV3Upsert{
					GithubAccessToken: utils.PointerTo("fake token"),
					userDirectlyEditableFields: userDirectlyEditableFields{
						Name: utils.PointerTo("directly set name"),
					},
				}),
				&got)
			s.Equal(http.StatusCreated, code)
			s.Equal(test_users.SuitableTestUserEmail, got.Email)
			if s.NotNil(got.NameFrom) {
				s.Equal("sherlock", *got.NameFrom)
			}
			if s.NotNil(got.Name) {
				s.Equal("directly set name", *got.Name)
			}
			if s.NotNil(got.SlackID) {
				s.Equal("slack ID", *got.SlackID)
			}
			if s.NotNil(got.SlackUsername) {
				s.Equal("slack username", *got.SlackUsername)
			}
			if s.NotNil(got.GithubID) {
				s.Equal("123", *got.GithubID)
			}
			if s.NotNil(got.GithubUsername) {
				s.Equal("github username", *got.GithubUsername)
			}
		})
	})
}

func (s *handlerSuite) TestUserV3Upsert_maximal_slackName() {
	slack.UseMockedClient(s.T(), func(c *slack_mocks.MockMockableClient) {
		c.EXPECT().GetUserByEmailContext(mock.Anything, test_users.SuitableTestUserEmail).Return(&slack2.User{
			ID:       "slack ID",
			Name:     "slack username",
			RealName: "name from slack",
		}, nil)
	}, func() {
		github.UseMockedClient(s.T(), func(c *github.MockClient) {
			c.Users.EXPECT().Get(mock.Anything, "").Return(&github2.User{
				ID:    utils.PointerTo[int64](123),
				Login: utils.PointerTo("github username"),
				Name:  utils.PointerTo("name from github"),
			}, nil, nil)
		}, func() {
			var got UserV3
			code := s.HandleRequest(
				s.NewRequest("PUT", "/api/users/v3", UserV3Upsert{
					GithubAccessToken: utils.PointerTo("fake token"),
				}),
				&got)
			s.Equal(http.StatusCreated, code)
			s.Equal(test_users.SuitableTestUserEmail, got.Email)
			if s.NotNil(got.NameFrom) {
				s.Equal("slack", *got.NameFrom)
			}
			if s.NotNil(got.Name) {
				s.Equal("name from slack", *got.Name)
			}
			if s.NotNil(got.SlackID) {
				s.Equal("slack ID", *got.SlackID)
			}
			if s.NotNil(got.SlackUsername) {
				s.Equal("slack username", *got.SlackUsername)
			}
			if s.NotNil(got.GithubID) {
				s.Equal("123", *got.GithubID)
			}
			if s.NotNil(got.GithubUsername) {
				s.Equal("github username", *got.GithubUsername)
			}
		})
	})
}

func (s *handlerSuite) TestUserV3Upsert_maximal_githubName() {
	slack.UseMockedClient(s.T(), func(c *slack_mocks.MockMockableClient) {
		c.EXPECT().GetUserByEmailContext(mock.Anything, test_users.SuitableTestUserEmail).Return(&slack2.User{
			ID:       "slack ID",
			Name:     "slack username",
			RealName: "name from slack",
		}, nil)
	}, func() {
		github.UseMockedClient(s.T(), func(c *github.MockClient) {
			c.Users.EXPECT().Get(mock.Anything, "").Return(&github2.User{
				ID:    utils.PointerTo[int64](123),
				Login: utils.PointerTo("github username"),
				Name:  utils.PointerTo("name from github"),
			}, nil, nil)
		}, func() {
			var got UserV3
			code := s.HandleRequest(
				s.NewRequest("PUT", "/api/users/v3", UserV3Upsert{
					GithubAccessToken: utils.PointerTo("fake token"),
					userDirectlyEditableFields: userDirectlyEditableFields{
						NameFrom: utils.PointerTo("github"),
					},
				}),
				&got)
			s.Equal(http.StatusCreated, code)
			s.Equal(test_users.SuitableTestUserEmail, got.Email)
			if s.NotNil(got.NameFrom) {
				s.Equal("github", *got.NameFrom)
			}
			if s.NotNil(got.Name) {
				s.Equal("name from github", *got.Name)
			}
			if s.NotNil(got.SlackID) {
				s.Equal("slack ID", *got.SlackID)
			}
			if s.NotNil(got.SlackUsername) {
				s.Equal("slack username", *got.SlackUsername)
			}
			if s.NotNil(got.GithubID) {
				s.Equal("123", *got.GithubID)
			}
			if s.NotNil(got.GithubUsername) {
				s.Equal("github username", *got.GithubUsername)
			}
		})
	})
}

func (s *handlerSuite) TestUserV3Upsert_dbConflict() {
	github.UseMockedClient(s.T(), func(c *github.MockClient) {
		c.Users.EXPECT().Get(mock.Anything, "").Return(&github2.User{
			ID:    utils.PointerTo[int64](123),
			Login: utils.PointerTo("github username"),
			Name:  utils.PointerTo("name from github"),
		}, nil, nil)
		c.Users.EXPECT().Get(mock.Anything, "").Return(&github2.User{
			ID:    utils.PointerTo[int64](123),
			Login: utils.PointerTo("github username"),
			Name:  utils.PointerTo("name from github"),
		}, nil, nil)
	}, func() {
		var got UserV3
		request := s.NewRequest("PUT", "/api/users/v3", UserV3Upsert{
			GithubAccessToken: utils.PointerTo("fake token"),
		})
		s.UseSuitableUserFor(request)
		code := s.HandleRequest(request, &got)
		s.Equal(http.StatusCreated, code)

		var gotError errors.ErrorResponse
		request = s.NewRequest("PUT", "/api/users/v3", UserV3Upsert{
			GithubAccessToken: utils.PointerTo("fake token"),
		})
		s.UseNonSuitableUserFor(request)
		code = s.HandleRequest(request, &gotError)
		s.Equal(http.StatusConflict, code)
		s.Equal(errors.Conflict, gotError.Type)
	})
}

func (s *handlerSuite) TestUserV3Upsert_maximal_swallowThirdPartyErrors() {
	zerolog.SetGlobalLevel(zerolog.Disabled)
	slack.UseMockedClient(s.T(), func(c *slack_mocks.MockMockableClient) {
		c.EXPECT().GetUserByEmailContext(mock.Anything, test_users.SuitableTestUserEmail).Return(nil, fmt.Errorf("some error"))
	}, func() {
		github.UseMockedClient(s.T(), func(c *github.MockClient) {
			c.Users.EXPECT().Get(mock.Anything, "").Return(nil, nil, fmt.Errorf("some other error"))
		}, func() {
			var got UserV3
			code := s.HandleRequest(
				s.NewRequest("PUT", "/api/users/v3", UserV3Upsert{
					GithubAccessToken: utils.PointerTo("fake token"),
					userDirectlyEditableFields: userDirectlyEditableFields{
						Name: utils.PointerTo("directly set name"),
					},
				}),
				&got)
			s.Equal(http.StatusCreated, code)
			s.Equal(test_users.SuitableTestUserEmail, got.Email)
			if s.NotNil(got.NameFrom) {
				s.Equal("sherlock", *got.NameFrom)
			}
			if s.NotNil(got.Name) {
				s.Equal("directly set name", *got.Name)
			}
			s.Nil(got.SlackID)
			s.Nil(got.SlackUsername)
			s.Nil(got.GithubID)
			s.Nil(got.GithubUsername)
		})
	})
}

func Test_processUserEdits(t *testing.T) {
	config.LoadTestConfig()
	type args struct {
		callingUser     *models.User
		directEdits     userDirectlyEditableFields
		userGithubToken *string
	}
	tests := []struct {
		name              string
		args              args
		slackMockConfig   func(c *slack_mocks.MockMockableClient)
		githubMockConfig  func(c *github.MockClient)
		wantResultingUser *models.User
		wantHasUpdates    bool
	}{
		{
			name: "can do nothing",
			args: args{
				callingUser: &models.User{Email: test_users.SuitableTestUserEmail},
			},
			slackMockConfig: func(c *slack_mocks.MockMockableClient) {
				c.EXPECT().GetUserByEmailContext(mock.Anything, test_users.SuitableTestUserEmail).Return(nil, nil)
			},
			githubMockConfig:  func(c *github.MockClient) {},
			wantResultingUser: &models.User{Email: test_users.SuitableTestUserEmail},
			wantHasUpdates:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			slack.UseMockedClient(t, tt.slackMockConfig, func() {
				github.UseMockedClient(t, tt.githubMockConfig, func() {
					gotResultingUser, gotHasUpdates := processUserEdits(tt.args.callingUser, tt.args.directEdits, tt.args.userGithubToken)
					assert.Equalf(t, tt.wantResultingUser, gotResultingUser, "processUserEdits(%v, %v, %v)", tt.args.callingUser, tt.args.directEdits, tt.args.userGithubToken)
					assert.Equalf(t, tt.wantHasUpdates, gotHasUpdates, "processUserEdits(%v, %v, %v)", tt.args.callingUser, tt.args.directEdits, tt.args.userGithubToken)
				})
			})
		})
	}
}
