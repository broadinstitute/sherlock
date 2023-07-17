package sherlock

import (
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"time"
)

type CiRunV3 struct {
	commonFields
	ciRunFields
	RelatedResources []CiIdentifierV3 `json:"relatedResources" form:"-"`
}

type ciRunFields struct {
	Platform                   string     `json:"platform" form:"platform"`
	GithubActionsOwner         string     `json:"githubActionsOwner" form:"githubActionsOwner"`
	GithubActionsRepo          string     `json:"githubActionsRepo" form:"githubActionsRepo"`
	GithubActionsRunID         uint       `json:"githubActionsRunID" form:"githubActionsRunID"`
	GithubActionsAttemptNumber uint       `json:"githubActionsAttemptNumber" form:"githubActionsAttemptNumber"`
	GithubActionsWorkflowPath  string     `json:"githubActionsWorkflowPath" form:"githubActionsWorkflowPath"`
	ArgoWorkflowsNamespace     string     `json:"argoWorkflowsNamespace" form:"argoWorkflowsNamespace"`
	ArgoWorkflowsName          string     `json:"argoWorkflowsName" form:"argoWorkflowsName"`
	ArgoWorkflowsTemplate      string     `json:"argoWorkflowsTemplate" form:"argoWorkflowsTemplate"`
	StartedAt                  *time.Time `json:"startedAt,omitempty" form:"startedAt"`
	TerminalAt                 *time.Time `json:"terminalAt,omitempty" form:"terminalAt"`
	Status                     *string    `json:"status,omitempty" form:"status"`
}

func (c CiRunV3) toModel() models.CiRun {
	return models.CiRun{
		Model:                      c.toGormModel(),
		Platform:                   c.Platform,
		GithubActionsOwner:         c.GithubActionsOwner,
		GithubActionsRepo:          c.GithubActionsRepo,
		GithubActionsRunID:         c.GithubActionsRunID,
		GithubActionsAttemptNumber: c.GithubActionsAttemptNumber,
		GithubActionsWorkflowPath:  c.GithubActionsWorkflowPath,
		ArgoWorkflowsNamespace:     c.ArgoWorkflowsNamespace,
		ArgoWorkflowsName:          c.ArgoWorkflowsName,
		ArgoWorkflowsTemplate:      c.ArgoWorkflowsTemplate,
		StartedAt:                  c.StartedAt,
		TerminalAt:                 c.TerminalAt,
		Status:                     c.Status,
	}
}

func ciRunFromModel(model models.CiRun) CiRunV3 {
	var relatedResources []CiIdentifierV3
	if len(model.RelatedResources) > 0 {
		relatedResources = make([]CiIdentifierV3, len(model.RelatedResources))
		for index, modelCiIdentifier := range model.RelatedResources {
			relatedResources[index] = ciIdentifierFromModel(modelCiIdentifier)
		}
	}
	return CiRunV3{
		commonFields: commonFieldsFromGormModel(model.Model),
		ciRunFields: ciRunFields{
			Platform:                   model.Platform,
			GithubActionsOwner:         model.GithubActionsOwner,
			GithubActionsRepo:          model.GithubActionsRepo,
			GithubActionsRunID:         model.GithubActionsRunID,
			GithubActionsAttemptNumber: model.GithubActionsAttemptNumber,
			GithubActionsWorkflowPath:  model.GithubActionsWorkflowPath,
			ArgoWorkflowsNamespace:     model.ArgoWorkflowsNamespace,
			ArgoWorkflowsName:          model.ArgoWorkflowsName,
			ArgoWorkflowsTemplate:      model.ArgoWorkflowsTemplate,
			StartedAt:                  model.StartedAt,
			TerminalAt:                 model.TerminalAt,
			Status:                     model.Status,
		},
		RelatedResources: relatedResources,
	}
}
