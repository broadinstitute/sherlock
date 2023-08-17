package deployhooks

import (
	"github.com/broadinstitute/sherlock/go-shared/pkg/testutils"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/broadinstitute/sherlock/sherlock/internal/slack"
	"gorm.io/gorm"
	"time"
)

func (s *deployHooksSuite) Test_generateSlackAttachment_triggerNotLoaded() {
	_, err := generateSlackAttachment(nil, models.SlackDeployHook{
		Model:        gorm.Model{ID: 123},
		Trigger:      models.DeployHookTriggerConfig{},
		SlackChannel: testutils.PointerTo("channel"),
	}, models.CiRun{})
	s.ErrorContains(err, "SlackDeployHook 123 didn't have Trigger fully loaded")
}

func (s *deployHooksSuite) Test_generateSlackAttachment_statusNotPresent() {
	_, err := generateSlackAttachment(nil, models.SlackDeployHook{
		Trigger: models.DeployHookTriggerConfig{
			OnEnvironment: &models.Environment{Name: "name"},
		},
	}, models.CiRun{
		Model: gorm.Model{ID: 123},
	})
	s.ErrorContains(err, "CiRun 123 didn't have status present")
}

func (s *deployHooksSuite) Test_generateSlackAttachment_environmentNoChangesets_success() {
	result, err := generateSlackAttachment(nil, models.SlackDeployHook{
		Trigger: models.DeployHookTriggerConfig{
			OnEnvironment: &models.Environment{Name: "dev"},
		},
	}, models.CiRun{
		Model:                      gorm.Model{ID: 123},
		TerminalAt:                 testutils.PointerTo(time.Now()),
		Status:                     testutils.PointerTo("success"),
		Platform:                   "github-actions",
		GithubActionsOwner:         "broadinstitute",
		GithubActionsRepo:          "terra-github-workflows",
		GithubActionsRunID:         123123,
		GithubActionsAttemptNumber: 1,
	})
	s.NoError(err)
	s.Equal(slack.GreenBlock{Text: "Deployment to dev: <https://github.com/broadinstitute/terra-github-workflows/actions/runs/123123/attempts/1|success>."}, result)
}

func (s *deployHooksSuite) Test_generateSlackAttachment_environmentNoChangesets_failure() {
	result, err := generateSlackAttachment(nil, models.SlackDeployHook{
		Trigger: models.DeployHookTriggerConfig{
			OnEnvironment: &models.Environment{Name: "dev"},
		},
	}, models.CiRun{
		Model:                      gorm.Model{ID: 123},
		TerminalAt:                 testutils.PointerTo(time.Now()),
		Status:                     testutils.PointerTo("failure"),
		Platform:                   "github-actions",
		GithubActionsOwner:         "broadinstitute",
		GithubActionsRepo:          "terra-github-workflows",
		GithubActionsRunID:         123123,
		GithubActionsAttemptNumber: 1,
	})
	s.NoError(err)
	s.Equal(slack.RedBlock{Text: "Deployment to dev: <https://github.com/broadinstitute/terra-github-workflows/actions/runs/123123/attempts/1|failure>."}, result)
}

func (s *deployHooksSuite) Test_generateSlackAttachment_chartReleaseNoChangesets_success() {
	result, err := generateSlackAttachment(nil, models.SlackDeployHook{
		Trigger: models.DeployHookTriggerConfig{
			OnChartRelease: &models.ChartRelease{Name: "leonardo-dev"},
		},
	}, models.CiRun{
		Model:                      gorm.Model{ID: 123},
		TerminalAt:                 testutils.PointerTo(time.Now()),
		Status:                     testutils.PointerTo("success"),
		Platform:                   "github-actions",
		GithubActionsOwner:         "broadinstitute",
		GithubActionsRepo:          "terra-github-workflows",
		GithubActionsRunID:         123123,
		GithubActionsAttemptNumber: 1,
	})
	s.NoError(err)
	s.Equal(slack.GreenBlock{Text: "Deployment to leonardo-dev: <https://github.com/broadinstitute/terra-github-workflows/actions/runs/123123/attempts/1|success>."}, result)
}

func (s *deployHooksSuite) Test_generateSlackAttachment_chartReleaseNoChangesets_failure() {
	result, err := generateSlackAttachment(nil, models.SlackDeployHook{
		Trigger: models.DeployHookTriggerConfig{
			OnChartRelease: &models.ChartRelease{Name: "leonardo-dev"},
		},
	}, models.CiRun{
		Model:                      gorm.Model{ID: 123},
		TerminalAt:                 testutils.PointerTo(time.Now()),
		Status:                     testutils.PointerTo("failure"),
		Platform:                   "github-actions",
		GithubActionsOwner:         "broadinstitute",
		GithubActionsRepo:          "terra-github-workflows",
		GithubActionsRunID:         123123,
		GithubActionsAttemptNumber: 1,
	})
	s.NoError(err)
	s.Equal(slack.RedBlock{Text: "Deployment to leonardo-dev: <https://github.com/broadinstitute/terra-github-workflows/actions/runs/123123/attempts/1|failure>."}, result)
}

