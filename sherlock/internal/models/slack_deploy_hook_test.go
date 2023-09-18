package models

import (
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (s *modelSuite) TestSlackDeployHookEnvironment() {
	s.SetSuitableTestUserForDB()
	environment := Environment{Name: "dev", Base: "live", Lifecycle: "static", RequiresSuitability: utils.PointerTo(false), HelmfileRef: utils.PointerTo("HEAD"), PreventDeletion: utils.PointerTo(false)}
	s.NoError(s.DB.Create(&environment).Error)

	hook := SlackDeployHook{SlackChannel: utils.PointerTo("channel"),
		Trigger: DeployHookTriggerConfig{OnEnvironmentID: &environment.ID}}
	s.NoError(s.DB.Create(&hook).Error)
}

func (s *modelSuite) TestSlackDeployHookChartRelease() {
	s.SetSuitableTestUserForDB()
	cluster := Cluster{Name: "terra-dev", Address: utils.PointerTo("0.0.0.0"), Base: utils.PointerTo("terra"), RequiresSuitability: utils.PointerTo(false), HelmfileRef: utils.PointerTo("HEAD")}
	s.NoError(s.DB.Create(&cluster).Error)
	environment := Environment{Name: "dev", Base: "live", Lifecycle: "static", RequiresSuitability: utils.PointerTo(false), HelmfileRef: utils.PointerTo("HEAD"), PreventDeletion: utils.PointerTo(false)}
	s.NoError(s.DB.Create(&environment).Error)
	chart := Chart{Name: "leonardo", ChartRepo: utils.PointerTo("terra-helm")}
	s.NoError(s.DB.Create(&chart).Error)
	chartRelease := ChartRelease{Name: "leonardo-dev", ChartID: chart.ID, EnvironmentID: &environment.ID, ClusterID: &cluster.ID, Namespace: "terra-dev",
		ChartReleaseVersion: ChartReleaseVersion{AppVersionResolver: utils.PointerTo("exact"), AppVersionExact: utils.PointerTo("v1.2.3"),
			ChartVersionResolver: utils.PointerTo("exact"), ChartVersionExact: utils.PointerTo("v2.3.4"), HelmfileRef: utils.PointerTo("HEAD"), HelmfileRefEnabled: utils.PointerTo(false)}}
	s.NoError(s.DB.Create(&chartRelease).Error)

	hook := SlackDeployHook{SlackChannel: utils.PointerTo("channel"),
		Trigger: DeployHookTriggerConfig{OnChartReleaseID: &chartRelease.ID}}
	s.NoError(s.DB.Create(&hook).Error)
}

func (s *modelSuite) TestSlackDeployHookNoChannel() {
	s.SetSuitableTestUserForDB()
	environment := Environment{Name: "dev", Base: "live", Lifecycle: "static", RequiresSuitability: utils.PointerTo(false), HelmfileRef: utils.PointerTo("HEAD"), PreventDeletion: utils.PointerTo(false)}
	s.NoError(s.DB.Create(&environment).Error)

	hook := SlackDeployHook{Trigger: DeployHookTriggerConfig{OnEnvironmentID: &environment.ID}}
	s.ErrorContains(s.DB.Create(&hook).Error, "slack_channel_present")
}

func (s *modelSuite) TestSlackDeployHookSuitability() {
	s.SetSuitableTestUserForDB()
	environment := Environment{Name: "prod", Base: "live", Lifecycle: "static", RequiresSuitability: utils.PointerTo(true), HelmfileRef: utils.PointerTo("HEAD"), PreventDeletion: utils.PointerTo(false)}
	s.NoError(s.DB.Create(&environment).Error)

	s.Run("when suitable", func() {
		s.SetSuitableTestUserForDB()
		hook := SlackDeployHook{SlackChannel: utils.PointerTo("channel"), Trigger: DeployHookTriggerConfig{OnEnvironmentID: &environment.ID}}
		s.NoError(s.DB.Create(&hook).Error)
	})
	s.Run("when not suitable", func() {
		s.SetNonSuitableTestUserForDB()
		hook := SlackDeployHook{SlackChannel: utils.PointerTo("channel"), Trigger: DeployHookTriggerConfig{OnEnvironmentID: &environment.ID}}
		s.ErrorContains(s.DB.Create(&hook).Error, errors.Forbidden)
	})
}

func (s *modelSuite) TestSlackDeployHookFlow() {
	s.SetSuitableTestUserForDB()
	environment := Environment{Name: "dev", Base: "live", Lifecycle: "static", RequiresSuitability: utils.PointerTo(false), HelmfileRef: utils.PointerTo("HEAD"), PreventDeletion: utils.PointerTo(false)}
	s.NoError(s.DB.Create(&environment).Error)

	hook := SlackDeployHook{SlackChannel: utils.PointerTo("channel"), Trigger: DeployHookTriggerConfig{OnEnvironmentID: &environment.ID}}
	s.NoError(s.DB.Create(&hook).Error)
	s.NotZero(hook.ID)
	s.NotZero(hook.Trigger.ID)

	s.Run("edits hook", func() {
		s.NoError(s.DB.Model(&hook).Updates(&SlackDeployHook{SlackChannel: utils.PointerTo("different channel")}).Error)
		if s.NotNil(hook.SlackChannel) {
			s.Equal("different channel", *hook.SlackChannel)
		}

		s.Run("read", func() {
			hookAgain := &SlackDeployHook{Model: gorm.Model{ID: hook.ID}}
			s.NoError(s.DB.Take(&hookAgain).Error)
			if s.NotNil(hookAgain.SlackChannel) {
				s.Equal("different channel", *hookAgain.SlackChannel)
			}
		})

	})

	s.Run("edits trigger", func() {
		s.NoError(s.DB.Model(&hook.Trigger).Updates(&DeployHookTriggerConfig{OnSuccess: utils.PointerTo(true)}).Error)
		if s.NotNil(hook.Trigger.OnSuccess) {
			s.True(*hook.Trigger.OnSuccess)
		}

		s.Run("read", func() {
			s.Run("empty without association", func() {
				hookAgain := &SlackDeployHook{Model: gorm.Model{ID: hook.ID}}
				s.NoError(s.DB.Take(&hookAgain).Error)
				s.Nil(hookAgain.Trigger.OnSuccess)
			})
			s.Run("filled with association", func() {
				hookAgain := &SlackDeployHook{Model: gorm.Model{ID: hook.ID}}
				s.NoError(s.DB.Preload("Trigger").Take(&hookAgain).Error)
				if s.NotNil(hookAgain.Trigger.OnSuccess) {
					s.True(*hookAgain.Trigger.OnSuccess)
				}
			})
			s.Run("filled with preload all", func() {
				hookAgain := &SlackDeployHook{Model: gorm.Model{ID: hook.ID}}
				s.NoError(s.DB.Preload(clause.Associations).Take(&hookAgain).Error)
				if s.NotNil(hookAgain.Trigger.OnSuccess) {
					s.True(*hookAgain.Trigger.OnSuccess)
				}
			})
		})
	})

	s.Run("delete", func() {
		s.NoError(s.DB.Select("Trigger").Delete(&SlackDeployHook{Model: gorm.Model{ID: hook.ID}}).Error)

		s.Run("hook gone", func() {
			var matchingHooks []SlackDeployHook
			s.NoError(s.DB.Where(&SlackDeployHook{Model: gorm.Model{ID: hook.ID}}).Find(&matchingHooks).Error)
			s.Len(matchingHooks, 0)
		})
		s.Run("trigger gone", func() {
			var matchingTriggers []DeployHookTriggerConfig
			s.NoError(s.DB.Where(&DeployHookTriggerConfig{Model: gorm.Model{ID: hook.Trigger.ID}}).Find(&matchingTriggers).Error)
			s.Len(matchingTriggers, 0)
		})
	})
}
