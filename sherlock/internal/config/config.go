package config

import (
	"fmt"
	embeddedFiles "github.com/broadinstitute/sherlock/sherlock/config"
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
)

// Config holds Sherlock's global configuration
var Config = koanf.New(".")

func init() {
	var infoMessages []string
	if err := Config.Load(fs.Provider(embeddedFiles.EmbeddedFiles, "default_config.yaml"), yaml.Parser()); err != nil {
		panic(fmt.Sprintf("failed to load default_config.yaml, panicking due to likely embedding issue: %w", err))
	}

	if err := Config.Load(file.Provider("/etc/sherlock/sherlock.yaml"), yaml.Parser()); err == nil {
		infoMessages = append(infoMessages, "CONF | loaded config from /etc/sherlock/sherlock.yaml")
	}

	if err := Config.Load(env.Provider("SHERLOCK_", ".", func(s string) string {
		return strings.Replace(strings.ToLower(
			strings.TrimPrefix(s, "SHERLOCK_")), "_", ".", -1)
	}), nil); err != nil {
		panic(fmt.Sprintf("failed to load config from environment, panicking due to likely runtime issue: %w", err))
	}

	// We handle environment variables both case-insensitively (transformed to lowercase) and case-sensitively.
	// Case-sensitive environment variables are ugly, but Koanf is case-sensitive and otherwise we can't set
	// camelCase properties from the environment.
	if err := Config.Load(env.Provider("SHERLOCK_", ".", func(s string) string {
		return strings.Replace(strings.TrimPrefix(s, "SHERLOCK_"), "_", ".", -1)
	}), nil); err != nil {
		panic(fmt.Sprintf("failed to load config from environment, panicking due to likely runtime issue: %w", err))
	}

	if Config.MustString("mode") != "debug" && Config.MustString("mode") != "release" {
		panic(fmt.Sprintf("mode was %s instead of either 'debug' or 'release'", Config.MustString("mode")))
	}

	configureLogging(infoMessages...)
}

// LoadTestConfig is an extra at-test-time-only configuration loading step. This package's init function will have
// already run and loaded Sherlock's default configuration, but we might have different configuration we'd want to
// apply while testing. This function loads this package's test_config.yaml and any TEST_SHERLOCK_* environment
// variables on top of whatever configuration currently exists.
func LoadTestConfig() {
	if err := Config.Load(fs.Provider(embeddedFiles.EmbeddedFiles, "test_config.yaml"), yaml.Parser()); err != nil {
		panic(fmt.Errorf("failed to load test configuration file test_config.yaml: %w", err))
	}

	if err := Config.Load(env.Provider("TEST_SHERLOCK", ".", func(s string) string {
		return strings.Replace(strings.ToLower(
			strings.TrimPrefix(s, "TEST_SHERLOCK_")), "_", ".", -1)
	}), nil); err != nil {
		panic(fmt.Errorf("failed to load test configuration environment variables: %w", err))
	}

	// We handle environment variables both case-insensitively (transformed to lowercase) and -sensitively.
	if err := Config.Load(env.Provider("TEST_SHERLOCK", ".", func(s string) string {
		return strings.Replace(strings.TrimPrefix(s, "TEST_SHERLOCK_"), "_", ".", -1)
	}), nil); err != nil {
		panic(fmt.Errorf("failed to load test configuration environment variables: %w", err))
	}

	configureLogging()
}

// configureLogging is abstracted init-like logic for updating global logging configuration.
// It accepts arbitrary messages to log at the info level once the logger is set up, in case
// there's messages about the loaded configuration.
func configureLogging(infoMessages ...string) {
	var logBuilder zerolog.Context
	if Config.String("mode") == "debug" {
		// Colored text for CLI
		logBuilder = log.Output(zerolog.ConsoleWriter{Out: os.Stderr}).With()
	} else {
		// JSON for GCP
		logBuilder = zerolog.New(os.Stderr).With()
		zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs
	}
	if Config.Bool("log.timestamp") {
		logBuilder = logBuilder.Timestamp()
	}
	if Config.Bool("log.caller") {
		logBuilder = logBuilder.Caller()
	}
	log.Logger = logBuilder.Logger()

	// Some other packages (looking at you, fvbock/endless) take it upon themselves to send
	// log messages using Go's built-in logging. We can at least format those messages
	// correctly by redirecting that into zerolog, though it won't have proper leveling
	// information
	stdlog.SetOutput(log.Logger)

	if logLevel := Config.String("log.level"); logLevel != "" {
		if parsedLevel, err := zerolog.ParseLevel(logLevel); err != nil {
			log.Warn().Msgf("log level '%s' couldn't be parsed by zerolog", logLevel)
		} else {
			zerolog.SetGlobalLevel(parsedLevel)
		}
	}

	for _, m := range infoMessages {
		log.Info().Msg(m)
	}
}
