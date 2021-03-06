// environments_handlers_test.go contains a number of test which verify environmment controller behavior
// without requiring an actual postgres instance to connect to

package v1handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/broadinstitute/sherlock/internal/controllers/v1controllers"
	"github.com/broadinstitute/sherlock/internal/models/v1models"
	"github.com/broadinstitute/sherlock/internal/serializers/v1serializers"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestListAllEnvironments(t *testing.T) {
	testCases := []struct {
		name                 string
		expectedEnvironments []v1models.Environment
		expectedError        error
		expectedCode         int
	}{
		{
			name:                 "no existing environments",
			expectedEnvironments: []v1models.Environment{},
			expectedCode:         http.StatusOK,
			expectedError:        nil,
		},
		{
			name: "one existing Environment",
			expectedEnvironments: []v1models.Environment{
				{
					Name: "test",
				},
			},
			expectedCode:  http.StatusOK,
			expectedError: nil,
		},
		{
			name: "multiple existing Environments",
			expectedEnvironments: []v1models.Environment{
				{
					Name: "dev",
				},
				{
					Name: "alpha",
				},
				{
					Name: "prod",
				},
			},
			expectedCode:  http.StatusOK,
			expectedError: nil,
		},
		{
			name:                 "internal error",
			expectedEnvironments: []v1models.Environment{},
			expectedCode:         http.StatusInternalServerError,
			expectedError:        errors.New("some internal error"),
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			controller, mock := setupMockController(testCase.expectedEnvironments, testCase.expectedError, "ListAll")

			context, response := setupTestContext()

			getEnvironments(controller)(context)

			mock.AssertCalled(t, "ListAll")
			assert.Equal(t, testCase.expectedCode, response.Code)

			responseMeetsExpectations(t, testCase.expectedEnvironments, testCase.expectedError, response.Body)
		})
	}
}

func TestGetEnvironmentByName(t *testing.T) {
	testCases := []struct {
		name                 string
		expectedEnvironments []v1models.Environment
		expectedError        error
		expectedCode         int
		environmentName      string
	}{
		{
			name: "successfully get exitsting environment",
			expectedEnvironments: []v1models.Environment{
				{
					Name:      "test",
					ID:        1,
					CreatedAt: time.Now(),
				},
			},
			expectedError:   nil,
			expectedCode:    http.StatusOK,
			environmentName: "test",
		},
		{
			name:                 "non-existent environment",
			expectedEnvironments: []v1models.Environment{},
			expectedError:        v1models.ErrEnvironmentNotFound,
			expectedCode:         http.StatusNotFound,
			environmentName:      "fake",
		},
		{
			name:                 "internal server error",
			expectedEnvironments: []v1models.Environment{},
			expectedError:        errors.New("some internal error"),
			expectedCode:         http.StatusInternalServerError,
			environmentName:      "testing",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			controller, mock := setupMockController(
				testCase.expectedEnvironments,
				testCase.expectedError, "GetByName",
				testCase.environmentName,
			)

			context, response := setupTestContext()
			context.Params = append(context.Params, gin.Param{
				Key:   "name",
				Value: testCase.environmentName,
			})

			getEnvironmentByName(controller)(context)

			mock.AssertCalled(t, "GetByName", testCase.environmentName)
			assert.Equal(t, testCase.expectedCode, response.Code)

			responseMeetsExpectations(t, testCase.expectedEnvironments, testCase.expectedError, response.Body)
		})
	}
}

