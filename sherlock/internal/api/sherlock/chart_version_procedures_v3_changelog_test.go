package sherlock

import (
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"net/http"
)

func (s *handlerSuite) TestChartVersionsProceduresV3Changelog_badChildSelector() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/chart-versions/procedures/v3/changelog?child=&parent=123", nil),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
}

func (s *handlerSuite) TestChartVersionsProceduresV3Changelog_badParentSelector() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/chart-versions/procedures/v3/changelog?child=123&parent=", nil),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
}

func (s *handlerSuite) TestChartVersionsProceduresV3Changelog_childNotFound() {
	chart := models.Chart{Name: "my-chart", ChartRepo: utils.PointerTo("some-repo")}
	s.NoError(s.DB.Create(&chart).Error)
	chartVersion := models.ChartVersion{ChartID: chart.ID, ChartVersion: "1"}
	s.NoError(s.DB.Create(&chartVersion).Error)

	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/chart-versions/procedures/v3/changelog?child=123&parent=my-chart%2F1", nil),
		&got)
	s.Equal(http.StatusNotFound, code)
	s.Equal(errors.NotFound, got.Type)
	s.Contains(got.Message, "123")
}

func (s *handlerSuite) TestChartVersionsProceduresV3Changelog_parentNotFound() {
	chart := models.Chart{Name: "my-chart", ChartRepo: utils.PointerTo("some-repo")}
	s.NoError(s.DB.Create(&chart).Error)
	chartVersion := models.ChartVersion{ChartID: chart.ID, ChartVersion: "1"}
	s.NoError(s.DB.Create(&chartVersion).Error)

	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/chart-versions/procedures/v3/changelog?child=my-chart%2F1&parent=123", nil),
		&got)
	s.Equal(http.StatusNotFound, code)
	s.Equal(errors.NotFound, got.Type)
	s.Contains(got.Message, "123")
}

func (s *handlerSuite) TestChartVersionsProceduresV3Changelog_sameChildAndParent() {
	chart := models.Chart{Name: "my-chart", ChartRepo: utils.PointerTo("some-repo")}
	s.NoError(s.DB.Create(&chart).Error)
	chartVersion := models.ChartVersion{ChartID: chart.ID, ChartVersion: "1"}
	s.NoError(s.DB.Create(&chartVersion).Error)

	var got ChartVersionV3ChangelogResponse
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/chart-versions/procedures/v3/changelog?child=my-chart%2F1&parent=my-chart%2F1", nil),
		&got)
	s.Equal(http.StatusOK, code)
	s.True(got.Complete)
	s.Empty(got.Changelog)
}

func (s *handlerSuite) TestChartVersionsProceduresV3Changelog_noPathFound() {
	chart := models.Chart{Name: "my-chart", ChartRepo: utils.PointerTo("some-repo")}
	s.NoError(s.DB.Create(&chart).Error)
	chartVersion1 := models.ChartVersion{ChartID: chart.ID, ChartVersion: "1"}
	s.NoError(s.DB.Create(&chartVersion1).Error)
	chartVersion2 := models.ChartVersion{ChartID: chart.ID, ChartVersion: "2"}
	s.NoError(s.DB.Create(&chartVersion2).Error)

	var got ChartVersionV3ChangelogResponse
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/chart-versions/procedures/v3/changelog?child=my-chart%2F1&parent=my-chart%2F2", nil),
		&got)
	s.Equal(http.StatusOK, code)
	s.False(got.Complete)
	if s.Len(got.Changelog, 1) {
		s.Equal(chartVersion1.ID, got.Changelog[0].ID)
	}

	code = s.HandleRequest(
		s.NewRequest("GET", "/api/chart-versions/procedures/v3/changelog?child=my-chart%2F2&parent=my-chart%2F1", nil),
		&got)
	s.Equal(http.StatusOK, code)
	s.False(got.Complete)
	if s.Len(got.Changelog, 1) {
		s.Equal(chartVersion2.ID, got.Changelog[0].ID)
	}
}

func (s *handlerSuite) TestChartVersionsProceduresV3Changelog_findsPath() {
	chart := models.Chart{Name: "my-chart", ChartRepo: utils.PointerTo("some-repo")}
	s.NoError(s.DB.Create(&chart).Error)
	chartVersion1 := models.ChartVersion{ChartID: chart.ID, ChartVersion: "1"}
	s.NoError(s.DB.Create(&chartVersion1).Error)
	chartVersion2 := models.ChartVersion{ChartID: chart.ID, ChartVersion: "2", ParentChartVersionID: &chartVersion1.ID}
	s.NoError(s.DB.Create(&chartVersion2).Error)
	chartVersion3 := models.ChartVersion{ChartID: chart.ID, ChartVersion: "3", ParentChartVersionID: &chartVersion2.ID}
	s.NoError(s.DB.Create(&chartVersion3).Error)

	var got ChartVersionV3ChangelogResponse
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/chart-versions/procedures/v3/changelog?child=my-chart/3&parent=my-chart/1", nil),
		&got)
	s.Equal(http.StatusOK, code)
	s.True(got.Complete)
	if s.Len(got.Changelog, 2) {
		// It's possible for these tests to run so fast that the ordering of the output can actually get messed up,
		// because we order by createdAt
		s.True((got.Changelog[0].ID == chartVersion3.ID && got.Changelog[1].ID == chartVersion2.ID) ||
			(got.Changelog[1].ID == chartVersion3.ID && got.Changelog[0].ID == chartVersion2.ID))
	}
}
