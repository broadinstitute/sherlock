package sherlock

import (
	"fmt"
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
	chartVersion := s.TestData.ChartVersion_Leonardo_V1()
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("GET", fmt.Sprintf("/api/chart-versions/procedures/v3/changelog?child=leonardo/abc&parent=%d", chartVersion.ID), nil),
		&got)
	s.Equal(http.StatusNotFound, code)
	s.Equal(errors.NotFound, got.Type)
	s.Contains(got.Message, "leonardo/abc")
}

func (s *handlerSuite) TestChartVersionsProceduresV3Changelog_parentNotFound() {
	chartVersion := s.TestData.ChartVersion_Leonardo_V1()
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("GET", fmt.Sprintf("/api/chart-versions/procedures/v3/changelog?child=%d&parent=leonardo/abc", chartVersion.ID), nil),
		&got)
	s.Equal(http.StatusNotFound, code)
	s.Equal(errors.NotFound, got.Type)
	s.Contains(got.Message, "leonardo/abc")
}

func (s *handlerSuite) TestChartVersionsProceduresV3Changelog_sameChildAndParent() {
	chartVersion := s.TestData.ChartVersion_Leonardo_V1()
	var got ChartVersionV3ChangelogResponse
	code := s.HandleRequest(
		s.NewRequest("GET", fmt.Sprintf("/api/chart-versions/procedures/v3/changelog?child=%d&parent=%d", chartVersion.ID, chartVersion.ID), nil),
		&got)
	s.Equal(http.StatusOK, code)
	s.True(got.Complete)
	s.Empty(got.Changelog)
}

func (s *handlerSuite) TestChartVersionsProceduresV3Changelog_noPathFound() {
	s.SetNonSuitableTestUserForDB()
	chart := s.TestData.Chart_Leonardo()
	chartVersion1 := models.ChartVersion{ChartID: chart.ID, ChartVersion: "1"}
	s.NoError(s.DB.Create(&chartVersion1).Error)
	chartVersion2 := models.ChartVersion{ChartID: chart.ID, ChartVersion: "2"}
	s.NoError(s.DB.Create(&chartVersion2).Error)

	var got ChartVersionV3ChangelogResponse
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/chart-versions/procedures/v3/changelog?child=leonardo%2F1&parent=leonardo%2F2", nil),
		&got)
	s.Equal(http.StatusOK, code)
	s.False(got.Complete)
	if s.Len(got.Changelog, 1) {
		s.Equal(chartVersion1.ID, got.Changelog[0].ID)
	}

	code = s.HandleRequest(
		s.NewRequest("GET", "/api/chart-versions/procedures/v3/changelog?child=leonardo%2F2&parent=leonardo%2F1", nil),
		&got)
	s.Equal(http.StatusOK, code)
	s.False(got.Complete)
	if s.Len(got.Changelog, 1) {
		s.Equal(chartVersion2.ID, got.Changelog[0].ID)
	}
}

func (s *handlerSuite) TestChartVersionsProceduresV3Changelog_findsPath() {
	s.SetNonSuitableTestUserForDB()
	chartVersion1 := s.TestData.ChartVersion_Leonardo_V1()
	chartVersion2 := s.TestData.ChartVersion_Leonardo_V2()
	chartVersion3 := s.TestData.ChartVersion_Leonardo_V3()

	var got ChartVersionV3ChangelogResponse
	code := s.HandleRequest(
		s.NewRequest("GET", fmt.Sprintf("/api/chart-versions/procedures/v3/changelog?child=%d&parent=%d", chartVersion3.ID, chartVersion1.ID), nil),
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
