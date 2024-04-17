package sherlock

import (
	"net/http"

	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/gin-gonic/gin"
)

func (s *handlerSuite) TestEnvironmentsProceduresV3UpsertBee_bee() {
	template := s.TestData.Environment_Swatomation()
	chartReleaseToUpdate := s.TestData.ChartRelease_LeonardoProd() // fix this
	var got EnvironmentV3
	code := s.HandleRequest(
		s.NewRequest("PUT", "/api/environments/procedures/v3/upsert-bee", EnvironmentProceduresV3UpsertBee{
			TemplateEnvironment: template.Name,
			ChartReleases: []ChangesetV3PlanRequestChartReleaseEntry{
				{
					ChangesetV3Create: ChangesetV3Create{
						ChartRelease:         chartReleaseToUpdate.Name,
						ToAppVersionResolver: utils.PointerTo("exact"),
						ToAppVersionExact:    utils.PointerTo("some new version"),
					},
				},
			},
		}),
		&got)
	s.Equal(http.StatusCreated, code)
	s.NotZero(got.ID)

	s.Run("copied chart releases", func() {
		var chartReleasesInTemplate []models.ChartRelease
		s.NoError(s.DB.Where("environment_id = ?", template.ID).Find(&chartReleasesInTemplate).Error)
		s.NotZero(len(chartReleasesInTemplate))
		var chartReleasesInNewEnvironment []models.ChartRelease
		s.NoError(s.DB.Where("environment_id = ?", got.ID).Find(&chartReleasesInNewEnvironment).Error)
		s.Equal(len(chartReleasesInTemplate), len(chartReleasesInNewEnvironment))
	})
}

func (s *handlerSuite) TestEnvironmentsProceduresV3UpsertBee_beeDuplicateOk() {
	template := s.TestData.Environment_Swatomation()
	chartReleaseToUpdate := s.TestData.ChartRelease_LeonardoProd() // fix this
	var got EnvironmentV3
	code := s.HandleRequest(
		s.NewRequest("PUT", "/api/environments/procedures/v3/upsert-bee", EnvironmentProceduresV3UpsertBee{
			TemplateEnvironment: template.Name,
			ChartReleases: []ChangesetV3PlanRequestChartReleaseEntry{
				{
					ChangesetV3Create: ChangesetV3Create{
						ChartRelease:         chartReleaseToUpdate.Name,
						ToAppVersionResolver: utils.PointerTo("exact"),
						ToAppVersionExact:    utils.PointerTo("some new version"),
					},
				},
			},
		}),
		&got)
	s.Equal(http.StatusCreated, code)
	s.NotZero(got.ID)

	s.Run("copied chart releases", func() {
		var chartReleasesInTemplate []models.ChartRelease
		s.NoError(s.DB.Where("environment_id = ?", template.ID).Find(&chartReleasesInTemplate).Error)
		s.NotZero(len(chartReleasesInTemplate))
		var chartReleasesInNewEnvironment []models.ChartRelease
		s.NoError(s.DB.Where("environment_id = ?", got.ID).Find(&chartReleasesInNewEnvironment).Error)
		s.Equal(len(chartReleasesInTemplate), len(chartReleasesInNewEnvironment))
	})

	// run the request again with the last name, should return the same ID
	var got2 EnvironmentV3
	code = s.HandleRequest(
		s.NewRequest("PUT", "/api/environments/procedures/v3/upsert-bee", EnvironmentProceduresV3UpsertBee{
			TemplateEnvironment: template.Name,
			Name:                got.Name,
			ChartReleases: []ChangesetV3PlanRequestChartReleaseEntry{
				{
					ChangesetV3Create: ChangesetV3Create{
						ChartRelease:         chartReleaseToUpdate.Name,
						ToAppVersionResolver: utils.PointerTo("exact"),
						ToAppVersionExact:    utils.PointerTo("some new version"),
					},
				},
			},
		}),
		&got2)
	s.Equal(http.StatusCreated, code)
	s.Equal(got.ID, got2.ID)
}

func (s *handlerSuite) TestEnvironmentsProceduresV3UpsertBee_badBody() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("PUT", "/api/environments/procedures/v3/upsert-bee", gin.H{
			"name": 123,
		}),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
	s.Contains(got.Message, "name")
}

func (s *handlerSuite) TestEnvironmentsProceduresV3UpsertBee_failToConvertToModel() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("PUT", "/api/environments/procedures/v3/upsert-bee", EnvironmentProceduresV3UpsertBee{
			EnvironmentV3Edit: EnvironmentV3Edit{
				DefaultCluster: utils.PointerTo("not-found"),
			},
		}),
		&got)
	s.Equal(http.StatusNotFound, code)
	s.Equal(errors.NotFound, got.Type)
	s.Contains(got.Message, "not-found")
}

func (s *handlerSuite) TestEnvironmentProceduresV3UpsertBee_template() {
	s.TestData.Chart_Honeycomb()
	s.TestData.ChartVersion_Honeycomb_V1()
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("PUT", "/api/environments/procedures/v3/upsert-bee", EnvironmentProceduresV3UpsertBee{
			Base:      "swatomation",
			Lifecycle: "template",
			Name:      "tempy-temp",
		}),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
	s.Contains(got.Message, "Lifecycle is not \"dynamic\"")
}

func (s *handlerSuite) TestEnvironmentProceduresV3UpsertBee_staticFail() {
	s.TestData.Environment_Swatomation()
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("PUT", "/api/environments/procedures/v3/upsert-bee", EnvironmentProceduresV3UpsertBee{
			Base:      "live",
			Lifecycle: "static",
			Name:      "staticy-static",
		}),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
	s.Contains(got.Message, "Lifecycle is not \"dynamic\"")
}

func (s *handlerSuite) TestEnvironmentsProceduresV3UpsertBee_suitability() {
	s.TestData.Chart_Honeycomb()
	s.TestData.ChartVersion_Honeycomb_V1()
	template := s.TestData.Environment_Swatomation()
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.UseNonSuitableUserFor(s.NewRequest("PUT", "/api/environments/procedures/v3/upsert-bee", EnvironmentProceduresV3UpsertBee{
			TemplateEnvironment: template.Name,
			Base:                "swatomation",
			Lifecycle:           "dynamic",
			Name:                "beey-bee",
			EnvironmentV3Edit: EnvironmentV3Edit{
				RequiresSuitability: utils.PointerTo(true),
			},
		})),
		&got)
	s.Equal(http.StatusForbidden, code)
	s.Equal(errors.Forbidden, got.Type)
}

func (s *handlerSuite) TestEnvironmentsProceduresV3UpsertBee_suitabilityAllowed() {
	s.TestData.Chart_Honeycomb()
	s.TestData.ChartVersion_Honeycomb_V1()
	template := s.TestData.Environment_Swatomation()
	var got EnvironmentV3
	code := s.HandleRequest(
		s.UseSuitableUserFor(s.NewRequest("PUT", "/api/environments/procedures/v3/upsert-bee", EnvironmentProceduresV3UpsertBee{
			TemplateEnvironment: template.Name,
			Base:                "swatomation",
			Lifecycle:           "dynamic",
			Name:                "beey-bee",
			EnvironmentV3Edit: EnvironmentV3Edit{
				RequiresSuitability: utils.PointerTo(true),
			},
		})),
		&got)
	s.Equal(http.StatusCreated, code)
	s.NotZero(got.ID)
}
