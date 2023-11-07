package models

import (
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (s *modelSuite) TestGithubActionsDeployHookEnvironment() {
	environment := s.TestData.Environment_Dev()
	hook := GithubActionsDeployHook{
		GithubActionsOwner:        utils.PointerTo("owner"),
		GithubActionsRepo:         utils.PointerTo("repo"),
		GithubActionsWorkflowPath: utils.PointerTo("path"),
		GithubActionsDefaultRef:   utils.PointerTo("head"),
		GithubActionsRefBehavior:  utils.PointerTo("always-use-default-ref"),
		Trigger:                   DeployHookTriggerConfig{OnEnvironmentID: &environment.ID}}
	s.NoError(s.DB.Create(&hook).Error)
}

func (s *modelSuite) TestGithubActionsDeployHookChartRelease() {
	chartRelease := s.TestData.ChartRelease_LeonardoDev()
	hook := GithubActionsDeployHook{
		GithubActionsOwner:        utils.PointerTo("owner"),
		GithubActionsRepo:         utils.PointerTo("repo"),
		GithubActionsWorkflowPath: utils.PointerTo("path"),
		GithubActionsDefaultRef:   utils.PointerTo("head"),
		GithubActionsRefBehavior:  utils.PointerTo("always-use-default-ref"),
		Trigger:                   DeployHookTriggerConfig{OnChartReleaseID: &chartRelease.ID}}
	s.NoError(s.DB.Create(&hook).Error)
}

func (s *modelSuite) TestGithubActionsDeployHookNoOwner() {
	environment := s.TestData.Environment_Dev()
	hook := GithubActionsDeployHook{
		GithubActionsRepo:         utils.PointerTo("repo"),
		GithubActionsWorkflowPath: utils.PointerTo("path"),
		GithubActionsDefaultRef:   utils.PointerTo("head"),
		GithubActionsRefBehavior:  utils.PointerTo("always-use-default-ref"),
		Trigger:                   DeployHookTriggerConfig{OnEnvironmentID: &environment.ID}}
	s.ErrorContains(s.DB.Create(&hook).Error, "github_actions_owner_present")
}

func (s *modelSuite) TestGithubActionsDeployHookNoRepo() {
	environment := s.TestData.Environment_Dev()
	hook := GithubActionsDeployHook{
		GithubActionsOwner:        utils.PointerTo("owner"),
		GithubActionsWorkflowPath: utils.PointerTo("path"),
		GithubActionsDefaultRef:   utils.PointerTo("head"),
		GithubActionsRefBehavior:  utils.PointerTo("always-use-default-ref"),
		Trigger:                   DeployHookTriggerConfig{OnEnvironmentID: &environment.ID}}
	s.ErrorContains(s.DB.Create(&hook).Error, "github_actions_repo_present")
}

func (s *modelSuite) TestGithubActionsDeployHookNoWorkflowPath() {
	environment := s.TestData.Environment_Dev()
	hook := GithubActionsDeployHook{
		GithubActionsOwner:       utils.PointerTo("owner"),
		GithubActionsRepo:        utils.PointerTo("repo"),
		GithubActionsDefaultRef:  utils.PointerTo("head"),
		GithubActionsRefBehavior: utils.PointerTo("always-use-default-ref"),
		Trigger:                  DeployHookTriggerConfig{OnEnvironmentID: &environment.ID}}
	s.ErrorContains(s.DB.Create(&hook).Error, "github_actions_workflow_path_present")
}

func (s *modelSuite) TestGithubActionsDeployHookNoDefaultRef() {
	environment := s.TestData.Environment_Dev()
	hook := GithubActionsDeployHook{
		GithubActionsOwner:        utils.PointerTo("owner"),
		GithubActionsRepo:         utils.PointerTo("repo"),
		GithubActionsWorkflowPath: utils.PointerTo("path"),
		GithubActionsRefBehavior:  utils.PointerTo("always-use-default-ref"),
		Trigger:                   DeployHookTriggerConfig{OnEnvironmentID: &environment.ID}}
	s.ErrorContains(s.DB.Create(&hook).Error, "github_actions_default_ref_present")
}

func (s *modelSuite) TestGithubActionsDeployHookNoRefBehavior() {
	environment := s.TestData.Environment_Dev()
	hook := GithubActionsDeployHook{
		GithubActionsOwner:        utils.PointerTo("owner"),
		GithubActionsRepo:         utils.PointerTo("repo"),
		GithubActionsWorkflowPath: utils.PointerTo("path"),
		GithubActionsDefaultRef:   utils.PointerTo("head"),
		Trigger:                   DeployHookTriggerConfig{OnEnvironmentID: &environment.ID}}
	s.ErrorContains(s.DB.Create(&hook).Error, "github_actions_ref_behavior_valid")
}

func (s *modelSuite) TestGithubActionsDeployHookBadRefBehavior() {
	environment := s.TestData.Environment_Dev()
	hook := GithubActionsDeployHook{
		GithubActionsOwner:        utils.PointerTo("owner"),
		GithubActionsRepo:         utils.PointerTo("repo"),
		GithubActionsWorkflowPath: utils.PointerTo("path"),
		GithubActionsDefaultRef:   utils.PointerTo("head"),
		GithubActionsRefBehavior:  utils.PointerTo("some nonsense string"),
		Trigger:                   DeployHookTriggerConfig{OnEnvironmentID: &environment.ID}}
	s.ErrorContains(s.DB.Create(&hook).Error, "github_actions_ref_behavior_valid")
}

func (s *modelSuite) TestGithubActionsDeployHookSuitability() {
	environment := s.TestData.Environment_Prod()
	s.Run("when suitable", func() {
		s.SetSuitableTestUserForDB()
		hook := GithubActionsDeployHook{
			GithubActionsOwner:        utils.PointerTo("owner"),
			GithubActionsRepo:         utils.PointerTo("repo"),
			GithubActionsWorkflowPath: utils.PointerTo("path"),
			GithubActionsDefaultRef:   utils.PointerTo("head"),
			GithubActionsRefBehavior:  utils.PointerTo("always-use-default-ref"),
			Trigger:                   DeployHookTriggerConfig{OnEnvironmentID: &environment.ID}}
		s.NoError(s.DB.Create(&hook).Error)
	})
	s.Run("when not suitable", func() {
		s.SetNonSuitableTestUserForDB()
		hook := GithubActionsDeployHook{
			GithubActionsOwner:        utils.PointerTo("owner"),
			GithubActionsRepo:         utils.PointerTo("repo"),
			GithubActionsWorkflowPath: utils.PointerTo("path"),
			GithubActionsDefaultRef:   utils.PointerTo("head"),
			GithubActionsRefBehavior:  utils.PointerTo("always-use-default-ref"),
			Trigger:                   DeployHookTriggerConfig{OnEnvironmentID: &environment.ID}}
		s.ErrorContains(s.DB.Create(&hook).Error, errors.Forbidden)
	})
}

func (s *modelSuite) TestGithubActionsDeployHookFlow() {
	environment := s.TestData.Environment_Dev()
	hook := GithubActionsDeployHook{
		GithubActionsOwner:        utils.PointerTo("owner"),
		GithubActionsRepo:         utils.PointerTo("repo"),
		GithubActionsWorkflowPath: utils.PointerTo("path"),
		GithubActionsDefaultRef:   utils.PointerTo("head"),
		GithubActionsRefBehavior:  utils.PointerTo("always-use-default-ref"),
		Trigger:                   DeployHookTriggerConfig{OnEnvironmentID: &environment.ID}}
	s.NoError(s.DB.Create(&hook).Error)
	s.NotZero(hook.ID)
	s.NotZero(hook.Trigger.ID)

	s.Run("edits hook", func() {
		s.NoError(s.DB.Model(&hook).Updates(&GithubActionsDeployHook{GithubActionsRefBehavior: utils.PointerTo("use-app-version-as-ref")}).Error)
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
		s.NoError(s.DB.Model(&hook.Trigger).Updates(&DeployHookTriggerConfig{OnSuccess: utils.PointerTo(true)}).Error)
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
