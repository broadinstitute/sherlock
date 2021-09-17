package cli

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/broadinstitute/sherlock/internal/builds"
	"github.com/broadinstitute/sherlock/internal/services"
	"github.com/stretchr/testify/assert"
)

func Test_createBuildCommand(t *testing.T) {
	testCases := []struct {
		name               string
		cliArgs            []string
		mockServerResponse http.HandlerFunc
		expectError        error
	}{
		{
			name: "successful create",
			cliArgs: []string{
				"builds",
				"create",
				"test-service",
				"--version-string",
				"gcr.io./broad/test-service:1.0.0",
				"--commit-sha",
				"l2kj34",
			},
			mockServerResponse: func(w http.ResponseWriter, r *http.Request) {
				_ = json.NewEncoder(w).Encode(&builds.Response{
					Builds: []builds.BuildResponse{
						{
							ID:            1,
							VersionString: "gcr.io./broad/test-service:1.0.0",
							CommitSha:     "l2kj34",
							Service: services.ServiceResponse{
								ID:   1,
								Name: "test-service",
							},
						},
					},
				})
			},
			expectError: nil,
		},
		{
			name: "error from server",
			cliArgs: []string{
				"builds",
				"create",
				"test-service",
				"--version-string",
				"gcr.io./broad/test-service:1.0.0",
				"--commit-sha",
				"l2kj34",
			},
			mockServerResponse: func(w http.ResponseWriter, r *http.Request) {
				_ = json.NewEncoder(w).Encode(&builds.Response{
					Error: "some error from sherlock server",
				})
			},
			expectError: errors.New("some error from sherlock server"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// set up a mock server of the sherlock api
			testServer := httptest.NewServer(testCase.mockServerResponse)

			sherlockServerURL = testServer.URL
			output, _ := executeCommand(rootCmd, "builds", "create", testCase.cliArgs[2], "--version-string", testCase.cliArgs[4], "--commit-sha", testCase.cliArgs[6])
			outputBytes := bytes.NewBufferString(output)

			if testCase.expectError == nil {
				// parse the output back into a builds.Response so that we can examine it
				var cliResponse builds.Response
				if err := json.NewDecoder(outputBytes).Decode(&cliResponse); err != nil {
					t.Errorf("error decoding cli output: %v", err)
				}

				assert.Equal(t, cliResponse.Builds[0].VersionString, testCase.cliArgs[4])
				assert.Equal(t, cliResponse.Builds[0].CommitSha, testCase.cliArgs[6])
				assert.Equal(t, cliResponse.Builds[0].Service.Name, testCase.cliArgs[2])
			} else {
				assert.Contains(t, output, testCase.expectError.Error())
			}
		})
	}

}
