package models

import (
	"gorm.io/gorm"
	"time"
)

type CiRun struct {
	gorm.Model
	Platform                   string
	GithubActionsOwner         string
	GithubActionsRepo          string
	GithubActionsRunID         uint
	GithubActionsAttemptNumber uint
	GithubActionsWorkflowPath  string
	ArgoWorkflowsNamespace     string
	ArgoWorkflowsName          string
	ArgoWorkflowsTemplate      string
	// Mutable
	RelatedResources []CiIdentifier `gorm:"many2many:v2_ci_runs_for_identifiers"`
	StartedAt        *time.Time
	TerminalAt       *time.Time
	Status           *string
}

func (c *CiRun) TableName() string {
	return "v2_ci_runs"
}
