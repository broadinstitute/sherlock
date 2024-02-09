package sherlock

import (
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *handlerSuite) TestDatabaseInstancesV3Upsert_badBody() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("PUT", "/api/database-instances/v3", gin.H{
			"platform": 123,
		}),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
	s.Contains(got.Message, "platform")
}

func (s *handlerSuite) TestDatabaseInstancesV3Upsert_requiresChartRelease() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("PUT", "/api/database-instances/v3", DatabaseInstanceV3Create{
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

func (s *handlerSuite) TestDatabaseInstancesV3Upsert_chartReleaseNotFound() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("PUT", "/api/database-instances/v3", DatabaseInstanceV3Create{
			ChartRelease: "not-found",
		}),
		&got)
	s.Equal(http.StatusNotFound, code)
	s.Contains(got.Message, "not-found")
}

func (s *handlerSuite) TestDatabaseInstancesV3Upsert_edits() {
	di := s.TestData.DatabaseInstance_LeonardoDev()
	var got DatabaseInstanceV3
	code := s.HandleRequest(
		s.NewRequest("PUT", "/api/database-instances/v3", DatabaseInstanceV3Create{
			ChartRelease: s.TestData.ChartRelease_LeonardoDev().Name,
			DatabaseInstanceV3Edit: DatabaseInstanceV3Edit{
				DefaultDatabase: utils.PointerTo("foo"),
			}}),
		&got)
	s.Equal(http.StatusOK, code)
	s.Equal("foo", *got.DefaultDatabase)
	s.Equal(di.ID, got.ID)
}

func (s *handlerSuite) TestDatabaseInstancesV3Upsert_creates() {
	var got DatabaseInstanceV3
	code := s.HandleRequest(
		s.NewRequest("PUT", "/api/database-instances/v3", DatabaseInstanceV3Create{
			ChartRelease: s.TestData.ChartRelease_LeonardoDev().Name,
		}),
		&got)
	s.Equal(http.StatusCreated, code)
	if s.NotNil(got.DefaultDatabase) {
		s.Equal("leonardo", *got.DefaultDatabase)
	}
}
