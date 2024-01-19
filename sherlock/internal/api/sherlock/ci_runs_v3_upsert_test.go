package sherlock

import (
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/authentication/gha_oidc"
	"github.com/broadinstitute/sherlock/sherlock/internal/authentication/gha_oidc/gha_oidc_claims"
	"github.com/broadinstitute/sherlock/sherlock/internal/authentication/gha_oidc/gha_oidc_mocks"
	"github.com/broadinstitute/sherlock/sherlock/internal/deprecated_controllers/v2controllers"
	"github.com/broadinstitute/sherlock/sherlock/internal/deprecated_models/v2models"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/hooks"
	"github.com/broadinstitute/sherlock/sherlock/internal/hooks/hooks_mocks"
	"github.com/broadinstitute/sherlock/sherlock/internal/slack"
	"github.com/broadinstitute/sherlock/sherlock/internal/slack/slack_mocks"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
	"net/http"
	"time"
)

func (s *handlerSuite) TestCiRunsV3Upsert_error() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("PUT", "/api/ci-runs/v3", CiRunV3Upsert{}),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
}

func (s *handlerSuite) TestCiRunsV3Upsert_edits() {
	startedAt := time.Now().Add(-time.Minute)
	var got1 CiRunV3
	code := s.HandleRequest(
		s.NewRequest("PUT", "/api/ci-runs/v3", CiRunV3Upsert{
			ciRunFields: ciRunFields{
				Platform:                   "github-actions",
				GithubActionsOwner:         "owner",
				GithubActionsRepo:          "repo",
				GithubActionsRunID:         1,
				GithubActionsAttemptNumber: 1,
				GithubActionsWorkflowPath:  "workflow",
				StartedAt:                  &startedAt,
			},
		}),
		&got1)
	s.Equal(http.StatusCreated, code)
	s.WithinDuration(startedAt, *got1.StartedAt, time.Second) // Database stores with less precision
	s.Nil(got1.NotifySlackCustomIcon)
	var got2 CiRunV3
	code = s.HandleRequest(
		s.NewRequest("PUT", "/api/ci-runs/v3", CiRunV3Upsert{
			ciRunFields: ciRunFields{
				Platform:                   "github-actions",
				GithubActionsOwner:         "owner",
				GithubActionsRepo:          "repo",
				GithubActionsRunID:         1,
				GithubActionsAttemptNumber: 1,
				GithubActionsWorkflowPath:  "workflow",
				Status:                     utils.PointerTo("in_progress"),
				NotifySlackCustomIcon:      utils.PointerTo(":smiley:"),
			},
		}),
		&got2)
	s.Equal(http.StatusCreated, code)
	s.Equal("in_progress", *got2.Status)
	s.Equal(got1.StartedAt, got2.StartedAt)
	s.Equal(got1.ID, got2.ID)
	s.NotEqual(got1.UpdatedAt, got2.UpdatedAt)
	if s.NotNil(got2.NotifySlackCustomIcon) {
		s.Equal(":smiley:", *got2.NotifySlackCustomIcon)
	}
}

func (s *handlerSuite) TestCiRunsV3Upsert_fieldValidation() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("PUT", "/api/ci-runs/v3", CiRunV3Upsert{
			ciRunFields: ciRunFields{
				Platform:                   "github-actions",
				GithubActionsOwner:         "owner",
				GithubActionsRepo:          "repo",
				GithubActionsRunID:         1,
				GithubActionsAttemptNumber: 1,
				GithubActionsWorkflowPath:  "workflow",
			},
			RelateToChangesetNewVersions: "some invalid value",
		}),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Contains(got.Message, "RelateToChangesetNewVersions")
}

