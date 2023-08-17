package deployhooks

import (
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCiRunIsDeploy(t *testing.T) {
	config.LoadTestConfig()
	assert.NoError(t, Init())
	type args struct {
		ciRun models.CiRun
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "matches",
			args: args{ciRun: models.CiRun{
				Platform:                   "github-actions",
				GithubActionsOwner:         "broadinstitute",
				GithubActionsRepo:          "terra-github-workflows",
				GithubActionsRunID:         123123,
				GithubActionsAttemptNumber: 1,
				GithubActionsWorkflowPath:  ".github/workflows/sync-release.yaml",
			}},
			want: true,
		},
		{
			name: "no platform match",
			args: args{ciRun: models.CiRun{
				Platform:               "argo-workflows",
				ArgoWorkflowsNamespace: "namespace",
				ArgoWorkflowsName:      "name",
				ArgoWorkflowsTemplate:  "template",
			}},
			want: false,
		},
		{
			name: "no owner match",
			args: args{ciRun: models.CiRun{
				Platform:                   "github-actions",
				GithubActionsOwner:         "DataBiosphere",
				GithubActionsRepo:          "terra-github-workflows",
				GithubActionsRunID:         123123,
				GithubActionsAttemptNumber: 1,
				GithubActionsWorkflowPath:  ".github/workflows/sync-release.yaml",
			}},
			want: false,
		},
		{
			name: "no repo match",
			args: args{ciRun: models.CiRun{
				Platform:                   "github-actions",
				GithubActionsOwner:         "broadinstitute",
				GithubActionsRepo:          "terra-helmfile",
				GithubActionsRunID:         123123,
				GithubActionsAttemptNumber: 1,
				GithubActionsWorkflowPath:  ".github/workflows/sync-release.yaml",
			}},
			want: false,
		},
		{
			name: "no path match",
			args: args{ciRun: models.CiRun{
				Platform:                   "github-actions",
				GithubActionsOwner:         "broadinstitute",
				GithubActionsRepo:          "terra-github-workflows",
				GithubActionsRunID:         123123,
				GithubActionsAttemptNumber: 1,
				GithubActionsWorkflowPath:  ".github/workflows/bee-create.yaml",
			}},
			want: false,
		},
		{
			name: "matches even if otherwise required fields are missing",
			args: args{ciRun: models.CiRun{
				Platform:                  "github-actions",
				GithubActionsOwner:        "broadinstitute",
				GithubActionsRepo:         "terra-github-workflows",
				GithubActionsWorkflowPath: ".github/workflows/sync-release.yaml",
			}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, CiRunIsDeploy(tt.args.ciRun), "CiRunIsDeploy(%v)", tt.args.ciRun)
		})
	}
}
