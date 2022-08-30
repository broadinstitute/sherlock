// v2 contains cobra cli command implementations for interact with sherlock v2 apis
package v2

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/posflag"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var (
	app     *sherlockClient
	cfgFile string
	config  = koanf.New(".")
	RootCmd = &cobra.Command{
		Use:   "v2",
		Short: "command subtree with support for v2 apis",
		Long:  "v2 contains subcommands for interacting with sherlock v2 apis",
		// ensures configuration and client initialization happens as a pre-run step before each child command
		PersistentPreRunE: initialize,
	}
)

func initialize(cmd *cobra.Command, args []string) error {
	// initialize global config flags passed
	if err := config.Load(posflag.Provider(cmd.Flags(), ".", config), nil); err != nil {
		return err
	}

	// Pull global config out of koanf config object
	credsFile := config.String("credentials-file")
	hostURL := config.String("sherlock-url")
	// remove https:// protocol prefix, needed for v1 cli support but not the v2 client lib
	hostURL = strings.TrimPrefix(hostURL, "https://")
	audience := config.String("oauth-audience")

	clientOptions := sherlockClientOptions{
		hostURL:         hostURL,
		credentialsPath: credsFile,
		schemes:         []string{"https"},
		audience:        audience,
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
func Initialize() {
	initConfig()
	buildV2CommandTree()
}

func initConfig() {
	if cfgFile == "" {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)
		cfgFile = filepath.Clean(filepath.Join(home, ".sherlock.yaml"))
	}

	if err := config.Load(env.Provider("SHERLOCK_", ".", func(s string) string {
		return strings.Replace(strings.ToLower(
			strings.TrimPrefix(s, "SHERLOCK_")), "_", "-", -1)
	}), nil); err != nil {
		log.Fatal().Msgf("failed to load config from environment: %v", err)
	}
}
