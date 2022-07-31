package config

import (
	"embed"
	"fmt"
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/providers/fs"
	"strings"
)

var (
	// Config holds Sherlock's global configuration
	Config = koanf.New(".")

	//go:embed default_config.yaml
	defaultConfig embed.FS
)

func init() {
	// Using an actual log package here results in unpredictable log formatting,
	// because that gets set in another init function and the execution order of
	// init functions is undefined

	if err := Config.Load(fs.Provider(defaultConfig, "default_config.yaml"), yaml.Parser()); err != nil {
		panic(fmt.Sprintf("failed to load config defaults: %v", err))
	}

	if err := Config.Load(file.Provider("/etc/sherlock/sherlock.yaml"), yaml.Parser()); err != nil {
		println("didn't load config from /etc/sherlock/sherlock.yaml,", err.Error())
	} else {
		println("loaded config from /etc/sherlock/sherlock.yaml")
	}

	if err := Config.Load(env.Provider("SHERLOCK_", ".", func(s string) string {
		return strings.Replace(strings.ToLower(
			strings.TrimPrefix(s, "SHERLOCK_")), "_", ".", -1)
	}), nil); err != nil {
		panic(fmt.Sprintf("failed to load config from environment: %v", err))
	}
}
