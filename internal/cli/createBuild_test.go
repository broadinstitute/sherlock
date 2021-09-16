package cli

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/broadinstitute/sherlock/internal/builds"
	"github.com/stretchr/testify/assert"
)

func Test_createBuildCommand(t *testing.T) {
	// set up command tree
	cmd := rootCmd
	cmd.AddCommand(buildCmd)
	buildCmd.AddCommand(createBuildCmd)

	// set up a mock server of the sherlock api
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(&builds.Response{})
	}))

	sherlockServerURL = testServer.URL
	_, err := executeCommand(cmd, "builds", "create", "testService", "--version-string", "blah", "--commit-sha", "alsoblah")
	assert.NoError(t, err)
}
