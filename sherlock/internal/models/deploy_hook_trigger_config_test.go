package models

import (
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
)

func (s *modelSuite) TestDeployHookTriggerConfigHookIDMissing() {
	environment := Environment{Name: "dev", Base: "live", Lifecycle: "static", RequiresSuitability: utils.PointerTo(false), HelmfileRef: utils.PointerTo("HEAD"), PreventDeletion: utils.PointerTo(false)}
	s.NoError(s.DB.Create(&environment).Error)

	s.ErrorContains(s.DB.Create(&DeployHookTriggerConfig{OnEnvironmentID: &environment.ID, HookType: "type"}).Error, "hook_id_present")
}

func (s *modelSuite) TestDeployHookTriggerConfigHookTypeMissing() {
	environment := Environment{Name: "dev", Base: "live", Lifecycle: "static", RequiresSuitability: utils.PointerTo(false), HelmfileRef: utils.PointerTo("HEAD"), PreventDeletion: utils.PointerTo(false)}
	s.NoError(s.DB.Create(&environment).Error)

	s.ErrorContains(s.DB.Create(&DeployHookTriggerConfig{OnEnvironmentID: &environment.ID, HookID: 1}).Error, "hook_type_present")
}

func (s *modelSuite) TestDeployHookTriggerConfigEnvironmentAndChartRelease() {
	s.SetNonSuitableTestUserForDB()
	cluster := Cluster{
		Name:                "terra-dev",
		Address:             utils.PointerTo("0.0.0.0"),
		Base:                utils.PointerTo("terra"),
		Location:            "some-location",
		Provider:            "google",
		GoogleProject:       "some-google-project",
		RequiresSuitability: utils.PointerTo(false),
		HelmfileRef:         utils.PointerTo("HEAD"),
	}
	s.NoError(s.DB.Create(&cluster).Error)
	environment := Environment{Name: "dev", Base: "live", Lifecycle: "static", RequiresSuitability: utils.PointerTo(false), HelmfileRef: utils.PointerTo("HEAD"), PreventDeletion: utils.PointerTo(false)}
	s.NoError(s.DB.Create(&environment).Error)
	chart := Chart{Name: "leonardo", ChartRepo: utils.PointerTo("terra-helm")}
	s.NoError(s.DB.Create(&chart).Error)
	chartRelease := ChartRelease{Name: "leonardo-dev", ChartID: chart.ID, EnvironmentID: &environment.ID, ClusterID: &cluster.ID, Namespace: "terra-dev",
		ChartReleaseVersion: ChartReleaseVersion{AppVersionResolver: utils.PointerTo("exact"), AppVersionExact: utils.PointerTo("v1.2.3"),
			ChartVersionResolver: utils.PointerTo("exact"), ChartVersionExact: utils.PointerTo("v2.3.4"), HelmfileRef: utils.PointerTo("HEAD"), HelmfileRefEnabled: utils.PointerTo(false)}}
	s.NoError(s.DB.Create(&chartRelease).Error)

	config := DeployHookTriggerConfig{OnEnvironmentID: &environment.ID, OnChartReleaseID: &chartRelease.ID}
	s.ErrorContains(s.DB.Create(&config).Error, "environment_or_chart_release_present")
}

func (s *modelSuite) TestDeployHookTriggerConfigEnvironmentSuitable() {
	environment := Environment{Name: "dev", Base: "live", Lifecycle: "static", RequiresSuitability: utils.PointerTo(true), HelmfileRef: utils.PointerTo("HEAD"), PreventDeletion: utils.PointerTo(false)}
	s.NoError(s.DB.Create(&environment).Error)

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
	s.SetNonSuitableTestUserForDB()
	cluster := Cluster{
		Name:                "terra-dev",
		Address:             utils.PointerTo("0.0.0.0"),
		Base:                utils.PointerTo("terra"),
		Location:            "some-location",
		Provider:            "google",
		GoogleProject:       "some-google-project",
		RequiresSuitability: utils.PointerTo(false),
		HelmfileRef:         utils.PointerTo("HEAD"),
	}
	s.NoError(s.DB.Create(&cluster).Error)
	environment := Environment{Name: "dev", Base: "live", Lifecycle: "static", RequiresSuitability: utils.PointerTo(true), HelmfileRef: utils.PointerTo("HEAD"), PreventDeletion: utils.PointerTo(false)}
	s.NoError(s.DB.Create(&environment).Error)
	chart := Chart{Name: "leonardo", ChartRepo: utils.PointerTo("terra-helm")}
	s.NoError(s.DB.Create(&chart).Error)
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
	s.SetSuitableTestUserForDB()
	environment := Environment{Name: "dev", Base: "live", Lifecycle: "static", RequiresSuitability: utils.PointerTo(false), HelmfileRef: utils.PointerTo("HEAD"), PreventDeletion: utils.PointerTo(false)}
	s.NoError(s.DB.Create(&environment).Error)
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
