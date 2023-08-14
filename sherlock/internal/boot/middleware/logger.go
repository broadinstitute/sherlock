package middleware

import (
	"github.com/broadinstitute/sherlock/sherlock/internal/authentication"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		t := time.Now()

		ctx.Next()

		path := ctx.Request.URL.Path
		if ctx.Request.URL.RawQuery != "" {
			path = path + "?" + ctx.Request.URL.RawQuery
		}
		identity := "client not identified"
		if user, err := authentication.ShouldUseUser(ctx); err == nil {
			identity = user.Email
		}
		var event *zerolog.Event
		switch code := ctx.Writer.Status(); {
		case code >= 500:
			event = log.Error()
		case code >= 400:
			event = log.Warn()
		default:
			event = log.Info()
		}

		if len(ctx.Errors) > 1 {
			for i, err := range ctx.Errors {
				log.Error().Err(err).Msgf("GIN  | request incurred a surprising number of errors (can't attach them all to a single log line), %d of %d: %v", i+1, len(ctx.Errors), err)
			}
		}
		if len(ctx.Errors) > 0 {
			event.Err(ctx.Errors[0])
		}

		event.Msgf("GIN  | %3d | %14s | %15s | %30s | %-7s %s",
			ctx.Writer.Status(), time.Since(t).String(), ctx.ClientIP(), identity, ctx.Request.Method, path)
	}
}
