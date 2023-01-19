package pagerduty

import (
	"context"
	"fmt"
	"github.com/PagerDuty/go-pagerduty"
	"github.com/broadinstitute/sherlock/internal/config"
	"github.com/rs/zerolog/log"
	"time"
)

type AlertSummary struct {
	Summary string `json:"summary"`
}

type SendAlertResponse struct {
	Message string `json:"message,omitempty"`
	Status  string `json:"status,omitempty"`
}

func SendAlert(integrationKey string, summary AlertSummary, sourceLink string) (SendAlertResponse, error) {
	if config.Config.Bool("pagerduty.enable") {
		ctx := context.Background()
		log.Info().Msgf("PDTY | sending alert '%s' due to pagerduty.enable = true", summary.Summary)
		event := pagerduty.V2Event{
			RoutingKey: integrationKey,
			Action:     "trigger",
			Payload: &pagerduty.V2Payload{
				Summary:   summary.Summary[:1024],
				Source:    sourceLink,
				Severity:  "critical",
				Timestamp: time.Now().Format(time.RFC3339),
			},
			// For some reason this isn't typed, so we borrow the struct from ChangeEvent so we get the right JSON
			Links: []any{
				pagerduty.ChangeEventLink{
					Href: sourceLink,
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
		log.Info().Msgf("PDTY | not sending alert '%s' due to pagerduty.enable = false", summary.Summary)
		return SendAlertResponse{}, nil
	}
}