func (s *handlerSuite) TestCiRunsV3Upsert_identifiers() {
	user := s.SetSuitableTestUserForDB()

	chart := s.TestData.Chart_Leonardo()
	chartVersion, created, err := v2models.InternalChartVersionStore.Create(s.DB, v2models.ChartVersion{
		ChartVersion: "v1.2.3",
		ChartID:      chart.ID,
	}, user)
	s.NoError(err)
	s.True(created)

	s.Run("basic upsert of identifiers", func() {
		var got CiRunV3
		code := s.HandleRequest(
			s.NewRequest("PUT", "/api/ci-runs/v3", CiRunV3Upsert{
				ciRunFields: ciRunFields{
					Platform:                   "github-actions",
					GithubActionsOwner:         "owner",
					GithubActionsRepo:          "repo",
					GithubActionsRunID:         1,
					GithubActionsAttemptNumber: 1,
					GithubActionsWorkflowPath:  "workflow",
				},
				Charts:        []string{"leonardo"},
				ChartVersions: []string{"leonardo/v1.2.3"},
			}),
			&got)
		s.Equal(http.StatusCreated, code)
		s.Len(got.RelatedResources, 2)
	})

	appVersion, created, err := v2models.InternalAppVersionStore.Create(s.DB, v2models.AppVersion{
		AppVersion: "v2.3.4",
		ChartID:    chart.ID,
	}, user)
	s.NoError(err)
	s.True(created)
	cluster, created, err := v2models.InternalClusterStore.Create(s.DB, v2models.Cluster{
		Name:                "terra-dev",
		Provider:            "google",
		GoogleProject:       "broad-dsde-dev",
		Base:                utils.PointerTo("live"),
		Address:             utils.PointerTo("0.0.0.0"),
		RequiresSuitability: utils.PointerTo(false),
		Location:            "us-central1-a",
		HelmfileRef:         utils.PointerTo("HEAD"),
	}, user)
	s.NoError(err)
	s.True(created)
	environment, created, err := v2models.InternalEnvironmentStore.Create(s.DB, v2models.Environment{
		Name:                       "dev",
		Lifecycle:                  "static",
		UniqueResourcePrefix:       "a1b2",
		Base:                       "live",
		DefaultClusterID:           &cluster.ID,
		DefaultNamespace:           "terra-dev",
		OwnerID:                    &user.ID,
		RequiresSuitability:        utils.PointerTo(false),
		HelmfileRef:                utils.PointerTo("HEAD"),
		DefaultFirecloudDevelopRef: utils.PointerTo("dev"),
		PreventDeletion:            utils.PointerTo(false),
	}, user)
	s.NoError(err)
	s.True(created)
	templateEnvironment, created, err := v2models.InternalEnvironmentStore.Create(s.DB, v2models.Environment{
		Name:                       "bee-template",
		Lifecycle:                  "template",
		UniqueResourcePrefix:       "a1b3",
		Base:                       "bee",
		DefaultClusterID:           &cluster.ID,
		DefaultNamespace:           "terra-bee-template",
		OwnerID:                    &user.ID,
		RequiresSuitability:        utils.PointerTo(false),
		HelmfileRef:                utils.PointerTo("HEAD"),
		DefaultFirecloudDevelopRef: utils.PointerTo("dev"),
		PreventDeletion:            utils.PointerTo(false),
	}, user)
	s.NoError(err)
	s.True(created)
	chartRelease, created, err := v2models.InternalChartReleaseStore.Create(s.DB, v2models.ChartRelease{
		Name:          "leonardo-dev",
		ChartID:       chart.ID,
		ClusterID:     &cluster.ID,
		EnvironmentID: &environment.ID,
		Namespace:     environment.DefaultNamespace,
		ChartReleaseVersion: v2models.ChartReleaseVersion{
			AppVersionResolver:   utils.PointerTo("exact"),
			AppVersionExact:      utils.PointerTo("app version blah"),
			ChartVersionResolver: utils.PointerTo("exact"),
			ChartVersionExact:    utils.PointerTo("chart version blah"),
			HelmfileRef:          utils.PointerTo("HEAD"),
			HelmfileRefEnabled:   utils.PointerTo(false),
			FirecloudDevelopRef:  utils.PointerTo("dev"),
		},
	}, user)
	s.NoError(err)
	s.True(created)
	templateChartRelease, created, err := v2models.InternalChartReleaseStore.Create(s.DB, v2models.ChartRelease{
		Name:          "leonardo-bee-template",
		ChartID:       chart.ID,
		ClusterID:     &cluster.ID,
		EnvironmentID: &templateEnvironment.ID,
		Namespace:     templateEnvironment.DefaultNamespace,
		ChartReleaseVersion: v2models.ChartReleaseVersion{
			AppVersionResolver:   utils.PointerTo("exact"),
			AppVersionExact:      utils.PointerTo("app version blah"),
			ChartVersionResolver: utils.PointerTo("exact"),
			ChartVersionExact:    utils.PointerTo("chart version blah"),
			HelmfileRef:          utils.PointerTo("HEAD"),
			HelmfileRefEnabled:   utils.PointerTo(false),
			FirecloudDevelopRef:  utils.PointerTo("dev"),
		},
	}, user)
	s.NoError(err)
	s.True(created)

	s.Run("chart release identifiers", func() {
		var got CiRunV3
		code := s.HandleRequest(
			s.NewRequest("PUT", "/api/ci-runs/v3", CiRunV3Upsert{
				ciRunFields: ciRunFields{
					Platform:                   "github-actions",
					GithubActionsOwner:         "owner",
					GithubActionsRepo:          "repo",
					GithubActionsRunID:         123123,
					GithubActionsAttemptNumber: 1,
					GithubActionsWorkflowPath:  "workflow",
				},
				ChartReleases: []string{chartRelease.Name},
			}),
			&got)
		s.Equal(http.StatusCreated, code)
		s.Len(got.RelatedResources, 3)
	})

	controllerChangesets, err := v2controllers.NewControllerSet(v2models.NewStoreSet(s.DB)).ChangesetController.PlanAndApply(v2controllers.ChangesetPlanRequest{
		ChartReleases: []v2controllers.ChangesetPlanRequestChartReleaseEntry{
			{
				CreatableChangeset: v2controllers.CreatableChangeset{
					ChartRelease:        chartRelease.Name,
					ToAppVersionExact:   &appVersion.AppVersion,
					ToChartVersionExact: &chartVersion.ChartVersion,
				},
			},
			{
				CreatableChangeset: v2controllers.CreatableChangeset{
					ChartRelease:        templateChartRelease.Name,
					ToAppVersionExact:   &appVersion.AppVersion,
					ToChartVersionExact: &chartVersion.ChartVersion,
				},
			},
		},
	}, user)
	s.NoError(err)
	s.Len(controllerChangesets, 2)

	// The changesets are in a somewhat unpredictable order when they come back, so we do this just to make sure
	// that we know which one is which when we get their database form
	var staticEnvironmentChangesetID, templateEnvironmentChangesetID uint
	if controllerChangesets[0].ChartReleaseInfo.ID == chartRelease.ID {
		staticEnvironmentChangesetID = controllerChangesets[0].ID
		templateEnvironmentChangesetID = controllerChangesets[1].ID
	} else {
		staticEnvironmentChangesetID = controllerChangesets[1].ID
		templateEnvironmentChangesetID = controllerChangesets[0].ID
	}

	changeset, err := v2models.InternalChangesetStore.Get(s.DB, v2models.Changeset{Model: gorm.Model{ID: staticEnvironmentChangesetID}})
	s.NoError(err)
	s.Equal(chartRelease.ID, changeset.ChartReleaseID)
	s.Len(changeset.NewAppVersions, 1)
	s.Len(changeset.NewChartVersions, 1)
	templateChangeset, err := v2models.InternalChangesetStore.Get(s.DB, v2models.Changeset{Model: gorm.Model{ID: templateEnvironmentChangesetID}})
	s.NoError(err)
	s.Equal(templateChartRelease.ID, templateChangeset.ChartReleaseID)
	s.Len(templateChangeset.NewAppVersions, 1)
	s.Len(templateChangeset.NewChartVersions, 1)

	s.Run("more advanced upsert of identifiers", func() {
		var got CiRunV3
		code := s.HandleRequest(
			s.NewRequest("PUT", "/api/ci-runs/v3", CiRunV3Upsert{
				ciRunFields: ciRunFields{
					Platform:                   "github-actions",
					GithubActionsOwner:         "owner",
					GithubActionsRepo:          "repo",
					GithubActionsRunID:         1,
					GithubActionsAttemptNumber: 2,
					GithubActionsWorkflowPath:  "workflow",
				},
				Changesets:                   []string{utils.UintToString(changeset.ID)},
				RelateToChangesetNewVersions: "never",
			}),
			&got)
		s.Equal(http.StatusCreated, code)
		s.Len(got.RelatedResources, 4)
	})
	s.Run("more advanced upsert of identifiers with more spreading", func() {
		var got CiRunV3
		code := s.HandleRequest(
			s.NewRequest("PUT", "/api/ci-runs/v3", CiRunV3Upsert{
				ciRunFields: ciRunFields{
					Platform:                   "github-actions",
					GithubActionsOwner:         "owner",
					GithubActionsRepo:          "repo",
					GithubActionsRunID:         1,
					GithubActionsAttemptNumber: 2,
					GithubActionsWorkflowPath:  "workflow",
				},
				Changesets: []string{utils.UintToString(changeset.ID)},
				// Use default for RelateToChangesetNewVersions, which spreads for static environments
			}),
			&got)
		s.Equal(http.StatusCreated, code)
		s.Len(got.RelatedResources, 6)
		s.Run("put loads identifiers", func() {
			var gotAgain CiRunV3
			code = s.HandleRequest(
				s.NewRequest("PUT", "/api/ci-runs/v3", CiRunV3Upsert{
					ciRunFields: ciRunFields{
						Platform:                   "github-actions",
						GithubActionsOwner:         "owner",
						GithubActionsRepo:          "repo",
						GithubActionsRunID:         1,
						GithubActionsAttemptNumber: 2,
						GithubActionsWorkflowPath:  "workflow",
					},
				}),
				&gotAgain)
			s.Equal(http.StatusCreated, code)
			s.Len(gotAgain.RelatedResources, 6)
		})
		s.Run("get loads identifiers too", func() {
			var gotAgain CiRunV3
			code = s.HandleRequest(
				s.NewRequest("GET", fmt.Sprintf("/api/ci-runs/v3/%d", got.ID), nil),
				&gotAgain)
			s.Equal(http.StatusOK, code)
			s.Len(gotAgain.RelatedResources, 6)
		})
	})
	s.Run("changeset spreading for non-static only when set to always", func() {
		var got CiRunV3
		code := s.HandleRequest(
			s.NewRequest("PUT", "/api/ci-runs/v3", CiRunV3Upsert{
				ciRunFields: ciRunFields{
					Platform:                   "github-actions",
					GithubActionsOwner:         "owner",
					GithubActionsRepo:          "repo",
					GithubActionsRunID:         1,
					GithubActionsAttemptNumber: 200,
					GithubActionsWorkflowPath:  "workflow",
				},
				Changesets:                   []string{utils.UintToString(templateChangeset.ID)},
				RelateToChangesetNewVersions: "never",
			}),
			&got)
		s.Run("never", func() {
			s.Equal(http.StatusCreated, code)
			s.Len(got.RelatedResources, 4)
		})
		code = s.HandleRequest(
			s.NewRequest("PUT", "/api/ci-runs/v3", CiRunV3Upsert{
				ciRunFields: ciRunFields{
					Platform:                   "github-actions",
					GithubActionsOwner:         "owner",
					GithubActionsRepo:          "repo",
					GithubActionsRunID:         1,
					GithubActionsAttemptNumber: 200,
					GithubActionsWorkflowPath:  "workflow",
				},
				Changesets: []string{utils.UintToString(templateChangeset.ID)},
			}),
			&got)
		s.Run("default", func() {
			s.Equal(http.StatusCreated, code)
			s.Len(got.RelatedResources, 4)
		})
		code = s.HandleRequest(
			s.NewRequest("PUT", "/api/ci-runs/v3", CiRunV3Upsert{
				ciRunFields: ciRunFields{
					Platform:                   "github-actions",
					GithubActionsOwner:         "owner",
					GithubActionsRepo:          "repo",
					GithubActionsRunID:         1,
					GithubActionsAttemptNumber: 200,
					GithubActionsWorkflowPath:  "workflow",
				},
				Changesets:                   []string{utils.UintToString(templateChangeset.ID)},
				RelateToChangesetNewVersions: "when-static",
			}),
			&got)
		s.Run("explicit when-static", func() {
			s.Equal(http.StatusCreated, code)
			s.Len(got.RelatedResources, 4)
		})
		code = s.HandleRequest(
			s.NewRequest("PUT", "/api/ci-runs/v3", CiRunV3Upsert{
				ciRunFields: ciRunFields{
					Platform:                   "github-actions",
					GithubActionsOwner:         "owner",
					GithubActionsRepo:          "repo",
					GithubActionsRunID:         1,
					GithubActionsAttemptNumber: 200,
					GithubActionsWorkflowPath:  "workflow",
				},
				Changesets:                   []string{utils.UintToString(templateChangeset.ID)},
				RelateToChangesetNewVersions: "always",
			}),
			&got)
		s.Run("always", func() {
			s.Equal(http.StatusCreated, code)
			s.Len(got.RelatedResources, 6)
		})
	})
	s.Run("eliminates duplicates and spreads downwards", func() {
		var got CiRunV3
		code := s.HandleRequest(
			s.NewRequest("PUT", "/api/ci-runs/v3", CiRunV3Upsert{
				ciRunFields: ciRunFields{
					Platform:                   "github-actions",
					GithubActionsOwner:         "owner",
					GithubActionsRepo:          "repo",
					GithubActionsRunID:         1,
					GithubActionsAttemptNumber: 5,
					GithubActionsWorkflowPath:  "workflow",
				},
				Environments: []string{"dev", "dev", utils.UintToString(environment.ID)},
			}),
			&got)
		s.Equal(http.StatusCreated, code)
		s.Len(got.RelatedResources, 3)
		code = s.HandleRequest(
			s.NewRequest("PUT", "/api/ci-runs/v3", CiRunV3Upsert{
				ciRunFields: ciRunFields{
					Platform:                   "github-actions",
					GithubActionsOwner:         "owner",
					GithubActionsRepo:          "repo",
					GithubActionsRunID:         1,
					GithubActionsAttemptNumber: 5,
					GithubActionsWorkflowPath:  "workflow",
				},
				Environments: []string{"dev"},
			}),
			&got)
		s.Equal(http.StatusCreated, code)
		s.Len(got.RelatedResources, 3)
		var envPresent, chartReleasePresent, clusterPresent bool
		for _, ciIdentifier := range got.RelatedResources {
			if ciIdentifier.ResourceType == "environment" {
				s.Equal(environment.ID, ciIdentifier.ResourceID)
				envPresent = true
			}
			if ciIdentifier.ResourceType == "chart-release" {
				s.Equal(chartRelease.ID, ciIdentifier.ResourceID)
				chartReleasePresent = true
			}
			if ciIdentifier.ResourceType == "cluster" {
				s.Equal(cluster.ID, ciIdentifier.ResourceID)
				clusterPresent = true
			}
		}
		s.True(envPresent)
		s.True(chartReleasePresent)
		s.True(clusterPresent)
	})
}

