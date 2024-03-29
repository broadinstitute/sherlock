package sherlock

import (
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"net/http"
	"time"
)

func (s *handlerSuite) TestCiIdentifiersV3Get_badSelector() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/ci-identifiers/v3/foo-bar", nil),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
}

func (s *handlerSuite) TestCiIdentifiersV3Get_notFound() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/ci-identifiers/v3/0", nil),
		&got)
	s.Equal(http.StatusNotFound, code)
	s.Equal(errors.NotFound, got.Type)
}

func (s *handlerSuite) TestCiIdentifiersV3Get() {
	chart := s.TestData.Chart_Leonardo()
	chartIdentifier := chart.GetCiIdentifier()
	err := s.DB.Create(&chartIdentifier).Error
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

	chartVersion := s.TestData.ChartVersion_Leonardo_V1()
	chartVersionIdentifier := chartVersion.GetCiIdentifier()
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
			s.NewRequest("GET", "/api/ci-identifiers/v3/chart-version/leonardo/"+chartVersion.ChartVersion, nil),
			&gotByVersion)
		s.Equal(http.StatusOK, code)
		s.Equal(chartVersionIdentifier.ID, gotByVersion.ID)
	})

	appVersion := s.TestData.AppVersion_Leonardo_V1()
	appVersionIdentifier := appVersion.GetCiIdentifier()
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
			s.NewRequest("GET", "/api/ci-identifiers/v3/app-version/leonardo/"+appVersion.AppVersion, nil),
			&gotByVersion)
		s.Equal(http.StatusOK, code)
		s.Equal(appVersionIdentifier.ID, gotByVersion.ID)
	})

	cluster := s.TestData.Cluster_TerraDev()
	clusterIdentifier := cluster.GetCiIdentifier()
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

	environment := s.TestData.Environment_Dev()
	environmentIdentifier := environment.GetCiIdentifier()
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
			s.NewRequest("GET", "/api/ci-identifiers/v3/environment/resource-prefix/"+environment.UniqueResourcePrefix, nil),
			&gotByPrefixSelector)
		s.Equal(http.StatusOK, code)
		s.Equal(environmentIdentifier.ID, gotByPrefixSelector.ID)
	})

	chartRelease := s.TestData.ChartRelease_LeonardoDev()
	chartReleaseIdentifier := chartRelease.GetCiIdentifier()
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

	changeset := s.TestData.Changeset_LeonardoDev_V1toV3()
	changesetIdentifier := changeset.GetCiIdentifier()
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
	s.TestData.Chart_Leonardo()

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
					StartedAt: utils.PointerTo(time.Now().Add(-time.Hour).Add(time.Minute * time.Duration(iteration))),
					Status:    utils.PointerTo("in progress"),
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

func (s *handlerSuite) TestCiIdentifiersV3Get_ResourceStatus() {
	s.TestData.CiRun_Deploy_LeonardoDev_V1toV3()
	var got CiIdentifierV3
	code := s.HandleRequest(
		s.NewRequest(http.MethodGet, "/api/ci-identifiers/v3/chart-release/dev/leonardo", nil),
		&got)
	s.Equal(http.StatusOK, code)
	s.NotEmpty(got.CiRuns)
	for _, cr := range got.CiRuns {
		s.NotNil(cr.ResourceStatus)
	}
}

func (s *handlerSuite) TestCiIdentifiersV3Get_allowStubCiRuns() {
	s.TestData.CiRun_Deploy_LeonardoDev_V1toV3()
	// Stub doesn't have status or started_at info
	s.TestData.CiRun_Stub_LeonardoDev()
	var got CiIdentifierV3
	code := s.HandleRequest(
		s.NewRequest(http.MethodGet, "/api/ci-identifiers/v3/chart-release/dev/leonardo?allowStubCiRuns=true", nil),
		&got)
	s.Equal(http.StatusOK, code)
	s.Len(got.CiRuns, 2)
	code = s.HandleRequest(
		s.NewRequest(http.MethodGet, "/api/ci-identifiers/v3/chart-release/dev/leonardo?allowStubCiRuns=false", nil),
		&got)
	s.Equal(http.StatusOK, code)
	s.Len(got.CiRuns, 1)
	code = s.HandleRequest(
		s.NewRequest(http.MethodGet, "/api/ci-identifiers/v3/chart-release/dev/leonardo", nil),
		&got)
	s.Equal(http.StatusOK, code)
	s.Len(got.CiRuns, 1)
}
