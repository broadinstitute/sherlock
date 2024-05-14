package sherlock

import (
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *handlerSuite) TestChartVersionsV3Upsert_badBody() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("PUT", "/api/chart-versions/v3", gin.H{
			"Description": 123,
		}),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
	s.Contains(got.Message, "description")
}

func (s *handlerSuite) TestChartVersionsV3Upsert_sqlValidation() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("PUT", "/api/chart-versions/v3", ChartVersionV3Create{}),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
	s.Contains(got.Message, "chart_version_present")
}

func (s *handlerSuite) TestChartVersionsV3Upsert_error() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("PUT", "/api/chart-versions/v3", ChartVersionV3Create{}),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
}

func (s *handlerSuite) TestChartVersionsV3Upsert() {
	s.SetNonSuitableTestUserForDB()
	chart := models.Chart{
		Name:      "chart-name",
		ChartRepo: utils.PointerTo("terra-helm"),
	}
	s.NoError(s.DB.Create(&chart).Error)
	chartVersion := models.ChartVersion{ChartID: chart.ID, ChartVersion: "1"}
	s.NoError(s.DB.Create(&chartVersion).Error)

	var got ChartVersionV3
	code := s.HandleRequest(
		s.NewRequest("PUT", "/api/chart-versions/v3", ChartVersionV3Create{
			Chart:              "chart-name",
			ChartVersion:       " 2 ",
			ParentChartVersion: utils.UintToString(chartVersion.ID),
			ChartVersionV3Edit: ChartVersionV3Edit{
				Description: "original description",
			},
		}),
		&got)
	s.Equal(http.StatusCreated, code)
	s.Equal("chart-name", got.Chart)
	if s.NotNil(got.ChartVersion) {
		s.Equal("2", got.ChartVersion)
	}
	if s.NotNil(got.ParentChartVersion) {
		s.Equal("chart-name/1", got.ParentChartVersion)
	}
	if s.NotNil(got.Description) {
		s.Equal("original description", got.Description)
	}
	if s.NotNil(got.AuthoredByInfo) {
		s.Equal(s.TestData.User_Suitable().Email, got.AuthoredByInfo.Email)
	}

	var got2 ChartVersionV3
	code = s.HandleRequest(
		s.NewRequest("PUT", "/api/chart-versions/v3", ChartVersionV3Create{
			Chart:              "chart-name",
			ChartVersion:       "2",
			ParentChartVersion: "chart-name/3",
			ChartVersionV3Edit: ChartVersionV3Edit{
				Description: "edited description",
			},
		}),
		&got2)
	s.Equal(http.StatusCreated, code)
	s.Equal("chart-name", got2.Chart)
	if s.NotNil(got2.ParentChartVersion) {
		s.Equal(got.ParentChartVersion, got2.ParentChartVersion)
	}
	if s.NotNil(got2.Description) {
		s.Equal("edited description", got2.Description)
	}
}
