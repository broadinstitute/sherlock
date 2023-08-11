package models

import (
	"github.com/broadinstitute/sherlock/go-shared/pkg/testutils"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (s *modelSuite) TestGithubActionsDeployHookEnvironment() {
	s.SetSuitableTestUserForDB()
	environment := Environment{Name: "dev", Base: "live", Lifecycle: "static", RequiresSuitability: testutils.PointerTo(false), HelmfileRef: testutils.PointerTo("HEAD"), PreventDeletion: testutils.PointerTo(false)}
	s.NoError(s.DB.Create(&environment).Error)

	hook := GithubActionsDeployHook{
		GithubActionsOwner:        testutils.PointerTo("owner"),
		GithubActionsRepo:         testutils.PointerTo("repo"),
		GithubActionsWorkflowPath: testutils.PointerTo("path"),
		GithubActionsDefaultRef:   testutils.PointerTo("head"),
		GithubActionsRefBehavior:  testutils.PointerTo("always-use-default-ref"),
		Trigger:                   DeployHookTriggerConfig{OnEnvironmentID: &environment.ID}}
	s.NoError(s.DB.Create(&hook).Error)
}

func (s *modelSuite) TestGithubActionsDeployHookChartRelease() {
	s.SetSuitableTestUserForDB()
	cluster := Cluster{Name: "terra-dev", Address: testutils.PointerTo("0.0.0.0"), Base: testutils.PointerTo("terra"), RequiresSuitability: testutils.PointerTo(false), HelmfileRef: testutils.PointerTo("HEAD")}
	s.NoError(s.DB.Create(&cluster).Error)
	environment := Environment{Name: "dev", Base: "live", Lifecycle: "static", RequiresSuitability: testutils.PointerTo(false), HelmfileRef: testutils.PointerTo("HEAD"), PreventDeletion: testutils.PointerTo(false)}
	s.NoError(s.DB.Create(&environment).Error)
	chart := Chart{Name: "leonardo", ChartRepo: testutils.PointerTo("terra-helm")}
	s.NoError(s.DB.Create(&chart).Error)
	chartRelease := ChartRelease{Name: "leonardo-dev", ChartID: chart.ID, EnvironmentID: &environment.ID, ClusterID: &cluster.ID, Namespace: "terra-dev",
		ChartReleaseVersion: ChartReleaseVersion{AppVersionResolver: testutils.PointerTo("exact"), AppVersionExact: testutils.PointerTo("v1.2.3"),
			ChartVersionResolver: testutils.PointerTo("exact"), ChartVersionExact: testutils.PointerTo("v2.3.4"), HelmfileRef: testutils.PointerTo("HEAD")}}
	s.NoError(s.DB.Create(&chartRelease).Error)

	hook := GithubActionsDeployHook{
		GithubActionsOwner:        testutils.PointerTo("owner"),
		GithubActionsRepo:         testutils.PointerTo("repo"),
		GithubActionsWorkflowPath: testutils.PointerTo("path"),
		GithubActionsDefaultRef:   testutils.PointerTo("head"),
		GithubActionsRefBehavior:  testutils.PointerTo("always-use-default-ref"),
		Trigger:                   DeployHookTriggerConfig{OnChartReleaseID: &chartRelease.ID}}
	s.NoError(s.DB.Create(&hook).Error)
}

func (s *modelSuite) TestGithubActionsDeployHookNoOwner() {
	s.SetSuitableTestUserForDB()
	environment := Environment{Name: "dev", Base: "live", Lifecycle: "static", RequiresSuitability: testutils.PointerTo(false), HelmfileRef: testutils.PointerTo("HEAD"), PreventDeletion: testutils.PointerTo(false)}
	s.NoError(s.DB.Create(&environment).Error)

	hook := GithubActionsDeployHook{
		GithubActionsRepo:         testutils.PointerTo("repo"),
		GithubActionsWorkflowPath: testutils.PointerTo("path"),
		GithubActionsDefaultRef:   testutils.PointerTo("head"),
		GithubActionsRefBehavior:  testutils.PointerTo("always-use-default-ref"),
		Trigger:                   DeployHookTriggerConfig{OnEnvironmentID: &environment.ID}}
	s.ErrorContains(s.DB.Create(&hook).Error, "github_actions_owner_present")
}

func (s *modelSuite) TestGithubActionsDeployHookNoRepo() {
	s.SetSuitableTestUserForDB()
	environment := Environment{Name: "dev", Base: "live", Lifecycle: "static", RequiresSuitability: testutils.PointerTo(false), HelmfileRef: testutils.PointerTo("HEAD"), PreventDeletion: testutils.PointerTo(false)}
	s.NoError(s.DB.Create(&environment).Error)

	hook := GithubActionsDeployHook{
		GithubActionsOwner:        testutils.PointerTo("owner"),
		GithubActionsWorkflowPath: testutils.PointerTo("path"),
		GithubActionsDefaultRef:   testutils.PointerTo("head"),
		GithubActionsRefBehavior:  testutils.PointerTo("always-use-default-ref"),
		Trigger:                   DeployHookTriggerConfig{OnEnvironmentID: &environment.ID}}
	s.ErrorContains(s.DB.Create(&hook).Error, "github_actions_repo_present")
}

func (s *modelSuite) TestGithubActionsDeployHookNoWorkflowPath() {
	s.SetSuitableTestUserForDB()
	environment := Environment{Name: "dev", Base: "live", Lifecycle: "static", RequiresSuitability: testutils.PointerTo(false), HelmfileRef: testutils.PointerTo("HEAD"), PreventDeletion: testutils.PointerTo(false)}
	s.NoError(s.DB.Create(&environment).Error)

	hook := GithubActionsDeployHook{
		GithubActionsOwner:       testutils.PointerTo("owner"),
		GithubActionsRepo:        testutils.PointerTo("repo"),
		GithubActionsDefaultRef:  testutils.PointerTo("head"),
		GithubActionsRefBehavior: testutils.PointerTo("always-use-default-ref"),
		Trigger:                  DeployHookTriggerConfig{OnEnvironmentID: &environment.ID}}
	s.ErrorContains(s.DB.Create(&hook).Error, "github_actions_workflow_path_present")
}

func (s *modelSuite) TestGithubActionsDeployHookNoDefaultRef() {
	s.SetSuitableTestUserForDB()
	environment := Environment{Name: "dev", Base: "live", Lifecycle: "static", RequiresSuitability: testutils.PointerTo(false), HelmfileRef: testutils.PointerTo("HEAD"), PreventDeletion: testutils.PointerTo(false)}
	s.NoError(s.DB.Create(&environment).Error)

	hook := GithubActionsDeployHook{
		GithubActionsOwner:        testutils.PointerTo("owner"),
		GithubActionsRepo:         testutils.PointerTo("repo"),
		GithubActionsWorkflowPath: testutils.PointerTo("path"),
		GithubActionsRefBehavior:  testutils.PointerTo("always-use-default-ref"),
		Trigger:                   DeployHookTriggerConfig{OnEnvironmentID: &environment.ID}}
	s.ErrorContains(s.DB.Create(&hook).Error, "github_actions_default_ref_present")
}

func (s *modelSuite) TestGithubActionsDeployHookNoRefBehavior() {
	s.SetSuitableTestUserForDB()
	environment := Environment{Name: "dev", Base: "live", Lifecycle: "static", RequiresSuitability: testutils.PointerTo(false), HelmfileRef: testutils.PointerTo("HEAD"), PreventDeletion: testutils.PointerTo(false)}
	s.NoError(s.DB.Create(&environment).Error)

	hook := GithubActionsDeployHook{
		GithubActionsOwner:        testutils.PointerTo("owner"),
		GithubActionsRepo:         testutils.PointerTo("repo"),
		GithubActionsWorkflowPath: testutils.PointerTo("path"),
		GithubActionsDefaultRef:   testutils.PointerTo("head"),
		Trigger:                   DeployHookTriggerConfig{OnEnvironmentID: &environment.ID}}
	s.ErrorContains(s.DB.Create(&hook).Error, "github_actions_ref_behavior_valid")
}

func (s *modelSuite) TestGithubActionsDeployHookBadRefBehavior() {
	s.SetSuitableTestUserForDB()
	environment := Environment{Name: "dev", Base: "live", Lifecycle: "static", RequiresSuitability: testutils.PointerTo(false), HelmfileRef: testutils.PointerTo("HEAD"), PreventDeletion: testutils.PointerTo(false)}
	s.NoError(s.DB.Create(&environment).Error)

	hook := GithubActionsDeployHook{
		GithubActionsOwner:        testutils.PointerTo("owner"),
		GithubActionsRepo:         testutils.PointerTo("repo"),
		GithubActionsWorkflowPath: testutils.PointerTo("path"),
		GithubActionsDefaultRef:   testutils.PointerTo("head"),
		GithubActionsRefBehavior:  testutils.PointerTo("some nonsense string"),
		Trigger:                   DeployHookTriggerConfig{OnEnvironmentID: &environment.ID}}
	s.ErrorContains(s.DB.Create(&hook).Error, "github_actions_ref_behavior_valid")
}

func (s *modelSuite) TestGithubActionsDeployHookSuitability() {
	s.SetSuitableTestUserForDB()
	environment := Environment{Name: "prod", Base: "live", Lifecycle: "static", RequiresSuitability: testutils.PointerTo(true), HelmfileRef: testutils.PointerTo("HEAD"), PreventDeletion: testutils.PointerTo(false)}
	s.NoError(s.DB.Create(&environment).Error)

	s.Run("when suitable", func() {
		s.SetSuitableTestUserForDB()
		hook := GithubActionsDeployHook{
			GithubActionsOwner:        testutils.PointerTo("owner"),
			GithubActionsRepo:         testutils.PointerTo("repo"),
			GithubActionsWorkflowPath: testutils.PointerTo("path"),
			GithubActionsDefaultRef:   testutils.PointerTo("head"),
			GithubActionsRefBehavior:  testutils.PointerTo("always-use-default-ref"),
			Trigger:                   DeployHookTriggerConfig{OnEnvironmentID: &environment.ID}}
		s.NoError(s.DB.Create(&hook).Error)
	})
	s.Run("when not suitable", func() {
		s.SetNonSuitableTestUserForDB()
		hook := GithubActionsDeployHook{
			GithubActionsOwner:        testutils.PointerTo("owner"),
			GithubActionsRepo:         testutils.PointerTo("repo"),
			GithubActionsWorkflowPath: testutils.PointerTo("path"),
			GithubActionsDefaultRef:   testutils.PointerTo("head"),
			GithubActionsRefBehavior:  testutils.PointerTo("always-use-default-ref"),
			Trigger:                   DeployHookTriggerConfig{OnEnvironmentID: &environment.ID}}
		s.ErrorContains(s.DB.Create(&hook).Error, errors.Forbidden)
	})
}

func (s *modelSuite) TestGithubActionsDeployHookFlow() {
	s.SetSuitableTestUserForDB()
	environment := Environment{Name: "dev", Base: "live", Lifecycle: "static", RequiresSuitability: testutils.PointerTo(false), HelmfileRef: testutils.PointerTo("HEAD"), PreventDeletion: testutils.PointerTo(false)}
	s.NoError(s.DB.Create(&environment).Error)

	hook := GithubActionsDeployHook{
		GithubActionsOwner:        testutils.PointerTo("owner"),
		GithubActionsRepo:         testutils.PointerTo("repo"),
		GithubActionsWorkflowPath: testutils.PointerTo("path"),
		GithubActionsDefaultRef:   testutils.PointerTo("head"),
		GithubActionsRefBehavior:  testutils.PointerTo("always-use-default-ref"),
		Trigger:                   DeployHookTriggerConfig{OnEnvironmentID: &environment.ID}}
	s.NoError(s.DB.Create(&hook).Error)
	s.NotZero(hook.ID)
	s.NotZero(hook.Trigger.ID)

	s.Run("edits hook", func() {
		s.NoError(s.DB.Model(&hook).Updates(&GithubActionsDeployHook{GithubActionsRefBehavior: testutils.PointerTo("use-app-version-as-ref")}).Error)
		if s.NotNil(hook.GithubActionsRefBehavior) {
			s.Equal("use-app-version-as-ref", *hook.GithubActionsRefBehavior)
		}

		s.Run("read", func() {
			hookAgain := &GithubActionsDeployHook{Model: gorm.Model{ID: hook.ID}}
			s.NoError(s.DB.Take(&hookAgain).Error)
			if s.NotNil(hookAgain.GithubActionsRefBehavior) {
				s.Equal("use-app-version-as-ref", *hookAgain.GithubActionsRefBehavior)
			}
		})

	})

	s.Run("edits trigger", func() {
		s.NoError(s.DB.Model(&hook.Trigger).Updates(&DeployHookTriggerConfig{OnSuccess: testutils.PointerTo(true)}).Error)
		if s.NotNil(hook.Trigger.OnSuccess) {
			s.True(*hook.Trigger.OnSuccess)
		}

		s.Run("read", func() {
			s.Run("empty without association", func() {
				hookAgain := &GithubActionsDeployHook{Model: gorm.Model{ID: hook.ID}}
				s.NoError(s.DB.Take(&hookAgain).Error)
				s.Nil(hookAgain.Trigger.OnSuccess)
			})
			s.Run("filled with association", func() {
				hookAgain := &GithubActionsDeployHook{Model: gorm.Model{ID: hook.ID}}
				s.NoError(s.DB.Preload("Trigger").Take(&hookAgain).Error)
				if s.NotNil(hookAgain.Trigger.OnSuccess) {
					s.True(*hookAgain.Trigger.OnSuccess)
				}
			})
			s.Run("filled with preload all", func() {
				hookAgain := &GithubActionsDeployHook{Model: gorm.Model{ID: hook.ID}}
				s.NoError(s.DB.Preload(clause.Associations).Take(&hookAgain).Error)
				if s.NotNil(hookAgain.Trigger.OnSuccess) {
					s.True(*hookAgain.Trigger.OnSuccess)
				}
			})
		})
	})

	s.Run("delete", func() {
		s.NoError(s.DB.Select("Trigger").Delete(&GithubActionsDeployHook{Model: gorm.Model{ID: hook.ID}}).Error)

		s.Run("hook gone", func() {
			var matchingHooks []GithubActionsDeployHook
			s.NoError(s.DB.Where(&GithubActionsDeployHook{Model: gorm.Model{ID: hook.ID}}).Find(&matchingHooks).Error)
			s.Len(matchingHooks, 0)
		})
		s.Run("trigger gone", func() {
			var matchingTriggers []DeployHookTriggerConfig
			s.NoError(s.DB.Where(&DeployHookTriggerConfig{Model: gorm.Model{ID: hook.Trigger.ID}}).Find(&matchingTriggers).Error)
			s.Len(matchingTriggers, 0)
		})
	})
}
