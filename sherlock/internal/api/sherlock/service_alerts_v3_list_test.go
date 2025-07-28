package sherlock

import (
	"fmt"
	"net/http"

	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
)

func (s *handlerSuite) TestServiceAlertsV3List() {
	s.SetSuitableTestUserForDB()
	alert1 := s.TestData.ServiceAlert_1()
	alertid := s.TestData.ServiceAlert_1().ID
	s.TestData.ServiceAlert_Prod()

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
	s.Run("deletedRecordWithoutFlag", func() {
		var got []ServiceAlertV3
		s.DB.Delete(&alert1)
		code := s.HandleRequest(
			s.NewRequest("GET", fmt.Sprintf("/api/service-alerts/v3?id=%d", alertid), nil),
			&got)
		s.Equal(http.StatusOK, code)
		s.Len(got, 0)
	})
	s.Run("deletedRecordWithFlag", func() {
		var got []ServiceAlertV3
		s.DB.Delete(&alert1)
		code := s.HandleRequest(
			s.NewRequest("GET", fmt.Sprintf("/api/service-alerts/v3?id=%d&include-deleted=true", alertid), nil),
			&got)
		s.Equal(http.StatusOK, code)
		s.Len(got, 1)
	})

}
