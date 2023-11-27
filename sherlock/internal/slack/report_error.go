package slack

import (
	"context"
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func ReportError(ctx context.Context, errs ...error) {
	if isEnabled() && config.Config.Bool("slack.behaviors.errors.enable") && len(errs) > 0 {

		var messageText string
		if len(errs) == 1 {
			messageText = "Sherlock encountered an unexpected error:"
			log.Info().Err(errs[0]).Msgf("SLCK | reporting error: %v", errs[0])
		} else {
			messageText = fmt.Sprintf("Sherlock encountered %d unexpected errors:", len(errs))
			log.Info().Errs("errors", errs).Msgf("SLCK | reporting %d errors, starting with: %v", len(errs), errs[0])
		}

		attachments := utils.Map(errs, func(e error) Attachment { return RedBlock{Text: e.Error()} })

		for _, channel := range config.Config.Strings("slack.behaviors.errors.channels") {
			SendMessage(ctx, channel, messageText, attachments...)
		}
	} else if config.Config.String("mode") == "debug" {
		log.Warn().Errs("errors", errs).Msg("Slack disabled in debug mode; would've reported errors if enabled")
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
