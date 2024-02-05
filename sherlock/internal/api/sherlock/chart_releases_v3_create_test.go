package sherlock

import (
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *handlerSuite) TestChartReleasesV3Create_badBody() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("POST", "/api/chart-releases/v3", gin.H{
			"chart": 123,
		}),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
	s.Contains(got.Message, "chart")
}

func (s *handlerSuite) TestChartReleasesV3Create_failToConvertToModel() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("POST", "/api/chart-releases/v3", ChartReleaseV3Create{
			Chart: "not-found",
		}),
		&got)
	s.Equal(http.StatusNotFound, code)
	s.Equal(errors.NotFound, got.Type)
	s.Contains(got.Message, "not-found")
}

func (s *handlerSuite) TestChartReleasesV3Create_missingChart() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("POST", "/api/chart-releases/v3", ChartReleaseV3Create{}),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
	s.Contains(got.Message, "chart is required")
}

func (s *handlerSuite) TestChartReleasesV3Create_missingDestination() {
	s.TestData.Chart_Leonardo()
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("POST", "/api/chart-releases/v3", ChartReleaseV3Create{
			Chart: "leonardo",
		}),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
	s.Contains(got.Message, "environment or cluster is required")
}

func (s *handlerSuite) TestChartReleasesV3Create_cannotFindBranchAppVersion() {
	s.TestData.Chart_Leonardo()
	s.TestData.Environment_Dev()
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("POST", "/api/chart-releases/v3", ChartReleaseV3Create{
			Chart:       "leonardo",
			Environment: "dev",
		}),
		&got)
	s.Equal(http.StatusNotFound, code)
	s.Equal(errors.NotFound, got.Type)
	s.Contains(got.Message, "no recorded app versions for leonardo come from a '"+*s.TestData.Chart_Leonardo().AppImageGitMainBranch+"' branch")
}

func (s *handlerSuite) TestChartReleasesV3Create_cannotFindLatestChartVersion() {
	s.TestData.Chart_Leonardo()
	s.TestData.AppVersion_Leonardo_V1()
	s.TestData.Environment_Dev()
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("POST", "/api/chart-releases/v3", ChartReleaseV3Create{
			Chart:       "leonardo",
			Environment: "dev",
		}),
		&got)
	s.Equal(http.StatusNotFound, code)
	s.Equal(errors.NotFound, got.Type)
	s.Contains(got.Message, "unable to query latest chart version for leonardo")
}

func (s *handlerSuite) TestChartReleasesV3Create_normalCase() {
	s.TestData.Chart_Leonardo()
	s.TestData.AppVersion_Leonardo_V1()
	s.TestData.ChartVersion_Leonardo_V1()
	s.TestData.Environment_Dev()
	var got ChartReleaseV3
	code := s.HandleRequest(
		s.NewRequest("POST", "/api/chart-releases/v3", ChartReleaseV3Create{
			Chart:       "leonardo",
			Environment: "dev",
		}),
		&got)
	s.Equal(http.StatusCreated, code)
	s.NotEmpty(got.ID)
	s.Equal("leonardo-dev", got.Name)
}

func (s *handlerSuite) TestChartReleasesV3Create_normalCaseCluster() {
	s.TestData.Chart_Leonardo()
	s.TestData.AppVersion_Leonardo_V1()
	s.TestData.ChartVersion_Leonardo_V1()
	s.TestData.Environment_Dev()
	var got ChartReleaseV3
	code := s.HandleRequest(
		s.NewRequest("POST", "/api/chart-releases/v3", ChartReleaseV3Create{
			Chart:     "leonardo",
			Cluster:   "terra-dev",
			Namespace: "terra-dev",
		}),
		&got)
	s.Equal(http.StatusCreated, code)
	s.NotEmpty(got.ID)
	s.Equal("leonardo-terra-dev", got.Name)
}

func (s *handlerSuite) TestChartReleasesV3Create_normalCaseClusterWeirdNamespace() {
	s.TestData.Chart_Leonardo()
	s.TestData.AppVersion_Leonardo_V1()
	s.TestData.ChartVersion_Leonardo_V1()
	s.TestData.Environment_Dev()
	var got ChartReleaseV3
	code := s.HandleRequest(
		s.NewRequest("POST", "/api/chart-releases/v3", ChartReleaseV3Create{
			Chart:     "leonardo",
			Cluster:   "terra-dev",
			Namespace: "some-namespace",
		}),
		&got)
	s.Equal(http.StatusCreated, code)
	s.NotEmpty(got.ID)
	s.Equal("leonardo-some-namespace-terra-dev", got.Name)
}