func (s *deployHooksSuite) Test_generateSlackAttachment_environmentChangesets_success() {
	result, err := generateSlackAttachment(nil, models.SlackDeployHook{
		Trigger: models.DeployHookTriggerConfig{
			OnEnvironment: &models.Environment{Name: "dev"},
		},
	}, models.CiRun{
		Model:                      gorm.Model{ID: 123},
		TerminalAt:                 testutils.PointerTo(time.Now()),
		Status:                     testutils.PointerTo("success"),
		Platform:                   "github-actions",
		GithubActionsOwner:         "broadinstitute",
		GithubActionsRepo:          "terra-github-workflows",
		GithubActionsRunID:         123123,
		GithubActionsAttemptNumber: 1,
		RelatedResources: []models.CiIdentifier{
			{
				ResourceType: "changeset",
				ResourceID:   1122,
			},
			{
				ResourceType: "chart-release",
				ResourceID:   1123,
			},
			{
				ResourceType: "changeset",
				ResourceID:   1124,
			},
		},
	})
	s.NoError(err)
	s.Equal(slack.GreenBlock{Text: "Deployment to dev: <https://github.com/broadinstitute/terra-github-workflows/actions/runs/123123/attempts/1|success>. Review all changes made by this deployment <https://beehive.dsp-devops.broadinstitute.org/review-changesets?changeset=1122&changeset=1124|here>."}, result)
}

func (s *deployHooksSuite) Test_generateSlackAttachment_environmentChangesets_failure() {
	result, err := generateSlackAttachment(nil, models.SlackDeployHook{
		Trigger: models.DeployHookTriggerConfig{
			OnEnvironment: &models.Environment{Name: "dev"},
		},
	}, models.CiRun{
		Model:                      gorm.Model{ID: 123},
		TerminalAt:                 testutils.PointerTo(time.Now()),
		Status:                     testutils.PointerTo("failure"),
		Platform:                   "github-actions",
		GithubActionsOwner:         "broadinstitute",
		GithubActionsRepo:          "terra-github-workflows",
		GithubActionsRunID:         123123,
		GithubActionsAttemptNumber: 1,
		RelatedResources: []models.CiIdentifier{
			{
				ResourceType: "changeset",
				ResourceID:   1122,
			},
			{
				ResourceType: "chart-release",
				ResourceID:   1123,
			},
			{
				ResourceType: "changeset",
				ResourceID:   1124,
			},
		},
	})
	s.NoError(err)
	s.Equal(slack.RedBlock{Text: "Deployment to dev: <https://github.com/broadinstitute/terra-github-workflows/actions/runs/123123/attempts/1|failure>. Review all changes made by this deployment <https://beehive.dsp-devops.broadinstitute.org/review-changesets?changeset=1122&changeset=1124|here>."}, result)
}

func (s *deployHooksSuite) Test_generateSlackAttachment_chartReleaseChangesets_success() {
	result, err := generateSlackAttachment(nil, models.SlackDeployHook{
		Trigger: models.DeployHookTriggerConfig{
			OnChartRelease: &models.ChartRelease{Name: "leonardo-dev"},
		},
	}, models.CiRun{
		Model:                      gorm.Model{ID: 123},
		TerminalAt:                 testutils.PointerTo(time.Now()),
		Status:                     testutils.PointerTo("success"),
		Platform:                   "github-actions",
		GithubActionsOwner:         "broadinstitute",
		GithubActionsRepo:          "terra-github-workflows",
		GithubActionsRunID:         123123,
		GithubActionsAttemptNumber: 1,
		RelatedResources: []models.CiIdentifier{
			{
				ResourceType: "changeset",
				ResourceID:   1122,
			},
			{
				ResourceType: "chart-release",
				ResourceID:   1123,
			},
			{
				ResourceType: "changeset",
				ResourceID:   1124,
			},
		},
	})
	s.NoError(err)
	s.Equal(slack.GreenBlock{Text: "Deployment to leonardo-dev: <https://github.com/broadinstitute/terra-github-workflows/actions/runs/123123/attempts/1|success>. Review all changes made by this deployment <https://beehive.dsp-devops.broadinstitute.org/review-changesets?changeset=1122&changeset=1124|here>."}, result)
}

func (s *deployHooksSuite) Test_generateSlackAttachment_chartReleaseChangesets_failure() {
	result, err := generateSlackAttachment(nil, models.SlackDeployHook{
		Trigger: models.DeployHookTriggerConfig{
			OnChartRelease: &models.ChartRelease{Name: "leonardo-dev"},
		},
	}, models.CiRun{
		Model:                      gorm.Model{ID: 123},
		TerminalAt:                 testutils.PointerTo(time.Now()),
		Status:                     testutils.PointerTo("failure"),
		Platform:                   "github-actions",
		GithubActionsOwner:         "broadinstitute",
		GithubActionsRepo:          "terra-github-workflows",
		GithubActionsRunID:         123123,
		GithubActionsAttemptNumber: 1,
		RelatedResources: []models.CiIdentifier{
			{
				ResourceType: "changeset",
				ResourceID:   1122,
			},
			{
				ResourceType: "chart-release",
				ResourceID:   1123,
			},
			{
				ResourceType: "changeset",
				ResourceID:   1124,
			},
		},
	})
	s.NoError(err)
	s.Equal(slack.RedBlock{Text: "Deployment to leonardo-dev: <https://github.com/broadinstitute/terra-github-workflows/actions/runs/123123/attempts/1|failure>. Review all changes made by this deployment <https://beehive.dsp-devops.broadinstitute.org/review-changesets?changeset=1122&changeset=1124|here>."}, result)
}
