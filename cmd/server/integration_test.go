package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/broadinstitute/sherlock/internal/db"
	"github.com/broadinstitute/sherlock/internal/sherlock"
	"github.com/jmoiron/sqlx"
)

var app *sherlock.Application

// This integration test patter is taken from https://www.ardanlabs.com/blog/2019/10/integration-testing-in-go-set-up-and-writing-tests.html
func TestMain(m *testing.M) {
	os.Exit(testMain(m))
}

func testMain(m *testing.M) int {
	dbConn, err := sqlx.Connect("pgx", os.Getenv("POSTGRESQL_URL"))
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
		return 1
	}
	defer dbConn.Close()

	// when running tests workdir is the package directory ie cmd/server
	// so a relative path to changelogs is needed.
	// TODO cleaner method to supply path to changelogs and run migration in tests
	if err := db.ApplyMigrations("../../db/migrations"); err != nil {
		log.Fatalf("error migrating database: %v", err)
		return 1
	}

	app = sherlock.New(dbConn)

	return m.Run()
}

func Test_getServicesIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	defer func() {
		if err := db.Truncate(app.DB); err != nil {
			t.Errorf("error truncating db in test run: %v", err)
		}
	}()

	_, err := db.SeedServices(app.DB)
	if err != nil {
		t.Fatalf("error seeding services: %v", err)
	}

	req, err := http.NewRequest(http.MethodGet, "/services", nil)
	if err != nil {
		t.Errorf("error creating request: %v", err)
	}

	responseWriter := httptest.NewRecorder()

	app.ServeHTTP(responseWriter, req)

	// verify status code is 200
	if responseWriter.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, responseWriter.Code)
	}
}
