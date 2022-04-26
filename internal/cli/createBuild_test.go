package cli

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/broadinstitute/sherlock/internal/serializers/v1serializers"
	"net/http"
	"net/http/httptest"
	"testing"

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
				_ = json.NewEncoder(w).Encode(&v1serializers.BuildsResponse{
					Builds: []v1serializers.BuildResponse{
						{
							ID:            1,
							VersionString: "gcr.io./broad/test-service:1.0.0",
							CommitSha:     "l2kj34",
							Service: v1serializers.ServiceResponse{
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
				_ = json.NewEncoder(w).Encode(&v1serializers.BuildsResponse{
					Error: "some error from sherlock server",
				})
			},
			expectError: errors.New("some error from sherlock server"),
		},
		{
			name: "unparseable response",
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
				fmt.Fprint(w, "invalid response")
			},
			expectError: errors.New("error parsing create build response"),
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
				// parse the output back into a BuildsResponse so that we can examine it
				var cliResponse v1serializers.BuildsResponse
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
