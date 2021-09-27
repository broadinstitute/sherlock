// Tests for the EnvironmentController

package environments

import (
	"testing"

	"github.com/broadinstitute/sherlock/internal/testutils"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestIntegrationCreateEnvironments(t *testing.T) {

	// skip integration tests if go test is invoked with -short flag
	if testing.Short() {
		t.Skip("skipping integration test")
	}

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
			testApp := initTestApp(t)
			defer testutils.Cleanup(t, testApp.db)

			newEnvironment, err := testApp.Environments.CreateNew(testCase.request)

			assert.Equal(t, testCase.expectedEnvironment.Name, newEnvironment.Name)
			assert.Equal(t, testCase.expectedError, err)
		})
	}
}

//
// Test Environment Setup
//

// only load the Controller we care about
type TestApplication struct {
	Environments *EnvironmentController
	db           *gorm.DB
}

func initTestApp(t *testing.T) *TestApplication {
	dbConn := testutils.ConnectAndMigrate(t)
	app := &TestApplication{
		Environments: NewController(dbConn),
		db:           dbConn,
	}

	return app
}
