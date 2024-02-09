package sherlock

import (
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *handlerSuite) TestPagerdutyIntegrationsV3Create_badBody() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("POST", "/api/pagerduty-integrations/v3", gin.H{
			"pagerdutyID": 123,
		}),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
	s.Contains(got.Message, "pagerdutyID")
}

func (s *handlerSuite) TestPagerdutyIntegrationsV3Create_sqlValidation() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("POST", "/api/pagerduty-integrations/v3", PagerdutyIntegrationV3Create{}),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
}

func (s *handlerSuite) TestPagerdutyIntegrationsV3Create_suitability() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewNonSuitableRequest("POST", "/api/pagerduty-integrations/v3", PagerdutyIntegrationV3Create{
			PagerdutyID: "pagerduty-id",
			PagerdutyIntegrationV3Edit: PagerdutyIntegrationV3Edit{
				Name: utils.PointerTo("pagerduty-integration-name"),
				Key:  utils.PointerTo("pagerduty-integration-key"),
				Type: utils.PointerTo("pagerduty-integration-type"),
			},
		}),
		&got)
	s.Equal(http.StatusForbidden, code)
	s.Equal(errors.Forbidden, got.Type)
}

func (s *handlerSuite) TestPagerdutyIntegrationsV3Create() {
	var got PagerdutyIntegrationV3
	code := s.HandleRequest(
		s.NewSuitableRequest("POST", "/api/pagerduty-integrations/v3", PagerdutyIntegrationV3Create{
			PagerdutyID: "pagerduty-id",
			PagerdutyIntegrationV3Edit: PagerdutyIntegrationV3Edit{
				Name: utils.PointerTo("pagerduty-integration-name"),
				Key:  utils.PointerTo("pagerduty-integration-key"),
				Type: utils.PointerTo("pagerduty-integration-type"),
			},
		}),
		&got)
	s.Equal(http.StatusCreated, code)
	s.NotZero(got.ID)
	if s.NotNil(got.Name) {
		s.Equal("pagerduty-integration-name", *got.Name)
	}
}

func (s *handlerSuite) TestPagerdutyIntegrationsV3Create_duplicate() {
	pdi := s.TestData.PagerdutyIntegration_ManuallyTriggeredTerraIncident()
	var got PagerdutyIntegrationV3
	code := s.HandleRequest(
		s.NewSuitableRequest("POST", "/api/pagerduty-integrations/v3", PagerdutyIntegrationV3Create{
			PagerdutyID: pdi.PagerdutyID,
			PagerdutyIntegrationV3Edit: PagerdutyIntegrationV3Edit{
				Name: utils.PointerTo("pagerduty-integration-name"),
				Key:  utils.PointerTo("pagerduty-integration-key"),
				Type: utils.PointerTo("pagerduty-integration-type"),
			},
		}),
		&got)
	s.Equal(http.StatusCreated, code)
	s.NotZero(got.ID)
	if s.NotNil(got.Name) {
		s.Equal("pagerduty-integration-name", *got.Name)
	}
	s.Equal(pdi.ID, got.ID)
}
