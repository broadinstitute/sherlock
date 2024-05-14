package sherlock

import (
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"net/http"
)

func (s *handlerSuite) TestAppVersionsV3List_none() {
	var got []AppVersionV3
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/app-versions/v3", nil),
		&got)
	s.Equal(http.StatusOK, code)
	s.Len(got, 0)
}

func (s *handlerSuite) TestAppVersionsV3List_badFilter() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/app-versions/v3?id=foo", nil),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
}

func (s *handlerSuite) TestAppVersionsV3List_badLimit() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/app-versions/v3?limit=foo", nil),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
}

func (s *handlerSuite) TestAppVersionsV3List_badOffset() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/app-versions/v3?offset=foo", nil),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
}

func (s *handlerSuite) TestAppVersionsV3List() {
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

	println("Chart1", chart1.ID)
	appVersion1 := models.AppVersion{ChartID: chart1.ID, AppVersion: "1"}
	appVersion2 := models.AppVersion{ChartID: chart2.ID, AppVersion: "2"}
	appVersion3 := models.AppVersion{ChartID: chart3.ID, AppVersion: "3"}

	for _, appVersion := range []models.AppVersion{appVersion1, appVersion2, appVersion3} {
		s.NoError(s.DB.Create(&appVersion).Error)
		s.NotZero(appVersion.ID)
	}

	s.Run("all", func() {
		var got []AppVersionV3
		code := s.HandleRequest(
			s.NewRequest("GET", "/api/app-versions/v3", nil),
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
		var got []AppVersionV3
		code := s.HandleRequest(
			s.NewRequest("GET", fmt.Sprintf("/api/app-versions/v3?authoredBy=%s", s.TestData.User_NonSuitable().Email), nil),
			&got)
		s.Equal(http.StatusOK, code)
		s.Len(got, 3)
	})
	s.Run("none", func() {
		var got []AppVersionV3
		code := s.HandleRequest(
			s.NewRequest("GET", "/api/app-versions/v3?appVersion=foo", nil),
			&got)
		s.Equal(http.StatusOK, code)
		s.Len(got, 0)
	})
	s.Run("some", func() {
		var got []AppVersionV3
		code := s.HandleRequest(
			s.NewRequest("GET", "/api/app-versions/v3?appVersion=1", nil),
			&got)
		s.Equal(http.StatusOK, code)
		s.Len(got, 1)
	})
	s.Run("limit and offset", func() {
		var got1 []AppVersionV3
		code := s.HandleRequest(
			s.NewRequest("GET", "/api/app-versions/v3?limit=1", nil),
			&got1)
		s.Equal(http.StatusOK, code)
		s.Len(got1, 1)
		var got2 []AppVersionV3
		code = s.HandleRequest(
			s.NewRequest("GET", "/api/app-versions/v3?limit=1&offset=1", nil),
			&got2)
		s.Equal(http.StatusOK, code)
		s.Len(got2, 1)
		s.NotEqual(got1[0].ID, got2[0].ID)
	})
}
