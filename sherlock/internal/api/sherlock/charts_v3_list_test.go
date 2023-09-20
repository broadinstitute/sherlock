package sherlock

import (
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"net/http"
)

func (s *handlerSuite) TestChartsV3List_none() {
	var got []ChartV3
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/charts/v3", nil),
		&got)
	s.Equal(http.StatusOK, code)
	s.Len(got, 0)
}

func (s *handlerSuite) TestChartsV3List_badFilter() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/charts/v3?id=foo", nil),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
}

func (s *handlerSuite) TestChartsV3List_badLimit() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/charts/v3?limit=foo", nil),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
}

func (s *handlerSuite) TestChartsV3List_badOffset() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/charts/v3?offset=foo", nil),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
}

func (s *handlerSuite) TestChartsV3List() {
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
	for _, chart := range []models.Chart{chart1, chart2, chart3} {
		s.NoError(s.DB.Create(&chart).Error)
		s.NotZero(chart.ID)
	}

	s.Run("all", func() {
		var got []ChartV3
		code := s.HandleRequest(
			s.NewRequest("GET", "/api/charts/v3", nil),
			&got)
		s.Equal(http.StatusOK, code)
		s.Len(got, 3)
	})
	s.Run("none", func() {
		var got []ChartV3
		code := s.HandleRequest(
			s.NewRequest("GET", "/api/charts/v3?name=foo", nil),
			&got)
		s.Equal(http.StatusOK, code)
		s.Len(got, 0)
	})
	s.Run("some", func() {
		var got []ChartV3
		code := s.HandleRequest(
			s.NewRequest("GET", "/api/charts/v3?name=name1", nil),
			&got)
		s.Equal(http.StatusOK, code)
		s.Len(got, 1)
	})
	s.Run("limit and offset", func() {
		var got1 []ChartV3
		code := s.HandleRequest(
			s.NewRequest("GET", "/api/charts/v3?limit=1", nil),
			&got1)
		s.Equal(http.StatusOK, code)
		s.Len(got1, 1)
		var got2 []ChartV3
		code = s.HandleRequest(
			s.NewRequest("GET", "/api/charts/v3?limit=1&offset=1", nil),
			&got2)
		s.Equal(http.StatusOK, code)
		s.Len(got2, 1)
		s.NotEqual(got1[0].ID, got2[0].ID)
	})
}
