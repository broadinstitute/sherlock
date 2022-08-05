package v2controllers

import (
	"fmt"
	"github.com/broadinstitute/sherlock/internal/auth"
	"github.com/broadinstitute/sherlock/internal/config"
	"github.com/broadinstitute/sherlock/internal/db"
	"github.com/broadinstitute/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/internal/models/v2models"
	"github.com/broadinstitute/sherlock/internal/testutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
	"testing"
)

//
// Test suite configuration
//

func TestClusterControllerSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping functional test")
	}
	suite.Run(t, new(ClusterControllerSuite))
}

type ClusterControllerSuite struct {
	suite.Suite
	*ControllerSet
	db *gorm.DB
}

func (suite *ClusterControllerSuite) SetupTest() {
	config.LoadTestConfig(suite.T())
	suite.db = db.ConnectFromTest(suite.T())
	suite.db.Begin()
	suite.ControllerSet = NewControllerSet(v2models.NewStoreSet(suite.db))
}

func (suite *ClusterControllerSuite) TearDownTest() {
	suite.db.Rollback()
}

//
// Controller seeding
//

var (
	terraDev = CreatableCluster{
		Name:          "terra-dev",
		GoogleProject: "broad-dsde-dev",
		EditableCluster: EditableCluster{
			Base:    testutils.PointerTo("terra"),
			Address: testutils.PointerTo("192.168.0.1"),
		},
	}
	terraProd = CreatableCluster{
		Name:          "terra-prod",
		GoogleProject: "broad-dsde-prod",
		EditableCluster: EditableCluster{
			Base:                testutils.PointerTo("terra"),
			Address:             testutils.PointerTo("192.168.0.10"),
			RequiresSuitability: testutils.PointerTo(true),
		},
	}
	terraDevBees = CreatableCluster{
		Name:          "terra-dev-bees",
		GoogleProject: "broad-dsde-dev",
		EditableCluster: EditableCluster{
			Base:    testutils.PointerTo("bee-cluster"),
			Address: testutils.PointerTo("192.168.1.1"),
		},
	}
	terraDevAzure = CreatableCluster{
		Name:              "terra-dev-azure",
		Provider:          "azure",
		AzureSubscription: "some-uuid",
		EditableCluster: EditableCluster{
			Base:    testutils.PointerTo("terra-azure"),
			Address: testutils.PointerTo("192.168.2.1"),
		},
	}
	clusterSeedList = []CreatableCluster{terraDev, terraProd, terraDevBees}
)

func (controllerSet *ControllerSet) seedClusters(t *testing.T) {
	for _, creatable := range clusterSeedList {
		if _, err := controllerSet.ClusterController.Create(creatable, auth.GenerateUser(t, true)); err != nil {
			t.Errorf("error seeding cluster %s: %v", creatable.Name, err)
		}
	}
}

//
// Controller tests
//

func (suite *ClusterControllerSuite) TestClusterCreate() {
	suite.Run("can create a new cluster", func() {
		db.Truncate(suite.T(), suite.db)

		cluster, err := suite.ClusterController.Create(terraDev, auth.GenerateUser(suite.T(), false))
		assert.NoError(suite.T(), err)
		assert.Equal(suite.T(), terraDev.Name, cluster.Name)
		assert.True(suite.T(), cluster.ID > 0)
		suite.Run("default provider google", func() {
			assert.Equal(suite.T(), "google", cluster.Provider)
		})
		suite.Run("default non-suitable", func() {
			assert.False(suite.T(), *cluster.RequiresSuitability)
		})
	})
	suite.Run("can create a new azure cluster", func() {
		db.Truncate(suite.T(), suite.db)

		cluster, err := suite.ClusterController.Create(terraDevAzure, auth.GenerateUser(suite.T(), false))
		assert.NoError(suite.T(), err)
		assert.Equal(suite.T(), terraDevAzure.Name, cluster.Name)
		assert.True(suite.T(), cluster.ID > 0)
		assert.Equal(suite.T(), terraDevAzure.AzureSubscription, cluster.AzureSubscription)
		assert.Equal(suite.T(), terraDevAzure.Provider, cluster.Provider)
	})
	suite.Run("won't create duplicates", func() {
		db.Truncate(suite.T(), suite.db)

		cluster, err := suite.ClusterController.Create(terraDev, auth.GenerateUser(suite.T(), false))
		assert.NoError(suite.T(), err)
		assert.True(suite.T(), cluster.ID > 0)
		_, err = suite.ClusterController.Create(terraDev, auth.GenerateUser(suite.T(), false))
		assert.ErrorContains(suite.T(), err, errors.Conflict)
	})
	suite.Run("validates incoming entries", func() {
		db.Truncate(suite.T(), suite.db)

		_, err := suite.ClusterController.Create(CreatableCluster{}, auth.GenerateUser(suite.T(), false))
		assert.ErrorContains(suite.T(), err, errors.BadRequest)
	})
	suite.Run("checks suitability", func() {
		db.Truncate(suite.T(), suite.db)

		suite.Run("blocks suitable creation for non-suitable", func() {
			_, err := suite.ClusterController.Create(terraProd, auth.GenerateUser(suite.T(), false))
			assert.ErrorContains(suite.T(), err, errors.Forbidden)
		})
		suite.Run("allows suitable creation for suitable", func() {
			cluster, err := suite.ClusterController.Create(terraProd, auth.GenerateUser(suite.T(), true))
			assert.NoError(suite.T(), err)
			assert.True(suite.T(), cluster.ID > 0)
		})
	})
}

