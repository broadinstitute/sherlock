package models

import (
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
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
