package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/broadinstitute/sherlock/internal/db"
	"github.com/broadinstitute/sherlock/internal/services"
	"github.com/broadinstitute/sherlock/internal/sherlock"
	"github.com/broadinstitute/sherlock/internal/tools"
	"github.com/google/go-cmp/cmp"
	"github.com/jmoiron/sqlx"
)

// expopses a common sherlock instance that can be shared in integration tests
var app *sherlock.Application

// This integration test patter is taken from https://www.ardanlabs.com/blog/2019/10/integration-testing-in-go-set-up-and-writing-tests.html

func Test_sherlockServerIntegration(t *testing.T) {
	// performs integration setup when -short flag is not supplied to go test
	integrationSetup(t)
	// ensure db connection is closed at end of integration tests
	defer app.DB.Close()

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
		services := make([]services.Service, 0)
		if err := json.NewDecoder(response.Body).Decode(&services); err != nil {
			t.Errorf("error decoding response body: %v", err)
		}

		// pretty prints a diff of 2 arbitrary structs
		if diff := cmp.Diff(expectedServices, services); diff != "" {
			t.Errorf("unexpected difference in response body:\n%v", diff)
		}
	})
}

func integrationSetup(t *testing.T) {
	t.Helper()
	// skip integration tests if go test is invoked with -short flag
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	// The following steps initialize the database for use in the
	// sherlock server integration test suite

	dbConn, err := sqlx.Connect("pgx", os.Getenv("POSTGRESQL_URL"))
	if err != nil {
		t.Fatalf("Unable to connect to database: %v\n", err)
	}

	// when running tests workdir is the package directory ie cmd/server
	// so a relative path to changelogs is needed.
	// TODO cleaner method to supply path to changelogs and run migration in tests
	if err := db.ApplyMigrations("../../db/migrations"); err != nil {
		t.Fatalf("error migrating database: %v", err)
	}

	app = sherlock.New(dbConn)
}
