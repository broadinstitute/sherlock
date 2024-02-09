package sherlock

import (
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/authentication/test_users"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"net/http"
)

func (s *handlerSuite) TestChartVersionsV3List_none() {
	var got []ChartVersionV3
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/chart-versions/v3", nil),
		&got)
	s.Equal(http.StatusOK, code)
	s.Len(got, 0)
}

func (s *handlerSuite) TestChartVersionsV3List_badFilter() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/chart-versions/v3?id=foo", nil),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
}

func (s *handlerSuite) TestChartVersionsV3List_badLimit() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/chart-versions/v3?limit=foo", nil),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
}

func (s *handlerSuite) TestChartVersionsV3List_badOffset() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/chart-versions/v3?offset=foo", nil),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
}

func (s *handlerSuite) TestChartVersionsV3List() {
	s.SetNonSuitableTestUserForDB()
	chart1 := models.Chart{
		Name:      "name1",
		ChartRepo: utils.PointerTo("terra-helm"),
	}
	chart2 := models.Chart{
		Name:      "name2",
		ChartRepo: utils.PointerTo("terra-helm"),
	}
	chart3 := models.Chart{
		Name:      "name3",
		ChartRepo: utils.PointerTo("terra-helm"),
	}

	for _, chart := range []*models.Chart{&chart1, &chart2, &chart3} {
		s.NoError(s.DB.Create(chart).Error)
		s.NotZero(chart.ID)
	}

	chartVersion1 := models.ChartVersion{ChartID: chart1.ID, ChartVersion: "1"}
	chartVersion2 := models.ChartVersion{ChartID: chart2.ID, ChartVersion: "2"}
	chartVersion3 := models.ChartVersion{ChartID: chart3.ID, ChartVersion: "3"}

	for _, chartVersion := range []models.ChartVersion{chartVersion1, chartVersion2, chartVersion3} {
		s.NoError(s.DB.Create(&chartVersion).Error)
		s.NotZero(chartVersion.ID)
	}

	s.Run("all", func() {
		var got []ChartVersionV3
		code := s.HandleRequest(
			s.NewRequest("GET", "/api/chart-versions/v3", nil),
			&got)
		s.Equal(http.StatusOK, code)
		s.Len(got, 3)
		s.Run("associations", func() {
			for _, chartVersion := range got {
				s.NotNil(chartVersion.ChartInfo)
			}
		})
	})
	s.Run("all via user filter", func() {
		var got []ChartVersionV3
		code := s.HandleRequest(
			s.NewRequest("GET", fmt.Sprintf("/api/chart-versions/v3?authoredBy=%s", test_users.NonSuitableTestUserEmail), nil),
			&got)
		s.Equal(http.StatusOK, code)
		s.Len(got, 3)
	})
	s.Run("none", func() {
		var got []ChartVersionV3
		code := s.HandleRequest(
			s.NewRequest("GET", "/api/chart-versions/v3?chartVersion=foo", nil),
			&got)
		s.Equal(http.StatusOK, code)
		s.Len(got, 0)
	})
	s.Run("some", func() {
		var got []ChartVersionV3
		code := s.HandleRequest(
			s.NewRequest("GET", "/api/chart-versions/v3?chartVersion=1", nil),
			&got)
		s.Equal(http.StatusOK, code)
		s.Len(got, 1)
	})
	s.Run("limit and offset", func() {
		var got1 []ChartVersionV3
		code := s.HandleRequest(
			s.NewRequest("GET", "/api/chart-versions/v3?limit=1", nil),
			&got1)
		s.Equal(http.StatusOK, code)
		s.Len(got1, 1)
		var got2 []ChartVersionV3
		code = s.HandleRequest(
			s.NewRequest("GET", "/api/chart-versions/v3?limit=1&offset=1", nil),
			&got2)
		s.Equal(http.StatusOK, code)
		s.Len(got2, 1)
		s.NotEqual(got1[0].ID, got2[0].ID)
	})
}
