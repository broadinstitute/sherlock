package sherlock

import (
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *handlerSuite) TestChangesetsProceduresV3PlanAndApply_badBody() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("POST", "/api/changesets/procedures/v3/plan-and-apply", gin.H{"environments": "bar"}),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
}

func (s *handlerSuite) TestChangesetsProceduresV3PlanAndApply_badVerboseOutput() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("POST", "/api/changesets/procedures/v3/plan-and-apply?verbose-output=foo", ChangesetV3PlanRequest{}),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
}

func (s *handlerSuite) TestChangesetsProceduresV3PlanAndApply_noChangesets() {
	var got []ChangesetV3
	code := s.HandleRequest(
		s.NewRequest("POST", "/api/changesets/procedures/v3/plan-and-apply", ChangesetV3PlanRequest{}),
		&got)
	s.Equal(http.StatusOK, code)
	s.Empty(got)
}

func (s *handlerSuite) TestChangesetsProceduresV3PlanAndApply_chartReleaseEntryError() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("POST", "/api/changesets/procedures/v3/plan-and-apply", ChangesetV3PlanRequest{
			ChartReleases: []ChangesetV3PlanRequestChartReleaseEntry{
				{
					ChangesetV3Create: ChangesetV3Create{
						ChartRelease: "!!!",
					},
				},
			},
		}),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
	s.Contains(got.Message, "!!!")
}

func (s *handlerSuite) TestChangesetsProceduresV3PlanAndApply_environmentEntryError() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("POST", "/api/changesets/procedures/v3/plan-and-apply", ChangesetV3PlanRequest{
			Environments: []ChangesetV3PlanRequestEnvironmentEntry{
				{
					Environment: "!!!",
				},
			},
		}),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
	s.Contains(got.Message, "!!!")
}

func (s *handlerSuite) TestChangesetsProceduresV3PlanAndApply_recreateEntryError() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("POST", "/api/changesets/procedures/v3/plan-and-apply", ChangesetV3PlanRequest{
			RecreateChangesets: []uint{0},
		}),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
	s.Contains(got.Message, "0")
}

func (s *handlerSuite) TestChangesetsProceduresV3PlanAndApply_conflict() {
	changesetToRecreate := s.TestData.Changeset_LeonardoDev_V1toV2Superseded()
	chartReleaseToUpdate := s.TestData.ChartRelease_LeonardoDev()

	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("POST", "/api/changesets/procedures/v3/plan-and-apply", ChangesetV3PlanRequest{
			ChartReleases: []ChangesetV3PlanRequestChartReleaseEntry{
				{
					ChangesetV3Create: ChangesetV3Create{
						ChartRelease:         chartReleaseToUpdate.Name,
						ToAppVersionResolver: utils.PointerTo("exact"),
						ToAppVersionExact:    utils.PointerTo("some new version"),
					},
				},
			},
			RecreateChangesets: []uint{changesetToRecreate.ID},
		}),
		&got)
	s.Equal(http.StatusConflict, code)
	s.Equal(errors.Conflict, got.Type)
}

func (s *handlerSuite) TestChangesetsProceduresV3Plan_forbidden() {
	chartReleaseToUpdate := s.TestData.ChartRelease_LeonardoProd()

	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.UseNonSuitableUserFor(s.NewRequest("POST", "/api/changesets/procedures/v3/plan-and-apply", ChangesetV3PlanRequest{
			ChartReleases: []ChangesetV3PlanRequestChartReleaseEntry{
				{
					ChangesetV3Create: ChangesetV3Create{
						ChartRelease:         chartReleaseToUpdate.Name,
						ToAppVersionResolver: utils.PointerTo("exact"),
						ToAppVersionExact:    utils.PointerTo("some new version"),
					},
				},
			},
		})),
		&got)
	s.Equal(http.StatusForbidden, code)
	s.Equal(errors.Forbidden, got.Type)
}

