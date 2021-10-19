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
	"github.com/golang-migrate/migrate/v4"
)

func main() {

	if err := db.ApplyMigrations("db/migrations", sherlock.Config); err != nil {
		// don't fail if there are no changes to apply
		if err == migrate.ErrNoChange {
			log.Println("no migration to apply, continuing...")
		} else {
			log.Println(err)
			os.Exit(1)
		}
	}

	app := sherlock.New()
	defer app.ShutdownStackdriver()

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
