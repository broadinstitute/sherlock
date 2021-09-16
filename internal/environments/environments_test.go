// Tests for the EnvironmentController

package environments

import (
	"log"
	"testing"

	//"github.com/broadinstitute/sherlock/internal/sherlock"
	"github.com/broadinstitute/sherlock/internal/db"
	"github.com/broadinstitute/sherlock/internal/tools"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestCreateEnvironments(t *testing.T) {
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

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testApp := initTestApp()
			newEnvironment, err := testApp.Environments.CreateNew(testCase.request)

			assert.Equal(t, testCase.expectedEnvironment.Name, newEnvironment.Name)
			assert.Equal(t, testCase.expectedError, err)
		})
	}
}

// Test Environment Setup
var (
	Config = viper.New()
)

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

	// initialize the gin router and store it in our app struct
	//app.buildRouter()

	tools.Truncate(dbConn)

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
	Config.SetDefault("dbhost", "localhost")
	Config.SetDefault("dbuser", "sherlock")
	Config.SetDefault("dbname", "sherlock")
	Config.SetDefault("dbport", "5432")
	Config.SetDefault("dbssl", "disable")
	Config.SetDefault("dbinit", true)
	// this is badness
	Config.SetDefault("dbpassword", "password")

	Config.AutomaticEnv()
}
