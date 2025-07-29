package sherlock

import (
	"fmt"
	"net/http"

	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
)

func (s *handlerSuite) TestServiceAlertV3Delete_badSelector() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("DELETE", "/api/service-alerts/v3/something/with/slashes", nil),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
}

func (s *handlerSuite) TestServiceAlertV3Delete_notFound() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("DELETE", "/api/service-alerts/v3/123", nil),
		&got)
	s.Equal(http.StatusNotFound, code)
	s.Equal(errors.NotFound, got.Type)
}

func (s *handlerSuite) TestServiceAlertV3Delete() {
	servicealert1 := s.TestData.ServiceAlert_1()

	var got ServiceAlertV3
	code := s.HandleRequest(
		s.NewRequest("DELETE", fmt.Sprintf("/api/service-alerts/v3/%d", servicealert1.ID), nil),
		&got)
	s.Equal(http.StatusOK, code)
	if s.NotNil(got.AlertMessage) {
		s.Equal(*servicealert1.AlertMessage, *got.AlertMessage)
	}
}

func (s *handlerSuite) TestServiceAlertV3DeleteNonSuitable() {
	servicealert := s.TestData.ServiceAlert_Prod()

	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.UseNonSuitableUserFor(s.NewRequest("DELETE", fmt.Sprintf("/api/service-alerts/v3/%d", servicealert.ID), nil)),
		&got)
	s.Equal(http.StatusForbidden, code)
	s.Equal(errors.Forbidden, got.Type)
}
