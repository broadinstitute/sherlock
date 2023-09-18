package sherlock

import (
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
	"time"
)

func TestCiRunV3_toModel(t *testing.T) {
	t1 := time.Now().Add(-(time.Minute * 4))
	t2 := time.Now().Add(-(time.Minute * 3))
	t3 := time.Now().Add(-(time.Minute * 2))
	t4 := time.Now().Add(-time.Minute)
	type fields struct {
		commonFields     CommonFields
		ciRunFields      ciRunFields
		RelatedResources []CiIdentifierV3
	}
	tests := []struct {
		name   string
		fields fields
		want   models.CiRun
	}{
		{
			name: "equal",
			fields: fields{
				commonFields: CommonFields{
					ID:        1,
					CreatedAt: t1,
					UpdatedAt: t2,
				},
				ciRunFields: ciRunFields{
					Platform:                   "platform",
					GithubActionsOwner:         "owner",
					GithubActionsRepo:          "repo",
					GithubActionsRunID:         2,
					GithubActionsAttemptNumber: 3,
					GithubActionsWorkflowPath:  "path",
					ArgoWorkflowsNamespace:     "namespace",
					ArgoWorkflowsName:          "name",
					ArgoWorkflowsTemplate:      "template",
					StartedAt:                  &t3,
					TerminalAt:                 &t4,
					Status:                     utils.PointerTo("status"),
				},
				RelatedResources: nil,
			},
			want: models.CiRun{
				Model: gorm.Model{
					ID:        1,
					CreatedAt: t1,
					UpdatedAt: t2,
				},
				Platform:                   "platform",
				GithubActionsOwner:         "owner",
				GithubActionsRepo:          "repo",
				GithubActionsRunID:         2,
				GithubActionsAttemptNumber: 3,
				GithubActionsWorkflowPath:  "path",
				ArgoWorkflowsNamespace:     "namespace",
				ArgoWorkflowsName:          "name",
				ArgoWorkflowsTemplate:      "template",
				StartedAt:                  &t3,
				TerminalAt:                 &t4,
				Status:                     utils.PointerTo("status"),
				RelatedResources:           nil,
			},
		},
		{
			name: "ignores ci identifiers",
			fields: fields{
				commonFields: CommonFields{
					ID:        1,
					CreatedAt: t1,
					UpdatedAt: t2,
				},
				ciRunFields: ciRunFields{
					Platform:                   "platform",
					GithubActionsOwner:         "owner",
					GithubActionsRepo:          "repo",
					GithubActionsRunID:         2,
					GithubActionsAttemptNumber: 3,
					GithubActionsWorkflowPath:  "path",
					ArgoWorkflowsNamespace:     "namespace",
					ArgoWorkflowsName:          "name",
					ArgoWorkflowsTemplate:      "template",
					StartedAt:                  &t3,
					TerminalAt:                 &t4,
					Status:                     utils.PointerTo("status"),
				},
				RelatedResources: []CiIdentifierV3{
					{
						CommonFields: CommonFields{
							ID: 4,
						},
					},
				},
			},
			want: models.CiRun{
				Model: gorm.Model{
					ID:        1,
					CreatedAt: t1,
					UpdatedAt: t2,
				},
				Platform:                   "platform",
				GithubActionsOwner:         "owner",
				GithubActionsRepo:          "repo",
				GithubActionsRunID:         2,
				GithubActionsAttemptNumber: 3,
				GithubActionsWorkflowPath:  "path",
				ArgoWorkflowsNamespace:     "namespace",
				ArgoWorkflowsName:          "name",
				ArgoWorkflowsTemplate:      "template",
				StartedAt:                  &t3,
				TerminalAt:                 &t4,
				Status:                     utils.PointerTo("status"),
				RelatedResources:           nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CiRunV3{
				CommonFields:     tt.fields.commonFields,
				ciRunFields:      tt.fields.ciRunFields,
				RelatedResources: tt.fields.RelatedResources,
			}
			assert.Equalf(t, tt.want, c.toModel(), "toModel()")
		})
	}
}

