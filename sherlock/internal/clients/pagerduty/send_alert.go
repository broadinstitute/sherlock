package pagerduty

import (
	"context"
	"fmt"
	"time"

	"github.com/PagerDuty/go-pagerduty"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/rs/zerolog/log"
)

type AlertSummary struct {
	Summary    string `json:"summary"`
	SourceLink string `json:"sourceLink"`
}

type SendAlertResponse struct {
	Message string `json:"message,omitempty"`
	Status  string `json:"status,omitempty"`
}

func SendAlert(integrationKey string, summary AlertSummary) (SendAlertResponse, error) {
	summaryText := summary.Summary
	if len(summaryText) > 1024 {
		summaryText = summaryText[:1024]
	}
	if config.Config.Bool("pagerduty.enable") {
		ctx := context.Background()
		log.Info().Msgf("PDTY | sending alert '%s' due to pagerduty.enable = true", summaryText)
		event := pagerduty.V2Event{
			RoutingKey: integrationKey,
			Action:     "trigger",
			Payload: &pagerduty.V2Payload{
				Summary:   summaryText,
				Source:    summary.SourceLink,
				Severity:  "critical",
				Timestamp: time.Now().Format(time.RFC3339),
			},
			// For some reason this isn't typed, so we borrow the struct from ChangeEvent so we get the right JSON
			Links: []any{
				pagerduty.ChangeEventLink{
					Href: summary.SourceLink,
					Text: "Beehive Link",
				},
			},
		}
		resp, err := client.ManageEventWithContext(ctx, &event)

		recordMetrics(ctx, "alert", err)

		if err != nil {
			return SendAlertResponse{}, err
		} else if resp == nil {
			return SendAlertResponse{}, fmt.Errorf("no error but response was nil")
		} else {
			return SendAlertResponse{
				Message: resp.Message,
				Status:  resp.Status,
			}, nil
		}
	} else {
		log.Info().Msgf("PDTY | not sending alert '%s' due to pagerduty.enable = false", summaryText)
		return SendAlertResponse{}, nil
	}
}
