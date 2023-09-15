package models

import (
	"github.com/broadinstitute/sherlock/go-shared/pkg/testutils"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
)

func (s *modelSuite) TestDeployHookTriggerConfigHookIDMissing() {
	environment := Environment{Name: "dev", Base: "live", Lifecycle: "static", RequiresSuitability: testutils.PointerTo(false), HelmfileRef: testutils.PointerTo("HEAD"), PreventDeletion: testutils.PointerTo(false)}
	s.NoError(s.DB.Create(&environment).Error)

	s.ErrorContains(s.DB.Create(&DeployHookTriggerConfig{OnEnvironmentID: &environment.ID, HookType: "type"}).Error, "hook_id_present")
}

func (s *modelSuite) TestDeployHookTriggerConfigHookTypeMissing() {
	environment := Environment{Name: "dev", Base: "live", Lifecycle: "static", RequiresSuitability: testutils.PointerTo(false), HelmfileRef: testutils.PointerTo("HEAD"), PreventDeletion: testutils.PointerTo(false)}
	s.NoError(s.DB.Create(&environment).Error)

	s.ErrorContains(s.DB.Create(&DeployHookTriggerConfig{OnEnvironmentID: &environment.ID, HookID: 1}).Error, "hook_type_present")
}

func (s *modelSuite) TestDeployHookTriggerConfigEnvironmentAndChartRelease() {
	cluster := Cluster{Name: "terra-dev", Address: testutils.PointerTo("0.0.0.0"), Base: testutils.PointerTo("terra"), RequiresSuitability: testutils.PointerTo(false), HelmfileRef: testutils.PointerTo("HEAD")}
	s.NoError(s.DB.Create(&cluster).Error)
	environment := Environment{Name: "dev", Base: "live", Lifecycle: "static", RequiresSuitability: testutils.PointerTo(false), HelmfileRef: testutils.PointerTo("HEAD"), PreventDeletion: testutils.PointerTo(false)}
	s.NoError(s.DB.Create(&environment).Error)
	chart := Chart{Name: "leonardo", ChartRepo: testutils.PointerTo("terra-helm")}
	s.NoError(s.DB.Create(&chart).Error)
	chartRelease := ChartRelease{Name: "leonardo-dev", ChartID: chart.ID, EnvironmentID: &environment.ID, ClusterID: &cluster.ID, Namespace: "terra-dev",
		ChartReleaseVersion: ChartReleaseVersion{AppVersionResolver: testutils.PointerTo("exact"), AppVersionExact: testutils.PointerTo("v1.2.3"),
			ChartVersionResolver: testutils.PointerTo("exact"), ChartVersionExact: testutils.PointerTo("v2.3.4"), HelmfileRef: testutils.PointerTo("HEAD"), HelmfileRefEnabled: testutils.PointerTo(false)}}
	s.NoError(s.DB.Create(&chartRelease).Error)

	config := DeployHookTriggerConfig{OnEnvironmentID: &environment.ID, OnChartReleaseID: &chartRelease.ID}
	s.ErrorContains(s.DB.Create(&config).Error, "environment_or_chart_release_present")
}

func (s *modelSuite) TestDeployHookTriggerConfigEnvironmentSuitable() {
	environment := Environment{Name: "dev", Base: "live", Lifecycle: "static", RequiresSuitability: testutils.PointerTo(true), HelmfileRef: testutils.PointerTo("HEAD"), PreventDeletion: testutils.PointerTo(false)}
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
	cluster := Cluster{Name: "terra-dev", Address: testutils.PointerTo("0.0.0.0"), Base: testutils.PointerTo("terra"), RequiresSuitability: testutils.PointerTo(false), HelmfileRef: testutils.PointerTo("HEAD")}
	s.NoError(s.DB.Create(&cluster).Error)
	environment := Environment{Name: "dev", Base: "live", Lifecycle: "static", RequiresSuitability: testutils.PointerTo(true), HelmfileRef: testutils.PointerTo("HEAD"), PreventDeletion: testutils.PointerTo(false)}
	s.NoError(s.DB.Create(&environment).Error)
	chart := Chart{Name: "leonardo", ChartRepo: testutils.PointerTo("terra-helm")}
	s.NoError(s.DB.Create(&chart).Error)
	chartRelease := ChartRelease{Name: "leonardo-dev", ChartID: chart.ID, EnvironmentID: &environment.ID, ClusterID: &cluster.ID, Namespace: "terra-dev",
		ChartReleaseVersion: ChartReleaseVersion{AppVersionResolver: testutils.PointerTo("exact"), AppVersionExact: testutils.PointerTo("v1.2.3"),
			ChartVersionResolver: testutils.PointerTo("exact"), ChartVersionExact: testutils.PointerTo("v2.3.4"), HelmfileRef: testutils.PointerTo("HEAD"), HelmfileRefEnabled: testutils.PointerTo(false)}}
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
	cluster := Cluster{Name: "terra-dev", Address: testutils.PointerTo("0.0.0.0"), Base: testutils.PointerTo("terra"), RequiresSuitability: testutils.PointerTo(true), HelmfileRef: testutils.PointerTo("HEAD")}
	s.NoError(s.DB.Create(&cluster).Error)
	environment := Environment{Name: "dev", Base: "live", Lifecycle: "static", RequiresSuitability: testutils.PointerTo(false), HelmfileRef: testutils.PointerTo("HEAD"), PreventDeletion: testutils.PointerTo(false)}
	s.NoError(s.DB.Create(&environment).Error)
	chart := Chart{Name: "leonardo", ChartRepo: testutils.PointerTo("terra-helm")}
	s.NoError(s.DB.Create(&chart).Error)
	chartRelease := ChartRelease{Name: "leonardo-dev", ChartID: chart.ID, EnvironmentID: &environment.ID, ClusterID: &cluster.ID, Namespace: "terra-dev",
		ChartReleaseVersion: ChartReleaseVersion{AppVersionResolver: testutils.PointerTo("exact"), AppVersionExact: testutils.PointerTo("v1.2.3"),
			ChartVersionResolver: testutils.PointerTo("exact"), ChartVersionExact: testutils.PointerTo("v2.3.4"), HelmfileRef: testutils.PointerTo("HEAD"), HelmfileRefEnabled: testutils.PointerTo(false)}}
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
