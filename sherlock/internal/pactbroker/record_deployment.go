package pactbroker

import (
	"errors"
	"fmt"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"net/http"
)

// Record deployment to pact broker
// https://docs.pact.io/pact_broker/recording_deployments_and_releases
func RecordDeployment(chartName string, appVersion string, eID uuid.UUID) {
	if config.Config.Bool("pactbroker.enabled") {

		request, err := http.NewRequest(http.MethodPost, config.Config.MustString("pactbroker.url")+"/pacticipants/"+chartName+
			"/versions/"+appVersion+"/deployed-versions/environment/"+eID.String(), nil)
		if err != nil {
			swallowErrors(err)
			return
		}
		request.Header.Set("Content-Type", "application/json; charset=utf-8")
		request.Header.Set("Accept", "application/hal+json")
		request.SetBasicAuth(config.Config.String("pactbroker.auth.username"), config.Config.String("pactbroker.auth.password"))
		// send the request
		client := &http.Client{}
		response, err := client.Do(request)
		if err != nil {
			PactSwallowErrors(err)
		}
		if response.StatusCode != 201 {
			swallowError(fmt.Errorf("deployment for %s app version %s was not recorded to pact successfully (return code %d)", 
				chartName, appVersion, response.StatusCode))
			return
		}
	}
}

func swallowError(err error) {
	log.Warn().Msgf("PACT | %v", err)
}
