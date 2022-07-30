package main

import (
	"github.com/broadinstitute/sherlock/internal/config"
	"github.com/broadinstitute/sherlock/internal/db"
	"github.com/broadinstitute/sherlock/internal/version"
	"github.com/fvbock/endless"
	"github.com/golang-migrate/migrate/v4"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	stdlog "log"
	"os"

	"github.com/broadinstitute/sherlock/internal/sherlock"
)

// BuildVersion is intended for use with Go's LDFlags compiler option, to
// set this value at compile time
var BuildVersion string = "development"

func main() {
	config.InitConfig()

	if err := db.ApplyMigrations("db/migrations", config.Config); err != nil {
		// don't fail if there are no changes to apply
		if err == migrate.ErrNoChange {
			log.Info().Msg("no migration to apply, continuing...")
		} else {
			log.Fatal().Msgf("%v", err)
			return
		}
	}

	app := sherlock.New()
	if app == nil {
		log.Fatal().Msg("failed to create an application instance")
		return
	}

	defer app.ShutdownStackdriver()
	defer app.CancelContexts()

	log.Info().Msg("starting sherlock server on :8080")
	if err := endless.ListenAndServe(":8080", app); err != nil {
		log.Warn().Msgf("%v", err)
	}
}

func init() {
	version.BuildVersion = BuildVersion
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr}).With().Timestamp().Logger()
	stdlog.SetOutput(log.Logger)
}
