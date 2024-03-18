package pactbroker

import (
	"fmt"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/broadinstitute/sherlock/sherlock/internal/pactbroker/pactbroker_mocks"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"net/http"
	"testing"
)

// PactBrokerErrorResponse is a struct that represents the response from the Pact Broker
type PactBrokerErrorResponse struct {
	Errors []string `json:"errors"`
}

type mockablePactBroker interface {
	RecordDeployment(chartName string, appVersion string, eid uuid.UUID)
}

type PactBrokerImpl struct{}

func (p PactBrokerImpl) RecordDeployment(chartName string, appVersion string, eid uuid.UUID) {
	RecordDeployment(chartName, appVersion, eid)
}

var (
	pactbroker mockablePactBroker
)

func init() {
	pactbroker = &PactBrokerImpl{}
}

// Record deployment to pact broker
// https://docs.pact.io/pact_broker/recording_deployments_and_releases
func RecordDeployment(chartName string, appVersion string, eID uuid.UUID) {
	if chartName == "" || appVersion == "" || eID == uuid.Nil {
		return
	}
	if config.Config.Bool("pactbroker.enable") {
		request, err := http.NewRequest(http.MethodPost, config.Config.MustString("pactbroker.url")+"/pacticipants/"+chartName+
			"/versions/"+appVersion+"/deployed-versions/environment/"+eID.String(), nil)
		if err != nil {
			swallowError(err)
			return
		}
		request.Header.Set("Content-Type", "application/json; charset=utf-8")
		request.Header.Set("Accept", "application/hal+json")
		request.SetBasicAuth(config.Config.String("pactbroker.auth.username"), config.Config.String("pactbroker.auth.password"))
		// send the request
		client := &http.Client{}
		response, err := client.Do(request)
		if err != nil {
			swallowError(err)
			return
		}
		if response.StatusCode != 201 {
			swallowError(fmt.Errorf("deployment for %s app version %s was not recorded to pact successfully (return code %d). URL: %s",
				chartName, appVersion, response.StatusCode, response.Request.URL.String()))
			return
		}
	}
}

func swallowError(err error) {
	log.Warn().Msgf("PACT | %v", err)
}

func UseMockedPactBroker(t *testing.T, config func(c *pactbroker_mocks.MockMockablePactBroker), callback func()) {
	if config == nil {
		callback()
		return
	}
	c := pactbroker_mocks.NewMockMockablePactBroker(t)
	config(c)
	temp := pactbroker
	pactbroker = c
	callback()
	pactbroker = temp
}
