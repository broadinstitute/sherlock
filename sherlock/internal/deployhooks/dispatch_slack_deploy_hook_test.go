package deployhooks

import (
	"github.com/broadinstitute/sherlock/go-shared/pkg/testutils"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/broadinstitute/sherlock/sherlock/internal/slack"
	"github.com/broadinstitute/sherlock/sherlock/internal/slack/slack_mocks"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
	"time"
)

func (s *deployHooksSuite) Test_dispatchSlackDeployHook_channelNil() {
	s.ErrorContains(dispatchSlackDeployHook(nil, models.SlackDeployHook{}, models.CiRun{}), "slack channel was nil")
}

func (s *deployHooksSuite) Test_dispatchSlackDeployHook_generateError() {
	s.ErrorContains(dispatchSlackDeployHook(nil, models.SlackDeployHook{
		Model:        gorm.Model{ID: 123},
		Trigger:      models.DeployHookTriggerConfig{},
		SlackChannel: testutils.PointerTo("channel"),
	}, models.CiRun{}), "SlackDeployHook 123 didn't have Trigger fully loaded")
}

func (s *deployHooksSuite) Test_dispatchSlackDeployHook() {
	slack.UseMockedClient(s.T(), func(c *slack_mocks.MockMockableClient) {
		c.EXPECT().
			SendMessageContext(s.DB.Statement.Context, "channel", mock.AnythingOfType("slack.MsgOption")).
			Return("", "", "", nil)
	}, func() {
		s.NoError(dispatchSlackDeployHook(s.DB, models.SlackDeployHook{
			Trigger: models.DeployHookTriggerConfig{
				OnEnvironment: &models.Environment{Name: "dev"},
			},
			SlackChannel: testutils.PointerTo("channel"),
		}, models.CiRun{
			Model:                      gorm.Model{ID: 123},
			TerminalAt:                 testutils.PointerTo(time.Now()),
			Status:                     testutils.PointerTo("success"),
			Platform:                   "github-actions",
			GithubActionsOwner:         "broadinstitute",
			GithubActionsRepo:          "terra-github-workflows",
			GithubActionsRunID:         123123,
			GithubActionsAttemptNumber: 1,
		}))
	})
}

func (s *deployHooksSuite) Test_generateSlackAttachment_triggerNotLoaded() {
	// Pass nil for the db because it shouldn't be used in this call
	_, err := generateSlackAttachment(nil, models.SlackDeployHook{
		Model:        gorm.Model{ID: 123},
		Trigger:      models.DeployHookTriggerConfig{},
		SlackChannel: testutils.PointerTo("channel"),
	}, models.CiRun{})
	s.ErrorContains(err, "SlackDeployHook 123 didn't have Trigger fully loaded")
}

