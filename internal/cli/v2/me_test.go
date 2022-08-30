package v2

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/broadinstitute/sherlock/internal/handlers/misc"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestMeCommandSuite(t *testing.T) {
	suite.Run(t, new(meCommandSuite))
}

type meCommandSuite struct {
	suite.Suite
}

func (suite *meCommandSuite) SetupSuite() {
	// initialize command parse tree
	buildV2CommandTree()
	// disable pre run intialization so test sherlock client doesn't get overwritten
	RootCmd.PersistentPreRunE = nil
}

func (suite *meCommandSuite) TestMeCommand() {
	testCases := []struct {
		name               string
		cliArgs            []string
		mockServerResponse http.HandlerFunc
		expectError        error
	}{
		{
			name: "successful response",
			cliArgs: []string{
				"me",
			},
			mockServerResponse: func(w http.ResponseWriter, r *http.Request) {
				w.Header().Add("Content-Type", "application/json")
				_ = json.NewEncoder(w).Encode(misc.MyUserResponse{
					Email:       "test@test.com",
					Suitability: "true",
				})
			},
			expectError: nil,
		},
		{
			name: "permissions error",
			cliArgs: []string{
				"me",
			},
			mockServerResponse: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusProxyAuthRequired)
				w.Header().Add("Content-Type", "application/json")
			},
			expectError: fmt.Errorf("error retrieving current user info"),
		},
	}

	for _, testCase := range testCases {
		suite.Run(testCase.name, func() {
			testServer := httptest.NewServer(testCase.mockServerResponse)
			defer testServer.Close()

			client, err := newTestClient(testServer.URL)
			if err != nil {
				suite.T().Fatalf("error building sherlock client: %v", err)
			}
			app = client

			output, err := executeCommand(RootCmd, testCase.cliArgs...)
			if testCase.expectError == nil {
				assert.NoError(suite.T(), err, "expected no error from me command but got one")
			} else {
				assert.Error(suite.T(), err)
				assert.ErrorContains(suite.T(), err, testCase.expectError.Error())
				// nothing else to assert on in error case
				return
			}

			assert.Contains(suite.T(), output, "test@test.com")
		})
	}
}

func newTestClient(url string) (*sherlockClient, error) {
	// urls from httptest include the scheme, strip this to work with the client lib
	testURL := strings.TrimPrefix(url, "http://")
	clientOptions := sherlockClientOptions{
		hostURL: testURL,
		schemes: []string{"http"},
	}
	return NewSherlockClient(clientOptions)
}

func executeCommand(root *cobra.Command, args ...string) (string, error) {
	buf := new(bytes.Buffer)
	root.SetOut(buf)
	root.SetErr(buf)
	root.SetArgs(args)

	_, err := root.ExecuteC()
	return buf.String(), err
}
