package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	createBuildCmd = &cobra.Command{
		Use:   "create",
		Short: "create a new build",
		Long:  `creates a new build entity in sherlock.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprintf(cmd.OutOrStdout(), "this is where we make a build")
		},
	}
)

func init() {
	buildCmd.AddCommand(createBuildCmd)
}
