package testutils

// db.go contains a variety of methods for helping connect to and setup a database for use in
// integration tests. It also facilitates the running of integration tests without dependencies
// on the top level package sherlock

import (
	"testing"

	"github.com/broadinstitute/sherlock/internal/db"
	"github.com/golang-migrate/migrate/v4"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

// db contains a utility for retrieving a gorm.DB connection on demand for use in test setup

var (
	config = viper.New()
)

func initConfig() {
	config.SetEnvPrefix("sherlock")

	config.SetDefault("dbhost", "postgres")
	config.SetDefault("dbuser", "sherlock")
	config.SetDefault("dbname", "sherlock")
	config.SetDefault("dbport", "5432")
	config.SetDefault("dbssl", "disable")
	config.SetDefault("dbinit", false)

	config.AutomaticEnv()
}

// ConnectAndMigrate will parse config to attempt to establish a connection to the test database
// It will cause the test to fail fatally if an error is encountered. It will then attempt to apply
// all migrations in /db/migrations to that database and will fail similarly if an error is encountered
func ConnectAndMigrate(t *testing.T) *gorm.DB {
	t.Helper()

	// setup config for db connection
	initConfig()

	dbConn, err := db.Connect(config)
	if err != nil {
		t.Fatalf("error retrieving db connection for testing: %v", err)
	}
	if err := db.ApplyMigrations("../../db/migrations", config); err == migrate.ErrNoChange {
		t.Log("no migration to apply, continuing...")
	} else if err != nil {
		t.Fatalf("error migrating database: %v", err)
	}

	return dbConn
}

// Truncate cleans up tables after integration tests
func Truncate(db *gorm.DB) error {
	// gorm doesn't seem to support truncate operations which are essential to cleaning up after
	// integration tests (and the only use case of this function so doing it with raw sql)
	truncateStatement := "TRUNCATE TABLE services, builds, environments, service_instances, deploys, clusters, allocation_pools"
	err := db.Exec(truncateStatement).Error

	return err
}

// Cleanup is intended to be deferred with each test run so that we can
// ensure each case starts with a clean database
func Cleanup(t *testing.T, dbConn *gorm.DB) {
	if err := Truncate(dbConn); err != nil {
		t.Fatalf("error cleaning db after test run: %v", err)
	}
}
