package sherlock

import (
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/authentication/test_users"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"net/http"
)

func (s *handlerSuite) TestChartVersionV3Get_badSelector() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/chart-versions/v3/mytextwithnoslashes", nil),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
}

func (s *handlerSuite) TestChartVersionV3Get_notFound() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/chart-versions/v3/1", nil),
		&got)
	s.Equal(http.StatusNotFound, code)
	s.Equal(errors.NotFound, got.Type)
}

func (s *handlerSuite) TestChartVersionV3Get() {
	s.SetNonSuitableTestUserForDB()
	chart := models.Chart{Name: "my-chart", ChartRepo: utils.PointerTo("some-repo")}
	s.NoError(s.DB.Create(&chart).Error)
	chartVersion := models.ChartVersion{ChartID: chart.ID, ChartVersion: "1"}
	s.NoError(s.DB.Create(&chartVersion).Error)

	var got ChartVersionV3
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/chart-versions/v3/my-chart/1", nil),
		&got)
	s.Equal(http.StatusOK, code)
	if s.NotNil(got.ChartInfo) {
		s.NotZero(got.ChartInfo.ID)
	}
	if s.NotNil(got.AuthoredByInfo) {
		s.Equal(test_users.NonSuitableTestUserEmail, got.AuthoredByInfo.Email)
	}
	s.Equal("1", got.ChartVersion)
}

func (s *handlerSuite) TestChartVersionV3Get_notFoundChart() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/chart-versions/v3/my-chart/1", nil),
		&got)
	s.Equal(http.StatusNotFound, code)
	s.Equal(errors.NotFound, got.Type)
}
