package models

import (
	"github.com/broadinstitute/sherlock/go-shared/pkg/testutils"
	"gorm.io/gorm"
	"testing"
	"time"
)

func (s *modelSuite) TestCiRunPlatformValidationSqlPlatformInvalid() {
	err := s.DB.Create(&CiRun{
		Platform: "invalid",
	}).Error
	s.ErrorContains(err, "violates check constraint \"platform_present\"")
}

func (s *modelSuite) TestCiRunPlatformValidationSqlGithubInvalid() {
	err := s.DB.Create(&CiRun{
		Platform:                   "github-actions",
		GithubActionsOwner:         "owner",
		GithubActionsRepo:          "repo",
		GithubActionsRunID:         123,
		GithubActionsAttemptNumber: 123,
	}).Error
	s.ErrorContains(err, "violates check constraint \"platform_present\"")
}

func (s *modelSuite) TestCiRunPlatformValidationSqlGithubValid() {
	err := s.DB.Create(&CiRun{
		Platform:                   "github-actions",
		GithubActionsOwner:         "owner",
		GithubActionsRepo:          "repo",
		GithubActionsRunID:         123,
		GithubActionsAttemptNumber: 123,
		GithubActionsWorkflowPath:  "path",
	}).Error
	s.NoError(err)
}

func (s *modelSuite) TestCiRunPlatformValidationSqlArgoInvalid() {
	err := s.DB.Create(&CiRun{
		Platform:               "argo-workflows",
		ArgoWorkflowsNamespace: "namespace",
		ArgoWorkflowsName:      "name",
	}).Error
	s.ErrorContains(err, "violates check constraint \"platform_present\"")
}

func (s *modelSuite) TestCiRunPlatformValidationSqlArgoValid() {
	err := s.DB.Create(&CiRun{
		Platform:               "argo-workflows",
		ArgoWorkflowsNamespace: "namespace",
		ArgoWorkflowsName:      "name",
		ArgoWorkflowsTemplate:  "template",
	}).Error
	s.NoError(err)
}

func (s *modelSuite) TestCiRunPlatformValidationSqlBoth() {
	err := s.DB.Create(&CiRun{
		Platform:                   "argo-workflows",
		GithubActionsOwner:         "owner",
		GithubActionsRepo:          "repo",
		GithubActionsRunID:         123,
		GithubActionsAttemptNumber: 123,
		GithubActionsWorkflowPath:  "path",
		ArgoWorkflowsNamespace:     "namespace",
		ArgoWorkflowsName:          "name",
		ArgoWorkflowsTemplate:      "template",
	}).Error
	s.ErrorContains(err, "violates check constraint \"platform_present\"")
}

func (s *modelSuite) TestCiRunTerminalValidationInvalid() {
	run := CiRun{
		Platform:                   "github-actions",
		GithubActionsOwner:         "owner",
		GithubActionsRepo:          "repo",
		GithubActionsRunID:         123,
		GithubActionsAttemptNumber: 123,
		GithubActionsWorkflowPath:  "path",
		TerminalAt:                 testutils.PointerTo(time.Now()),
	}
	err := s.DB.Create(&run).Error
	s.ErrorContains(err, "violates check constraint \"terminal_status_present\"")
}

func (s *modelSuite) TestCiRunTerminalValidationValid() {
	err := s.DB.Create(&CiRun{
		Platform:                   "github-actions",
		GithubActionsOwner:         "owner",
		GithubActionsRepo:          "repo",
		GithubActionsRunID:         123,
		GithubActionsAttemptNumber: 123,
		GithubActionsWorkflowPath:  "path",
		TerminalAt:                 testutils.PointerTo(time.Now()),
		Status:                     testutils.PointerTo("status"),
	}).Error
	s.NoError(err)
}

func (s *modelSuite) TestCiRunUniquenessSqlInvalid() {
	err := s.DB.Create(&CiRun{
		Platform:                   "github-actions",
		GithubActionsOwner:         "owner",
		GithubActionsRepo:          "repo",
		GithubActionsRunID:         123,
		GithubActionsAttemptNumber: 123,
		GithubActionsWorkflowPath:  "path",
	}).Error
	s.NoError(err)
	err = s.DB.Create(&CiRun{
		Platform:                   "github-actions",
		GithubActionsOwner:         "owner",
		GithubActionsRepo:          "repo",
		GithubActionsRunID:         123,
		GithubActionsAttemptNumber: 123,
		GithubActionsWorkflowPath:  "path",
	}).Error
	s.ErrorContains(err, "violates unique constraint")
}

func (s *modelSuite) TestCiRunUniquenessSqlValid() {
	run1 := CiRun{
		Platform:                   "github-actions",
		GithubActionsOwner:         "owner",
		GithubActionsRepo:          "repo",
		GithubActionsRunID:         123,
		GithubActionsAttemptNumber: 123,
		GithubActionsWorkflowPath:  "path",
	}
	err := s.DB.Create(&run1).Error
	s.NoError(err)
	s.NotZero(run1.ID)
	err = s.DB.Delete(&run1).Error
	s.NoError(err)
	run2 := CiRun{
		Platform:                   "github-actions",
		GithubActionsOwner:         "owner",
		GithubActionsRepo:          "repo",
		GithubActionsRunID:         123,
		GithubActionsAttemptNumber: 123,
		GithubActionsWorkflowPath:  "path",
	}
	err = s.DB.Create(&run2).Error
	s.NoError(err)
	s.NotZero(run2.ID)
	s.NotEqual(run1.ID, run2.ID)
}

