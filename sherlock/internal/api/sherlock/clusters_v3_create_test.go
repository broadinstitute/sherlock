package sherlock

import (
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *handlerSuite) TestClustersV3Create_badBody() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("POST", "/api/clusters/v3", gin.H{
			"name": 123,
		}),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
	s.Contains(got.Message, "name")
}

func (s *handlerSuite) TestClustersV3Create_sqlValidation() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("POST", "/api/clusters/v3", ClusterV3Create{}),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
}

func (s *handlerSuite) TestClustersV3Create_defaults() {
	var got ClusterV3
	code := s.HandleRequest(
		s.NewRequest("POST", "/api/clusters/v3", ClusterV3Create{
			Name:          "cluster-name",
			Provider:      "google",
			GoogleProject: "google-project",
			ClusterV3Edit: ClusterV3Edit{
				Base:    utils.PointerTo("some-base"),
				Address: utils.PointerTo("0.0.0.0"),
			},
		}),
		&got)
	s.Equal(http.StatusCreated, code)
	s.Equal("cluster-name", got.Name)
	s.Equal("us-central1-a", got.Location)
	if s.NotNil(got.HelmfileRef) {
		s.Equal("HEAD", *got.HelmfileRef)
	}
}

func (s *handlerSuite) TestClustersV3Create_overrideDefaults() {
	var got ClusterV3
	code := s.HandleRequest(
		s.NewRequest("POST", "/api/clusters/v3", ClusterV3Create{
			Name:          "cluster-name",
			Provider:      "google",
			GoogleProject: "google-project",
			Location:      "some location",
			ClusterV3Edit: ClusterV3Edit{
				Base:                utils.PointerTo("some-base"),
				Address:             utils.PointerTo("0.0.0.0"),
				RequiresSuitability: utils.PointerTo(true),
				HelmfileRef:         utils.PointerTo("some-ref"),
			},
		}),
		&got)
	s.Equal(http.StatusCreated, code)
	s.Equal("cluster-name", got.Name)
	s.Equal("some location", got.Location)
	if s.NotNil(got.RequiresSuitability) {
		s.True(*got.RequiresSuitability)
	}
	if s.NotNil(got.HelmfileRef) {
		s.Equal("some-ref", *got.HelmfileRef)
	}
}

func (s *handlerSuite) TestClustersV3Create_suitability() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.UseNonSuitableUserFor(s.NewRequest("POST", "/api/clusters/v3", ClusterV3Create{
			Name:          "cluster-name",
			Provider:      "google",
			GoogleProject: "google-project",
			Location:      "some location",
			ClusterV3Edit: ClusterV3Edit{
				Base:         utils.PointerTo("some-base"),
				Address:      utils.PointerTo("0.0.0.0"),
				RequiredRole: s.TestData.Role_TerraSuitableEngineer().Name,
				HelmfileRef:  utils.PointerTo("some-ref"),
			},
		})),
		&got)
	s.Equal(http.StatusForbidden, code)
	s.Equal(errors.Forbidden, got.Type)
}
