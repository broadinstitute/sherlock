package builds

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/broadinstitute/sherlock/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/google/go-cmp/cmp"
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
		{
			name: "one existing build",
			expectedBuilds: []Build{
				{
					VersionString: "imagerepo.io/test:0.1.0",
					CommitSha:     "f2f2f23",
					BuildURL:      "https://build.job.out/blah",
					BuiltAt:       time.Now(),
					ServiceID:     1,
					Service: services.Service{
						ID:      1,
						Name:    "tester",
						RepoURL: "https://tester.repo",
					},
				},
			},
			expectedError: nil,
			expectedCode:  http.StatusOK,
		},
		{
			name: "multiple builds and services",
			expectedBuilds: []Build{
				{
					VersionString: "imagerepo.io/test:0.1.0",
					CommitSha:     "f2f2f23",
					BuildURL:      "https://build.job.out/blah",
					BuiltAt:       time.Now(),
					ServiceID:     1,
					Service: services.Service{
						ID:      1,
						Name:    "tester",
						RepoURL: "https://tester.repo",
					},
				},
				{
					VersionString: "imagerepo.io/test:0.2.0",
					CommitSha:     "l2k3j4",
					BuildURL:      "https://build.job.out/blah",
					BuiltAt:       time.Now(),
					ServiceID:     1,
					Service: services.Service{
						ID:      1,
						Name:    "tester",
						RepoURL: "https://tester.repo",
					},
				},
				{
					VersionString: "imagerepo.io/dummyService:0.1.0",
					CommitSha:     "f2f2f23",
					BuildURL:      "https://build.job.out/234",
					BuiltAt:       time.Now(),
					ServiceID:     2,
					Service: services.Service{
						ID:      2,
						Name:    "dummyServcie",
						RepoURL: "https://dummy.repo",
					},
				},
				{
					VersionString: "imagerepo.io/cromwell:1.45.0",
					CommitSha:     "k2j3h43k",
					BuildURL:      "https://build.job.out/2345",
					BuiltAt:       time.Now(),
					ServiceID:     3,
					Service: services.Service{
						ID:      3,
						Name:    "cromwell",
						RepoURL: "https://cromwell.repo",
					},
				},
			},
			expectedError: nil,
			expectedCode:  http.StatusOK,
		},
		{
			name:           "internal error",
			expectedBuilds: []Build{},
			expectedError:  errors.New("some internal error"),
			expectedCode:   http.StatusInternalServerError,
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

			var gotResponse Response
			if err := json.NewDecoder(response.Body).Decode(&gotResponse); err != nil {
				t.Fatalf("error decoding response body: %v\n", err)
			}

			var expectedResponse Response
			if testCase.expectedError != nil {
				expectedResponse = Response{Error: testCase.expectedError.Error()}
			} else {
				expectedResponse = Response{Builds: testCase.expectedBuilds}
			}

			if diff := cmp.Diff(gotResponse, expectedResponse); diff != "" {
				t.Errorf("unexpected difference in response body: \n%v\n", diff)
			}
		})
	}
}

