package sherlock

import (
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *handlerSuite) TestEnvironmentsV3Edit_badSelector() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("PATCH", "/api/environments/v3/!!!!!", nil),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
}

func (s *handlerSuite) TestEnvironmentsV3Edit_badBody() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("PATCH", "/api/environments/v3/123", gin.H{
			"helmfileRef": 123,
		}),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
	s.Contains(got.Message, "helmfileRef")
}

func (s *handlerSuite) TestEnvironmentsV3Edit_notFound() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("PATCH", "/api/environments/v3/123", EnvironmentV3Edit{
			HelmfileRef: utils.PointerTo("some-ref"),
		}),
		&got)
	s.Equal(http.StatusNotFound, code)
	s.Equal(errors.NotFound, got.Type)
}

func (s *handlerSuite) TestEnvironmentsV3Edit_sqlValidation() {
	edit := s.TestData.Environment_Dev()
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("PATCH", fmt.Sprintf("/api/environments/v3/%d", edit.ID), EnvironmentV3Edit{
			HelmfileRef: utils.PointerTo(""),
		}),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
	s.Contains(got.Message, "helmfile")
}

func (s *handlerSuite) TestEnvironmentsV3Edit() {
	edit := s.TestData.Environment_Dev()
	var got EnvironmentV3
	code := s.HandleRequest(
		s.NewRequest("PATCH", fmt.Sprintf("/api/environments/v3/%d", edit.ID), EnvironmentV3Edit{
			HelmfileRef: utils.PointerTo("some-ref"),
		}),
		&got)
	s.Equal(http.StatusOK, code)
	if s.NotNil(got.HelmfileRef) {
		s.Equal("some-ref", *got.HelmfileRef)
	}
}

func (s *handlerSuite) TestEnvironmentsV3Edit_suitability() {
	edit := s.TestData.Environment_Prod()
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.UseNonSuitableUserFor(s.NewRequest("PATCH", fmt.Sprintf("/api/environments/v3/%d", edit.ID), EnvironmentV3Edit{
			HelmfileRef: utils.PointerTo("some-ref"),
		})),
		&got)
	s.Equal(http.StatusForbidden, code)
	s.Equal(errors.Forbidden, got.Type)
}

func (s *handlerSuite) TestEnvironmentsV3Edit_suitabilityAllowed() {
	edit := s.TestData.Environment_Prod()
	var got EnvironmentV3
	code := s.HandleRequest(
		s.UseSuitableUserFor(s.NewRequest("PATCH", fmt.Sprintf("/api/environments/v3/%d", edit.ID), EnvironmentV3Edit{
			HelmfileRef: utils.PointerTo("some-ref"),
		})),
		&got)
	s.Equal(http.StatusOK, code)
	if s.NotNil(got.HelmfileRef) {
		s.Equal("some-ref", *got.HelmfileRef)
	}
}

func (s *handlerSuite) TestEnvironmentsV3Edit_suitabilityBefore() {
	edit := s.TestData.Environment_Prod()
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.UseNonSuitableUserFor(s.NewRequest("PATCH", fmt.Sprintf("/api/environments/v3/%d", edit.ID), EnvironmentV3Edit{
			RequiresSuitability: utils.PointerTo(false),
		})),
		&got)
	s.Equal(http.StatusForbidden, code)
	s.Equal(errors.Forbidden, got.Type)
}

func (s *handlerSuite) TestEnvironmentsV3Edit_suitabilityAfter() {
	edit := s.TestData.Environment_Dev()
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.UseNonSuitableUserFor(s.NewRequest("PATCH", fmt.Sprintf("/api/environments/v3/%d", edit.ID), EnvironmentV3Edit{
			RequiresSuitability: utils.PointerTo(true),
		})),
		&got)
	s.Equal(http.StatusForbidden, code)
	s.Equal(errors.Forbidden, got.Type)
}
