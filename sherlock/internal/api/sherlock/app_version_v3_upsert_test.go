package sherlock

import (
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *handlerSuite) TestAppVersionsV3Upsert_badBody() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("PUT", "/api/app-versions/v3", gin.H{
			"Description": 123,
		}),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
	s.Contains(got.Message, "description")
}

func (s *handlerSuite) TestAppVersionsV3Upsert_sqlValidation() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("PUT", "/api/app-versions/v3", AppVersionV3Create{}),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
	s.Contains(got.Message, "app_version_present")
}

func (s *handlerSuite) TestAppVersionsV3Upsert_error() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("PUT", "/api/app-versions/v3", AppVersionV3Create{}),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
}

func (s *handlerSuite) TestAppVersionsV3Upsert() {
	var got ChartV3
	code := s.HandleRequest(
		s.NewRequest("POST", "/api/charts/v3", ChartV3Create{
			Name: "chart-name",
		}),
		&got)
	s.Equal(http.StatusCreated, code)
	s.Equal("chart-name", got.Name)
	if s.NotNil(got.ChartRepo) {
		s.Equal("terra-helm", *got.ChartRepo)
	}
	if s.NotNil(got.DefaultSubdomain) {
		s.Equal("chart-name", *got.DefaultSubdomain)
	}
	// check that i can edit description, specify git branch, but it doesn't actually edit
}
