package sherlock

import (
	"fmt"
	"net/http"

	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/gin-gonic/gin"
)

func (s *handlerSuite) TestServiceAlertV3Edit_badSelector() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("PATCH", "/api/service-alerts/v3/something/with/slashes", nil),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
}

func (s *handlerSuite) TestServiceAlertV3Edit_badBody() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("PATCH", "/api/service-alerts/v3/123", gin.H{
			"severity": 123,
		}),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
	s.Contains(got.Message, "severity")
}

func (s *handlerSuite) TestServiceAlertV3Edit_notFound() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("PATCH", "/api/service-alerts/v3/123", ServiceAlertV3EditableFields{
			Title:        utils.PointerTo("test tile"),
			AlertMessage: utils.PointerTo("this message is a test"),
		}),
		&got)
	s.Equal(http.StatusNotFound, code)
	s.Equal(errors.NotFound, got.Type)
}

func (s *handlerSuite) TestServiceAlertV3Edit() {
	testAlert := s.TestData.ServiceAlert_1()

	edit := models.ServiceAlert{
		Title:           utils.PointerTo("original title"),
		AlertMessage:    utils.PointerTo("original message"),
		Link:            utils.PointerTo("link"),
		Severity:        utils.PointerTo("minor"),
		OnEnvironmentID: utils.PointerTo(*testAlert.OnEnvironmentID),
	}
	s.NoError(s.DB.Create(&edit).Error)

	var got ServiceAlertV3
	code := s.HandleRequest(
		s.NewRequest("PATCH", fmt.Sprintf("/api/service-alerts/v3/%d", edit.ID), ServiceAlertV3EditableFields{
			Title:        utils.PointerTo("a whole new title"),
			AlertMessage: utils.PointerTo("updated message"),
		}),
		&got)
	s.Equal(http.StatusOK, code)
	if s.NotNil(got.AlertMessage) {
		s.Equal("updated message", *got.AlertMessage)
	}
}
