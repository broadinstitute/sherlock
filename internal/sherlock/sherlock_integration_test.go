package sherlock_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/broadinstitute/sherlock/internal/builds"
	"github.com/broadinstitute/sherlock/internal/db"
	"github.com/broadinstitute/sherlock/internal/services"
	"github.com/broadinstitute/sherlock/internal/sherlock"
	"github.com/broadinstitute/sherlock/internal/tools"
	"github.com/gin-gonic/gin"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

// exposes a common sherlock instance that can be shared in integration tests
var app *sherlock.Application

// This integration test pattern is taken from https://www.ardanlabs.com/blog/2019/10/integration-testing-in-go-set-up-and-writing-tests.html

func Test_sherlockServerIntegration(t *testing.T) {
	// performs integration setup when -short flag is not supplied to go test
	integrationSetup(t)
	gin.SetMode(gin.TestMode)

	t.Run("GET /services integration test", func(t *testing.T) {
		// ensure db cleanup will always run at end of test
		defer func() {
			if err := tools.Truncate(app.DB); err != nil {
				t.Errorf("error truncating db in test run: %v", err)
			}
		}()

		// seed test db with sample data. seeded data is also returned
		// for ease of testing
		expectedServices, err := tools.SeedServices(app.DB)
		if err != nil {
			t.Fatalf("error seeding services: %v", err)
		}

		expectedServicesResponse := &services.Response{Services: expectedServices}

		req, err := http.NewRequest(http.MethodGet, "/services", nil)
		if err != nil {
			t.Errorf("error creating request: %v", err)
		}

		response := httptest.NewRecorder()

		app.ServeHTTP(response, req)

		// verify status code is 200
		if response.Code != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, response.Code)
		}

		// decode the resonse body into a slice of Services
		servicesResponse := &services.Response{}
		if err := json.NewDecoder(response.Body).Decode(servicesResponse); err != nil {
			t.Errorf("error decoding response body: %v", err)
		}

		// pretty prints a diff of 2 arbitrary structs
		if diff := cmp.Diff(expectedServicesResponse, servicesResponse); diff != "" {
			t.Errorf("unexpected difference in response body:\n%v", diff)
		}
	})

	t.Run("POST /services integration test", func(t *testing.T) {
		defer func() {
			if err := tools.Truncate(app.DB); err != nil {
				t.Errorf("error truncating db in test run: %v", err)
			}
		}()

		// declare a new service to be included in an http post request body
		newService := &services.CreateServiceRequest{
			Name:    "agora",
			RepoURL: "https://github.com/broadinstitute/agora",
		}

		payload := new(bytes.Buffer)
		if err := json.NewEncoder(payload).Encode(newService); err != nil {
			t.Errorf("error encoding post payload: %v", err)
		}

		req, err := http.NewRequest(http.MethodPost, "/services", payload)
		if err != nil {
			t.Errorf("error generating test request: %v", err)
		}

		response := httptest.NewRecorder()

		app.ServeHTTP(response, req)

		if response.Code != http.StatusCreated {
			t.Errorf("Expected status code %d, got %d", http.StatusCreated, response.Code)
		}

		var savedServiceResponse services.Response
		if err := json.NewDecoder(response.Body).Decode(&savedServiceResponse); err != nil {
			t.Errorf("error decoding response body: %v", err)
		}

		// use a GET /services request to verify the new entity was persisted
		req, err = http.NewRequest(http.MethodGet, "/services", nil)
		if err != nil {
			t.Errorf("error creating get /services request: %v", err)
		}

		response = httptest.NewRecorder()

		app.ServeHTTP(response, req)

		// decode the resonse body into a slice of Services
		service := services.Response{}
		if err := json.NewDecoder(response.Body).Decode(&service); err != nil {
			t.Errorf("error decoding response body: %v", err)
		}

		// verify the service returned from the list endpoint is the same as that
		// returned in the post request response
		if diff := cmp.Diff(savedServiceResponse, service); diff != "" {
			t.Errorf("unexpected difference in response body:\n%v", diff)
		}
	})
	t.Run("GET /services by ID", func(t *testing.T) {
		defer func() {
			if err := tools.Truncate(app.DB); err != nil {
				t.Errorf("error truncating db in test run: %v", err)
			}
		}()

		expectedServices, err := tools.SeedServices(app.DB)
		if err != nil {
			t.Fatalf("error seeding services: %v", err)
		}
		expectedServicesResponse := &services.Response{Services: []*services.Service{expectedServices[0]}}

		req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/services/%s", expectedServices[0].Name), nil)
		if err != nil {
			t.Errorf("error creating request: %v", err)
		}

		response := httptest.NewRecorder()

		app.ServeHTTP(response, req)

		if response.Code != http.StatusOK {
			t.Errorf("expected response code %d, got %d", http.StatusOK, response.Code)
		}

		var gotService services.Response
		if err := json.NewDecoder(response.Body).Decode(&gotService); err != nil {
			t.Errorf("error decoding reponse body %v", err)
		}

		if diff := cmp.Diff(&gotService, expectedServicesResponse); diff != "" {
			t.Errorf("received unexpected response %v\n", diff)
		}

		// Make sure we get 404 for nonexistent service
		req, err = http.NewRequest(http.MethodGet, fmt.Sprintf("/services/%s", "non-existent"), nil)
		if err != nil {
			t.Errorf("error creating request: %v", err)
		}

		response = httptest.NewRecorder()

		app.ServeHTTP(response, req)

		if response.Code != http.StatusNotFound {
			t.Errorf("Expected to receive code %d for non-existent service, got %d", http.StatusNotFound, response.Code)
		}
	})

	t.Run("GET /builds", func(t *testing.T) {
		defer func() {
			if err := tools.Truncate(app.DB); err != nil {
				t.Errorf("error truncatingdb in test run : %v", err)
			}
		}()

		_, err := tools.SeedServices(app.DB)
		if err != nil {
			t.Fatalf("error seeding services: %v", err)
		}

		expectedBuilds, err := tools.SeedBuilds(app.DB)
		if err != nil {
			t.Fatalf("error seeding builds: %v", err)
		}

		expectedBuildsResponse := &builds.Response{Builds: expectedBuilds}

		req, err := http.NewRequest(http.MethodGet, "/builds", nil)
		if err != nil {
			t.Errorf("error generating test GET /builds request: %v", err)
		}

		response := httptest.NewRecorder()

		app.ServeHTTP(response, req)

		assert.Equal(t, http.StatusOK, response.Code)

		result := &builds.Response{}
		if err := json.NewDecoder(response.Body).Decode(result); err != nil {
			t.Errorf("error decoding response body: %v", err)
		}

		if diff := cmp.Diff(expectedBuildsResponse, result); diff != "" {
			t.Errorf("unexpected difference in response body:\n%v", diff)
		}
	})
}

func integrationSetup(t *testing.T) {
	// skip integration tests if go test is invoked with -short flag
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	// The following steps initialize the database for use in the
	// sherlock server integration test suite
	// TODO pull this from config with viper

	// when running tests workdir is the package directory ie cmd/server
	// so a relative path to changelogs is needed.
	// TODO cleaner method to supply path to changelogs and run migration in tests
	if err := db.ApplyMigrations("../../db/migrations", sherlock.Config); err != nil {
		t.Fatalf("error migrating database: %v", err)
	}

	app = sherlock.New()
}
