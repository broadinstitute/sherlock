// Package db contains a variety of utility functions and test helpers for
// working with Sherlock's database
package db

import (
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	// indirect import to set the database driver to use when applying migrations
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// ApplyMigrations is a utility function intended for use in integration tests and
// local development where changelogs can be applied to a local postgres instance
// during startup
func ApplyMigrations(changeLogPath string) error {
	// check for environment flag whether to run migrations on app start up or not
	if _, ok := os.LookupEnv("SHERLOCK_INIT_DB"); !ok {
		log.Println("skipping database migration on startup, starting server...")
		return nil
	}

	// TODO use viper to handle the db config rather than pulling it
	// directly out of env

	changelogLocation := fmt.Sprintf("file://%s", changeLogPath)

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