func (s *handlerSuite) TestChangesetsProceduresV3PlanAndApply() {
	changesetToRecreate := s.TestData.Changeset_LeonardoDev_V1toV2Superseded()
	chartReleaseToUpdate := s.TestData.ChartRelease_LeonardoProd()
	environmentToUpdate := s.TestData.Environment_Staging()
	s.TestData.ChartRelease_LeonardoStaging()

	var got []ChangesetV3
	code := s.HandleRequest(
		s.NewRequest("POST", "/api/changesets/procedures/v3/plan-and-apply", ChangesetV3PlanRequest{
			ChartReleases: []ChangesetV3PlanRequestChartReleaseEntry{
				{
					ChangesetV3Create: ChangesetV3Create{
						ChartRelease:         chartReleaseToUpdate.Name,
						ToAppVersionResolver: utils.PointerTo("exact"),
						ToAppVersionExact:    utils.PointerTo("some new version"),
					},
				},
			},
			Environments: []ChangesetV3PlanRequestEnvironmentEntry{
				{
					Environment:                          environmentToUpdate.Name,
					UseExactVersionsFromOtherEnvironment: utils.PointerTo(s.TestData.Environment_Dev().Name),
				},
			},
			RecreateChangesets: []uint{changesetToRecreate.ID},
		}),
		&got)
	s.Equal(http.StatusOK, code)
	s.Len(got, 3)
	for _, changeset := range got {
		s.NotEmpty(changeset.ID)
		s.NotEmpty(changeset.ChartReleaseInfo.Name)
		s.Equal(changeset.ChartReleaseInfo.AppVersionExact, changeset.ToAppVersionExact)
	}
}

func (s *handlerSuite) TestChangesetProceduresV3PlanAndApply_noneGiven() {
	var got []ChangesetV3
	code := s.HandleRequest(
		s.NewRequest("POST", "/api/changesets/procedures/v3/plan-and-apply", ChangesetV3PlanRequest{}),
		&got)
	s.Equal(http.StatusOK, code)
	s.NotNil(got)
	s.Empty(got)
}

func (s *handlerSuite) TestChangesetsProceduresV3PlanAndApply_nonePlanned() {
	var got []ChangesetV3
	code := s.HandleRequest(
		s.NewRequest("POST", "/api/changesets/procedures/v3/plan-and-apply", ChangesetV3PlanRequest{
			Environments: []ChangesetV3PlanRequestEnvironmentEntry{
				{
					Environment: s.TestData.Environment_Staging().Name,
				},
			},
		}),
		&got)
	s.Equal(http.StatusOK, code)
	s.NotNil(got)
	s.Empty(got)
}

func (s *handlerSuite) TestChangesetsProceduresV3PlanAndApply_notVerbose() {
	changesetToRecreate := s.TestData.Changeset_LeonardoDev_V1toV2Superseded()
	chartReleaseToUpdate := s.TestData.ChartRelease_LeonardoProd()
	environmentToUpdate := s.TestData.Environment_Staging()
	s.TestData.ChartRelease_LeonardoStaging()

	var got []ChangesetV3
	code := s.HandleRequest(
		s.NewRequest("POST", "/api/changesets/procedures/v3/plan-and-apply?verbose-output=false", ChangesetV3PlanRequest{
			ChartReleases: []ChangesetV3PlanRequestChartReleaseEntry{
				{
					ChangesetV3Create: ChangesetV3Create{
						ChartRelease:         chartReleaseToUpdate.Name,
						ToAppVersionResolver: utils.PointerTo("exact"),
						ToAppVersionExact:    utils.PointerTo("some new version"),
					},
				},
			},
			Environments: []ChangesetV3PlanRequestEnvironmentEntry{
				{
					Environment:                          environmentToUpdate.Name,
					UseExactVersionsFromOtherEnvironment: utils.PointerTo(s.TestData.Environment_Dev().Name),
				},
			},
			RecreateChangesets: []uint{changesetToRecreate.ID},
		}),
		&got)
	s.Equal(http.StatusOK, code)
	s.Len(got, 3)
	for _, changeset := range got {
		s.NotEmpty(changeset.ID)
		s.Nil(changeset.ChartReleaseInfo)
	}
}
