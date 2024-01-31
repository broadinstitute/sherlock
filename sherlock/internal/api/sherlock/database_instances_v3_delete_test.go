package sherlock

import (
	"fmt"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"net/http"
)

func (s *handlerSuite) TestDatabaseInstancesV3Delete_badSelector() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("DELETE", "/api/database-instances/v3/!!!!!", nil),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
}

func (s *handlerSuite) TestDatabaseInstancesV3Delete_notFound() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("DELETE", "/api/database-instances/v3/chart-release/not-found", nil),
		&got)
	s.Equal(http.StatusNotFound, code)
	s.Equal(errors.NotFound, got.Type)
}

func (s *handlerSuite) TestDatabaseInstancesV3Delete() {
	di := s.TestData.DatabaseInstance_LeonardoDev()
	var got DatabaseInstanceV3
	code := s.HandleRequest(
		s.NewRequest("DELETE", fmt.Sprintf("/api/database-instances/v3/%d", di.ID), nil),
		&got)
	s.Equal(http.StatusOK, code)
	s.Equal(di.ID, got.ID)
}

func (s *handlerSuite) TestDatabaseInstancesV3Delete_suitability() {
	s.TestData.DatabaseInstance_LeonardoProd()
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.UseNonSuitableUserFor(s.NewRequest("DELETE", "/api/database-instances/v3/chart-release/leonardo-prod", nil)),
		&got)
	s.Equal(http.StatusForbidden, code)
	s.Equal(errors.Forbidden, got.Type)
}

func (s *handlerSuite) TestDatabaseInstancesV3Delete_suitabilityAllowed() {
	di := s.TestData.DatabaseInstance_LeonardoProd()
	var got DatabaseInstanceV3
	code := s.HandleRequest(
		s.UseSuitableUserFor(s.NewRequest("DELETE", "/api/database-instances/v3/chart-release/leonardo-prod", nil)),
		&got)
	s.Equal(http.StatusOK, code)
	s.Equal(di.ID, got.ID)
}
