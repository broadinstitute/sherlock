package sherlock

import (
	"net/http"

	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/gin-gonic/gin"
)

func (s *handlerSuite) TestServiceAlertV3Create_badBody() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("POST", "/api/service-alerts/v3", gin.H{
			"severity": 123,
		}),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
	s.Contains(got.Message, "severity")
}

func (s *handlerSuite) TestServiceAlertV3CreateSuitable() {
	var got ServiceAlertV3
	code := s.HandleRequest(
		s.NewRequest("POST", "/api/service-alerts/v3", ServiceAlertV3Create{
			OnEnvironment: utils.PointerTo(s.TestData.Environment_Prod().Name),
			ServiceAlertV3EditableFields: ServiceAlertV3EditableFields{
				AlertMessage: utils.PointerTo("Alert message here"),
				Title:        utils.PointerTo("Service Alert Title"),
				Link:         utils.PointerTo("Link here"),
				Severity:     utils.PointerTo("blocker"),
			},
		}),
		&got)
	s.Equal(http.StatusCreated, code)
	if s.NotNil(got.Title) {
		s.Equal("Service Alert Title", *got.Title)
	}
}

func (s *handlerSuite) TestServiceAlertV3CreateNonSuitableUser() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.UseNonSuitableUserFor(s.NewRequest("POST", "/api/service-alerts/v3", ServiceAlertV3Create{
			OnEnvironment: utils.PointerTo(s.TestData.Environment_Prod().Name),
			ServiceAlertV3EditableFields: ServiceAlertV3EditableFields{
				AlertMessage: utils.PointerTo("Alert message here"),
				Title:        utils.PointerTo("Service Alert Title"),
				Link:         utils.PointerTo("Link here"),
				Severity:     utils.PointerTo("blocker"),
			},
		})),
		&got)
	s.Equal(http.StatusForbidden, code)
	s.Equal(errors.Forbidden, got.Type)
}
