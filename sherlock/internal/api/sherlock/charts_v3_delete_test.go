package sherlock

import (
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"net/http"
)

func (s *handlerSuite) TestChartV3Delete_badSelector() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("DELETE", "/api/charts/v3/something/with/slashes", nil),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
}

func (s *handlerSuite) TestChartV3Delete_notFound() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("DELETE", "/api/charts/v3/my-chart", nil),
		&got)
	s.Equal(http.StatusNotFound, code)
	s.Equal(errors.NotFound, got.Type)
}

func (s *handlerSuite) TestChartV3Delete() {
	s.NoError(s.DB.Create(&models.Chart{Name: "my-chart", ChartRepo: utils.PointerTo("some-repo")}).Error)

	var got ChartV3
	code := s.HandleRequest(
		s.NewRequest("DELETE", "/api/charts/v3/my-chart", nil),
		&got)
	s.Equal(http.StatusOK, code)
	if s.NotNil(got.ChartRepo) {
		s.Equal("some-repo", *got.ChartRepo)
	}
}
