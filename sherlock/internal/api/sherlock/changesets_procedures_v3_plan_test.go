package sherlock

import (
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *handlerSuite) TestChangesetsProceduresV3Plan_badBody() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("POST", "/api/changesets/procedures/v3/plan", gin.H{"environments": "bar"}),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
}

func (s *handlerSuite) TestChangesetsProceduresV3Plan_badVerboseOutput() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("POST", "/api/changesets/procedures/v3/plan?verbose-output=foo", ChangesetV3PlanRequest{}),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
}

func (s *handlerSuite) TestChangesetsProceduresV3Plan_noChangesets() {
	var got []ChangesetV3
	code := s.HandleRequest(
		s.NewRequest("POST", "/api/changesets/procedures/v3/plan", ChangesetV3PlanRequest{}),
		&got)
	s.Equal(http.StatusOK, code)
	s.Empty(got)
}

func (s *handlerSuite) TestChangesetsProceduresV3Plan_chartReleaseEntryError() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("POST", "/api/changesets/procedures/v3/plan", ChangesetV3PlanRequest{
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

func (s *handlerSuite) TestChangesetsProceduresV3Plan_environmentEntryError() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("POST", "/api/changesets/procedures/v3/plan", ChangesetV3PlanRequest{
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

func (s *handlerSuite) TestChangesetsProceduresV3Plan_recreateEntryError() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("POST", "/api/changesets/procedures/v3/plan", ChangesetV3PlanRequest{
			RecreateChangesets: []uint{0},
		}),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
	s.Contains(got.Message, "0")
}

func (s *handlerSuite) TestChangesetsProceduresV3Plan_allowsNonSuitable() {
	chartReleaseToUpdate := s.TestData.ChartRelease_LeonardoProd()

	var got []ChangesetV3
	code := s.HandleRequest(
		s.UseNonSuitableUserFor(s.NewRequest("POST", "/api/changesets/procedures/v3/plan", ChangesetV3PlanRequest{
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
	s.Equal(http.StatusOK, code)
	s.Len(got, 1)
	for _, changeset := range got {
		s.NotEmpty(changeset.ID)
		s.NotEmpty(changeset.ChartReleaseInfo.Name)
	}
}

func (s *handlerSuite) TestChangesetsProceduresV3Plan() {
	changesetToRecreate := s.TestData.Changeset_LeonardoDev_V1toV2Superseded()
	chartReleaseToUpdate := s.TestData.ChartRelease_LeonardoProd()
	environmentToUpdate := s.TestData.Environment_Staging()
	s.TestData.ChartRelease_LeonardoStaging()

	var got []ChangesetV3
	code := s.HandleRequest(
		s.NewRequest("POST", "/api/changesets/procedures/v3/plan", ChangesetV3PlanRequest{
			ChartReleases: []ChangesetV3PlanRequestChartReleaseEntry{
				{
					ChangesetV3Create: ChangesetV3Create{
						ChartRelease:         chartReleaseToUpdate.Name,
						ToAppVersionResolver: utils.PointerTo("exact"),
						ToAppVersionExact:    utils.PointerTo(" some new version "),
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
		if changeset.ChartReleaseInfo.Name == chartReleaseToUpdate.Name {
			s.Equal("some new version", *changeset.ToAppVersionExact, "(perhaps unexpected whitespace?)")
			s.Equal("exact", *changeset.ToAppVersionResolver)
		}
	}
}

func (s *handlerSuite) TestChangesetProceduresV3Plan_noneGiven() {
	var got []ChangesetV3
	code := s.HandleRequest(
		s.NewRequest("POST", "/api/changesets/procedures/v3/plan", ChangesetV3PlanRequest{}),
		&got)
	s.Equal(http.StatusOK, code)
	s.NotNil(got)
	s.Empty(got)
}

func (s *handlerSuite) TestChangesetsProceduresV3Plan_nonePlanned() {
	var got []ChangesetV3
	code := s.HandleRequest(
		s.NewRequest("POST", "/api/changesets/procedures/v3/plan", ChangesetV3PlanRequest{
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

func (s *handlerSuite) TestChangesetsProceduresV3Plan_notVerbose() {
	changesetToRecreate := s.TestData.Changeset_LeonardoDev_V1toV2Superseded()
	chartReleaseToUpdate := s.TestData.ChartRelease_LeonardoDev()
	environmentToUpdate := s.TestData.Environment_Staging()
	s.TestData.ChartRelease_LeonardoStaging()

	var got []ChangesetV3
	code := s.HandleRequest(
		s.NewRequest("POST", "/api/changesets/procedures/v3/plan?verbose-output=false", ChangesetV3PlanRequest{
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

func (s *handlerSuite) TestChangesetV3PlanRequest_parseChartReleaseEntries_badModel() {
	r := ChangesetV3PlanRequest{
		ChartReleases: []ChangesetV3PlanRequestChartReleaseEntry{
			{
				ChangesetV3Create: ChangesetV3Create{
					ChartRelease: "foo",
				},
			},
		},
	}
	_, err := r.parseChartReleaseEntries(s.DB)
	s.ErrorContains(err, "not found")
}

func (s *handlerSuite) TestChangesetV3PlanRequest_parseChartReleaseEntries_bothSpecified() {
	r := ChangesetV3PlanRequest{
		ChartReleases: []ChangesetV3PlanRequestChartReleaseEntry{
			{
				ChangesetV3Create: ChangesetV3Create{
					ChartRelease: s.TestData.ChartRelease_LeonardoDev().Name,
				},
				UseExactVersionsFromOtherChartRelease: utils.PointerTo(s.TestData.ChartRelease_LeonardoProd().Name),
				FollowVersionsFromOtherChartRelease:   utils.PointerTo(s.TestData.ChartRelease_LeonardoProd().Name),
			},
		},
	}
	_, err := r.parseChartReleaseEntries(s.DB)
	s.ErrorContains(err, "both")
}

func (s *handlerSuite) TestChangesetV3PlanRequest_parseChartReleaseEntries_useExactVersionsFromOtherChartRelease_badSelector() {
	r := ChangesetV3PlanRequest{
		ChartReleases: []ChangesetV3PlanRequestChartReleaseEntry{
			{
				ChangesetV3Create: ChangesetV3Create{
					ChartRelease: s.TestData.ChartRelease_LeonardoDev().Name,
				},
				UseExactVersionsFromOtherChartRelease: utils.PointerTo("!!!"),
			},
		},
	}
	_, err := r.parseChartReleaseEntries(s.DB)
	s.ErrorContains(err, "!!!")
}

func (s *handlerSuite) TestChangesetV3PlanRequest_parseChartReleaseEntries_followVersionsFromOtherChartRelease_badSelector() {
	r := ChangesetV3PlanRequest{
		ChartReleases: []ChangesetV3PlanRequestChartReleaseEntry{
			{
				ChangesetV3Create: ChangesetV3Create{
					ChartRelease: s.TestData.ChartRelease_LeonardoDev().Name,
				},
				FollowVersionsFromOtherChartRelease: utils.PointerTo("!!!"),
			},
		},
	}
	_, err := r.parseChartReleaseEntries(s.DB)
	s.ErrorContains(err, "!!!")
}

func (s *handlerSuite) TestChangesetV3PlanRequest_parseChartReleaseEntries_useExactVersionsFromOtherChartRelease_notFound() {
	r := ChangesetV3PlanRequest{
		ChartReleases: []ChangesetV3PlanRequestChartReleaseEntry{
			{
				ChangesetV3Create: ChangesetV3Create{
					ChartRelease: s.TestData.ChartRelease_LeonardoDev().Name,
				},
				UseExactVersionsFromOtherChartRelease: utils.PointerTo("not-found"),
			},
		},
	}
	_, err := r.parseChartReleaseEntries(s.DB)
	s.ErrorContains(err, "not-found")
}

func (s *handlerSuite) TestChangesetV3PlanRequest_parseChartReleaseEntries_followVersionsFromOtherChartRelease_notFound() {
	r := ChangesetV3PlanRequest{
		ChartReleases: []ChangesetV3PlanRequestChartReleaseEntry{
			{
				ChangesetV3Create: ChangesetV3Create{
					ChartRelease: s.TestData.ChartRelease_LeonardoDev().Name,
				},
				FollowVersionsFromOtherChartRelease: utils.PointerTo("not-found"),
			},
		},
	}
	_, err := r.parseChartReleaseEntries(s.DB)
	s.ErrorContains(err, "not-found")
}

func (s *handlerSuite) TestChangesetV3PlanRequest_parseChartReleaseEntries_useExactVersionsFromOtherChartRelease() {
	r := ChangesetV3PlanRequest{
		ChartReleases: []ChangesetV3PlanRequestChartReleaseEntry{
			{
				ChangesetV3Create: ChangesetV3Create{
					ChartRelease: s.TestData.ChartRelease_LeonardoDev().Name,
				},
				UseExactVersionsFromOtherChartRelease: utils.PointerTo(s.TestData.ChartRelease_LeonardoProd().Name),
			},
		},
	}
	changesets, err := r.parseChartReleaseEntries(s.DB)
	s.NoError(err)
	s.Len(changesets, 1)
	s.Equal(s.TestData.ChartRelease_LeonardoProd().AppVersionExact, changesets[0].To.AppVersionExact)
	s.Equal("exact", *changesets[0].To.AppVersionResolver)
}

func (s *handlerSuite) TestChangesetV3PlanRequest_parseChartReleaseEntries_followVersionsFromOtherChartRelease() {
	r := ChangesetV3PlanRequest{
		ChartReleases: []ChangesetV3PlanRequestChartReleaseEntry{
			{
				ChangesetV3Create: ChangesetV3Create{
					ChartRelease: s.TestData.ChartRelease_LeonardoDev().Name,
				},
				FollowVersionsFromOtherChartRelease: utils.PointerTo(s.TestData.ChartRelease_LeonardoProd().Name),
			},
		},
	}
	changesets, err := r.parseChartReleaseEntries(s.DB)
	s.NoError(err)
	s.Len(changesets, 1)
	s.Equal(s.TestData.ChartRelease_LeonardoProd().AppVersionExact, changesets[0].To.AppVersionExact)
	s.Equal("follow", *changesets[0].To.AppVersionResolver)
}

func (s *handlerSuite) TestChangesetV3PlanRequest_parseEnvironmentEntries_badEnvironmentSelector() {
	r := ChangesetV3PlanRequest{
		Environments: []ChangesetV3PlanRequestEnvironmentEntry{
			{
				Environment: "!!!",
			},
		},
	}
	_, err := r.parseEnvironmentEntries(s.DB)
	s.ErrorContains(err, "!!!")
}

func (s *handlerSuite) TestChangesetV3PlanRequest_parseEnvironmentEntries_environmentNotFound() {
	r := ChangesetV3PlanRequest{
		Environments: []ChangesetV3PlanRequestEnvironmentEntry{
			{
				Environment: "not-found",
			},
		},
	}
	_, err := r.parseEnvironmentEntries(s.DB)
	s.ErrorContains(err, "not-found")
}

func (s *handlerSuite) TestChangesetV3PlanRequest_parseEnvironmentEntries_includeChartBadSelector() {
	r := ChangesetV3PlanRequest{
		Environments: []ChangesetV3PlanRequestEnvironmentEntry{
			{
				Environment: s.TestData.Environment_Dev().Name,
				IncludeCharts: []string{
					"!!!",
				},
			},
		},
	}
	_, err := r.parseEnvironmentEntries(s.DB)
	s.ErrorContains(err, "!!!")
}

func (s *handlerSuite) TestChangesetV3PlanRequest_parseEnvironmentEntries_includeChartNotFound() {
	r := ChangesetV3PlanRequest{
		Environments: []ChangesetV3PlanRequestEnvironmentEntry{
			{
				Environment: s.TestData.Environment_Dev().Name,
				IncludeCharts: []string{
					"not-found",
				},
			},
		},
	}
	_, err := r.parseEnvironmentEntries(s.DB)
	s.ErrorContains(err, "not-found")
}

func (s *handlerSuite) TestChangesetV3PlanRequest_parseEnvironmentEntries_excludeChartBadSelector() {
	r := ChangesetV3PlanRequest{
		Environments: []ChangesetV3PlanRequestEnvironmentEntry{
			{
				Environment: s.TestData.Environment_Dev().Name,
				ExcludeCharts: []string{
					"!!!",
				},
			},
		},
	}
	_, err := r.parseEnvironmentEntries(s.DB)
	s.ErrorContains(err, "!!!")
}

func (s *handlerSuite) TestChangesetV3PlanRequest_parseEnvironmentEntries_excludeChartNotFound() {
	r := ChangesetV3PlanRequest{
		Environments: []ChangesetV3PlanRequestEnvironmentEntry{
			{
				Environment: s.TestData.Environment_Dev().Name,
				ExcludeCharts: []string{
					"not-found",
				},
			},
		},
	}
	_, err := r.parseEnvironmentEntries(s.DB)
	s.ErrorContains(err, "not-found")
}

func (s *handlerSuite) TestChangesetV3PlanRequest_parseEnvironmentEntries_excludedAll() {
	s.TestData.ChartRelease_LeonardoDev()
	s.TestData.ChartRelease_LeonardoStaging()
	r := ChangesetV3PlanRequest{
		Environments: []ChangesetV3PlanRequestEnvironmentEntry{
			{
				Environment:                          s.TestData.Environment_Staging().Name,
				UseExactVersionsFromOtherEnvironment: utils.PointerTo(s.TestData.Environment_Dev().Name),
				ExcludeCharts:                        []string{s.TestData.Chart_Leonardo().Name},
			},
		},
	}
	changesets, err := r.parseEnvironmentEntries(s.DB)
	s.NoError(err)
	s.Empty(changesets)
}

func (s *handlerSuite) TestChangesetV3PlanRequest_parseEnvironmentEntries_includedAll() {
	s.TestData.ChartRelease_LeonardoDev()
	s.TestData.ChartRelease_LeonardoStaging()
	r := ChangesetV3PlanRequest{
		Environments: []ChangesetV3PlanRequestEnvironmentEntry{
			{
				Environment:                          s.TestData.Environment_Staging().Name,
				UseExactVersionsFromOtherEnvironment: utils.PointerTo(s.TestData.Environment_Dev().Name),
				IncludeCharts:                        []string{s.TestData.Chart_Leonardo().Name},
			},
		},
	}
	changesets, err := r.parseEnvironmentEntries(s.DB)
	s.NoError(err)
	s.Len(changesets, 1)
}

func (s *handlerSuite) TestChangesetV3PlanRequest_parseEnvironmentEntries_both() {
	s.TestData.ChartRelease_LeonardoDev()
	s.TestData.ChartRelease_LeonardoStaging()
	r := ChangesetV3PlanRequest{
		Environments: []ChangesetV3PlanRequestEnvironmentEntry{
			{
				Environment:                          s.TestData.Environment_Staging().Name,
				UseExactVersionsFromOtherEnvironment: utils.PointerTo(s.TestData.Environment_Dev().Name),
				FollowVersionsFromOtherEnvironment:   utils.PointerTo(s.TestData.Environment_Dev().Name),
			},
		},
	}
	_, err := r.parseEnvironmentEntries(s.DB)
	s.ErrorContains(err, "both")
}

func (s *handlerSuite) TestChangesetV3PlanRequest_parseEnvironmentEntries_useExactVersionsFromOtherEnvironment_badSelector() {
	s.TestData.ChartRelease_LeonardoDev()
	s.TestData.ChartRelease_LeonardoStaging()
	r := ChangesetV3PlanRequest{
		Environments: []ChangesetV3PlanRequestEnvironmentEntry{
			{
				Environment:                          s.TestData.Environment_Staging().Name,
				UseExactVersionsFromOtherEnvironment: utils.PointerTo("!!!"),
			},
		},
	}
	_, err := r.parseEnvironmentEntries(s.DB)
	s.ErrorContains(err, "!!!")
}

func (s *handlerSuite) TestChangesetV3PlanRequest_parseEnvironmentEntries_useExactVersionsFromOtherEnvironment_notFound() {
	s.TestData.ChartRelease_LeonardoDev()
	s.TestData.ChartRelease_LeonardoStaging()
	r := ChangesetV3PlanRequest{
		Environments: []ChangesetV3PlanRequestEnvironmentEntry{
			{
				Environment:                          s.TestData.Environment_Staging().Name,
				UseExactVersionsFromOtherEnvironment: utils.PointerTo("not-found"),
			},
		},
	}
	_, err := r.parseEnvironmentEntries(s.DB)
	s.ErrorContains(err, "not-found")
}

func (s *handlerSuite) TestChangesetV3PlanRequest_parseEnvironmentEntries_followVersionsFromOtherEnvironment_badSelector() {
	s.TestData.ChartRelease_LeonardoDev()
	s.TestData.ChartRelease_LeonardoStaging()
	r := ChangesetV3PlanRequest{
		Environments: []ChangesetV3PlanRequestEnvironmentEntry{
			{
				Environment:                        s.TestData.Environment_Staging().Name,
				FollowVersionsFromOtherEnvironment: utils.PointerTo("!!!"),
			},
		},
	}
	_, err := r.parseEnvironmentEntries(s.DB)
	s.ErrorContains(err, "!!!")
}

func (s *handlerSuite) TestChangesetV3PlanRequest_parseEnvironmentEntries_followVersionsFromOtherEnvironment_notFound() {
	s.TestData.ChartRelease_LeonardoDev()
	s.TestData.ChartRelease_LeonardoStaging()
	r := ChangesetV3PlanRequest{
		Environments: []ChangesetV3PlanRequestEnvironmentEntry{
			{
				Environment:                        s.TestData.Environment_Staging().Name,
				FollowVersionsFromOtherEnvironment: utils.PointerTo("not-found"),
			},
		},
	}
	_, err := r.parseEnvironmentEntries(s.DB)
	s.ErrorContains(err, "not-found")
}

func (s *handlerSuite) TestChangesetV3PlanRequest_parseEnvironmentEntries_useExactVersionsFromOtherEnvironment() {
	s.TestData.ChartRelease_LeonardoDev()
	s.TestData.ChartRelease_LeonardoStaging()
	r := ChangesetV3PlanRequest{
		Environments: []ChangesetV3PlanRequestEnvironmentEntry{
			{
				Environment:                          s.TestData.Environment_Staging().Name,
				UseExactVersionsFromOtherEnvironment: utils.PointerTo(s.TestData.Environment_Dev().Name),
			},
		},
	}
	changesets, err := r.parseEnvironmentEntries(s.DB)
	s.NoError(err)
	s.Len(changesets, 1)
	s.Equal("exact", *changesets[0].To.AppVersionResolver)
	s.Equal(s.TestData.ChartRelease_LeonardoDev().AppVersionExact, changesets[0].To.AppVersionExact)
}

func (s *handlerSuite) TestChangesetV3PlanRequest_parseEnvironmentEntries_followVersionsFromOtherEnvironment() {
	s.TestData.ChartRelease_LeonardoDev()
	s.TestData.ChartRelease_LeonardoStaging()
	r := ChangesetV3PlanRequest{
		Environments: []ChangesetV3PlanRequestEnvironmentEntry{
			{
				Environment:                        s.TestData.Environment_Staging().Name,
				FollowVersionsFromOtherEnvironment: utils.PointerTo(s.TestData.Environment_Dev().Name),
			},
		},
	}
	changesets, err := r.parseEnvironmentEntries(s.DB)
	s.NoError(err)
	s.Len(changesets, 1)
	s.Equal(s.TestData.ChartRelease_LeonardoDev().ID, *changesets[0].To.AppVersionFollowChartReleaseID)
	s.Equal("follow", *changesets[0].To.AppVersionResolver)
	s.Equal(s.TestData.ChartRelease_LeonardoDev().AppVersionExact, changesets[0].To.AppVersionExact)
}
