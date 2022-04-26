package cli

import (
	"github.com/spf13/cobra"
)

var (
	deployCmd = &cobra.Command{
		Use:   "deploys",
		Short: "deploys is a group of commands for interacting with sherlock deploy events",
		Long: `deploys contains a group of commands for viewing existing builds and creating new deploys.
Currently supported commands:
	1. create - creates a new deploy event`,
	}
)

func init() {
	rootCmd.AddCommand(deployCmd)
}