func Test_ciRunFromModel(t *testing.T) {
	t1 := time.Now().Add(-(time.Minute * 4))
	t2 := time.Now().Add(-(time.Minute * 3))
	t3 := time.Now().Add(-(time.Minute * 2))
	t4 := time.Now().Add(-time.Minute)
	type args struct {
		model models.CiRun
	}
	tests := []struct {
		name string
		args args
		want CiRunV3
	}{
		{
			name: "equal",
			args: args{model: models.CiRun{
				Model: gorm.Model{
					ID:        1,
					CreatedAt: t1,
					UpdatedAt: t2,
					DeletedAt: gorm.DeletedAt{
						Valid: true,
						Time:  time.Now(),
					},
				},
				Platform:                   "platform",
				GithubActionsOwner:         "owner",
				GithubActionsRepo:          "repo",
				GithubActionsRunID:         2,
				GithubActionsAttemptNumber: 3,
				GithubActionsWorkflowPath:  "path",
				ArgoWorkflowsNamespace:     "namespace",
				ArgoWorkflowsName:          "name",
				ArgoWorkflowsTemplate:      "template",
				StartedAt:                  &t3,
				TerminalAt:                 &t4,
				Status:                     utils.PointerTo("status"),
				RelatedResources:           nil,
			}},
			want: CiRunV3{
				CommonFields: CommonFields{
					ID:        1,
					CreatedAt: t1,
					UpdatedAt: t2,
				},
				ciRunFields: ciRunFields{
					Platform:                   "platform",
					GithubActionsOwner:         "owner",
					GithubActionsRepo:          "repo",
					GithubActionsRunID:         2,
					GithubActionsAttemptNumber: 3,
					GithubActionsWorkflowPath:  "path",
					ArgoWorkflowsNamespace:     "namespace",
					ArgoWorkflowsName:          "name",
					ArgoWorkflowsTemplate:      "template",
					StartedAt:                  &t3,
					TerminalAt:                 &t4,
					Status:                     utils.PointerTo("status"),
				},
				RelatedResources: nil,
			},
		},
		{
			name: "respects related resources",
			args: args{model: models.CiRun{
				Model: gorm.Model{
					ID:        1,
					CreatedAt: t1,
					UpdatedAt: t2,
				},
				Platform:                   "platform",
				GithubActionsOwner:         "owner",
				GithubActionsRepo:          "repo",
				GithubActionsRunID:         2,
				GithubActionsAttemptNumber: 3,
				GithubActionsWorkflowPath:  "path",
				ArgoWorkflowsNamespace:     "namespace",
				ArgoWorkflowsName:          "name",
				ArgoWorkflowsTemplate:      "template",
				StartedAt:                  &t3,
				TerminalAt:                 &t4,
				Status:                     utils.PointerTo("status"),
				RelatedResources: []models.CiIdentifier{
					{
						Model: gorm.Model{
							ID: 4,
						},
					},
				},
			}},
			want: CiRunV3{
				CommonFields: CommonFields{
					ID:        1,
					CreatedAt: t1,
					UpdatedAt: t2,
				},
				ciRunFields: ciRunFields{
					Platform:                   "platform",
					GithubActionsOwner:         "owner",
					GithubActionsRepo:          "repo",
					GithubActionsRunID:         2,
					GithubActionsAttemptNumber: 3,
					GithubActionsWorkflowPath:  "path",
					ArgoWorkflowsNamespace:     "namespace",
					ArgoWorkflowsName:          "name",
					ArgoWorkflowsTemplate:      "template",
					StartedAt:                  &t3,
					TerminalAt:                 &t4,
					Status:                     utils.PointerTo("status"),
				},
				RelatedResources: []CiIdentifierV3{
					{
						CommonFields: CommonFields{
							ID: 4,
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, ciRunFromModel(tt.args.model), "ciRunFromModel(%v)", tt.args.model)
		})
	}
}
