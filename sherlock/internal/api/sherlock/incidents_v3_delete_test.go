package sherlock

import (
	"fmt"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"net/http"
)

func (s *handlerSuite) TestIncidentV3Delete_badSelector() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("DELETE", "/api/incidents/v3/something/with/slashes", nil),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
}

func (s *handlerSuite) TestIncidentV3Delete_notFound() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("DELETE", "/api/incidents/v3/123", nil),
		&got)
	s.Equal(http.StatusNotFound, code)
	s.Equal(errors.NotFound, got.Type)
}

func (s *handlerSuite) TestIncidentV3Delete() {
	s.SetNonSuitableTestUserForDB()
	incident1 := s.TestData.Incident_1()

	var got IncidentV3
	code := s.HandleRequest(
		s.NewRequest("DELETE", fmt.Sprintf("/api/incidents/v3/%d", incident1.ID), nil),
		&got)
	s.Equal(http.StatusOK, code)
	if s.NotNil(got.Description) {
		s.Equal(*incident1.Description, *got.Description)
	}
}