func (suite *ClusterControllerSuite) TestClusterListAllMatching() {
	db.Truncate(suite.T(), suite.db)
	suite.ControllerSet.seedClusters(suite.T())

	suite.Run("lists all clusters", func() {
		matching, err := suite.ClusterController.ListAllMatching(Cluster{}, 0)
		assert.NoError(suite.T(), err)
		assert.Equal(suite.T(), len(clusterSeedList), len(matching))
		suite.Run("orders by latest updated", func() {
			latestUpdated := matching[0].UpdatedAt
			for _, cluster := range matching {
				assert.GreaterOrEqual(suite.T(), latestUpdated, cluster.UpdatedAt)
			}
		})
	})
	suite.Run("limits", func() {
		limit := 2
		matching, err := suite.ClusterController.ListAllMatching(Cluster{}, limit)
		assert.NoError(suite.T(), err)
		assert.Equal(suite.T(), limit, len(matching))
	})
	suite.Run("filters exactly", func() {
		matching, err := suite.ClusterController.ListAllMatching(Cluster{CreatableCluster: terraDev}, 0)
		assert.NoError(suite.T(), err)
		assert.Equal(suite.T(), 1, len(matching))
		assert.Equal(suite.T(), terraDev.Name, matching[0].Name)
	})
	suite.Run("filters multiple", func() {
		matching, err := suite.ClusterController.ListAllMatching(
			Cluster{CreatableCluster: CreatableCluster{EditableCluster: EditableCluster{Base: testutils.PointerTo("terra")}}}, 0)
		assert.NoError(suite.T(), err)
		assert.True(suite.T(), len(matching) > 1)
		for _, cluster := range matching {
			assert.Equal(suite.T(), testutils.PointerTo("terra"), cluster.Base)
		}
	})
}

func (suite *ClusterControllerSuite) TestClusterGet() {
	db.Truncate(suite.T(), suite.db)
	suite.ControllerSet.seedClusters(suite.T())

	suite.Run("successfully", func() {
		byName, err := suite.ClusterController.Get(terraDev.Name)
		assert.NoError(suite.T(), err)
		assert.True(suite.T(), byName.ID > 0)
		byID, err := suite.ClusterController.Get(fmt.Sprintf("%d", byName.ID))
		assert.NoError(suite.T(), err)
		assert.Equal(suite.T(), terraDev.Name, byID.Name)
	})
	suite.Run("unsuccessfully for non-present", func() {
		_, err := suite.ClusterController.Get("foobar")
		assert.ErrorContains(suite.T(), err, errors.NotFound)
	})
	suite.Run("unsuccessfully for invalid selector", func() {
		_, err := suite.ClusterController.Get("something obviously invalid")
		assert.ErrorContains(suite.T(), err, errors.BadRequest)
	})
}

func (suite *ClusterControllerSuite) TestClusterGetOtherValidSelectors() {
	db.Truncate(suite.T(), suite.db)
	suite.ControllerSet.seedClusters(suite.T())

	suite.Run("successfully", func() {
		selectors, err := suite.ClusterController.GetOtherValidSelectors(terraDev.Name)
		assert.NoError(suite.T(), err)
		assert.Equal(suite.T(), 2, len(selectors))
		assert.Equal(suite.T(), terraDev.Name, selectors[0])
	})
	suite.Run("unsuccessfully for not found", func() {
		_, err := suite.ClusterController.GetOtherValidSelectors("foobar")
		assert.ErrorContains(suite.T(), err, errors.NotFound)
	})
	suite.Run("unsuccessfully for invalid", func() {
		_, err := suite.ClusterController.GetOtherValidSelectors("something obviously invalid")
		assert.ErrorContains(suite.T(), err, errors.BadRequest)
	})
}

