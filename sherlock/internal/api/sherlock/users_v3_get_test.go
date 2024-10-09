package sherlock

import (
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/testutils"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"net/http"
	"testing"
	"time"
)

func (s *handlerSuite) TestUserV3Get_badSelector() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/users/v3/foo-bar", nil),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
}

func (s *handlerSuite) TestUserV3Get_notFound() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/users/v3/foo@example.com", nil),
		&got)
	s.Equal(http.StatusNotFound, code)
	s.Equal(errors.NotFound, got.Type)
}

// TestUserV3Get_deactivated is more of a smoke test -- we actually test the code that powers this
// where it's defined in the middleware package, and it technically applies across every single API
// route in this package, but it's unnecessary to exhaustively test that. We just check here that
// it's wired up correctly.
func (s *handlerSuite) TestUserV3Get_deactivated() {
	// Have to switch to super admin to make this user deactivated
	s.SetUserForDB(utils.PointerTo(s.TestData.User_SuperAdmin()))
	s.NoError(s.DB.Model(utils.PointerTo(s.TestData.User_NonSuitable())).Omit(clause.Associations).Update("deactivated_at", utils.PointerTo(time.Now())).Error)

	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewNonSuitableRequest("GET", "/api/users/v3/foo@example.com", nil),
		&got)
	s.Equal(http.StatusForbidden, code)
	s.Equal(errors.Forbidden, got.Type)
}

func (s *handlerSuite) TestUsersV3Get_self() {
	for _, selector := range []string{"self", "me", s.TestData.User_Suitable().Email, fmt.Sprintf("google-id/%s", s.TestData.User_Suitable().GoogleID)} {
		s.Run(fmt.Sprintf("get own user via '%s'", selector), func() {
			var got UserV3
			code := s.HandleRequest(
				s.NewRequest("GET", fmt.Sprintf("/api/users/v3/%s", selector), nil),
				&got)
			s.Equal(http.StatusOK, code)
			s.Equal(s.TestData.User_Suitable().Email, got.Email)
			s.True(*got.Suitable)
		})
	}
}

func (s *handlerSuite) TestUserV3Get_others() {
	dummyUser := models.User{
		Email:          "dummy@example.com",
		GoogleID:       "some-fake-google-id",
		GithubUsername: utils.PointerTo("some-fake-github-username"),
		GithubID:       utils.PointerTo("some-fake-github-id"),
	}
	s.NoError(s.DB.Create(&dummyUser).Error)
	s.NotZero(dummyUser.ID)
	for _, selector := range []string{
		utils.UintToString(dummyUser.ID),
		dummyUser.Email,
		fmt.Sprintf("google-id/%s", dummyUser.GoogleID),
		fmt.Sprintf("github/%s", *dummyUser.GithubUsername),
		fmt.Sprintf("github-id/%s", *dummyUser.GithubID),
	} {
		s.Run(fmt.Sprintf("get dummy user via '%s'", selector), func() {
			var got UserV3
			code := s.HandleRequest(
				s.NewRequest("GET", fmt.Sprintf("/api/users/v3/%s", selector), nil),
				&got)
			s.Equal(http.StatusOK, code)
			s.Equal(dummyUser.ID, got.ID)
			s.Equal(dummyUser.Email, got.Email)
			s.False(*got.Suitable)
		})
	}
}

func Test_userModelFromSelector(t *testing.T) {
	type args struct {
		selector string
	}
	tests := []struct {
		name      string
		args      args
		wantQuery models.User
		wantErr   assert.ErrorAssertionFunc
	}{
		{
			name:    "empty",
			args:    args{selector: ""},
			wantErr: assert.Error,
		},
		{
			name:      "id",
			args:      args{selector: "123"},
			wantQuery: models.User{Model: gorm.Model{ID: 123}},
			wantErr:   assert.NoError,
		},
		{
			name:    "invalid id",
			args:    args{selector: testutils.StringNumberTooBigForInt},
			wantErr: assert.Error,
		},
		{
			name:      "email",
			args:      args{selector: "foo@example.com"},
			wantQuery: models.User{Email: "foo@example.com"},
			wantErr:   assert.NoError,
		},
		{
			name:      "google id",
			args:      args{selector: "google-id/foo"},
			wantQuery: models.User{GoogleID: "foo"},
			wantErr:   assert.NoError,
		},
		{
			name:    "empty google id",
			args:    args{selector: "google-id/"},
			wantErr: assert.Error,
		},
		{
			name:      "github username",
			args:      args{selector: "github/foo"},
			wantQuery: models.User{GithubUsername: utils.PointerTo("foo")},
			wantErr:   assert.NoError,
		},
		{
			name:    "empty github username",
			args:    args{selector: "github/"},
			wantErr: assert.Error,
		},
		{
			name:      "github id",
			args:      args{selector: "github-id/foo"},
			wantQuery: models.User{GithubID: utils.PointerTo("foo")},
			wantErr:   assert.NoError,
		},
		{
			name:    "empty github id",
			args:    args{selector: "github-id/"},
			wantErr: assert.Error,
		},
		{
			name:    "invalid",
			args:    args{selector: "foo"},
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotQuery, err := userModelFromSelector(tt.args.selector)
			if !tt.wantErr(t, err, fmt.Sprintf("userModelFromSelector(%v)", tt.args.selector)) {
				return
			}
			assert.Equalf(t, tt.wantQuery, gotQuery, "userModelFromSelector(%v)", tt.args.selector)
		})
	}
}