func TestCreateEnvironment(t *testing.T) {
	testCases := []struct {
		name                 string
		expectedError        error
		expectedCode         int
		expectedEnvironments []v1models.Environment
		createRequest        v1models.CreateEnvironmentRequest
	}{
		{
			name:          "successful create env",
			expectedError: nil,
			expectedCode:  http.StatusCreated,
			expectedEnvironments: []v1models.Environment{
				{
					ID:        0,
					Name:      "test",
					CreatedAt: time.Now(),
				},
			},
			createRequest: v1models.CreateEnvironmentRequest{
				Name: "test",
			},
		},
		{
			name:                 "empty create request",
			expectedError:        ErrBadEnvironmentCreateRequest,
			expectedCode:         http.StatusBadRequest,
			expectedEnvironments: []v1models.Environment{},
			createRequest:        v1models.CreateEnvironmentRequest{},
		},
		{
			name:                 "empty environment name",
			expectedError:        ErrBadEnvironmentCreateRequest,
			expectedCode:         http.StatusBadRequest,
			expectedEnvironments: []v1models.Environment{},
			createRequest: v1models.CreateEnvironmentRequest{
				Name: "",
			},
		},
		{
			name:                 "internal error",
			expectedError:        errors.New("some internal error"),
			expectedCode:         http.StatusInternalServerError,
			expectedEnvironments: []v1models.Environment{},
			createRequest: v1models.CreateEnvironmentRequest{
				Name: "test",
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			controller, mock := setupMockController(testCase.expectedEnvironments, testCase.expectedError, "CreateNew", testCase.createRequest)

			context, response := setupTestContext()

			buildCreateEnvironmentRequest(t, context, testCase.createRequest)
			createEnvironment(controller)(context)

			if testCase.expectedError == ErrBadEnvironmentCreateRequest {
				mock.AssertNotCalled(t, "CreateNew")
			} else {
				mock.AssertCalled(t, "CreateNew", testCase.createRequest)
			}

			assert.Equal(t, testCase.expectedCode, response.Code)

			responseMeetsExpectations(t, testCase.expectedEnvironments, testCase.expectedError, response.Body)
		})
	}
}

// setupTestContext creates a gin.Context for use in setting up test request
// and a ResponseRecorder to inspect response contents
func setupTestContext() (*gin.Context, *httptest.ResponseRecorder) {
	response := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(response)

	return c, response
}

// setupMockController return an EnvironmentController with the internal store mocked with the desired behavior passed as expectedEnvironments and expectedError
func setupMockController(expectedEnvironments []v1models.Environment, expectedError error, methodName string, methodArgs ...interface{}) (*v1controllers.EnvironmentController, *v1controllers.MockEnvironmentStore) {
	mockStore := new(v1controllers.MockEnvironmentStore)
	if methodName == "ListAll" {
		mockStore.On(methodName, methodArgs...).Return(expectedEnvironments, expectedError)
	} else {
		if len(expectedEnvironments) < 1 {
			mockStore.On(methodName, methodArgs...).Return(v1models.Environment{}, expectedError)
		} else {
			mockStore.On(methodName, methodArgs...).Return(expectedEnvironments[0], expectedError)
		}
	}
	return v1controllers.NewEnvironmentMockController(mockStore), mockStore
}

func buildCreateEnvironmentRequest(t *testing.T, c *gin.Context, createRequest v1models.CreateEnvironmentRequest) {
	t.Helper()

	reqBody := new(bytes.Buffer)
	if err := json.NewEncoder(reqBody).Encode(createRequest); err != nil {
		t.Fatalf("error building create environment request body: %v", err)
	}

	req, err := http.NewRequest(http.MethodPost, "/environments", reqBody)
	if err != nil {
		t.Fatalf("error building create environment request: %v", err)
	}

	c.Request = req
}

// accepts an io.Reader type and attempts to json decode it into an environments.Response
func decodeResponseBody(t *testing.T, body io.Reader) v1controllers.Response {
	t.Helper()

	var response v1controllers.Response
	if err := json.NewDecoder(body).Decode(&response); err != nil {
		t.Fatalf("error decoding listAll response body: %v", err)
	}
	return response
}

// responseMeetes Expectations takes a list of Expected Environment structs and an http response body
// It then decodes the response body and checks that it matches the epected result
func responseMeetsExpectations(t *testing.T, expectedEnvironments []v1models.Environment, expectedError error, gotBody io.Reader) {
	t.Helper()

	gotResponse := decodeResponseBody(t, gotBody)

	var expectedResponse v1controllers.Response
	if expectedError != nil {
		expectedResponse = v1controllers.Response{Error: expectedError.Error()}
	} else {
		expectationSerializer := v1serializers.EnvironmentsSerializer{Environments: expectedEnvironments}
		expectedResponse = v1controllers.Response{Environments: expectationSerializer.Response()}
	}

	if diff := cmp.Diff(gotResponse, expectedResponse); diff != "" {
		t.Errorf("unexpected difference in response body: \n%v\n", diff)
	}
}
