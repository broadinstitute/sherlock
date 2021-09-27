package main

import (
	"net/http"
	"os"

	"log"

	"github.com/broadinstitute/sherlock/internal/builds"
	"github.com/broadinstitute/sherlock/internal/db"
	"github.com/broadinstitute/sherlock/internal/deploys"
	"github.com/broadinstitute/sherlock/internal/environments"
	"github.com/broadinstitute/sherlock/internal/services"
	"github.com/broadinstitute/sherlock/internal/sherlock"
)

func main() {

	// just logging error rather than failing here as
	// a no changes migration is treated as an error..
	if err := db.ApplyMigrations("db/migrations", sherlock.Config); err != nil {
		log.Println(err)
		os.Exit(1)
	}

	app := sherlock.New()

	if _, err := services.Seed(app.DB); err != nil {
		log.Println(err)
	}

	if _, err := builds.Seed(app.DB); err != nil {
		log.Println(err)
	}

	if _, err := environments.Seed(app.DB); err != nil {
		log.Println(err)
	}

	if _, err := deploys.SeedServiceInstances(app.DB); err != nil {
		log.Println(err)
	}

	log.Println("starting sherlock server")

	log.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", app))

}
