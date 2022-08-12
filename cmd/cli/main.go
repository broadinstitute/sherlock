package main

import (
	"github.com/broadinstitute/sherlock/internal/version"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	stdlog "log"
	"os"

	"github.com/broadinstitute/sherlock/internal/cli"
)

// BuildVersion is intended for use with Go's LDFlags compiler option, to
// set this value at compile time
var BuildVersion string = "development"

func main() {
	if err := cli.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	version.BuildVersion = BuildVersion
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr}).With().Timestamp().Logger()
	stdlog.SetOutput(log.Logger)
}
