package cli

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/providers/posflag"
	"github.com/rs/zerolog/log"

	v2 "github.com/broadinstitute/sherlock/internal/cli/v2"
	"github.com/spf13/cobra"
)

var (
	config                = koanf.New(".")
	cfgFile               string
	sherlockServerURL     string
	clientCredentials     string
	useServiceAccountAuth bool
	rootCmd               = &cobra.Command{
		Use:   "sherlock",
		Short: "sherlock tracks and manages Terra's environments",
		Long: `Sherlock is an inventory and tracking service for Terra's
persistent environments.

It also acts as a control plane for on demand ephemeral environments.
The primary purposes of this CLI tool are:

1. To report build and deployment events from ci/cd pipelines.
2. To quickly query deployment history of particular services or environments
3. To request on demand preview environments`,
	}
)

// Execute initalizes the cobra command processing tree for sherlock
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.sherlock.yaml")
	rootCmd.PersistentFlags().StringVar(&sherlockServerURL, "sherlock-url", "https://sherlock.dsp-devops.broadinstitute.org", "Address of the sherlock server")
	rootCmd.PersistentFlags().StringVar(&clientCredentials, "credentials-file", "/app/sherlock/client-sa.json", "Path to the file containing service account credentials for auth in automated workflows")
	rootCmd.PersistentFlags().BoolVar(&useServiceAccountAuth, "use-sa-auth", false, "Whether or not to use service account credentials for oauth")

	err := config.Load(posflag.Provider(rootCmd.PersistentFlags(), ".", config), nil)

	// perform initialization stuff for v2 cli
	v2.Initialize()
	rootCmd.AddCommand(v2.RootCmd)
	cobra.CheckErr(err)
}

func initConfig() {
	if cfgFile == "" {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)
		cfgFile = filepath.Clean(filepath.Join(home, ".sherlock.yaml"))
	}
	if err := config.Load(file.Provider(cfgFile), yaml.Parser()); err != nil {
		log.Info().Msgf("not using a configuration file, looked at '%s': %v", cfgFile, err)
	}

	if err := config.Load(env.Provider("SHERLOCK_", ".", func(s string) string {
		return strings.Replace(strings.ToLower(
			strings.TrimPrefix(s, "SHERLOCK_")), "_", "-", -1)
	}), nil); err != nil {
		log.Fatal().Msgf("failed to load config from environment: %v", err)
	}
}
