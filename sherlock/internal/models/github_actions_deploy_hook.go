package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type GithubActionsDeployHook struct {
	gorm.Model
	Trigger                     DeployHookTriggerConfig `gorm:"polymorphic:Hook;polymorphicValue:github-actions"`
	GithubActionsOwner          *string
	GithubActionsRepo           *string
	GithubActionsWorkflowPath   *string
	GithubActionsDefaultRef     *string
	GithubActionsRefBehavior    *string // Must be "always-use-default-ref", "use-app-version-as-ref", or "use-app-version-commit-as-ref"
	GithubActionsWorkflowInputs *datatypes.JSON
}

func (g *GithubActionsDeployHook) TableName() string {
	return "github_actions_deploy_hooks"
}

func (g *GithubActionsDeployHook) AfterSave(tx *gorm.DB) error {
	return g.Trigger.ErrorIfForbidden(tx)
}

func (g *GithubActionsDeployHook) AfterDelete(tx *gorm.DB) error {
	return g.Trigger.ErrorIfForbidden(tx)
}