func TestCreateBuild(t *testing.T) {
	testCases := []struct {
		name                    string
		expectedError           error
		expectedCode            int
		createRequest           CreateBuildRequest
		simulateServiceCreation bool
	}{
		{
			name:          "build for existing service",
			expectedError: nil,
			expectedCode:  http.StatusCreated,
			createRequest: CreateBuildRequest{
				VersionString: "gcr.io/broad/cromwell:1.0.0",
				CommitSha:     "lk23j44",
				ServiceName:   "cromwell",
				ServiceRepo:   "github.com/broadinstitute/cromwell",
				BuildURL:      "https://jenkins.job/123",
				BuiltAt:       time.Now(),
			},
			simulateServiceCreation: false,
		},
		{
			name:          "build for new service",
			expectedError: nil,
			expectedCode:  http.StatusCreated,
			createRequest: CreateBuildRequest{
				VersionString: "gcr.io/broad/workspacemanager:0.1.0",
				CommitSha:     "k3l42j",
				ServiceName:   "workspacemanager",
				ServiceRepo:   "github.com/databiosphere/workspacemanager",
				BuildURL:      "https://github.com/workspacemanager/actions/2",
				BuiltAt:       time.Now(),
			},
			simulateServiceCreation: true,
		},
		{
			name:          "missing version string",
			expectedError: ErrBadCreateRequest,
			expectedCode:  http.StatusBadRequest,
			createRequest: CreateBuildRequest{
				CommitSha:   "k2j34",
				ServiceName: "leonardo",
				ServiceRepo: "github.com/databiosphere/leonardo",
				BuiltAt:     time.Now(),
			},
		},
		{
			name:          "missing commit sha",
			expectedError: ErrBadCreateRequest,
			expectedCode:  http.StatusBadRequest,
			createRequest: CreateBuildRequest{
				VersionString: "docker.io/asdf/lskdf:1.0.1",
				ServiceName:   "leonardo",
				ServiceRepo:   "github.com/databiosphere/leonardo",
				BuiltAt:       time.Now(),
			},
		},
		{
			name:          "missing service name",
			expectedError: ErrBadCreateRequest,
			expectedCode:  http.StatusBadRequest,
			createRequest: CreateBuildRequest{
				VersionString: "docker.io/asdf/lskdf:1.0.1",
				CommitSha:     "k234lj2",
				ServiceRepo:   "github.com/databiosphere/leonardo",
				BuiltAt:       time.Now(),
			},
		},
		{
			name:          "empty create request",
			expectedError: ErrBadCreateRequest,
			expectedCode:  http.StatusBadRequest,
			createRequest: CreateBuildRequest{},
		},
		{
			name:          "internal error",
			expectedError: errors.New("some internal error"),
			expectedCode:  http.StatusInternalServerError,
			createRequest: CreateBuildRequest{
				VersionString: "gcr.io/broad/cromwell:1.0.0",
				CommitSha:     "lk23j44",
				ServiceName:   "cromwell",
				ServiceRepo:   "github.com/broadinstitute/cromwell",
				BuildURL:      "https://jenkins.job/123",
				BuiltAt:       time.Now(),
			},
			simulateServiceCreation: false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			service := services.Service{
				ID:      1,
				Name:    testCase.createRequest.ServiceName,
				RepoURL: testCase.createRequest.ServiceRepo,
			}

			expectedBuild := &Build{
				VersionString: testCase.createRequest.VersionString,
				CommitSha:     testCase.createRequest.CommitSha,
				BuildURL:      testCase.createRequest.BuildURL,
				BuiltAt:       testCase.createRequest.BuiltAt,
				ServiceID:     service.ID,
				Service:       service,
			}
			// TODO try to simplify or DRY some of this logic

			mockBuildStore := new(mockBuildStore)
			mockBuildStore.On("createNew", mock.Anything).Return(expectedBuild, testCase.expectedError)
			mockServiceStore := new(services.MockServiceStore)

			// set up behavior for the serviceStore mock
			if testCase.simulateServiceCreation {
				mockServiceStore.On("GetByName", service.Name).Return(&services.Service{}, services.ErrServiceNotFound)
				mockServiceStore.On("CreateNew", mock.Anything).Return(&service, nil)
			} else {
				mockServiceStore.On("GetByName", service.Name).Return(&service, nil)
			}

			mockServiceController := services.ServiceController{Store: mockServiceStore}

			controller := BuildController{
				store:    mockBuildStore,
				services: &mockServiceController,
			}

			response := httptest.NewRecorder()
			gin.SetMode(gin.TestMode)
			c, _ := gin.CreateTestContext(response)

			addBuildRequestToContext(t, c, testCase.createRequest)

			controller.createBuild(c)
			assert.Equal(t, testCase.expectedCode, response.Code)

			if testCase.expectedError == ErrBadCreateRequest {
				mockServiceStore.AssertNotCalled(t, "GetByName")
				mockBuildStore.AssertNotCalled(t, "createNew")
			} else {
				mockServiceStore.AssertCalled(t, "GetByName", testCase.createRequest.ServiceName)
			}

			// ensure the create service method on the mock store is called in
			// case where a new service needs to be created
			if testCase.simulateServiceCreation {
				mockServiceStore.AssertCalled(t, "CreateNew", mock.Anything)
			}

			var expectedResponse Response
			if testCase.expectedError != nil {
				expectedResponse = Response{Error: testCase.expectedError.Error()}
			} else {
				expectedResponse = Response{Builds: []Build{*expectedBuild}}
			}

			validateResponse(t, response, expectedResponse)
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

func (m *mockBuildStore) createNew(newBuild *Build) (*Build, error) {
	retval := m.Called(newBuild)
	return retval.Get(0).(*Build), retval.Error(1)
}

func addBuildRequestToContext(t *testing.T, c *gin.Context, bodyData CreateBuildRequest) {
	t.Helper()

	reqBody := new(bytes.Buffer)
	if err := json.NewEncoder(reqBody).Encode(bodyData); err != nil {
		t.Fatalf("error parsing request body: %v", err)
	}

	req, err := http.NewRequest(http.MethodPost, "/builds", reqBody)
	if err != nil {
		t.Fatalf("error generating test request: %v", err)
	}
	c.Request = req
}

func validateResponse(t *testing.T, response *httptest.ResponseRecorder, expectedResponse Response) {
	t.Helper()

	var gotResponse Response
	if err := json.NewDecoder(response.Body).Decode(&gotResponse); err != nil {
		t.Fatalf("error decoding response body: %v", err)
	}

	if diff := cmp.Diff(gotResponse, expectedResponse); diff != "" {
		t.Errorf("unexpected difference in response body:\n%v\n", diff)
	}
}
