package sherlock

import (
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"time"
)

type GitCommitV3 struct {
	CommonFields
	GitRepo      string    `json:"gitRepo"`
	GitCommit    string    `json:"gitCommit"`
	GitBranch    string    `json:"gitBranch"`
	IsMainBranch bool      `json:"isMainBranch"`
	SecSincePrev *uint     `json:"secSincePrev"`
	CreatedAt    time.Time `json:"createdAt" form:"createdAt" format:"date-time"`
}

//nolint:unused
func (g GitCommitV3) toModel() models.GitCommit {
	return models.GitCommit{
		Model:        g.toGormModel(),
		GitRepo:      g.GitRepo,
		GitCommit:    g.GitCommit,
		GitBranch:    g.GitBranch,
		IsMainBranch: g.IsMainBranch,
		SecSincePrev: g.SecSincePrev,
	}
}

func gitCommitFromModel(model models.GitCommit) GitCommitV3 {
	return GitCommitV3{
		CommonFields: commonFieldsFromGormModel(model.Model),
		GitRepo:      model.GitRepo,
		GitCommit:    model.GitCommit,
		GitBranch:    model.GitBranch,
		IsMainBranch: model.IsMainBranch,
		SecSincePrev: model.SecSincePrev,
	}
}
