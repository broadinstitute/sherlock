package cli

import (
	"github.com/spf13/cobra"
)

var (
	buildCmd = &cobra.Command{
		Use:   "builds",
		Short: "builds is a group of commands for interacting with sherlock build entities",
		Long: `builds contains a group of commands for viewing existing builds and creating new builds.
Currently supported commands:
	1. create - creates a new build entity`,
	}
)

func init() {
	rootCmd.AddCommand(buildCmd)
}
