package sherlock

import (
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/testutils"
	"github.com/broadinstitute/sherlock/sherlock/internal/deprecated_controllers/v2controllers"
	"github.com/broadinstitute/sherlock/sherlock/internal/deprecated_models/v2models"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"gorm.io/gorm"
	"net/http"
	"time"
)

func (s *handlerSuite) TestCiIdentifiersV3GetBadSelector() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/ci-identifiers/v3/foo-bar", nil),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
}

func (s *handlerSuite) TestCiIdentifiersV3GetNotFound() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/ci-identifiers/v3/0", nil),
		&got)
	s.Equal(http.StatusNotFound, code)
	s.Equal(errors.NotFound, got.Type)
}

func (s *handlerSuite) TestCiIdentifiersV3Get() {
	user := s.SetSuitableTestUserForDB()

	chart, created, err := v2models.InternalChartStore.Create(s.DB, v2models.Chart{
		Name:      "leonardo",
		ChartRepo: testutils.PointerTo("terra-helm"),
	}, user)
	s.NoError(err)
	s.True(created)
	chartIdentifier := ciIdentifierModelFromOldModel(chart)
	err = s.DB.Create(&chartIdentifier).Error
	s.NoError(err)
	s.Equal(chart.ID, chartIdentifier.ResourceID)

	s.Run("chart identifier", func() {
		var gotByID, gotByName CiIdentifierV3
		code := s.HandleRequest(
			s.NewRequest("GET", fmt.Sprintf("/api/ci-identifiers/v3/%d", chartIdentifier.ID), nil),
			&gotByID)
		s.Equal(http.StatusOK, code)
		s.Equal(chartIdentifier.ID, gotByID.ID)
		code = s.HandleRequest(
			s.NewRequest("GET", "/api/ci-identifiers/v3/chart/leonardo", nil),
			&gotByName)
		s.Equal(http.StatusOK, code)
		s.Equal(chartIdentifier.ID, gotByName.ID)
	})

	chartVersion, created, err := v2models.InternalChartVersionStore.Create(s.DB, v2models.ChartVersion{
		ChartVersion: "v1.2.3",
		ChartID:      chart.ID,
	}, user)
	s.NoError(err)
	s.True(created)
	chartVersionIdentifier := ciIdentifierModelFromOldModel(chartVersion)
	err = s.DB.Create(&chartVersionIdentifier).Error
	s.NoError(err)
	s.Equal(chartVersion.ID, chartVersionIdentifier.ResourceID)

	s.Run("chart version identifier", func() {
		var gotByID, gotByVersion CiIdentifierV3
		code := s.HandleRequest(
			s.NewRequest("GET", fmt.Sprintf("/api/ci-identifiers/v3/%d", chartVersionIdentifier.ID), nil),
			&gotByID)
		s.Equal(http.StatusOK, code)
		s.Equal(chartVersionIdentifier.ID, gotByID.ID)
		code = s.HandleRequest(
			s.NewRequest("GET", "/api/ci-identifiers/v3/chart-version/leonardo/v1.2.3", nil),
			&gotByVersion)
		s.Equal(http.StatusOK, code)
		s.Equal(chartVersionIdentifier.ID, gotByVersion.ID)
	})

	appVersion, created, err := v2models.InternalAppVersionStore.Create(s.DB, v2models.AppVersion{
		AppVersion: "v2.3.4",
		ChartID:    chart.ID,
	}, user)
	s.NoError(err)
	s.True(created)
	appVersionIdentifier := ciIdentifierModelFromOldModel(appVersion)
	err = s.DB.Create(&appVersionIdentifier).Error
	s.NoError(err)
	s.Equal(appVersion.ID, appVersionIdentifier.ResourceID)

	s.Run("app version identifier", func() {
		var gotByID, gotByVersion CiIdentifierV3
		code := s.HandleRequest(
			s.NewRequest("GET", fmt.Sprintf("/api/ci-identifiers/v3/%d", appVersionIdentifier.ID), nil),
			&gotByID)
		s.Equal(http.StatusOK, code)
		s.Equal(appVersionIdentifier.ID, gotByID.ID)
		code = s.HandleRequest(
			s.NewRequest("GET", "/api/ci-identifiers/v3/app-version/leonardo/v2.3.4", nil),
			&gotByVersion)
		s.Equal(http.StatusOK, code)
		s.Equal(appVersionIdentifier.ID, gotByVersion.ID)
	})

	cluster, created, err := v2models.InternalClusterStore.Create(s.DB, v2models.Cluster{
		Name:                "terra-dev",
		Provider:            "google",
		GoogleProject:       "broad-dsde-dev",
		Base:                testutils.PointerTo("live"),
		Address:             testutils.PointerTo("1.2.3.4"),
		RequiresSuitability: testutils.PointerTo(false),
		Location:            "us-central1-a",
		HelmfileRef:         testutils.PointerTo("HEAD"),
	}, user)
	s.NoError(err)
	s.True(created)
	clusterIdentifier := ciIdentifierModelFromOldModel(cluster)
	err = s.DB.Create(&clusterIdentifier).Error
	s.NoError(err)
	s.Equal(cluster.ID, clusterIdentifier.ResourceID)

	s.Run("cluster identifier", func() {
		var gotByID, gotByName CiIdentifierV3
		code := s.HandleRequest(
			s.NewRequest("GET", fmt.Sprintf("/api/ci-identifiers/v3/%d", clusterIdentifier.ID), nil),
			&gotByID)
		s.Equal(http.StatusOK, code)
		s.Equal(clusterIdentifier.ID, gotByID.ID)
		code = s.HandleRequest(
			s.NewRequest("GET", "/api/ci-identifiers/v3/cluster/terra-dev", nil),
			&gotByName)
		s.Equal(http.StatusOK, code)
		s.Equal(clusterIdentifier.ID, gotByName.ID)
	})

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
	environmentIdentifier := ciIdentifierModelFromOldModel(environment)
	err = s.DB.Create(&environmentIdentifier).Error
	s.NoError(err)
	s.Equal(environment.ID, environmentIdentifier.ResourceID)

	s.Run("environment identifier", func() {
		var gotByID, gotByName, gotByPrefixSelector CiIdentifierV3
		code := s.HandleRequest(
			s.NewRequest("GET", fmt.Sprintf("/api/ci-identifiers/v3/%d", environmentIdentifier.ID), nil),
			&gotByID)
		s.Equal(http.StatusOK, code)
		s.Equal(environmentIdentifier.ID, gotByID.ID)
		code = s.HandleRequest(
			s.NewRequest("GET", "/api/ci-identifiers/v3/environment/dev", nil),
			&gotByName)
		s.Equal(http.StatusOK, code)
		s.Equal(environmentIdentifier.ID, gotByName.ID)
		code = s.HandleRequest(
			s.NewRequest("GET", "/api/ci-identifiers/v3/environment/resource-prefix/a1b2", nil),
			&gotByPrefixSelector)
		s.Equal(http.StatusOK, code)
		s.Equal(environmentIdentifier.ID, gotByPrefixSelector.ID)
	})

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
	chartReleaseIdentifier := ciIdentifierModelFromOldModel(chartRelease)
	err = s.DB.Create(&chartReleaseIdentifier).Error
	s.NoError(err)
	s.Equal(chartRelease.ID, chartReleaseIdentifier.ResourceID)

	s.Run("chart release identifier", func() {
		var gotByID, gotByName, gotByEnvironment, gotByCluster CiIdentifierV3
		code := s.HandleRequest(
			s.NewRequest("GET", fmt.Sprintf("/api/ci-identifiers/v3/%d", chartReleaseIdentifier.ID), nil),
			&gotByID)
		s.Equal(http.StatusOK, code)
		s.Equal(chartReleaseIdentifier.ID, gotByID.ID)
		code = s.HandleRequest(
			s.NewRequest("GET", "/api/ci-identifiers/v3/chart-release/leonardo-dev", nil),
			&gotByName)
		s.Equal(http.StatusOK, code)
		s.Equal(chartReleaseIdentifier.ID, gotByName.ID)
		code = s.HandleRequest(
			s.NewRequest("GET", "/api/ci-identifiers/v3/chart-release/dev/leonardo", nil),
			&gotByEnvironment)
		s.Equal(http.StatusOK, code)
		s.Equal(chartReleaseIdentifier.ID, gotByEnvironment.ID)
		code = s.HandleRequest(
			s.NewRequest("GET", "/api/ci-identifiers/v3/chart-release/terra-dev/terra-dev/leonardo", nil),
			&gotByCluster)
		s.Equal(http.StatusOK, code)
		s.Equal(chartReleaseIdentifier.ID, gotByCluster.ID)
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
		},
	}, user)
	s.NoError(err)
	s.Len(controllerChangesets, 1)
	changeset, err := v2models.InternalChangesetStore.Get(s.DB, v2models.Changeset{Model: gorm.Model{ID: controllerChangesets[0].ID}})
	s.NoError(err)
	changesetIdentifier := ciIdentifierModelFromOldModel(changeset)
	err = s.DB.Create(&changesetIdentifier).Error
	s.NoError(err)
	s.Equal(changeset.ID, changesetIdentifier.ResourceID)

	s.Run("changeset identifier", func() {
		var gotByID, gotByChangesetID CiIdentifierV3
		code := s.HandleRequest(
			s.NewRequest("GET", fmt.Sprintf("/api/ci-identifiers/v3/%d", changesetIdentifier.ID), nil),
			&gotByID)
		s.Equal(http.StatusOK, code)
		s.Equal(changesetIdentifier.ID, gotByID.ID)
		code = s.HandleRequest(
			s.NewRequest("GET", fmt.Sprintf("/api/ci-identifiers/v3/changeset/%d", changeset.ID), nil),
			&gotByChangesetID)
		s.Equal(http.StatusOK, code)
		s.Equal(changesetIdentifier.ID, gotByChangesetID.ID)
	})
}

