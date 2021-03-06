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
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_createDeployCommand(t *testing.T) {
	testCases := []struct {
		name               string
		cliArgs            []string
		mockServerResponse http.HandlerFunc
		expectError        error
	}{
		{
			name: "successful create",
			cliArgs: []string{
				"deploys",
				"create",
				"docker.io/my-repo/my-app:1.0.0",
				"--environment",
				"dev",
				"--service",
				"my-app",
			},
			mockServerResponse: func(w http.ResponseWriter, r *http.Request) {
				_ = json.NewEncoder(w).Encode(&v1serializers.DeploysResponse{
					Deploys: []v1serializers.DeployResponse{
						{
							ID: 1,
							ServiceInstance: v1serializers.ServiceInstanceResponse{
								ID: 1,
								Service: v1serializers.ServiceResponse{
									ID:   1,
									Name: "my-app",
								},
								Environment: v1serializers.EnvironmentResponse{
									ID:   1,
									Name: "dev",
								},
							},
							Build: v1serializers.BuildResponse{
								ID:            1,
								VersionString: "docker.io/my-repo/my-app:1.0.0",
								BuiltAt:       time.Now(),
								Service: v1serializers.ServiceResponse{
									ID:   1,
									Name: "my-app",
								},
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
				"deploys",
				"create",
				"gcr.io./broad/test-service:1.0.0",
				"--environment",
				"qa",
				"--service",
				"test-service",
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
				"deploys",
				"create",
				"gcr.io./broad/test-service:1.0.0",
				"--environment",
				"terra-prod",
				"--service",
				"cromwell",
			},
			mockServerResponse: func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprint(w, "invalid response")
			},
			expectError: errors.New("error parsing create deploy response"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// set up a mock server of the sherlock api
			testServer := httptest.NewServer(testCase.mockServerResponse)

			sherlockServerURL = testServer.URL
			output, _ := executeCommand(rootCmd, "deploys", "create", testCase.cliArgs[2], "--environment", testCase.cliArgs[4], "--service", testCase.cliArgs[6])
			outputBytes := bytes.NewBufferString(output)

			if testCase.expectError == nil {
				// parse the output back into a v1serializers.DeploysResponse so that we can examine it
				var cliResponse v1serializers.DeploysResponse
				if err := json.NewDecoder(outputBytes).Decode(&cliResponse); err != nil {
					t.Errorf("error decoding cli output: %v", err)
				}

				assert.Equal(t, cliResponse.Deploys[0].ServiceInstance.Environment.Name, testCase.cliArgs[4])
				assert.Equal(t, cliResponse.Deploys[0].ServiceInstance.Service.Name, testCase.cliArgs[6])
				assert.Equal(t, cliResponse.Deploys[0].Build.VersionString, testCase.cliArgs[2])
			} else {
				assert.Contains(t, output, testCase.expectError.Error())
			}
		})
	}

}
