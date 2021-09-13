package builds

import (
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

// below is boilerplate code for the testify/mock library

type mockBuildStore struct {
	mock.Mock
}

func (m *mockBuildStore) listAll() ([]Build, error) {
	retVal := m.Called()
	return retVal.Get(0).([]Build), retVal.Error(1)
}

func (m *mockBuildStore) createNew(newBuild CreateBuildRequest) (Build, error) {
	retval := m.Called(newBuild)
	return retval.Get(0).(Build), retval.Error(1)
}
