package middleware

import (
	"context"
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/authentication"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/broadinstitute/sherlock/sherlock/internal/metrics"
	"github.com/broadinstitute/sherlock/sherlock/internal/slack"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go.opencensus.io/stats"
	"go.opencensus.io/tag"
	"strconv"
	"time"
)

func Logger() gin.HandlerFunc {
	// We're breaking a bit of encapsulation by reaching into config like this, but
	// we're providing a speed-up for an extremely hot path by making the logger
	// simply not have to check the config each time.
	debugMode := config.Config.String("mode") == "debug"
	var errorCodesToReport []int
	if config.Config.Bool("slack.behaviors.errors.enable") {
		errorCodesToReport = config.Config.Ints("slack.behaviors.errors.statusCodes")
	}
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

		errs := utils.Map(ctx.Errors, func(e *gin.Error) error { return e })

		var event *zerolog.Event
		switch code := ctx.Writer.Status(); {
		case code >= 500:
			event = log.Error()
		case code >= 400:
			event = log.Warn()
		default:
			event = log.Info()
		}

		if len(errs) > 0 {
			event.Errs("errors", errs)
		}
		event.Int("status", ctx.Writer.Status())
		event.Str("principal", principal)
		event.Str("method", ctx.Request.Method)
		if debugMode {
			event.Stringer("latency", latency)
			event.Msgf("GIN  | %-50s", path)
		} else {
			if claims, err := authentication.ShouldUseGithubClaims(ctx); err == nil {
				// If we have GitHub claims -- meaning this request definitely came from GitHub Actions --
				// we might as well log some info from it to help with debugging.
				event.Str("githubWorkflowFile", claims.WorkflowRef)
				event.Str("githubWorkflowLink", claims.WorkflowURL())
			}
			event.Dur("latency", latency)
			event.Str("route", ctx.FullPath())
			event.Str("client", ctx.ClientIP())
			event.Msgf("GIN  | %3d | %14s | %50s | %-7s %s",
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

		if utils.Contains(errorCodesToReport, ctx.Writer.Status()) {
			description := fmt.Sprintf("%s's %s %s returned %d", principal, ctx.Request.Method, ctx.Request.URL.Path, ctx.Writer.Status())
			if debugMode {
				slack.ReportError(context.Background(), description, errs...)
			} else {
				go slack.ReportError(context.Background(), description, errs...)
			}
		}
	}
}
