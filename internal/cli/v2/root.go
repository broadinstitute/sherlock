// v2 contains cobra cli command implementations for interact with sherlock v2 apis
package v2

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	app     *SherlockClient
	RootCmd = &cobra.Command{
		Use:   "v2",
		Short: "command subtree with support for v2 apis",
		Long:  "v2 contains subcommands for interacting with sherlock v2 apis",
		// ensures
		PersistentPreRunE: initialize,
	}
)

func initialize(cmd *cobra.Command, args []string) error {
	credsFile, err := cmd.Flags().GetString("credentials-file")
	if err != nil {
		return fmt.Errorf("credentials-file flag not specified: %v", err)
	}

	client, err := NewSherlockClient(credsFile)
	if err != nil {
		return fmt.Errorf("error constructing v2 client: %v", err)
	}

	app = client
	return nil
}

func buildV2CommandTree() {
	RootCmd.AddCommand(MeCmd)
}

// initialize the sub command parse tree for v2 apis
func init() {
	buildV2CommandTree()
}
