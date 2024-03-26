package sherlock

import (
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"time"
)

type GithubActionsJobV3 struct {
	CommonFields
	GithubActionsJobV3Create
}

type GithubActionsJobV3Create struct {
	GithubActionsOwner         string `json:"githubActionsOwner"`
	GithubActionsRepo          string `json:"githubActionsRepo"`
	GithubActionsRunID         uint   `json:"githubActionsRunID"`
	GithubActionsAttemptNumber uint   `json:"githubActionsAttemptNumber"`
	GithubActionsJobID         uint   `json:"githubActionsJobID"`
	GithubActionsJobV3Edit
}

type GithubActionsJobV3Edit struct {
	JobCreatedAt  *time.Time `json:"jobCreatedAt"`
	JobStartedAt  *time.Time `json:"jobStartedAt"`
	JobTerminalAt *time.Time `json:"jobTerminalAt"`
	Status        *string    `json:"status"`
}

func (j GithubActionsJobV3) toModel() models.GithubActionsJob {
	return models.GithubActionsJob{
		Model:                      j.toGormModel(),
		GithubActionsOwner:         j.GithubActionsOwner,
		GithubActionsRepo:          j.GithubActionsRepo,
		GithubActionsRunID:         j.GithubActionsRunID,
		GithubActionsAttemptNumber: j.GithubActionsAttemptNumber,
		GithubActionsJobID:         j.GithubActionsJobID,
		JobCreatedAt:               j.JobCreatedAt,
		JobStartedAt:               j.JobStartedAt,
		JobTerminalAt:              j.JobTerminalAt,
		Status:                     j.Status,
	}
}

func (j GithubActionsJobV3Create) toModel() models.GithubActionsJob {
	return GithubActionsJobV3{GithubActionsJobV3Create: j}.toModel()
}

func (j GithubActionsJobV3Edit) toModel() models.GithubActionsJob {
	return GithubActionsJobV3Create{GithubActionsJobV3Edit: j}.toModel()
}

func githubActionsJobFromModel(model models.GithubActionsJob) GithubActionsJobV3 {
	return GithubActionsJobV3{
		CommonFields: commonFieldsFromGormModel(model.Model),
		GithubActionsJobV3Create: GithubActionsJobV3Create{
			GithubActionsOwner:         model.GithubActionsOwner,
			GithubActionsRepo:          model.GithubActionsRepo,
			GithubActionsRunID:         model.GithubActionsRunID,
			GithubActionsAttemptNumber: model.GithubActionsAttemptNumber,
			GithubActionsJobID:         model.GithubActionsJobID,
			GithubActionsJobV3Edit: GithubActionsJobV3Edit{
				JobCreatedAt:  model.JobCreatedAt,
				JobStartedAt:  model.JobStartedAt,
				JobTerminalAt: model.JobTerminalAt,
				Status:        model.Status,
			},
		},
	}
}
