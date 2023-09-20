package deployhooks

import (
	"encoding/json"
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/github"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	github2 "github.com/google/go-github/v50/github"
	"gorm.io/datatypes"
	"time"
)

func (s *deployHooksSuite) Test_dispatchGithubActionsDeployHook_jsonFailure() {
	arrayRatherThanObject := []any{1, "hi"}
	bytes, err := json.Marshal(arrayRatherThanObject)
	s.NoError(err)
	var inputs datatypes.JSON
	err = json.Unmarshal(bytes, &inputs)
	s.NoError(err)

	err = dispatchGithubActionsDeployHook(nil, models.GithubActionsDeployHook{
		GithubActionsWorkflowInputs: &inputs,
	}, models.CiRun{})
	s.ErrorContains(err, "couldn't unmarshall inputs")
}

func (s *deployHooksSuite) Test_dispatchGithubActionsDeployHook_missingFields() {
	err := dispatchGithubActionsDeployHook(nil, models.GithubActionsDeployHook{}, models.CiRun{})
	s.ErrorContains(err, "GithubActionsDeployHook 0 lacked fields")
}

func (s *deployHooksSuite) Test_dispatchGithubActionsDeployHook_basic() {
	github.UseMockedClient(s.T(), func(c *github.MockClient) {
		c.Actions.EXPECT().CreateWorkflowDispatchEventByFileName(
			s.DB.Statement.Context, "owner", "repo", "path", github2.CreateWorkflowDispatchEventRequest{
				Ref: "main",
			}).Return(nil, nil)
	}, func() {
		s.NoError(dispatchGithubActionsDeployHook(s.DB, models.GithubActionsDeployHook{
			GithubActionsOwner:        utils.PointerTo("owner"),
			GithubActionsRepo:         utils.PointerTo("repo"),
			GithubActionsWorkflowPath: utils.PointerTo("path"),
			GithubActionsDefaultRef:   utils.PointerTo("main"),
		}, models.CiRun{}))
	})
}

func (s *deployHooksSuite) Test_dispatchGithubActionsDeployHook_passesThroughErrors() {
	err := fmt.Errorf("some error")
	github.UseMockedClient(s.T(), func(c *github.MockClient) {
		c.Actions.EXPECT().CreateWorkflowDispatchEventByFileName(
			s.DB.Statement.Context, "owner", "repo", "path", github2.CreateWorkflowDispatchEventRequest{
				Ref: "main",
			}).Return(nil, err)
	}, func() {
		s.ErrorIs(dispatchGithubActionsDeployHook(s.DB, models.GithubActionsDeployHook{
			GithubActionsOwner:        utils.PointerTo("owner"),
			GithubActionsRepo:         utils.PointerTo("repo"),
			GithubActionsWorkflowPath: utils.PointerTo("path"),
			GithubActionsDefaultRef:   utils.PointerTo("main"),
		}, models.CiRun{}), err)
	})
}

func (s *deployHooksSuite) Test_dispatchGithubActionsDeployHook_appVersionRef() {
	github.UseMockedClient(s.T(), func(c *github.MockClient) {
		c.Actions.EXPECT().CreateWorkflowDispatchEventByFileName(
			s.DB.Statement.Context, "owner", "repo", "path", github2.CreateWorkflowDispatchEventRequest{
				Ref: "app version",
			}).Return(nil, nil)
	}, func() {
		s.NoError(dispatchGithubActionsDeployHook(s.DB, models.GithubActionsDeployHook{
			GithubActionsOwner:        utils.PointerTo("owner"),
			GithubActionsRepo:         utils.PointerTo("repo"),
			GithubActionsWorkflowPath: utils.PointerTo("path"),
			GithubActionsDefaultRef:   utils.PointerTo("custom ref"),
			GithubActionsRefBehavior:  utils.PointerTo("use-app-version-as-ref"),
			Trigger: models.DeployHookTriggerConfig{
				OnChartRelease: &models.ChartRelease{ChartReleaseVersion: models.ChartReleaseVersion{AppVersionExact: utils.PointerTo("app version")}},
			},
		}, models.CiRun{}))
	})
}

