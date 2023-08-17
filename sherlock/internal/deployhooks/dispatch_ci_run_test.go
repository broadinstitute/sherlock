package deployhooks

import (
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/testutils"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"gorm.io/gorm"
)

func (s *deployHooksSuite) Test_dispatchCiRun() {
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
			AppVersionResolver:   testutils.PointerTo("exact"),
			AppVersionExact:      testutils.PointerTo("v1.2.3"),
			ChartVersionResolver: testutils.PointerTo("exact"),
			ChartVersionExact:    testutils.PointerTo("v2.3.4"),
		},
	}
	s.NoError(s.DB.Create(&chartRelease).Error)
	chartReleaseCiIdentifier := chartRelease.GetCiIdentifier()
	s.NoError(s.DB.Create(&chartReleaseCiIdentifier).Error)

	chartReleaseGithubHook := models.GithubActionsDeployHook{
		Trigger: models.DeployHookTriggerConfig{
			OnChartReleaseID: &chartRelease.ID,
			OnSuccess:        testutils.PointerTo(true),
			OnFailure:        testutils.PointerTo(true),
		},
		GithubActionsOwner:        testutils.PointerTo("broadinstitute"),
		GithubActionsRepo:         testutils.PointerTo("sam"),
		GithubActionsWorkflowPath: testutils.PointerTo(".github/workflows/some-workflow.yaml"),
		GithubActionsDefaultRef:   testutils.PointerTo("HEAD"),
		GithubActionsRefBehavior:  testutils.PointerTo("always-use-default-ref"),
	}
	s.NoError(s.DB.Create(&chartReleaseGithubHook).Error)

	chartReleaseSlackHook := models.SlackDeployHook{
		Trigger: models.DeployHookTriggerConfig{
			OnChartReleaseID: &chartRelease.ID,
			OnSuccess:        testutils.PointerTo(true),
			OnFailure:        testutils.PointerTo(true),
		},
		SlackChannel: testutils.PointerTo("#dsp-identiteam"),
	}
	s.NoError(s.DB.Create(&chartReleaseSlackHook).Error)

	ciRun := models.CiRun{
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
	}
	s.NoError(s.DB.Create(&ciRun).Error)

	s.Run("normal case", func() {
		var calledEnvironmentGithub, calledEnvironmentSlack, calledChartReleaseGithub, calledChartReleaseSlack bool
		errs := dispatchCiRun(
			s.DB,
			models.CiRun{Model: gorm.Model{ID: ciRun.ID}},
			func(db *gorm.DB, hook models.SlackDeployHook, ciRun models.CiRun) error {
				if hook.ID == environmentSlackHook.ID {
					calledEnvironmentSlack = true
				} else if hook.ID == chartReleaseSlackHook.ID {
					calledChartReleaseSlack = true
				}
				return nil
			},
			func(db *gorm.DB, hook models.GithubActionsDeployHook, ciRun models.CiRun) error {
				if hook.ID == environmentGithubHook.ID {
					calledEnvironmentGithub = true
				} else if hook.ID == chartReleaseGithubHook.ID {
					calledChartReleaseGithub = true
				}
				return nil
			})
		s.Empty(errs)
		s.True(calledEnvironmentGithub)
		s.True(calledEnvironmentSlack)
		s.True(calledChartReleaseGithub)
		s.True(calledChartReleaseSlack)
	})
	s.Run("collects errors", func() {
		errs := dispatchCiRun(
			s.DB,
			models.CiRun{Model: gorm.Model{ID: ciRun.ID}},
			func(db *gorm.DB, hook models.SlackDeployHook, ciRun models.CiRun) error {
				return fmt.Errorf("error")
			},
			func(db *gorm.DB, hook models.GithubActionsDeployHook, ciRun models.CiRun) error {
				return fmt.Errorf("error")
			})
		s.Len(errs, 4)
	})
	s.Run("collects errors from deleted hooks", func() {
		s.NoError(s.DB.Where(&environmentGithubHook).Delete(&environmentGithubHook).Error)
		s.NoError(s.DB.Where(&environmentSlackHook).Delete(&environmentSlackHook).Error)
		s.NoError(s.DB.Where(&chartReleaseGithubHook).Delete(&chartReleaseGithubHook).Error)
		s.NoError(s.DB.Where(&chartReleaseSlackHook).Delete(&chartReleaseSlackHook).Error)
		errs := dispatchCiRun(
			s.DB,
			models.CiRun{Model: gorm.Model{ID: ciRun.ID}},
			func(db *gorm.DB, hook models.SlackDeployHook, ciRun models.CiRun) error {
				s.Fail("should have errored")
				return nil
			},
			func(db *gorm.DB, hook models.GithubActionsDeployHook, ciRun models.CiRun) error {
				return fmt.Errorf("should have errored")
			})
		s.Len(errs, 4)
	})
}
