package sherlock

import (
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *handlerSuite) TestChartVersionsV3Edit_badSelector() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("PATCH", "/api/chart-versions/v3/something/with/slashes", nil),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
}

func (s *handlerSuite) TestChartVersionsV3Edit_badBody() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("PATCH", "/api/chart-versions/v3/1", gin.H{
			"description": 123,
		}),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
	s.Contains(got.Message, "description")
}

func (s *handlerSuite) TestChartVersionsV3Edit_notFound() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("PATCH", "/api/chart-versions/v3/123", ChartVersionV3Edit{}),
		&got)
	s.Equal(http.StatusNotFound, code)
	s.Equal(errors.NotFound, got.Type)
}

func (s *handlerSuite) TestChartVersionsV3Edit() {
	s.SetNonSuitableTestUserForDB()
	chart := models.Chart{
		Name:      "name",
		ChartRepo: utils.PointerTo("terra-helm"),
	}
	s.NoError(s.DB.Create(&chart).Error)
	chartVersion := models.ChartVersion{ChartID: chart.ID, ChartVersion: "1"}
	s.NoError(s.DB.Create(&chartVersion).Error)
	s.NotZero(&chartVersion.ID)

	var got ChartVersionV3
	code := s.HandleRequest(
		s.NewRequest("PATCH", fmt.Sprintf("/api/chart-versions/v3/%d", chartVersion.ID), ChartVersionV3Edit{
			Description: "description",
		}),
		&got)
	s.Equal(http.StatusOK, code)
	if s.NotNil(got.ChartVersion) {
		s.Equal("description", got.Description)
	}
}