func (s *deployHooksSuite) Test_dispatchGithubActionsDeployHook_appVersionCommitRef() {
	github.UseMockedClient(s.T(), func(c *github.MockClient) {
		c.Actions.EXPECT().CreateWorkflowDispatchEventByFileName(
			s.DB.Statement.Context, "owner", "repo", "path", github2.CreateWorkflowDispatchEventRequest{
				Ref: "app version commit",
			}).Return(nil, nil)
	}, func() {
		s.NoError(dispatchGithubActionsDeployHook(s.DB, models.GithubActionsDeployHook{
			GithubActionsOwner:        utils.PointerTo("owner"),
			GithubActionsRepo:         utils.PointerTo("repo"),
			GithubActionsWorkflowPath: utils.PointerTo("path"),
			GithubActionsDefaultRef:   utils.PointerTo("custom ref"),
			GithubActionsRefBehavior:  utils.PointerTo("use-app-version-commit-as-ref"),
			Trigger: models.DeployHookTriggerConfig{
				OnChartRelease: &models.ChartRelease{ChartReleaseVersion: models.ChartReleaseVersion{AppVersionCommit: utils.PointerTo("app version commit")}},
			},
		}, models.CiRun{}))
	})
}

func (s *deployHooksSuite) Test_dispatchGithubActionsDeployHook_appVersionRefFromChangesets() {
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

	changeset1 := models.Changeset{
		ChartReleaseID: chartRelease.ID,
		AppliedAt:      utils.PointerTo(time.Now().Add(-time.Minute)),
		From: models.ChartReleaseVersion{
			HelmfileRef:          utils.PointerTo("HEAD"),
			HelmfileRefEnabled:   utils.PointerTo(false),
			AppVersionResolver:   utils.PointerTo("exact"),
			AppVersionExact:      utils.PointerTo("v1.1.3"),
			ChartVersionResolver: utils.PointerTo("exact"),
			ChartVersionExact:    utils.PointerTo("v2.3.4"),
		},
		To: models.ChartReleaseVersion{
			HelmfileRef:          utils.PointerTo("HEAD"),
			HelmfileRefEnabled:   utils.PointerTo(false),
			AppVersionResolver:   utils.PointerTo("exact"),
			AppVersionExact:      utils.PointerTo("v1.1.4"),
			ChartVersionResolver: utils.PointerTo("exact"),
			ChartVersionExact:    utils.PointerTo("v2.3.4"),
		},
	}
	s.NoError(s.DB.Create(&changeset1).Error)
	changeset2 := models.Changeset{
		ChartReleaseID: chartRelease.ID,
		AppliedAt:      utils.PointerTo(time.Now()),
		From: models.ChartReleaseVersion{
			HelmfileRef:          utils.PointerTo("HEAD"),
			HelmfileRefEnabled:   utils.PointerTo(false),
			AppVersionResolver:   utils.PointerTo("exact"),
			AppVersionExact:      utils.PointerTo("v1.1.4"),
			ChartVersionResolver: utils.PointerTo("exact"),
			ChartVersionExact:    utils.PointerTo("v2.3.4"),
		},
		To: models.ChartReleaseVersion{
			HelmfileRef:          utils.PointerTo("HEAD"),
			HelmfileRefEnabled:   utils.PointerTo(false),
			AppVersionResolver:   utils.PointerTo("exact"),
			AppVersionExact:      utils.PointerTo("v1.1.5"),
			ChartVersionResolver: utils.PointerTo("exact"),
			ChartVersionExact:    utils.PointerTo("v2.3.4"),
		},
	}
	s.NoError(s.DB.Create(&changeset2).Error)

	github.UseMockedClient(s.T(), func(c *github.MockClient) {
		c.Actions.EXPECT().CreateWorkflowDispatchEventByFileName(
			s.DB.Statement.Context, "owner", "repo", "path", github2.CreateWorkflowDispatchEventRequest{
				Ref: "v1.1.5",
			}).Return(nil, nil)
	}, func() {
		s.NoError(dispatchGithubActionsDeployHook(s.DB, models.GithubActionsDeployHook{
			GithubActionsOwner:        utils.PointerTo("owner"),
			GithubActionsRepo:         utils.PointerTo("repo"),
			GithubActionsWorkflowPath: utils.PointerTo("path"),
			GithubActionsDefaultRef:   utils.PointerTo("custom ref"),
			GithubActionsRefBehavior:  utils.PointerTo("use-app-version-as-ref"),
			Trigger: models.DeployHookTriggerConfig{
				OnChartRelease: &chartRelease,
			},
		}, models.CiRun{
			RelatedResources: []models.CiIdentifier{
				{ResourceType: "changeset", ResourceID: changeset2.ID},
				{ResourceType: "changeset", ResourceID: changeset1.ID},
			},
		}))
	})
}