func (s *deployHooksSuite) Test_generateSlackAttachment_statusNotPresent() {
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

func (s *deployHooksSuite) Test_generateSlackAttachment_environmentNoChangesets_success() {
	// Pass nil for the db because it shouldn't be used in this call
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
	s.Equal(slack.GreenBlock{Text: "Deployment to <https://beehive.dsp-devops.broadinstitute.org/r/environment/dev|dev>: <https://github.com/broadinstitute/terra-github-workflows/actions/runs/123123/attempts/1|success>."}, result)
}

func (s *deployHooksSuite) Test_generateSlackAttachment_environmentNoChangesets_failure() {
	// Pass nil for the db because it shouldn't be used in this call
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
	s.Equal(slack.RedBlock{Text: "Deployment to <https://beehive.dsp-devops.broadinstitute.org/r/environment/dev|dev>: <https://github.com/broadinstitute/terra-github-workflows/actions/runs/123123/attempts/1|failure>."}, result)
}

func (s *deployHooksSuite) Test_generateSlackAttachment_chartReleaseNoChangesets_success() {
	// Pass nil for the db because it shouldn't be used in this call
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
	s.Equal(slack.GreenBlock{Text: "Deployment to <https://beehive.dsp-devops.broadinstitute.org/r/chart-release/leonardo-dev|leonardo-dev>: <https://github.com/broadinstitute/terra-github-workflows/actions/runs/123123/attempts/1|success>."}, result)
}

func (s *deployHooksSuite) Test_generateSlackAttachment_chartReleaseNoChangesets_failure() {
	// Pass nil for the db because it shouldn't be used in this call
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
	s.Equal(slack.RedBlock{Text: "Deployment to <https://beehive.dsp-devops.broadinstitute.org/r/chart-release/leonardo-dev|leonardo-dev>: <https://github.com/broadinstitute/terra-github-workflows/actions/runs/123123/attempts/1|failure>."}, result)
}

func (s *deployHooksSuite) Test_generateSlackAttachment_environmentChangesets_success() {
	// Pass nil for the db because it shouldn't be used in this call
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

func (s *deployHooksSuite) Test_generateSlackAttachment_environmentChangesets_failure() {
	// Pass nil for the db because it shouldn't be used in this call
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

func (s *deployHooksSuite) Test_generateSlackAttachment_chartReleaseChangesets_success() {
	// Pass nil for the db because it shouldn't be used in this call
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

func (s *deployHooksSuite) Test_generateSlackAttachment_chartReleaseChangesets_failure() {
	// Pass nil for the db because it shouldn't be used in this call
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

func (s *deployHooksSuite) Test_generateSlackAttachment_environmentChangesets_chartReleaseNotFound() {
	_, err := generateSlackAttachment(s.DB, models.SlackDeployHook{
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
		},
	})
	s.ErrorIs(err, gorm.ErrRecordNotFound)
}

func (s *deployHooksSuite) Test_generateSlackAttachment_environmentChangesets_chartReleaseNames() {
	user := s.SetSuitableTestUserForDB()

	chart := models.Chart{
		Name:      "leonardo",
		ChartRepo: testutils.PointerTo("terra-helm"),
	}
	s.NoError(s.DB.Create(&chart).Error)

	cluster := models.Cluster{
		Name:                "terra-prod",
		HelmfileRef:         testutils.PointerTo("HEAD"),
		RequiresSuitability: testutils.PointerTo(true),
		Provider:            "google",
		GoogleProject:       "broad-dsde-prod",
		Location:            "us-central1-a",
		Base:                testutils.PointerTo("terra"),
		Address:             testutils.PointerTo("0.0.0.0"),
	}
	s.NoError(s.DB.Create(&cluster).Error)

	environment := models.Environment{
		Name:                 "prod",
		ValuesName:           "prod",
		UniqueResourcePrefix: "a123",
		HelmfileRef:          testutils.PointerTo("HEAD"),
		RequiresSuitability:  testutils.PointerTo(true),
		Base:                 "live",
		DefaultClusterID:     &cluster.ID,
		DefaultNamespace:     "terra-prod",
		Lifecycle:            "static",
		PreventDeletion:      testutils.PointerTo(true),
		OwnerID:              &user.ID,
	}
	s.NoError(s.DB.Create(&environment).Error)
	environmentCiIdentifier := environment.GetCiIdentifier()
	s.NoError(s.DB.Create(&environmentCiIdentifier).Error)

	environmentGithubHook := models.GithubActionsDeployHook{
		Trigger: models.DeployHookTriggerConfig{
			OnEnvironmentID: &environment.ID,
			OnSuccess:       testutils.PointerTo(true),
			OnFailure:       testutils.PointerTo(true),
		},
		GithubActionsOwner:        testutils.PointerTo("broadinstitute"),
		GithubActionsRepo:         testutils.PointerTo("terra-github-workflows"),
		GithubActionsWorkflowPath: testutils.PointerTo(".github/workflows/some-workflow.yaml"),
		GithubActionsDefaultRef:   testutils.PointerTo("HEAD"),
		GithubActionsRefBehavior:  testutils.PointerTo("always-use-default-ref"),
	}
	s.NoError(s.DB.Create(&environmentGithubHook).Error)

	environmentSlackHook := models.SlackDeployHook{
		Trigger: models.DeployHookTriggerConfig{
			OnEnvironmentID: &environment.ID,
			OnSuccess:       testutils.PointerTo(true),
			OnFailure:       testutils.PointerTo(true),
		},
		SlackChannel: testutils.PointerTo("#workbench-release"),
	}
	s.NoError(s.DB.Create(&environmentSlackHook).Error)

	chartRelease := models.ChartRelease{
		Name:            "leonardo-prod",
		ChartID:         chart.ID,
		ClusterID:       &cluster.ID,
		EnvironmentID:   &environment.ID,
		Namespace:       "terra-prod",
		DestinationType: "environment",
		ChartReleaseVersion: models.ChartReleaseVersion{
			HelmfileRef:          testutils.PointerTo("HEAD"),
			HelmfileRefEnabled:   testutils.PointerTo(false),
			AppVersionResolver:   testutils.PointerTo("exact"),
			AppVersionExact:      testutils.PointerTo("v1.2.3"),
			ChartVersionResolver: testutils.PointerTo("exact"),
			ChartVersionExact:    testutils.PointerTo("v2.3.4"),
		},
	}
	s.NoError(s.DB.Create(&chartRelease).Error)
	chartReleaseCiIdentifier := chartRelease.GetCiIdentifier()
	s.NoError(s.DB.Create(&chartReleaseCiIdentifier).Error)

	result, err := generateSlackAttachment(s.DB, models.SlackDeployHook{
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
			chartReleaseCiIdentifier, environmentCiIdentifier,
		},
	})
	s.NoError(err)
	s.Equal(slack.GreenBlock{Text: "Deployment to <https://beehive.dsp-devops.broadinstitute.org/r/environment/dev|dev> (<https://beehive.dsp-devops.broadinstitute.org/r/chart-release/leonardo-prod|leonardo-prod>): <https://github.com/broadinstitute/terra-github-workflows/actions/runs/123123/attempts/1|success>. Review all changes made by this deployment <https://beehive.dsp-devops.broadinstitute.org/review-changesets?changeset=1122|here>."}, result)
}
