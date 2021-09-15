package cli

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/broadinstitute/sherlock/internal/version"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func Test_VersionCommand(t *testing.T) {
	cmd := rootCmd
	cmd.AddCommand(versionCmd)
	got, err := executeCommand(cmd, "version")
	assert.NoError(t, err)
	expected := fmt.Sprintf(versionFormatString, version.BuildVersion)

	assert.Equal(t, expected, got)

}

func executeCommand(root *cobra.Command, args ...string) (string, error) {
	buf := new(bytes.Buffer)
	root.SetOut(buf)
	root.SetErr(buf)
	root.SetArgs(args)

	_, err := root.ExecuteC()
	return buf.String(), err
}
