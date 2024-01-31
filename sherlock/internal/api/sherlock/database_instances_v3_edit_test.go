package sherlock

import (
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *handlerSuite) TestDatabaseInstancesV3Edit_badSelector() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("PATCH", "/api/database-instances/v3/!!!!!", nil),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
}

func (s *handlerSuite) TestDatabaseInstancesV3Edit_badBody() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("PATCH", "/api/database-instances/v3/123", gin.H{
			"platform": 123,
		}),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
	s.Contains(got.Message, "platform")
}

func (s *handlerSuite) TestDatabaseInstancesV3Edit_notFound() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("PATCH", "/api/database-instances/v3/chart-release/not-found", DatabaseInstanceV3Edit{
			Platform: utils.PointerTo("some-platform"),
		}),
		&got)
	s.Equal(http.StatusNotFound, code)
	s.Equal(errors.NotFound, got.Type)
}

func (s *handlerSuite) TestDatabaseInstancesV3Edit() {
	edit := s.TestData.DatabaseInstance_LeonardoDev()
	var got DatabaseInstanceV3
	code := s.HandleRequest(
		s.NewRequest("PATCH", fmt.Sprintf("/api/database-instances/v3/%d", edit.ID), DatabaseInstanceV3Edit{
			GoogleProject: utils.PointerTo("some-project"),
		}),
		&got)
	s.Equal(http.StatusOK, code)
	if s.NotNil(got.GoogleProject) {
		s.Equal("some-project", *got.GoogleProject)
	}
}

func (s *handlerSuite) TestDatabaseInstancesV3Edit_sqlValidation() {
	edit := s.TestData.DatabaseInstance_LeonardoDev()
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("PATCH", fmt.Sprintf("/api/database-instances/v3/%d", edit.ID), DatabaseInstanceV3Edit{
			Platform: utils.PointerTo("something super invalid"),
		}),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
}

func (s *handlerSuite) TestDatabaseInstancesV3Edit_suitability() {
	edit := s.TestData.DatabaseInstance_LeonardoProd()
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.UseNonSuitableUserFor(s.NewRequest("PATCH", fmt.Sprintf("/api/database-instances/v3/%d", edit.ID), DatabaseInstanceV3Edit{
			GoogleProject: utils.PointerTo("some-project"),
		})),
		&got)
	s.Equal(http.StatusForbidden, code)
	s.Equal(errors.Forbidden, got.Type)
}

func (s *handlerSuite) TestDatabaseInstancesV3Edit_suitabilityAllowed() {
	edit := s.TestData.DatabaseInstance_LeonardoProd()
	var got DatabaseInstanceV3
	code := s.HandleRequest(
		s.UseSuitableUserFor(s.NewRequest("PATCH", fmt.Sprintf("/api/database-instances/v3/%d", edit.ID), DatabaseInstanceV3Edit{
			GoogleProject: utils.PointerTo("some-project"),
		})),
		&got)
	s.Equal(http.StatusOK, code)
	if s.NotNil(got.GoogleProject) {
		s.Equal("some-project", *got.GoogleProject)
	}
}
