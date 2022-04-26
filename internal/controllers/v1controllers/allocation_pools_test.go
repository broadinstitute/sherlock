// Tests for the AllocationPoolController

package v1controllers

import (
	"errors"
	"github.com/broadinstitute/sherlock/internal/models/v1models"
	"testing"

	"github.com/broadinstitute/sherlock/internal/environments"
	"github.com/broadinstitute/sherlock/internal/testutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context
type AllocationPoolTestSuite struct {
	suite.Suite
	testApp                      *testApplication
	goodAllocationPoolRequest    v1models.CreateAllocationPoolRequest
	goodEnvironmentRequest       v1models.CreateEnvironmentRequest
	anotherAllocationPoolRequest v1models.CreateAllocationPoolRequest
	badAllocationPoolRequest     v1models.CreateAllocationPoolRequest
	notFoundID                   int
}

// Test entry point
func TestIntegrationAllocationPoolsSuite(t *testing.T) {
	// skip integration tests if go test is invoked with -short flag
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	suite.Run(t, new(AllocationPoolTestSuite))
}

// between-test initialization
func (suite *AllocationPoolTestSuite) SetupTest() {
	suite.testApp = initTestAllocationPoolsApp(suite.T())
	suite.goodAllocationPoolRequest = v1models.CreateAllocationPoolRequest{
		Name: "swatomation 1.0",
	}
	suite.goodEnvironmentRequest = v1models.CreateEnvironmentRequest{
		Name: "terra-juyang-prime-sawfly",
	}
	suite.anotherAllocationPoolRequest = v1models.CreateAllocationPoolRequest{
		Name: "new swatomation-FiaB",
	}
	suite.badAllocationPoolRequest = v1models.CreateAllocationPoolRequest{}
	suite.notFoundID = 1234567890 //unsure of a way to guarantee not-found-ness
}

func (suite *AllocationPoolTestSuite) TearDownTest() {
	// each test runs in its own isolated transaction
	// this ensures we cleanup after each test as it completes
	suite.testApp.db.Rollback()
}

//
// Test AllocationPool Setup
//

// connect to DB and create the Application
func initTestAllocationPoolsApp(t *testing.T) *testApplication {
	dbConn := testutils.ConnectAndMigrate(t)

	// ensures each test will run in it's own isolated transaction
	// The transaction will be rolled back after each test
	// regardless of pass or fail
	dbConn = dbConn.Begin()
	app := &testApplication{
		AllocationPools: NewAllocationPoolController(dbConn),
		Environments:    environments.NewController(dbConn),
		db:              dbConn,
	}

	return app
}

//
// The Actual Tests
//

func (suite *AllocationPoolTestSuite) TestIntegrationCreateAllocationPools() {
	suite.Run("creates a valid allocationPool", func() {
		testutils.Cleanup(suite.T(), suite.testApp.db)

		newAllocationPool, err := suite.testApp.AllocationPools.CreateNew(suite.goodAllocationPoolRequest)
		assert.NoError(suite.T(), err)

		assert.Equal(suite.T(), suite.goodAllocationPoolRequest.Name, newAllocationPool.Name)
		assert.NoError(suite.T(), err)
	})

	suite.Run("fails to create a allocationPool with no name", func() {
		testutils.Cleanup(suite.T(), suite.testApp.db)

		expectedError := errors.New("error saving to database: ERROR: null value in column \"name\" of relation \"allocation_pools\" violates not-null constraint (SQLSTATE 23502)")
		namelessAllocationPoolRequest := suite.goodAllocationPoolRequest
		namelessAllocationPoolRequest.Name = ""

		newAllocationPool, err := suite.testApp.AllocationPools.CreateNew(namelessAllocationPoolRequest)

		assert.Equal(suite.T(), "", newAllocationPool.Name)
		assert.Equal(suite.T(), expectedError, err)
	})

	suite.Run("fails to create a allocationPool with duplicate name", func() {
		testutils.Cleanup(suite.T(), suite.testApp.db)

		expectedError := errors.New("error saving to database: ERROR: duplicate key value violates unique constraint \"allocation_pools_name_key\" (SQLSTATE 23505)")

		_, err := suite.testApp.AllocationPools.CreateNew(suite.goodAllocationPoolRequest)
		assert.NoError(suite.T(), err)
		newAllocationPool, err := suite.testApp.AllocationPools.CreateNew(suite.goodAllocationPoolRequest)

		assert.Empty(suite.T(), newAllocationPool.Name)
		assert.Equal(suite.T(), expectedError, err)
	})

	suite.Run("fails to create a allocationPool with duplicate name", func() {
		testutils.Cleanup(suite.T(), suite.testApp.db)

		_, err := suite.testApp.AllocationPools.CreateNew(suite.goodAllocationPoolRequest)
		assert.NoError(suite.T(), err)
		_, err = suite.testApp.AllocationPools.CreateNew(suite.anotherAllocationPoolRequest)
		assert.NoError(suite.T(), err)
	})

	suite.Run("create a cluster with a new embedded environment", func() {
		testutils.Cleanup(suite.T(), suite.testApp.db)

		suite.goodAllocationPoolRequest.Environments = []v1models.Environment{suite.goodEnvironmentRequest.EnvironmentReq()}

		newAllocationPool, err := suite.testApp.AllocationPools.CreateNew(suite.goodAllocationPoolRequest)
		assert.NoError(suite.T(), err)

		assert.Equal(suite.T(), suite.goodAllocationPoolRequest.Name, newAllocationPool.Name)
		assert.Equal(suite.T(), suite.goodAllocationPoolRequest.Environments[0].Name, suite.goodEnvironmentRequest.Name)
	})
}

func (suite *AllocationPoolTestSuite) TestAddByEnvironmentID() {
	suite.Run("creates an allocationPool and environment seperately and then joins", func() {
		testutils.Cleanup(suite.T(), suite.testApp.db)

		newAllocationPool, _ := suite.testApp.AllocationPools.CreateNew(suite.goodAllocationPoolRequest)
		newEnvironment, _ := suite.testApp.Environments.CreateNew(suite.goodEnvironmentRequest)

		// add the environment to the allocation pool
		updatedAllocationPool, err := suite.testApp.AllocationPools.AddEnvironmentByID(newAllocationPool, newEnvironment.ID)
		assert.NoError(suite.T(), err)

		// update the objects from the db
		updatedEnvironment, _ := suite.testApp.Environments.GetByName(newEnvironment.Name)

		assert.Equal(suite.T(), newAllocationPool.ID, *updatedEnvironment.AllocationPoolID)
		require.NotEmpty(suite.T(), updatedAllocationPool.Environments)
		assert.Equal(suite.T(), updatedEnvironment.ID, (updatedAllocationPool.Environments)[0].ID)
	})

	suite.Run("reassigns environment to different allocation pool", func() {
		testutils.Cleanup(suite.T(), suite.testApp.db)

		newAllocationPool, _ := suite.testApp.AllocationPools.CreateNew(suite.goodAllocationPoolRequest)
		anotherAllocationPool, _ := suite.testApp.AllocationPools.CreateNew(suite.anotherAllocationPoolRequest)
		newEnvironment, _ := suite.testApp.Environments.CreateNew(suite.goodEnvironmentRequest)

		// add the environment to the cluster
		_, err := suite.testApp.AllocationPools.AddEnvironmentByID(newAllocationPool, newEnvironment.ID)
		assert.NoError(suite.T(), err)
		updatedEnvironment, _ := suite.testApp.Environments.GetByName(newEnvironment.Name)
		assert.Equal(suite.T(), newAllocationPool.ID, *updatedEnvironment.AllocationPoolID)

		// change to a new cluster
		updatedAllocationPool, _ := suite.testApp.AllocationPools.AddEnvironmentByID(anotherAllocationPool, newEnvironment.ID)
		updatedEnvironment, _ = suite.testApp.Environments.GetByName(newEnvironment.Name)
		assert.Equal(suite.T(), anotherAllocationPool.ID, *updatedEnvironment.AllocationPoolID)
		assert.Equal(suite.T(), updatedEnvironment.ID, (updatedAllocationPool.Environments)[0].ID)
	})
}

func (suite *AllocationPoolTestSuite) TestIntegrationAllocationPoolGetByName() {
	suite.Run("GetByName gets an allocationPool by name", func() {
		testutils.Cleanup(suite.T(), suite.testApp.db)

		_, err := suite.testApp.AllocationPools.CreateNew(suite.goodAllocationPoolRequest)
		assert.NoError(suite.T(), err)

		foundAllocationPool, err := suite.testApp.AllocationPools.GetByName(suite.goodAllocationPoolRequest.Name)

		assert.Equal(suite.T(), foundAllocationPool.Name, suite.goodAllocationPoolRequest.Name)

		assert.NoError(suite.T(), err)
	})

	suite.Run("GetByName returns error if not found", func() {
		testutils.Cleanup(suite.T(), suite.testApp.db)

		_, err := suite.testApp.AllocationPools.CreateNew(suite.goodAllocationPoolRequest)
		assert.NoError(suite.T(), err)

		foundAllocationPool, err := suite.testApp.AllocationPools.GetByName("this-doesnt-exist")

		assert.Equal(suite.T(), foundAllocationPool.Name, "")
		assert.Equal(suite.T(), errors.New("record not found"), err)
	})
}

func (suite *AllocationPoolTestSuite) TestIntegrationAllocationPoolGetByID() {
	suite.Run("GetByID gets an allocationPool by ID", func() {
		testutils.Cleanup(suite.T(), suite.testApp.db)

		newAllocationPool, err := suite.testApp.AllocationPools.CreateNew(suite.goodAllocationPoolRequest)
		assert.NoError(suite.T(), err)

		foundAllocationPool, err := suite.testApp.AllocationPools.GetByID(newAllocationPool.ID)

		assert.Equal(suite.T(), foundAllocationPool.ID, newAllocationPool.ID)

		assert.NoError(suite.T(), err)
	})

	suite.Run("GetByName returns error if not found", func() {
		testutils.Cleanup(suite.T(), suite.testApp.db)

		_, err := suite.testApp.AllocationPools.CreateNew(suite.goodAllocationPoolRequest)
		assert.NoError(suite.T(), err)

		foundAllocationPool, err := suite.testApp.AllocationPools.GetByID(suite.notFoundID)

		assert.Equal(suite.T(), foundAllocationPool.ID, 0)
		assert.Equal(suite.T(), errors.New("record not found"), err)
	})
}

func (suite *AllocationPoolTestSuite) TestIntegrationAllocationPoolListAll() {
	suite.Run("ListAll returns nothing", func() {
		testutils.Cleanup(suite.T(), suite.testApp.db)

		foundAllocationPools, err := suite.testApp.AllocationPools.ListAll()

		assert.Equal(suite.T(), len(foundAllocationPools), 0)
		assert.NoError(suite.T(), err)
	})

	suite.Run("ListAll returns one AllocationPool", func() {
		testutils.Cleanup(suite.T(), suite.testApp.db)

		_, err := suite.testApp.AllocationPools.CreateNew(suite.goodAllocationPoolRequest)
		assert.NoError(suite.T(), err)

		foundAllocationPools, err := suite.testApp.AllocationPools.ListAll()

		assert.Equal(suite.T(), len(foundAllocationPools), 1)
		assert.Equal(suite.T(), foundAllocationPools[0].Name, suite.goodAllocationPoolRequest.Name)
		assert.NoError(suite.T(), err)
	})

	suite.Run("ListAll returns many AllocationPools", func() {
		testutils.Cleanup(suite.T(), suite.testApp.db)

		_, err := suite.testApp.AllocationPools.CreateNew(suite.goodAllocationPoolRequest)
		assert.NoError(suite.T(), err)
		_, err = suite.testApp.AllocationPools.CreateNew(suite.anotherAllocationPoolRequest)
		assert.NoError(suite.T(), err)

		foundAllocationPools, err := suite.testApp.AllocationPools.ListAll()

		assert.Equal(suite.T(), len(foundAllocationPools), 2)
		assert.NoError(suite.T(), err)
	})
}

func (suite *AllocationPoolTestSuite) TestIntegrationAllocationPoolDoesAllocationPoolExist() {
	suite.Run("AllocationPoolDoesExist returns true when exists", func() {
		testutils.Cleanup(suite.T(), suite.testApp.db)

		newAllocationPool, _ := suite.testApp.AllocationPools.CreateNew(suite.goodAllocationPoolRequest)

		allocationPoolID, doesAllocationPoolExist := suite.testApp.AllocationPools.DoesAllocationPoolExist(suite.goodAllocationPoolRequest.Name)

		assert.Equal(suite.T(), allocationPoolID, newAllocationPool.ID)
		assert.Equal(suite.T(), doesAllocationPoolExist, true)
	})

	suite.Run("AllocationPoolDoesExist returns false when not exists", func() {
		testutils.Cleanup(suite.T(), suite.testApp.db)

		_, err := suite.testApp.AllocationPools.CreateNew(suite.goodAllocationPoolRequest)
		assert.NoError(suite.T(), err)

		allocationPoolID, doesAllocationPoolExist := suite.testApp.AllocationPools.DoesAllocationPoolExist("no-allocationPool-here")

		assert.Equal(suite.T(), allocationPoolID, 0)
		assert.Equal(suite.T(), doesAllocationPoolExist, false)
	})
}
