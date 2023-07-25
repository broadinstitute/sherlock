package pactbroker

import (
	"errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"net/http"
)

const pburl = `https://pact-broker.dsp-eng-tools.broadinstitute.org`
const alphaID = "19e0aeee-4fc5-49a7-b502-065e3fbce967"
const productionID = "354e0bfa-9634-417c-b10a-4beea2ffc3bd"

// Record deployment to pact broker
// https://docs.pact.io/pact_broker/recording_deployments_and_releases
func RecordDeployment(chartName string, chartversion string, environmentName string) error {
	var eID string
	if config.Config.Bool("pactbroker.enabled") {

		switch en := environmentName; en {
		case "alpha":
			eID = alphaID
		case "prod":
			eID = productionID
		}
		request, err := http.NewRequest(http.MethodPost, pburl+"/pacticipants/"+chartName+
			"/versions/"+chartversion+"/deployed-versions/environment/"+eID, nil)
		if err != nil {
			return err
		}
		request.Header.Set("Content-Type", "application/json; charset=utf-8")
		request.Header.Set("Accept", "application/hal+json")
		request.SetBasicAuth(config.Config.MustString("pactbroker.auth.username"), config.Config.MustString("pactbroker.auth.password"))
		// send the request
		client := &http.Client{}
		response, err := client.Do(request)
		if err != nil {
			return err
		}
		if response.StatusCode != 201 {
			return errors.New("Deployment was not recorded to pact successfully.")
		}
	}
	return nil
}