func TestCiRun_WebURL(t *testing.T) {
	type fields struct {
		Model                      gorm.Model
		Platform                   string
		GithubActionsOwner         string
		GithubActionsRepo          string
		GithubActionsRunID         uint
		GithubActionsAttemptNumber uint
		GithubActionsWorkflowPath  string
		ArgoWorkflowsNamespace     string
		ArgoWorkflowsName          string
		ArgoWorkflowsTemplate      string
		DeployHooksDispatchedAt    *string
		RelatedResources           []CiIdentifier
		StartedAt                  *time.Time
		TerminalAt                 *time.Time
		Status                     *string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "github",
			fields: fields{
				Platform:                   "github-actions",
				GithubActionsOwner:         "owner",
				GithubActionsRepo:          "repo",
				GithubActionsRunID:         123,
				GithubActionsAttemptNumber: 1,
			},
			want: "https://github.com/owner/repo/actions/runs/123/attempts/1",
		},
		{
			name: "argo",
			fields: fields{
				Platform:               "argo-workflows",
				ArgoWorkflowsNamespace: "namespace",
				ArgoWorkflowsName:      "name",
			},
			want: "https://workflows.dsp-devops.broadinstitute.org/workflows/namespace/name",
		},
		{
			name: "invalid? non-enum",
			fields: fields{
				Model:    gorm.Model{ID: 123},
				Platform: "something invalid",
			},
			want: "https://sherlock.dsp-devops.broadinstitute.org/api/ci-runs/v3/123",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CiRun{
				Model:                        tt.fields.Model,
				Platform:                     tt.fields.Platform,
				GithubActionsOwner:           tt.fields.GithubActionsOwner,
				GithubActionsRepo:            tt.fields.GithubActionsRepo,
				GithubActionsRunID:           tt.fields.GithubActionsRunID,
				GithubActionsAttemptNumber:   tt.fields.GithubActionsAttemptNumber,
				GithubActionsWorkflowPath:    tt.fields.GithubActionsWorkflowPath,
				ArgoWorkflowsNamespace:       tt.fields.ArgoWorkflowsNamespace,
				ArgoWorkflowsName:            tt.fields.ArgoWorkflowsName,
				ArgoWorkflowsTemplate:        tt.fields.ArgoWorkflowsTemplate,
				TerminationHooksDispatchedAt: tt.fields.DeployHooksDispatchedAt,
				RelatedResources:             tt.fields.RelatedResources,
				StartedAt:                    tt.fields.StartedAt,
				TerminalAt:                   tt.fields.TerminalAt,
				Status:                       tt.fields.Status,
			}
			if got := c.WebURL(); got != tt.want {
				t.Errorf("WebURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCiRun_Succeeded(t *testing.T) {
	type fields struct {
		Model                      gorm.Model
		Platform                   string
		GithubActionsOwner         string
		GithubActionsRepo          string
		GithubActionsRunID         uint
		GithubActionsAttemptNumber uint
		GithubActionsWorkflowPath  string
		ArgoWorkflowsNamespace     string
		ArgoWorkflowsName          string
		ArgoWorkflowsTemplate      string
		DeployHooksDispatchedAt    *string
		RelatedResources           []CiIdentifier
		StartedAt                  *time.Time
		TerminalAt                 *time.Time
		Status                     *string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "success",
			fields: fields{
				TerminalAt: testutils.PointerTo(time.Now()),
				Status:     testutils.PointerTo("success"),
			},
			want: true,
		},
		{
			name: "not success",
			fields: fields{
				TerminalAt: testutils.PointerTo(time.Now()),
				Status:     testutils.PointerTo("something else"),
			},
			want: false,
		},
		{
			name: "no status",
			fields: fields{
				TerminalAt: testutils.PointerTo(time.Now()),
			},
			want: false,
		},
		{
			name: "not terminal",
			fields: fields{
				Status: testutils.PointerTo("success"),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CiRun{
				Model:                        tt.fields.Model,
				Platform:                     tt.fields.Platform,
				GithubActionsOwner:           tt.fields.GithubActionsOwner,
				GithubActionsRepo:            tt.fields.GithubActionsRepo,
				GithubActionsRunID:           tt.fields.GithubActionsRunID,
				GithubActionsAttemptNumber:   tt.fields.GithubActionsAttemptNumber,
				GithubActionsWorkflowPath:    tt.fields.GithubActionsWorkflowPath,
				ArgoWorkflowsNamespace:       tt.fields.ArgoWorkflowsNamespace,
				ArgoWorkflowsName:            tt.fields.ArgoWorkflowsName,
				ArgoWorkflowsTemplate:        tt.fields.ArgoWorkflowsTemplate,
				TerminationHooksDispatchedAt: tt.fields.DeployHooksDispatchedAt,
				RelatedResources:             tt.fields.RelatedResources,
				StartedAt:                    tt.fields.StartedAt,
				TerminalAt:                   tt.fields.TerminalAt,
				Status:                       tt.fields.Status,
			}
			if got := c.Succeeded(); got != tt.want {
				t.Errorf("Succeeded() = %v, want %v", got, tt.want)
			}
		})
	}
}
