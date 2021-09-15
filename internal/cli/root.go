package cli

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "sherlock",
	Short: "sherlock tracks and manages Terra's environments",
	Long: `Sherlock is an inventory and tracking service for Terra's
persistent environments.

It also acts as a control plane for on demand ephemeral environments.
The primary purposes of this CLI tool are:

1. To report build and deployment events from ci/cd pipelines.
2. To quickly query deployment history of particular services or environments
3. To request on demand previewenvironments`,
}

// Execute initalizes the cobra command processing tree for sherlock
func Execute() error {
	return rootCmd.Execute()
}
