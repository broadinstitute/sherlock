package sherlock

import (
	"fmt"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"net/http"
)

func (s *handlerSuite) TestAppVersionsProceduresV3Changelog_badChildSelector() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/app-versions/procedures/v3/changelog?child=&parent=123", nil),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
}

func (s *handlerSuite) TestAppVersionsProceduresV3Changelog_badParentSelector() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/app-versions/procedures/v3/changelog?child=123&parent=", nil),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
}

func (s *handlerSuite) TestAppVersionsProceduresV3Changelog_childNotFound() {
	appVersion := s.TestData.AppVersion_Leonardo_V1()
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("GET", fmt.Sprintf("/api/app-versions/procedures/v3/changelog?child=leonardo/abc&parent=%d", appVersion.ID), nil),
		&got)
	s.Equal(http.StatusNotFound, code)
	s.Equal(errors.NotFound, got.Type)
	s.Contains(got.Message, "leonardo/abc")
}

func (s *handlerSuite) TestAppVersionsProceduresV3Changelog_parentNotFound() {
	appVersion := s.TestData.AppVersion_Leonardo_V1()
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("GET", fmt.Sprintf("/api/app-versions/procedures/v3/changelog?child=%d&parent=leonardo/abc", appVersion.ID), nil),
		&got)
	s.Equal(http.StatusNotFound, code)
	s.Equal(errors.NotFound, got.Type)
	s.Contains(got.Message, "leonardo/abc")
}

func (s *handlerSuite) TestAppVersionsProceduresV3Changelog_sameChildAndParent() {
	appVersion := s.TestData.AppVersion_Leonardo_V1()
	var got AppVersionV3ChangelogResponse
	code := s.HandleRequest(
		s.NewRequest("GET", fmt.Sprintf("/api/app-versions/procedures/v3/changelog?child=%d&parent=%d", appVersion.ID, appVersion.ID), nil),
		&got)
	s.Equal(http.StatusOK, code)
	s.True(got.Complete)
	s.Empty(got.Changelog)
}

func (s *handlerSuite) TestAppVersionsProceduresV3Changelog_noPathFound() {
	s.SetNonSuitableTestUserForDB()
	chart := s.TestData.Chart_Leonardo()
	appVersion1 := models.AppVersion{ChartID: chart.ID, AppVersion: "1"}
	s.NoError(s.DB.Create(&appVersion1).Error)
	appVersion2 := models.AppVersion{ChartID: chart.ID, AppVersion: "2"}
	s.NoError(s.DB.Create(&appVersion2).Error)

	var got AppVersionV3ChangelogResponse
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/app-versions/procedures/v3/changelog?child=leonardo%2F1&parent=leonardo%2F2", nil),
		&got)
	s.Equal(http.StatusOK, code)
	s.False(got.Complete)
	if s.Len(got.Changelog, 1) {
		s.Equal(appVersion1.ID, got.Changelog[0].ID)
	}

	code = s.HandleRequest(
		s.NewRequest("GET", "/api/app-versions/procedures/v3/changelog?child=leonardo%2F2&parent=leonardo%2F1", nil),
		&got)
	s.Equal(http.StatusOK, code)
	s.False(got.Complete)
	if s.Len(got.Changelog, 1) {
		s.Equal(appVersion2.ID, got.Changelog[0].ID)
	}
}

func (s *handlerSuite) TestAppVersionsProceduresV3Changelog_findsPath() {
	s.SetNonSuitableTestUserForDB()
	appVersion1 := s.TestData.AppVersion_Leonardo_V1()
	appVersion2 := s.TestData.AppVersion_Leonardo_V2()
	appVersion3 := s.TestData.AppVersion_Leonardo_V3()

	var got AppVersionV3ChangelogResponse
	code := s.HandleRequest(
		s.NewRequest("GET", fmt.Sprintf("/api/app-versions/procedures/v3/changelog?child=%d&parent=%d", appVersion3.ID, appVersion1.ID), nil),
		&got)
	s.Equal(http.StatusOK, code)
	s.True(got.Complete)
	if s.Len(got.Changelog, 2) {
		// It's possible for these tests to run so fast that the ordering of the output can actually get messed up,
		// because we order by createdAt
		s.True((got.Changelog[0].ID == appVersion3.ID && got.Changelog[1].ID == appVersion2.ID) ||
			(got.Changelog[1].ID == appVersion3.ID && got.Changelog[0].ID == appVersion2.ID))
	}
}
