// Tests for the EnvironmentController

package environments

import (
	"log"
	"testing"

	"github.com/broadinstitute/sherlock/internal/db"
	"github.com/broadinstitute/sherlock/internal/tools"

	"github.com/golang-migrate/migrate/v4"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestIntegrationCreateEnvironments(t *testing.T) {

	// skip integration tests if go test is invoked with -short flag
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	runMigrations(t)

	testCases := []struct {
		name                string
		request             CreateEnvironmentRequest
		expectedError       error
		expectedEnvironment Environment
	}{
		{
			name: "creates a valid environment",
			request: CreateEnvironmentRequest{
				Name: "terra-juyang-opera-fish",
			},
			expectedError: nil,
			expectedEnvironment: Environment{
				Name: "terra-juyang-opera-fish",
			},
		},
	}

	// Testing Code
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testApp := initTestApp()
			newEnvironment, err := testApp.Environments.CreateNew(testCase.request)

			assert.Equal(t, testCase.expectedEnvironment.Name, newEnvironment.Name)
			assert.Equal(t, testCase.expectedError, err)
		})
	}
}

//
// Test Environment Setup
//

// TODO: this is technically all integration testing-related and could be shared
// and moved into a testing package

var (
	Config = viper.New()
)

// we need to check that the db is in a sane state and run migrations if it's not
func runMigrations(t *testing.T) {
	// The following steps initialize the database for use in the
	// sherlock server integration test suite
	// TODO pull this from config with viper

	// when running tests workdir is the package directory ie cmd/server
	// so a relative path to changelogs is needed.
	// TODO cleaner method to supply path to changelogs and run migration in tests
	if err := db.ApplyMigrations("../../db/migrations", Config); err == migrate.ErrNoChange {
		t.Log("no migration to apply, continuing...")
	} else if err != nil {
		t.Fatalf("error migrating database: %v", err)
	}
}

// only load the Controller we care about
type TestApplication struct {
	Environments *EnvironmentController
	DB           *gorm.DB
}

// New returns a new instance of the core sherlock application
func initTestApp() *TestApplication {
	dbConn, err := db.Connect(Config)
	if err != nil {
		log.Fatalf("error connecting to database: %v", err)
	}

	app := &TestApplication{
		DB: dbConn,
	}

	app.registerControllers()

	// nuke the DB in case it's dirty

	if err := tools.Truncate(dbConn); err != nil {
		log.Fatalf("error truncating database: %v", err)
	}

	return app
}

func (testApp *TestApplication) registerControllers() {
	testApp.Environments = NewController(testApp.DB)
}

// initialization for Tests, mostly for Viper so we can get the proper DB conn.
func init() {
	// viper will auto parse ENV VARS prefixed with SHERLOCK
	// into config
	Config.SetEnvPrefix("sherlock")

	// TODO: this is supposed to be postgres, but POC
	Config.SetDefault("dbhost", "postgres")
	Config.SetDefault("dbuser", "sherlock")
	Config.SetDefault("dbname", "sherlock")
	Config.SetDefault("dbport", "5432")
	Config.SetDefault("dbssl", "disable")
	Config.SetDefault("dbinit", true)
	// this is badness
	Config.SetDefault("dbpassword", "password")

	Config.AutomaticEnv()
}
