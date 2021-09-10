package services

import (
	"bytes"
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

func TestListServices(t *testing.T) {
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

func TestGetServiceByID(t *testing.T) {
	testCases := []struct {
		name            string
		id              string
		expectedService *Service
		expectedError   error
		expectedCode    int
	}{
		{
			name: "successful get by id",
			expectedService: &Service{
				Name:    "tester",
				RepoURL: "https://test.repo",
				ID:      1,
			},
			expectedError: nil,
			expectedCode:  http.StatusOK,
			id:            "1",
		},
		{
			name:            "id not found",
			expectedService: nil,
			expectedCode:    http.StatusNotFound,
			expectedError:   ErrServiceNotFound,
			id:              "1",
		},
		{
			name:            "internal error",
			expectedService: nil,
			expectedCode:    http.StatusInternalServerError,
			expectedError:   errors.New("some internal error"),
			id:              "2",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// setup mock

			mockStore := new(mockServiceStore)
			mockStore.On("getByID", testCase.id).Return(testCase.expectedService, testCase.expectedError)
			controller := ServiceController{store: mockStore}

			response := httptest.NewRecorder()

			gin.SetMode(gin.TestMode)
			c, _ := gin.CreateTestContext(response)
			c.Params = []gin.Param{
				{
					Key:   "id",
					Value: testCase.id,
				},
			}

			controller.getServiceByID(c)
			mockStore.AssertCalled(t, "getByID", testCase.id)

			assert.Equal(t, testCase.expectedCode, response.Code)

			var gotResponse Response
			if err := json.NewDecoder(response.Body).Decode(&gotResponse); err != nil {
				t.Fatalf("error decoding response body: %v\n", err)
			}

			var expectedResponse Response
			if testCase.expectedError != nil {
				expectedResponse = Response{Error: testCase.expectedError.Error()}
			} else {
				expectedResponse = Response{Services: []*Service{testCase.expectedService}}
			}

			if diff := cmp.Diff(gotResponse, expectedResponse); diff != "" {
				t.Errorf("unexpected difference in response body: \n%v\n", diff)
			}
		})
	}
}

func TestCreateService(t *testing.T) {
	testCases := []struct {
		name            string
		expectedError   error
		expectedCode    int
		expectedService *Service
		createRequest   CreateServiceRequest
	}{
		{
			name:          "successful create",
			expectedError: nil,
			expectedCode:  http.StatusCreated,
			createRequest: CreateServiceRequest{
				Name:    "tester",
				RepoURL: "https://test.repo",
			},
			expectedService: &Service{
				ID:      1,
				Name:    "tester",
				RepoURL: "https://test.repo",
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			mockStore := new(mockServiceStore)
			mockStore.On("createNew", testCase.createRequest).Return(testCase.expectedService, testCase.expectedError)
			controller := ServiceController{store: mockStore}

			response := httptest.NewRecorder()

			gin.SetMode(gin.TestMode)
			c, _ := gin.CreateTestContext(response)

			reqBody := new(bytes.Buffer)
			if err := json.NewEncoder(reqBody).Encode(testCase.createRequest); err != nil {
				t.Fatalf("error encoding create service request body: %v", err)
			}

			req, err := http.NewRequest(http.MethodPost, "/services", reqBody)
			if err != nil {
				t.Errorf("error building create service request: %v", err)
			}
			c.Request = req

			controller.createService(c)
			mockStore.AssertCalled(t, "createNew", testCase.createRequest)

			assert.Equal(t, testCase.expectedCode, response.Code)

			var gotResponse Response
			if err := json.NewDecoder(response.Body).Decode(&gotResponse); err != nil {
				t.Fatalf("error decoding response body: %v\n", err)
			}

			var expectedResponse Response
			if testCase.expectedError != nil {
				expectedResponse = Response{Error: testCase.expectedError.Error()}
			} else {
				expectedResponse = Response{Services: []*Service{testCase.expectedService}}
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
