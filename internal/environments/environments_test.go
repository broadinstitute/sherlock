// Tests for the EnvironmentController

package environments

import (
	"errors"
	"testing"

	"github.com/broadinstitute/sherlock/internal/testutils"
	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context
type EnvironmentTestSuite struct {
	suite.Suite
	testApp                   *TestApplication
	goodEnvironmentRequest    CreateEnvironmentRequest
	anotherEnvironmentRequest CreateEnvironmentRequest
	badEnvironmentRequest     CreateEnvironmentRequest
}

// Test entry point
func TestIntegrationEnvironmentsSuite(t *testing.T) {
	// skip integration tests if go test is invoked with -short flag
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	suite.Run(t, new(EnvironmentTestSuite))
}

// between-test initialization
func (suite *EnvironmentTestSuite) SetupTest() {
	suite.testApp = initTestApp(suite.T())
	suite.goodEnvironmentRequest = CreateEnvironmentRequest{
		Name: faker.UUIDHyphenated(),
	}
	suite.anotherEnvironmentRequest = CreateEnvironmentRequest{
		Name: faker.UUIDHyphenated(),
	}
	suite.badEnvironmentRequest = CreateEnvironmentRequest{}
}

func (suite *EnvironmentTestSuite) TearDownTest() {
	// each test runs in its own isolated transaction
	// this ensures we cleanup after each test as it completes
	suite.testApp.db.Rollback()
}

//
// Test Environment Setup
//

// only load the Controller we care about
type TestApplication struct {
	Environments *EnvironmentController
	db           *gorm.DB
}

// connect to DB and create the Application
func initTestApp(t *testing.T) *TestApplication {
	dbConn := testutils.ConnectAndMigrate(t)
	// ensures each test will run in it's own isolated transaction
	// The transaction will be rolled back after each test
	// regardless of pass or fail
	dbConn = dbConn.Begin()
	return &TestApplication{
		Environments: NewController(dbConn),
		db:           dbConn,
	}
}

//
// The Actual Tests
//

func (suite *EnvironmentTestSuite) TestIntegrationCreateEnvironments() {
	testCases := []struct {
		name                string
		requests            []CreateEnvironmentRequest
		expectedError       error
		expectedEnvironment Environment
	}{
		{
			name: "creates a valid environment",
			requests: []CreateEnvironmentRequest{
				{
					Name: "terra-juyang-opera-fish",
				},
			},
			expectedError: nil,
			expectedEnvironment: Environment{
				Name: "terra-juyang-opera-fish",
			},
		},
		{
			name: "fails to create an environment with no name",
			requests: []CreateEnvironmentRequest{
				{},
			},
			expectedError: errors.New("error saving to database: ERROR: null value in column \"name\" of relation \"environments\" violates not-null constraint (SQLSTATE 23502)"),
			expectedEnvironment: Environment{
				Name: "",
			},
		},
		{
			name: "fails to create an environment with duplicate name",
			requests: []CreateEnvironmentRequest{
				{
					Name: "terra-juyang-opera-fish",
				},
				{
					Name: "terra-juyang-opera-fish",
				},
			},
			expectedError: errors.New("error saving to database: ERROR: duplicate key value violates unique constraint \"environments_name_key\" (SQLSTATE 23505)"),
			expectedEnvironment: Environment{
				Name: "terra-juyang-opera-fish",
			},
		},
		{
			name: "creates environments with unique names",
			requests: []CreateEnvironmentRequest{
				{
					Name: "terra-juyang-opera-fish",
				},
				{
					Name: "terra-juyang-maggot-sawfly",
				},
			},
			expectedError: nil,
			expectedEnvironment: Environment{
				Name: "terra-juyang-maggot-sawfly",
			},
		},
	}

	// Testing Code
	for _, testCase := range testCases {
		// creating a temporary test app instance with its own transaction for each
		// testcase so they don't step on eachother
		tempApp := initTestApp(suite.T())
		suite.Run(testCase.name, func() {

			// create all non-final environments as setup
			for _, request := range testCase.requests[:len(testCase.requests)-1] {
				_, err := tempApp.Environments.CreateNew(request)
				assert.NoError(suite.T(), err)
			}

			// create and test the last environment
			newEnvironment, err := tempApp.Environments.CreateNew(testCase.requests[len(testCase.requests)-1])
			assert.Equal(suite.T(), testCase.expectedEnvironment.Name, newEnvironment.Name)
			assert.Equal(suite.T(), testCase.expectedError, err)
		})
		tempApp.db.Rollback()
	}
}

func (suite *EnvironmentTestSuite) TestIntegrationEnvironmentGetByName() {
	suite.Run("GetByName gets an environment by name", func() {

		_, err := suite.testApp.Environments.CreateNew(suite.goodEnvironmentRequest)
		assert.NoError(suite.T(), err)

		foundEnvironment, err := suite.testApp.Environments.GetByName(suite.goodEnvironmentRequest.Name)

		assert.Equal(suite.T(), foundEnvironment.Name, suite.goodEnvironmentRequest.Name)

		assert.NoError(suite.T(), err)
	})

	suite.Run("GetByName returns error if not found", func() {
		var envRequest CreateEnvironmentRequest
		err := faker.FakeData(&envRequest)
		suite.Require().NoError(err)

		_, err = suite.testApp.Environments.CreateNew(envRequest)
		assert.NoError(suite.T(), err)

		foundEnvironment, err := suite.testApp.Environments.GetByName("this-doesnt-exist")

		assert.Equal(suite.T(), foundEnvironment.Name, "")
		assert.Equal(suite.T(), errors.New("record not found"), err)
	})
}

func (suite *EnvironmentTestSuite) TestIntegrationEnvironmentListAll() {
	suite.Run("ListAll returns nothing", func() {

		foundEnvironments, err := suite.testApp.Environments.ListAll()

		assert.GreaterOrEqual(suite.T(), len(foundEnvironments), 0)
		assert.NoError(suite.T(), err)
	})

	suite.Run("ListAll returns one Environment", func() {

		_, err := suite.testApp.Environments.CreateNew(suite.goodEnvironmentRequest)
		assert.NoError(suite.T(), err)

		foundEnvironments, err := suite.testApp.Environments.ListAll()

		assert.GreaterOrEqual(suite.T(), len(foundEnvironments), 1)
		assert.NoError(suite.T(), err)
	})

	suite.Run("ListAll returns many Environments", func() {

		var randomEnvRequest CreateEnvironmentRequest
		err := faker.FakeData(&randomEnvRequest)
		suite.Require().NoError(err)

		_, err = suite.testApp.Environments.CreateNew(randomEnvRequest)
		assert.NoError(suite.T(), err)
		_, err = suite.testApp.Environments.CreateNew(suite.anotherEnvironmentRequest)
		assert.NoError(suite.T(), err)

		foundEnvironments, err := suite.testApp.Environments.ListAll()

		assert.GreaterOrEqual(suite.T(), len(foundEnvironments), 3)
		assert.NoError(suite.T(), err)
	})
}

func (suite *EnvironmentTestSuite) TestIntegrationEnvironmentDoesEnvironmentExist() {
	suite.Run("EnvironmentDoesExist returns true when exists", func() {

		newEnvironment, _ := suite.testApp.Environments.CreateNew(suite.goodEnvironmentRequest)

		environmentID, doesEnvironmentExist := suite.testApp.Environments.DoesEnvironmentExist(suite.goodEnvironmentRequest.Name)

		assert.Equal(suite.T(), environmentID, newEnvironment.ID)
		assert.Equal(suite.T(), doesEnvironmentExist, true)
	})

	suite.Run("EnvironmentDoesExist returns false when not exists", func() {

		_, err := suite.testApp.Environments.CreateNew(suite.anotherEnvironmentRequest)
		assert.NoError(suite.T(), err)

		environmentID, doesEnvironmentExist := suite.testApp.Environments.DoesEnvironmentExist("no-environment-here")

		assert.Equal(suite.T(), environmentID, 0)
		assert.Equal(suite.T(), doesEnvironmentExist, false)
	})
}

// Note: Since serialize is it's own file and rather big, we're limiting tests here
// to expecting the correct response types.
func (suite *EnvironmentTestSuite) TestIntegrationEnvironmentSerialize() {
	suite.Run("Serialize returns JSON one environment", func() {

		newEnvironment, _ := suite.testApp.Environments.CreateNew(suite.goodEnvironmentRequest)

		environmentResponses := suite.testApp.Environments.serialize(newEnvironment)

		// check that we get expected number of elements in the slice + correct object type
		assert.Equal(suite.T(), 1, len(environmentResponses))
		assert.IsType(suite.T(), EnvironmentResponse{}, environmentResponses[0])
	})

	suite.Run("Serialize returns JSON for many environments", func() {

		var environments []Environment
		var newEnvironment Environment

		newEnvironment, _ = suite.testApp.Environments.CreateNew(suite.goodEnvironmentRequest)
		environments = append(environments, newEnvironment)
		newEnvironment, _ = suite.testApp.Environments.CreateNew(suite.anotherEnvironmentRequest)
		environments = append(environments, newEnvironment)

		environmentResponses := suite.testApp.Environments.serialize(environments...)

		// check that we get expected number of elements in the slice + correct object type
		assert.Equal(suite.T(), 2, len(environmentResponses))
		assert.IsType(suite.T(), EnvironmentResponse{}, environmentResponses[0])
	})

	suite.Run("Serialize returns empty environment for bad environments", func() {

		newEnvironment, _ := suite.testApp.Environments.CreateNew(suite.badEnvironmentRequest)
		environmentResponses := suite.testApp.Environments.serialize(newEnvironment)

		assert.Equal(suite.T(), 1, len(environmentResponses))
		assert.IsType(suite.T(), EnvironmentResponse{}, environmentResponses[0])
	})
}

func (suite *EnvironmentTestSuite) TestFindOrCreate() {
	suite.Run("retrieves an environment that already exists", func() {

		existingEnvironment, err := suite.testApp.Environments.CreateNew(suite.goodEnvironmentRequest)
		suite.Require().NoError(err)

		foundEnvironmentID, err := suite.testApp.Environments.FindOrCreate(suite.goodEnvironmentRequest.Name)
		suite.Require().NoError(err)

		suite.Assert().Equal(existingEnvironment.ID, foundEnvironmentID)
	})

	suite.Run("creates a new environment that doesn't exist already", func() {

		newEnvironmentID, err := suite.testApp.Environments.FindOrCreate(faker.Word())
		suite.Assert().NoError(err)

		// assert the env was created by verifying it has a non-zero id
		suite.Assert().NotEqual(0, newEnvironmentID)
	})
}
