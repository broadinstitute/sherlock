package slack

import (
	"context"
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/slack-go/slack"
)

func ReportError(ctx context.Context, err error) {
	if isEnabled() && config.Config.Bool("slack.behaviors.errors.enable") {

		log.Info().Err(err).Msgf("SLCK | reporting error: %v", err)

		for _, channel := range config.Config.Strings("slack.behaviors.errors.channels") {
			_, _, _, sendErr := client.SendMessageContext(ctx, channel,
				slack.MsgOptionText("Sherlock encountered an unexpected error:", true),
				slack.MsgOptionAttachments(slack.Attachment{
					Color: config.Config.String("slack.behaviors.errors.color"),
					Text:  err.Error(),
				}))
			if sendErr != nil {
				log.Warn().Err(sendErr).Msgf("SLCK | unable to send error message to %s: %v", channel, sendErr)
			}
		}
	}
}

func ErrorReportingMiddleware(outerCtx context.Context) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		if utils.Contains(config.Config.Ints("slack.behaviors.errors.statusCodes"), ctx.Writer.Status()) &&
			isEnabled() && config.Config.Bool("slack.behaviors.errors.enable") {

			callback := func(ctx context.Context, status int, errors []*gin.Error) {
				if len(errors) > 0 {
					for _, err := range errors {
						ReportError(ctx, err)
					}
				} else {
					ReportError(ctx, fmt.Errorf("unknown %d error (handler didn't attach errors to Gin context)", status))
				}
			}

			// Offline, call this synchronously so we can actually test it.
			// We're using the outerCtx mainly for the goroutine case, where we're running this function outside
			// the literal and figurative context of the handler.
			if config.Config.String("mode") == "debug" {
				callback(outerCtx, ctx.Writer.Status(), ctx.Errors)
			} else {
				go callback(outerCtx, ctx.Writer.Status(), ctx.Errors)
			}
		}
	}
}
