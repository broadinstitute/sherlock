package cli

import (
	"fmt"

	"github.com/broadinstitute/sherlock/internal/version"
	"github.com/spf13/cobra"
)

const (
	versionFormatString     string = "sherlock version: %s\n"
	shortVersionDescription string = "Get Sherlock's recorded build version"
	longVersionDescription  string = `Get Sherlock's internal BuildVersion, usually set via LDFlags during
	compilation.`
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: shortVersionDescription,
	Long:  longVersionDescription,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Fprintf(cmd.OutOrStdout(), versionFormatString, version.BuildVersion)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
