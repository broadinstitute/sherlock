package sherlock

import (
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"net/http"
)

func (s *handlerSuite) TestSlackDeployHooksV3List_none() {
	var got []UserV3
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/deploy-hooks/slack/v3", nil),
		&got)
	s.Equal(http.StatusOK, code)
	s.Len(got, 0)
}

func (s *handlerSuite) TestSlackDeployHooksV3List_badFilter() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/deploy-hooks/slack/v3?id=foo", nil),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
}

func (s *handlerSuite) TestSlackDeployHooksV3List_notFoundFilter() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/deploy-hooks/slack/v3?onEnvironment=foo", nil),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
}

func (s *handlerSuite) TestSlackDeployHooksV3List_badLimit() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/deploy-hooks/slack/v3?limit=foo", nil),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
}

func (s *handlerSuite) TestSlackDeployHooksV3List_badOffset() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/deploy-hooks/slack/v3?offset=foo", nil),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
}

func (s *handlerSuite) TestSlackDeployHooksV3List() {
	environment := s.TestData.Environment_Dev()
	environment2 := s.TestData.Environment_Staging()
	hook1 := models.SlackDeployHook{
		Trigger: models.DeployHookTriggerConfig{
			OnEnvironmentID: &environment.ID,
		},
		SlackChannel: utils.PointerTo("channel"),
	}
	hook2 := models.SlackDeployHook{
		Trigger: models.DeployHookTriggerConfig{
			OnEnvironmentID: &environment.ID,
		},
		SlackChannel: utils.PointerTo("different channel"),
	}
	hook3 := models.SlackDeployHook{
		Trigger: models.DeployHookTriggerConfig{
			OnEnvironmentID: &environment2.ID,
		},
		SlackChannel: utils.PointerTo("another channel"),
	}
	for _, hook := range []models.SlackDeployHook{hook1, hook2, hook3} {
		s.NoError(s.DB.Create(&hook).Error)
		s.NotZero(hook.ID)
	}

	s.Run("all", func() {
		var got []SlackDeployHookV3
		code := s.HandleRequest(
			s.NewRequest("GET", "/api/deploy-hooks/slack/v3", nil),
			&got)
		s.Equal(http.StatusOK, code)
		s.Len(got, 3)
	})
	s.Run("none", func() {
		var got []SlackDeployHookV3
		code := s.HandleRequest(
			s.NewRequest("GET", "/api/deploy-hooks/slack/v3?slackChannel=nothing", nil),
			&got)
		s.Equal(http.StatusOK, code)
		s.Len(got, 0)
	})
	s.Run("some", func() {
		var got []SlackDeployHookV3
		code := s.HandleRequest(
			s.NewRequest("GET", "/api/deploy-hooks/slack/v3?slackChannel=channel", nil),
			&got)
		s.Equal(http.StatusOK, code)
		s.Len(got, 1)
	})
	s.Run("some by environment", func() {
		var got []SlackDeployHookV3
		code := s.HandleRequest(
			s.NewRequest("GET", "/api/deploy-hooks/slack/v3?onEnvironment=dev", nil),
			&got)
		s.Equal(http.StatusOK, code)
		s.Len(got, 2)
	})
	s.Run("limit and offset", func() {
		var got1 []SlackDeployHookV3
		code := s.HandleRequest(
			s.NewRequest("GET", "/api/deploy-hooks/slack/v3?limit=1", nil),
			&got1)
		s.Equal(http.StatusOK, code)
		s.Len(got1, 1)
		var got2 []SlackDeployHookV3
		code = s.HandleRequest(
			s.NewRequest("GET", "/api/deploy-hooks/slack/v3?limit=1&offset=1", nil),
			&got2)
		s.Equal(http.StatusOK, code)
		s.Len(got2, 1)
		s.NotEqual(got1[0].ID, got2[0].ID)
	})
}
