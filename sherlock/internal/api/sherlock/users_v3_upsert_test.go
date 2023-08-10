package sherlock

import (
	"github.com/broadinstitute/sherlock/go-shared/pkg/testutils"
	"github.com/broadinstitute/sherlock/sherlock/internal/authentication/test_users"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/google/go-github/v50/github"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
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

func (s *handlerSuite) TestUserV3Upsert_name() {

	var got UserV3
	code := s.HandleRequest(
		s.NewRequest("PUT", "/api/users/v3", UserV3Upsert{
			userDirectlyEditableFields: userDirectlyEditableFields{
				Name: testutils.PointerTo("a name"),
			},
		}),
		&got)
	s.Equal(http.StatusCreated, code)
	s.Equal(test_users.SuitableTestUserEmail, got.Email)
	s.Equal("a name", *got.Name)
	s.False(*got.NameInferredFromGithub)

	s.Run("update name", func() {
		code = s.HandleRequest(
			s.NewRequest("PUT", "/api/users/v3", UserV3Upsert{
				userDirectlyEditableFields: userDirectlyEditableFields{
					Name: testutils.PointerTo("a different name"),
				},
			}),
			&got)
		s.Equal(http.StatusCreated, code)
		s.Equal(test_users.SuitableTestUserEmail, got.Email)
		s.Equal("a different name", *got.Name)
		s.False(*got.NameInferredFromGithub)
	})
}

func (s *handlerSuite) TestUserV3Upsert_nameInferredFromGithub() {
	var got UserV3
	code := s.HandleRequest(
		s.NewRequest("PUT", "/api/users/v3", UserV3Upsert{
			userDirectlyEditableFields: userDirectlyEditableFields{
				NameInferredFromGithub: testutils.PointerTo(true),
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
					Name: testutils.PointerTo("a different name"),
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
					NameInferredFromGithub: testutils.PointerTo(false),
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
						Name: testutils.PointerTo("a different name"),
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

func Test_processUserEdits(t *testing.T) {
	// Load the test config to make sure we silence the log messages here
	config.LoadTestConfig()
	type args struct {
		callingUser *models.User
		githubUser  *github.User
		directEdits userDirectlyEditableFields
	}
	tests := []struct {
		name        string
		args        args
		want        *models.User
		wantChanged bool
	}{
		{
			name:        "keeps same model",
			args:        args{callingUser: &models.User{Model: gorm.Model{ID: 123}, Email: "foo@bar.com"}},
			want:        &models.User{Model: gorm.Model{ID: 123}, Email: "foo@bar.com"},
			wantChanged: false,
		},
		{
			name: "set name inferred from github when empty",
			args: args{
				callingUser: &models.User{},
				directEdits: userDirectlyEditableFields{NameInferredFromGithub: testutils.PointerTo(true)},
			},
			want:        &models.User{NameInferredFromGithub: testutils.PointerTo(true)},
			wantChanged: true,
		},
		{
			name: "set name inferred from github when different value",
			args: args{
				callingUser: &models.User{NameInferredFromGithub: testutils.PointerTo(true)},
				directEdits: userDirectlyEditableFields{NameInferredFromGithub: testutils.PointerTo(false)},
			},
			want:        &models.User{NameInferredFromGithub: testutils.PointerTo(false)},
			wantChanged: true,
		},
		{
			name: "keeps name inferred from github setting when same",
			args: args{
				callingUser: &models.User{NameInferredFromGithub: testutils.PointerTo(true)},
				directEdits: userDirectlyEditableFields{NameInferredFromGithub: testutils.PointerTo(true)},
			},
			want:        &models.User{NameInferredFromGithub: testutils.PointerTo(true)},
			wantChanged: false,
		},
		{
			name: "set name when empty",
			args: args{
				callingUser: &models.User{},
				directEdits: userDirectlyEditableFields{Name: testutils.PointerTo("name")},
			},
			want: &models.User{
				Name:                   testutils.PointerTo("name"),
				NameInferredFromGithub: testutils.PointerTo(false),
			},
			wantChanged: true,
		},
		{
			name: "set name when different value",
			args: args{
				callingUser: &models.User{Name: testutils.PointerTo("different name")},
				directEdits: userDirectlyEditableFields{Name: testutils.PointerTo("name")},
			},
			want: &models.User{
				Name:                   testutils.PointerTo("name"),
				NameInferredFromGithub: testutils.PointerTo(false),
			},
			wantChanged: true,
		},
		{
			name: "keep name when same",
			args: args{
				callingUser: &models.User{Name: testutils.PointerTo("name")},
				directEdits: userDirectlyEditableFields{Name: testutils.PointerTo("name")},
			},
			want:        &models.User{Name: testutils.PointerTo("name")},
			wantChanged: false,
		},
		{
			name: "ignores name change when inferred from github",
			args: args{
				callingUser: &models.User{
					Name:                   testutils.PointerTo("different name"),
					NameInferredFromGithub: testutils.PointerTo(true),
				},
				directEdits: userDirectlyEditableFields{Name: testutils.PointerTo("name")},
			},
			want: &models.User{
				Name:                   testutils.PointerTo("different name"),
				NameInferredFromGithub: testutils.PointerTo(true),
			},
			wantChanged: false,
		},
		{
			name: "ignores name change when also setting inferred from github",
			args: args{
				callingUser: &models.User{
					Name: testutils.PointerTo("different name"),
				},
				directEdits: userDirectlyEditableFields{
					Name:                   testutils.PointerTo("name"),
					NameInferredFromGithub: testutils.PointerTo(true),
				},
			},
			want: &models.User{
				Name:                   testutils.PointerTo("different name"),
				NameInferredFromGithub: testutils.PointerTo(true),
			},
			wantChanged: true,
		},
		{
			name: "respects name change when also setting inferred from github",
			args: args{
				callingUser: &models.User{
					Name:                   testutils.PointerTo("different name"),
					NameInferredFromGithub: testutils.PointerTo(true),
				},
				directEdits: userDirectlyEditableFields{
					Name:                   testutils.PointerTo("name"),
					NameInferredFromGithub: testutils.PointerTo(false),
				},
			},
			want: &models.User{
				Name:                   testutils.PointerTo("name"),
				NameInferredFromGithub: testutils.PointerTo(false),
			},
			wantChanged: true,
		},
		{
			name: "sets github info with name inference to true when name absent",
			args: args{
				callingUser: &models.User{},
				githubUser: &github.User{
					ID:    testutils.PointerTo[int64](123),
					Login: testutils.PointerTo("username"),
				},
			},
			want: &models.User{
				GithubID:               testutils.PointerTo("123"),
				GithubUsername:         testutils.PointerTo("username"),
				NameInferredFromGithub: testutils.PointerTo(true),
			},
			wantChanged: true,
		},
		{
			name: "sets github info with name inference to false when name present",
			args: args{
				callingUser: &models.User{
					Name: testutils.PointerTo("name"),
				},
				githubUser: &github.User{
					ID:    testutils.PointerTo[int64](123),
					Login: testutils.PointerTo("username"),
				},
			},
			want: &models.User{
				Name:                   testutils.PointerTo("name"),
				GithubID:               testutils.PointerTo("123"),
				GithubUsername:         testutils.PointerTo("username"),
				NameInferredFromGithub: testutils.PointerTo(false),
			},
			wantChanged: true,
		},
		{
			name: "sets name from github",
			args: args{
				callingUser: &models.User{},
				githubUser: &github.User{
					ID:    testutils.PointerTo[int64](123),
					Login: testutils.PointerTo("username"),
					Name:  testutils.PointerTo("name"),
				},
			},
			want: &models.User{
				Name:                   testutils.PointerTo("name"),
				GithubID:               testutils.PointerTo("123"),
				GithubUsername:         testutils.PointerTo("username"),
				NameInferredFromGithub: testutils.PointerTo(true),
			},
			wantChanged: true,
		},
		{
			name: "ignores name from github if disabled",
			args: args{
				callingUser: &models.User{
					NameInferredFromGithub: testutils.PointerTo(false),
				},
				githubUser: &github.User{
					ID:    testutils.PointerTo[int64](123),
					Login: testutils.PointerTo("username"),
					Name:  testutils.PointerTo("name"),
				},
			},
			want: &models.User{
				GithubID:               testutils.PointerTo("123"),
				GithubUsername:         testutils.PointerTo("username"),
				NameInferredFromGithub: testutils.PointerTo(false),
			},
			wantChanged: true,
		},
		{
			name: "updates name from github",
			args: args{
				callingUser: &models.User{
					Name:                   testutils.PointerTo("different name"),
					NameInferredFromGithub: testutils.PointerTo(true),
				},
				githubUser: &github.User{
					ID:    testutils.PointerTo[int64](123),
					Login: testutils.PointerTo("username"),
					Name:  testutils.PointerTo("name"),
				},
			},
			want: &models.User{
				Name:                   testutils.PointerTo("name"),
				GithubID:               testutils.PointerTo("123"),
				GithubUsername:         testutils.PointerTo("username"),
				NameInferredFromGithub: testutils.PointerTo(true),
			},
			wantChanged: true,
		},
		{
			name: "updates github info if account changes",
			args: args{
				callingUser: &models.User{
					GithubID:       testutils.PointerTo("old ID"),
					GithubUsername: testutils.PointerTo("old username"),
				},
				githubUser: &github.User{
					ID:    testutils.PointerTo[int64](123),
					Login: testutils.PointerTo("username"),
				},
			},
			want: &models.User{
				GithubID:               testutils.PointerTo("123"),
				GithubUsername:         testutils.PointerTo("username"),
				NameInferredFromGithub: testutils.PointerTo(true),
			},
			wantChanged: true,
		},
		{
			name: "updates github info if just username changes",
			args: args{
				callingUser: &models.User{
					GithubID:       testutils.PointerTo("123"),
					GithubUsername: testutils.PointerTo("old username"),
				},
				githubUser: &github.User{
					ID:    testutils.PointerTo[int64](123),
					Login: testutils.PointerTo("username"),
				},
			},
			want: &models.User{
				GithubID:               testutils.PointerTo("123"),
				GithubUsername:         testutils.PointerTo("username"),
				NameInferredFromGithub: testutils.PointerTo(true),
			},
			wantChanged: true,
		},
		{
			name: "doesn't update if github info equivalent",
			args: args{
				callingUser: &models.User{
					GithubID:               testutils.PointerTo("123"),
					GithubUsername:         testutils.PointerTo("username"),
					NameInferredFromGithub: testutils.PointerTo(true),
				},
				githubUser: &github.User{
					ID:    testutils.PointerTo[int64](123),
					Login: testutils.PointerTo("username"),
				},
			},
			want: &models.User{
				GithubID:               testutils.PointerTo("123"),
				GithubUsername:         testutils.PointerTo("username"),
				NameInferredFromGithub: testutils.PointerTo(true),
			},
			wantChanged: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := processUserEdits(tt.args.callingUser, tt.args.githubUser, tt.args.directEdits)
			assert.Equalf(t, tt.want, got, "processUserEdits(%v, %v, %v)", tt.args.callingUser, tt.args.githubUser, tt.args.directEdits)
			assert.Equalf(t, tt.wantChanged, got1, "processUserEdits(%v, %v, %v)", tt.args.callingUser, tt.args.githubUser, tt.args.directEdits)
		})
	}
}