func (s *handlerSuite) TestCiRunsV3Upsert_identifiersInvalid() {
	s.TestData.Chart_Leonardo()

	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("PUT", "/api/ci-runs/v3", CiRunV3Upsert{
			ciRunFields: ciRunFields{
				Platform:                   "github-actions",
				GithubActionsOwner:         "owner",
				GithubActionsRepo:          "repo",
				GithubActionsRunID:         1,
				GithubActionsAttemptNumber: 1,
				GithubActionsWorkflowPath:  "workflow",
			},
			Charts: []string{"leonardo-that-doesn't-exist"},
		}),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(got.Type, errors.BadRequest)
}

func (s *handlerSuite) TestCiRunsV3Upsert_identifiersInvalidIgnore() {
	chart := s.TestData.Chart_Leonardo()

	s.Run("flat ignore", func() {
		var got CiRunV3
		code := s.HandleRequest(
			s.NewRequest("PUT", "/api/ci-runs/v3", CiRunV3Upsert{
				ciRunFields: ciRunFields{
					Platform:                   "github-actions",
					GithubActionsOwner:         "owner",
					GithubActionsRepo:          "repo",
					GithubActionsRunID:         1,
					GithubActionsAttemptNumber: 1,
					GithubActionsWorkflowPath:  "workflow",
				},
				Charts:             []string{"leonardo-that-doesn't-exist"},
				IgnoreBadSelectors: true,
			}),
			&got)
		s.Equal(http.StatusCreated, code)
		s.Len(got.RelatedResources, 0)
		s.NotZero(got.ID)
	})
	s.Run("opportunistic when ignoring", func() {
		var got CiRunV3
		code := s.HandleRequest(
			s.NewRequest("PUT", "/api/ci-runs/v3", CiRunV3Upsert{
				ciRunFields: ciRunFields{
					Platform:                   "github-actions",
					GithubActionsOwner:         "owner",
					GithubActionsRepo:          "repo",
					GithubActionsRunID:         1,
					GithubActionsAttemptNumber: 1,
					GithubActionsWorkflowPath:  "workflow",
				},
				Charts:             []string{"leonardo-that-doesn't-exist", "leonardo"},
				IgnoreBadSelectors: true,
			}),
			&got)
		s.Equal(http.StatusCreated, code)
		s.Len(got.RelatedResources, 1)
		s.Equal(chart.ID, got.RelatedResources[0].ResourceID)
		s.NotZero(got.ID)
	})
}

