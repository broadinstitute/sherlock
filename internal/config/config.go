package config

import (
	"embed"
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/providers/fs"
	"github.com/rs/zerolog/log"
	"os"
	"path/filepath"
	"strings"
)

var (
	// Config holds Sherlock's global configuration
	Config = koanf.New(".")

	//go:embed default_config.yaml
	defaultConfig embed.FS
)

func InitConfig() {
	if err := Config.Load(fs.Provider(defaultConfig, "default_config.yaml"), yaml.Parser()); err != nil {
		log.Fatal().Msgf("failed to load config defaults: %v", err)
	}

	if err := Config.Load(file.Provider(filepath.Clean("/etc/sherlock/sherlock.yaml")), yaml.Parser()); err != nil {
		log.Info().Msgf("didn't load config from /etc/sherlock/sherlock.yaml: %v", err)
	}

	if home, err := os.UserHomeDir(); err != nil {
		log.Warn().Msgf("failed to get user home directory to look for config there: %v", err)
	} else {
		if err = Config.Load(file.Provider(filepath.Clean(filepath.Join(home, "sherlock.yaml"))), yaml.Parser()); err != nil {
			log.Info().Msgf("didn't load config from ~/sherlock.yaml: %v", err)
		}
	}

	if err := Config.Load(env.Provider("SHERLOCK_", ".", func(s string) string {
		return strings.Replace(strings.ToLower(
			strings.TrimPrefix(s, "SHERLOCK_")), "_", ".", -1)
	}), nil); err != nil {
		log.Fatal().Msgf("failed to load config from environment: %v", err)
	}
}
