package hooks

import (
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/broadinstitute/sherlock/sherlock/internal/slack"
	"github.com/broadinstitute/sherlock/sherlock/internal/slack/slack_mocks"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
	"time"
)

func (s *hooksSuite) Test_dispatchSlackDeployHook_channelNil() {
	s.ErrorContains(dispatcher.DispatchSlackDeployHook(nil, models.SlackDeployHook{}, models.CiRun{
		TerminalAt: utils.PointerTo(time.Now()),
	}), "slack channel was nil")
}

func (s *hooksSuite) Test_dispatchSlackDeployHook_generateError() {
	s.ErrorContains(dispatcher.DispatchSlackDeployHook(nil, models.SlackDeployHook{
		Model:        gorm.Model{ID: 123},
		Trigger:      models.DeployHookTriggerConfig{},
		SlackChannel: utils.PointerTo("channel"),
	}, models.CiRun{
		TerminalAt: utils.PointerTo(time.Now()),
	}), "SlackDeployHook 123 didn't have Trigger fully loaded")
}

func (s *hooksSuite) Test_dispatchSlackDeployHook() {
	slack.UseMockedClient(s.T(), func(c *slack_mocks.MockMockableClient) {
		c.EXPECT().
			SendMessageContext(s.DB.Statement.Context, "channel", mock.AnythingOfType("slack.MsgOption")).
			Return("", "", "", nil)
	}, func() {
		s.NoError(dispatcher.DispatchSlackDeployHook(s.DB, models.SlackDeployHook{
			Trigger: models.DeployHookTriggerConfig{
				OnEnvironment: &models.Environment{Name: "dev"},
			},
			SlackChannel: utils.PointerTo("channel"),
		}, models.CiRun{
			Model:                      gorm.Model{ID: 123},
			TerminalAt:                 utils.PointerTo(time.Now()),
			Status:                     utils.PointerTo("success"),
			Platform:                   "github-actions",
			GithubActionsOwner:         "broadinstitute",
			GithubActionsRepo:          "terra-github-workflows",
			GithubActionsRunID:         123123,
			GithubActionsAttemptNumber: 1,
		}))
	})
}

func (s *hooksSuite) Test_generateSlackAttachment_triggerNotLoaded() {
	// Pass nil for the db because it shouldn't be used in this call
	_, err := generateSlackAttachment(nil, models.SlackDeployHook{
		Model:        gorm.Model{ID: 123},
		Trigger:      models.DeployHookTriggerConfig{},
		SlackChannel: utils.PointerTo("channel"),
	}, models.CiRun{})
	s.ErrorContains(err, "SlackDeployHook 123 didn't have Trigger fully loaded")
}

func (s *hooksSuite) Test_generateSlackAttachment_statusNotPresent() {
	// Pass nil for the db because it shouldn't be used in this call
	_, err := generateSlackAttachment(nil, models.SlackDeployHook{
		Trigger: models.DeployHookTriggerConfig{
			OnEnvironment: &models.Environment{Name: "name"},
		},
	}, models.CiRun{
		Model: gorm.Model{ID: 123},
	})
	s.ErrorContains(err, "CiRun 123 didn't have status present")
}

func (s *hooksSuite) Test_generateSlackAttachment_environmentNoChangesets_success() {
	// Pass nil for the db because it shouldn't be used in this call
	result, err := generateSlackAttachment(nil, models.SlackDeployHook{
		Trigger: models.DeployHookTriggerConfig{
			OnEnvironment: &models.Environment{Name: "dev"},
		},
	}, models.CiRun{
		Model:                      gorm.Model{ID: 123},
		TerminalAt:                 utils.PointerTo(time.Now()),
		Status:                     utils.PointerTo("success"),
		Platform:                   "github-actions",
		GithubActionsOwner:         "broadinstitute",
		GithubActionsRepo:          "terra-github-workflows",
		GithubActionsRunID:         123123,
		GithubActionsAttemptNumber: 1,
	})
	s.NoError(err)
	s.Equal(slack.GreenBlock{Text: "Deployment to <https://beehive.dsp-devops.broadinstitute.org/r/environment/dev|dev>: <https://github.com/broadinstitute/terra-github-workflows/actions/runs/123123/attempts/1|success>."}, result)
}

