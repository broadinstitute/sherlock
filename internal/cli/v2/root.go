// v2 contains cobra cli command implementations for interact with sherlock v2 apis
package v2

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var (
	app     *sherlockClient
	RootCmd = &cobra.Command{
		Use:   "v2",
		Short: "command subtree with support for v2 apis",
		Long:  "v2 contains subcommands for interacting with sherlock v2 apis",
		// ensures
		PersistentPreRunE: initialize,
	}
)

func initialize(cmd *cobra.Command, args []string) error {
	// The error case cannot happen in these lookups, defaults are set by the top level root command
	credsFile, _ := cmd.Flags().GetString("credentials-file")
	hostURL, _ := cmd.Flags().GetString("sherlock-url")
	// remove https:// protocol prefix, needed for v1 cli support but not the v2 client lib
	hostURL = strings.TrimPrefix(hostURL, "https://")
	useSaAuth, _ := cmd.Flags().GetBool("use-sa-auth")

	clientOptions := sherlockClientOptions{
		hostURL:               hostURL,
		credentialsPath:       credsFile,
		schemes:               []string{"https"},
		useServiceAccountAuth: useSaAuth,
	}

	client, err := NewSherlockClient(clientOptions)
	if err != nil {
		return fmt.Errorf("error constructing v2 client: %v", err)
	}

	app = client
	return nil
}

func buildV2CommandTree() {
	RootCmd.AddCommand(meCmd)
}

// initialize the sub command parse tree for v2 apis
func init() {
	buildV2CommandTree()
}
