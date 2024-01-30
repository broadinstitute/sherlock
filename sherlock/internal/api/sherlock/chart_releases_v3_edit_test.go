package sherlock

import (
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *handlerSuite) TestChartReleasesV3Edit_badSelector() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("PATCH", "/api/chart-releases/v3/!!!!!", nil),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
}

func (s *handlerSuite) TestChartReleasesV3Edit_badBody() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("PATCH", "/api/chart-releases/v3/123", gin.H{
			"subdomain": 123,
		}),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
	s.Contains(got.Message, "subdomain")
}

func (s *handlerSuite) TestChartReleasesV3Edit_notFound() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("PATCH", "/api/chart-releases/v3/123", ChartReleaseV3Edit{
			Subdomain: utils.PointerTo("some-subdomain"),
		}),
		&got)
	s.Equal(http.StatusNotFound, code)
	s.Equal(errors.NotFound, got.Type)
}

func (s *handlerSuite) TestChartReleasesV3Edit() {
	edit := s.TestData.ChartRelease_LeonardoDev()
	var got ChartReleaseV3
	code := s.HandleRequest(
		s.NewRequest("PATCH", fmt.Sprintf("/api/chart-releases/v3/%d", edit.ID), ChartReleaseV3Edit{
			Subdomain: utils.PointerTo("some-subdomain"),
		}),
		&got)
	s.Equal(http.StatusOK, code)
	if s.NotNil(got.Subdomain) {
		s.Equal("some-subdomain", *got.Subdomain)
	}
}

func (s *handlerSuite) TestChartReleasesV3Edit_suitability() {
	edit := s.TestData.ChartRelease_LeonardoProd()
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.UseNonSuitableUserFor(s.NewRequest("PATCH", fmt.Sprintf("/api/chart-releases/v3/%d", edit.ID), ChartReleaseV3Edit{
			Subdomain: utils.PointerTo("some-subdomain"),
		})),
		&got)
	s.Equal(http.StatusForbidden, code)
	s.Equal(errors.Forbidden, got.Type)
}

func (s *handlerSuite) TestChartReleasesV3Edit_suitabilityAllowed() {
	edit := s.TestData.ChartRelease_LeonardoProd()
	var got ChartReleaseV3
	code := s.HandleRequest(
		s.UseSuitableUserFor(s.NewRequest("PATCH", fmt.Sprintf("/api/chart-releases/v3/%d", edit.ID), ChartReleaseV3Edit{
			Subdomain: utils.PointerTo("some-subdomain"),
		})),
		&got)
	s.Equal(http.StatusOK, code)
	if s.NotNil(got.Subdomain) {
		s.Equal("some-subdomain", *got.Subdomain)
	}
}
