package config

import (
	"embed"
	"fmt"
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/providers/fs"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	stdlog "log"
	"os"
	"strings"
	"testing"
)

var (
	// Config holds Sherlock's global configuration
	Config = koanf.New(".")

	//go:embed default_config.yaml test_config.yaml
	embeddedFiles embed.FS
)

func init() {
	var infoMessages []string
	if err := Config.Load(fs.Provider(embeddedFiles, "default_config.yaml"), yaml.Parser()); err != nil {
		panic(fmt.Sprintf("failed to load default_config.yaml, panicking due to likely embedding issue: %v", err))
	}

	if err := Config.Load(file.Provider("/etc/sherlock/sherlock.yaml"), yaml.Parser()); err != nil {
		infoMessages = append(infoMessages, fmt.Sprintf("didn't load config from /etc/sherlock/sherlock.yaml: %v", err.Error()))
	} else {
		infoMessages = append(infoMessages, "loaded config from /etc/sherlock/sherlock.yaml")
	}

	if err := Config.Load(env.Provider("SHERLOCK_", ".", func(s string) string {
		return strings.Replace(strings.ToLower(
			strings.TrimPrefix(s, "SHERLOCK_")), "_", ".", -1)
	}), nil); err != nil {
		panic(fmt.Sprintf("failed to load config from environment, panicking due to likely runtime issue: %v", err))
	}

	configureLogging(infoMessages...)
}

// LoadTestConfig is an extra at-test-time-only configuration loading step. This package's init function will have
// already run and loaded Sherlock's default configuration, but we might have different configuration we'd want to
// apply while testing. This function loads this package's test_config.yaml and any TEST_SHERLOCK_* environment
// variables on top of whatever configuration currently exists.
func LoadTestConfig(t *testing.T) {
	if err := Config.Load(fs.Provider(embeddedFiles, "test_config.yaml"), yaml.Parser()); err != nil {
		t.Fatalf("failed to load test configuration file test_config.yaml: %v", err)
		return
	}

	if err := Config.Load(env.Provider("TEST_SHERLOCK", ".", func(s string) string {
		return strings.Replace(strings.ToLower(
			strings.TrimPrefix(s, "TEST_SHERLOCK_")), "_", ".", -1)
	}), nil); err != nil {
		t.Fatalf("failed to load test configuration environment variables: %v", err)
		return
	}

	configureLogging()
}

func configureLogging(infoMessages ...string) {
	if Config.String("mode") == "debug" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr}).With().Timestamp().Logger()
	} else {
		log.Logger = zerolog.New(os.Stderr).With().Timestamp().Logger()
	}
	stdlog.SetOutput(log.Logger)
	for _, m := range infoMessages {
		log.Info().Msg(m)
	}
	if logLevel := Config.String("logLevel"); logLevel != "" {
		if parsedLevel, err := zerolog.ParseLevel(logLevel); err != nil {
			log.Warn().Msgf("log level '%s' couldn't be parsed by zerolog", logLevel)
		} else {
			zerolog.SetGlobalLevel(parsedLevel)
		}
	}
}
