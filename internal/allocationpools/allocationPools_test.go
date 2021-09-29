// Tests for the AllocationPoolController

package allocationpools

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
type AllocationPoolTestSuite struct {
	suite.Suite
	testApp                      *TestApplication
	goodAllocationPoolRequest    CreateAllocationPoolRequest
	anotherAllocationPoolRequest CreateAllocationPoolRequest
	badAllocationPoolRequest     CreateAllocationPoolRequest
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
	suite.testApp = initTestApp(suite.T())
	suite.goodAllocationPoolRequest = CreateAllocationPoolRequest{
		Name: "terra-prod",
	}
	suite.anotherAllocationPoolRequest = CreateAllocationPoolRequest{
		Name: "terra-microprod",
	}
	suite.badAllocationPoolRequest = CreateAllocationPoolRequest{}
}

//
// Test AllocationPool Setup
//

// only load the Controller we care about
type TestApplication struct {
	AllocationPools *AllocationPoolController
	db              *gorm.DB
}

// connect to DB and create the Application
func initTestApp(t *testing.T) *TestApplication {
	dbConn := testutils.ConnectAndMigrate(t)
	app := &TestApplication{
		AllocationPools: NewController(dbConn),
		db:              dbConn,
	}

	testutils.Cleanup(t, app.db)

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

		assert.Equal(suite.T(), suite.goodAllocationPoolRequest.Name, newAllocationPool.Name)
		assert.Equal(suite.T(), expectedError, err)
	})

	suite.Run("fails to create a allocationPool with duplicate name", func() {
		testutils.Cleanup(suite.T(), suite.testApp.db)

		_, err := suite.testApp.AllocationPools.CreateNew(suite.goodAllocationPoolRequest)
		assert.NoError(suite.T(), err)
		_, err = suite.testApp.AllocationPools.CreateNew(suite.anotherAllocationPoolRequest)
		assert.NoError(suite.T(), err)
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
