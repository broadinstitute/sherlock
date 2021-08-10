package main

import (
	"net/http"

	"log"

	"github.com/broadinstitute/sherlock/internal/db"
	"github.com/broadinstitute/sherlock/internal/sherlock"
	"github.com/broadinstitute/sherlock/internal/tools"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {

	// just logging error rather than failing here as
	// a no changes migration is treated as an error..
	if err := db.ApplyMigrations("db/migrations"); err != nil {
		log.Println(err)
	}

	app := sherlock.New()

	if _, err := tools.SeedServices(app.Repository); err != nil {
		log.Println(err)
	}

	log.Println("starting sherlock server")

	log.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", app))

}
