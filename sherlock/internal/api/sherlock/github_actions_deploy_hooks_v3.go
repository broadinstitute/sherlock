package sherlock

import (
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type GithubActionsDeployHookV3 struct {
	commonFields
	deployHookTriggerConfigV3
	githubActionsDeployHookFields
}

type githubActionsDeployHookFields struct {
	GithubActionsOwner        *string `json:"githubActionsOwner,omitempty" form:"githubActionsOwner"`
	GithubActionsRepo         *string `json:"githubActionsRepo,omitempty" form:"githubActionsRepo"`
	GithubActionsWorkflowPath *string `json:"githubActionsWorkflowPath,omitempty" form:"githubActionsWorkflowPath"`
	GithubActionsDefaultRef   *string `json:"githubActionsDefaultRef,omitempty" form:"githubActionsDefaultRef"`
	// This field determines what git ref the workflow will be run on. The default of always-use-default-ref always uses the default ref; use-app-version-as-ref will use the app version (when available) as the ref, useful when versions are always commit hashes or tags; use-app-version-commit-as-ref will use the app version's commit (when available) as the ref, useful when the repo is configured to fully report app versions to Sherlock.
	GithubActionsRefBehavior *string `json:"githubActionsRefBehavior,omitempty" form:"githubActionsRefBehavior" enums:"always-use-default-ref,use-app-version-as-ref,use-app-version-commit-as-ref" default:"always-use-default-ref" binding:"omitempty,oneof=always-use-default-ref use-app-version-as-ref use-app-version-commit-as-ref"`
	// These workflow inputs will be passed statically as-is to GitHub's workflow dispatch API (https://docs.github.com/en/rest/actions/workflows#create-a-workflow-dispatch-event) as the `inputs` parameter object.
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
		commonFields:              commonFieldsFromGormModel(model.Model),
		deployHookTriggerConfigV3: deployHookTriggerConfigFromModel(model.Trigger),
		githubActionsDeployHookFields: githubActionsDeployHookFields{
			GithubActionsOwner:          model.GithubActionsOwner,
			GithubActionsRepo:           model.GithubActionsRepo,
			GithubActionsWorkflowPath:   model.GithubActionsWorkflowPath,
			GithubActionsDefaultRef:     model.GithubActionsDefaultRef,
			GithubActionsRefBehavior:    model.GithubActionsRefBehavior,
			GithubActionsWorkflowInputs: model.GithubActionsWorkflowInputs,
		},
	}
}
