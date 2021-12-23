package main

import (
	"net/http"
	"os"

	"log"

	"github.com/broadinstitute/sherlock/internal/db"
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

	log.Println("starting sherlock server")

	log.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", app))

}
