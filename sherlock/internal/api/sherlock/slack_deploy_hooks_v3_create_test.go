package sherlock

import (
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *handlerSuite) TestSlackDeployHooksV3Create_badBody() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("POST", "/api/deploy-hooks/slack/v3", gin.H{
			"slackChannel": 123,
		}),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
	s.Contains(got.Message, "slackChannel")
}

func (s *handlerSuite) TestSlackDeployHooksV3Create_notFoundBody() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("POST", "/api/deploy-hooks/slack/v3", SlackDeployHookV3Create{
			DeployHookTriggerConfigV3: DeployHookTriggerConfigV3{
				OnEnvironment: utils.PointerTo("foo"),
			},
			SlackDeployHookFields: SlackDeployHookFields{
				SlackChannel: utils.PointerTo("channel"),
			},
		}),
		&got)
	s.Equal(http.StatusNotFound, code)
	s.Equal(errors.NotFound, got.Type)
}

func (s *handlerSuite) TestSlackDeployHooksV3Create_sqlValidation() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("POST", "/api/deploy-hooks/slack/v3", SlackDeployHookV3Create{
			DeployHookTriggerConfigV3: DeployHookTriggerConfigV3{
				OnEnvironment: utils.PointerTo(s.TestData.Environment_Dev().Name),
			},
		}),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
	s.Contains(got.Message, "slack_channel_present")
}

func (s *handlerSuite) TestSlackDeployHooksV3Create_forbidden() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewNonSuitableRequest("POST", "/api/deploy-hooks/slack/v3", SlackDeployHookV3Create{
			DeployHookTriggerConfigV3: DeployHookTriggerConfigV3{
				OnEnvironment: utils.PointerTo(s.TestData.Environment_Prod().Name),
			},
			SlackDeployHookFields: SlackDeployHookFields{
				SlackChannel: utils.PointerTo("channel"),
			},
		}),
		&got)
	s.Equal(http.StatusForbidden, code)
	s.Equal(errors.Forbidden, got.Type)
}

func (s *handlerSuite) TestSlackDeployHooksV3Create() {
	var got SlackDeployHookV3
	code := s.HandleRequest(
		s.NewRequest("POST", "/api/deploy-hooks/slack/v3", SlackDeployHookV3Create{
			DeployHookTriggerConfigV3: DeployHookTriggerConfigV3{
				OnEnvironment: utils.PointerTo(s.TestData.Environment_Dev().Name),
			},
			SlackDeployHookFields: SlackDeployHookFields{
				SlackChannel: utils.PointerTo("channel"),
			},
		}),
		&got)
	s.Equal(http.StatusCreated, code)
	if s.NotNil(got.SlackChannel) {
		s.Equal("channel", *got.SlackChannel)
	}
	if s.NotNil(got.OnEnvironment) {
		s.Equal(s.TestData.Environment_Dev().Name, *got.OnEnvironment)
	}
	if s.NotNil(got.MentionPeople) {
		s.False(*got.MentionPeople)
	}
}

func (s *handlerSuite) TestSlackDeployHooksV3Create_mentionPeople() {
	var got SlackDeployHookV3
	code := s.HandleRequest(
		s.NewRequest("POST", "/api/deploy-hooks/slack/v3", SlackDeployHookV3Create{
			DeployHookTriggerConfigV3: DeployHookTriggerConfigV3{
				OnEnvironment: utils.PointerTo(s.TestData.Environment_Dev().Name),
			},
			SlackDeployHookFields: SlackDeployHookFields{
				SlackChannel:  utils.PointerTo("channel"),
				MentionPeople: utils.PointerTo(true),
			},
		}),
		&got)
	s.Equal(http.StatusCreated, code)
	if s.NotNil(got.SlackChannel) {
		s.Equal("channel", *got.SlackChannel)
	}
	if s.NotNil(got.OnEnvironment) {
		s.Equal(s.TestData.Environment_Dev().Name, *got.OnEnvironment)
	}
	if s.NotNil(got.MentionPeople) {
		s.True(*got.MentionPeople)
	}
}
