package environments

import (
	"bytes"
	"encoding/json"
	"errors"
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
		expectedEnvironments []Environment
		expectedError        error
		expectedCode         int
	}{
		{
			name:                 "no existing environments",
			expectedEnvironments: []Environment{},
			expectedCode:         http.StatusOK,
			expectedError:        nil,
		},
		{
			name: "one existing Environment",
			expectedEnvironments: []Environment{
				{
					Name: "test",
				},
			},
			expectedCode:  http.StatusOK,
			expectedError: nil,
		},
		{
			name: "multiple existing Environments",
			expectedEnvironments: []Environment{
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
			expectedEnvironments: []Environment{},
			expectedCode:         http.StatusInternalServerError,
			expectedError:        errors.New("some internal error"),
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			controller, mock := setupMockController(testCase.expectedEnvironments, testCase.expectedError, "listAll")

			context, response := setupTestContext()

			controller.getEnvironments(context)

			mock.AssertCalled(t, "listAll")
			assert.Equal(t, testCase.expectedCode, response.Code)

			responseMeetsExpectations(t, testCase.expectedEnvironments, testCase.expectedError, response.Body)
		})
	}
}

func TestGetEnvironmentByName(t *testing.T) {
	testCases := []struct {
		name                 string
		expectedEnvironments []Environment
		expectedError        error
		expectedCode         int
		environmentName      string
	}{
		{
			name: "successfully get exitsting environment",
			expectedEnvironments: []Environment{
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
			expectedEnvironments: []Environment{},
			expectedError:        ErrEnvironmentNotFound,
			expectedCode:         http.StatusNotFound,
			environmentName:      "fake",
		},
		{
			name:                 "internal server error",
			expectedEnvironments: []Environment{},
			expectedError:        errors.New("some internal error"),
			expectedCode:         http.StatusInternalServerError,
			environmentName:      "testing",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			controller, mock := setupMockController(
				testCase.expectedEnvironments,
				testCase.expectedError, "getByName",
				testCase.environmentName,
			)

			context, response := setupTestContext()
			context.Params = append(context.Params, gin.Param{
				Key:   "name",
				Value: testCase.environmentName,
			})

			controller.getEnvironmentByName(context)

			mock.AssertCalled(t, "getByName", testCase.environmentName)
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
		expectedEnvironments []Environment
		createRequest        CreateEnvironmentRequest
	}{
		{
			name:          "successful create env",
			expectedError: nil,
			expectedCode:  http.StatusCreated,
			expectedEnvironments: []Environment{
				{
					ID:        0,
					Name:      "test",
					CreatedAt: time.Now(),
				},
			},
			createRequest: CreateEnvironmentRequest{
				Name: "test",
			},
		},
		{
			name:                 "empty create request",
			expectedError:        ErrBadCreateRequest,
			expectedCode:         http.StatusBadRequest,
			expectedEnvironments: []Environment{},
			createRequest:        CreateEnvironmentRequest{},
		},
		{
			name:                 "empty environment name",
			expectedError:        ErrBadCreateRequest,
			expectedCode:         http.StatusBadRequest,
			expectedEnvironments: []Environment{},
			createRequest: CreateEnvironmentRequest{
				Name: "",
			},
		},
		{
			name:                 "internal error",
			expectedError:        errors.New("some internal error"),
			expectedCode:         http.StatusInternalServerError,
			expectedEnvironments: []Environment{},
			createRequest: CreateEnvironmentRequest{
				Name: "test",
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			controller, mock := setupMockController(testCase.expectedEnvironments, testCase.expectedError, "createNew", testCase.createRequest)

			context, response := setupTestContext()

			buildCreateEnvironmentRequest(t, context, testCase.createRequest)
			controller.createEnvironment(context)

			if testCase.expectedError == ErrBadCreateRequest {
				mock.AssertNotCalled(t, "createNew")
			} else {
				mock.AssertCalled(t, "createNew", testCase.createRequest)
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
func setupMockController(expectedEnvironments []Environment, expectedError error, methodName string, methodArgs ...interface{}) (*EnvironmentController, *MockEnvironmentStore) {
	mockStore := new(MockEnvironmentStore)
	if methodName == "listAll" {
		mockStore.On(methodName, methodArgs...).Return(expectedEnvironments, expectedError)
	} else {
		if len(expectedEnvironments) < 1 {
			mockStore.On(methodName, methodArgs...).Return(Environment{}, expectedError)
		} else {
			mockStore.On(methodName, methodArgs...).Return(expectedEnvironments[0], expectedError)
		}
	}
	return NewMockController(mockStore), mockStore
}

func buildCreateEnvironmentRequest(t *testing.T, c *gin.Context, createRequest CreateEnvironmentRequest) {
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
func decodeResponseBody(t *testing.T, body io.Reader) Response {
	t.Helper()

	var response Response
	if err := json.NewDecoder(body).Decode(&response); err != nil {
		t.Fatalf("error decoding listAll response body: %v", err)
	}
	return response
}

// responseMeetes Expectations takes a list of Expected Environment structs and an http response body
// It then decodes the response body and checks that it matches the epected result
func responseMeetsExpectations(t *testing.T, expectedEnvironments []Environment, expectedError error, gotBody io.Reader) {
	t.Helper()

	gotResponse := decodeResponseBody(t, gotBody)

	var expectedResponse Response
	if expectedError != nil {
		expectedResponse = Response{Error: expectedError.Error()}
	} else {
		expectationSerializer := EnvironmentsSerializer{expectedEnvironments}
		expectedResponse = Response{Environments: expectationSerializer.Response()}
	}

	if diff := cmp.Diff(gotResponse, expectedResponse); diff != "" {
		t.Errorf("unexpected difference in response body: \n%v\n", diff)
	}
}