func (s *hooksSuite) Test_generateSlackAttachment_environmentNoChangesets_failure() {
	// Pass nil for the db because it shouldn't be used in this call
	result, err := generateSlackAttachment(nil, models.SlackDeployHook{
		Trigger: models.DeployHookTriggerConfig{
			OnEnvironment: &models.Environment{Name: "dev"},
		},
	}, models.CiRun{
		Model:                      gorm.Model{ID: 123},
		TerminalAt:                 utils.PointerTo(time.Now()),
		Status:                     utils.PointerTo("failure"),
		Platform:                   "github-actions",
		GithubActionsOwner:         "broadinstitute",
		GithubActionsRepo:          "terra-github-workflows",
		GithubActionsRunID:         123123,
		GithubActionsAttemptNumber: 1,
	})
	s.NoError(err)
	s.Equal(slack.RedBlock{Text: "Deployment to <https://beehive.dsp-devops.broadinstitute.org/r/environment/dev|dev>: <https://github.com/broadinstitute/terra-github-workflows/actions/runs/123123/attempts/1|failure>."}, result)
}

func (s *hooksSuite) Test_generateSlackAttachment_chartReleaseNoChangesets_success() {
	// Pass nil for the db because it shouldn't be used in this call
	result, err := generateSlackAttachment(nil, models.SlackDeployHook{
		Trigger: models.DeployHookTriggerConfig{
			OnChartRelease: &models.ChartRelease{Name: "leonardo-dev"},
		},
	}, models.CiRun{
		Model:                      gorm.Model{ID: 123},
		TerminalAt:                 utils.PointerTo(time.Now()),
		Status:                     utils.PointerTo("success"),
		Platform:                   "github-actions",
		GithubActionsOwner:         "broadinstitute",
		GithubActionsRepo:          "terra-github-workflows",
		GithubActionsRunID:         123123,
		GithubActionsAttemptNumber: 1,
	})
	s.NoError(err)
	s.Equal(slack.GreenBlock{Text: "Deployment to <https://beehive.dsp-devops.broadinstitute.org/r/chart-release/leonardo-dev|leonardo-dev>: <https://github.com/broadinstitute/terra-github-workflows/actions/runs/123123/attempts/1|success>."}, result)
}

func (s *hooksSuite) Test_generateSlackAttachment_chartReleaseNoChangesets_failure() {
	// Pass nil for the db because it shouldn't be used in this call
	result, err := generateSlackAttachment(nil, models.SlackDeployHook{
		Trigger: models.DeployHookTriggerConfig{
			OnChartRelease: &models.ChartRelease{Name: "leonardo-dev"},
		},
	}, models.CiRun{
		Model:                      gorm.Model{ID: 123},
		TerminalAt:                 utils.PointerTo(time.Now()),
		Status:                     utils.PointerTo("failure"),
		Platform:                   "github-actions",
		GithubActionsOwner:         "broadinstitute",
		GithubActionsRepo:          "terra-github-workflows",
		GithubActionsRunID:         123123,
		GithubActionsAttemptNumber: 1,
	})
	s.NoError(err)
	s.Equal(slack.RedBlock{Text: "Deployment to <https://beehive.dsp-devops.broadinstitute.org/r/chart-release/leonardo-dev|leonardo-dev>: <https://github.com/broadinstitute/terra-github-workflows/actions/runs/123123/attempts/1|failure>."}, result)
}

