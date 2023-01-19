package pagerduty

import (
	"context"
	"github.com/PagerDuty/go-pagerduty"
	"github.com/broadinstitute/sherlock/internal/config"
	"github.com/rs/zerolog/log"
	"time"
)

func SendChange(integrationKey string, summary string, sourceLink string) error {
	if config.Config.Bool("pagerduty.enable") {
		ctx := context.Background()
		log.Info().Msgf("PDTY | sending change '%s' due to pagerduty.enable = true", summary)
		event := pagerduty.ChangeEvent{
			RoutingKey: integrationKey,
			Payload: pagerduty.ChangeEventPayload{
				Summary:   summary[:1024],
				Source:    sourceLink,
				Timestamp: time.Now().Format(time.RFC3339),
			},
			Links: []pagerduty.ChangeEventLink{
				{
					Href: sourceLink,
					Text: "Beehive Link",
				},
			},
		}
		_, err := client.CreateChangeEventWithContext(ctx, event)

		recordMetrics(ctx, "change", err)

		return err
	} else {
		log.Info().Msgf("PDTY | not sending change '%s' due to pagerduty.enable = false", summary)
		return nil
	}
}

func SendChangeSwallowErrors(integrationKey string, summary string, sourceLink string) {
	err := SendChange(integrationKey, summary, sourceLink)
	if err != nil {
		log.Warn().Msgf("PDTY | error sending change: %v", err)
	}
}
