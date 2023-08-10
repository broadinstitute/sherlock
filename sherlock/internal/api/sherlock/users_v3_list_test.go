package sherlock

import (
	"github.com/broadinstitute/sherlock/go-shared/pkg/testutils"
	"github.com/broadinstitute/sherlock/sherlock/internal/authentication/test_users"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"net/http"
)

func (s *handlerSuite) TestUsersV3List_minimal() {
	var got []UserV3
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/users/v3", nil),
		&got)
	s.Equal(http.StatusOK, code)
	s.Len(got, 1)
	// This'll never return 0 because the calling user will be upserted, but
	// we can check that the only entry here is ourselves
	s.Run("inserted user is self", func() {
		s.Equal(test_users.SuitableTestUserEmail, got[0].Email)
	})
}

func (s *handlerSuite) TestUsersV3List() {
	user1 := models.User{
		Email:    "email1@example.com",
		GoogleID: "google-id-1",
	}
	user2 := models.User{
		Email:    "email2@example.com",
		GoogleID: "google-id-2",
		Name:     testutils.PointerTo("a name"),
	}
	user3 := models.User{
		Email:                  "email3@example.com",
		GoogleID:               "google-id-3",
		Name:                   testutils.PointerTo("a name"),
		NameInferredFromGithub: testutils.PointerTo(false),
		GithubID:               testutils.PointerTo("github-id-3"),
		GithubUsername:         testutils.PointerTo("github-username-3"),
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
		s.Len(got, 4)
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
