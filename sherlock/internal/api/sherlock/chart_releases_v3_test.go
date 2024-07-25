package sherlock

import (
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
	"time"
)

func (s *handlerSuite) TestChartReleaseV3_toModel() {
	now := time.Now()
	chart := s.TestData.Chart_Leonardo()
	cluster := s.TestData.Cluster_TerraProd()
	environment := s.TestData.Environment_Prod()
	chartReleaseDev := s.TestData.ChartRelease_LeonardoDev()
	chartReleaseStaging := s.TestData.ChartRelease_LeonardoStaging()
	appVersion := s.TestData.AppVersion_Leonardo_V1()
	chartVersion := s.TestData.ChartVersion_Leonardo_V1()
	pagerdutyIntegration := s.TestData.PagerdutyIntegration_ManuallyTriggeredTerraIncident()
	type fields struct {
		CommonFields          CommonFields
		AppVersionReference   string
		ChartVersionReference string
		DestinationType       string
		ResolvedAt            *time.Time
		ChartReleaseV3Create  ChartReleaseV3Create
	}
	tests := []struct {
		name    string
		fields  fields
		want    models.ChartRelease
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:    "empty",
			fields:  fields{},
			wantErr: assert.NoError,
			want:    models.ChartRelease{},
		},
		{
			name: "invalid chart selector",
			fields: fields{
				ChartReleaseV3Create: ChartReleaseV3Create{
					Chart: "!!!!",
				},
			},
			wantErr: assert.Error,
		},
		{
			name: "not found chart",
			fields: fields{
				ChartReleaseV3Create: ChartReleaseV3Create{
					Chart: "not-found",
				},
			},
			wantErr: assert.Error,
		},
		{
			name: "invalid cluster selector",
			fields: fields{
				ChartReleaseV3Create: ChartReleaseV3Create{
					Cluster: "!!!!!",
				},
			},
			wantErr: assert.Error,
		},
		{
			name: "not found cluster",
			fields: fields{
				ChartReleaseV3Create: ChartReleaseV3Create{
					Cluster: "not-found",
				},
			},
			wantErr: assert.Error,
		},
		{
			name: "invalid environment selector",
			fields: fields{
				ChartReleaseV3Create: ChartReleaseV3Create{
					Environment: "!!!!",
				},
			},
			wantErr: assert.Error,
		},
		{
			name: "not found environment",
			fields: fields{
				ChartReleaseV3Create: ChartReleaseV3Create{
					Environment: "not-found",
				},
			},
			wantErr: assert.Error,
		},
		{
			name: "invalid chart release for app version follow",
			fields: fields{
				ChartReleaseV3Create: ChartReleaseV3Create{
					AppVersionFollowChartRelease: "!!!!!",
				},
			},
			wantErr: assert.Error,
		},
		{
			name: "not found chart release for app version follow",
			fields: fields{
				ChartReleaseV3Create: ChartReleaseV3Create{
					AppVersionFollowChartRelease: "not-found",
				},
			},
			wantErr: assert.Error,
		},
		{
			name: "invalid chart release for chart version follow",
			fields: fields{
				ChartReleaseV3Create: ChartReleaseV3Create{
					ChartVersionFollowChartRelease: "!!!!!",
				},
			},
			wantErr: assert.Error,
		},
		{
			name: "not found chart release for chart version follow",
			fields: fields{
				ChartReleaseV3Create: ChartReleaseV3Create{
					ChartVersionFollowChartRelease: "not-found",
				},
			},
			wantErr: assert.Error,
		},
		{
			name: "invalid app version reference",
			fields: fields{
				AppVersionReference: "!!!!",
			},
			wantErr: assert.Error,
		},
		{
			name: "not found app version",
			fields: fields{
				AppVersionReference: "leonardo/not-found",
			},
			wantErr: assert.Error,
		},
		{
			name: "invalid chart version reference",
			fields: fields{
				ChartVersionReference: "!!!!",
			},
			wantErr: assert.Error,
		},
		{
			name: "not found chart version",
			fields: fields{
				ChartVersionReference: "leonardo/not-found",
			},
			wantErr: assert.Error,
		},
		{
			name: "invalid pagerduty integration",
			fields: fields{
				ChartReleaseV3Create: ChartReleaseV3Create{
					ChartReleaseV3Edit: ChartReleaseV3Edit{
						PagerdutyIntegration: utils.PointerTo("!!!!!"),
					},
				},
			},
			wantErr: assert.Error,
		},
		{
			name: "not found pagerduty integration",
			fields: fields{
				ChartReleaseV3Create: ChartReleaseV3Create{
					ChartReleaseV3Edit: ChartReleaseV3Edit{
						PagerdutyIntegration: utils.PointerTo("pd-id/not-found"),
					},
				},
			},
			wantErr: assert.Error,
		},
		{
			name: "valid",
			fields: fields{
				CommonFields: CommonFields{
					ID:        1,
					CreatedAt: now.Add(-time.Hour),
					UpdatedAt: now.Add(-time.Minute),
				},
				AppVersionReference:   fmt.Sprintf("leonardo/%s", appVersion.AppVersion),
				ChartVersionReference: fmt.Sprintf("leonardo/%s", chartVersion.ChartVersion),
				DestinationType:       "environment",
				ResolvedAt:            &now,
				ChartReleaseV3Create: ChartReleaseV3Create{
					Chart:                          chart.Name,
					Cluster:                        cluster.Name,
					Environment:                    environment.Name,
					Name:                           "name",
					Namespace:                      "namespace",
					AppVersionResolver:             utils.PointerTo("exact"),
					AppVersionExact:                utils.PointerTo("v1.0.0"),
					AppVersionBranch:               utils.PointerTo("branch"),
					AppVersionCommit:               utils.PointerTo("commit"),
					AppVersionFollowChartRelease:   chartReleaseDev.Name,
					ChartVersionResolver:           utils.PointerTo("exact"),
					ChartVersionExact:              utils.PointerTo("2.0.0"),
					ChartVersionFollowChartRelease: chartReleaseStaging.Name,
					HelmfileRef:                    utils.PointerTo("HEAD"),
					HelmfileRefEnabled:             utils.PointerTo(true),
					ChartReleaseV3Edit: ChartReleaseV3Edit{
						Subdomain:               utils.PointerTo("subdomain"),
						Protocol:                utils.PointerTo("https"),
						Port:                    utils.PointerTo[uint](443),
						PagerdutyIntegration:    utils.PointerTo(utils.UintToString(pagerdutyIntegration.ID)),
						IncludeInBulkChangesets: utils.PointerTo(true),
					},
				},
			},
			wantErr: assert.NoError,
			want: models.ChartRelease{
				Model: gorm.Model{
					ID:        1,
					CreatedAt: now.Add(-time.Hour),
					UpdatedAt: now.Add(-time.Minute),
				},
				ChartID:         chart.ID,
				ClusterID:       &cluster.ID,
				DestinationType: "environment",
				EnvironmentID:   &environment.ID,
				Name:            "name",
				Namespace:       "namespace",
				ChartReleaseVersion: models.ChartReleaseVersion{
					ResolvedAt:                       &now,
					AppVersionResolver:               utils.PointerTo("exact"),
					AppVersionExact:                  utils.PointerTo("v1.0.0"),
					AppVersionBranch:                 utils.PointerTo("branch"),
					AppVersionCommit:                 utils.PointerTo("commit"),
					AppVersionFollowChartReleaseID:   &chartReleaseDev.ID,
					AppVersionID:                     &appVersion.ID,
					ChartVersionResolver:             utils.PointerTo("exact"),
					ChartVersionExact:                utils.PointerTo("2.0.0"),
					ChartVersionFollowChartReleaseID: &chartReleaseStaging.ID,
					ChartVersionID:                   &chartVersion.ID,
					HelmfileRef:                      utils.PointerTo("HEAD"),
					HelmfileRefEnabled:               utils.PointerTo(true),
				},
				Subdomain:               utils.PointerTo("subdomain"),
				Protocol:                utils.PointerTo("https"),
				Port:                    utils.PointerTo[uint](443),
				PagerdutyIntegrationID:  &pagerdutyIntegration.ID,
				IncludeInBulkChangesets: utils.PointerTo(true),
			},
		},
		{
			name: "full with whitespace",
			fields: fields{
				CommonFields: CommonFields{
					ID:        1,
					CreatedAt: now.Add(-time.Hour),
					UpdatedAt: now.Add(-time.Minute),
				},
				AppVersionReference:   fmt.Sprintf("leonardo/%s", appVersion.AppVersion),
				ChartVersionReference: fmt.Sprintf("leonardo/%s", chartVersion.ChartVersion),
				DestinationType:       "environment",
				ResolvedAt:            &now,
				ChartReleaseV3Create: ChartReleaseV3Create{
					Chart:                          chart.Name,
					Cluster:                        cluster.Name,
					Environment:                    environment.Name,
					Name:                           "name",
					Namespace:                      "namespace",
					AppVersionResolver:             utils.PointerTo("exact"),
					AppVersionExact:                utils.PointerTo(" v1.0.0 "),
					AppVersionBranch:               utils.PointerTo(" branch "),
					AppVersionCommit:               utils.PointerTo(" commit "),
					AppVersionFollowChartRelease:   chartReleaseDev.Name,
					ChartVersionResolver:           utils.PointerTo("exact"),
					ChartVersionExact:              utils.PointerTo(" 2.0.0 "),
					ChartVersionFollowChartRelease: chartReleaseStaging.Name,
					HelmfileRef:                    utils.PointerTo(" HEAD "),
					HelmfileRefEnabled:             utils.PointerTo(true),
					ChartReleaseV3Edit: ChartReleaseV3Edit{
						Subdomain:               utils.PointerTo("subdomain"),
						Protocol:                utils.PointerTo("https"),
						Port:                    utils.PointerTo[uint](443),
						PagerdutyIntegration:    utils.PointerTo(utils.UintToString(pagerdutyIntegration.ID)),
						IncludeInBulkChangesets: utils.PointerTo(true),
					},
				},
			},
			wantErr: assert.NoError,
			want: models.ChartRelease{
				Model: gorm.Model{
					ID:        1,
					CreatedAt: now.Add(-time.Hour),
					UpdatedAt: now.Add(-time.Minute),
				},
				ChartID:         chart.ID,
				ClusterID:       &cluster.ID,
				DestinationType: "environment",
				EnvironmentID:   &environment.ID,
				Name:            "name",
				Namespace:       "namespace",
				ChartReleaseVersion: models.ChartReleaseVersion{
					ResolvedAt:                       &now,
					AppVersionResolver:               utils.PointerTo("exact"),
					AppVersionExact:                  utils.PointerTo("v1.0.0"),
					AppVersionBranch:                 utils.PointerTo("branch"),
					AppVersionCommit:                 utils.PointerTo("commit"),
					AppVersionFollowChartReleaseID:   &chartReleaseDev.ID,
					AppVersionID:                     &appVersion.ID,
					ChartVersionResolver:             utils.PointerTo("exact"),
					ChartVersionExact:                utils.PointerTo("2.0.0"),
					ChartVersionFollowChartReleaseID: &chartReleaseStaging.ID,
					ChartVersionID:                   &chartVersion.ID,
					HelmfileRef:                      utils.PointerTo("HEAD"),
					HelmfileRefEnabled:               utils.PointerTo(true),
				},
				Subdomain:               utils.PointerTo("subdomain"),
				Protocol:                utils.PointerTo("https"),
				Port:                    utils.PointerTo[uint](443),
				PagerdutyIntegrationID:  &pagerdutyIntegration.ID,
				IncludeInBulkChangesets: utils.PointerTo(true),
			},
		},
	}
	for _, tt := range tests {
		s.Run(tt.name, func() {
			c := ChartReleaseV3{
				CommonFields:          tt.fields.CommonFields,
				AppVersionReference:   tt.fields.AppVersionReference,
				ChartVersionReference: tt.fields.ChartVersionReference,
				DestinationType:       tt.fields.DestinationType,
				ResolvedAt:            tt.fields.ResolvedAt,
				ChartReleaseV3Create:  tt.fields.ChartReleaseV3Create,
			}
			got, err := c.toModel(s.DB)
			if !tt.wantErr(s.T(), err, "toModel()") {
				return
			}
			s.Equalf(tt.want, got, "toModel()")
		})
	}
}

