// Tests for the EnvironmentController

package environments

import (
	"errors"
	"testing"

	"github.com/broadinstitute/sherlock/internal/testutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context
type EnvironmentTestSuite struct {
	suite.Suite
	VariableThatShouldStartAtFive int
}

func TestIntegrationCreateEnvironmentsSuite(t *testing.T) {
	suite.Run(t, new(EnvironmentTestSuite))
}

func (suite *EnvironmentTestSuite) SetupTest() {
	suite.VariableThatShouldStartAtFive = 556
}

func (suite *EnvironmentTestSuite) CreateEnvironmentTest() {
	assert.Equal(suite.T(), 600, suite.VariableThatShouldStartAtFive)
}

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
		{
			name:          "fails to create an environment with no name",
			request:       CreateEnvironmentRequest{},
			expectedError: errors.New("error saving to database: ERROR: null value in column \"name\" of relation \"environments\" violates not-null constraint (SQLSTATE 23502)"),
			expectedEnvironment: Environment{
				Name: "",
			},
		},
	}

	// Testing Code
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testApp := initTestApp(t)
			//defer testutils.Cleanup(t, testApp.db)

			newEnvironment, err := testApp.Environments.CreateNew(testCase.request)

			assert.Equal(t, testCase.expectedEnvironment.Name, newEnvironment.Name)
			assert.Equal(t, err, testCase.expectedError)

			environments, errTest := testApp.Environments.ListAll()
			for _, environment := range environments {
				t.Logf("%+v", environment)
			}
			t.Logf("%+v", newEnvironment)
			t.Log(errTest)
			t.Log(err)
			//t.Fatalf("debug test: %v", err)
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
	testutils.Truncate(app.db)

	return app
}
