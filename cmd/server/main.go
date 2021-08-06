package main

import (
	"net/http"
	"os"

	"log"

	"github.com/broadinstitute/sherlock/internal/db"
	"github.com/broadinstitute/sherlock/internal/sherlock"
	"github.com/broadinstitute/sherlock/internal/tools"
	_ "github.com/jackc/pgx/v4/stdlib"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=postgres user=sherlock password=password dbname=sherlock port=5432 sslmode=disable"
	dbConn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	log.Println("Successfully connected to database")

	// just logging error rather than failing here as
	// a no changes migration is treated as an error..
	if err := db.ApplyMigrations("db/migrations"); err != nil {
		log.Println(err)
	}

	repository := db.NewRepository(dbConn)
	app := sherlock.New(repository)

	if _, err := tools.SeedServices(app.Repository); err != nil {
		log.Println(err)
	}

	log.Println("starting sherlock server")

	log.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", app))

}
