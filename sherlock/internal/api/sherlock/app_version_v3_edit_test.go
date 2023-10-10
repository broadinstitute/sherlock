package sherlock

import (
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *handlerSuite) TestAppVersionsV3Edit_badSelector() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("PATCH", "/api/app-versions/v3/something/with/slashes", nil),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
}

func (s *handlerSuite) TestAppVersionsV3Edit_badBody() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("PATCH", "/api/app-versions/v3/1", gin.H{
			"description": 123,
		}),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
	s.Contains(got.Message, "description")
}

func (s *handlerSuite) TestAppVersionsV3Edit_notFound() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("PATCH", "/api/app-versions/v3/123", AppVersionV3Edit{}),
		&got)
	s.Equal(http.StatusNotFound, code)
	s.Equal(errors.NotFound, got.Type)
}

func (s *handlerSuite) TestAppVersionsV3Edit() {
	s.SetNonSuitableTestUserForDB()
	chart := models.Chart{
		Name:      "name",
		ChartRepo: utils.PointerTo("terra-helm"),
	}
	s.NoError(s.DB.Create(&chart).Error)
	appVersion := models.AppVersion{ChartID: chart.ID, AppVersion: "1"}
	s.NoError(s.DB.Create(&appVersion).Error)
	s.NotZero(&appVersion.ID)

	var got AppVersionV3
	code := s.HandleRequest(
		s.NewRequest("PATCH", fmt.Sprintf("/api/app-versions/v3/%d", appVersion.ID), AppVersionV3Edit{
			Description: "description",
		}),
		&got)
	s.Equal(http.StatusOK, code)
	if s.NotNil(got.AppVersion) {
		s.Equal("description", got.Description)
	}
}
