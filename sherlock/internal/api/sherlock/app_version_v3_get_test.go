package sherlock

import (
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"net/http"
)

func (s *handlerSuite) TestAppVersionV3Get_badSelector() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/app-versions/v3/mytextwithnoslashes", nil),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
}

func (s *handlerSuite) TestAppVersionV3Get_notFound() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/app-versions/v3/1", nil),
		&got)
	s.Equal(http.StatusNotFound, code)
	s.Equal(errors.NotFound, got.Type)
}

func (s *handlerSuite) TestAppVersionV3Get() {
	s.SetNonSuitableTestUserForDB()
	chart := models.Chart{Name: "my-chart", ChartRepo: utils.PointerTo("some-repo")}
	s.NoError(s.DB.Create(&chart).Error)
	appVersion := models.AppVersion{ChartID: chart.ID, AppVersion: "1"}
	s.NoError(s.DB.Create(&appVersion).Error)

	var got AppVersionV3
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/app-versions/v3/my-chart/1", nil),
		&got)
	s.Equal(http.StatusOK, code)
	if s.NotNil(got.ChartInfo) {
		s.NotZero(got.ChartInfo.ID)
	}
	if s.NotNil(got.AuthoredByInfo) {
		s.Equal(s.TestData.User_NonSuitable().Email, got.AuthoredByInfo.Email)
	}
	s.Equal("1", got.AppVersion)
}

func (s *handlerSuite) TestAppVersionV3Get_notFoundChart() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/app-versions/v3/my-chart/1", nil),
		&got)
	s.Equal(http.StatusNotFound, code)
	s.Equal(errors.NotFound, got.Type)
}
