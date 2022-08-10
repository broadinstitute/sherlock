package main

import (
	"github.com/broadinstitute/sherlock/internal/db"
	"github.com/broadinstitute/sherlock/internal/sherlock"
	"github.com/broadinstitute/sherlock/internal/version"
	"github.com/fvbock/endless"
	"github.com/rs/zerolog/log"
)

// BuildVersion is intended for use with Go's LDFlags compiler option, to
// set this value at compile time
var BuildVersion string = "development"

func main() {
	sqlDB, err := db.Connect()
	if err != nil {
		log.Fatal().Msgf("%v", err)
	}

	// Spin up the liveness check once we've got a good database connection
	go func() { _ = endless.ListenAndServe(":8081", livenessHandler{}) }()

	gormDB, err := db.Configure(sqlDB)
	if err != nil {
		log.Fatal().Msgf("%v", err)
	}

	app := sherlock.New(gormDB)
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
}
