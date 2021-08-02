package main

import (
	"fmt"
	"net/http"
	"os"

	"log"

	"github.com/broadinstitute/sherlock/internal/db"
	"github.com/broadinstitute/sherlock/internal/sherlock"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

func main() {
	dbConn, err := sqlx.Connect("pgx", os.Getenv("POSTGRESQL_URL"))
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer dbConn.Close()

	log.Println("Successfully connected to database")

	// just logging error rather than failing here as
	// a no changes migration is treated as an error..
	if err := db.ApplyMigrations("db/migrations"); err != nil {
		log.Println(err)
	}

	log.Println("starting sherlock server")
	app := sherlock.New(dbConn)

	log.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", app))

}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello sherlock!")
}
