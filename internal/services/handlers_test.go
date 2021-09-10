package services

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockServiceStore struct {
	mock.Mock
}

func TestListServicesHandler(t *testing.T) {
	testCases := []struct {
		name             string
		expectedServices []*Service
		expectedError    error
		expectedCode     int
	}{
		{
			name:             "no existing services",
			expectedServices: []*Service{},
			expectedCode:     http.StatusOK,
			expectedError:    nil,
		},
		{
			name: "one existing service",
			expectedServices: []*Service{
				{
					Name:    "test",
					RepoURL: "http://test.repo",
				},
			},
			expectedCode:  http.StatusOK,
			expectedError: nil,
		},
		{
			name: "multiple existing services",
			expectedServices: []*Service{
				{
					Name:    "cromwell",
					RepoURL: "https://github.com/broadinstitute/cromwell",
				},
				{
					Name:    "leonardo",
					RepoURL: "https://github.com/DataBiosphere/leonardo",
				},
				{
					Name:    "workspacemanager",
					RepoURL: "https://github.com/DataBiosphere/terra-workspace-manager",
				},
			},
			expectedCode:  http.StatusOK,
			expectedError: nil,
		},
		{
			name:             "internal error",
			expectedServices: []*Service{},
			expectedCode:     http.StatusInternalServerError,
			expectedError:    errors.New("some internal error"),
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// setup mock
			mockStore := new(mockServiceStore)
			mockStore.On("listAll").Return(testCase.expectedServices, testCase.expectedError)
			controller := ServiceController{store: mockStore}

			// setup response recorder and request response := httptest.NewRecorder()
			response := httptest.NewRecorder()
			gin.SetMode(gin.TestMode)
			c, _ := gin.CreateTestContext(response)

			controller.getServices(c)

			mockStore.AssertCalled(t, "listAll")
			assert.Equal(t, testCase.expectedCode, response.Code)

			var gotResponse Response
			if err := json.NewDecoder(response.Body).Decode(&gotResponse); err != nil {
				t.Fatalf("error decoding response body: %v\n", err)
			}

			var expectedResponse Response
			if testCase.expectedError != nil {
				expectedResponse = Response{Error: testCase.expectedError.Error()}
			} else {
				expectedResponse = Response{Services: testCase.expectedServices}
			}

			if diff := cmp.Diff(gotResponse, expectedResponse); diff != "" {
				t.Errorf("unexpected difference in response body: \n%v\n", diff)
			}
		})
	}
}

// this is boilerplate code for the testify mock library
func (m *mockServiceStore) listAll() ([]*Service, error) {
	retVal := m.Called()
	return retVal.Get(0).([]*Service), retVal.Error(1)
}

func (m *mockServiceStore) createNew(newService CreateServiceRequest) (*Service, error) {
	retVal := m.Called(newService)
	return retVal.Get(0).(*Service), retVal.Error(1)
}

func (m *mockServiceStore) getByID(id string) (*Service, error) {
	retVal := m.Called(id)
	return retVal.Get(0).(*Service), retVal.Error(1)
}
