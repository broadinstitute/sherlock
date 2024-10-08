package models

import (
	"context"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/clients/github"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	github2 "github.com/google/go-github/v58/github"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"net/http"
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
		TerminalAt:                 utils.PointerTo(time.Now()),
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
		TerminalAt:                 utils.PointerTo(time.Now()),
		Status:                     utils.PointerTo("status"),
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

func (s *modelSuite) TestCiRun_FillRelatedResourceStatuses() {
	ciRun := s.TestData.CiRun_Deploy_LeonardoDev_V1toV3()
	resourcesWithStatusSet := 0
	for _, rr := range ciRun.RelatedResources {
		if rr.ResourceStatus != nil {
			resourcesWithStatusSet++
		}
	}
	s.Zero(resourcesWithStatusSet)
	s.NoError(ciRun.FillRelatedResourceStatuses(s.DB))
	resourcesWithStatusSet = 0
	for _, rr := range ciRun.RelatedResources {
		if rr.ResourceStatus != nil {
			resourcesWithStatusSet++
		}
	}
	s.NotZero(resourcesWithStatusSet)
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
			want: "https://sherlock.dsp-devops-prod.broadinstitute.org/api/ci-runs/v3/123",
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
				TerminalAt: utils.PointerTo(time.Now()),
				Status:     utils.PointerTo("success"),
			},
			want: true,
		},
		{
			name: "not success",
			fields: fields{
				TerminalAt: utils.PointerTo(time.Now()),
				Status:     utils.PointerTo("something else"),
			},
			want: false,
		},
		{
			name: "no status",
			fields: fields{
				TerminalAt: utils.PointerTo(time.Now()),
			},
			want: false,
		},
		{
			name: "not terminal",
			fields: fields{
				Status: utils.PointerTo("success"),
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

func TestCiRun_Nickname(t *testing.T) {
	type fields struct {
		Model                          gorm.Model
		Platform                       string
		GithubActionsOwner             string
		GithubActionsRepo              string
		GithubActionsRunID             uint
		GithubActionsAttemptNumber     uint
		GithubActionsWorkflowPath      string
		ArgoWorkflowsNamespace         string
		ArgoWorkflowsName              string
		ArgoWorkflowsTemplate          string
		TerminationHooksDispatchedAt   *string
		RelatedResources               []CiIdentifier
		StartedAt                      *time.Time
		TerminalAt                     *time.Time
		Status                         *string
		NotifySlackChannelsUponSuccess datatypes.JSONSlice[string]
		NotifySlackChannelsUponFailure datatypes.JSONSlice[string]
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "gha",
			fields: fields{
				Platform:                  "github-actions",
				GithubActionsRepo:         "repo",
				GithubActionsWorkflowPath: "path/to/file.yaml",
			},
			want: "repo's file workflow",
		},
		{
			name: "argo",
			fields: fields{
				Platform:              "argo-workflows",
				ArgoWorkflowsTemplate: "template",
			},
			want: "template Argo workflow",
		},
		{
			name: "unknown",
			fields: fields{
				Model:    gorm.Model{ID: 1},
				Platform: "platform",
			},
			want: "unknown platform workflow 1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CiRun{
				Model:                          tt.fields.Model,
				Platform:                       tt.fields.Platform,
				GithubActionsOwner:             tt.fields.GithubActionsOwner,
				GithubActionsRepo:              tt.fields.GithubActionsRepo,
				GithubActionsRunID:             tt.fields.GithubActionsRunID,
				GithubActionsAttemptNumber:     tt.fields.GithubActionsAttemptNumber,
				GithubActionsWorkflowPath:      tt.fields.GithubActionsWorkflowPath,
				ArgoWorkflowsNamespace:         tt.fields.ArgoWorkflowsNamespace,
				ArgoWorkflowsName:              tt.fields.ArgoWorkflowsName,
				ArgoWorkflowsTemplate:          tt.fields.ArgoWorkflowsTemplate,
				TerminationHooksDispatchedAt:   tt.fields.TerminationHooksDispatchedAt,
				RelatedResources:               tt.fields.RelatedResources,
				StartedAt:                      tt.fields.StartedAt,
				TerminalAt:                     tt.fields.TerminalAt,
				Status:                         tt.fields.Status,
				NotifySlackChannelsUponSuccess: tt.fields.NotifySlackChannelsUponSuccess,
				NotifySlackChannelsUponFailure: tt.fields.NotifySlackChannelsUponFailure,
			}
			if got := c.Nickname(); got != tt.want {
				t.Errorf("Nickname() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCiRun_IsDeploy(t *testing.T) {
	config.LoadTestConfig()
	assert.NoError(t, initDeployMatchers())
	type fields struct {
		Model                          gorm.Model
		Platform                       string
		GithubActionsOwner             string
		GithubActionsRepo              string
		GithubActionsRunID             uint
		GithubActionsAttemptNumber     uint
		GithubActionsWorkflowPath      string
		ArgoWorkflowsNamespace         string
		ArgoWorkflowsName              string
		ArgoWorkflowsTemplate          string
		TerminationHooksDispatchedAt   *string
		RelatedResources               []CiIdentifier
		StartedAt                      *time.Time
		TerminalAt                     *time.Time
		Status                         *string
		NotifySlackChannelsUponSuccess datatypes.JSONSlice[string]
		NotifySlackChannelsUponFailure datatypes.JSONSlice[string]
		ResourceStatus                 *string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "matches",
			fields: fields{
				Platform:                   "github-actions",
				GithubActionsOwner:         "broadinstitute",
				GithubActionsRepo:          "terra-github-workflows",
				GithubActionsRunID:         123123,
				GithubActionsAttemptNumber: 1,
				GithubActionsWorkflowPath:  ".github/workflows/sync-release.yaml",
			},
			want: true,
		},
		{
			name: "no platform match",
			fields: fields{
				Platform:               "argo-workflows",
				ArgoWorkflowsNamespace: "namespace",
				ArgoWorkflowsName:      "name",
				ArgoWorkflowsTemplate:  "template",
			},
			want: false,
		},
		{
			name: "no owner match",
			fields: fields{
				Platform:                   "github-actions",
				GithubActionsOwner:         "DataBiosphere",
				GithubActionsRepo:          "terra-github-workflows",
				GithubActionsRunID:         123123,
				GithubActionsAttemptNumber: 1,
				GithubActionsWorkflowPath:  ".github/workflows/sync-release.yaml",
			},
			want: false,
		},
		{
			name: "no repo match",
			fields: fields{
				Platform:                   "github-actions",
				GithubActionsOwner:         "broadinstitute",
				GithubActionsRepo:          "terra-helmfile",
				GithubActionsRunID:         123123,
				GithubActionsAttemptNumber: 1,
				GithubActionsWorkflowPath:  ".github/workflows/sync-release.yaml",
			},
			want: false,
		},
		{
			name: "no path match",
			fields: fields{
				Platform:                   "github-actions",
				GithubActionsOwner:         "broadinstitute",
				GithubActionsRepo:          "terra-github-workflows",
				GithubActionsRunID:         123123,
				GithubActionsAttemptNumber: 1,
				GithubActionsWorkflowPath:  ".github/workflows/bee-create.yaml",
			},
			want: false,
		},
		{
			name: "matches even if otherwise required fields are missing",
			fields: fields{
				Platform:                  "github-actions",
				GithubActionsOwner:        "broadinstitute",
				GithubActionsRepo:         "terra-github-workflows",
				GithubActionsWorkflowPath: ".github/workflows/sync-release.yaml",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CiRun{
				Model:                          tt.fields.Model,
				Platform:                       tt.fields.Platform,
				GithubActionsOwner:             tt.fields.GithubActionsOwner,
				GithubActionsRepo:              tt.fields.GithubActionsRepo,
				GithubActionsRunID:             tt.fields.GithubActionsRunID,
				GithubActionsAttemptNumber:     tt.fields.GithubActionsAttemptNumber,
				GithubActionsWorkflowPath:      tt.fields.GithubActionsWorkflowPath,
				ArgoWorkflowsNamespace:         tt.fields.ArgoWorkflowsNamespace,
				ArgoWorkflowsName:              tt.fields.ArgoWorkflowsName,
				ArgoWorkflowsTemplate:          tt.fields.ArgoWorkflowsTemplate,
				TerminationHooksDispatchedAt:   tt.fields.TerminationHooksDispatchedAt,
				RelatedResources:               tt.fields.RelatedResources,
				StartedAt:                      tt.fields.StartedAt,
				TerminalAt:                     tt.fields.TerminalAt,
				Status:                         tt.fields.Status,
				NotifySlackChannelsUponSuccess: tt.fields.NotifySlackChannelsUponSuccess,
				NotifySlackChannelsUponFailure: tt.fields.NotifySlackChannelsUponFailure,
				ResourceStatus:                 tt.fields.ResourceStatus,
			}
			if got := c.IsDeploy(); got != tt.want {
				t.Errorf("IsDeploy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func (s *modelSuite) TestCiRun_SlackCompletionText() {
	environment := s.TestData.Environment_Dev()
	chartRelease := s.TestData.ChartRelease_LeonardoDev()

	type fields struct {
		Model                          gorm.Model
		Platform                       string
		GithubActionsOwner             string
		GithubActionsRepo              string
		GithubActionsRunID             uint
		GithubActionsAttemptNumber     uint
		GithubActionsWorkflowPath      string
		ArgoWorkflowsNamespace         string
		ArgoWorkflowsName              string
		ArgoWorkflowsTemplate          string
		TerminationHooksDispatchedAt   *string
		RelatedResources               []CiIdentifier
		StartedAt                      *time.Time
		TerminalAt                     *time.Time
		Status                         *string
		NotifySlackChannelsUponSuccess datatypes.JSONSlice[string]
		NotifySlackChannelsUponFailure datatypes.JSONSlice[string]
		ResourceStatus                 *string
	}
	tests := []struct {
		name             string
		fields           fields
		githubMockConfig func(c *github.MockClient)
		want             string
	}{
		{
			name: "chart release",
			fields: fields{
				Platform:                   "github-actions",
				GithubActionsOwner:         "owner",
				GithubActionsRepo:          "repo",
				GithubActionsRunID:         1,
				GithubActionsAttemptNumber: 3,
				GithubActionsWorkflowPath:  "workflow",
				StartedAt:                  utils.PointerTo(time.Now().Add(-time.Minute)),
				TerminalAt:                 utils.PointerTo(time.Now()),
				Status:                     utils.PointerTo("success"),
				RelatedResources: []CiIdentifier{
					{ResourceType: "chart-release", ResourceID: chartRelease.ID},
					{ResourceType: "environment", ResourceID: environment.ID},
				},
			},
			want: "repo's workflow workflow against <https://beehive.dsp-devops-prod.broadinstitute.org/r/chart-release/leonardo-dev|leonardo-dev> (attempt 3): <https://github.com/owner/repo/actions/runs/1/attempts/3|success>",
		},
		{
			name: "environment",
			fields: fields{
				Platform:                   "github-actions",
				GithubActionsOwner:         "owner",
				GithubActionsRepo:          "repo",
				GithubActionsRunID:         1,
				GithubActionsAttemptNumber: 1,
				GithubActionsWorkflowPath:  "workflow",
				StartedAt:                  utils.PointerTo(time.Now().Add(-time.Minute)),
				TerminalAt:                 utils.PointerTo(time.Now()),
				Status:                     utils.PointerTo("success"),
				RelatedResources: []CiIdentifier{
					{ResourceType: "environment", ResourceID: environment.ID},
				},
			},
			want: "repo's workflow workflow against <https://beehive.dsp-devops-prod.broadinstitute.org/r/environment/dev|dev>: <https://github.com/owner/repo/actions/runs/1/attempts/1|success>",
		},
		{
			name: "with jobs",
			fields: fields{
				Platform:                   "github-actions",
				GithubActionsOwner:         "owner",
				GithubActionsRepo:          "repo",
				GithubActionsRunID:         1,
				GithubActionsAttemptNumber: 3,
				GithubActionsWorkflowPath:  "workflow",
				StartedAt:                  utils.PointerTo(time.Now().Add(-time.Minute)),
				TerminalAt:                 utils.PointerTo(time.Now()),
				Status:                     utils.PointerTo("success"),
				RelatedResources: []CiIdentifier{
					{ResourceType: "environment", ResourceID: environment.ID},
				},
			},
			githubMockConfig: func(c *github.MockClient) {
				request, err := http.NewRequest(http.MethodGet, "repos/owner/repo/actions/runs/1/attempts/3/jobs", nil)
				if err != nil {
					panic(err)
				}
				c.EXPECT().NewRequest(http.MethodGet, "repos/owner/repo/actions/runs/1/attempts/3/jobs", nil).Return(request, nil).Once()
				c.EXPECT().Do(mock.Anything, request, mock.AnythingOfType("*github.Jobs")).Run(
					func(ctx context.Context, req *http.Request, v interface{}) {
						jobs := v.(*github2.Jobs)
						jobs.Jobs = []*github2.WorkflowJob{
							{
								ID:         utils.PointerTo(int64(1)),
								Name:       utils.PointerTo("job1"),
								Conclusion: utils.PointerTo("success"),
							},
							{
								ID:     utils.PointerTo(int64(2)),
								Name:   utils.PointerTo("job2"),
								Status: utils.PointerTo("in_progress"),
							},
							{
								ID:         utils.PointerTo(int64(3)),
								Name:       utils.PointerTo("job3"),
								Conclusion: utils.PointerTo("failure"),
							},
						}
					},
				).Return(&github2.Response{
					Response: &http.Response{StatusCode: http.StatusOK},
				}, nil).Once()
			},
			want: "repo's workflow workflow against <https://beehive.dsp-devops-prod.broadinstitute.org/r/environment/dev|dev> (attempt 3): <https://github.com/owner/repo/actions/runs/1/attempts/3|success> (job3: <https://github.com/owner/repo/actions/runs/1/job/3|failure>)",
		},
	}
	for _, tt := range tests {
		s.Run(tt.name, func() {
			github.UseMockedClient(s.T(), tt.githubMockConfig, func() {
				c := &CiRun{
					Model:                          tt.fields.Model,
					Platform:                       tt.fields.Platform,
					GithubActionsOwner:             tt.fields.GithubActionsOwner,
					GithubActionsRepo:              tt.fields.GithubActionsRepo,
					GithubActionsRunID:             tt.fields.GithubActionsRunID,
					GithubActionsAttemptNumber:     tt.fields.GithubActionsAttemptNumber,
					GithubActionsWorkflowPath:      tt.fields.GithubActionsWorkflowPath,
					ArgoWorkflowsNamespace:         tt.fields.ArgoWorkflowsNamespace,
					ArgoWorkflowsName:              tt.fields.ArgoWorkflowsName,
					ArgoWorkflowsTemplate:          tt.fields.ArgoWorkflowsTemplate,
					TerminationHooksDispatchedAt:   tt.fields.TerminationHooksDispatchedAt,
					RelatedResources:               tt.fields.RelatedResources,
					StartedAt:                      tt.fields.StartedAt,
					TerminalAt:                     tt.fields.TerminalAt,
					Status:                         tt.fields.Status,
					NotifySlackChannelsUponSuccess: tt.fields.NotifySlackChannelsUponSuccess,
					NotifySlackChannelsUponFailure: tt.fields.NotifySlackChannelsUponFailure,
					ResourceStatus:                 tt.fields.ResourceStatus,
				}
				got, errs := c.SlackCompletionText(s.DB)
				s.Empty(errs)
				s.Equalf(tt.want, got, "SlackCompletionText()")
			})
		})
	}
}

func TestCiRun_DoneOrUnderway(t *testing.T) {
	type fields struct {
		TerminalAt *time.Time
		Status     *string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "not terminal",
			want: "underway",
		},
		{
			name: "success",
			fields: fields{
				TerminalAt: utils.PointerTo(time.Now()),
				Status:     utils.PointerTo("success"),
			},
			want: "done",
		},
		{
			name: "failure",
			fields: fields{
				TerminalAt: utils.PointerTo(time.Now()),
				Status:     utils.PointerTo("failure"),
			},
			want: "done",
		},
		{
			name: "some other finished status",
			fields: fields{
				TerminalAt: utils.PointerTo(time.Now()),
				Status:     utils.PointerTo("cancelled"),
			},
			want: "cancelled",
		},
		{
			name: "terminal but no status yet",
			fields: fields{
				TerminalAt: utils.PointerTo(time.Now()),
			},
			want: "waiting for status",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CiRun{
				TerminalAt: tt.fields.TerminalAt,
				Status:     tt.fields.Status,
			}
			assert.Equalf(t, tt.want, c.DoneOrUnderway(), "DoneOrUnderway()")
		})
	}
}
