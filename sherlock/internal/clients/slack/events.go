package slack

import (
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/rs/zerolog/log"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
	"github.com/slack-go/slack/socketmode"
)

// socketmode.EventTypeConnecting
func handleConnecting(_ *socketmode.Event, _ *socketmode.Client) {
	log.Info().Msgf("SLCK | connecting socket...")
}

// socketmode.EventTypeConnectionError
func handleConnectionError(_ *socketmode.Event, _ *socketmode.Client) {
	log.Info().Msgf("SLCK | socket connection failed, will retry...")
}

// socketmode.EventTypeConnected
func handleConnected(_ *socketmode.Event, _ *socketmode.Client) {
	log.Info().Msgf("SLCK | socket connected")
}

// socketmode.EventTypeHello
func handleHello(_ *socketmode.Event, _ *socketmode.Client) {
	log.Info().Msgf("SLCK | successfully received hello message over socket")
}

// socketmode.EventTypeEventsAPI
func handleEvents(event *socketmode.Event, client *socketmode.Client) {
	eventsAPIEvent, ok := event.Data.(slackevents.EventsAPIEvent)
	if !ok {
		log.Warn().Msgf("SLCK | handleEvents ignoring non-event %+v", event)
		return
	}

	client.Ack(*event.Request)

	switch eventsAPIEvent.Type {

	case slackevents.CallbackEvent:
		switch e := eventsAPIEvent.InnerEvent.Data.(type) {
		case *slackevents.AppMentionEvent:
			handleAppMentionEvent(client, e)
		default:
			log.Warn().Msgf("SLCK | unsupported event received (\"%s\")", eventsAPIEvent.InnerEvent.Type)
		}

	case slackevents.URLVerification:
		log.Error().Msgf("SLCK | was asked to perform URL verification, but that isn't expected or possible in socket mode")

	case slackevents.AppRateLimited:
		log.Warn().Msgf("SLCK | was notified of rate limiting, except over socket instead of in response to a request")

	default:
		log.Warn().Msgf("SLCK | unsupported event wrapper outer type (\"%s\")", eventsAPIEvent.Type)
	}
}

func handleAppMentionEvent(client mockableClient, event *slackevents.AppMentionEvent) {
	if config.Config.Bool("slack.behaviors.reactToMentionsWithEmoji.enable") {
		err := client.AddReaction(
			config.Config.String("slack.behaviors.reactToMentionsWithEmoji.emoji"),
			slack.ItemRef{
				Channel:   event.Channel,
				Timestamp: event.EventTimeStamp,
			},
		)
		if err != nil {
			log.Error().Err(err).Msgf("SLCK | error handling slackevents.AppMentionEvent: %v", err)
		} else {
			log.Info().Msgf("SLCK | handled slackevents.AppMentionEvent in %s", event.Channel)
		}
	}
}
