package sherlock

import (
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
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

func (s *handlerSuite) TestSlackDeployHooksV3Delete_forbidden() {
	hook := models.SlackDeployHook{
		Trigger: models.DeployHookTriggerConfig{
			OnEnvironmentID: utils.PointerTo(s.TestData.Environment_Prod().ID),
		},
		SlackChannel: utils.PointerTo("channel"),
	}
	s.NoError(s.DB.Create(&hook).Error)

	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewNonSuitableRequest("DELETE", fmt.Sprintf("/api/deploy-hooks/slack/v3/%d", hook.ID), nil),
		&got)
	s.Equal(http.StatusForbidden, code)
	s.Equal(errors.Forbidden, got.Type)
}

func (s *handlerSuite) TestSlackDeployHooksV3Delete() {

	hook := models.SlackDeployHook{
		Trigger: models.DeployHookTriggerConfig{
			OnEnvironmentID: utils.PointerTo(s.TestData.Environment_Dev().ID),
		},
		SlackChannel: utils.PointerTo("channel"),
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
		s.Equal(s.TestData.Environment_Dev().Name, *got.OnEnvironment)
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
			s.Equal(s.TestData.Environment_Dev().ID, remaining[0].ID)
		}
	})
}
