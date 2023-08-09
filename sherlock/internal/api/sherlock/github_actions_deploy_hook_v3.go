package sherlock

import (
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type GithubActionsDeployHookV3 struct {
	commonFields
	deployHookTriggerConfigV3
	GithubActionsOwner          *string         `json:"githubActionsOwner,omitempty" form:"githubActionsOwner"`
	GithubActionsRepo           *string         `json:"githubActionsRepo,omitempty" form:"githubActionsRepo"`
	GithubActionsWorkflowPath   *string         `json:"githubActionsWorkflowPath,omitempty" form:"githubActionsWorkflowPath"`
	GithubActionsDefaultRef     *string         `json:"githubActionsDefaultRef,omitempty" form:"githubActionsDefaultRef"`
	GithubActionsRefBehavior    *string         `json:"githubActionsRefBehavior,omitempty" form:"githubActionsRefBehavior" enums:"always-use-default-ref,use-app-version-as-ref,use-app-version-commit-as-ref" default:"always-use-default-ref" binding:"oneof=always-use-default-ref use-app-version-as-ref use-app-version-commit-as-ref ''"`
	GithubActionsWorkflowInputs *datatypes.JSON `json:"githubActionsWorkflowInputs,omitempty" form:"githubActionsWorkflowInputs" swaggertype:"object"`
}

func (g GithubActionsDeployHookV3) toModel(db *gorm.DB) (models.GithubActionsDeployHook, error) {
	ret := models.GithubActionsDeployHook{
		Model:                       g.commonFields.toGormModel(),
		GithubActionsOwner:          g.GithubActionsOwner,
		GithubActionsRepo:           g.GithubActionsRepo,
		GithubActionsWorkflowPath:   g.GithubActionsWorkflowPath,
		GithubActionsDefaultRef:     g.GithubActionsDefaultRef,
		GithubActionsRefBehavior:    g.GithubActionsRefBehavior,
		GithubActionsWorkflowInputs: g.GithubActionsWorkflowInputs,
	}
	var err error
	ret.Trigger, err = g.deployHookTriggerConfigV3.toModel(db)
	return ret, err
}

func githubActionsDeployHookFromModel(model models.GithubActionsDeployHook) GithubActionsDeployHookV3 {
	return GithubActionsDeployHookV3{
		commonFields:                commonFieldsFromGormModel(model.Model),
		deployHookTriggerConfigV3:   deployHookTriggerConfigFromModel(model.Trigger),
		GithubActionsOwner:          model.GithubActionsOwner,
		GithubActionsRepo:           model.GithubActionsRepo,
		GithubActionsWorkflowPath:   model.GithubActionsWorkflowPath,
		GithubActionsDefaultRef:     model.GithubActionsDefaultRef,
		GithubActionsRefBehavior:    model.GithubActionsRefBehavior,
		GithubActionsWorkflowInputs: model.GithubActionsWorkflowInputs,
	}
}
