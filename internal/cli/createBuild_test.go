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
	// testCases := []struct {
	// 	name        string
	// 	cliArgs     []string
	// 	expectError bool
	// }{}

	// set up command tree
	cmd := rootCmd
	cmd.AddCommand(buildCmd)
	buildCmd.AddCommand(createBuildCmd)

	// set up a mock server of the sherlock api
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(&builds.Response{})
	}))

	sherlockServerURL = testServer.URL
	output, err := executeCommand(cmd, "builds", "create", "testService", "--version-string", "blah", "--commit-sha", "alsoblah")
	outputBytes := bytes.NewBufferString(output)
	assert.NoError(t, err)

	// parse the output back into a builds.Response so that we can examine it
	var cliResponse builds.Response
	if err := json.NewDecoder(outputBytes).Decode(&cliResponse); err != nil {
		t.Errorf("error decoding cli output: %v", err)
	}
}
