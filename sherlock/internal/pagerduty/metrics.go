package pagerduty

import (
	"context"
	"github.com/PagerDuty/go-pagerduty"
	"github.com/broadinstitute/sherlock/sherlock/internal/metrics"
	"go.opencensus.io/stats"
	"go.opencensus.io/tag"
	"net/http"
	"strconv"
)

func recordMetrics(ctx context.Context, requestType string, err error) {
	ctx, _ = tag.New(ctx, tag.Upsert(metrics.PagerdutyRequestTypeKey, requestType))
	if err != nil {
		if customError, ok := err.(pagerduty.EventsAPIV2Error); ok {
			ctx, _ = tag.New(ctx, tag.Upsert(metrics.PagerdutyResponseCodeKey, strconv.Itoa(customError.StatusCode)))
		} else if customError, ok := err.(pagerduty.APIError); ok {
			ctx, _ = tag.New(ctx, tag.Upsert(metrics.PagerdutyResponseCodeKey, strconv.Itoa(customError.StatusCode)))
		} else {
			ctx, _ = tag.New(ctx, tag.Delete(metrics.PagerdutyResponseCodeKey))
		}
	} else {
		ctx, _ = tag.New(ctx, tag.Upsert(metrics.PagerdutyResponseCodeKey, strconv.Itoa(http.StatusAccepted)))
	}
	stats.Record(ctx, metrics.PagerdutyRequestCount.M(1))
}
