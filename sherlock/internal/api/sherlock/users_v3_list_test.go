package sherlock

import (
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/broadinstitute/sherlock/sherlock/internal/self"
	"net/http"
)

func (s *handlerSuite) TestUsersV3List_minimal() {
	var got []UserV3
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/users/v3", nil),
		&got)
	s.Equal(http.StatusOK, code)
	s.Len(got, 4) // admin user + nonsuitable user + suitable user + self
	// This'll never return 0 because the calling user will be upserted, but
	// we can check that the only entries here are ourselves and the self
	// user
	s.Run("first alphabetical user is the nonsuitable user in the test data", func() {
		s.Equal(s.TestData.User_NonSuitable().Email, got[0].Email)
	})
	s.Run("second alphabetical user is the super admin user in the test data", func() {
		s.Equal(s.TestData.User_SuperAdmin().Email, got[1].Email)
	})
	s.Run("third alphabetical user is Sherlock's own self", func() {
		s.Equal(self.Email, got[2].Email)
	})
	s.Run("fourth alphabetical user is the suitable user we made the request as", func() {
		s.Equal(s.TestData.User_Suitable().Email, got[3].Email)
	})
}

func (s *handlerSuite) TestUsersV3List_badFilter() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/users/v3?id=foo", nil),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
}

func (s *handlerSuite) TestUsersV3List_badLimit() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/users/v3?limit=foo", nil),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
}

func (s *handlerSuite) TestUsersV3List_badOffset() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/users/v3?offset=foo", nil),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
}

func (s *handlerSuite) TestUsersV3List() {
	user1 := models.User{
		Email:    "email1@example.com",
		GoogleID: "google-id-1",
	}
	user2 := models.User{
		Email:    "email2@example.com",
		GoogleID: "google-id-2",
		Name:     utils.PointerTo("a name"),
	}
	user3 := models.User{
		Email:          "email3@example.com",
		GoogleID:       "google-id-3",
		Name:           utils.PointerTo("a name"),
		NameFrom:       utils.PointerTo("sherlock"),
		GithubID:       utils.PointerTo("github-id-3"),
		GithubUsername: utils.PointerTo("github-username-3"),
	}
	for _, user := range []models.User{user1, user2, user3} {
		s.NoError(s.DB.Create(&user).Error)
		s.NotZero(user.ID)
	}

	s.Run("all", func() {
		var got []UserV3
		code := s.HandleRequest(
			s.NewRequest("GET", "/api/users/v3", nil),
			&got)
		s.Equal(http.StatusOK, code)
		s.Len(got, 7) // test admin user + test suitable user + test nonsuitable user + self + 3
		s.Run("each has suitability", func() {
			for _, user := range got {
				s.NotZero(user.Suitable)
				s.NotZero(user.SuitabilityDescription)
			}
		})
	})
	s.Run("none", func() {
		var got []UserV3
		code := s.HandleRequest(
			s.NewRequest("GET", "/api/users/v3?name=foo", nil),
			&got)
		s.Equal(http.StatusOK, code)
		s.Len(got, 0)
	})
	s.Run("some", func() {
		var got []UserV3
		code := s.HandleRequest(
			s.NewRequest("GET", "/api/users/v3?name=a%20name", nil),
			&got)
		s.Equal(http.StatusOK, code)
		s.Len(got, 2)
	})
	s.Run("limit and offset", func() {
		var got1 []UserV3
		code := s.HandleRequest(
			s.NewRequest("GET", "/api/users/v3?limit=1", nil),
			&got1)
		s.Equal(http.StatusOK, code)
		s.Len(got1, 1)
		var got2 []UserV3
		code = s.HandleRequest(
			s.NewRequest("GET", "/api/users/v3?limit=1&offset=1", nil),
			&got2)
		s.Equal(http.StatusOK, code)
		s.Len(got2, 1)
		s.NotEqual(got1[0].ID, got2[0].ID)
	})
}
