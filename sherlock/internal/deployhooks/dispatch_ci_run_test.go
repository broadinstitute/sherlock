package deployhooks

import (
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"gorm.io/gorm"
	"time"
)

func (s *deployHooksSuite) Test_dispatchCiRun() {
	user := s.SetSuitableTestUserForDB()

	chart := models.Chart{
		Name:      "leonardo",
		ChartRepo: utils.PointerTo("terra-helm"),
	}
	s.NoError(s.DB.Create(&chart).Error)

	cluster := models.Cluster{
		Name:                "terra-prod",
		HelmfileRef:         utils.PointerTo("HEAD"),
		RequiresSuitability: utils.PointerTo(true),
		Provider:            "google",
		GoogleProject:       "broad-dsde-prod",
		Location:            "us-central1-a",
		Base:                utils.PointerTo("terra"),
		Address:             utils.PointerTo("0.0.0.0"),
	}
	s.NoError(s.DB.Create(&cluster).Error)

	environment := models.Environment{
		Name:                 "prod",
		ValuesName:           "prod",
		UniqueResourcePrefix: "a123",
		HelmfileRef:          utils.PointerTo("HEAD"),
		RequiresSuitability:  utils.PointerTo(true),
		Base:                 "live",
		DefaultClusterID:     &cluster.ID,
		DefaultNamespace:     "terra-prod",
		Lifecycle:            "static",
		PreventDeletion:      utils.PointerTo(true),
		OwnerID:              &user.ID,
	}
	s.NoError(s.DB.Create(&environment).Error)
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

	environmentSuccessOnlyGithubHook := models.GithubActionsDeployHook{
		Trigger: models.DeployHookTriggerConfig{
			OnEnvironmentID: &environment.ID,
			OnSuccess:       utils.PointerTo(true),
			OnFailure:       utils.PointerTo(false),
		},
		GithubActionsOwner:        utils.PointerTo("broadinstitute"),
		GithubActionsRepo:         utils.PointerTo("terra-github-workflows"),
		GithubActionsWorkflowPath: utils.PointerTo(".github/workflows/some-success-workflow.yaml"),
		GithubActionsDefaultRef:   utils.PointerTo("HEAD"),
		GithubActionsRefBehavior:  utils.PointerTo("always-use-default-ref"),
	}
	s.NoError(s.DB.Create(&environmentSuccessOnlyGithubHook).Error)

	environmentFailureOnlyGithubHook := models.GithubActionsDeployHook{
		Trigger: models.DeployHookTriggerConfig{
			OnEnvironmentID: &environment.ID,
			OnSuccess:       utils.PointerTo(false),
			OnFailure:       utils.PointerTo(true),
		},
		GithubActionsOwner:        utils.PointerTo("broadinstitute"),
		GithubActionsRepo:         utils.PointerTo("terra-github-workflows"),
		GithubActionsWorkflowPath: utils.PointerTo(".github/workflows/some-failure-workflow.yaml"),
		GithubActionsDefaultRef:   utils.PointerTo("HEAD"),
		GithubActionsRefBehavior:  utils.PointerTo("always-use-default-ref"),
	}
	s.NoError(s.DB.Create(&environmentFailureOnlyGithubHook).Error)

	environmentSlackHook := models.SlackDeployHook{
		Trigger: models.DeployHookTriggerConfig{
			OnEnvironmentID: &environment.ID,
			OnSuccess:       utils.PointerTo(true),
			OnFailure:       utils.PointerTo(true),
		},
		SlackChannel: utils.PointerTo("#workbench-release"),
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
			HelmfileRef:          utils.PointerTo("HEAD"),
			HelmfileRefEnabled:   utils.PointerTo(false),
			AppVersionResolver:   utils.PointerTo("exact"),
			AppVersionExact:      utils.PointerTo("v1.2.3"),
			ChartVersionResolver: utils.PointerTo("exact"),
			ChartVersionExact:    utils.PointerTo("v2.3.4"),
		},
	}
	s.NoError(s.DB.Create(&chartRelease).Error)
	chartReleaseCiIdentifier := chartRelease.GetCiIdentifier()
	s.NoError(s.DB.Create(&chartReleaseCiIdentifier).Error)

	chartReleaseGithubHook := models.GithubActionsDeployHook{
		Trigger: models.DeployHookTriggerConfig{
			OnChartReleaseID: &chartRelease.ID,
			OnSuccess:        utils.PointerTo(true),
			OnFailure:        utils.PointerTo(true),
		},
		GithubActionsOwner:        utils.PointerTo("broadinstitute"),
		GithubActionsRepo:         utils.PointerTo("sam"),
		GithubActionsWorkflowPath: utils.PointerTo(".github/workflows/some-workflow.yaml"),
		GithubActionsDefaultRef:   utils.PointerTo("HEAD"),
		GithubActionsRefBehavior:  utils.PointerTo("always-use-default-ref"),
	}
	s.NoError(s.DB.Create(&chartReleaseGithubHook).Error)

	chartReleaseSlackHook := models.SlackDeployHook{
		Trigger: models.DeployHookTriggerConfig{
			OnChartReleaseID: &chartRelease.ID,
			OnSuccess:        utils.PointerTo(true),
			OnFailure:        utils.PointerTo(true),
		},
		SlackChannel: utils.PointerTo("#dsp-identiteam"),
	}
	s.NoError(s.DB.Create(&chartReleaseSlackHook).Error)

	chartReleaseSuccessOnlySlackHook := models.SlackDeployHook{
		Trigger: models.DeployHookTriggerConfig{
			OnChartReleaseID: &chartRelease.ID,
			OnSuccess:        utils.PointerTo(true),
			OnFailure:        utils.PointerTo(false),
		},
		SlackChannel: utils.PointerTo("#dsp-identiteam-successes"),
	}
	s.NoError(s.DB.Create(&chartReleaseSuccessOnlySlackHook).Error)

	chartReleaseFailureOnlySlackHook := models.SlackDeployHook{
		Trigger: models.DeployHookTriggerConfig{
			OnChartReleaseID: &chartRelease.ID,
			OnSuccess:        utils.PointerTo(false),
			OnFailure:        utils.PointerTo(true),
		},
		SlackChannel: utils.PointerTo("#dsp-identiteam-failures"),
	}
	s.NoError(s.DB.Create(&chartReleaseFailureOnlySlackHook).Error)

	successfulCiRun := models.CiRun{
		Platform:                   "github-actions",
		GithubActionsOwner:         "broadinstitute",
		GithubActionsRepo:          "terra-github-workflows",
		GithubActionsRunID:         123123,
		GithubActionsAttemptNumber: 1,
		GithubActionsWorkflowPath:  ".github/workflows/sync-release.yaml",
		RelatedResources: []models.CiIdentifier{
			environmentCiIdentifier,
			chartReleaseCiIdentifier,
		},
		Status:     utils.PointerTo("success"),
		TerminalAt: utils.PointerTo(time.Now()),
	}
	s.NoError(s.DB.Create(&successfulCiRun).Error)

	failureCiRun := models.CiRun{
		Platform:                   "github-actions",
		GithubActionsOwner:         "broadinstitute",
		GithubActionsRepo:          "terra-github-workflows",
		GithubActionsRunID:         123123,
		GithubActionsAttemptNumber: 2,
		GithubActionsWorkflowPath:  ".github/workflows/sync-release.yaml",
		RelatedResources: []models.CiIdentifier{
			environmentCiIdentifier,
			chartReleaseCiIdentifier,
		},
		Status:     utils.PointerTo("failure"),
		TerminalAt: utils.PointerTo(time.Now()),
	}
	s.NoError(s.DB.Create(&failureCiRun).Error)

	for _, ciRunSuccess := range []bool{true, false} {
		s.Run(fmt.Sprintf("normal case for success (%v)", ciRunSuccess), func() {
			var calledEnvironmentGithub, calledEnvironmentSuccessGithub, calledEnvironmentFailureGithub, calledEnvironmentSlack,
				calledChartReleaseGithub, calledChartReleaseSlack, calledChartReleaseSuccessSlack, calledChartReleaseFailureSlack bool
			var ciRunIdToUse uint
			if ciRunSuccess {
				ciRunIdToUse = successfulCiRun.ID
			} else {
				ciRunIdToUse = failureCiRun.ID
			}
			errs := dispatchCiRun(
				s.DB,
				models.CiRun{Model: gorm.Model{ID: ciRunIdToUse}},
				func(db *gorm.DB, hook models.SlackDeployHook, ciRun models.CiRun) error {
					if hook.ID == environmentSlackHook.ID {
						calledEnvironmentSlack = true
					} else if hook.ID == chartReleaseSlackHook.ID {
						calledChartReleaseSlack = true
					} else if hook.ID == chartReleaseSuccessOnlySlackHook.ID {
						calledChartReleaseSuccessSlack = true
					} else if hook.ID == chartReleaseFailureOnlySlackHook.ID {
						calledChartReleaseFailureSlack = true
					}
					return nil
				},
				func(db *gorm.DB, hook models.GithubActionsDeployHook, ciRun models.CiRun) error {
					if hook.ID == environmentGithubHook.ID {
						calledEnvironmentGithub = true
					} else if hook.ID == environmentSuccessOnlyGithubHook.ID {
						calledEnvironmentSuccessGithub = true
					} else if hook.ID == environmentFailureOnlyGithubHook.ID {
						calledEnvironmentFailureGithub = true
					} else if hook.ID == chartReleaseGithubHook.ID {
						calledChartReleaseGithub = true
					}
					return nil
				})
			s.Empty(errs)
			s.True(calledEnvironmentGithub)
			s.Equal(ciRunSuccess, calledEnvironmentSuccessGithub)
			s.Equal(ciRunSuccess, !calledEnvironmentFailureGithub)
			s.True(calledEnvironmentSlack)
			s.True(calledChartReleaseGithub)
			s.True(calledChartReleaseSlack)
			s.Equal(ciRunSuccess, calledChartReleaseSuccessSlack)
			s.Equal(ciRunSuccess, !calledChartReleaseFailureSlack)
		})
	}
	s.Run("collects errors", func() {
		errs := dispatchCiRun(
			s.DB,
			models.CiRun{Model: gorm.Model{ID: successfulCiRun.ID}},
			func(db *gorm.DB, hook models.SlackDeployHook, ciRun models.CiRun) error {
				return fmt.Errorf("error")
			},
			func(db *gorm.DB, hook models.GithubActionsDeployHook, ciRun models.CiRun) error {
				return fmt.Errorf("error")
			})
		s.Len(errs, 6)
	})
	s.Run("collects errors from deleted hooks", func() {
		s.NoError(s.DB.Where(&environmentGithubHook).Delete(&environmentGithubHook).Error)
		s.NoError(s.DB.Where(&environmentSlackHook).Delete(&environmentSlackHook).Error)
		s.NoError(s.DB.Where(&chartReleaseGithubHook).Delete(&chartReleaseGithubHook).Error)
		s.NoError(s.DB.Where(&chartReleaseSlackHook).Delete(&chartReleaseSlackHook).Error)
		errs := dispatchCiRun(
			s.DB,
			models.CiRun{Model: gorm.Model{ID: successfulCiRun.ID}},
			func(db *gorm.DB, hook models.SlackDeployHook, ciRun models.CiRun) error {
				return nil
			},
			func(db *gorm.DB, hook models.GithubActionsDeployHook, ciRun models.CiRun) error {
				return nil
			})
		s.Len(errs, 4)
	})
}
