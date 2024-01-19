package hooks

import (
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/broadinstitute/sherlock/sherlock/internal/hooks/hooks_mocks"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/broadinstitute/sherlock/sherlock/internal/slack"
	"github.com/broadinstitute/sherlock/sherlock/internal/slack/slack_mocks"
	"github.com/stretchr/testify/mock"
)

func (s *hooksSuite) TestDispatch_hooks() {
	ciRun := s.TestData.CiRun_Deploy_LeonardoDev_V1toV3()
	s.TestData.SlackDeployHook_Dev()
	s.TestData.GithubActionsDeployHook_LeonardoDev()
	completionText, errs := ciRun.SlackCompletionText(s.DB)
	s.Empty(errs)
	slack.UseMockedClient(s.T(), func(c *slack_mocks.MockMockableClient) {}, func() {
		UseMockedDispatcher(s.T(), func(d *hooks_mocks.MockMockableDispatcher) {
			for _, channel := range ciRun.NotifySlackChannelsUponSuccess {
				d.EXPECT().DispatchSlackCompletionNotification(
					mock.Anything,
					channel,
					completionText,
					ciRun.Succeeded(),
					ciRun.NotifySlackCustomIcon).
					Return(nil).Once()
			}
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

func (s *hooksSuite) TestDispatch_duplicateHooks() {
	ciRun := s.TestData.CiRun_Deploy_LeonardoDev_V1toV3()

	// Create two Slack hooks
	slackHook := s.TestData.SlackDeployHook_Dev()
	s.NoError(s.DB.Create(&models.SlackDeployHook{
		Trigger: models.DeployHookTriggerConfig{
			OnEnvironmentID: slackHook.Trigger.OnEnvironmentID,
			OnSuccess:       slackHook.Trigger.OnSuccess,
			OnFailure:       slackHook.Trigger.OnFailure,
		},
		SlackChannel:  slackHook.SlackChannel,
		MentionPeople: slackHook.MentionPeople,
	}).Error)

	// Create two Github Actions hooks
	ghaHook := s.TestData.GithubActionsDeployHook_LeonardoDev()
	s.NoError(s.DB.Create(&models.GithubActionsDeployHook{
		Trigger: models.DeployHookTriggerConfig{
			OnChartReleaseID: ghaHook.Trigger.OnChartReleaseID,
			OnSuccess:        ghaHook.Trigger.OnSuccess,
			OnFailure:        ghaHook.Trigger.OnFailure,
		},
		GithubActionsOwner:          ghaHook.GithubActionsOwner,
		GithubActionsRepo:           ghaHook.GithubActionsRepo,
		GithubActionsWorkflowPath:   ghaHook.GithubActionsWorkflowPath,
		GithubActionsDefaultRef:     ghaHook.GithubActionsDefaultRef,
		GithubActionsRefBehavior:    ghaHook.GithubActionsRefBehavior,
		GithubActionsWorkflowInputs: ghaHook.GithubActionsWorkflowInputs,
	}).Error)

	completionText, errs := ciRun.SlackCompletionText(s.DB)
	s.Empty(errs)
	slack.UseMockedClient(s.T(), func(c *slack_mocks.MockMockableClient) {}, func() {
		UseMockedDispatcher(s.T(), func(d *hooks_mocks.MockMockableDispatcher) {
			for _, channel := range ciRun.NotifySlackChannelsUponSuccess {
				d.EXPECT().DispatchSlackCompletionNotification(
					mock.Anything,
					channel,
					completionText,
					ciRun.Succeeded(),
					ciRun.NotifySlackCustomIcon).
					Return(nil).Once()
			}
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

func (s *hooksSuite) TestDispatch_completionMessageWithIcon() {
	ciRun := s.TestData.CiRun_Deploy_LeonardoDev_V1toV3()
	s.NoError(s.DB.Model(&ciRun).Updates(&models.CiRun{NotifySlackCustomIcon: utils.PointerTo(":smiley:")}).Error)
	completionText, errs := ciRun.SlackCompletionText(s.DB)
	s.Empty(errs)
	slack.UseMockedClient(s.T(), func(c *slack_mocks.MockMockableClient) {}, func() {
		UseMockedDispatcher(s.T(), func(d *hooks_mocks.MockMockableDispatcher) {
			for _, channel := range ciRun.NotifySlackChannelsUponSuccess {
				d.EXPECT().DispatchSlackCompletionNotification(
					mock.Anything,
					channel,
					completionText,
					ciRun.Succeeded(),
					utils.PointerTo(":smiley:")).
					Return(nil).Once()
			}
		}, func() {
			Dispatch(s.DB, ciRun)
		})
	})
}

func (s *hooksSuite) TestDispatch_noHooks() {
	ciRun := s.TestData.CiRun_Deploy_LeonardoDev_V1toV3()
	completionText, errs := ciRun.SlackCompletionText(s.DB)
	s.Empty(errs)
	slack.UseMockedClient(s.T(), func(c *slack_mocks.MockMockableClient) {}, func() {
		UseMockedDispatcher(s.T(), func(d *hooks_mocks.MockMockableDispatcher) {
			for _, channel := range ciRun.NotifySlackChannelsUponSuccess {
				d.EXPECT().DispatchSlackCompletionNotification(
					mock.Anything,
					channel,
					completionText,
					ciRun.Succeeded(),
					ciRun.NotifySlackCustomIcon).
					Return(nil).Once()
			}
		}, func() {
			Dispatch(s.DB, ciRun)
		})
	})
}

func (s *hooksSuite) TestDispatch_errors() {
	ciRun := s.TestData.CiRun_Deploy_LeonardoDev_V1toV3()
	s.TestData.SlackDeployHook_Dev()
	s.TestData.GithubActionsDeployHook_LeonardoDev()
	completionText, errs := ciRun.SlackCompletionText(s.DB)
	s.Empty(errs)
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
			for _, channel := range ciRun.NotifySlackChannelsUponSuccess {
				d.EXPECT().DispatchSlackCompletionNotification(
					mock.Anything,
					channel,
					completionText,
					ciRun.Succeeded(),
					ciRun.NotifySlackCustomIcon).
					Return(fmt.Errorf("error 1")).Once()
			}
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

func (s *hooksSuite) TestDispatch_panics() {
	ciRun := s.TestData.CiRun_Deploy_LeonardoDev_V1toV3()
	s.TestData.SlackDeployHook_Dev()
	s.TestData.GithubActionsDeployHook_LeonardoDev()
	completionText, errs := ciRun.SlackCompletionText(s.DB)
	s.Empty(errs)
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
			for _, channel := range ciRun.NotifySlackChannelsUponSuccess {
				d.EXPECT().DispatchSlackCompletionNotification(
					mock.Anything,
					channel,
					completionText,
					ciRun.Succeeded(),
					ciRun.NotifySlackCustomIcon).
					Panic("error 1").Once()
			}
			d.EXPECT().DispatchSlackDeployHook(
				mock.Anything,
				mock.Anything,
				ciRun).Panic("error 2").Once()
			d.EXPECT().DispatchGithubActionsDeployHook(
				mock.Anything,
				mock.Anything,
				ciRun).Panic("error 3").Once()
		}, func() {
			Dispatch(s.DB, ciRun)
		})
	})
}
