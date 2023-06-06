package v2models

import (
	"fmt"
	"github.com/broadinstitute/sherlock/internal/testutils"
	"gorm.io/gorm"
	"reflect"
	"testing"
	"time"
)

func Test_ciRunSelectorToQuery(t *testing.T) {
	type args struct {
		db       *gorm.DB
		selector string
	}
	tests := []struct {
		name    string
		args    args
		want    CiRun
		wantErr bool
	}{
		{
			name:    "empty",
			args:    args{selector: ""},
			wantErr: true,
		},
		{
			name:    "invalid",
			args:    args{selector: "something obviously invalid!"},
			wantErr: true,
		},
		{
			name: "valid id",
			args: args{selector: "123"},
			want: CiRun{Model: gorm.Model{ID: 123}},
		},
		{
			name:    "invalid id",
			args:    args{selector: testutils.StringNumberTooBigForInt},
			wantErr: true,
		},
		{
			name:    "gha + empty owner",
			args:    args{selector: "github-actions//beehive/1/1"},
			wantErr: true,
		},
		{
			name:    "gha + empty repo",
			args:    args{selector: "github-actions/broadinstitute//1/1"},
			wantErr: true,
		},
		{
			name:    "gha + bad run id",
			args:    args{selector: fmt.Sprintf("github-actions/broadinstitute/beehive/%s/1", testutils.StringNumberTooBigForInt)},
			wantErr: true,
		},
		{
			name:    "gha + bad attempt number",
			args:    args{selector: fmt.Sprintf("github-actions/broadinstitute/beehive/1/%s", testutils.StringNumberTooBigForInt)},
			wantErr: true,
		},
		{
			name: "valid gha",
			args: args{selector: "github-actions/broadinstitute/beehive/2/1"},
			want: CiRun{
				Platform:                   "github-actions",
				GithubActionsOwner:         "broadinstitute",
				GithubActionsRepo:          "beehive",
				GithubActionsRunID:         2,
				GithubActionsAttemptNumber: 1,
			},
		},
		{
			name:    "argo + empty namespace",
			args:    args{selector: "argo-workflows//name"},
			wantErr: true,
		},
		{
			name:    "argo + empty name",
			args:    args{selector: "argo-workflows/namespace/"},
			wantErr: true,
		},
		{
			name: "valid argo",
			args: args{selector: "argo-workflows/namespace/name"},
			want: CiRun{
				Platform:               "argo-workflows",
				ArgoWorkflowsNamespace: "namespace",
				ArgoWorkflowsName:      "name",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ciRunSelectorToQuery(tt.args.db, tt.args.selector)
			if (err != nil) != tt.wantErr {
				t.Errorf("ciRunSelectorToQuery() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ciRunSelectorToQuery() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ciRunToSelectors(t *testing.T) {
	type args struct {
		ciRun *CiRun
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "nil",
			args: args{ciRun: nil},
			want: nil,
		},
		{
			name: "none",
			args: args{ciRun: &CiRun{}},
			want: nil,
		},
		{
			name: "id",
			args: args{ciRun: &CiRun{
				Model: gorm.Model{ID: 1},
			}},
			want: []string{"1"},
		},
		{
			name: "gha",
			args: args{ciRun: &CiRun{
				Platform:                   "github-actions",
				GithubActionsOwner:         "broadinstitute",
				GithubActionsRepo:          "beehive",
				GithubActionsRunID:         123,
				GithubActionsAttemptNumber: 1,
			}},
			want: []string{"github-actions/broadinstitute/beehive/123/1"},
		},
		{
			name: "argo",
			args: args{ciRun: &CiRun{
				Platform:               "argo-workflows",
				ArgoWorkflowsNamespace: "namespace",
				ArgoWorkflowsName:      "name",
			}},
			want: []string{"argo-workflows/namespace/name"},
		},
		{
			name: "id + gha",
			args: args{ciRun: &CiRun{
				Model:                      gorm.Model{ID: 1},
				Platform:                   "github-actions",
				GithubActionsOwner:         "broadinstitute",
				GithubActionsRepo:          "beehive",
				GithubActionsRunID:         123,
				GithubActionsAttemptNumber: 1,
			}},
			want: []string{"1", "github-actions/broadinstitute/beehive/123/1"},
		},
		{
			name: "id + argo",
			args: args{ciRun: &CiRun{
				Model:                  gorm.Model{ID: 1},
				Platform:               "argo-workflows",
				ArgoWorkflowsNamespace: "namespace",
				ArgoWorkflowsName:      "name",
			}},
			want: []string{"1", "argo-workflows/namespace/name"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ciRunToSelectors(tt.args.ciRun); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ciRunToSelectors() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validateCiRun(t *testing.T) {
	type args struct {
		ciRun *CiRun
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "nil",
			args:    args{ciRun: nil},
			wantErr: true,
		},
		{
			name:    "empty",
			args:    args{ciRun: &CiRun{}},
			wantErr: true,
		},
		{
			name: "invalid platform",
			args: args{ciRun: &CiRun{
				Platform: "foobar",
			}},
			wantErr: true,
		},
		{
			name: "valid GHA",
			args: args{ciRun: &CiRun{
				Platform:                   "github-actions",
				GithubActionsOwner:         "broadinstitute",
				GithubActionsRepo:          "beehive",
				GithubActionsRunID:         123,
				GithubActionsAttemptNumber: 1,
				GithubActionsWorkflowPath:  ".github/workflows/build.yaml",
			}},
			wantErr: false,
		},
		{
			name: "GHA missing owner",
			args: args{ciRun: &CiRun{
				Platform:                   "github-actions",
				GithubActionsRepo:          "beehive",
				GithubActionsRunID:         123,
				GithubActionsAttemptNumber: 1,
				GithubActionsWorkflowPath:  ".github/workflows/build.yaml",
			}},
			wantErr: true,
		},
		{
			name: "GHA missing repo",
			args: args{ciRun: &CiRun{
				Platform:                   "github-actions",
				GithubActionsOwner:         "broadinstitute",
				GithubActionsRunID:         123,
				GithubActionsAttemptNumber: 1,
				GithubActionsWorkflowPath:  ".github/workflows/build.yaml",
			}},
			wantErr: true,
		},
		{
			name: "GHA missing run ID",
			args: args{ciRun: &CiRun{
				Platform:                   "github-actions",
				GithubActionsOwner:         "broadinstitute",
				GithubActionsRepo:          "beehive",
				GithubActionsAttemptNumber: 1,
				GithubActionsWorkflowPath:  ".github/workflows/build.yaml",
			}},
			wantErr: true,
		},
		{
			name: "GHA missing attempt number",
			args: args{ciRun: &CiRun{
				Platform:                  "github-actions",
				GithubActionsOwner:        "broadinstitute",
				GithubActionsRepo:         "beehive",
				GithubActionsRunID:        123,
				GithubActionsWorkflowPath: ".github/workflows/build.yaml",
			}},
			wantErr: true,
		},
		{
			name: "GHA missing workflow path",
			args: args{ciRun: &CiRun{
				Platform:                   "github-actions",
				GithubActionsOwner:         "broadinstitute",
				GithubActionsRepo:          "beehive",
				GithubActionsRunID:         123,
				GithubActionsAttemptNumber: 1,
			}},
			wantErr: true,
		},
		{
			name: "GHA with argo namespace",
			args: args{ciRun: &CiRun{
				Platform:                   "github-actions",
				GithubActionsOwner:         "broadinstitute",
				GithubActionsRepo:          "beehive",
				GithubActionsRunID:         123,
				GithubActionsAttemptNumber: 1,
				GithubActionsWorkflowPath:  ".github/workflows/build.yaml",
				ArgoWorkflowsNamespace:     "namespace",
			}},
			wantErr: true,
		},
		{
			name: "GHA with argo name",
			args: args{ciRun: &CiRun{
				Platform:                   "github-actions",
				GithubActionsOwner:         "broadinstitute",
				GithubActionsRepo:          "beehive",
				GithubActionsRunID:         123,
				GithubActionsAttemptNumber: 1,
				GithubActionsWorkflowPath:  ".github/workflows/build.yaml",
				ArgoWorkflowsName:          "name",
			}},
			wantErr: true,
		},
		{
			name: "GHA with argo template",
			args: args{ciRun: &CiRun{
				Platform:                   "github-actions",
				GithubActionsOwner:         "broadinstitute",
				GithubActionsRepo:          "beehive",
				GithubActionsRunID:         123,
				GithubActionsAttemptNumber: 1,
				GithubActionsWorkflowPath:  ".github/workflows/build.yaml",
				ArgoWorkflowsTemplate:      "template",
			}},
			wantErr: true,
		},
		{
			name: "valid argo",
			args: args{ciRun: &CiRun{
				Platform:               "argo-workflows",
				ArgoWorkflowsNamespace: "namespace",
				ArgoWorkflowsName:      "name",
				ArgoWorkflowsTemplate:  "template",
			}},
			wantErr: false,
		},
		{
			name: "argo missing namespace",
			args: args{ciRun: &CiRun{
				Platform:              "argo-workflows",
				ArgoWorkflowsName:     "name",
				ArgoWorkflowsTemplate: "template",
			}},
			wantErr: true,
		},
		{
			name: "argo missing name",
			args: args{ciRun: &CiRun{
				Platform:               "argo-workflows",
				ArgoWorkflowsNamespace: "namespace",
				ArgoWorkflowsTemplate:  "template",
			}},
			wantErr: true,
		},
		{
			name: "argo missing template",
			args: args{ciRun: &CiRun{
				Platform:               "argo-workflows",
				ArgoWorkflowsNamespace: "namespace",
				ArgoWorkflowsName:      "name",
			}},
			wantErr: true,
		},
		{
			name: "argo with GHA owner",
			args: args{ciRun: &CiRun{
				Platform:               "argo-workflows",
				ArgoWorkflowsNamespace: "namespace",
				ArgoWorkflowsName:      "name",
				ArgoWorkflowsTemplate:  "template",
				GithubActionsOwner:     "broadinstitute",
			}},
			wantErr: true,
		},
		{
			name: "argo with GHA repo",
			args: args{ciRun: &CiRun{
				Platform:               "argo-workflows",
				ArgoWorkflowsNamespace: "namespace",
				ArgoWorkflowsName:      "name",
				ArgoWorkflowsTemplate:  "template",
				GithubActionsRepo:      "beehive",
			}},
			wantErr: true,
		},
		{
			name: "argo with GHA run ID",
			args: args{ciRun: &CiRun{
				Platform:               "argo-workflows",
				ArgoWorkflowsNamespace: "namespace",
				ArgoWorkflowsName:      "name",
				ArgoWorkflowsTemplate:  "template",
				GithubActionsRunID:     123,
			}},
			wantErr: true,
		},
		{
			name: "argo with GHA attempt number",
			args: args{ciRun: &CiRun{
				Platform:                   "argo-workflows",
				ArgoWorkflowsNamespace:     "namespace",
				ArgoWorkflowsName:          "name",
				ArgoWorkflowsTemplate:      "template",
				GithubActionsAttemptNumber: 1,
			}},
			wantErr: true,
		},
		{
			name: "argo with GHA workflow path",
			args: args{ciRun: &CiRun{
				Platform:                  "argo-workflows",
				ArgoWorkflowsNamespace:    "namespace",
				ArgoWorkflowsName:         "name",
				ArgoWorkflowsTemplate:     "template",
				GithubActionsWorkflowPath: ".github/workflows/build.yaml",
			}},
			wantErr: true,
		},
		{
			name: "valid terminal GHA",
			args: args{ciRun: &CiRun{
				Platform:                   "github-actions",
				GithubActionsOwner:         "broadinstitute",
				GithubActionsRepo:          "beehive",
				GithubActionsRunID:         123,
				GithubActionsAttemptNumber: 1,
				GithubActionsWorkflowPath:  ".github/workflows/build.yaml",
				TerminalAt:                 testutils.PointerTo(time.Now()),
				Status:                     testutils.PointerTo("failed"),
			}},
			wantErr: false,
		},
		{
			name: "valid terminal argo",
			args: args{ciRun: &CiRun{
				Platform:               "argo-workflows",
				ArgoWorkflowsNamespace: "namespace",
				ArgoWorkflowsName:      "name",
				ArgoWorkflowsTemplate:  "template",
				TerminalAt:             testutils.PointerTo(time.Now()),
				Status:                 testutils.PointerTo("failed"),
			}},
			wantErr: false,
		},
		{
			name: "terminal GHA without status",
			args: args{ciRun: &CiRun{
				Platform:                   "github-actions",
				GithubActionsOwner:         "broadinstitute",
				GithubActionsRepo:          "beehive",
				GithubActionsRunID:         123,
				GithubActionsAttemptNumber: 1,
				GithubActionsWorkflowPath:  ".github/workflows/build.yaml",
				TerminalAt:                 testutils.PointerTo(time.Now()),
			}},
			wantErr: true,
		},
		{
			name: "terminal argo without status",
			args: args{ciRun: &CiRun{
				Platform:               "argo-workflows",
				ArgoWorkflowsNamespace: "namespace",
				ArgoWorkflowsName:      "name",
				ArgoWorkflowsTemplate:  "template",
				TerminalAt:             testutils.PointerTo(time.Now()),
			}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateCiRun(tt.args.ciRun); (err != nil) != tt.wantErr {
				t.Errorf("validateCiRun() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
