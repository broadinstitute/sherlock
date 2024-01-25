package sherlock

import (
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"net/http"
)

func (s *handlerSuite) TestIncidentsV3List_none() {
	var got []IncidentV3
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/incidents/v3", nil),
		&got)
	s.Equal(http.StatusOK, code)
	s.Len(got, 0)
}

func (s *handlerSuite) TestIncidentsV3List_badFilter() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/incidents/v3?id=foo", nil),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
}

func (s *handlerSuite) TestIncidentsV3List_badLimit() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/incidents/v3?limit=foo", nil),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
}

func (s *handlerSuite) TestIncidentsV3List_badOffset() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/incidents/v3?offset=foo", nil),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
}

func (s *handlerSuite) TestIncidentsV3List() {
	s.TestData.Incident_1()
	s.TestData.Incident_2()
	s.TestData.Incident_3()

	s.Run("all", func() {
		var got []IncidentV3
		code := s.HandleRequest(
			s.NewRequest("GET", "/api/incidents/v3", nil),
			&got)
		s.Equal(http.StatusOK, code)
		s.Len(got, 3)
	})
	s.Run("none", func() {
		var got []IncidentV3
		code := s.HandleRequest(
			s.NewRequest("GET", "/api/incidents/v3?ticket=foo", nil),
			&got)
		s.Equal(http.StatusOK, code)
		s.Len(got, 0)
	})
	s.Run("some", func() {
		var got []IncidentV3
		code := s.HandleRequest(
			s.NewRequest("GET", "/api/incidents/v3?ticket=PROD-1", nil),
			&got)
		s.Equal(http.StatusOK, code)
		s.Len(got, 1)
	})
	s.Run("limit and offset", func() {
		var got1 []IncidentV3
		code := s.HandleRequest(
			s.NewRequest("GET", "/api/incidents/v3?limit=1", nil),
			&got1)
		s.Equal(http.StatusOK, code)
		s.Len(got1, 1)
		var got2 []IncidentV3
		code = s.HandleRequest(
			s.NewRequest("GET", "/api/incidents/v3?limit=1&offset=1", nil),
			&got2)
		s.Equal(http.StatusOK, code)
		s.Len(got2, 1)
		s.NotEqual(got1[0].ID, got2[0].ID)
	})
}