func (s *handlerSuite) TestCiIdentifiersV3GetLimitRuns() {
	user := s.SetSuitableTestUserForDB()

	_, created, err := v2models.InternalChartStore.Create(s.DB, v2models.Chart{
		Name:      "leonardo",
		ChartRepo: testutils.PointerTo("terra-helm"),
	}, user)
	s.NoError(err)
	s.True(created)

	totalIterations := uint(15)
	for iteration := uint(1); iteration <= totalIterations; iteration++ {
		var got CiRunV3
		code := s.HandleRequest(
			s.NewRequest("PUT", "/api/ci-runs/v3", CiRunV3Upsert{
				ciRunFields: ciRunFields{
					Platform:                   "github-actions",
					GithubActionsOwner:         "owner",
					GithubActionsRepo:          "repo",
					GithubActionsRunID:         iteration,
					GithubActionsAttemptNumber: 1,
					GithubActionsWorkflowPath:  "workflow",
					// Higher IDs started more recently, just for convenience in testing
					StartedAt: testutils.PointerTo(time.Now().Add(-time.Hour).Add(time.Minute * time.Duration(iteration))),
				},
				Charts: []string{"leonardo"},
			}),
			&got)
		s.Equalf(http.StatusCreated, code, "iteration %d", iteration)
	}

	var got CiIdentifierV3
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/ci-identifiers/v3/chart/leonardo", nil),
		&got)
	s.Equal(http.StatusOK, code)
	s.Len(got.CiRuns, 10) // default
	code = s.HandleRequest(
		s.NewRequest("GET", "/api/ci-identifiers/v3/chart/leonardo?limitCiRuns=20", nil),
		&got)
	s.Equal(http.StatusOK, code)
	s.Len(got.CiRuns, int(totalIterations))
	code = s.HandleRequest(
		s.NewRequest("GET", "/api/ci-identifiers/v3/chart/leonardo?limitCiRuns=5", nil),
		&got)
	s.Equal(http.StatusOK, code)
	s.Len(got.CiRuns, 5)
	s.Greater(*got.CiRuns[0].StartedAt, *got.CiRuns[4].StartedAt)
	s.Equal(got.CiRuns[0].GithubActionsRunID, totalIterations)
	s.Equal(got.CiRuns[1].GithubActionsRunID, totalIterations-1)
	s.Equal(got.CiRuns[2].GithubActionsRunID, totalIterations-2)
	s.Equal(got.CiRuns[3].GithubActionsRunID, totalIterations-3)
	s.Equal(got.CiRuns[4].GithubActionsRunID, totalIterations-4)
	code = s.HandleRequest(
		s.NewRequest("GET", "/api/ci-identifiers/v3/chart/leonardo?limitCiRuns=5&offsetCiRuns=5", nil),
		&got)
	s.Equal(http.StatusOK, code)
	s.Len(got.CiRuns, 5)
	s.Greater(*got.CiRuns[0].StartedAt, *got.CiRuns[4].StartedAt)
	s.Equal(got.CiRuns[0].GithubActionsRunID, totalIterations-5)
	s.Equal(got.CiRuns[1].GithubActionsRunID, totalIterations-6)
	s.Equal(got.CiRuns[2].GithubActionsRunID, totalIterations-7)
	s.Equal(got.CiRuns[3].GithubActionsRunID, totalIterations-8)
	s.Equal(got.CiRuns[4].GithubActionsRunID, totalIterations-9)
	code = s.HandleRequest(
		s.NewRequest("GET", "/api/ci-identifiers/v3/chart/leonardo?limitCiRuns=10&offsetCiRuns=10", nil),
		&got)
	s.Equal(http.StatusOK, code)
	s.Len(got.CiRuns, 5)
	s.Greater(*got.CiRuns[0].StartedAt, *got.CiRuns[4].StartedAt)
	s.Equal(got.CiRuns[0].GithubActionsRunID, totalIterations-10)
	s.Equal(got.CiRuns[1].GithubActionsRunID, totalIterations-11)
	s.Equal(got.CiRuns[2].GithubActionsRunID, totalIterations-12)
	s.Equal(got.CiRuns[3].GithubActionsRunID, totalIterations-13)
	s.Equal(got.CiRuns[4].GithubActionsRunID, totalIterations-14)
}
