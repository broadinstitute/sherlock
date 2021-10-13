// Tests for the ClusterController

package controllers

import (
	"errors"
	"testing"

	"github.com/broadinstitute/sherlock/internal/models"
	"github.com/broadinstitute/sherlock/internal/testutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context
type ClusterTestSuite struct {
	suite.Suite
	testApp               *TestApplication
	goodClusterRequest    models.CreateClusterRequest
	anotherClusterRequest models.CreateClusterRequest
	badClusterRequest     models.CreateClusterRequest
	notFoundID            int
}

// Test entry point
func TestIntegrationClustersSuite(t *testing.T) {
	// skip integration tests if go test is invoked with -short flag
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	suite.Run(t, new(ClusterTestSuite))
}

// between-test initialization
func (suite *ClusterTestSuite) SetupTest() {
	suite.testApp = initTestApp(suite.T())
	suite.goodClusterRequest = models.CreateClusterRequest{
		Name: "terra-prod",
	}
	suite.anotherClusterRequest = models.CreateClusterRequest{
		Name: "terra-microprod",
	}
	suite.badClusterRequest = models.CreateClusterRequest{}
	suite.notFoundID = 1234567890 //unsure of a way to guarantee not-found-ness
}

//
// Test Cluster Setup
//

// only load the Controller we care about
type TestApplication struct {
	Clusters *ClusterController
	db       *gorm.DB
}

// connect to DB and create the Application
func initTestApp(t *testing.T) *TestApplication {
	dbConn := testutils.ConnectAndMigrate(t)
	app := &TestApplication{
		Clusters: NewClusterController(dbConn),
		db:       dbConn,
	}

	testutils.Cleanup(t, app.db)

	return app
}

//
// The Actual Tests
//

func (suite *ClusterTestSuite) TestIntegrationCreateClusters() {
	suite.Run("creates a valid cluster", func() {
		testutils.Cleanup(suite.T(), suite.testApp.db)

		newCluster, err := suite.testApp.Clusters.CreateNew(suite.goodClusterRequest)
		assert.NoError(suite.T(), err)

		assert.Equal(suite.T(), suite.goodClusterRequest.Name, newCluster.Name)
		assert.NoError(suite.T(), err)
	})

	suite.Run("fails to create a cluster with no name", func() {
		testutils.Cleanup(suite.T(), suite.testApp.db)

		expectedError := errors.New("error saving to database: ERROR: null value in column \"name\" of relation \"clusters\" violates not-null constraint (SQLSTATE 23502)")
		namelessClusterRequest := suite.goodClusterRequest
		namelessClusterRequest.Name = ""

		newCluster, err := suite.testApp.Clusters.CreateNew(namelessClusterRequest)

		assert.Equal(suite.T(), "", newCluster.Name)
		assert.Equal(suite.T(), expectedError, err)
	})

	suite.Run("fails to create a cluster with duplicate name", func() {
		testutils.Cleanup(suite.T(), suite.testApp.db)

		expectedError := errors.New("error saving to database: ERROR: duplicate key value violates unique constraint \"clusters_name_key\" (SQLSTATE 23505)")

		_, err := suite.testApp.Clusters.CreateNew(suite.goodClusterRequest)
		assert.NoError(suite.T(), err)
		newCluster, err := suite.testApp.Clusters.CreateNew(suite.goodClusterRequest)

		assert.Equal(suite.T(), "", newCluster.Name)
		assert.Equal(suite.T(), expectedError, err)
	})
}

func (suite *ClusterTestSuite) TestIntegrationClusterGetByName() {
	suite.Run("GetByName gets an cluster by name", func() {
		testutils.Cleanup(suite.T(), suite.testApp.db)

		_, err := suite.testApp.Clusters.CreateNew(suite.goodClusterRequest)
		assert.NoError(suite.T(), err)

		foundCluster, err := suite.testApp.Clusters.GetByName(suite.goodClusterRequest.Name)

		assert.Equal(suite.T(), suite.goodClusterRequest.Name, foundCluster.Name)

		assert.NoError(suite.T(), err)
	})

	suite.Run("GetByName returns error if not found", func() {
		testutils.Cleanup(suite.T(), suite.testApp.db)

		_, err := suite.testApp.Clusters.CreateNew(suite.goodClusterRequest)
		assert.NoError(suite.T(), err)

		foundCluster, err := suite.testApp.Clusters.GetByName("this-doesnt-exist")

		assert.Equal(suite.T(), foundCluster.Name, "")
		assert.Equal(suite.T(), errors.New("record not found"), err)
	})
}

func (suite *ClusterTestSuite) TestIntegrationClusterGetByID() {
	suite.Run("GetByID gets an allocationPool by ID", func() {
		testutils.Cleanup(suite.T(), suite.testApp.db)

		newCluster, err := suite.testApp.Clusters.CreateNew(suite.goodClusterRequest)
		assert.NoError(suite.T(), err)

		foundCluster, err := suite.testApp.Clusters.GetByID(newCluster.ID)

		assert.Equal(suite.T(), foundCluster.ID, newCluster.ID)

		assert.NoError(suite.T(), err)
	})

	suite.Run("GetByName returns error if not found", func() {
		testutils.Cleanup(suite.T(), suite.testApp.db)

		_, err := suite.testApp.Clusters.CreateNew(suite.goodClusterRequest)
		assert.NoError(suite.T(), err)

		foundCluster, err := suite.testApp.Clusters.GetByID(suite.notFoundID)

		assert.Equal(suite.T(), foundCluster.ID, 0)
		assert.Equal(suite.T(), errors.New("record not found"), err)
	})
}

func (suite *ClusterTestSuite) TestIntegrationClusterListAll() {
	suite.Run("ListAll returns nothing", func() {
		testutils.Cleanup(suite.T(), suite.testApp.db)

		foundClusters, err := suite.testApp.Clusters.ListAll()

		assert.Equal(suite.T(), len(foundClusters), 0)
		assert.NoError(suite.T(), err)
	})

	suite.Run("ListAll returns one Cluster", func() {
		testutils.Cleanup(suite.T(), suite.testApp.db)

		_, err := suite.testApp.Clusters.CreateNew(suite.goodClusterRequest)
		assert.NoError(suite.T(), err)

		foundClusters, err := suite.testApp.Clusters.ListAll()

		assert.Equal(suite.T(), len(foundClusters), 1)
		assert.Equal(suite.T(), foundClusters[0].Name, suite.goodClusterRequest.Name)
		assert.NoError(suite.T(), err)
	})

	suite.Run("ListAll returns many Clusters", func() {
		testutils.Cleanup(suite.T(), suite.testApp.db)

		_, err := suite.testApp.Clusters.CreateNew(suite.goodClusterRequest)
		assert.NoError(suite.T(), err)
		_, err = suite.testApp.Clusters.CreateNew(suite.anotherClusterRequest)
		assert.NoError(suite.T(), err)

		foundClusters, err := suite.testApp.Clusters.ListAll()

		assert.Equal(suite.T(), len(foundClusters), 2)
		assert.NoError(suite.T(), err)
	})
}

func (suite *ClusterTestSuite) TestIntegrationClusterDoesClusterExist() {
	suite.Run("ClusterDoesExist returns true when exists", func() {
		testutils.Cleanup(suite.T(), suite.testApp.db)

		newCluster, _ := suite.testApp.Clusters.CreateNew(suite.goodClusterRequest)

		clusterID, doesClusterExist := suite.testApp.Clusters.DoesClusterExist(suite.goodClusterRequest.Name)

		assert.Equal(suite.T(), clusterID, newCluster.ID)
		assert.Equal(suite.T(), doesClusterExist, true)
	})

	suite.Run("ClusterDoesExist returns false when not exists", func() {
		testutils.Cleanup(suite.T(), suite.testApp.db)

		_, err := suite.testApp.Clusters.CreateNew(suite.goodClusterRequest)
		assert.NoError(suite.T(), err)

		clusterID, doesClusterExist := suite.testApp.Clusters.DoesClusterExist("no-cluster-here")

		assert.Equal(suite.T(), clusterID, 0)
		assert.Equal(suite.T(), doesClusterExist, false)
	})
}
