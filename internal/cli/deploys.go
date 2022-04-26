package cli

import (
	"github.com/spf13/cobra"
)

var (
	deployCmd = &cobra.Command{
		Use:   "v1mocks",
		Short: "v1mocks is a group of commands for interacting with sherlock deploy events",
		Long: `v1mocks contains a group of commands for viewing existing builds and creating new v1mocks.
Currently supported commands:
	1. create - creates a new deploy event`,
	}
)

func init() {
	rootCmd.AddCommand(deployCmd)
}
