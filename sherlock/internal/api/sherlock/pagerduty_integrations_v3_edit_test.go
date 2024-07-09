package sherlock

import (
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *handlerSuite) TestPagerdutyIntegrationV3Edit_badSelector() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewSuperAdminRequest("PATCH", "/api/pagerduty-integrations/v3/something/with/slashes", nil),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
}

func (s *handlerSuite) TestPagerdutyIntegrationV3Edit_badBody() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewSuperAdminRequest("PATCH", "/api/pagerduty-integrations/v3/123", gin.H{
			"key": 123,
		}),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
}

func (s *handlerSuite) TestPagerdutyIntegrationV3Edit_notFound() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewSuperAdminRequest("PATCH", "/api/pagerduty-integrations/v3/pd-id/blahblahblah", PagerdutyIntegrationV3Edit{
			Key: utils.PointerTo("pagerduty-integration-key"),
		}),
		&got)
	s.Equal(http.StatusNotFound, code)
	s.Equal(errors.NotFound, got.Type)
}

func (s *handlerSuite) TestPagerdutyIntegrationV3Edit_sqlValidation() {
	pdi := s.TestData.PagerdutyIntegration_ManuallyTriggeredTerraIncident()
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewSuperAdminRequest("PATCH", fmt.Sprintf("/api/pagerduty-integrations/v3/%d", pdi.ID), PagerdutyIntegrationV3Edit{
			Key: utils.PointerTo(""),
		}),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
	s.Contains(got.Message, "key")
}

func (s *handlerSuite) TestPagerdutyIntegrationV3Edit_suitability() {
	pdi := s.TestData.PagerdutyIntegration_ManuallyTriggeredTerraIncident()
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewNonSuitableRequest("PATCH", fmt.Sprintf("/api/pagerduty-integrations/v3/%d", pdi.ID), PagerdutyIntegrationV3Edit{
			Key: utils.PointerTo("foo"),
		}),
		&got)
	s.Equal(http.StatusForbidden, code)
	s.Equal(errors.Forbidden, got.Type)
}

func (s *handlerSuite) TestPagerdutyIntegrationV3Edit() {
	pdi := s.TestData.PagerdutyIntegration_ManuallyTriggeredTerraIncident()
	var got PagerdutyIntegrationV3
	code := s.HandleRequest(
		s.NewSuperAdminRequest("PATCH", fmt.Sprintf("/api/pagerduty-integrations/v3/%d", pdi.ID), PagerdutyIntegrationV3Edit{
			Key:  utils.PointerTo("a key"),
			Name: utils.PointerTo("a name"),
		}),
		&got)
	s.Equal(http.StatusOK, code)
	if s.NotNil(got.Name) {
		s.Equal("a name", *got.Name)
	}
}
