package sherlock

import (
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"net/http"
)

func (s *handlerSuite) TestEnvironmentsV3List_none() {
	var got []EnvironmentV3
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/environments/v3", nil),
		&got)
	s.Equal(http.StatusOK, code)
	s.Len(got, 0)
}

func (s *handlerSuite) TestEnvironmentsV3List_badFilter() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/environments/v3?id=foo", nil),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
}

func (s *handlerSuite) TestEnvironmentsV3List_badLimit() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/environments/v3?limit=foo", nil),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
}

func (s *handlerSuite) TestEnvironmentsV3List_badOffset() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/environments/v3?offset=foo", nil),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
}

func (s *handlerSuite) TestEnvironmentsV3List() {
	s.TestData.Environment_Prod()
	s.TestData.Environment_Dev()
	s.TestData.Environment_Swatomation()
	s.TestData.Environment_Swatomation_DevBee()

	s.Run("all", func() {
		var got []EnvironmentV3
		code := s.HandleRequest(
			s.NewRequest("GET", "/api/environments/v3", nil),
			&got)
		s.Equal(http.StatusOK, code)
		s.Len(got, 4)
		s.Run("associations", func() {
			for _, env := range got {
				s.NotNil(env.DefaultCluster)
			}
		})
	})
	s.Run("none", func() {
		var got []EnvironmentV3
		code := s.HandleRequest(
			s.NewRequest("GET", "/api/environments/v3?name=foo", nil),
			&got)
		s.Equal(http.StatusOK, code)
		s.Len(got, 0)
	})
	s.Run("some", func() {
		var got []EnvironmentV3
		code := s.HandleRequest(
			s.NewRequest("GET", "/api/environments/v3?name=prod", nil),
			&got)
		s.Equal(http.StatusOK, code)
		s.Len(got, 1)
	})
	s.Run("limit and offset", func() {
		var got1 []EnvironmentV3
		code := s.HandleRequest(
			s.NewRequest("GET", "/api/environments/v3?limit=1", nil),
			&got1)
		s.Equal(http.StatusOK, code)
		s.Len(got1, 1)
		var got2 []EnvironmentV3
		code = s.HandleRequest(
			s.NewRequest("GET", "/api/environments/v3?limit=1&offset=1", nil),
			&got2)
		s.Equal(http.StatusOK, code)
		s.Len(got2, 1)
		s.NotEqual(got1[0].ID, got2[0].ID)
	})
}
