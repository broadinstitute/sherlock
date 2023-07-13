package models

import (
	"github.com/broadinstitute/sherlock/go-shared/pkg/testutils"
	"time"
)

func (s *modelSuite) TestCiRunPlatformValidationSqlPlatformInvalid() {
	err := s.DB.Create(&CiRun{
		Platform: "invalid",
	}).Error
	s.ErrorContains(err, "violates check constraint \"platform_present\"")
}

func (s *modelSuite) TestCiRunPlatformValidationSqlGithubInvalid() {
	err := s.DB.Create(&CiRun{
		Platform:                   "github-actions",
		GithubActionsOwner:         "owner",
		GithubActionsRepo:          "repo",
		GithubActionsRunID:         123,
		GithubActionsAttemptNumber: 123,
	}).Error
	s.ErrorContains(err, "violates check constraint \"platform_present\"")
}

func (s *modelSuite) TestCiRunPlatformValidationSqlGithubValid() {
	err := s.DB.Create(&CiRun{
		Platform:                   "github-actions",
		GithubActionsOwner:         "owner",
		GithubActionsRepo:          "repo",
		GithubActionsRunID:         123,
		GithubActionsAttemptNumber: 123,
		GithubActionsWorkflowPath:  "path",
	}).Error
	s.NoError(err)
}

func (s *modelSuite) TestCiRunPlatformValidationSqlArgoInvalid() {
	err := s.DB.Create(&CiRun{
		Platform:               "argo-workflows",
		ArgoWorkflowsNamespace: "namespace",
		ArgoWorkflowsName:      "name",
	}).Error
	s.ErrorContains(err, "violates check constraint \"platform_present\"")
}

func (s *modelSuite) TestCiRunPlatformValidationSqlArgoValid() {
	err := s.DB.Create(&CiRun{
		Platform:               "argo-workflows",
		ArgoWorkflowsNamespace: "namespace",
		ArgoWorkflowsName:      "name",
		ArgoWorkflowsTemplate:  "template",
	}).Error
	s.NoError(err)
}

func (s *modelSuite) TestCiRunPlatformValidationSqlBoth() {
	err := s.DB.Create(&CiRun{
		Platform:                   "argo-workflows",
		GithubActionsOwner:         "owner",
		GithubActionsRepo:          "repo",
		GithubActionsRunID:         123,
		GithubActionsAttemptNumber: 123,
		GithubActionsWorkflowPath:  "path",
		ArgoWorkflowsNamespace:     "namespace",
		ArgoWorkflowsName:          "name",
		ArgoWorkflowsTemplate:      "template",
	}).Error
	s.ErrorContains(err, "violates check constraint \"platform_present\"")
}

func (s *modelSuite) TestCiRunTerminalValidationInvalid() {
	run := CiRun{
		Platform:                   "github-actions",
		GithubActionsOwner:         "owner",
		GithubActionsRepo:          "repo",
		GithubActionsRunID:         123,
		GithubActionsAttemptNumber: 123,
		GithubActionsWorkflowPath:  "path",
		TerminalAt:                 testutils.PointerTo(time.Now()),
	}
	err := s.DB.Create(&run).Error
	s.ErrorContains(err, "violates check constraint \"terminal_status_present\"")
}

func (s *modelSuite) TestCiRunTerminalValidationValid() {
	err := s.DB.Create(&CiRun{
		Platform:                   "github-actions",
		GithubActionsOwner:         "owner",
		GithubActionsRepo:          "repo",
		GithubActionsRunID:         123,
		GithubActionsAttemptNumber: 123,
		GithubActionsWorkflowPath:  "path",
		TerminalAt:                 testutils.PointerTo(time.Now()),
		Status:                     testutils.PointerTo("status"),
	}).Error
	s.NoError(err)
}

func (s *modelSuite) TestCiRunUniquenessSqlInvalid() {
	err := s.DB.Create(&CiRun{
		Platform:                   "github-actions",
		GithubActionsOwner:         "owner",
		GithubActionsRepo:          "repo",
		GithubActionsRunID:         123,
		GithubActionsAttemptNumber: 123,
		GithubActionsWorkflowPath:  "path",
	}).Error
	s.NoError(err)
	err = s.DB.Create(&CiRun{
		Platform:                   "github-actions",
		GithubActionsOwner:         "owner",
		GithubActionsRepo:          "repo",
		GithubActionsRunID:         123,
		GithubActionsAttemptNumber: 123,
		GithubActionsWorkflowPath:  "path",
	}).Error
	s.ErrorContains(err, "violates unique constraint")
}

func (s *modelSuite) TestCiRunUniquenessSqlValid() {
	run1 := CiRun{
		Platform:                   "github-actions",
		GithubActionsOwner:         "owner",
		GithubActionsRepo:          "repo",
		GithubActionsRunID:         123,
		GithubActionsAttemptNumber: 123,
		GithubActionsWorkflowPath:  "path",
	}
	err := s.DB.Create(&run1).Error
	s.NoError(err)
	s.NotZero(run1.ID)
	err = s.DB.Delete(&run1).Error
	s.NoError(err)
	run2 := CiRun{
		Platform:                   "github-actions",
		GithubActionsOwner:         "owner",
		GithubActionsRepo:          "repo",
		GithubActionsRunID:         123,
		GithubActionsAttemptNumber: 123,
		GithubActionsWorkflowPath:  "path",
	}
	err = s.DB.Create(&run2).Error
	s.NoError(err)
	s.NotZero(run2.ID)
	s.NotEqual(run1.ID, run2.ID)
}
