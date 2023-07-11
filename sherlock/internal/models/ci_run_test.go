package models

import (
	"github.com/broadinstitute/sherlock/go-shared/pkg/testutils"
	"time"
)

func (s *modelSuite) TestCiRunPlatformValidationSqlPlatformInvalid() {
	err := s.db.Create(&CiRun{
		Platform: "invalid",
	}).Error
	s.ErrorContains(err, "violates check constraint \"platform_present\"")
}

func (s *modelSuite) TestCiRunPlatformValidationSqlGithubInvalid() {
	err := s.db.Create(&CiRun{
		Platform:                   "github-actions",
		GithubActionsOwner:         "owner",
		GithubActionsRepo:          "repo",
		GithubActionsRunID:         123,
		GithubActionsAttemptNumber: 123,
	}).Error
	s.ErrorContains(err, "violates check constraint \"platform_present\"")
}

func (s *modelSuite) TestCiRunPlatformValidationSqlGithubValid() {
	err := s.db.Create(&CiRun{
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
	err := s.db.Create(&CiRun{
		Platform:               "argo-workflows",
		ArgoWorkflowsNamespace: "namespace",
		ArgoWorkflowsName:      "name",
	}).Error
	s.ErrorContains(err, "violates check constraint \"platform_present\"")
}

func (s *modelSuite) TestCiRunPlatformValidationSqlArgoValid() {
	err := s.db.Create(&CiRun{
		Platform:               "argo-workflows",
		ArgoWorkflowsNamespace: "namespace",
		ArgoWorkflowsName:      "name",
		ArgoWorkflowsTemplate:  "template",
	}).Error
	s.NoError(err)
}

func (s *modelSuite) TestCiRunPlatformValidationSqlBoth() {
	err := s.db.Create(&CiRun{
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
	err := s.db.Create(&run).Error
	s.ErrorContains(err, "violates check constraint \"terminal_status_present\"")
}

func (s *modelSuite) TestCiRunTerminalValidationValid() {
	err := s.db.Create(&CiRun{
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
