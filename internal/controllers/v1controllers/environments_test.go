// Tests for the EnvironmentController

package v1controllers

import (
	"errors"
	"github.com/broadinstitute/sherlock/internal/db"
	"github.com/broadinstitute/sherlock/internal/models/v1models"
	"github.com/broadinstitute/sherlock/internal/serializers/v1serializers"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context
type EnvironmentTestSuite struct {
	suite.Suite
	testApp                   *TestApplication
	goodEnvironmentRequest    v1models.CreateEnvironmentRequest
	anotherEnvironmentRequest v1models.CreateEnvironmentRequest
	badEnvironmentRequest     v1models.CreateEnvironmentRequest
	notFoundID                int
}

// Test entry point
func TestFunctionalEnvironmentsSuite(t *testing.T) {
	// skip functional tests if go test is invoked with -short flag
	if testing.Short() {
		t.Skip("skipping functional test")
	}
	suite.Run(t, new(EnvironmentTestSuite))
}

// between-test initialization
func (suite *EnvironmentTestSuite) SetupTest() {
	suite.testApp = initEnvironmentsTestApp(suite.T())
	suite.goodEnvironmentRequest = v1models.CreateEnvironmentRequest{
		Name: faker.UUIDHyphenated(),
	}
	suite.anotherEnvironmentRequest = v1models.CreateEnvironmentRequest{
		Name: faker.UUIDHyphenated(),
	}
	suite.badEnvironmentRequest = v1models.CreateEnvironmentRequest{}
	suite.notFoundID = 1234567890 //unsure of a way to guarantee not-found-ness
}

func (suite *EnvironmentTestSuite) TearDownTest() {
	// each test runs in its own isolated transaction
	// this ensures we cleanup after each test as it completes
	suite.testApp.DB.Rollback()
}

//
// Test Environment Setup
//

// connect to DB and create the Application
func initEnvironmentsTestApp(t *testing.T) *TestApplication {
	dbConn := db.ConnectFromTest(t)
	// ensures each test will run in it's own isolated transaction
	// The transaction will be rolled back after each test
	// regardless of pass or fail
	dbConn = dbConn.Begin()
	return &TestApplication{
		Environments: NewEnvironmentController(dbConn),
		DB:           dbConn,
	}
}

//
// The Actual Tests
//

func (suite *EnvironmentTestSuite) TestFunctionalCreateEnvironments() {
	testCases := []struct {
		name                string
		requests            []v1models.CreateEnvironmentRequest
		expectedError       error
		expectedEnvironment v1models.Environment
	}{
		{
			name: "creates a valid environment",
			requests: []v1models.CreateEnvironmentRequest{
				{
					Name: "terra-juyang-opera-fish",
				},
			},
			expectedError: nil,
			expectedEnvironment: v1models.Environment{
				Name: "terra-juyang-opera-fish",
			},
		},
		{
			name: "fails to create an environment with no name",
			requests: []v1models.CreateEnvironmentRequest{
				{},
			},
			expectedError: errors.New("error saving to database: ERROR: null value in column \"name\" of relation \"environments\" violates not-null constraint (SQLSTATE 23502)"),
			expectedEnvironment: v1models.Environment{
				Name: "",
			},
		},
		{
			name: "fails to create an environment with duplicate name",
			requests: []v1models.CreateEnvironmentRequest{
				{
					Name: "terra-juyang-opera-fish",
				},
				{
					Name: "terra-juyang-opera-fish",
				},
			},
			expectedError: errors.New("error saving to database: ERROR: duplicate key value violates unique constraint \"environments_name_key\" (SQLSTATE 23505)"),
			expectedEnvironment: v1models.Environment{
				Name: "",
			},
		},
		{
			name: "creates environments with unique names",
			requests: []v1models.CreateEnvironmentRequest{
				{
					Name: "terra-juyang-opera-fish",
				},
				{
					Name: "terra-juyang-maggot-sawfly",
				},
			},
			expectedError: nil,
			expectedEnvironment: v1models.Environment{
				Name: "terra-juyang-maggot-sawfly",
			},
		},
	}

	// Testing Code
	for _, testCase := range testCases {
		// creating a temporary test app instance with its own transaction for each
		// testcase so they don't step on eachother
		tempApp := initEnvironmentsTestApp(suite.T())
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
		tempApp.DB.Rollback()
	}
}

func (suite *EnvironmentTestSuite) TestFunctionalEnvironmentGetByName() {
	suite.Run("GetByName gets an environment by name", func() {

		_, err := suite.testApp.Environments.CreateNew(suite.goodEnvironmentRequest)
		assert.NoError(suite.T(), err)

		foundEnvironment, err := suite.testApp.Environments.GetByName(suite.goodEnvironmentRequest.Name)

		assert.Equal(suite.T(), foundEnvironment.Name, suite.goodEnvironmentRequest.Name)

		assert.NoError(suite.T(), err)
	})

	suite.Run("GetByName returns error if not found", func() {
		var envRequest v1models.CreateEnvironmentRequest
		err := faker.FakeData(&envRequest)
		suite.Require().NoError(err)

		_, err = suite.testApp.Environments.CreateNew(envRequest)
		assert.NoError(suite.T(), err)

		foundEnvironment, err := suite.testApp.Environments.GetByName("this-doesnt-exist")

		assert.Equal(suite.T(), foundEnvironment.Name, "")
		assert.Equal(suite.T(), errors.New("record not found"), err)
	})
}

func (suite *EnvironmentTestSuite) TestFunctionalEnvironmentGetByID() {
	suite.Run("GetByID gets an environment by name", func() {
		db.Truncate(suite.T(), suite.testApp.DB)

		newEnvironment, err := suite.testApp.Environments.CreateNew(suite.goodEnvironmentRequest)
		assert.NoError(suite.T(), err)

		foundEnvironment, err := suite.testApp.Environments.GetByID(newEnvironment.ID)

		assert.Equal(suite.T(), foundEnvironment.ID, newEnvironment.ID)

		assert.NoError(suite.T(), err)
	})

	suite.Run("GetByID returns error if not found", func() {
		db.Truncate(suite.T(), suite.testApp.DB)

		_, err := suite.testApp.Environments.CreateNew(suite.goodEnvironmentRequest)
		assert.NoError(suite.T(), err)

		foundEnvironment, err := suite.testApp.Environments.GetByID(suite.notFoundID)

		assert.Equal(suite.T(), foundEnvironment.ID, 0)
		assert.Equal(suite.T(), errors.New("record not found"), err)
	})
}

func (suite *EnvironmentTestSuite) TestFunctionalEnvironmentListAll() {
	suite.Run("ListAll returns nothing", func() {
		db.Truncate(suite.T(), suite.testApp.DB)

		foundEnvironments, err := suite.testApp.Environments.ListAll()

		assert.Equal(suite.T(), 0, len(foundEnvironments))
		assert.NoError(suite.T(), err)
	})

	suite.Run("ListAll returns one Environment", func() {
		db.Truncate(suite.T(), suite.testApp.DB)

		_, err := suite.testApp.Environments.CreateNew(suite.goodEnvironmentRequest)
		assert.NoError(suite.T(), err)

		foundEnvironments, err := suite.testApp.Environments.ListAll()

		assert.Equal(suite.T(), len(foundEnvironments), 1)
		assert.Equal(suite.T(), suite.goodEnvironmentRequest.Name, foundEnvironments[0].Name)
		assert.NoError(suite.T(), err)
	})

	suite.Run("ListAll returns many Environments", func() {
		db.Truncate(suite.T(), suite.testApp.DB)

		var randomEnvRequest v1models.CreateEnvironmentRequest
		err := faker.FakeData(&randomEnvRequest)
		suite.Require().NoError(err)

		startingEnvironments, _ := suite.testApp.Environments.ListAll()

		_, err = suite.testApp.Environments.CreateNew(randomEnvRequest)
		assert.NoError(suite.T(), err)
		_, err = suite.testApp.Environments.CreateNew(suite.anotherEnvironmentRequest)
		assert.NoError(suite.T(), err)

		foundEnvironments, err := suite.testApp.Environments.ListAll()

		assert.Equal(suite.T(), len(startingEnvironments)+2, len(foundEnvironments))
		assert.NoError(suite.T(), err)
	})
}

func (suite *EnvironmentTestSuite) TestFunctionalEnvironmentDoesEnvironmentExist() {
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
func (suite *EnvironmentTestSuite) TestFunctionalEnvironmentSerialize() {
	suite.Run("Serialize returns JSON one environment", func() {

		newEnvironment, _ := suite.testApp.Environments.CreateNew(suite.goodEnvironmentRequest)

		environmentResponses := suite.testApp.Environments.Serialize(newEnvironment)

		// check that we get expected number of elements in the slice + correct object type
		assert.Equal(suite.T(), 1, len(environmentResponses))
		assert.IsType(suite.T(), v1serializers.EnvironmentResponse{}, environmentResponses[0])
	})

	suite.Run("Serialize returns JSON for many environments", func() {

		var environments []v1models.Environment
		var newEnvironment v1models.Environment

		newEnvironment, _ = suite.testApp.Environments.CreateNew(suite.goodEnvironmentRequest)
		environments = append(environments, newEnvironment)
		newEnvironment, _ = suite.testApp.Environments.CreateNew(suite.anotherEnvironmentRequest)
		environments = append(environments, newEnvironment)

		environmentResponses := suite.testApp.Environments.Serialize(environments...)

		// check that we get expected number of elements in the slice + correct object type
		assert.Equal(suite.T(), 2, len(environmentResponses))
		assert.IsType(suite.T(), v1serializers.EnvironmentResponse{}, environmentResponses[0])
	})

	suite.Run("Serialize returns empty environment for bad environments", func() {

		newEnvironment, _ := suite.testApp.Environments.CreateNew(suite.badEnvironmentRequest)
		environmentResponses := suite.testApp.Environments.Serialize(newEnvironment)

		assert.Equal(suite.T(), 1, len(environmentResponses))
		assert.IsType(suite.T(), v1serializers.EnvironmentResponse{}, environmentResponses[0])
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
