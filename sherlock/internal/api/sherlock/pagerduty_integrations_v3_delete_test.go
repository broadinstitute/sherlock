package sherlock

import (
	"fmt"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"net/http"
)

func (s *handlerSuite) TestPagerdutyIntegrationV3Delete_badSelector() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("DELETE", "/api/pagerduty-integrations/v3/something/with/slashes", nil),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
}

func (s *handlerSuite) TestPagerdutyIntegrationV3Delete_notFound() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("DELETE", "/api/pagerduty-integrations/v3/pd-id/blahblahblah", nil),
		&got)
	s.Equal(http.StatusNotFound, code)
	s.Equal(errors.NotFound, got.Type)
}

func (s *handlerSuite) TestPagerdutyIntegrationV3Delete_suitability() {
	pdi := s.TestData.PagerdutyIntegration_ManuallyTriggeredTerraIncident()
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewNonSuitableRequest("DELETE", fmt.Sprintf("/api/pagerduty-integrations/v3/%d", pdi.ID), nil),
		&got)
	s.Equal(http.StatusForbidden, code)
	s.Equal(errors.Forbidden, got.Type)
}

func (s *handlerSuite) TestPagerdutyIntegrationV3Delete() {
	pdi := s.TestData.PagerdutyIntegration_ManuallyTriggeredTerraIncident()
	var got PagerdutyIntegrationV3
	code := s.HandleRequest(
		s.NewSuitableRequest("DELETE", fmt.Sprintf("/api/pagerduty-integrations/v3/%d", pdi.ID), nil),
		&got)
	s.Equal(http.StatusOK, code)
	s.Equal(pdi.ID, got.ID)
}
