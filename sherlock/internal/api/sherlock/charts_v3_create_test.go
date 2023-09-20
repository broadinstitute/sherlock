package sherlock

import (
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *handlerSuite) TestChartsV3Create_badBody() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("POST", "/api/charts/v3", gin.H{
			"name": 123,
		}),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
	s.Contains(got.Message, "name")
}

func (s *handlerSuite) TestChartsV3Create_sqlValidation() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("POST", "/api/charts/v3", ChartV3Create{}),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
	s.Contains(got.Message, "name")
}

func (s *handlerSuite) TestChartsV3Create_defaults() {
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
}

func (s *handlerSuite) TestChartsV3Create_overrideDefaults() {
	var got ChartV3
	code := s.HandleRequest(
		s.NewRequest("POST", "/api/charts/v3", ChartV3Create{
			Name: "chart-name",
			ChartV3Edit: ChartV3Edit{
				ChartRepo:        utils.PointerTo("different-chart-repo"),
				DefaultSubdomain: utils.PointerTo("different-subdomain"),
			},
		}),
		&got)
	s.Equal(http.StatusCreated, code)
	s.Equal("chart-name", got.Name)
	if s.NotNil(got.ChartRepo) {
		s.Equal("different-chart-repo", *got.ChartRepo)
	}
	if s.NotNil(got.DefaultSubdomain) {
		s.Equal("different-subdomain", *got.DefaultSubdomain)
	}
}
