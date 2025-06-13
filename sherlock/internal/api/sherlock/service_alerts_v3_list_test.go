package sherlock

import (
	"net/http"

	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
)

func (s *handlerSuite) TestServiceAlertsV3List() {
	s.TestData.ServiceAlert_1()
	s.TestData.ServiceAlert_2()

	s.Run("listAll", func() {
		var got []ServiceAlertV3
		code := s.HandleRequest(
			s.NewRequest("GET", "/api/service-alerts/v3", nil),
			&got)
		s.Equal(http.StatusOK, code)
		s.Len(got, 2)
	})
	s.Run("returnsNone", func() {
		var got []ServiceAlertV3
		code := s.HandleRequest(
			s.NewRequest("GET", "/api/service-alerts/v3?title=foo", nil),
			&got)
		s.Equal(http.StatusOK, code)
		s.Len(got, 0)
	})
	s.Run("returnSome", func() {
		var got []ServiceAlertV3
		code := s.HandleRequest(
			s.NewRequest("GET", "/api/service-alerts/v3?severity=minor", nil),
			&got)
		s.Equal(http.StatusOK, code)
		s.Len(got, 1)
	})
	s.Run("badFilter", func() {
		var got errors.ErrorResponse
		code := s.HandleRequest(
			s.NewRequest("GET", "/api/service-alerts/v3?id=foo", nil),
			&got)
		s.Equal(http.StatusBadRequest, code)
		s.Equal(errors.BadRequest, got.Type)
	})
	s.Run("invalidSeverity", func() {
		var got errors.ErrorResponse
		code := s.HandleRequest(
			s.NewRequest("GET", "/api/service-alerts/v3?severity=foo", nil),
			&got)
		s.Equal(http.StatusBadRequest, code)
		s.Equal(errors.BadRequest, got.Type)
	})

}
