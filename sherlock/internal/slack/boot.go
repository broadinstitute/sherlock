package slack

import (
	"context"
	"errors"
	"fmt"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/rs/zerolog/log"
	"github.com/slack-go/slack/socketmode"

	"github.com/slack-go/slack"
)

type slackZerologShim struct {
	prefix string
}

func (s *slackZerologShim) Output(_ int, message string) error {
	log.Info().Msg(s.prefix + message)
	return nil
}

func Init(ctx context.Context) error {
	if config.Config.Bool("slack.enable") {
		var api *slack.Client
		if appToken := config.Config.String("slack.appToken"); len(appToken) == 0 {
			return fmt.Errorf("slack enabled but app token not provided")
		} else if botToken := config.Config.String("slack.botToken"); len(botToken) == 0 {
			return fmt.Errorf("slack enabled but bot token not provided")
		} else {
			api = slack.New(
				botToken,
				slack.OptionAppLevelToken(appToken),
				slack.OptionDebug(config.Config.Bool("slack.debug")),
				slack.OptionLog(&slackZerologShim{"SLCK | LIBRARY (via API) | "}),
			)
		}

		if response, err := api.AuthTestContext(ctx); err != nil {
			return fmt.Errorf("slack authentication failed self-test, see https://api.slack.com/methods/auth.test#errors for more information: %v", err)
		} else {
			log.Info().Msgf("SLCK | successfully authenticated to Slack as \"%s\" in the \"%s\" workspace", response.User, response.Team)
		}

		rawClient = socketmode.New(
			api,
			socketmode.OptionDebug(config.Config.Bool("slack.debug")),
			socketmode.OptionLog(&slackZerologShim{"SLCK | LIBRARY (via socket) | "}),
		)
		client = rawClient
	}
	return nil
}

func Start(ctx context.Context) {
	if config.Config.Bool("slack.enable") {
		socketmodeHandler := socketmode.NewSocketmodeHandler(rawClient)

		socketmodeHandler.Handle(socketmode.EventTypeConnecting, handleConnecting)
		socketmodeHandler.Handle(socketmode.EventTypeConnectionError, handleConnectionError)
		socketmodeHandler.Handle(socketmode.EventTypeConnected, handleConnected)
		socketmodeHandler.Handle(socketmode.EventTypeHello, handleHello)
		socketmodeHandler.Handle(socketmode.EventTypeEventsAPI, handleEvents)

		if err := socketmodeHandler.RunEventLoopContext(ctx); err != nil && !errors.Is(err, context.Canceled) {
			log.Error().Err(err).Msgf("SLCK | INCOMING | socket event loop exited with error: %v", err)
		}
	}
}
