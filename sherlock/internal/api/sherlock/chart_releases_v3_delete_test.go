package sherlock

import (
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"net/http"
)

func (s *handlerSuite) TestChartReleasesV3Delete_badSelector() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("DELETE", "/api/chart-releases/v3/!!!!!", nil),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
}

func (s *handlerSuite) TestChartReleasesV3Delete_notFound() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("DELETE", "/api/chart-releases/v3/leonardo-dev", nil),
		&got)
	s.Equal(http.StatusNotFound, code)
	s.Equal(errors.NotFound, got.Type)
}

func (s *handlerSuite) TestChartReleasesV3Delete() {
	cr := s.TestData.ChartRelease_LeonardoDev()
	var got ChartReleaseV3
	code := s.HandleRequest(
		s.NewRequest("DELETE", "/api/chart-releases/v3/"+cr.Name, nil),
		&got)
	s.Equal(http.StatusOK, code)
	s.Equal(cr.ID, got.ID)
}

func (s *handlerSuite) TestChartReleasesV3Delete_suitability() {
	s.TestData.ChartRelease_LeonardoProd()
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.UseNonSuitableUserFor(s.NewRequest("DELETE", "/api/chart-releases/v3/leonardo-prod", nil)),
		&got)
	s.Equal(http.StatusForbidden, code)
	s.Equal(errors.Forbidden, got.Type)
}

func (s *handlerSuite) TestChartReleasesV3Delete_suitabilityAllowed() {
	cr := s.TestData.ChartRelease_LeonardoProd()
	var got ChartReleaseV3
	code := s.HandleRequest(
		s.UseSuitableUserFor(s.NewRequest("DELETE", "/api/chart-releases/v3/leonardo-prod", nil)),
		&got)
	s.Equal(http.StatusOK, code)
	s.Equal(cr.Name, got.Name)
}
