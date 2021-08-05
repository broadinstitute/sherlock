package sherlock

import (
	"database/sql"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetServicesFailures(t *testing.T) {
	app := New(&mockDB{})
	req, _ := http.NewRequest(http.MethodGet, "/services", nil)

	response := httptest.NewRecorder()
	app.ServeHTTP(response, req)

	expectedCode := http.StatusInternalServerError

	if response.Code != expectedCode {
		t.Errorf("Expected status code %d, got %d", expectedCode, response.Code)
	}

	body := response.Body.String()
	if !strings.Contains(body, "error") {
		t.Errorf("Expected body to contain an error message and didn't receive one. got %s\n", body)
	}
}

// used to abstract away database interactions in unit tests for route handlers
type mockDB struct {
}

// This is intended to test the failure modes of the get services route. The success
// cases are already exersized in the integration tests
func (m *mockDB) Select(dest interface{}, query string, args ...interface{}) error {
	return fmt.Errorf("Error retrieving services from datastore")
}

// This is just to satisfy the querier interface required for a mock sherlock application
// this method is not actually used in current routes, only in utility helpers
func (m *mockDB) Exec(query string, args ...interface{}) (sql.Result, error) {
	return nil, nil
}

// same as above
func (m *mockDB) Prepare(query string) (*sql.Stmt, error) {
	return nil, nil
}

// Same deal as exec
func (m *mockDB) Close() error {
	return nil
}