func (s *handlerSuite) TestChartReleasesV3Create_noDefaults() {
	s.TestData.Chart_Leonardo()
	s.TestData.AppVersion_Leonardo_V1()
	s.TestData.ChartVersion_Leonardo_V1()
	s.TestData.Environment_Dev()
	var got ChartReleaseV3
	code := s.HandleRequest(
		s.NewRequest("POST", "/api/chart-releases/v3", ChartReleaseV3Create{
			Chart:                "leonardo",
			Cluster:              "terra-dev",
			Environment:          "dev",
			Name:                 "name",
			Namespace:            "namespace",
			AppVersionResolver:   utils.PointerTo("exact"),
			AppVersionExact:      utils.PointerTo(s.TestData.AppVersion_Leonardo_V1().AppVersion),
			ChartVersionResolver: utils.PointerTo("exact"),
			ChartVersionExact:    utils.PointerTo(s.TestData.ChartVersion_Leonardo_V1().ChartVersion),
			HelmfileRef:          utils.PointerTo("HEAD"),
			HelmfileRefEnabled:   utils.PointerTo(true),
			FirecloudDevelopRef:  utils.PointerTo("develop"),
			ChartReleaseV3Edit: ChartReleaseV3Edit{
				Subdomain:               utils.PointerTo("subdomain"),
				Protocol:                utils.PointerTo("protocol"),
				Port:                    utils.PointerTo[uint](123),
				IncludeInBulkChangesets: utils.PointerTo(false),
			},
		}),
		&got)
	s.Equal(http.StatusCreated, code)
	s.NotEmpty(got.ID)
	s.Equal("leonardo", got.Chart)
	s.Equal("terra-dev", got.Cluster)
	s.Equal("dev", got.Environment)
	s.Equal("name", got.Name)
	s.Equal("namespace", got.Namespace)
	s.Equal("exact", *got.AppVersionResolver)
	s.Equal(s.TestData.AppVersion_Leonardo_V1().AppVersion, *got.AppVersionExact)
	s.Equal("exact", *got.ChartVersionResolver)
	s.Equal(s.TestData.ChartVersion_Leonardo_V1().ChartVersion, *got.ChartVersionExact)
	s.Equal("HEAD", *got.HelmfileRef)
	s.Equal(true, *got.HelmfileRefEnabled)
	s.Equal("develop", *got.FirecloudDevelopRef)
	s.Equal("subdomain", *got.Subdomain)
	s.Equal("protocol", *got.Protocol)
	s.Equal(uint(123), *got.Port)
	s.Equal(false, *got.IncludeInBulkChangesets)
}

func (s *handlerSuite) TestChartReleasesV3Create_blocksDuplicates() {
	s.TestData.Chart_Leonardo()
	s.TestData.AppVersion_Leonardo_V1()
	s.TestData.ChartVersion_Leonardo_V1()
	s.TestData.Environment_Dev()
	var got ChartReleaseV3
	code := s.HandleRequest(
		s.NewRequest("POST", "/api/chart-releases/v3", ChartReleaseV3Create{
			Chart:       "leonardo",
			Environment: "dev",
		}),
		&got)
	s.Equal(http.StatusCreated, code)
	s.NotEmpty(got.ID)
	s.Equal("leonardo-dev", got.Name)
	var got2 errors.ErrorResponse
	code = s.HandleRequest(
		s.NewRequest("POST", "/api/chart-releases/v3", ChartReleaseV3Create{
			Chart:       "leonardo",
			Environment: "dev",
		}),
		&got2)
	s.Equal(http.StatusConflict, code)
	s.Equal(errors.Conflict, got2.Type)
}

func (s *handlerSuite) TestChartReleasesV3Create_suitability() {
	s.TestData.Chart_Leonardo()
	s.TestData.AppVersion_Leonardo_V1()
	s.TestData.ChartVersion_Leonardo_V1()
	s.TestData.Environment_Prod()
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.UseNonSuitableUserFor(s.NewRequest("POST", "/api/chart-releases/v3", ChartReleaseV3Create{
			Chart:       "leonardo",
			Environment: "prod",
		})),
		&got)
	s.Equal(http.StatusForbidden, code)
	s.Equal(errors.Forbidden, got.Type)
}

func (s *handlerSuite) TestChartReleasesV3Create_suitabilityAllowed() {
	s.TestData.Chart_Leonardo()
	s.TestData.AppVersion_Leonardo_V1()
	s.TestData.ChartVersion_Leonardo_V1()
	s.TestData.Environment_Prod()
	var got ChartReleaseV3
	code := s.HandleRequest(
		s.UseSuitableUserFor(s.NewRequest("POST", "/api/chart-releases/v3", ChartReleaseV3Create{
			Chart:       "leonardo",
			Environment: "prod",
		})),
		&got)
	s.Equal(http.StatusCreated, code)
	s.NotEmpty(got.ID)
}