func (s *handlerSuite) TestCiRunsV3Upsert_slackNotifications() {
	s.Run("failure", func() {
		slack.UseMockedClient(s.T(), func(c *slack_mocks.MockMockableClient) {
			c.EXPECT().
				SendMessageContext(mock.Anything, "#my-failure-channel", mock.AnythingOfType("slack.MsgOption")).
				Return("", "", "", nil)
		}, func() {
			var got CiRunV3
			code := s.HandleRequest(
				s.NewRequest("PUT", "/api/ci-runs/v3", CiRunV3Upsert{
					ciRunFields: ciRunFields{
						Platform:                       "github-actions",
						GithubActionsOwner:             "owner",
						GithubActionsRepo:              "repo",
						GithubActionsRunID:             1,
						GithubActionsAttemptNumber:     1,
						GithubActionsWorkflowPath:      "workflow",
						StartedAt:                      utils.PointerTo(time.Now().Add(-time.Minute)),
						TerminalAt:                     utils.PointerTo(time.Now()),
						Status:                         utils.PointerTo("failure"),
						NotifySlackChannelsUponSuccess: []string{"#my-success-channel"},
						NotifySlackChannelsUponFailure: []string{"#my-failure-channel"},
					},
				}),
				&got)
			s.Equal(http.StatusCreated, code)
			s.Equal([]string{"#my-success-channel"}, got.NotifySlackChannelsUponSuccess)
			s.Equal([]string{"#my-failure-channel"}, got.NotifySlackChannelsUponFailure)
			s.NotNil(got.TerminationHooksDispatchedAt)
		})
	})
	s.Run("success", func() {
		slack.UseMockedClient(s.T(), func(c *slack_mocks.MockMockableClient) {
			c.EXPECT().
				SendMessageContext(mock.Anything, "#my-success-channel", mock.AnythingOfType("slack.MsgOption")).
				Return("", "", "", nil)
		}, func() {
			var got CiRunV3
			code := s.HandleRequest(
				s.NewRequest("PUT", "/api/ci-runs/v3", CiRunV3Upsert{
					ciRunFields: ciRunFields{
						Platform:                       "github-actions",
						GithubActionsOwner:             "owner",
						GithubActionsRepo:              "repo",
						GithubActionsRunID:             1,
						GithubActionsAttemptNumber:     2,
						GithubActionsWorkflowPath:      "workflow",
						StartedAt:                      utils.PointerTo(time.Now().Add(-time.Minute)),
						TerminalAt:                     utils.PointerTo(time.Now()),
						Status:                         utils.PointerTo("success"),
						NotifySlackChannelsUponSuccess: []string{"#my-success-channel"},
						NotifySlackChannelsUponFailure: []string{"#my-failure-channel"},
					},
				}),
				&got)
			s.Equal(http.StatusCreated, code)
			s.Equal([]string{"#my-success-channel"}, got.NotifySlackChannelsUponSuccess)
			s.Equal([]string{"#my-failure-channel"}, got.NotifySlackChannelsUponFailure)
			s.NotNil(got.TerminationHooksDispatchedAt)
		})
	})
	s.Run("with against string", func() {
		chartRelease := s.TestData.ChartRelease_LeonardoDev()
		slack.UseMockedClient(s.T(), func(c *slack_mocks.MockMockableClient) {
			c.EXPECT().
				// Unfortunately mockery isn't smart enough for us to actually test the text output here but at least
				// we can validate that this scenario doesn't blow things up.
				SendMessageContext(mock.Anything, "#my-success-channel", mock.AnythingOfType("slack.MsgOption")).
				Return("", "", "", nil)
		}, func() {
			var got CiRunV3
			code := s.HandleRequest(
				s.NewRequest("PUT", "/api/ci-runs/v3", CiRunV3Upsert{
					ciRunFields: ciRunFields{
						Platform:                       "github-actions",
						GithubActionsOwner:             "owner",
						GithubActionsRepo:              "repo",
						GithubActionsRunID:             1,
						GithubActionsAttemptNumber:     3,
						GithubActionsWorkflowPath:      "workflow",
						StartedAt:                      utils.PointerTo(time.Now().Add(-time.Minute)),
						TerminalAt:                     utils.PointerTo(time.Now()),
						Status:                         utils.PointerTo("success"),
						NotifySlackChannelsUponSuccess: []string{"#my-success-channel"},
						NotifySlackChannelsUponFailure: []string{"#my-failure-channel"},
					},
					ChartReleases: []string{chartRelease.Name},
				}),
				&got)
			s.Equal(http.StatusCreated, code)
			s.Equal([]string{"#my-success-channel"}, got.NotifySlackChannelsUponSuccess)
			s.Equal([]string{"#my-failure-channel"}, got.NotifySlackChannelsUponFailure)
			s.NotNil(got.TerminationHooksDispatchedAt)
		})
	})
}

