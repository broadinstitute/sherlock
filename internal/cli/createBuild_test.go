package cli

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/broadinstitute/sherlock/internal/builds"
	"github.com/stretchr/testify/assert"
)

func Test_createBuildCommand(t *testing.T) {
	testCases := []struct {
		name        string
		cliArgs     []string
		expectError bool
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
			expectError: false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// set up a mock server of the sherlock api
			testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				json.NewEncoder(w).Encode(&builds.Response{
					Builds: []builds.BuildResponse{
						{},
					},
				})
			}))

			sherlockServerURL = testServer.URL
			output, err := executeCommand(rootCmd, "builds", "create", testCase.cliArgs[2], "--version-string", testCase.cliArgs[4], "--commit-sha", testCase.cliArgs[6])
			outputBytes := bytes.NewBufferString(output)

			if testCase.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			// parse the output back into a builds.Response so that we can examine it
			var cliResponse builds.Response
			if err := json.NewDecoder(outputBytes).Decode(&cliResponse); err != nil {
				t.Errorf("error decoding cli output: %v", err)
			}
		})
	}

}
