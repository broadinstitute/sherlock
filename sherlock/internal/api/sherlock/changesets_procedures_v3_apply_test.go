package sherlock

import (
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *handlerSuite) TestChangesetsProceduresV3Apply_badBody() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("POST", "/api/changesets/procedures/v3/apply", gin.H{"foo": "bar"}),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
}

func (s *handlerSuite) TestChangesetsProceduresV3Apply_badID() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("POST", "/api/changesets/procedures/v3/apply", []string{"foo"}),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
}

func (s *handlerSuite) TestChangesetsProceduresV3Apply_zeroID() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("POST", "/api/changesets/procedures/v3/apply", []string{"0"}),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
}

func (s *handlerSuite) TestChangesetsProceduresV3Apply_noChangesets() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("POST", "/api/changesets/procedures/v3/apply", []string{}),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
}

func (s *handlerSuite) TestChangesetsProceduresV3Apply_badVerboseOutput() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("POST", "/api/changesets/procedures/v3/apply?verbose-output=foo", []string{"1"}),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
}

func (s *handlerSuite) TestChangesetsProceduresV3Apply() {
	chartRelease := s.TestData.ChartRelease_LeonardoDev()
	changeset := models.Changeset{
		ChartReleaseID: chartRelease.ID,
		To:             chartRelease.ChartReleaseVersion,
	}
	changeset.To.AppVersionExact = utils.PointerTo("different version a")
	someOtherChartRelease := s.TestData.ChartRelease_LeonardoProd()
	someOtherChangeset := models.Changeset{
		ChartReleaseID: someOtherChartRelease.ID,
		To:             someOtherChartRelease.ChartReleaseVersion,
	}
	someOtherChangeset.To.AppVersionExact = utils.PointerTo("different version b")
	ids, err := models.PlanChangesets(s.DB, []models.Changeset{changeset, someOtherChangeset})
	s.NoError(err)

	var got []ChangesetV3
	code := s.HandleRequest(
		s.NewRequest("POST", "/api/changesets/procedures/v3/apply", utils.Map(ids, func(id uint) string {
			return utils.UintToString(id)
		})),
		&got)
	s.Equal(http.StatusOK, code)
	s.Len(got, 2)
	s.Equal(chartRelease.Name, got[0].ChartRelease)
	s.Equal("different version a", *got[0].ChartReleaseInfo.AppVersionExact)
	s.Equal(someOtherChartRelease.Name, got[1].ChartRelease)
	s.Equal("different version b", *got[1].ChartReleaseInfo.AppVersionExact)
}

func (s *handlerSuite) TestChangesetsProceduresV3Apply_forbidden() {
	chartRelease := s.TestData.ChartRelease_LeonardoDev()
	changeset := models.Changeset{
		ChartReleaseID: chartRelease.ID,
		To:             chartRelease.ChartReleaseVersion,
	}
	changeset.To.AppVersionExact = utils.PointerTo("different version a")
	someOtherChartRelease := s.TestData.ChartRelease_LeonardoProd()
	someOtherChangeset := models.Changeset{
		ChartReleaseID: someOtherChartRelease.ID,
		To:             someOtherChartRelease.ChartReleaseVersion,
	}
	someOtherChangeset.To.AppVersionExact = utils.PointerTo("different version b")
	ids, err := models.PlanChangesets(s.DB, []models.Changeset{changeset, someOtherChangeset})
	s.NoError(err)

	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.UseNonSuitableUserFor(s.NewRequest("POST", "/api/changesets/procedures/v3/apply", utils.Map(ids, func(id uint) string {
			return utils.UintToString(id)
		}))),
		&got)
	s.Equal(http.StatusForbidden, code)
	s.Equal(errors.Forbidden, got.Type)

	s.NoError(s.DB.Take(&chartRelease, chartRelease.ID).Error)
	s.NotEqual("different version a", *chartRelease.AppVersionExact)
	s.NoError(s.DB.Take(&someOtherChartRelease, someOtherChartRelease.ID).Error)
	s.NotEqual("different version b", *someOtherChartRelease.AppVersionExact)
}

func (s *handlerSuite) TestChangesetsProceduresV3Apply_notVerbose() {
	chartRelease := s.TestData.ChartRelease_LeonardoDev()
	changeset := models.Changeset{
		ChartReleaseID: chartRelease.ID,
		To:             chartRelease.ChartReleaseVersion,
	}
	changeset.To.AppVersionExact = utils.PointerTo("different version a")
	someOtherChartRelease := s.TestData.ChartRelease_LeonardoProd()
	someOtherChangeset := models.Changeset{
		ChartReleaseID: someOtherChartRelease.ID,
		To:             someOtherChartRelease.ChartReleaseVersion,
	}
	someOtherChangeset.To.AppVersionExact = utils.PointerTo("different version b")
	ids, err := models.PlanChangesets(s.DB, []models.Changeset{changeset, someOtherChangeset})
	s.NoError(err)

	var got []ChangesetV3
	code := s.HandleRequest(
		s.NewRequest("POST", "/api/changesets/procedures/v3/apply?verbose-output=false", utils.Map(ids, func(id uint) string {
			return utils.UintToString(id)
		})),
		&got)
	s.Equal(http.StatusOK, code)
	s.Len(got, 2)
	s.Empty(got[0].ChartRelease)
	s.Empty(got[1].ChartRelease)

	s.NoError(s.DB.Take(&chartRelease, chartRelease.ID).Error)
	s.Equal("different version a", *chartRelease.AppVersionExact)
	s.NoError(s.DB.Take(&someOtherChartRelease, someOtherChartRelease.ID).Error)
	s.Equal("different version b", *someOtherChartRelease.AppVersionExact)
}
