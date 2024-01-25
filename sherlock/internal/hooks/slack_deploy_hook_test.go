package hooks

import (
	"github.com/broadinstitute/sherlock/sherlock/internal/slack"
	"github.com/broadinstitute/sherlock/sherlock/internal/slack/slack_mocks"
	"github.com/stretchr/testify/mock"
)

func (s *hooksSuite) Test_dispatcherImpl_DispatchSlackDeployHook_emptyChannel() {
	ciRun := s.TestData.CiRun_Deploy_LeonardoDev_V1toV3()
	s.NoError(ciRun.FillRelatedResourceStatuses(s.DB))
	hook := s.TestData.SlackDeployHook_Dev()
	hook.SlackChannel = nil
	s.ErrorContains((&dispatcherImpl{}).DispatchSlackDeployHook(s.DB, hook, ciRun), "slack channel was nil")
}

func (s *hooksSuite) Test_dispatcherImpl_DispatchSlackDeployHook() {
	ciRun := s.TestData.CiRun_Deploy_LeonardoDev_V1toV3()
	s.NoError(ciRun.FillRelatedResourceStatuses(s.DB))
	hook := s.TestData.SlackDeployHook_Dev()
	slack.UseMockedClient(s.T(), func(c *slack_mocks.MockMockableClient) {
		c.EXPECT().SendMessageContext(s.DB.Statement.Context, *hook.SlackChannel, mock.AnythingOfType("slack.MsgOption")).
			Return(*hook.SlackChannel+"-different", "123", "some text", nil).Once()
		c.EXPECT().SendMessageContext(s.DB.Statement.Context, *hook.SlackChannel+"-different", mock.AnythingOfType("slack.MsgOption"), mock.AnythingOfType("slack.MsgOption")).
			Return(*hook.SlackChannel+"-different", "123", "some text", nil).Once()
	}, func() {
		s.NoError((&dispatcherImpl{}).DispatchSlackDeployHook(s.DB, hook, ciRun))
	})
}
