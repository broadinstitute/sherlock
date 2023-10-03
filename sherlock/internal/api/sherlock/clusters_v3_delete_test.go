package sherlock

import (
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"net/http"
)

func (s *handlerSuite) TestClusterV3Delete_badSelector() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("DELETE", "/api/clusters/v3/something/with/slashes", nil),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
}

func (s *handlerSuite) TestClusterV3Delete_notFound() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("DELETE", "/api/clusters/v3/my-cluster", nil),
		&got)
	s.Equal(http.StatusNotFound, code)
	s.Equal(errors.NotFound, got.Type)
}

func (s *handlerSuite) TestClusterV3Delete() {
	s.SetNonSuitableTestUserForDB()
	s.NoError(s.DB.Create(&models.Cluster{
		Name:                "some-name",
		Provider:            "azure",
		AzureSubscription:   "some-subscription",
		Location:            "some-location",
		Base:                utils.PointerTo("some base"),
		Address:             utils.PointerTo("0.0.0.0"),
		RequiresSuitability: utils.PointerTo(false),
		HelmfileRef:         utils.PointerTo("some-ref"),
	}).Error)

	var got ClusterV3
	code := s.HandleRequest(
		s.NewRequest("DELETE", "/api/clusters/v3/some-name", nil),
		&got)
	s.Equal(http.StatusOK, code)
	if s.NotNil(got.Base) {
		s.Equal("some base", *got.Base)
	}
}