func Test_chartReleaseFromModel(t *testing.T) {
	config.LoadTestConfig()
	now := time.Now()
	type args struct {
		model models.ChartRelease
	}
	tests := []struct {
		name string
		args args
		want ChartReleaseV3
	}{
		{
			name: "empty",
			args: args{},
			want: ChartReleaseV3{},
		},
		{
			name: "full",
			args: args{model: models.ChartRelease{
				Model: gorm.Model{
					ID:        1,
					CreatedAt: now.Add(-time.Hour),
					UpdatedAt: now.Add(-time.Minute),
				},
				CiIdentifier:            &models.CiIdentifier{Model: gorm.Model{ID: 2}},
				Chart:                   &models.Chart{Model: gorm.Model{ID: 3}, Name: "leonardo"},
				ChartID:                 3,
				Cluster:                 &models.Cluster{Model: gorm.Model{ID: 4}, Name: "terra-prod"},
				ClusterID:               utils.PointerTo[uint](4),
				DestinationType:         "environment",
				Environment:             &models.Environment{Model: gorm.Model{ID: 5}, Name: "prod"},
				EnvironmentID:           utils.PointerTo[uint](5),
				Name:                    "name",
				Namespace:               "namespace",
				Subdomain:               utils.PointerTo("subdomain"),
				Protocol:                utils.PointerTo("https"),
				Port:                    utils.PointerTo[uint](443),
				PagerdutyIntegration:    &models.PagerdutyIntegration{Model: gorm.Model{ID: 6}, PagerdutyID: "blah"},
				PagerdutyIntegrationID:  utils.PointerTo[uint](6),
				IncludeInBulkChangesets: utils.PointerTo(true),
				ChartReleaseVersion: models.ChartReleaseVersion{
					ResolvedAt:                       &now,
					AppVersionResolver:               utils.PointerTo("exact"),
					AppVersionExact:                  utils.PointerTo("v1.0.0"),
					AppVersionBranch:                 utils.PointerTo("branch"),
					AppVersionCommit:                 utils.PointerTo("commit"),
					AppVersionFollowChartRelease:     &models.ChartRelease{Model: gorm.Model{ID: 7}, Name: "leonardo-dev"},
					AppVersionFollowChartReleaseID:   utils.PointerTo[uint](7),
					AppVersion:                       &models.AppVersion{Model: gorm.Model{ID: 8}, AppVersion: "v1.0.0"},
					AppVersionID:                     utils.PointerTo[uint](8),
					ChartVersionResolver:             utils.PointerTo("exact"),
					ChartVersionExact:                utils.PointerTo("2.0.0"),
					ChartVersionFollowChartRelease:   &models.ChartRelease{Model: gorm.Model{ID: 9}, Name: "leonardo-staging"},
					ChartVersionFollowChartReleaseID: utils.PointerTo[uint](9),
					ChartVersion:                     &models.ChartVersion{Model: gorm.Model{ID: 10}, ChartVersion: "2.0.0"},
					ChartVersionID:                   utils.PointerTo[uint](10),
					HelmfileRef:                      utils.PointerTo("HEAD"),
					HelmfileRefEnabled:               utils.PointerTo(true),
				},
			}},
			want: ChartReleaseV3{
				CommonFields: CommonFields{
					ID:        1,
					CreatedAt: now.Add(-time.Hour),
					UpdatedAt: now.Add(-time.Minute),
				},
				CiIdentifier:             &CiIdentifierV3{CommonFields: CommonFields{ID: 2}},
				ChartInfo:                &ChartV3{CommonFields: CommonFields{ID: 3}, ChartV3Create: ChartV3Create{Name: "leonardo"}},
				ClusterInfo:              &ClusterV3{CommonFields: CommonFields{ID: 4}, ClusterV3Create: ClusterV3Create{Name: "terra-prod", ClusterV3Edit: ClusterV3Edit{RequiredRole: utils.PointerTo(config.Config.String("model.roles.substituteEmptyRequiredRoleWithValue"))}}},
				EnvironmentInfo:          &EnvironmentV3{CommonFields: CommonFields{ID: 5}, EnvironmentV3Create: EnvironmentV3Create{Name: "prod", EnvironmentV3Edit: EnvironmentV3Edit{RequiredRole: utils.PointerTo(config.Config.String("model.roles.substituteEmptyRequiredRoleWithValue"))}}},
				AppVersionReference:      "leonardo/v1.0.0",
				AppVersionInfo:           &AppVersionV3{CommonFields: CommonFields{ID: 8}, AppVersionV3Create: AppVersionV3Create{AppVersion: "v1.0.0"}},
				ChartVersionReference:    "leonardo/2.0.0",
				ChartVersionInfo:         &ChartVersionV3{CommonFields: CommonFields{ID: 10}, ChartVersionV3Create: ChartVersionV3Create{ChartVersion: "2.0.0"}},
				PagerdutyIntegrationInfo: &PagerdutyIntegrationV3{CommonFields: CommonFields{ID: 6}, PagerdutyID: "blah"},
				DestinationType:          "environment",
				ResolvedAt:               &now,
				ChartReleaseV3Create: ChartReleaseV3Create{
					Chart:                          "leonardo",
					Cluster:                        "terra-prod",
					Environment:                    "prod",
					Name:                           "name",
					Namespace:                      "namespace",
					AppVersionResolver:             utils.PointerTo("exact"),
					AppVersionExact:                utils.PointerTo("v1.0.0"),
					AppVersionBranch:               utils.PointerTo("branch"),
					AppVersionCommit:               utils.PointerTo("commit"),
					AppVersionFollowChartRelease:   "leonardo-dev",
					ChartVersionResolver:           utils.PointerTo("exact"),
					ChartVersionExact:              utils.PointerTo("2.0.0"),
					ChartVersionFollowChartRelease: "leonardo-staging",
					HelmfileRef:                    utils.PointerTo("HEAD"),
					HelmfileRefEnabled:             utils.PointerTo(true),
					ChartReleaseV3Edit: ChartReleaseV3Edit{
						Subdomain:               utils.PointerTo("subdomain"),
						Protocol:                utils.PointerTo("https"),
						Port:                    utils.PointerTo[uint](443),
						PagerdutyIntegration:    utils.PointerTo("pd-id/blah"),
						IncludeInBulkChangesets: utils.PointerTo(true),
					},
				},
			},
		},
		{
			name: "pagerduty ID case",
			args: args{model: models.ChartRelease{
				PagerdutyIntegrationID: utils.PointerTo[uint](6),
			}},
			want: ChartReleaseV3{
				ChartReleaseV3Create: ChartReleaseV3Create{
					ChartReleaseV3Edit: ChartReleaseV3Edit{
						PagerdutyIntegration: utils.PointerTo("6"),
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, chartReleaseFromModel(tt.args.model), "chartReleaseFromModel(%v)", tt.args.model)
		})
	}
}
