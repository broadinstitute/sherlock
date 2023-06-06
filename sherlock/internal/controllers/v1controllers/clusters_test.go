// Tests for the ClusterController

package v1controllers

import (
	"errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/broadinstitute/sherlock/sherlock/internal/db"
	"github.com/broadinstitute/sherlock/sherlock/internal/models/v1models"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context
type ClusterTestSuite struct {
	suite.Suite
	testApp               *TestApplication
	goodClusterRequest    v1models.CreateClusterRequest
	anotherClusterRequest v1models.CreateClusterRequest
	badClusterRequest     v1models.CreateClusterRequest
	notFoundID            int
}

// Test entry point
func TestFunctionalClustersSuite(t *testing.T) {
	// skip functional tests if go test is invoked with -short flag
	if testing.Short() {
		t.Skip("skipping functional test")
	}
	suite.Run(t, new(ClusterTestSuite))
}

// between-test initialization
func (suite *ClusterTestSuite) SetupTest() {
	suite.testApp = initTestClusterApp(suite.T())
	suite.goodClusterRequest = v1models.CreateClusterRequest{
		Name: "terra-prod",
	}
	suite.anotherClusterRequest = v1models.CreateClusterRequest{
		Name: "terra-microprod",
	}
	suite.badClusterRequest = v1models.CreateClusterRequest{}
	suite.notFoundID = 1234567890 //unsure of a way to guarantee not-found-ness
}

func (suite *ClusterTestSuite) TearDownTest() {
	// each test runs in its own isolated transaction
	// this ensures we cleanup after each test as it completes
	suite.testApp.DB.Rollback()
}

//
// Test ClusterController Setup
//

// connect to DB and create the Application
func initTestClusterApp(t *testing.T) *TestApplication {
	config.LoadTestConfig(t)
	dbConn := db.ConnectAndConfigureFromTest(t)

	// ensures each test will run in it's own isolated transaction
	// The transaction will be rolled back after each test
	// regardless of pass or fail
	dbConn = dbConn.Begin()
	app := &TestApplication{
		Clusters: NewClusterController(dbConn),
		DB:       dbConn,
	}

	return app
}

//
// The Actual Tests
//

func (suite *ClusterTestSuite) TestFunctionalCreateClusters() {
	suite.Run("creates a valid cluster", func() {
		db.Truncate(suite.T(), suite.testApp.DB)

		newCluster, err := suite.testApp.Clusters.CreateNew(suite.goodClusterRequest)
		assert.NoError(suite.T(), err)

		assert.Equal(suite.T(), suite.goodClusterRequest.Name, newCluster.Name)
		assert.NoError(suite.T(), err)
	})

	suite.Run("fails to create a cluster with no name", func() {
		db.Truncate(suite.T(), suite.testApp.DB)

		expectedError := errors.New("error saving to database: ERROR: null value in column \"name\" of relation \"clusters\" violates not-null constraint (SQLSTATE 23502)")
		namelessClusterRequest := suite.goodClusterRequest
		namelessClusterRequest.Name = ""

		newCluster, err := suite.testApp.Clusters.CreateNew(namelessClusterRequest)

		assert.Equal(suite.T(), "", newCluster.Name)
		assert.Equal(suite.T(), expectedError, err)
	})

	suite.Run("fails to create a cluster with duplicate name", func() {
		db.Truncate(suite.T(), suite.testApp.DB)

		expectedError := errors.New("error saving to database: ERROR: duplicate key value violates unique constraint \"clusters_name_key\" (SQLSTATE 23505)")

		_, err := suite.testApp.Clusters.CreateNew(suite.goodClusterRequest)
		assert.NoError(suite.T(), err)
		newCluster, err := suite.testApp.Clusters.CreateNew(suite.goodClusterRequest)

		assert.Equal(suite.T(), "", newCluster.Name)
		assert.Equal(suite.T(), expectedError, err)
	})
}

func (suite *ClusterTestSuite) TestFunctionalClusterGetByName() {
	suite.Run("GetByName gets an cluster by name", func() {
		db.Truncate(suite.T(), suite.testApp.DB)

		_, err := suite.testApp.Clusters.CreateNew(suite.goodClusterRequest)
		assert.NoError(suite.T(), err)

		foundCluster, err := suite.testApp.Clusters.GetByName(suite.goodClusterRequest.Name)

		assert.Equal(suite.T(), suite.goodClusterRequest.Name, foundCluster.Name)

		assert.NoError(suite.T(), err)
	})

	suite.Run("GetByName returns error if not found", func() {
		db.Truncate(suite.T(), suite.testApp.DB)

		_, err := suite.testApp.Clusters.CreateNew(suite.goodClusterRequest)
		assert.NoError(suite.T(), err)

		foundCluster, err := suite.testApp.Clusters.GetByName("this-doesnt-exist")

		assert.Equal(suite.T(), foundCluster.Name, "")
		assert.Equal(suite.T(), errors.New("record not found"), err)
	})
}

func (suite *ClusterTestSuite) TestFunctionalClusterGetByID() {
	suite.Run("GetByID gets an allocationPool by ID", func() {
		db.Truncate(suite.T(), suite.testApp.DB)

		newCluster, err := suite.testApp.Clusters.CreateNew(suite.goodClusterRequest)
		assert.NoError(suite.T(), err)

		foundCluster, err := suite.testApp.Clusters.GetByID(newCluster.ID)

		assert.Equal(suite.T(), foundCluster.ID, newCluster.ID)

		assert.NoError(suite.T(), err)
	})

	suite.Run("GetByName returns error if not found", func() {
		db.Truncate(suite.T(), suite.testApp.DB)

		_, err := suite.testApp.Clusters.CreateNew(suite.goodClusterRequest)
		assert.NoError(suite.T(), err)

		foundCluster, err := suite.testApp.Clusters.GetByID(suite.notFoundID)

		assert.Equal(suite.T(), foundCluster.ID, 0)
		assert.Equal(suite.T(), errors.New("record not found"), err)
	})
}

func (suite *ClusterTestSuite) TestFunctionalClusterListAll() {
	suite.Run("ListAll returns nothing", func() {
		db.Truncate(suite.T(), suite.testApp.DB)

		foundClusters, err := suite.testApp.Clusters.ListAll()

		assert.Equal(suite.T(), len(foundClusters), 0)
		assert.NoError(suite.T(), err)
	})

	suite.Run("ListAll returns one Cluster", func() {
		db.Truncate(suite.T(), suite.testApp.DB)

		_, err := suite.testApp.Clusters.CreateNew(suite.goodClusterRequest)
		assert.NoError(suite.T(), err)

		foundClusters, err := suite.testApp.Clusters.ListAll()

		assert.Equal(suite.T(), len(foundClusters), 1)
		assert.Equal(suite.T(), foundClusters[0].Name, suite.goodClusterRequest.Name)
		assert.NoError(suite.T(), err)
	})

	suite.Run("ListAll returns many Clusters", func() {
		db.Truncate(suite.T(), suite.testApp.DB)

		_, err := suite.testApp.Clusters.CreateNew(suite.goodClusterRequest)
		assert.NoError(suite.T(), err)
		_, err = suite.testApp.Clusters.CreateNew(suite.anotherClusterRequest)
		assert.NoError(suite.T(), err)

		foundClusters, err := suite.testApp.Clusters.ListAll()

		assert.Equal(suite.T(), len(foundClusters), 2)
		assert.NoError(suite.T(), err)
	})
}

func (suite *ClusterTestSuite) TestFunctionalClusterDoesClusterExist() {
	suite.Run("ClusterDoesExist returns true when exists", func() {
		db.Truncate(suite.T(), suite.testApp.DB)

		newCluster, _ := suite.testApp.Clusters.CreateNew(suite.goodClusterRequest)

		clusterID, doesClusterExist := suite.testApp.Clusters.DoesClusterExist(suite.goodClusterRequest.Name)

		assert.Equal(suite.T(), clusterID, newCluster.ID)
		assert.Equal(suite.T(), doesClusterExist, true)
	})

	suite.Run("ClusterDoesExist returns false when not exists", func() {
		db.Truncate(suite.T(), suite.testApp.DB)

		_, err := suite.testApp.Clusters.CreateNew(suite.goodClusterRequest)
		assert.NoError(suite.T(), err)

		clusterID, doesClusterExist := suite.testApp.Clusters.DoesClusterExist("no-cluster-here")

		assert.Equal(suite.T(), clusterID, 0)
		assert.Equal(suite.T(), doesClusterExist, false)
	})
}
