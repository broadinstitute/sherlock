package builds

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/broadinstitute/sherlock/internal/models/v1models"
	"net/http"
	"net/http/httptest"
	"strconv"
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
		expectedBuilds []v1models.Build
		expectedError  error
		expectedCode   int
	}{
		{
			name:           "no existing builds",
			expectedBuilds: []v1models.Build{},
			expectedCode:   http.StatusOK,
			expectedError:  nil,
		},
		{
			name: "one existing build",
			expectedBuilds: []v1models.Build{
				{
					VersionString: "imagerepo.io/test:0.1.0",
					CommitSha:     "f2f2f23",
					BuildURL:      "https://build.job.out/blah",
					BuiltAt:       time.Now(),
					ServiceID:     1,
					Service: v1models.Service{
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
			expectedBuilds: []v1models.Build{
				{
					VersionString: "imagerepo.io/test:0.1.0",
					CommitSha:     "f2f2f23",
					BuildURL:      "https://build.job.out/blah",
					BuiltAt:       time.Now(),
					ServiceID:     1,
					Service: v1models.Service{
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
					Service: v1models.Service{
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
					Service: v1models.Service{
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
					Service: v1models.Service{
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
			expectedBuilds: []v1models.Build{},
			expectedError:  errors.New("some internal error"),
			expectedCode:   http.StatusInternalServerError,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// setup mock store
			mockStore := new(mockBuildStore)
			mockStore.On("ListAll").Return(testCase.expectedBuilds, testCase.expectedError)

			controller := BuildController{store: mockStore}

			response := httptest.NewRecorder()
			gin.SetMode(gin.TestMode)
			c, _ := gin.CreateTestContext(response)

			controller.getBuilds(c)

			mockStore.AssertCalled(t, "ListAll")
			assert.Equal(t, testCase.expectedCode, response.Code)

			var gotResponse Response
			if err := json.NewDecoder(response.Body).Decode(&gotResponse); err != nil {
				t.Fatalf("error decoding response body: %v\n", err)
			}

			var expectedResponse Response
			if testCase.expectedError != nil {
				expectedResponse = Response{Error: testCase.expectedError.Error()}
			} else {
				expectationSerializer := BuildsSerializer{Builds: testCase.expectedBuilds}
				expectedResponse = Response{Builds: expectationSerializer.Response()}
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
			expectedError: v1models.ErrBadCreateRequest,
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
			expectedError: v1models.ErrBadCreateRequest,
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
			expectedError: v1models.ErrBadCreateRequest,
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
			expectedError: v1models.ErrBadCreateRequest,
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
		{
			name:          "non-unique version string",
			expectedError: v1models.ErrDuplicateVersionString,
			expectedCode:  http.StatusBadRequest,
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

			service := v1models.Service{
				ID:      1,
				Name:    testCase.createRequest.ServiceName,
				RepoURL: testCase.createRequest.ServiceRepo,
			}

			expectedBuild := v1models.Build{
				VersionString: testCase.createRequest.VersionString,
				CommitSha:     testCase.createRequest.CommitSha,
				BuildURL:      testCase.createRequest.BuildURL,
				BuiltAt:       testCase.createRequest.BuiltAt,
				ServiceID:     service.ID,
				Service:       service,
			}
			// TODO try to simplify or DRY some of this logic

			mockBuildStore := new(mockBuildStore)
			mockBuildStore.On("CreateNew", mock.Anything).Return(expectedBuild, testCase.expectedError)
			mockServiceStore := new(services.MockServiceStore)

			// set up behavior for the serviceStore mock
			if testCase.simulateServiceCreation {
				mockServiceStore.On("GetByName", service.Name).Return(v1models.Service{}, v1models.ErrServiceNotFound)
				mockServiceStore.On("CreateNew", mock.Anything).Return(service, nil)
			} else {
				mockServiceStore.On("GetByName", service.Name).Return(service, nil)
			}

			mockServiceController := services.NewMockController(mockServiceStore)

			controller := BuildController{
				store:    mockBuildStore,
				services: mockServiceController,
			}

			response := httptest.NewRecorder()
			gin.SetMode(gin.TestMode)
			c, _ := gin.CreateTestContext(response)

			addBuildRequestToContext(t, c, testCase.createRequest)

			controller.createBuild(c)
			assert.Equal(t, testCase.expectedCode, response.Code)

			if testCase.expectedError == v1models.ErrBadCreateRequest {
				mockServiceStore.AssertNotCalled(t, "GetByName")
				mockBuildStore.AssertNotCalled(t, "CreateNew")
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
				expectationSerializer := BuildSerializer{expectedBuild}
				expectedResponse = Response{Builds: []BuildResponse{expectationSerializer.Response()}}
			}

			validateResponse(t, response, expectedResponse)
		})
	}
}

func TestGetBuildByID(t *testing.T) {
	testCases := []struct {
		name          string
		buildID       string
		expectedBuild v1models.Build
		expectedError error
		expectedCode  int
	}{
		{
			name:    "successfully get build by id",
			buildID: "1",
			expectedBuild: v1models.Build{
				ID:            1,
				VersionString: "imagerepo.io/test:0.1.0",
				CommitSha:     "f2f2f23",
				BuildURL:      "https://build.job.out/blah",
				BuiltAt:       time.Now(),
				ServiceID:     1,
				Service: v1models.Service{
					ID:      1,
					Name:    "tester",
					RepoURL: "https://tester.repo",
				},
			},
			expectedCode:  http.StatusOK,
			expectedError: nil,
		},
		{
			name:          "non-existent build id",
			buildID:       "100",
			expectedBuild: v1models.Build{},
			expectedCode:  http.StatusNotFound,
			expectedError: v1models.ErrBuildNotFound,
		},
		{
			name:          "invalid id param",
			buildID:       "abc",
			expectedBuild: v1models.Build{},
			expectedCode:  http.StatusBadRequest,
			expectedError: ErrInvalidBuildID,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			mockStore := new(mockBuildStore)
			buildID, err := strconv.Atoi(testCase.buildID)
			mockStore.On("GetByID", buildID).Return(testCase.expectedBuild, testCase.expectedError)

			controller := BuildController{store: mockStore}

			response := httptest.NewRecorder()
			gin.SetMode(gin.TestMode)
			c, _ := gin.CreateTestContext(response)

			// gin stores url params including numeric ids as strings
			c.Params = append(c.Params, gin.Param{
				Key:   "id",
				Value: testCase.buildID,
			})

			controller.getByID(c)

			// err will have a value when id param cannot be successfully parsed as an int and thus is invalid
			// so getByID method should not be called
			if err != nil {
				mockStore.AssertNotCalled(t, "GetByID")
			} else {
				mockStore.AssertCalled(t, "GetByID", buildID)
			}

			assert.Equal(t, testCase.expectedCode, response.Code)

			var expectedResponse Response
			if testCase.expectedError != nil {
				expectedResponse = Response{Error: testCase.expectedError.Error()}
			} else {
				expectationSerializer := BuildSerializer{testCase.expectedBuild}
				expectedResponse = Response{Builds: []BuildResponse{expectationSerializer.Response()}}
			}

			validateResponse(t, response, expectedResponse)
		})
	}
}

// below is boilerplate code for the testify/mock library
type mockBuildStore struct {
	mock.Mock
}

func (m *mockBuildStore) ListAll() ([]v1models.Build, error) {
	retVal := m.Called()
	return retVal.Get(0).([]v1models.Build), retVal.Error(1)
}

func (m *mockBuildStore) CreateNew(newBuild v1models.Build) (v1models.Build, error) {
	retval := m.Called(newBuild)
	return retval.Get(0).(v1models.Build), retval.Error(1)
}

func (m *mockBuildStore) GetByID(id int) (v1models.Build, error) {
	retVal := m.Called(id)
	return retVal.Get(0).(v1models.Build), retVal.Error(1)
}

func (m *mockBuildStore) GetByVersionString(versionString string) (v1models.Build, error) {
	retVal := m.Called(versionString)
	return retVal.Get(0).(v1models.Build), retVal.Error(1)
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
