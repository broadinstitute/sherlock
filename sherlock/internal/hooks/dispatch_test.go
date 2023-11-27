package hooks

import (
	"fmt"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/broadinstitute/sherlock/sherlock/internal/hooks/hooks_mocks"
	"github.com/broadinstitute/sherlock/sherlock/internal/slack"
	"github.com/broadinstitute/sherlock/sherlock/internal/slack/slack_mocks"
	"github.com/stretchr/testify/mock"
)

func (s *hooksSuite) TestDispatch_hooks() {
	ciRun := s.TestData.CiRun_Deploy_LeonardoDev_V1toV3()
	s.TestData.SlackDeployHook_Dev()
	s.TestData.GithubActionsDeployHook_LeonardoDev()
	slack.UseMockedClient(s.T(), func(c *slack_mocks.MockMockableClient) {}, func() {
		UseMockedDispatcher(s.T(), func(d *hooks_mocks.MockMockableDispatcher) {
			d.EXPECT().DispatchSlackCompletionNotification(
				mock.Anything,
				ciRun.NotifySlackChannelsUponSuccess[0],
				ciRun.SlackCompletionText(s.DB),
				ciRun.Succeeded()).
				Return(nil).Once()
			d.EXPECT().DispatchSlackDeployHook(
				mock.Anything,
				mock.Anything,
				ciRun).Return(nil).Once()
			d.EXPECT().DispatchGithubActionsDeployHook(
				mock.Anything,
				mock.Anything,
				ciRun).Return(nil).Once()
		}, func() {
			Dispatch(s.DB, ciRun)
		})
	})
}

func (s *hooksSuite) TestDispatch_noHooks() {
	ciRun := s.TestData.CiRun_Deploy_LeonardoDev_V1toV3()
	slack.UseMockedClient(s.T(), func(c *slack_mocks.MockMockableClient) {}, func() {
		UseMockedDispatcher(s.T(), func(d *hooks_mocks.MockMockableDispatcher) {
			d.EXPECT().DispatchSlackCompletionNotification(
				mock.Anything,
				ciRun.NotifySlackChannelsUponSuccess[0],
				ciRun.SlackCompletionText(s.DB),
				ciRun.Succeeded()).
				Return(nil).Once()
		}, func() {
			Dispatch(s.DB, ciRun)
		})
	})
}

func (s *hooksSuite) TestDispatch_errors() {
	ciRun := s.TestData.CiRun_Deploy_LeonardoDev_V1toV3()
	s.TestData.SlackDeployHook_Dev()
	s.TestData.GithubActionsDeployHook_LeonardoDev()
	slack.UseMockedClient(s.T(), func(c *slack_mocks.MockMockableClient) {
		c.EXPECT().SendMessageContext(
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything).
			Return("", "", "", nil).
			Times(len(config.Config.Strings("slack.behaviors.errors.channels")))
	}, func() {
		UseMockedDispatcher(s.T(), func(d *hooks_mocks.MockMockableDispatcher) {
			d.EXPECT().DispatchSlackCompletionNotification(
				mock.Anything,
				ciRun.NotifySlackChannelsUponSuccess[0],
				ciRun.SlackCompletionText(s.DB),
				ciRun.Succeeded()).
				Return(fmt.Errorf("error 1")).Once()
			d.EXPECT().DispatchSlackDeployHook(
				mock.Anything,
				mock.Anything,
				ciRun).Return(fmt.Errorf("error 2")).Once()
			d.EXPECT().DispatchGithubActionsDeployHook(
				mock.Anything,
				mock.Anything,
				ciRun).Return(fmt.Errorf("error 3")).Once()
		}, func() {
			Dispatch(s.DB, ciRun)
		})
	})
}
