// Package db contains a variety of utility functions and test helpers for
// working with Sherlock's database
package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/spf13/viper"
	gormpg "gorm.io/driver/postgres"
	"gorm.io/gorm"

	// indirect import used to set proper migration data source
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v4/stdlib"
)

const databaseName string = "sherlock"

// Repository is a wrapper around a connection to a data base
//  that can support mocking database connections in tests
type Repository struct {
	DB *gorm.DB
}

// NewRepository takes a gorm db connection and returns it
// wrapped in a repository struct
func NewRepository(db *gorm.DB) *Repository {
	return &Repository{DB: db}
}

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
	// The below code is to ensure migrations run using the same
	// postgres driver (pgx) that gorm uses. golang-migrate uses
	// a different postgres driver by default
	db, err := sql.Open("pgx", os.Getenv("POSTGRESQL_URL"))
	if err != nil {
		return err
	}
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return err
	}
	m, err := migrate.NewWithDatabaseInstance(
		changelogLocation,
		databaseName,
		driver,
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

func Connect(config *viper.Viper) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		config.GetString("dbhost"),
		config.GetString("dbuser"),
		config.GetString("dbpassword"),
		config.GetString("dbname"),
		config.GetString("dbport"),
		config.GetString("dbssl"),
	)
	dbConn, err := gorm.Open(gormpg.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %v", err)
	}
	return dbConn, nil
}
