package sherlock

import (
	"github.com/broadinstitute/sherlock/internal/auth"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"time"
)

func logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		t := time.Now()

		ctx.Next()

		path := ctx.Request.URL.Path
		if ctx.Request.URL.RawQuery != "" {
			path = path + "?" + ctx.Request.URL.RawQuery
		}
		identity := "client not identified"
		if user, err := auth.ExtractUserFromContext(ctx); err == nil {
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
		event.Msgf("GIN  | %3d | %14s | %15s | %30s | %-7s %s",
			ctx.Writer.Status(), time.Since(t).String(), ctx.ClientIP(), identity, ctx.Request.Method, path)
	}
}
