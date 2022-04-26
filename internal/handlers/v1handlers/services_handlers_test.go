package v1handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/broadinstitute/sherlock/internal/controllers/v1controllers"
	"github.com/broadinstitute/sherlock/internal/models/v1models"
	"github.com/broadinstitute/sherlock/internal/serializers/v1serializers"
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
		expectedServices []v1models.Service
		expectedError    error
		expectedCode     int
	}{
		{
			name:             "no existing services",
			expectedServices: []v1models.Service{},
			expectedCode:     http.StatusOK,
			expectedError:    nil,
		},
		{
			name: "one existing service",
			expectedServices: []v1models.Service{
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
			expectedServices: []v1models.Service{
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
			expectedServices: []v1models.Service{},
			expectedCode:     http.StatusInternalServerError,
			expectedError:    errors.New("some internal error"),
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// setup mock
			mockStore := new(v1controllers.MockServiceStore)
			mockStore.On("ListAll").Return(testCase.expectedServices, testCase.expectedError)
			controller := v1controllers.ServiceController{Store: mockStore}

			// setup response recorder and request response := httptest.NewRecorder()
			response := httptest.NewRecorder()
			gin.SetMode(gin.TestMode)
			c, _ := gin.CreateTestContext(response)

			getServices(&controller)(c)

			mockStore.AssertCalled(t, "ListAll")
			assert.Equal(t, testCase.expectedCode, response.Code)

			var gotResponse v1serializers.Response
			if err := json.NewDecoder(response.Body).Decode(&gotResponse); err != nil {
				t.Fatalf("error decoding response body: %v\n", err)
			}

			// serialize the expectations using to get expected json response data

			var expectedResponse v1serializers.Response
			if testCase.expectedError != nil {
				expectedResponse = v1serializers.Response{Error: testCase.expectedError.Error()}
			} else {
				// Serialize the expected service entity to a Service Response type
				expectationSerializer := v1serializers.ServicesSerializer{testCase.expectedServices}
				expectedServices := expectationSerializer.Response()
				expectedResponse = v1serializers.Response{Services: expectedServices}
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
		expectedService v1models.Service
		expectedError   error
		expectedCode    int
		serviceName     string
	}{
		{
			name: "successful get by name",
			expectedService: v1models.Service{
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
			expectedService: v1models.Service{},
			expectedCode:    http.StatusNotFound,
			expectedError:   v1models.ErrServiceNotFound,
			serviceName:     "blah",
		},
		{
			name:            "internal error",
			expectedService: v1models.Service{},
			expectedCode:    http.StatusInternalServerError,
			expectedError:   errors.New("some internal error"),
			serviceName:     "test-service",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// setup mock

			mockStore := new(v1controllers.MockServiceStore)
			mockStore.On("GetByName", testCase.serviceName).Return(testCase.expectedService, testCase.expectedError)
			controller := v1controllers.ServiceController{Store: mockStore}

			response := httptest.NewRecorder()

			gin.SetMode(gin.TestMode)
			c, _ := gin.CreateTestContext(response)
			c.Params = []gin.Param{
				{
					Key:   "name",
					Value: testCase.serviceName,
				},
			}

			getServiceByName(&controller)(c)
			mockStore.AssertCalled(t, "GetByName", testCase.serviceName)

			assert.Equal(t, testCase.expectedCode, response.Code)

			var gotResponse v1serializers.Response
			if err := json.NewDecoder(response.Body).Decode(&gotResponse); err != nil {
				t.Fatalf("error decoding response body: %v\n", err)
			}

			var expectedResponse v1serializers.Response
			if testCase.expectedError != nil {
				expectedResponse = v1serializers.Response{Error: testCase.expectedError.Error()}
			} else {
				expectationSerializer := v1serializers.ServiceSerializer{testCase.expectedService}
				expectedService := expectationSerializer.Response()
				expectedResponse = v1serializers.Response{Services: []v1serializers.ServiceResponse{expectedService}}
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
		expectedService v1models.Service
		createRequest   v1models.CreateServiceRequest
	}{
		{
			name:          "successful create",
			expectedError: nil,
			expectedCode:  http.StatusCreated,
			createRequest: v1models.CreateServiceRequest{
				Name:    "tester",
				RepoURL: "https://test.repo",
			},
			expectedService: v1models.Service{
				Name:    "tester",
				RepoURL: "https://test.repo",
			},
		},
		{
			name:          "missing service name",
			expectedError: ErrBadServiceCreateRequest,
			expectedCode:  http.StatusBadRequest,
			createRequest: v1models.CreateServiceRequest{
				RepoURL: "https://tester.repo",
			},
			expectedService: v1models.Service{},
		},
		{
			name:          "missing repo url",
			expectedError: ErrBadServiceCreateRequest,
			expectedCode:  http.StatusBadRequest,
			createRequest: v1models.CreateServiceRequest{
				Name: "tester",
			},
			expectedService: v1models.Service{},
		},
		{
			name:            "empty create request",
			expectedError:   ErrBadServiceCreateRequest,
			expectedCode:    http.StatusBadRequest,
			createRequest:   v1models.CreateServiceRequest{},
			expectedService: v1models.Service{},
		},
		{
			name:          "empty service name",
			expectedError: ErrBadServiceCreateRequest,
			expectedCode:  http.StatusBadRequest,
			createRequest: v1models.CreateServiceRequest{
				Name:    "",
				RepoURL: "https://tester.repo",
			},
			expectedService: v1models.Service{},
		},
		{
			name:          "empty repo url",
			expectedError: ErrBadServiceCreateRequest,
			expectedCode:  http.StatusBadRequest,
			createRequest: v1models.CreateServiceRequest{
				Name:    "tester",
				RepoURL: "",
			},
			expectedService: v1models.Service{},
		},
		{
			name:          "internal error",
			expectedError: errors.New("some internal error"),
			expectedCode:  http.StatusInternalServerError,
			createRequest: v1models.CreateServiceRequest{
				Name:    "tester",
				RepoURL: "https://tester.repo",
			},
			expectedService: v1models.Service{},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			mockStore := new(v1controllers.MockServiceStore)
			mockStore.On("CreateNew", testCase.createRequest).Return(testCase.expectedService, testCase.expectedError)
			controller := v1controllers.ServiceController{Store: mockStore}

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

			createService(&controller)(c)
			if testCase.expectedError == ErrBadServiceCreateRequest {
				mockStore.AssertNotCalled(t, "CreateNew")
			} else {
				mockStore.AssertCalled(t, "CreateNew", testCase.createRequest)
			}

			assert.Equal(t, testCase.expectedCode, response.Code)

			var gotResponse v1serializers.Response
			if err := json.NewDecoder(response.Body).Decode(&gotResponse); err != nil {
				t.Fatalf("error decoding response body: %v\n", err)
			}

			var expectedResponse v1serializers.Response
			if testCase.expectedError != nil {
				expectedResponse = v1serializers.Response{Error: testCase.expectedError.Error()}
			} else {
				expectationSerializer := v1serializers.ServiceSerializer{testCase.expectedService}
				expectedService := expectationSerializer.Response()
				expectedResponse = v1serializers.Response{Services: []v1serializers.ServiceResponse{expectedService}}
			}

			if diff := cmp.Diff(gotResponse, expectedResponse); diff != "" {
				t.Errorf("unexpected difference in response body: \n%v\n", diff)
			}
		})
	}
}