func (s *hooksSuite) Test_generateSlackAttachment_environmentChangesets_success() {
	// Pass nil for the db because it shouldn't be used in this call
	result, err := generateSlackAttachment(nil, models.SlackDeployHook{
		Trigger: models.DeployHookTriggerConfig{
			OnEnvironment: &models.Environment{Name: "dev"},
		},
	}, models.CiRun{
		Model:                      gorm.Model{ID: 123},
		TerminalAt:                 utils.PointerTo(time.Now()),
		Status:                     utils.PointerTo("success"),
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
				ResourceType: "environment",
				ResourceID:   1123,
			},
			{
				ResourceType: "changeset",
				ResourceID:   1124,
			},
		},
	})
	s.NoError(err)
	s.Equal(slack.GreenBlock{Text: "Deployment to <https://beehive.dsp-devops.broadinstitute.org/r/environment/dev|dev>: <https://github.com/broadinstitute/terra-github-workflows/actions/runs/123123/attempts/1|success>. Review all changes made by this deployment <https://beehive.dsp-devops.broadinstitute.org/review-changesets?changeset=1122&changeset=1124|here>."}, result)
}

func (s *hooksSuite) Test_generateSlackAttachment_environmentChangesets_failure() {
	// Pass nil for the db because it shouldn't be used in this call
	result, err := generateSlackAttachment(nil, models.SlackDeployHook{
		Trigger: models.DeployHookTriggerConfig{
			OnEnvironment: &models.Environment{Name: "dev"},
		},
	}, models.CiRun{
		Model:                      gorm.Model{ID: 123},
		TerminalAt:                 utils.PointerTo(time.Now()),
		Status:                     utils.PointerTo("failure"),
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
				ResourceType: "environment",
				ResourceID:   1123,
			},
			{
				ResourceType: "changeset",
				ResourceID:   1124,
			},
		},
	})
	s.NoError(err)
	s.Equal(slack.RedBlock{Text: "Deployment to <https://beehive.dsp-devops.broadinstitute.org/r/environment/dev|dev>: <https://github.com/broadinstitute/terra-github-workflows/actions/runs/123123/attempts/1|failure>. Review all changes made by this deployment <https://beehive.dsp-devops.broadinstitute.org/review-changesets?changeset=1122&changeset=1124|here>."}, result)
}

func (s *hooksSuite) Test_generateSlackAttachment_chartReleaseChangesets_success() {
	// Pass nil for the db because it shouldn't be used in this call
	result, err := generateSlackAttachment(nil, models.SlackDeployHook{
		Trigger: models.DeployHookTriggerConfig{
			OnChartRelease: &models.ChartRelease{Name: "leonardo-dev"},
		},
	}, models.CiRun{
		Model:                      gorm.Model{ID: 123},
		TerminalAt:                 utils.PointerTo(time.Now()),
		Status:                     utils.PointerTo("success"),
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
				ResourceType: "environment",
				ResourceID:   1123,
			},
			{
				ResourceType: "changeset",
				ResourceID:   1124,
			},
		},
	})
	s.NoError(err)
	s.Equal(slack.GreenBlock{Text: "Deployment to <https://beehive.dsp-devops.broadinstitute.org/r/chart-release/leonardo-dev|leonardo-dev>: <https://github.com/broadinstitute/terra-github-workflows/actions/runs/123123/attempts/1|success>. Review all changes made by this deployment <https://beehive.dsp-devops.broadinstitute.org/review-changesets?changeset=1122&changeset=1124|here>."}, result)
}

func (s *hooksSuite) Test_generateSlackAttachment_chartReleaseChangesets_failure() {
	// Pass nil for the db because it shouldn't be used in this call
	result, err := generateSlackAttachment(nil, models.SlackDeployHook{
		Trigger: models.DeployHookTriggerConfig{
			OnChartRelease: &models.ChartRelease{Name: "leonardo-dev"},
		},
	}, models.CiRun{
		Model:                      gorm.Model{ID: 123},
		TerminalAt:                 utils.PointerTo(time.Now()),
		Status:                     utils.PointerTo("failure"),
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
				ResourceType: "environment",
				ResourceID:   1123,
			},
			{
				ResourceType: "changeset",
				ResourceID:   1124,
			},
		},
	})
	s.NoError(err)
	s.Equal(slack.RedBlock{Text: "Deployment to <https://beehive.dsp-devops.broadinstitute.org/r/chart-release/leonardo-dev|leonardo-dev>: <https://github.com/broadinstitute/terra-github-workflows/actions/runs/123123/attempts/1|failure>. Review all changes made by this deployment <https://beehive.dsp-devops.broadinstitute.org/review-changesets?changeset=1122&changeset=1124|here>."}, result)
}

