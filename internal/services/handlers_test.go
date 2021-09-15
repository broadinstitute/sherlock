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
)

func TestListServices(t *testing.T) {
	testCases := []struct {
		name             string
		expectedServices []Service
		expectedError    error
		expectedCode     int
	}{
		{
			name:             "no existing services",
			expectedServices: []Service{},
			expectedCode:     http.StatusOK,
			expectedError:    nil,
		},
		{
			name: "one existing service",
			expectedServices: []Service{
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
			expectedServices: []Service{
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
			expectedServices: []Service{},
			expectedCode:     http.StatusInternalServerError,
			expectedError:    errors.New("some internal error"),
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// setup mock
			mockStore := new(MockServiceStore)
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

			// serialize the expectations using to get expected json response data

			var expectedResponse Response
			if testCase.expectedError != nil {
				expectedResponse = Response{Error: testCase.expectedError.Error()}
			} else {
				expectationSerializer := ServicesSerializer{testCase.expectedServices}
				expectedServices := expectationSerializer.Response()
				expectedResponse = Response{Services: expectedServices}
			}

			if diff := cmp.Diff(gotResponse, expectedResponse); diff != "" {
				t.Errorf("unexpected difference in response body: \n%v\n", diff)
			}
		})
	}
}

func TestGetServiceByName(t *testing.T) {
	testCases := []struct {
		name            string
		expectedService *Service
		expectedError   error
		expectedCode    int
		serviceName     string
	}{
		{
			name: "successful get by name",
			expectedService: &Service{
				Name:    "tester",
				RepoURL: "https://test.repo",
				ID:      1,
			},
			expectedError: nil,
			expectedCode:  http.StatusOK,
			serviceName:   "tester",
		},
		{
			name:            "name not found",
			expectedService: nil,
			expectedCode:    http.StatusNotFound,
			expectedError:   ErrServiceNotFound,
			serviceName:     "blah",
		},
		{
			name:            "internal error",
			expectedService: nil,
			expectedCode:    http.StatusInternalServerError,
			expectedError:   errors.New("some internal error"),
			serviceName:     "test-service",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// setup mock

			mockStore := new(MockServiceStore)
			mockStore.On("getByName", testCase.serviceName).Return(testCase.expectedService, testCase.expectedError)
			controller := ServiceController{store: mockStore}

			response := httptest.NewRecorder()

			gin.SetMode(gin.TestMode)
			c, _ := gin.CreateTestContext(response)
			c.Params = []gin.Param{
				{
					Key:   "name",
					Value: testCase.serviceName,
				},
			}

			controller.getServiceByName(c)
			mockStore.AssertCalled(t, "getByName", testCase.serviceName)

			assert.Equal(t, testCase.expectedCode, response.Code)

			var gotResponse Response
			if err := json.NewDecoder(response.Body).Decode(&gotResponse); err != nil {
				t.Fatalf("error decoding response body: %v\n", err)
			}

			var expectedResponse Response
			if testCase.expectedError != nil {
				expectedResponse = Response{Error: testCase.expectedError.Error()}
			} else {
				expectationSerializer := ServiceSerializer{*testCase.expectedService}
				expectedService := expectationSerializer.response()
				expectedResponse = Response{Services: []ServiceResponse{expectedService}}
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
				Name:    "tester",
				RepoURL: "https://test.repo",
			},
		},
		{
			name:          "missing service name",
			expectedError: ErrBadCreateRequest,
			expectedCode:  http.StatusBadRequest,
			createRequest: CreateServiceRequest{
				RepoURL: "https://tester.repo",
			},
			expectedService: nil,
		},
		{
			name:          "missing repo url",
			expectedError: ErrBadCreateRequest,
			expectedCode:  http.StatusBadRequest,
			createRequest: CreateServiceRequest{
				Name: "tester",
			},
			expectedService: nil,
		},
		{
			name:            "empty create request",
			expectedError:   ErrBadCreateRequest,
			expectedCode:    http.StatusBadRequest,
			createRequest:   CreateServiceRequest{},
			expectedService: nil,
		},
		{
			name:          "empty service name",
			expectedError: ErrBadCreateRequest,
			expectedCode:  http.StatusBadRequest,
			createRequest: CreateServiceRequest{
				Name:    "",
				RepoURL: "https://tester.repo",
			},
			expectedService: nil,
		},
		{
			name:          "empty repo url",
			expectedError: ErrBadCreateRequest,
			expectedCode:  http.StatusBadRequest,
			createRequest: CreateServiceRequest{
				Name:    "tester",
				RepoURL: "",
			},
			expectedService: nil,
		},
		{
			name:          "internal error",
			expectedError: errors.New("some internal error"),
			expectedCode:  http.StatusInternalServerError,
			createRequest: CreateServiceRequest{
				Name:    "tester",
				RepoURL: "https://tester.repo",
			},
			expectedService: nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			mockStore := new(MockServiceStore)
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
			if testCase.expectedError == ErrBadCreateRequest {
				mockStore.AssertNotCalled(t, "createNew")
			} else {
				mockStore.AssertCalled(t, "createNew", testCase.createRequest)
			}

			assert.Equal(t, testCase.expectedCode, response.Code)

			var gotResponse Response
			if err := json.NewDecoder(response.Body).Decode(&gotResponse); err != nil {
				t.Fatalf("error decoding response body: %v\n", err)
			}

			var expectedResponse Response
			if testCase.expectedError != nil {
				expectedResponse = Response{Error: testCase.expectedError.Error()}
			} else {
				expectationSerializer := ServiceSerializer{*testCase.expectedService}
				expectedService := expectationSerializer.response()
				expectedResponse = Response{Services: []ServiceResponse{expectedService}}
			}

			if diff := cmp.Diff(gotResponse, expectedResponse); diff != "" {
				t.Errorf("unexpected difference in response body: \n%v\n", diff)
			}
		})
	}
}
