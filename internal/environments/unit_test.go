package environments

import (
	"encoding/json"
	"errors"
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
			mockStore := new(MockEnvironmentStore)
			mockStore.On("listAll").Return(testCase.expectedEnvironments, testCase.expectedError)
			controller := NewMockController(mockStore)

			response := httptest.NewRecorder()
			gin.SetMode(gin.TestMode)
			c, _ := gin.CreateTestContext(response)

			controller.getEnvironments(c)

			mockStore.AssertCalled(t, "listAll")
			assert.Equal(t, testCase.expectedCode, response.Code)

			var gotResponse Response
			if err := json.NewDecoder(response.Body).Decode(&gotResponse); err != nil {
				t.Fatalf("error decoding listAll response body: %v", err)
			}

			var expectedResponse Response
			if testCase.expectedError != nil {
				expectedResponse = Response{Error: testCase.expectedError.Error()}
			} else {
				expectationSerializer := EnvironmentsSerializer{testCase.expectedEnvironments}
				expectedResponse = Response{Environments: expectationSerializer.Response()}
			}

			if diff := cmp.Diff(gotResponse, expectedResponse); diff != "" {
				t.Errorf("unexpected difference in response body: \n%v\n", diff)
			}
		})
	}
}
