package v2

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/broadinstitute/sherlock/internal/handlers/misc"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func Test_meCommand(t *testing.T) {
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
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testServer := httptest.NewServer(testCase.mockServerResponse)
			defer testServer.Close()

			testURL := strings.TrimPrefix(testServer.URL, "http://")
			clientOptions := sherlockClientOptions{
				hostURL: testURL,
				schemes: []string{"http"},
			}
			client, err := NewSherlockClient(clientOptions)
			if err != nil {
				t.Fatalf("error building sherlock client: %v", err)
			}

			app = client

			buildV2CommandTree()
			// disable pre run intialization so test sherlock client doesn't get overwritten
			RootCmd.PersistentPreRunE = nil

			output, err := executeCommand(RootCmd, testCase.cliArgs...)
			log.Println(output)
			assert.NoError(t, err, "expected no error from me command but got one")
		})
	}
}

func executeCommand(root *cobra.Command, args ...string) (string, error) {
	buf := new(bytes.Buffer)
	root.SetOut(buf)
	root.SetErr(buf)
	root.SetArgs(args)

	_, err := root.ExecuteC()
	return buf.String(), err
}
