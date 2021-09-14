package main

import (
	"net/http"
	"os"

	"log"

	"github.com/broadinstitute/sherlock/internal/db"
	"github.com/broadinstitute/sherlock/internal/sherlock"
	"github.com/broadinstitute/sherlock/internal/tools"
)

func main() {

	// just logging error rather than failing here as
	// a no changes migration is treated as an error..
	if err := db.ApplyMigrations("db/migrations", sherlock.Config); err != nil {
		log.Println(err)
		os.Exit(1)
	}

	app := sherlock.New()

	if _, err := tools.SeedServices(app.DB); err != nil {
		log.Println(err)
	}

	if _, err := tools.SeedBuilds(app.DB); err != nil {
		log.Println(err)
	}

	log.Println("starting sherlock server")

	log.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", app))

}