func (s *handlerSuite) TestCiRunsV3Upsert_githubActionsClaimDefaults() {
	// Note that the request body is empty!
	// Normally this would result in an error due to missing fields, but suppose a GHA OIDC JWT was passed...
	request := s.NewRequest(http.MethodPut, "/api/ci-runs/v3", CiRunV3Upsert{})
	request.Header.Set(gha_oidc.Header, "some GHA OIDC token")

	var got CiRunV3
	var code int
	gha_oidc.UseMockedVerifier(s.T(), func(v *gha_oidc_mocks.MockMockableVerifier) {
		v.EXPECT().VerifyAndParseClaims(mock.AnythingOfType("*gin.Context"), "some GHA OIDC token").
			Return(gha_oidc_claims.Claims{
				RepositoryOwner: "broadinstitute",
				Repository:      "broadinstitute/terra-github-workflows",
				WorkflowRef:     "broadinstitute/terra-github-workflows/.github/workflows/bee-create.yaml@refs/heads/main",
				RunID:           "123456",
				RunAttempt:      "1",
			}, nil)
	}, func() {
		code = s.HandleRequest(request, &got)
	})

	s.Equal(http.StatusCreated, code)
	s.NotZero(got.ID)
	s.Equal(got.Platform, "github-actions")
	s.Equal(got.GithubActionsOwner, "broadinstitute")
	s.Equal(got.GithubActionsRepo, "terra-github-workflows")
	s.Equal(got.GithubActionsRunID, uint(123456))
	s.Equal(got.GithubActionsAttemptNumber, uint(1))
	s.Equal(got.GithubActionsWorkflowPath, ".github/workflows/bee-create.yaml")
}

