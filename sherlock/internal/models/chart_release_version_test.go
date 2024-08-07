package models

import (
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func (s *modelSuite) TestChartReleaseVersion_resolve_follow() {
	s.SetSuitableTestUserForDB()
	chartRelease := s.TestData.ChartRelease_LeonardoDev()
	s.NotNil(chartRelease.AppVersionResolver)
	s.NotNil(chartRelease.AppVersionBranch)
	s.NotNil(chartRelease.AppVersionCommit)
	s.NotNil(chartRelease.AppVersionExact)
	s.NotNil(chartRelease.AppVersionID)
	s.NotNil(chartRelease.ChartVersionResolver)
	s.NotNil(chartRelease.ChartVersionExact)
	s.NotNil(chartRelease.ChartVersionID)
	s.NotNil(chartRelease.HelmfileRef)
	s.NotNil(chartRelease.HelmfileRefEnabled)
	chartReleaseVersion := ChartReleaseVersion{
		AppVersionResolver:               utils.PointerTo("follow"),
		AppVersionFollowChartReleaseID:   utils.PointerTo(chartRelease.ID),
		ChartVersionResolver:             utils.PointerTo("follow"),
		ChartVersionFollowChartReleaseID: utils.PointerTo(chartRelease.ID),
	}
	err := chartReleaseVersion.resolve(s.DB, s.TestData.Chart_Leonardo().ID)
	s.NoError(err)
	s.Equal(chartRelease.AppVersionBranch, chartReleaseVersion.AppVersionBranch)
	s.Equal(chartRelease.AppVersionCommit, chartReleaseVersion.AppVersionCommit)
	s.Equal(chartRelease.AppVersionExact, chartReleaseVersion.AppVersionExact)
	s.Equal(chartRelease.AppVersionID, chartReleaseVersion.AppVersionID)
	s.Equal(chartRelease.ChartVersionExact, chartReleaseVersion.ChartVersionExact)
	s.Equal(chartRelease.ChartVersionID, chartReleaseVersion.ChartVersionID)
	s.Equal(chartRelease.HelmfileRef, chartReleaseVersion.HelmfileRef)
	s.Equal(chartRelease.HelmfileRefEnabled, chartReleaseVersion.HelmfileRefEnabled)
}

func (s *modelSuite) TestChartReleaseVersion_resolve_exactHotfix() {
	s.SetSuitableTestUserForDB()
	chartReleaseVersion := ChartReleaseVersion{
		AppVersionResolver:   utils.PointerTo("exact"),
		AppVersionExact:      utils.PointerTo("v1.2.4"),
		ChartVersionResolver: utils.PointerTo("exact"),
		ChartVersionExact:    utils.PointerTo("v1.2.3"),
		HelmfileRefEnabled:   utils.PointerTo(true),
		HelmfileRef:          utils.PointerTo("v1.2.2"),
	}
	err := chartReleaseVersion.resolve(s.DB, s.TestData.Chart_Leonardo().ID)
	s.NoError(err)
	s.Equal("exact", *chartReleaseVersion.AppVersionResolver)
	s.Nil(chartReleaseVersion.AppVersionBranch)
	s.Nil(chartReleaseVersion.AppVersionCommit)
	s.Equal("v1.2.4", *chartReleaseVersion.AppVersionExact)
	s.Nil(chartReleaseVersion.AppVersionID)
	s.Equal("exact", *chartReleaseVersion.ChartVersionResolver)
	s.Equal("v1.2.3", *chartReleaseVersion.ChartVersionExact)
	s.Nil(chartReleaseVersion.ChartVersionID)
	s.Equal("v1.2.2", *chartReleaseVersion.HelmfileRef)
	s.True(*chartReleaseVersion.HelmfileRefEnabled)
}

func (s *modelSuite) TestChartReleaseVersion_resolve_exact() {
	s.SetSuitableTestUserForDB()
	chartReleaseVersion := ChartReleaseVersion{
		AppVersionResolver:   utils.PointerTo("exact"),
		AppVersionExact:      utils.PointerTo(s.TestData.AppVersion_Leonardo_V2().AppVersion),
		ChartVersionResolver: utils.PointerTo("exact"),
		ChartVersionExact:    utils.PointerTo(s.TestData.ChartVersion_Leonardo_V2().ChartVersion),
	}
	err := chartReleaseVersion.resolve(s.DB, s.TestData.Chart_Leonardo().ID)
	s.NoError(err)
	s.Equal("exact", *chartReleaseVersion.AppVersionResolver)
	s.Equal(s.TestData.AppVersion_Leonardo_V2().GitBranch, *chartReleaseVersion.AppVersionBranch)
	s.Equal(s.TestData.AppVersion_Leonardo_V2().GitCommit, *chartReleaseVersion.AppVersionCommit)
	s.Equal(s.TestData.AppVersion_Leonardo_V2().AppVersion, *chartReleaseVersion.AppVersionExact)
	s.Equal(s.TestData.AppVersion_Leonardo_V2().ID, *chartReleaseVersion.AppVersionID)
	s.Equal("exact", *chartReleaseVersion.ChartVersionResolver)
	s.Equal(s.TestData.ChartVersion_Leonardo_V2().ChartVersion, *chartReleaseVersion.ChartVersionExact)
	s.Equal(s.TestData.ChartVersion_Leonardo_V2().ID, *chartReleaseVersion.ChartVersionID)
	s.False(*chartReleaseVersion.HelmfileRefEnabled)
	s.Equal("charts/"+s.TestData.Chart_Leonardo().Name+"-"+s.TestData.ChartVersion_Leonardo_V2().ChartVersion, *chartReleaseVersion.HelmfileRef)
}

func (s *modelSuite) TestChartReleaseVersion_resolve_latest() {
	s.SetSuitableTestUserForDB()
	s.TestData.AppVersion_Leonardo_V1()
	s.TestData.ChartVersion_Leonardo_V1()
	chartReleaseVersion := ChartReleaseVersion{
		AppVersionResolver:   utils.PointerTo("branch"),
		AppVersionBranch:     utils.PointerTo(s.TestData.AppVersion_Leonardo_V1().GitBranch),
		ChartVersionResolver: utils.PointerTo("latest"),
	}
	err := chartReleaseVersion.resolve(s.DB, s.TestData.Chart_Leonardo().ID)
	s.NoError(err)
	s.Equal("branch", *chartReleaseVersion.AppVersionResolver)
	s.Equal(s.TestData.AppVersion_Leonardo_V1().GitBranch, *chartReleaseVersion.AppVersionBranch)
	s.Equal(s.TestData.AppVersion_Leonardo_V1().GitCommit, *chartReleaseVersion.AppVersionCommit)
	s.Equal(s.TestData.AppVersion_Leonardo_V1().AppVersion, *chartReleaseVersion.AppVersionExact)
	s.Equal(s.TestData.AppVersion_Leonardo_V1().ID, *chartReleaseVersion.AppVersionID)
	s.Nil(chartReleaseVersion.AppVersionFollowChartReleaseID)
	s.Equal("latest", *chartReleaseVersion.ChartVersionResolver)
	s.Equal(s.TestData.ChartVersion_Leonardo_V1().ChartVersion, *chartReleaseVersion.ChartVersionExact)
	s.Equal(s.TestData.ChartVersion_Leonardo_V1().ID, *chartReleaseVersion.ChartVersionID)
	s.False(*chartReleaseVersion.HelmfileRefEnabled)
	s.Equal("charts/"+s.TestData.Chart_Leonardo().Name+"-"+s.TestData.ChartVersion_Leonardo_V1().ChartVersion, *chartReleaseVersion.HelmfileRef)
}

func (s *modelSuite) TestChartReleaseVersion_resolveAppVersion_failToFindBranch() {
	s.SetSuitableTestUserForDB()
	chartReleaseVersion := ChartReleaseVersion{
		AppVersionResolver: utils.PointerTo("branch"),
		AppVersionBranch:   utils.PointerTo("nonexistent"),
	}
	err := chartReleaseVersion.resolve(s.DB, s.TestData.Chart_Leonardo().ID)
	s.ErrorContains(err, "no recorded app versions for leonardo come from a 'nonexistent' branch")
}

func (s *modelSuite) TestChartReleaseVersion_resolveAppVersion_commit() {
	s.SetSuitableTestUserForDB()
	s.TestData.AppVersion_Leonardo_V1()
	chartReleaseVersion := ChartReleaseVersion{
		AppVersionResolver: utils.PointerTo("commit"),
		AppVersionCommit:   utils.PointerTo(s.TestData.AppVersion_Leonardo_V1().GitCommit),
	}
	err := chartReleaseVersion.resolveAppVersion(s.DB, s.TestData.Chart_Leonardo())
	s.NoError(err)
	s.Equal("commit", *chartReleaseVersion.AppVersionResolver)
	s.Equal(s.TestData.AppVersion_Leonardo_V1().GitBranch, *chartReleaseVersion.AppVersionBranch)
	s.Equal(s.TestData.AppVersion_Leonardo_V1().GitCommit, *chartReleaseVersion.AppVersionCommit)
	s.Equal(s.TestData.AppVersion_Leonardo_V1().AppVersion, *chartReleaseVersion.AppVersionExact)
	s.Equal(s.TestData.AppVersion_Leonardo_V1().ID, *chartReleaseVersion.AppVersionID)
	s.Nil(chartReleaseVersion.AppVersionFollowChartReleaseID)
}

func (s *modelSuite) TestChartReleaseVersion_resolveAppVersion_commitPartial() {
	s.SetSuitableTestUserForDB()
	s.TestData.AppVersion_Leonardo_V1()
	chartReleaseVersion := ChartReleaseVersion{
		AppVersionResolver: utils.PointerTo("commit"),
		AppVersionCommit:   utils.PointerTo(s.TestData.AppVersion_Leonardo_V1().GitCommit[:3]),
	}
	err := chartReleaseVersion.resolveAppVersion(s.DB, s.TestData.Chart_Leonardo())
	s.NoError(err)
	s.Equal("commit", *chartReleaseVersion.AppVersionResolver)
	s.Equal(s.TestData.AppVersion_Leonardo_V1().GitBranch, *chartReleaseVersion.AppVersionBranch)
	s.Equal(s.TestData.AppVersion_Leonardo_V1().GitCommit, *chartReleaseVersion.AppVersionCommit)
	s.Equal(s.TestData.AppVersion_Leonardo_V1().AppVersion, *chartReleaseVersion.AppVersionExact)
	s.Equal(s.TestData.AppVersion_Leonardo_V1().ID, *chartReleaseVersion.AppVersionID)
	s.Nil(chartReleaseVersion.AppVersionFollowChartReleaseID)
}

func (s *modelSuite) TestChartReleaseVersion_resolveAppVersion_failToFindCommit() {
	s.SetSuitableTestUserForDB()
	chartReleaseVersion := ChartReleaseVersion{
		AppVersionResolver: utils.PointerTo("commit"),
		AppVersionCommit:   utils.PointerTo("nonexistent"),
	}
	err := chartReleaseVersion.resolve(s.DB, s.TestData.Chart_Leonardo().ID)
	s.ErrorContains(err, "no recorded app versions for leonardo have a commit starting with 'nonexistent'")
}

func (s *modelSuite) TestChartReleaseVersion_resolveAppVersion_branchNoBranch() {
	chartReleaseVersion := ChartReleaseVersion{
		AppVersionResolver: utils.PointerTo("branch"),
	}
	err := chartReleaseVersion.resolveAppVersion(s.DB, s.TestData.Chart_Leonardo())
	s.ErrorContains(err, "appVersionResolver was set to 'branch' but no app branch was supplied")
}

func (s *modelSuite) TestChartReleaseVersion_resolveAppVersion_commitNoCommit() {
	chartReleaseVersion := ChartReleaseVersion{
		AppVersionResolver: utils.PointerTo("commit"),
	}
	err := chartReleaseVersion.resolveAppVersion(s.DB, s.TestData.Chart_Leonardo())
	s.ErrorContains(err, "appVersionResolver was set to 'commit' but no app commit was supplied")
}

func (s *modelSuite) TestChartReleaseVersion_resolveAppVersion_exactNoExact() {
	s.SetSuitableTestUserForDB()
	chartReleaseVersion := ChartReleaseVersion{
		AppVersionResolver: utils.PointerTo("exact"),
	}
	err := chartReleaseVersion.resolveAppVersion(s.DB, s.TestData.Chart_Leonardo())
	s.ErrorContains(err, "appVersionResolver was set to 'exact' but no exact chart version was supplied")
}

func (s *modelSuite) TestChartReleaseVersion_resolveAppVersion_followNoID() {
	chartReleaseVersion := ChartReleaseVersion{
		AppVersionResolver: utils.PointerTo("follow"),
	}
	err := chartReleaseVersion.resolveAppVersion(s.DB, s.TestData.Chart_Leonardo())
	s.ErrorContains(err, "appVersionResolver was set to 'follow' but no chart release ID was given to follow")
}

func (s *modelSuite) TestChartReleaseVersion_resolveChartVersion_failToFindLatest() {
	chartReleaseVersion := ChartReleaseVersion{
		ChartVersionResolver: utils.PointerTo("latest"),
	}
	err := chartReleaseVersion.resolveChartVersion(s.DB, s.TestData.Chart_Leonardo())
	s.ErrorContains(err, "no recorded chart versions for leonardo")
}

func (s *modelSuite) TestChartReleaseVersion_resolveChartVersion_exactNoExact() {
	s.SetSuitableTestUserForDB()
	chartReleaseVersion := ChartReleaseVersion{
		ChartVersionResolver: utils.PointerTo("exact"),
	}
	err := chartReleaseVersion.resolveChartVersion(s.DB, s.TestData.Chart_Leonardo())
	s.ErrorContains(err, "chartVersionResolver was set to 'exact' but no exact chart version was supplied")
}

func (s *modelSuite) TestChartReleaseVersion_resolveChartVersion_followNoID() {
	chartReleaseVersion := ChartReleaseVersion{
		ChartVersionResolver: utils.PointerTo("follow"),
	}
	err := chartReleaseVersion.resolveChartVersion(s.DB, s.TestData.Chart_Leonardo())
	s.ErrorContains(err, "chartVersionResolver was set to 'follow' but no chart release ID was given to follow")
}

func TestChartReleaseVersion_ClearAppVersion(t *testing.T) {
	type fields struct {
		ResolvedAt                       *time.Time
		AppVersionResolver               *string
		AppVersionExact                  *string
		AppVersionBranch                 *string
		AppVersionCommit                 *string
		AppVersionFollowChartReleaseID   *uint
		AppVersionID                     *uint
		ChartVersionResolver             *string
		ChartVersionExact                *string
		ChartVersionFollowChartReleaseID *uint
		ChartVersionID                   *uint
		HelmfileRef                      *string
		HelmfileRefEnabled               *bool
	}
	tests := []struct {
		name   string
		fields fields
		want   fields
	}{
		{
			name: "all fields set",
			fields: fields{
				AppVersionResolver:               utils.PointerTo("exact"),
				AppVersionExact:                  utils.PointerTo("v1.2.4"),
				AppVersionBranch:                 utils.PointerTo("main"),
				AppVersionCommit:                 utils.PointerTo("1234567"),
				AppVersionFollowChartReleaseID:   utils.PointerTo(uint(1)),
				AppVersionID:                     utils.PointerTo(uint(1)),
				ChartVersionResolver:             utils.PointerTo("exact"),
				ChartVersionExact:                utils.PointerTo("v1.2.4"),
				ChartVersionFollowChartReleaseID: utils.PointerTo(uint(1)),
				ChartVersionID:                   utils.PointerTo(uint(1)),
				HelmfileRef:                      utils.PointerTo("v1.2.4"),
				HelmfileRefEnabled:               utils.PointerTo(true),
			},
			want: fields{
				AppVersionResolver:               nil,
				AppVersionExact:                  nil,
				AppVersionBranch:                 nil,
				AppVersionCommit:                 nil,
				AppVersionFollowChartReleaseID:   nil,
				AppVersionID:                     nil,
				ChartVersionResolver:             utils.PointerTo("exact"),
				ChartVersionExact:                utils.PointerTo("v1.2.4"),
				ChartVersionFollowChartReleaseID: utils.PointerTo(uint(1)),
				ChartVersionID:                   utils.PointerTo(uint(1)),
				HelmfileRef:                      utils.PointerTo("v1.2.4"),
				HelmfileRefEnabled:               utils.PointerTo(true),
			},
		},
		{
			name:   "no fields set",
			fields: fields{},
			want:   fields{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			crv := &ChartReleaseVersion{
				ResolvedAt:                       tt.fields.ResolvedAt,
				AppVersionResolver:               tt.fields.AppVersionResolver,
				AppVersionExact:                  tt.fields.AppVersionExact,
				AppVersionBranch:                 tt.fields.AppVersionBranch,
				AppVersionCommit:                 tt.fields.AppVersionCommit,
				AppVersionFollowChartReleaseID:   tt.fields.AppVersionFollowChartReleaseID,
				AppVersionID:                     tt.fields.AppVersionID,
				ChartVersionResolver:             tt.fields.ChartVersionResolver,
				ChartVersionExact:                tt.fields.ChartVersionExact,
				ChartVersionFollowChartReleaseID: tt.fields.ChartVersionFollowChartReleaseID,
				ChartVersionID:                   tt.fields.ChartVersionID,
				HelmfileRef:                      tt.fields.HelmfileRef,
				HelmfileRefEnabled:               tt.fields.HelmfileRefEnabled,
			}
			crv.ClearAppVersion()
			assert.Equal(t, &ChartReleaseVersion{
				ResolvedAt:                       tt.want.ResolvedAt,
				AppVersionResolver:               tt.want.AppVersionResolver,
				AppVersionExact:                  tt.want.AppVersionExact,
				AppVersionBranch:                 tt.want.AppVersionBranch,
				AppVersionCommit:                 tt.want.AppVersionCommit,
				AppVersionFollowChartReleaseID:   tt.want.AppVersionFollowChartReleaseID,
				AppVersionID:                     tt.want.AppVersionID,
				ChartVersionResolver:             tt.want.ChartVersionResolver,
				ChartVersionExact:                tt.want.ChartVersionExact,
				ChartVersionFollowChartReleaseID: tt.want.ChartVersionFollowChartReleaseID,
				ChartVersionID:                   tt.want.ChartVersionID,
				HelmfileRef:                      tt.want.HelmfileRef,
				HelmfileRefEnabled:               tt.want.HelmfileRefEnabled,
			}, crv)
		})
	}
}
