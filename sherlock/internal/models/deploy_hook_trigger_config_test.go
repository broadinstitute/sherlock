package models

import (
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
)

func (s *modelSuite) TestDeployHookTriggerConfigHookIDMissing() {
	environment := s.TestData.Environment_Dev()
	s.ErrorContains(s.DB.Create(&DeployHookTriggerConfig{OnEnvironmentID: &environment.ID, HookType: "type"}).Error, "hook_id_present")
}

func (s *modelSuite) TestDeployHookTriggerConfigHookTypeMissing() {
	environment := s.TestData.Environment_Dev()
	s.ErrorContains(s.DB.Create(&DeployHookTriggerConfig{OnEnvironmentID: &environment.ID, HookID: 1}).Error, "hook_type_present")
}

func (s *modelSuite) TestDeployHookTriggerConfigEnvironmentAndChartRelease() {
	environment := s.TestData.Environment_Dev()
	chartRelease := s.TestData.ChartRelease_LeonardoDev()
	s.SetNonSuitableTestUserForDB()
	config := DeployHookTriggerConfig{OnEnvironmentID: &environment.ID, OnChartReleaseID: &chartRelease.ID}
	s.ErrorContains(s.DB.Create(&config).Error, "environment_or_chart_release_present")
}

func (s *modelSuite) TestDeployHookTriggerConfigEnvironmentSuitable() {
	environment := s.TestData.Environment_Prod()

	s.Run("when suitable", func() {
		s.SetSuitableTestUserForDB()
		s.NoError(s.DB.Create(&DeployHookTriggerConfig{OnEnvironmentID: &environment.ID, HookID: 1, HookType: "type"}).Error)
	})
	s.Run("not suitable", func() {
		s.SetNonSuitableTestUserForDB()
		s.ErrorContains(s.DB.Create(&DeployHookTriggerConfig{OnEnvironmentID: &environment.ID, HookID: 1, HookType: "type"}).Error, errors.Forbidden)
	})
}

func (s *modelSuite) TestDeployHookTriggerConfigChartReleaseSuitableViaEnvironment() {
	cluster := s.TestData.Cluster_TerraDev()
	chart := s.TestData.Chart_Leonardo()
	environment := s.TestData.Environment_Prod()
	s.SetSuitableTestUserForDB()
	chartRelease := ChartRelease{Name: "leonardo-dev", ChartID: chart.ID, EnvironmentID: &environment.ID, ClusterID: &cluster.ID, Namespace: "terra-dev",
		ChartReleaseVersion: ChartReleaseVersion{AppVersionResolver: utils.PointerTo("exact"), AppVersionExact: utils.PointerTo("v1.2.3"),
			ChartVersionResolver: utils.PointerTo("exact"), ChartVersionExact: utils.PointerTo("v2.3.4"), HelmfileRef: utils.PointerTo("HEAD"), HelmfileRefEnabled: utils.PointerTo(false)}}
	s.NoError(s.DB.Create(&chartRelease).Error)

	s.Run("when suitable", func() {
		s.SetSuitableTestUserForDB()
		s.NoError(s.DB.Create(&DeployHookTriggerConfig{OnChartReleaseID: &chartRelease.ID, HookID: 1, HookType: "type"}).Error)
	})
	s.Run("not suitable", func() {
		s.SetNonSuitableTestUserForDB()
		s.ErrorContains(s.DB.Create(&DeployHookTriggerConfig{OnChartReleaseID: &chartRelease.ID, HookID: 1, HookType: "type"}).Error, errors.Forbidden)
	})
}