func (s *handlerSuite) TestCiRunsV3Upsert_dispatch() {
	chartRelease := s.TestData.ChartRelease_LeonardoDev()
	changeset := s.TestData.Changeset_LeonardoDev_V1toV3()
	s.TestData.SlackDeployHook_Dev()
	s.TestData.GithubActionsDeployHook_LeonardoDev()

	var got CiRunV3
	var code int

	// 1. Suppose a CiRun is upserted via webhook
	hooks.UseMockedDispatcher(s.T(), func(d *hooks_mocks.MockMockableDispatcher) {}, func() {
		code = s.HandleRequest(
			s.NewRequest(http.MethodPut, "/api/ci-runs/v3", CiRunV3Upsert{
				ciRunFields: ciRunFields{
					Platform:                   "github-actions",
					GithubActionsOwner:         "broadinstitute",
					GithubActionsRepo:          "terra-github-workflows",
					GithubActionsRunID:         123,
					GithubActionsAttemptNumber: 1,
					GithubActionsWorkflowPath:  ".github/workflows/sync-release.yaml",
					StartedAt:                  utils.PointerTo(changeset.AppliedAt.Add(30 * time.Second)),
					Status:                     utils.PointerTo("in_progress"),
				},
			}),
			&got)
	})
	s.Equal(http.StatusCreated, code)
	if s.NotNil(got.Status) {
		s.Equal("in_progress", *got.Status)
	}

	// 2. Suppose the relation to the changeset is reported by the action itself, plus a notification channel
	//    (now deploy hooks start firing because we can match the resources)
	hooks.UseMockedDispatcher(s.T(), func(d *hooks_mocks.MockMockableDispatcher) {
		d.EXPECT().DispatchSlackDeployHook(mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()
	}, func() {
		code = s.HandleRequest(
			s.NewRequest(http.MethodPut, "/api/ci-runs/v3", CiRunV3Upsert{
				ciRunFields: ciRunFields{
					Platform:                       "github-actions",
					GithubActionsOwner:             "broadinstitute",
					GithubActionsRepo:              "terra-github-workflows",
					GithubActionsRunID:             123,
					GithubActionsAttemptNumber:     1,
					GithubActionsWorkflowPath:      ".github/workflows/sync-release.yaml",
					NotifySlackChannelsUponSuccess: []string{"#workbench-resilience-dev"},
					NotifySlackChannelsUponFailure: []string{"#workbench-resilience-dev", "#ap-k8s-monitor"},
				},
				Changesets: []string{utils.UintToString(changeset.ID)},
			}),
			&got)
	})
	s.Equal(http.StatusCreated, code)

	s.Run("check CiIdentifier creation", func() {
		relatedResourceCounts := map[string]int{}
		for _, rr := range got.RelatedResources {
			if s.NotZero(rr.ResourceType) && s.NotZero(rr.ResourceID) && s.NotZero(rr.ID) {
				relatedResourceCounts[rr.ResourceType]++
			}
			s.Nil(rr.ResourceStatus)
		}
		s.Equal(1, relatedResourceCounts["changeset"])
		s.Equal(1, relatedResourceCounts["chart-release"])
		s.Equal(1, relatedResourceCounts["cluster"])
		s.Equal(1, relatedResourceCounts["environment"])
		s.Equal(len(changeset.NewAppVersions), relatedResourceCounts["app-version"])
		s.Equal(len(changeset.NewChartVersions), relatedResourceCounts["chart-version"])
	})

	// 3. Suppose Thelma reported the status for the chart release itself
	hooks.UseMockedDispatcher(s.T(), func(d *hooks_mocks.MockMockableDispatcher) {
		d.EXPECT().DispatchSlackDeployHook(mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()
	}, func() {
		code = s.HandleRequest(
			s.NewRequest(http.MethodPut, "/api/ci-runs/v3", CiRunV3Upsert{
				ciRunFields: ciRunFields{
					Platform:                   "github-actions",
					GithubActionsOwner:         "broadinstitute",
					GithubActionsRepo:          "terra-github-workflows",
					GithubActionsRunID:         123,
					GithubActionsAttemptNumber: 1,
					GithubActionsWorkflowPath:  ".github/workflows/sync-release.yaml",
				},
				ChartReleaseStatuses: map[string]string{
					chartRelease.Name: "success: healthy",
				},
			}),
			&got)
	})
	s.Equal(http.StatusCreated, code)

	s.Run("check resource statuses", func() {
		for _, rr := range got.RelatedResources {
			if rr.ResourceType == "chart-release" || rr.ResourceType == "changeset" || rr.ResourceType == "app-version" || rr.ResourceType == "chart-version" {
				if s.NotNil(rr.ResourceStatus) {
					s.Equal("success: healthy", *rr.ResourceStatus)
				}
			} else {
				s.Nil(rr.ResourceStatus)
			}
		}
	})

	// 4. Suppose CiRun marked as completed by via webhook
	hooks.UseMockedDispatcher(s.T(), func(d *hooks_mocks.MockMockableDispatcher) {
		d.EXPECT().DispatchSlackDeployHook(mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()
		d.EXPECT().DispatchGithubActionsDeployHook(mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()
		d.EXPECT().DispatchSlackCompletionNotification(mock.Anything, "#workbench-resilience-dev", mock.Anything, true, mock.AnythingOfType("*string")).Return(nil).Once()
	}, func() {
		code = s.HandleRequest(
			s.NewRequest(http.MethodPut, "/api/ci-runs/v3", CiRunV3Upsert{
				ciRunFields: ciRunFields{
					Platform:                   "github-actions",
					GithubActionsOwner:         "broadinstitute",
					GithubActionsRepo:          "terra-github-workflows",
					GithubActionsRunID:         123,
					GithubActionsAttemptNumber: 1,
					GithubActionsWorkflowPath:  ".github/workflows/sync-release.yaml",
					TerminalAt:                 utils.PointerTo(changeset.AppliedAt.Add(time.Minute)),
					Status:                     utils.PointerTo("success"),
				},
			}),
			&got)
	})
	s.Equal(http.StatusCreated, code)
	if s.NotNil(got.Status) {
		s.Equal("success", *got.Status)
	}
	s.NotNil(got.TerminationHooksDispatchedAt)
}
