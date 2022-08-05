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
	go func() { _ = endless.ListenAndServe(":8081", livenessHandler{}) }()

	dbConn, err := db.Connect()
	if err != nil {
		log.Fatal().Msgf("%v", err)
	}

	app := sherlock.New(dbConn)
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
