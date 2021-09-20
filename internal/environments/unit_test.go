package environments

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

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
			controller, mock := setupMockController("listAll", testCase.expectedEnvironments, testCase.expectedError)

			context, response := setupTestContext()

			controller.getEnvironments(context)

			mock.AssertCalled(t, "listAll")
			assert.Equal(t, testCase.expectedCode, response.Code)

			gotResponse := decodeResponseBody(t, response.Body)

			responseMeetsExpectations(t, testCase.expectedEnvironments, testCase.expectedError, gotResponse)
		})
	}
}

func setupTestContext() (*gin.Context, *httptest.ResponseRecorder) {
	response := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(response)

	return c, response
}

func setupMockController(methodName string, expectedEnvironments []Environment, expectedError error) (*EnvironmentController, *MockEnvironmentStore) {
	mockStore := new(MockEnvironmentStore)
	mockStore.On(methodName).Return(expectedEnvironments, expectedError)
	return NewMockController(mockStore), mockStore
}

func decodeResponseBody(t *testing.T, body io.Reader) Response {
	t.Helper()

	var response Response
	if err := json.NewDecoder(body).Decode(&response); err != nil {
		t.Fatalf("error decoding listAll response body: %v", err)
	}
	return response
}

func responseMeetsExpectations(t *testing.T, expectedEnvironments []Environment, expectedError error, got Response) {
	var expectedResponse Response
	if expectedError != nil {
		expectedResponse = Response{Error: expectedError.Error()}
	} else {
		expectationSerializer := EnvironmentsSerializer{expectedEnvironments}
		expectedResponse = Response{Environments: expectationSerializer.Response()}
	}

	if diff := cmp.Diff(got, expectedResponse); diff != "" {
		t.Errorf("unexpected difference in response body: \n%v\n", diff)
	}
}
