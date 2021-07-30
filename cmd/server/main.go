package main

import (
	"fmt"
	"net/http"
	"os"

	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

const changelogLocation = "file://db/migrations"

func main() {
	db, err := sqlx.Connect("pgx", os.Getenv("POSTGRESQL_URL"))
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer db.Close()

	log.Println("Successfully connected to database")

	// just logging error rather than failing here as
	// a no changes migration is treated as an error..
	if err := applyMigrations(); err != nil {
		log.Println(err)
	}

	log.Println("starting sherlock server")
	http.HandleFunc("/", handler)

	log.Fatal(http.ListenAndServe(":8080", nil))
	log.Println("Listening on port 8080")

}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello sherlock!")
}

func applyMigrations() error {
	// check for environment flag whether to run migrations on app start up or not
	if _, ok := os.LookupEnv("SHERLOCK_INIT_DB"); !ok {
		log.Println("skipping database migration on startup, starting server...")
		return nil
	}

	log.Println("Executing database migration")
	m, err := migrate.New(
		changelogLocation,
		os.Getenv("POSTGRESQL_URL"),
	)
	if err != nil {
		return err
	}

	if err := m.Up(); err != nil {
		return err
	}

	log.Println("database migration complete")
	return nil
}
