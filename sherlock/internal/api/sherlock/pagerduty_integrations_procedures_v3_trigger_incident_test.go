package sherlock

import (
	"github.com/broadinstitute/sherlock/sherlock/internal/clients/pagerduty"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *handlerSuite) TestPagerdutyIntegrationsProceduresV3TriggerIncident_badSelector() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("POST", "/api/pagerduty-integrations/procedures/v3/trigger-incident/something/with/slashes", &pagerduty.AlertSummary{}),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
}

func (s *handlerSuite) TestPagerdutyIntegrationsProceduresV3TriggerIncident_notFound() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("POST", "/api/pagerduty-integrations/procedures/v3/trigger-incident/pd-id/blahblahblah", &pagerduty.AlertSummary{}),
		&got)
	s.Equal(http.StatusNotFound, code)
	s.Equal(errors.NotFound, got.Type)
}

func (s *handlerSuite) TestPagerdutyIntegrationsProceduresV3TriggerIncident_badRequest() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("POST", "/api/pagerduty-integrations/procedures/v3/trigger-incident/pd-id/blahblahblah", gin.H{"summary": 123}),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
	s.Contains(got.Message, "summary")
}

func (s *handlerSuite) TestPagerdutyIntegrationsProceduresV3TriggerIncident_fillsSourceLink() {
	pdi := s.TestData.PagerdutyIntegration_ManuallyTriggeredTerraIncident()
	var got pagerduty.SendAlertResponse
	code := s.HandleRequest(
		s.NewRequest("POST", "/api/pagerduty-integrations/procedures/v3/trigger-incident/pd-id/"+pdi.PagerdutyID, &pagerduty.AlertSummary{
			Summary: "some summary",
		}),
		&got)
	s.Equal(http.StatusOK, code)
}

func (s *handlerSuite) TestPagerdutyIntegrationsProceduresV3TriggerIncident() {
	pdi := s.TestData.PagerdutyIntegration_ManuallyTriggeredTerraIncident()
	var got pagerduty.SendAlertResponse
	code := s.HandleRequest(
		s.NewRequest("POST", "/api/pagerduty-integrations/procedures/v3/trigger-incident/pd-id/"+pdi.PagerdutyID, &pagerduty.AlertSummary{
			Summary:    "some summary",
			SourceLink: "https://broad.io/beehive/1234",
		}),
		&got)
	s.Equal(http.StatusOK, code)
}
