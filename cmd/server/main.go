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

func main() {
	db, err := sqlx.Connect("pgx", os.Getenv("POSTGRESQL_URL"))
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer db.Close()

	log.Println("Successfully connected to database")

	log.Println("Executing database migration")
	m, err := migrate.New(
		"file://db/migrations",
		os.Getenv("POSTGRESQL_URL"),
	)
	if err != nil {
		log.Println(err)
	}

	if err := m.Up(); err != nil {
		log.Println(err)
	}

	log.Println("database migration complete")

	log.Println("starting sherlock server")
	http.HandleFunc("/", handler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello sherlock!")
}
