package sherlock

import (
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/testutils"
	"github.com/broadinstitute/sherlock/sherlock/internal/deprecated_models/v2models"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"net/http"
)

func (s *handlerSuite) TestSlackDeployHooksV3Delete_badSelector() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("DELETE", "/api/deploy-hooks/slack/v3/foo-bar", nil),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
}

func (s *handlerSuite) TestSlackDeployHooksV3Delete_notFound() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("DELETE", "/api/deploy-hooks/slack/v3/0", nil),
		&got)
	s.Equal(http.StatusNotFound, code)
	s.Equal(errors.NotFound, got.Type)
}

func (s *handlerSuite) TestSlackDeployHooksV3Delete() {
	user := s.SetSuitableTestUserForDB()
	cluster, created, err := v2models.InternalClusterStore.Create(s.DB, v2models.Cluster{
		Name:                "terra-dev",
		Provider:            "google",
		GoogleProject:       "broad-dsde-dev",
		Base:                testutils.PointerTo("live"),
		Address:             testutils.PointerTo("0.0.0.0"),
		RequiresSuitability: testutils.PointerTo(false),
		Location:            "us-central1-a",
		HelmfileRef:         testutils.PointerTo("HEAD"),
	}, user)
	s.NoError(err)
	s.True(created)
	environment, created, err := v2models.InternalEnvironmentStore.Create(s.DB, v2models.Environment{
		Name:                       "dev",
		Lifecycle:                  "static",
		UniqueResourcePrefix:       "a1b2",
		Base:                       "live",
		DefaultClusterID:           &cluster.ID,
		DefaultNamespace:           "terra-dev",
		OwnerID:                    &user.ID,
		RequiresSuitability:        testutils.PointerTo(false),
		HelmfileRef:                testutils.PointerTo("HEAD"),
		DefaultFirecloudDevelopRef: testutils.PointerTo("dev"),
		PreventDeletion:            testutils.PointerTo(false),
	}, user)
	s.NoError(err)
	s.True(created)

	hook := models.SlackDeployHook{
		Trigger: models.DeployHookTriggerConfig{
			OnEnvironmentID: &environment.ID,
		},
		SlackChannel: testutils.PointerTo("channel"),
	}
	s.NoError(s.DB.Create(&hook).Error)

	var got SlackDeployHookV3
	code := s.HandleRequest(
		s.NewRequest("DELETE", fmt.Sprintf("/api/deploy-hooks/slack/v3/%d", hook.ID), nil),
		&got)
	s.Equal(http.StatusOK, code)
	if s.NotNil(got.SlackChannel) {
		s.Equal("channel", *got.SlackChannel)
	}
	if s.NotNil(got.OnEnvironment) {
		s.Equal(environment.Name, *got.OnEnvironment)
	}

	s.Run("was deleted", func() {
		var remaining []models.SlackDeployHook
		s.NoError(s.DB.Where(&models.SlackDeployHook{}).Find(&remaining).Error)
		s.Len(remaining, 0)
	})
	s.Run("inner trigger config was deleted", func() {
		var remaining []models.DeployHookTriggerConfig
		s.NoError(s.DB.Where(&models.DeployHookTriggerConfig{}).Find(&remaining).Error)
		s.Len(remaining, 0)
	})
	s.Run("environment still exists", func() {
		var remaining []models.Environment
		s.NoError(s.DB.Where(&models.Environment{}).Find(&remaining).Error)
		if s.Len(remaining, 1) {
			s.Equal(environment.ID, remaining[0].ID)
		}
	})
}
