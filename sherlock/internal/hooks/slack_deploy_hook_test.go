package hooks

import (
	"context"
	"fmt"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/broadinstitute/sherlock/sherlock/internal/slack"
	"github.com/broadinstitute/sherlock/sherlock/internal/slack/slack_mocks"
	slack2 "github.com/slack-go/slack"
	"github.com/stretchr/testify/mock"
	"sync"
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

func (s *hooksSuite) Test_dispatcherImpl_DispatchSlackDeployHook_initialMessageRace() {
	hook := s.TestData.SlackDeployHook_Dev()
	ciRun := s.TestData.CiRun_Deploy_LeonardoDev_V1toV3()
	// This ciRun is a good starting place but we want to make it look like it's in progress for the purposes of
	// simulating this race condition.
	// In a real scenario, the final dispatch is guarded by the TerminationHooksDispatchedAt field, so we wouldn't
	// have a race there. For the initial dispatch, we guard based on who created the slack state in the database,
	// and that's what we're testing here.
	ciRun.TerminalAt = nil
	ciRun.TerminationHooksDispatchedAt = nil

	// We're trying to simulate a race between two dispatches.
	// Here's the timeline we're trying to create:
	//
	//     1 --------------------------------> 4
	//                2 ----------> 3
	//
	// 1. First dispatch starts
	// 2. Second dispatch starts
	// 3. Second dispatch ends
	// 4. First dispatch ends

	// To accomplish this without sleeping, we use two mutex.
	// The first blocks the second dispatch from starting until the first dispatch has started "talking to the Slack API".
	// The second blocks the first dispatch from ending until the second dispatch has ended.
	//
	// Note that this is just barely safe with how the Slack client treats getting mocked:
	// The entire package becomes mocked for the duration of the callback, with whatever client was there previously
	// being inaccessible until the callback ends. We're mocking it twice over -- once for the first dispatch, then
	// again for the second dispatch. That means the second dispatch's mock overlays the first. That's okay in this
	// particular situation because we're not trying to interleave calls between the two dispatches.
	var firstDispatchMayEnd, secondDispatchMayStart sync.Mutex
	firstDispatchMayEnd.Lock()
	secondDispatchMayStart.Lock()

	// To make sure our goroutines complete, we use a WaitGroup.
	var wg sync.WaitGroup
	wg.Add(2)

	// The crux of our test is that the second dispatch won't ever try to talk to the (mock) Slack API.
	// We check this by asserting against this boolean, so that an errant won't cause our test to hang.
	var secondDispatchTalkedToSlack bool

	go func() {
		slack.UseMockedClient(s.T(), func(c *slack_mocks.MockMockableClient) {
			c.EXPECT().SendMessageContext(s.DB.Statement.Context, *hook.SlackChannel, mock.AnythingOfType("slack.MsgOption")).
				Run(func(ctx context.Context, channelID string, options ...slack2.MsgOption) {
					secondDispatchMayStart.Unlock()
					firstDispatchMayEnd.Lock()
				}).
				Return(*hook.SlackChannel, "123", "some text", nil).Once()
		}, func() {
			s.NoError((&dispatcherImpl{}).DispatchSlackDeployHook(s.DB, hook, ciRun))
		})
		wg.Done()
	}()

	go func() {
		secondDispatchMayStart.Lock()
		slack.UseMockedClient(s.T(), func(c *slack_mocks.MockMockableClient) {
			c.EXPECT().SendMessageContext(s.DB.Statement.Context, *hook.SlackChannel, mock.AnythingOfType("slack.MsgOption")).
				Run(func(ctx context.Context, channelID string, options ...slack2.MsgOption) {
					secondDispatchTalkedToSlack = true
				}).
				Return(*hook.SlackChannel, "123", "some text", nil).
				// All important, while we do configure the mock we don't actually require that it run.
				// If it runs, we set the boolean to true so we fail in a sec. Again, we provide this
				// just so that if there's a failure the test doesn't actually hang.
				Maybe()

		}, func() {
			s.NoError((&dispatcherImpl{}).DispatchSlackDeployHook(s.DB, hook, ciRun))
		})
		firstDispatchMayEnd.Unlock()
		wg.Done()
	}()

	wg.Wait()
	s.False(secondDispatchTalkedToSlack, "second dispatch talked to Slack API, it shouldn't have")
}

func (s *hooksSuite) Test_dispatcherImpl_DispatchSlackDeployHook_stateDeletedUponInitialFailure() {
	hook := s.TestData.SlackDeployHook_Dev()
	ciRun := s.TestData.CiRun_Deploy_LeonardoDev_V1toV3()
	slack.UseMockedClient(s.T(), func(c *slack_mocks.MockMockableClient) {
		c.EXPECT().SendMessageContext(s.DB.Statement.Context, *hook.SlackChannel, mock.AnythingOfType("slack.MsgOption")).
			Return("", "", "", fmt.Errorf("some error")).Once()
	}, func() {
		s.Error((&dispatcherImpl{}).DispatchSlackDeployHook(s.DB, hook, ciRun))
	})

	var messageStates []models.SlackDeployHookState
	s.NoError(s.DB.Find(&messageStates).Error)
	s.Empty(messageStates)
}

func (s *hooksSuite) Test_dispatcherImpl_DispatchSlackDeployHook_stateNotDeletedUponUpdateFailure() {
	hook := s.TestData.SlackDeployHook_Dev()
	ciRun := s.TestData.CiRun_Deploy_LeonardoDev_V1toV3()
	// We create the state here because we want to pretend that a message was already sent.
	s.NoError(s.DB.Create(&models.SlackDeployHookState{
		SlackDeployHookID: hook.ID,
		CiRunID:           ciRun.ID,
		MessageChannel:    *hook.SlackChannel,
		MessageTimestamp:  "1234",
	}).Error)
	slack.UseMockedClient(s.T(), func(c *slack_mocks.MockMockableClient) {
		c.EXPECT().UpdateMessageContext(s.DB.Statement.Context, *hook.SlackChannel, "1234", mock.AnythingOfType("slack.MsgOption")).
			Return("", "", "", fmt.Errorf("some error")).Once()
	}, func() {
		s.Error((&dispatcherImpl{}).DispatchSlackDeployHook(s.DB, hook, ciRun))
	})

	var messageStates []models.SlackDeployHookState
	s.NoError(s.DB.Find(&messageStates).Error)
	s.Len(messageStates, 1)
}
