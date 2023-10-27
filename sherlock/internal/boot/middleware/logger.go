package middleware

import (
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/authentication"
	"github.com/broadinstitute/sherlock/sherlock/internal/metrics"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go.opencensus.io/stats"
	"go.opencensus.io/tag"
	"strconv"
	"time"
)

func Logger(consoleLogging bool) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		t := time.Now()

		ctx.Next()

		latency := time.Since(t)

		path := ctx.Request.URL.Path
		if ctx.Request.URL.RawQuery != "" {
			path = path + "?" + ctx.Request.URL.RawQuery
		}

		var principal string
		if user, err := authentication.ShouldUseUser(ctx); err == nil {
			principal = user.Email
		} else {
			principal = "unevaluated"
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

		if len(ctx.Errors) > 0 {
			event.Errs("errors", utils.Map(ctx.Errors, func(e *gin.Error) error { return e }))
		}

		event.Int("status", ctx.Writer.Status())
		event.Str("principal", principal)
		event.Str("method", ctx.Request.Method)
		if consoleLogging {
			event.Stringer("latency", latency)
			event.Msgf("GIN  | %-50s", path)
		} else {
			event.Dur("latency", latency)
			event.Str("route", ctx.FullPath())
			event.Str("client", ctx.ClientIP())
			event.Msgf("GIN  | %3d | %14s | %40s | %-7s %s",
				ctx.Writer.Status(), time.Since(t).String(), principal, ctx.Request.Method, path)
		}

		if tagCtx, err := tag.New(ctx,
			tag.Upsert(metrics.StatusKey, strconv.Itoa(ctx.Writer.Status())),
			tag.Upsert(metrics.MethodKey, ctx.Request.Method),
			tag.Upsert(metrics.RouteKey, ctx.FullPath())); err == nil {
			stats.Record(tagCtx, metrics.ResponseLatencyMeasure.M(latency.Milliseconds()))
		} else {
			log.Warn().Err(err).Msg("unable to record latency")
		}
	}
}
