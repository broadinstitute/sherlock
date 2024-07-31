package sherlock

import (
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *handlerSuite) TestClustersV3Edit_badSelector() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("PATCH", "/api/clusters/v3/something/with/slashes", nil),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
}

func (s *handlerSuite) TestClustersV3Edit_badBody() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("PATCH", "/api/clusters/v3/123", gin.H{
			"base": 123,
		}),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
	s.Contains(got.Message, "base")
}

func (s *handlerSuite) TestClustersV3Edit_notFound() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("PATCH", "/api/clusters/v3/123", ClusterV3Edit{
			Base:    utils.PointerTo("some-base"),
			Address: utils.PointerTo("0.0.0.0"),
		}),
		&got)
	s.Equal(http.StatusNotFound, code)
	s.Equal(errors.NotFound, got.Type)
}

func (s *handlerSuite) TestClustersV3Edit_sqlValidation() {
	s.SetNonSuitableTestUserForDB()
	edit := models.Cluster{
		Name:                "some-name",
		Provider:            "azure",
		AzureSubscription:   "some-subscription",
		Location:            "some-location",
		Base:                utils.PointerTo("some base"),
		Address:             utils.PointerTo("0.0.0.0"),
		RequiresSuitability: utils.PointerTo(false),
		HelmfileRef:         utils.PointerTo("some-ref"),
	}
	s.NoError(s.DB.Create(&edit).Error)

	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("PATCH", fmt.Sprintf("/api/clusters/v3/%d", edit.ID), ClusterV3Edit{
			Base:    utils.PointerTo(""),
			Address: utils.PointerTo("0.0.0.0"),
		}),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
	s.Contains(got.Message, "base")
}

func (s *handlerSuite) TestClustersV3Edit() {
	s.SetNonSuitableTestUserForDB()
	edit := models.Cluster{
		Name:                "some-name",
		Provider:            "azure",
		AzureSubscription:   "some-subscription",
		Location:            "some-location",
		Base:                utils.PointerTo("some base"),
		Address:             utils.PointerTo("0.0.0.0"),
		RequiresSuitability: utils.PointerTo(false),
		HelmfileRef:         utils.PointerTo("some-ref"),
	}
	s.NoError(s.DB.Create(&edit).Error)

	var got ClusterV3
	code := s.HandleRequest(
		s.NewRequest("PATCH", fmt.Sprintf("/api/clusters/v3/%d", edit.ID), ClusterV3Edit{
			Base:    utils.PointerTo("some other base"),
			Address: utils.PointerTo("0.0.0.0"),
		}),
		&got)
	s.Equal(http.StatusOK, code)
	if s.NotNil(got.Base) {
		s.Equal("some other base", *got.Base)
	}
}

func (s *handlerSuite) TestClustersV3Edit_suitability() {
	s.SetSuitableTestUserForDB()
	edit := models.Cluster{
		Name:              "some-name",
		Provider:          "azure",
		AzureSubscription: "some-subscription",
		Location:          "some-location",
		Base:              utils.PointerTo("some base"),
		Address:           utils.PointerTo("0.0.0.0"),
		RequiredRoleID:    utils.PointerTo(s.TestData.Role_TerraSuitableEngineer().ID),
		HelmfileRef:       utils.PointerTo("some-ref"),
	}
	s.NoError(s.DB.Create(&edit).Error)

	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.UseNonSuitableUserFor(s.NewRequest("PATCH", fmt.Sprintf("/api/clusters/v3/%d", edit.ID), ClusterV3Edit{
			Base:    utils.PointerTo("some other base"),
			Address: utils.PointerTo("0.0.0.0"),
		})),
		&got)
	s.Equal(http.StatusForbidden, code)
	s.Equal(errors.Forbidden, got.Type)
}

func (s *handlerSuite) TestClustersV3Edit_suitabilityBefore() {
	s.SetSuitableTestUserForDB()
	edit := models.Cluster{
		Name:              "some-name",
		Provider:          "azure",
		AzureSubscription: "some-subscription",
		Location:          "some-location",
		Base:              utils.PointerTo("some base"),
		Address:           utils.PointerTo("0.0.0.0"),
		RequiredRoleID:    utils.PointerTo(s.TestData.Role_TerraSuitableEngineer().ID),
		HelmfileRef:       utils.PointerTo("some-ref"),
	}
	s.NoError(s.DB.Create(&edit).Error)

	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.UseNonSuitableUserFor(s.NewRequest("PATCH", fmt.Sprintf("/api/clusters/v3/%d", edit.ID), ClusterV3Edit{
			Base:                utils.PointerTo("some other base"),
			Address:             utils.PointerTo("0.0.0.0"),
			RequiresSuitability: utils.PointerTo(false),
		})),
		&got)
	s.Equal(http.StatusForbidden, code)
	s.Equal(errors.Forbidden, got.Type)
}

func (s *handlerSuite) TestClustersV3Edit_suitabilityAfter() {
	s.SetNonSuitableTestUserForDB()
	edit := models.Cluster{
		Name:              "some-name",
		Provider:          "azure",
		AzureSubscription: "some-subscription",
		Location:          "some-location",
		Base:              utils.PointerTo("some base"),
		Address:           utils.PointerTo("0.0.0.0"),
		RequiredRoleID:    utils.PointerTo(s.TestData.Role_TerraEngineer().ID),
		HelmfileRef:       utils.PointerTo("some-ref"),
	}
	s.NoError(s.DB.Create(&edit).Error)

	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.UseNonSuitableUserFor(s.NewRequest("PATCH", fmt.Sprintf("/api/clusters/v3/%d", edit.ID), ClusterV3Edit{
			Base:         utils.PointerTo("some other base"),
			Address:      utils.PointerTo("0.0.0.0"),
			RequiredRole: s.TestData.Role_TerraSuitableEngineer().Name,
		})),
		&got)
	s.Equal(http.StatusForbidden, code)
	s.Equal(errors.Forbidden, got.Type)
}

func (s *handlerSuite) TestClusterV3Edit_clearRequiredRole() {
	toEdit := s.TestData.Cluster_TerraProd()
	s.NotNil(toEdit.RequiredRoleID)
	var got ClusterV3
	code := s.HandleRequest(
		s.NewSuitableRequest("PATCH", fmt.Sprintf("/api/clusters/v3/%d", toEdit.ID), ClusterV3Edit{
			RequiredRole: utils.PointerTo(""),
		}),
		&got)
	s.Equal(http.StatusOK, code)
	if s.NotNil(got.RequiredRole) {
		s.Equal(config.Config.String("model.roles.substituteEmptyRequiredRoleWithValue"), *got.RequiredRole)
	}
	var inDB models.Cluster
	s.NoError(s.DB.First(&inDB, toEdit.ID).Error)
	s.Nil(inDB.RequiredRoleID)
}
