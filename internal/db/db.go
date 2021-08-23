// Package db contains a variety of utility functions and test helpers for
// working with Sherlock's database
package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/spf13/viper"
	gormpg "gorm.io/driver/postgres"
	"gorm.io/gorm"

	// indirect import used to set proper migration data source
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v4/stdlib"
)

// ApplyMigrations is a utility function intended for use in integration tests and
// local development where changelogs can be applied to a local postgres instance
// during startup
func ApplyMigrations(changeLogPath string, config *viper.Viper) error {
	// check for environment flag whether to run migrations on app start up or not
	if dbInit := config.GetBool("dbinit"); !dbInit {
		log.Println("skipping database migration on startup, starting server...")
		return nil
	}

	// TODO use viper to handle the db config rather than pulling it
	// directly out of env

	changelogLocation := fmt.Sprintf("file://%s", changeLogPath)
	dbURL := buildDBConnectionString(config)

	log.Println("Executing database migration")
	// The below code is to ensure migrations run using the same
	// postgres driver (pgx) that gorm uses. golang-migrate uses
	// a different postgres driver by default
	db, err := sql.Open("pgx", dbURL)
	if err != nil {
		return err
	}
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return err
	}
	migrationPlan, err := migrate.NewWithDatabaseInstance(
		changelogLocation,
		config.GetString("dbname"),
		driver,
	)
	if err != nil {
		return err
	}

	if err := migrationPlan.Up(); err != nil {
		return err
	}

	log.Println("database migration complete")
	return nil
}

// Connect is a utility function that accepts a viper instance containing
// database configs and returns a gorm database connection
func Connect(config *viper.Viper) (*gorm.DB, error) {
	dbURL := buildDBConnectionString(config)
	dbConn, err := sql.Open("pgx", dbURL)
	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %v", err)
	}

	gormDB, err := gorm.Open(
		gormpg.New(gormpg.Config{
			Conn: dbConn,
		}),
		// This is to account for the fact that go and postgres have different
		// time stamp precision which causes issues in testing.
		// This is a fix to have gorm round down timestamps to postgres' millisecond
		// precision
		&gorm.Config{
			NowFunc: func() time.Time {
				return time.Now().Round(time.Millisecond)
			},
		},
	)
	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %v", err)
	}
	return gormDB, nil
}

func buildDBConnectionString(config *viper.Viper) string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		config.GetString("dbuser"),
		config.GetString("dbpassword"),
		config.GetString("dbhost"),
		config.GetString("dbport"),
		config.GetString("dbname"),
		config.GetString("dbssl"),
	)
}
