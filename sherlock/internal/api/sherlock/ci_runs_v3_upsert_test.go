package sherlock

import (
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/testutils"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/deprecated_controllers/v2controllers"
	"github.com/broadinstitute/sherlock/sherlock/internal/deprecated_models/v2models"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
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

func (s *handlerSuite) TestCiRunsV3Upsert() {
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
				Status:                     testutils.PointerTo("in_progress"),
			},
		}),
		&got2)
	s.Equal(http.StatusCreated, code)
	s.Equal("in_progress", *got2.Status)
	s.Equal(got1.StartedAt, got2.StartedAt)
	s.Equal(got1.ID, got2.ID)
	s.NotEqual(got1.UpdatedAt, got2.UpdatedAt)
}

func (s *handlerSuite) TestCiRunsV3UpsertFieldValidation() {
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

func (s *handlerSuite) TestCiRunsV3UpsertIdentifiers() {
	user := s.SetSuitableTestUserForDB()

	chart, created, err := v2models.InternalChartStore.Create(s.DB, v2models.Chart{
		Name:      "leonardo",
		ChartRepo: testutils.PointerTo("terra-helm"),
	}, user)
	s.NoError(err)
	s.True(created)
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
		Base:                testutils.PointerTo("live"),
		Address:             testutils.PointerTo("0.0.0.0"),
		RequiresSuitability: testutils.PointerTo(false),
		Location:            "us-central1-a",
		HelmfileRef:         testutils.PointerTo("HEAD"),
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
		RequiresSuitability:        testutils.PointerTo(false),
		HelmfileRef:                testutils.PointerTo("HEAD"),
		DefaultFirecloudDevelopRef: testutils.PointerTo("dev"),
		PreventDeletion:            testutils.PointerTo(false),
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
		RequiresSuitability:        testutils.PointerTo(false),
		HelmfileRef:                testutils.PointerTo("HEAD"),
		DefaultFirecloudDevelopRef: testutils.PointerTo("dev"),
		PreventDeletion:            testutils.PointerTo(false),
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
			AppVersionResolver:   testutils.PointerTo("exact"),
			AppVersionExact:      testutils.PointerTo("app version blah"),
			ChartVersionResolver: testutils.PointerTo("exact"),
			ChartVersionExact:    testutils.PointerTo("chart version blah"),
			HelmfileRef:          testutils.PointerTo("HEAD"),
			FirecloudDevelopRef:  testutils.PointerTo("dev"),
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
			AppVersionResolver:   testutils.PointerTo("exact"),
			AppVersionExact:      testutils.PointerTo("app version blah"),
			ChartVersionResolver: testutils.PointerTo("exact"),
			ChartVersionExact:    testutils.PointerTo("chart version blah"),
			HelmfileRef:          testutils.PointerTo("HEAD"),
			FirecloudDevelopRef:  testutils.PointerTo("dev"),
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
	user := s.SetSuitableTestUserForDB()

	_, created, err := v2models.InternalChartStore.Create(s.DB, v2models.Chart{
		Name:      "leonardo",
		ChartRepo: testutils.PointerTo("terra-helm"),
	}, user)
	s.NoError(err)
	s.True(created)

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
	user := s.SetSuitableTestUserForDB()

	chart, created, err := v2models.InternalChartStore.Create(s.DB, v2models.Chart{
		Name:      "leonardo",
		ChartRepo: testutils.PointerTo("terra-helm"),
	}, user)
	s.NoError(err)
	s.True(created)

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