func (s *modelSuite) TestDeployHookTriggerConfigChartReleaseSuitableViaCluster() {
	cluster := s.TestData.Cluster_TerraProd()
	chart := s.TestData.Chart_Leonardo()
	environment := s.TestData.Environment_Dev()
	s.SetSuitableTestUserForDB()
	chartRelease := ChartRelease{
		Name:          "leonardo-dev",
		ChartID:       chart.ID,
		EnvironmentID: &environment.ID,
		ClusterID:     &cluster.ID,
		Namespace:     "terra-dev",
		ChartReleaseVersion: ChartReleaseVersion{
			AppVersionResolver:   utils.PointerTo("exact"),
			AppVersionExact:      utils.PointerTo("v1.2.3"),
			ChartVersionResolver: utils.PointerTo("exact"),
			ChartVersionExact:    utils.PointerTo("v2.3.4"),
			HelmfileRef:          utils.PointerTo("HEAD"),
			HelmfileRefEnabled:   utils.PointerTo(false),
		},
	}
	s.NoError(s.DB.Model(&ChartRelease{}).Create(&chartRelease).Error)
	s.Run("when suitable", func() {
		s.SetSuitableTestUserForDB()
		s.NoError(s.DB.Create(&DeployHookTriggerConfig{OnChartReleaseID: &chartRelease.ID, HookID: 1, HookType: "type"}).Error)
	})
	s.Run("not suitable", func() {
		s.SetNonSuitableTestUserForDB()
		s.ErrorContains(s.DB.Create(&DeployHookTriggerConfig{OnChartReleaseID: &chartRelease.ID, HookID: 1, HookType: "type"}).Error, errors.Forbidden)
	})
}

func TestDeployHookTriggerConfig_SlackBeehiveLink(t *testing.T) {
	type fields struct {
		Model          gorm.Model
		OnEnvironment  *Environment
		OnChartRelease *ChartRelease
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "normal environment",
			fields: fields{
				OnEnvironment: &Environment{
					Name: "dev",
				},
			},
			want: "<" + fmt.Sprintf(config.Config.String("beehive.environmentUrlFormatString"), "dev") + "|dev>",
		},
		{
			name: "empty environment",
			fields: fields{
				OnEnvironment: &Environment{},
			},
			want: "(unknown environment)",
		},
		{
			name: "normal chart release",
			fields: fields{
				OnChartRelease: &ChartRelease{
					Name: "leonardo-dev",
				},
			},
			want: "<" + fmt.Sprintf(config.Config.String("beehive.chartReleaseUrlFormatString"), "leonardo-dev") + "|leonardo-dev>",
		},
		{
			name: "empty chart release",
			fields: fields{
				OnChartRelease: &ChartRelease{},
			},
			want: "(unknown chart release)",
		},
		{
			name: "empty",
			fields: fields{
				Model: gorm.Model{ID: 123},
			},
			want: "(orphaned deploy hook trigger config 123)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DeployHookTriggerConfig{
				Model:          tt.fields.Model,
				OnEnvironment:  tt.fields.OnEnvironment,
				OnChartRelease: tt.fields.OnChartRelease,
			}
			assert.Equalf(t, tt.want, d.SlackBeehiveLink(), "SlackBeehiveLink()")
		})
	}
}

func TestDeployHookTriggerConfig_ArgoCdUrl(t *testing.T) {
	type fields struct {
		OnEnvironment  *Environment
		OnChartRelease *ChartRelease
	}
	tests := []struct {
		name   string
		fields fields
		want   string
		wantOk bool
	}{
		{
			name: "normal environment",
			fields: fields{
				OnEnvironment: &Environment{
					Name: "dev",
				},
			},
			want:   fmt.Sprintf(config.Config.String("argoCd.environmentUrlFormatString"), "dev"),
			wantOk: true,
		},
		{
			name: "empty environment",
			fields: fields{
				OnEnvironment: &Environment{},
			},
			want:   "",
			wantOk: false,
		},
		{
			name: "normal chart release",
			fields: fields{
				OnChartRelease: &ChartRelease{
					Name: "leonardo-dev",
				},
			},
			want:   fmt.Sprintf(config.Config.String("argoCd.chartReleaseUrlFormatString"), "leonardo-dev"),
			wantOk: true,
		},
		{
			name: "empty chart release",
			fields: fields{
				OnChartRelease: &ChartRelease{},
			},
			want:   "",
			wantOk: false,
		},
		{
			name:   "empty",
			fields: fields{},
			want:   "",
			wantOk: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DeployHookTriggerConfig{
				OnEnvironment:  tt.fields.OnEnvironment,
				OnChartRelease: tt.fields.OnChartRelease,
			}
			got, gotOk := d.ArgoCdUrl()
			assert.Equalf(t, tt.want, got, "SlackBeehiveLink()")
			assert.Equalf(t, tt.wantOk, gotOk, "ArgoCdUrl()")
		})
	}
}
