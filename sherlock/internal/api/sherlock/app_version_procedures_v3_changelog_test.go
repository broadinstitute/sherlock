package sherlock

import (
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
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
	chart := models.Chart{Name: "my-chart", ChartRepo: utils.PointerTo("some-repo")}
	s.NoError(s.DB.Create(&chart).Error)
	appVersion := models.AppVersion{ChartID: chart.ID, AppVersion: "1"}
	s.NoError(s.DB.Create(&appVersion).Error)

	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/app-versions/procedures/v3/changelog?child=123&parent=my-chart%2F1", nil),
		&got)
	s.Equal(http.StatusNotFound, code)
	s.Equal(errors.NotFound, got.Type)
	s.Contains(got.Message, "123")
}

func (s *handlerSuite) TestAppVersionsProceduresV3Changelog_parentNotFound() {
	chart := models.Chart{Name: "my-chart", ChartRepo: utils.PointerTo("some-repo")}
	s.NoError(s.DB.Create(&chart).Error)
	appVersion := models.AppVersion{ChartID: chart.ID, AppVersion: "1"}
	s.NoError(s.DB.Create(&appVersion).Error)

	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/app-versions/procedures/v3/changelog?child=my-chart%2F1&parent=123", nil),
		&got)
	s.Equal(http.StatusNotFound, code)
	s.Equal(errors.NotFound, got.Type)
	s.Contains(got.Message, "123")
}

func (s *handlerSuite) TestAppVersionsProceduresV3Changelog_sameChildAndParent() {
	chart := models.Chart{Name: "my-chart", ChartRepo: utils.PointerTo("some-repo")}
	s.NoError(s.DB.Create(&chart).Error)
	appVersion := models.AppVersion{ChartID: chart.ID, AppVersion: "1"}
	s.NoError(s.DB.Create(&appVersion).Error)

	var got AppVersionV3ChangelogResponse
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/app-versions/procedures/v3/changelog?child=my-chart%2F1&parent=my-chart%2F1", nil),
		&got)
	s.Equal(http.StatusOK, code)
	s.True(got.Complete)
	s.Empty(got.Changelog)
}

func (s *handlerSuite) TestAppVersionsProceduresV3Changelog_noPathFound() {
	chart := models.Chart{Name: "my-chart", ChartRepo: utils.PointerTo("some-repo")}
	s.NoError(s.DB.Create(&chart).Error)
	appVersion1 := models.AppVersion{ChartID: chart.ID, AppVersion: "1"}
	s.NoError(s.DB.Create(&appVersion1).Error)
	appVersion2 := models.AppVersion{ChartID: chart.ID, AppVersion: "2"}
	s.NoError(s.DB.Create(&appVersion2).Error)

	var got AppVersionV3ChangelogResponse
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/app-versions/procedures/v3/changelog?child=my-chart%2F1&parent=my-chart%2F2", nil),
		&got)
	s.Equal(http.StatusOK, code)
	s.False(got.Complete)
	if s.Len(got.Changelog, 1) {
		s.Equal(appVersion1.ID, got.Changelog[0].ID)
	}

	code = s.HandleRequest(
		s.NewRequest("GET", "/api/app-versions/procedures/v3/changelog?child=my-chart%2F2&parent=my-chart%2F1", nil),
		&got)
	s.Equal(http.StatusOK, code)
	s.False(got.Complete)
	if s.Len(got.Changelog, 1) {
		s.Equal(appVersion2.ID, got.Changelog[0].ID)
	}
}

func (s *handlerSuite) TestAppVersionsProceduresV3Changelog_findsPath() {
	chart := models.Chart{Name: "my-chart", ChartRepo: utils.PointerTo("some-repo")}
	s.NoError(s.DB.Create(&chart).Error)
	appVersion1 := models.AppVersion{ChartID: chart.ID, AppVersion: "1"}
	s.NoError(s.DB.Create(&appVersion1).Error)
	appVersion2 := models.AppVersion{ChartID: chart.ID, AppVersion: "2", ParentAppVersionID: &appVersion1.ID}
	s.NoError(s.DB.Create(&appVersion2).Error)
	appVersion3 := models.AppVersion{ChartID: chart.ID, AppVersion: "3", ParentAppVersionID: &appVersion2.ID}
	s.NoError(s.DB.Create(&appVersion3).Error)

	var got AppVersionV3ChangelogResponse
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/app-versions/procedures/v3/changelog?child=my-chart/3&parent=my-chart/1", nil),
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
