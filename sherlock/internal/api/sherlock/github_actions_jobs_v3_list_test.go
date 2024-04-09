package sherlock

import (
	"fmt"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"net/http"
)

func (s *handlerSuite) TestGithubActionJobsV3List_none() {
	var got []GithubActionsJobV3
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/github-actions-jobs/v3", nil),
		&got)
	s.Equal(http.StatusOK, code)
	s.Len(got, 0)
}

func (s *handlerSuite) TestGithubActionJobsV3List_badFilter() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/github-actions-jobs/v3?id=foo", nil),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
}

func (s *handlerSuite) TestGithubActionJobsV3List_badLimit() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/github-actions-jobs/v3?limit=foo", nil),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
}

func (s *handlerSuite) TestGithubActionJobsV3List_badOffset() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/github-actions-jobs/v3?offset=foo", nil),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
}

func (s *handlerSuite) TestGithubActionJobsV3List() {
	s.TestData.GithubActionsJob_1()
	s.TestData.GithubActionsJob_2()

	s.Run("all", func() {
		var got []GithubActionsJobV3
		code := s.HandleRequest(
			s.NewRequest("GET", "/api/github-actions-jobs/v3", nil),
			&got)
		s.Equal(http.StatusOK, code)
		s.Len(got, 2)
	})
	s.Run("none", func() {
		var got []GithubActionsJobV3
		code := s.HandleRequest(
			s.NewRequest("GET", "/api/github-actions-jobs/v3?githubActionsJobID=123", nil),
			&got)
		s.Equal(http.StatusOK, code)
		s.Len(got, 0)
	})
	s.Run("some", func() {
		var got []GithubActionsJobV3
		code := s.HandleRequest(
			s.NewRequest("GET", fmt.Sprintf("/api/github-actions-jobs/v3?githubActionsJobID=%d", s.TestData.GithubActionsJob_1().GithubActionsJobID), nil),
			&got)
		s.Equal(http.StatusOK, code)
		s.Len(got, 1)
	})
	s.Run("limit and offset", func() {
		var got1 []GithubActionsJobV3
		code := s.HandleRequest(
			s.NewRequest("GET", "/api/github-actions-jobs/v3?limit=1", nil),
			&got1)
		s.Equal(http.StatusOK, code)
		s.Len(got1, 1)
		var got2 []GithubActionsJobV3
		code = s.HandleRequest(
			s.NewRequest("GET", "/api/github-actions-jobs/v3?limit=1&offset=1", nil),
			&got2)
		s.Equal(http.StatusOK, code)
		s.Len(got2, 1)
		s.NotEqual(got1[0].ID, got2[0].ID)
	})
}
