package sherlock

import (
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *handlerSuite) TestDatabaseInstancesV3Create_badBody() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("POST", "/api/database-instances/v3", gin.H{
			"platform": 123,
		}),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
	s.Contains(got.Message, "platform")
}

func (s *handlerSuite) TestDatabaseInstancesV3Create() {
	var got DatabaseInstanceV3
	code := s.HandleRequest(
		s.NewRequest("POST", "/api/database-instances/v3", DatabaseInstanceV3Create{
			ChartRelease: s.TestData.ChartRelease_LeonardoSwatomation().Name,
		}),
		&got)
	s.Equal(http.StatusCreated, code)
	if s.NotNil(got.DefaultDatabase) {
		s.Equal("leonardo", *got.DefaultDatabase)
	}
}

func (s *handlerSuite) TestDatabaseInstancesV3Create_overrideDefault() {
	var got DatabaseInstanceV3
	code := s.HandleRequest(
		s.NewRequest("POST", "/api/database-instances/v3", DatabaseInstanceV3Create{
			ChartRelease: s.TestData.ChartRelease_LeonardoSwatomation().Name,
			DatabaseInstanceV3Edit: DatabaseInstanceV3Edit{
				Platform:        utils.PointerTo("google"),
				GoogleProject:   utils.PointerTo("some-project"),
				InstanceName:    utils.PointerTo("some-instance"),
				DefaultDatabase: utils.PointerTo("stairway or something idk"),
			},
		}),
		&got)
	s.Equal(http.StatusCreated, code)
	if s.NotNil(got.Platform) {
		s.Equal("google", *got.Platform)
	}
	if s.NotNil(got.DefaultDatabase) {
		s.Equal("stairway or something idk", *got.DefaultDatabase)
	}
}

func (s *handlerSuite) TestDatabaseInstancesV3Create_sqlValidation() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("POST", "/api/database-instances/v3", DatabaseInstanceV3Create{
			ChartRelease: s.TestData.ChartRelease_LeonardoSwatomation().Name,
			DatabaseInstanceV3Edit: DatabaseInstanceV3Edit{
				Platform: utils.PointerTo("blah"),
			},
		}),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Contains(got.Message, "platform")
}

func (s *handlerSuite) TestDatabaseInstancesV3Create_requiresChartRelease() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("POST", "/api/database-instances/v3", DatabaseInstanceV3Create{
			DatabaseInstanceV3Edit: DatabaseInstanceV3Edit{
				Platform:        utils.PointerTo("google"),
				GoogleProject:   utils.PointerTo("some-project"),
				InstanceName:    utils.PointerTo("some-instance"),
				DefaultDatabase: utils.PointerTo("stairway or something idk"),
			},
		}),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Contains(got.Message, "chartRelease")
}

func (s *handlerSuite) TestDatabaseInstancesV3Create_chartReleaseNotFound() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("POST", "/api/database-instances/v3", DatabaseInstanceV3Create{
			ChartRelease: "not-found",
		}),
		&got)
	s.Equal(http.StatusNotFound, code)
	s.Contains(got.Message, "not-found")
}

func (s *handlerSuite) TestDatabaseInstancesV3Create_suitability() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.UseNonSuitableUserFor(s.NewRequest("POST", "/api/database-instances/v3", DatabaseInstanceV3Create{
			ChartRelease: s.TestData.ChartRelease_LeonardoProd().Name,
		})),
		&got)
	s.Equal(http.StatusForbidden, code)
}

func (s *handlerSuite) TestDatabaseInstancesV3Create_suitabilityAllowed() {
	var got DatabaseInstanceV3
	code := s.HandleRequest(
		s.UseSuitableUserFor(s.NewRequest("POST", "/api/database-instances/v3", DatabaseInstanceV3Create{
			ChartRelease: s.TestData.ChartRelease_LeonardoProd().Name,
		})),
		&got)
	s.Equal(http.StatusCreated, code)
}
