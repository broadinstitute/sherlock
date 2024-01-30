package sherlock

import (
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"net/http"
)

func (s *handlerSuite) TestDatabaseInstancesV3List_none() {
	var got []DatabaseInstanceV3
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/database-instances/v3", nil),
		&got)
	s.Equal(http.StatusOK, code)
	s.Len(got, 0)
}

func (s *handlerSuite) TestDatabaseInstancesV3List_badFilter() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/database-instances/v3?id=foo", nil),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
}

func (s *handlerSuite) TestDatabaseInstancesV3List_badLimit() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/database-instances/v3?limit=foo", nil),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
}

func (s *handlerSuite) TestDatabaseInstancesV3List_badOffset() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/database-instances/v3?offset=foo", nil),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
}

func (s *handlerSuite) TestDatabaseInstancesV3List() {
	s.TestData.DatabaseInstance_LeonardoDev()
	s.TestData.DatabaseInstance_LeonardoProd()
	s.TestData.DatabaseInstance_LeonardoSwatomation()

	s.Run("all", func() {
		var got []DatabaseInstanceV3
		code := s.HandleRequest(
			s.NewRequest("GET", "/api/database-instances/v3", nil),
			&got)
		s.Equal(http.StatusOK, code)
		s.Len(got, 3)
	})
	s.Run("none", func() {
		var got []DatabaseInstanceV3
		code := s.HandleRequest(
			s.NewRequest("GET", "/api/database-instances/v3?platform=foo", nil),
			&got)
		s.Equal(http.StatusOK, code)
		s.Len(got, 0)
	})
	s.Run("some", func() {
		var got []DatabaseInstanceV3
		code := s.HandleRequest(
			s.NewRequest("GET", "/api/database-instances/v3?platform=kubernetes", nil),
			&got)
		s.Equal(http.StatusOK, code)
		s.Len(got, 1)
	})
	s.Run("limit and offset", func() {
		var got1 []DatabaseInstanceV3
		code := s.HandleRequest(
			s.NewRequest("GET", "/api/database-instances/v3?limit=1", nil),
			&got1)
		s.Equal(http.StatusOK, code)
		s.Len(got1, 1)
		var got2 []DatabaseInstanceV3
		code = s.HandleRequest(
			s.NewRequest("GET", "/api/database-instances/v3?limit=1&offset=1", nil),
			&got2)
		s.Equal(http.StatusOK, code)
		s.Len(got2, 1)
		s.NotEqual(got1[0].ID, got2[0].ID)
	})
}
