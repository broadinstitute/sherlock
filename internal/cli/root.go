package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile           string
	sherlockServerURL string
	rootCmd           = &cobra.Command{
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
	rootCmd.PersistentFlags().Bool("viper", true, "use Viper for configuration")
	rootCmd.PersistentFlags().StringVar(&sherlockServerURL, "addr", "http://localhost:8080", "Address of the sherlock server")

	err := viper.BindPFlags(rootCmd.PersistentFlags())
	cobra.CheckErr(err)
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		// find home dir
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".sherlock")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Printf("not using a configuration file: %v\n", err)
		} else {
			cobra.CheckErr(err)
		}
	}
}
