package sherlock

import (
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *handlerSuite) TestChartsV3Edit_badSelector() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("PATCH", "/api/charts/v3/something/with/slashes", nil),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
}

func (s *handlerSuite) TestChartsV3Edit_badBody() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("PATCH", "/api/charts/v3/123", gin.H{
			"chartRepo": 123,
		}),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
	s.Contains(got.Message, "chartRepo")
}

func (s *handlerSuite) TestChartsV3Edit_notFound() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("PATCH", "/api/charts/v3/123", ChartV3Edit{
			ChartRepo: utils.PointerTo("terra-helm"),
		}),
		&got)
	s.Equal(http.StatusNotFound, code)
	s.Equal(errors.NotFound, got.Type)
}

func (s *handlerSuite) TestChartsV3Edit_sqlValidation() {
	edit := models.Chart{
		Name:      "name",
		ChartRepo: utils.PointerTo("terra-helm"),
	}
	s.NoError(s.DB.Create(&edit).Error)

	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("PATCH", fmt.Sprintf("/api/charts/v3/%d", edit.ID), ChartV3Edit{
			ChartRepo: utils.PointerTo(""),
		}),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
	s.Contains(got.Message, "chart_repo_present")
}

func (s *handlerSuite) TestChartsV3Edit() {
	edit := models.Chart{
		Name:      "name",
		ChartRepo: utils.PointerTo("terra-helm"),
	}
	s.NoError(s.DB.Create(&edit).Error)

	var got ChartV3
	code := s.HandleRequest(
		s.NewRequest("PATCH", fmt.Sprintf("/api/charts/v3/%d", edit.ID), ChartV3Edit{
			ChartRepo: utils.PointerTo("repo"),
		}),
		&got)
	s.Equal(http.StatusOK, code)
	if s.NotNil(got.ChartRepo) {
		s.Equal("repo", *got.ChartRepo)
	}
}
