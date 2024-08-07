package sherlock

import (
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/clients/slack"
	"github.com/broadinstitute/sherlock/sherlock/internal/clients/slack/slack_mocks"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/stretchr/testify/mock"
	"net/http"
)

func (s *handlerSuite) TestSlackDeployHooksV3TestRun_badSelector() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("POST", "/api/deploy-hooks/slack/procedures/v3/test/foo-bar",
			SlackDeployHookTestRunRequest{Execute: utils.PointerTo(true)}),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
}

func (s *handlerSuite) TestSlackDeployHooksV3TestRun_missingBody() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("POST", "/api/deploy-hooks/slack/procedures/v3/test/0", nil),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
	s.Contains(got.Message, "execute")
}

func (s *handlerSuite) TestSlackDeployHooksV3TestRun_notFound() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("POST", "/api/deploy-hooks/slack/procedures/v3/test/0",
			SlackDeployHookTestRunRequest{Execute: utils.PointerTo(true)}),
		&got)
	s.Equal(http.StatusNotFound, code)
	s.Equal(errors.NotFound, got.Type)
}

func (s *handlerSuite) TestSlackDeployHooksV3TestRun() {
	hook := models.SlackDeployHook{
		Trigger: models.DeployHookTriggerConfig{
			OnEnvironmentID: utils.PointerTo(s.TestData.Environment_Dev().ID),
		},
		SlackChannel: utils.PointerTo("channel"),
	}
	s.NoError(s.DB.Create(&hook).Error)

	s.Run("no error", func() {
		var got SlackDeployHookTestRunResponse
		slack.UseMockedClient(s.T(), func(c *slack_mocks.MockMockableClient) {
			c.On("SendMessageContext", mock.Anything, "channel",
				mock.AnythingOfType("slack.MsgOption")).Return("", "", "", nil)
		}, func() {
			code := s.HandleRequest(
				s.NewRequest("POST", fmt.Sprintf("/api/deploy-hooks/slack/procedures/v3/test/%d", hook.ID),
					SlackDeployHookTestRunRequest{Execute: utils.PointerTo(true)}),
				&got)
			s.Equal(http.StatusOK, code)
		})
	})
	s.Run("don't execute", func() {
		var got SlackDeployHookTestRunResponse
		slack.UseMockedClient(s.T(), func(c *slack_mocks.MockMockableClient) {}, func() {
			code := s.HandleRequest(
				s.NewRequest("POST", fmt.Sprintf("/api/deploy-hooks/slack/procedures/v3/test/%d", hook.ID),
					SlackDeployHookTestRunRequest{Execute: utils.PointerTo(false)}),
				&got)
			s.Equal(http.StatusOK, code)
		})
	})
	s.Run("returns error", func() {
		var got errors.ErrorResponse
		// We use errors.BadRequest so that a 500 doesn't get logged during the test
		slack.UseMockedClient(s.T(), func(c *slack_mocks.MockMockableClient) {
			c.On("SendMessageContext", mock.Anything, "channel",
				mock.AnythingOfType("slack.MsgOption")).Return("", "", "", fmt.Errorf(errors.BadRequest))
		}, func() {
			code := s.HandleRequest(
				s.NewRequest("POST", fmt.Sprintf("/api/deploy-hooks/slack/procedures/v3/test/%d", hook.ID),
					SlackDeployHookTestRunRequest{Execute: utils.PointerTo(true)}),
				&got)
			s.Equal(http.StatusBadRequest, code)
			s.Equal(fmt.Sprintf("error between Sherlock and Slack: %s", errors.BadRequest), got.Message)
		})
	})
}
