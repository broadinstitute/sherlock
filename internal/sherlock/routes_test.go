package sherlock

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/broadinstitute/sherlock/internal/db"
)

func TestGetServicesFailures(t *testing.T) {
	// setup an app instance with mock db that will always fail to verify failure mode behavior
	model, mock := db.NewMockServiceModel(t, true)
	app := &Application{}
	app.buildRouter()
	app.ServiceModel = model

	// ensure mock db will error out on any query
	mock.ExpectQuery(".*").WillReturnError(fmt.Errorf("unable to select all services"))
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
