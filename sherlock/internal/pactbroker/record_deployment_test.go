package pactbroker

import (
	"encoding/json"
	"fmt"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/broadinstitute/sherlock/sherlock/internal/pactbroker/pactbroker_mocks"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"io"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
)

func BadRequestError(message string) *ErrorResponse {
	return &ErrorResponse{Errors: []string{message}}
}

func NotFoundError(message string) *ErrorResponse {
	return &ErrorResponse{Errors: []string{message}}
}

// NonEmptyStringMatcher is a custom matcher function that matches non-empty strings.
func NonEmptyStringMatcher(v interface{}) bool {
	str, ok := v.(string)
	return ok && str != ""
}

// UUIDMatcher is a custom matcher function that matches any UUID.
func UUIDMatcher(v interface{}) bool {
	_, ok := v.(uuid.UUID)
	return ok
}

func handler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	// Compile the regex patterns
	reChartName := regexp.MustCompile(`/pacticipants/([\w-]+)/versions/`)
	reVersion := regexp.MustCompile(`versions/([\w.-]+)/deployed-versions/environment/`)
	reEnv := regexp.MustCompile(`environment/([0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12})`)

	// Match the patterns against the request path
	mChartName := reChartName.FindStringSubmatch(path)
	mAppVersion := reVersion.FindStringSubmatch(path)
	mEID := reEnv.FindStringSubmatch(path)

	w.Header().Set("Content-Type", "application/hal+json;charset=utf-8")
	if len(mChartName) == 0 || len(mAppVersion) == 0 || len(mEID) == 0 {
		fmt.Println("Bad request")
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(BadRequestError("Bad request"))
		return
	} else if mEID[1] == uuid.Nil.String() {
		fmt.Println("Environment ID " + mEID[1] + " is not a valid UUID")
		w.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(w).Encode(NotFoundError("Not found"))
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func Test_RecordDeployment(t *testing.T) {
	// Define a configuration function for the mock object
	c := func(m *pactbroker_mocks.MockMockablePactBroker) {
		// Set expectations for the RecordDeployment method
		m.EXPECT().
			RecordDeployment(
				mock.AnythingOfType("string"),
				mock.AnythingOfType("string"),
				mock.AnythingOfType("uuid.UUID")).Run(
			func(chartName string, appVersion string, eID uuid.UUID) {
				fmt.Println(config.Config.Bool("pactbroker.enable"))
				// Create a mock HTTP request with the desired URL path
				request := httptest.NewRequest("POST", "/pacticipants/"+chartName+"/versions/"+appVersion+"/deployed-versions/environment/"+eID.String(), nil)
				request.Header.Set("Content-Type", "application/json; charset=utf-8")
				request.Header.Set("Accept", "application/hal+json")
				request.SetBasicAuth("username", "password")

				// Create a mock HTTP response writer
				w := httptest.NewRecorder()

				// Call the handler function with the mock request and response writer
				handler(w, request)

				// Retrieve the response from the mock response writer
				response := w.Result()

				// Read the response body
				_, err := io.ReadAll(response.Body)
				if err != nil {
					swallowError(err)
					return
				}
				if response.StatusCode == 201 {
					fmt.Printf("deployment for %s app version %s recorded to pact successfully (return code %d). URL: %s",
						chartName, appVersion, response.StatusCode, request.URL.String())
				} else {
					swallowError(fmt.Errorf("deployment for %s app version %s was not recorded to pact successfully (return code %d). URL: %s",
						chartName, appVersion, response.StatusCode, request.URL.String()))
				}
			}).Return()
	}

	// Use the mocked PactBroker API in the test case
	t.Run("invalid UUID", func(t *testing.T) {
		// assert.Contains(t, reflect.TypeOf(pactbroker).String(), "pactBrokerImpl")
		UseMockedPactBroker(t, c, func() {
			// assert.Contains(t, reflect.TypeOf(pactbroker).String(), "mock")
			// Call the RecordDeployment method with the mocked PactBroker API
			uuid00, err := uuid.Parse("00000000-0000-0000-0000-000000000000")
			if err != nil {
				return
			}
			pactbroker.RecordDeployment("anyChart", "v0.0.8", uuid00)
		})
		// assert.Contains(t, reflect.TypeOf(pactbroker).String(), "pactBrokerImpl")
	})

	t.Run("malformed appVersion", func(t *testing.T) {
		// assert.Contains(t, reflect.TypeOf(pactbroker).String(), "pactBrokerImpl")
		UseMockedPactBroker(t, c, func() {
			// assert.Contains(t, reflect.TypeOf(pactbroker).String(), "mock")
			// Call the RecordDeployment method with the mocked PactBroker API
			pactbroker.RecordDeployment("anyChart", "", uuid.New())
		})
		// assert.Contains(t, reflect.TypeOf(pactbroker).String(), "pactBrokerImpl")
	})

	t.Run("values ok", func(t *testing.T) {
		// assert.Contains(t, reflect.TypeOf(pactbroker).String(), "pactBrokerImpl")
		UseMockedPactBroker(t, c, func() {
			// assert.Contains(t, reflect.TypeOf(pactbroker).String(), "mock")
			// Call the RecordDeployment method with the mocked PactBroker API
			pactbroker.RecordDeployment("anyChart", "v0.0.8", uuid.New())
		})
		// assert.Contains(t, reflect.TypeOf(pactbroker).String(), "pactBrokerImpl")
	})
}