func (suite *ClusterControllerSuite) TestClusterEdit() {
	suite.Run("successfully", func() {
		db.Truncate(suite.T(), suite.db)
		suite.ControllerSet.seedClusters(suite.T())

		before, err := suite.ClusterController.Get(terraDev.Name)
		assert.NoError(suite.T(), err)
		assert.Equal(suite.T(), terraDev.Base, before.Base)
		newBase := testutils.PointerTo("new")
		edited, err := suite.ClusterController.Edit(terraDev.Name, EditableCluster{Base: newBase}, auth.GenerateUser(suite.T(), false))
		assert.NoError(suite.T(), err)
		assert.Equal(suite.T(), newBase, edited.Base)
		after, err := suite.ClusterController.Get(terraDev.Name)
		assert.NoError(suite.T(), err)
		assert.Equal(suite.T(), newBase, after.Base)
	})
	suite.Run("edit to suitable environment", func() {
		db.Truncate(suite.T(), suite.db)
		suite.ControllerSet.seedClusters(suite.T())
		newBase := testutils.PointerTo("new")

		suite.Run("unsuccessfully if not suitable", func() {
			_, err := suite.ClusterController.Edit(terraProd.Name, EditableCluster{Base: newBase}, auth.GenerateUser(suite.T(), false))
			assert.ErrorContains(suite.T(), err, errors.Forbidden)
			notEdited, err := suite.ClusterController.Get(terraProd.Name)
			assert.NoError(suite.T(), err)
			assert.Equal(suite.T(), terraProd.Base, notEdited.Base)
		})
		suite.Run("successfully if suitable", func() {
			edited, err := suite.ClusterController.Edit(terraProd.Name, EditableCluster{Base: newBase}, auth.GenerateUser(suite.T(), true))
			assert.NoError(suite.T(), err)
			assert.Equal(suite.T(), newBase, edited.Base)
		})
	})
	suite.Run("edit that would make environment suitable", func() {
		db.Truncate(suite.T(), suite.db)
		suite.ControllerSet.seedClusters(suite.T())

		suite.Run("unsuccessfully if not suitable", func() {
			_, err := suite.ClusterController.Edit(terraDev.Name, EditableCluster{RequiresSuitability: testutils.PointerTo(true)}, auth.GenerateUser(suite.T(), false))
			assert.ErrorContains(suite.T(), err, errors.Forbidden)
			notEdited, err := suite.ClusterController.Get(terraDev.Name)
			assert.NoError(suite.T(), err)
			assert.False(suite.T(), *notEdited.RequiresSuitability)
		})
		suite.Run("successfully if suitable", func() {
			edited, err := suite.ClusterController.Edit(terraDev.Name, EditableCluster{RequiresSuitability: testutils.PointerTo(true)}, auth.GenerateUser(suite.T(), true))
			assert.NoError(suite.T(), err)
			assert.True(suite.T(), *edited.RequiresSuitability)
		})
	})
	suite.Run("unsuccessfully if invalid", func() {
		db.Truncate(suite.T(), suite.db)
		suite.ControllerSet.seedClusters(suite.T())

		_, err := suite.ClusterController.Edit(terraDev.Name, EditableCluster{Base: testutils.PointerTo("")}, auth.GenerateUser(suite.T(), false))
		assert.ErrorContains(suite.T(), err, errors.BadRequest)
	})
}

func (suite *ClusterControllerSuite) TestClusterDelete() {
	suite.Run("successfully", func() {
		db.Truncate(suite.T(), suite.db)
		suite.ControllerSet.seedClusters(suite.T())

		deleted, err := suite.ClusterController.Delete(terraDev.Name, auth.GenerateUser(suite.T(), false))
		assert.NoError(suite.T(), err)
		assert.Equal(suite.T(), terraDev.Name, deleted.Name)
		_, err = suite.ClusterController.Get(terraDev.Name)
		assert.ErrorContains(suite.T(), err, errors.NotFound)
		suite.Run("sql constraints ignore soft deletion", func() {
			_, err = suite.ClusterController.Create(terraDev, auth.GenerateUser(suite.T(), false))
			assert.ErrorContains(suite.T(), err, errors.BadRequest)
			assert.ErrorContains(suite.T(), err, "Contact DevOps")
		})
	})
	suite.Run("delete suitable environment", func() {
		db.Truncate(suite.T(), suite.db)
		suite.ControllerSet.seedClusters(suite.T())

		suite.Run("unsuccessfully if not suitable", func() {
			_, err := suite.ClusterController.Delete(terraProd.Name, auth.GenerateUser(suite.T(), false))
			assert.ErrorContains(suite.T(), err, errors.Forbidden)
		})
		suite.Run("successfully if suitable", func() {
			deleted, err := suite.ClusterController.Delete(terraProd.Name, auth.GenerateUser(suite.T(), true))
			assert.NoError(suite.T(), err)
			assert.Equal(suite.T(), terraProd.Name, deleted.Name)
		})
	})
}
