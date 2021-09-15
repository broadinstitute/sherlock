package cli

import (
	"fmt"

	"github.com/broadinstitute/sherlock/internal/version"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Get Sherlock's recorded build version",
	Long: `Get Sherlock's internal BuildVersion, usually set via LDFlags during
compilation.'`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("sherlock version: %s\n", version.BuildVersion)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