func (s *hooksSuite) Test_generateSlackAttachment_environmentChangesets_chartReleaseNotFound() {
	_, err := generateSlackAttachment(s.DB, models.SlackDeployHook{
		Trigger: models.DeployHookTriggerConfig{
			OnEnvironment: &models.Environment{Name: "dev"},
		},
	}, models.CiRun{
		Model:                      gorm.Model{ID: 123},
		TerminalAt:                 utils.PointerTo(time.Now()),
		Status:                     utils.PointerTo("failure"),
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
		},
	})
	s.ErrorIs(err, gorm.ErrRecordNotFound)
}

func (s *hooksSuite) Test_generateSlackAttachment_environmentChangesets_chartReleaseNames() {
	environment := s.TestData.Environment_Dev()
	environmentCiIdentifier := environment.GetCiIdentifier()
	s.NoError(s.DB.Create(&environmentCiIdentifier).Error)

	environmentGithubHook := models.GithubActionsDeployHook{
		Trigger: models.DeployHookTriggerConfig{
			OnEnvironmentID: &environment.ID,
			OnSuccess:       utils.PointerTo(true),
			OnFailure:       utils.PointerTo(true),
		},
		GithubActionsOwner:        utils.PointerTo("broadinstitute"),
		GithubActionsRepo:         utils.PointerTo("terra-github-workflows"),
		GithubActionsWorkflowPath: utils.PointerTo(".github/workflows/some-workflow.yaml"),
		GithubActionsDefaultRef:   utils.PointerTo("HEAD"),
		GithubActionsRefBehavior:  utils.PointerTo("always-use-default-ref"),
	}
	s.NoError(s.DB.Create(&environmentGithubHook).Error)

	environmentSlackHook := models.SlackDeployHook{
		Trigger: models.DeployHookTriggerConfig{
			OnEnvironmentID: &environment.ID,
			OnSuccess:       utils.PointerTo(true),
			OnFailure:       utils.PointerTo(true),
		},
		SlackChannel: utils.PointerTo("#workbench-release"),
	}
	s.NoError(s.DB.Create(&environmentSlackHook).Error)

	chartRelease := s.TestData.ChartRelease_LeonardoDev()
	chartReleaseCiIdentifier := chartRelease.GetCiIdentifier()
	s.NoError(s.DB.Create(&chartReleaseCiIdentifier).Error)

	result, err := generateSlackAttachment(s.DB, models.SlackDeployHook{
		Trigger: models.DeployHookTriggerConfig{
			OnEnvironment: &models.Environment{Name: "dev"},
		},
	}, models.CiRun{
		Model:                      gorm.Model{ID: 123},
		TerminalAt:                 utils.PointerTo(time.Now()),
		Status:                     utils.PointerTo("success"),
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
			chartReleaseCiIdentifier, environmentCiIdentifier,
		},
	})
	s.NoError(err)
	s.Equal(slack.GreenBlock{Text: "Deployment to <https://beehive.dsp-devops.broadinstitute.org/r/environment/dev|dev> (<https://beehive.dsp-devops.broadinstitute.org/r/chart-release/leonardo-dev|leonardo-dev>): <https://github.com/broadinstitute/terra-github-workflows/actions/runs/123123/attempts/1|success>. Review all changes made by this deployment <https://beehive.dsp-devops.broadinstitute.org/review-changesets?changeset=1122|here>."}, result)
}
