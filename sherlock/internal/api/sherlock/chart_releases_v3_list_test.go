package sherlock

import (
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"net/http"
)

func (s *handlerSuite) TestChartReleasesV3List_none() {
	var got []ChartReleaseV3
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/chart-releases/v3", nil),
		&got)
	s.Equal(http.StatusOK, code)
	s.Len(got, 0)
}

func (s *handlerSuite) TestChartReleasesV3List_badFilter() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/chart-releases/v3?id=foo", nil),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
}

func (s *handlerSuite) TestChartReleasesV3List_badLimit() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/chart-releases/v3?limit=foo", nil),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
}

func (s *handlerSuite) TestChartReleasesV3List_badOffset() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/chart-releases/v3?offset=foo", nil),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
}

func (s *handlerSuite) TestChartReleasesV3List() {
	s.TestData.ChartRelease_LeonardoDev()
	s.TestData.ChartRelease_LeonardoProd()
	s.TestData.ChartRelease_D2pDdpAzureDev()

	s.Run("all", func() {
		var got []ChartReleaseV3
		code := s.HandleRequest(
			s.NewRequest("GET", "/api/chart-releases/v3", nil),
			&got)
		s.Equal(http.StatusOK, code)
		s.Len(got, 3)
		s.Run("associations", func() {
			for _, release := range got {
				s.NotNil(release.ChartInfo)
				s.NotNil(release.EnvironmentInfo)
			}
		})
	})
	s.Run("none", func() {
		var got []ChartReleaseV3
		code := s.HandleRequest(
			s.NewRequest("GET", "/api/chart-releases/v3?name=foo", nil),
			&got)
		s.Equal(http.StatusOK, code)
		s.Len(got, 0)
	})
	s.Run("some", func() {
		var got []ChartReleaseV3
		code := s.HandleRequest(
			s.NewRequest("GET", "/api/chart-releases/v3?name=leonardo-prod", nil),
			&got)
		s.Equal(http.StatusOK, code)
		s.Len(got, 1)
	})
	s.Run("limit and offset", func() {
		var got1 []ChartReleaseV3
		code := s.HandleRequest(
			s.NewRequest("GET", "/api/chart-releases/v3?limit=1", nil),
			&got1)
		s.Equal(http.StatusOK, code)
		s.Len(got1, 1)
		var got2 []ChartReleaseV3
		code = s.HandleRequest(
			s.NewRequest("GET", "/api/chart-releases/v3?limit=1&offset=1", nil),
			&got2)
		s.Equal(http.StatusOK, code)
		s.Len(got2, 1)
		s.NotEqual(got1[0].ID, got2[0].ID)
	})
}