func (s *deployHooksSuite) Test_dispatchGithubActionsDeployHook_appVersionCommitRefFromChangesets() {
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
			AppVersionCommit:     utils.PointerTo("commit a"),
			ChartVersionResolver: utils.PointerTo("exact"),
			ChartVersionExact:    utils.PointerTo("v2.3.4"),
		},
	}
	s.NoError(s.DB.Create(&chartRelease).Error)

	changeset1 := models.Changeset{
		ChartReleaseID: chartRelease.ID,
		AppliedAt:      utils.PointerTo(time.Now().Add(-time.Minute)),
		From: models.ChartReleaseVersion{
			HelmfileRef:          utils.PointerTo("HEAD"),
			HelmfileRefEnabled:   utils.PointerTo(false),
			AppVersionResolver:   utils.PointerTo("exact"),
			AppVersionExact:      utils.PointerTo("v1.1.3"),
			AppVersionCommit:     utils.PointerTo("commit b"),
			ChartVersionResolver: utils.PointerTo("exact"),
			ChartVersionExact:    utils.PointerTo("v2.3.4"),
		},
		To: models.ChartReleaseVersion{
			HelmfileRef:          utils.PointerTo("HEAD"),
			HelmfileRefEnabled:   utils.PointerTo(false),
			AppVersionResolver:   utils.PointerTo("exact"),
			AppVersionExact:      utils.PointerTo("v1.1.4"),
			AppVersionCommit:     utils.PointerTo("commit c"),
			ChartVersionResolver: utils.PointerTo("exact"),
			ChartVersionExact:    utils.PointerTo("v2.3.4"),
		},
	}
	s.NoError(s.DB.Create(&changeset1).Error)
	changeset2 := models.Changeset{
		ChartReleaseID: chartRelease.ID,
		AppliedAt:      utils.PointerTo(time.Now()),
		From: models.ChartReleaseVersion{
			HelmfileRef:          utils.PointerTo("HEAD"),
			HelmfileRefEnabled:   utils.PointerTo(false),
			AppVersionResolver:   utils.PointerTo("exact"),
			AppVersionExact:      utils.PointerTo("v1.1.4"),
			AppVersionCommit:     utils.PointerTo("commit d"),
			ChartVersionResolver: utils.PointerTo("exact"),
			ChartVersionExact:    utils.PointerTo("v2.3.4"),
		},
		To: models.ChartReleaseVersion{
			HelmfileRef:          utils.PointerTo("HEAD"),
			HelmfileRefEnabled:   utils.PointerTo(false),
			AppVersionResolver:   utils.PointerTo("exact"),
			AppVersionExact:      utils.PointerTo("v1.1.5"),
			AppVersionCommit:     utils.PointerTo("commit e"),
			ChartVersionResolver: utils.PointerTo("exact"),
			ChartVersionExact:    utils.PointerTo("v2.3.4"),
		},
	}
	s.NoError(s.DB.Create(&changeset2).Error)

	github.UseMockedClient(s.T(), func(c *github.MockClient) {
		c.Actions.EXPECT().CreateWorkflowDispatchEventByFileName(
			s.DB.Statement.Context, "owner", "repo", "path", github2.CreateWorkflowDispatchEventRequest{
				Ref: "commit e",
			}).Return(nil, nil)
	}, func() {
		s.NoError(dispatchGithubActionsDeployHook(s.DB, models.GithubActionsDeployHook{
			GithubActionsOwner:        utils.PointerTo("owner"),
			GithubActionsRepo:         utils.PointerTo("repo"),
			GithubActionsWorkflowPath: utils.PointerTo("path"),
			GithubActionsDefaultRef:   utils.PointerTo("custom ref"),
			GithubActionsRefBehavior:  utils.PointerTo("use-app-version-commit-as-ref"),
			Trigger: models.DeployHookTriggerConfig{
				OnChartRelease: &chartRelease,
			},
		}, models.CiRun{
			RelatedResources: []models.CiIdentifier{
				{ResourceType: "changeset", ResourceID: changeset2.ID},
				{ResourceType: "changeset", ResourceID: changeset1.ID},
			},
		}))
	})
}
