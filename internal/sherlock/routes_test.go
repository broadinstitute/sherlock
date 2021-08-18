package sherlock

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/broadinstitute/sherlock/internal/db"
	"github.com/broadinstitute/sherlock/internal/services"
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

func TestPostServiceDBFailure(t *testing.T) {
	// setup an app instance with mock db that will always fail to verify failure mode behavior
	model, mock := db.NewMockServiceModel(t, true)
	app := &Application{}
	app.buildRouter()
	app.ServiceModel = model

	// simulate an error persisting service to db
	mock.ExpectBegin()
	mock.ExpectQuery(".*").WillReturnError(fmt.Errorf("unable to create service"))
	mock.ExpectRollback()

	newService := &services.Service{
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

	expectedCode := http.StatusInternalServerError

	if response.Code != expectedCode {
		t.Errorf("Expected status code %d, got %d", expectedCode, response.Code)
	}

	body := response.Body.String()
	if !strings.Contains(body, "error") {
		t.Errorf("Expected body to contain an error message and didn't receive one. got %s\n", body)
	}
}
