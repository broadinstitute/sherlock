package builds

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestListBuilds(t *testing.T) {
	testCases := []struct {
		name           string
		expectedBuilds []Build
		expectedError  error
		expectedCode   int
	}{
		{
			name:           "no existing builds",
			expectedBuilds: []Build{},
			expectedCode:   http.StatusOK,
			expectedError:  nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// setup mock store
			mockStore := new(mockBuildStore)
			mockStore.On("listAll").Return(testCase.expectedBuilds, testCase.expectedError)

			controller := BuildController{store: mockStore}

			response := httptest.NewRecorder()
			gin.SetMode(gin.TestMode)
			c, _ := gin.CreateTestContext(response)

			controller.getBuilds(c)

			mockStore.AssertCalled(t, "listAll")
			assert.Equal(t, testCase.expectedCode, response.Code)
		})
	}
}

// below is boilerplate code for the testify/mock library

type mockBuildStore struct {
	mock.Mock
}

func (m *mockBuildStore) listAll() ([]Build, error) {
	retVal := m.Called()
	return retVal.Get(0).([]Build), retVal.Error(1)
}
