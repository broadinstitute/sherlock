package sherlock

import (
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
