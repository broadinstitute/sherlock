package v2controllers

import (
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/testutils"
	"github.com/broadinstitute/sherlock/sherlock/internal/deprecated_db"
	"github.com/broadinstitute/sherlock/sherlock/internal/deprecated_models/v2models"
	"testing"

	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

//
// Test suite configuration
//

func TestClusterControllerSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping functional test")
	}
	suite.Run(t, new(clusterControllerSuite))
}

type clusterControllerSuite struct {
	suite.Suite
	*ControllerSet
	db *gorm.DB
}

func (suite *clusterControllerSuite) SetupTest() {
	config.LoadTestConfig()
	suite.db = deprecated_db.ConnectAndConfigureFromTest(suite.T())
	suite.db.Begin()
	suite.ControllerSet = NewControllerSet(v2models.NewStoreSet(suite.db))
}

func (suite *clusterControllerSuite) TearDownTest() {
	suite.db.Rollback()
}

//
// Controller seeding
//

var (
	terraDevCluster = CreatableCluster{
		Name:          "terra-dev",
		GoogleProject: "broad-dsde-dev",
		Location:      "us-central1-a",
		EditableCluster: EditableCluster{
			Base:    testutils.PointerTo("terra"),
			Address: testutils.PointerTo("192.168.0.1"),
		},
	}
	datarepoDevCluster = CreatableCluster{
		Name:          "datarepo-dev",
		GoogleProject: "datarepo-dev",
		Location:      "us-central1-a",
		EditableCluster: EditableCluster{
			Base:    testutils.PointerTo("datarepo"),
			Address: testutils.PointerTo("192.168.30.1"),
		},
	}
	terraStagingCluster = CreatableCluster{
		Name:          "terra-staging",
		GoogleProject: "broad-dsde-staging",
		Location:      "us-central1-a",
		EditableCluster: EditableCluster{
			Base:                testutils.PointerTo("terra"),
			Address:             testutils.PointerTo("192.168.0.10"),
			RequiresSuitability: testutils.PointerTo(true),
		},
	}
	terraProdCluster = CreatableCluster{
		Name:          "terra-prod",
		GoogleProject: "broad-dsde-prod",
		Location:      "us-central1",
		EditableCluster: EditableCluster{
			Base:                testutils.PointerTo("terra"),
			Address:             testutils.PointerTo("192.168.0.10"),
			RequiresSuitability: testutils.PointerTo(true),
		},
	}
	datarepoProdCluster = CreatableCluster{
		Name:          "datarepo-prod",
		GoogleProject: "datarepo-prod",
		Location:      "us-central1",
		EditableCluster: EditableCluster{
			Base:                testutils.PointerTo("datarepo"),
			Address:             testutils.PointerTo("192.168.30.10"),
			RequiresSuitability: testutils.PointerTo(true),
		},
	}
	terraDevBeesCluster = CreatableCluster{
		Name:          "terra-dev-bees",
		GoogleProject: "broad-dsde-dev",
		Location:      "us-central1-a",
		EditableCluster: EditableCluster{
			Base:    testutils.PointerTo("bee-cluster"),
			Address: testutils.PointerTo("192.168.1.1"),
		},
	}
	terraQaBeesCluster = CreatableCluster{
		Name:          "terra-qa-bees",
		GoogleProject: "broad-dsde-qa",
		Location:      "us-central1-a",
		EditableCluster: EditableCluster{
			Base:    testutils.PointerTo("bee-cluster"),
			Address: testutils.PointerTo("192.168.1.2"),
		},
	}
	terraDevAzureCluster = CreatableCluster{
		Name:              "terra-dev-azure",
		Provider:          "azure",
		AzureSubscription: "some-uuid",
		Location:          "US-EAST",
		EditableCluster: EditableCluster{
			Base:    testutils.PointerTo("terra-azure"),
			Address: testutils.PointerTo("192.168.2.1"),
		},
	}
	clusterSeedList = []CreatableCluster{
		terraDevCluster,
		terraStagingCluster,
		terraProdCluster,
		datarepoDevCluster,
		datarepoProdCluster,
		terraDevBeesCluster,
		terraQaBeesCluster,
	}
)

func (controllerSet *ControllerSet) seedClusters(t *testing.T, db *gorm.DB) {
	for _, creatable := range clusterSeedList {
		if _, _, err := controllerSet.ClusterController.Create(creatable, generateUser(t, db, true)); err != nil {
			t.Errorf("error seeding cluster %s: %v", creatable.Name, err)
		}
	}
}

//
// Controller tests
//

func (suite *clusterControllerSuite) TestClusterCreate() {
	suite.Run("can create a new cluster", func() {
		deprecated_db.Truncate(suite.T(), suite.db)

		cluster, created, err := suite.ClusterController.Create(terraDevCluster, generateUser(suite.T(), suite.db, false))
		assert.NoError(suite.T(), err)
		assert.True(suite.T(), created)
		assert.Equal(suite.T(), terraDevCluster.Name, cluster.Name)
		assert.True(suite.T(), cluster.ID > 0)
		suite.Run("default provider google", func() {
			assert.Equal(suite.T(), "google", cluster.Provider)
		})
		suite.Run("default non-suitable", func() {
			assert.False(suite.T(), *cluster.RequiresSuitability)
		})
		suite.Run("default terra-helmfile ref", func() {
			suite.Assert().Equal("HEAD", *cluster.HelmfileRef)
		})
	})
	suite.Run("can create a new azure cluster", func() {
		deprecated_db.Truncate(suite.T(), suite.db)

		cluster, created, err := suite.ClusterController.Create(terraDevAzureCluster, generateUser(suite.T(), suite.db, false))
		assert.NoError(suite.T(), err)
		assert.True(suite.T(), created)
		assert.Equal(suite.T(), terraDevAzureCluster.Name, cluster.Name)
		assert.True(suite.T(), cluster.ID > 0)
		assert.Equal(suite.T(), terraDevAzureCluster.AzureSubscription, cluster.AzureSubscription)
		assert.Equal(suite.T(), terraDevAzureCluster.Provider, cluster.Provider)
	})
	suite.Run("won't create duplicates", func() {
		deprecated_db.Truncate(suite.T(), suite.db)

		cluster, created, err := suite.ClusterController.Create(terraDevCluster, generateUser(suite.T(), suite.db, false))
		assert.NoError(suite.T(), err)
		assert.True(suite.T(), created)
		assert.True(suite.T(), cluster.ID > 0)
		_, created, err = suite.ClusterController.Create(terraDevCluster, generateUser(suite.T(), suite.db, false))
		assert.ErrorContains(suite.T(), err, errors.Conflict)
		assert.False(suite.T(), created)
	})
	suite.Run("validates incoming entries", func() {
		deprecated_db.Truncate(suite.T(), suite.db)

		_, created, err := suite.ClusterController.Create(CreatableCluster{}, generateUser(suite.T(), suite.db, false))
		assert.ErrorContains(suite.T(), err, errors.BadRequest)
		assert.False(suite.T(), created)
	})
	suite.Run("checks suitability", func() {
		deprecated_db.Truncate(suite.T(), suite.db)

		suite.Run("blocks suitable creation for non-suitable", func() {
			_, created, err := suite.ClusterController.Create(terraProdCluster, generateUser(suite.T(), suite.db, false))
			assert.ErrorContains(suite.T(), err, errors.Forbidden)
			assert.False(suite.T(), created)
		})
		suite.Run("allows suitable creation for suitable", func() {
			cluster, created, err := suite.ClusterController.Create(terraProdCluster, generateUser(suite.T(), suite.db, true))
			assert.NoError(suite.T(), err)
			assert.True(suite.T(), created)
			assert.True(suite.T(), cluster.ID > 0)
		})
	})
}

func (suite *clusterControllerSuite) TestClusterListAllMatching() {
	deprecated_db.Truncate(suite.T(), suite.db)
	suite.seedClusters(suite.T(), suite.db)

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
		matching, err := suite.ClusterController.ListAllMatching(Cluster{CreatableCluster: terraDevCluster}, 0)
		assert.NoError(suite.T(), err)
		assert.Equal(suite.T(), 1, len(matching))
		assert.Equal(suite.T(), terraDevCluster.Name, matching[0].Name)
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
	suite.Run("none is an empty list, not null", func() {
		matching, err := suite.ClusterController.ListAllMatching(
			Cluster{CreatableCluster: CreatableCluster{Name: "blah"}}, 0)
		assert.NoError(suite.T(), err)
		assert.NotNil(suite.T(), matching)
		assert.Empty(suite.T(), matching)
	})
}

func (suite *clusterControllerSuite) TestClusterGet() {
	deprecated_db.Truncate(suite.T(), suite.db)
	suite.seedClusters(suite.T(), suite.db)

	suite.Run("successfully", func() {
		byName, err := suite.ClusterController.Get(terraDevCluster.Name)
		assert.NoError(suite.T(), err)
		assert.True(suite.T(), byName.ID > 0)
		byID, err := suite.ClusterController.Get(fmt.Sprintf("%d", byName.ID))
		assert.NoError(suite.T(), err)
		assert.Equal(suite.T(), terraDevCluster.Name, byID.Name)
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

func (suite *clusterControllerSuite) TestClusterGetOtherValidSelectors() {
	deprecated_db.Truncate(suite.T(), suite.db)
	suite.seedClusters(suite.T(), suite.db)

	suite.Run("successfully", func() {
		selectors, err := suite.ClusterController.GetOtherValidSelectors(terraDevCluster.Name)
		assert.NoError(suite.T(), err)
		assert.Equal(suite.T(), 2, len(selectors))
		assert.Equal(suite.T(), terraDevCluster.Name, selectors[0])
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

func (suite *clusterControllerSuite) TestClusterEdit() {
	suite.Run("successfully", func() {
		deprecated_db.Truncate(suite.T(), suite.db)
		suite.seedClusters(suite.T(), suite.db)

		before, err := suite.ClusterController.Get(terraDevCluster.Name)
		assert.NoError(suite.T(), err)
		assert.Equal(suite.T(), terraDevCluster.Base, before.Base)
		newBase := testutils.PointerTo("new")
		edited, err := suite.ClusterController.Edit(terraDevCluster.Name, EditableCluster{Base: newBase}, generateUser(suite.T(), suite.db, false))
		assert.NoError(suite.T(), err)
		assert.Equal(suite.T(), newBase, edited.Base)
		after, err := suite.ClusterController.Get(terraDevCluster.Name)
		assert.NoError(suite.T(), err)
		assert.Equal(suite.T(), newBase, after.Base)
	})
	suite.Run("edit to suitable cluster", func() {
		deprecated_db.Truncate(suite.T(), suite.db)
		suite.seedClusters(suite.T(), suite.db)
		newBase := testutils.PointerTo("new")

		suite.Run("unsuccessfully if not suitable", func() {
			_, err := suite.ClusterController.Edit(terraProdCluster.Name, EditableCluster{Base: newBase}, generateUser(suite.T(), suite.db, false))
			assert.ErrorContains(suite.T(), err, errors.Forbidden)
			notEdited, err := suite.ClusterController.Get(terraProdCluster.Name)
			assert.NoError(suite.T(), err)
			assert.Equal(suite.T(), terraProdCluster.Base, notEdited.Base)
		})
		suite.Run("successfully if suitable", func() {
			edited, err := suite.ClusterController.Edit(terraProdCluster.Name, EditableCluster{Base: newBase}, generateUser(suite.T(), suite.db, true))
			assert.NoError(suite.T(), err)
			assert.Equal(suite.T(), newBase, edited.Base)
		})
	})
	suite.Run("edit that would make cluster suitable", func() {
		deprecated_db.Truncate(suite.T(), suite.db)
		suite.seedClusters(suite.T(), suite.db)

		suite.Run("unsuccessfully if not suitable", func() {
			_, err := suite.ClusterController.Edit(terraDevCluster.Name, EditableCluster{RequiresSuitability: testutils.PointerTo(true)}, generateUser(suite.T(), suite.db, false))
			assert.ErrorContains(suite.T(), err, errors.Forbidden)
			notEdited, err := suite.ClusterController.Get(terraDevCluster.Name)
			assert.NoError(suite.T(), err)
			assert.False(suite.T(), *notEdited.RequiresSuitability)
		})
		suite.Run("successfully if suitable", func() {
			edited, err := suite.ClusterController.Edit(terraDevCluster.Name, EditableCluster{RequiresSuitability: testutils.PointerTo(true)}, generateUser(suite.T(), suite.db, true))
			assert.NoError(suite.T(), err)
			assert.True(suite.T(), *edited.RequiresSuitability)
		})
	})
	suite.Run("unsuccessfully if invalid", func() {
		deprecated_db.Truncate(suite.T(), suite.db)
		suite.seedClusters(suite.T(), suite.db)

		_, err := suite.ClusterController.Edit(terraDevCluster.Name, EditableCluster{Base: testutils.PointerTo("")}, generateUser(suite.T(), suite.db, false))
		assert.ErrorContains(suite.T(), err, errors.BadRequest)
	})
}

func (suite *clusterControllerSuite) TestClusterDelete() {
	suite.Run("successfully", func() {
		deprecated_db.Truncate(suite.T(), suite.db)
		suite.seedClusters(suite.T(), suite.db)

		deleted, err := suite.ClusterController.Delete(terraDevCluster.Name, generateUser(suite.T(), suite.db, false))
		assert.NoError(suite.T(), err)
		assert.Equal(suite.T(), terraDevCluster.Name, deleted.Name)
		_, err = suite.ClusterController.Get(terraDevCluster.Name)
		assert.ErrorContains(suite.T(), err, errors.NotFound)
		suite.Run("sql constraints ignore soft deletion", func() {
			_, created, err := suite.ClusterController.Create(terraDevCluster, generateUser(suite.T(), suite.db, false))
			assert.ErrorContains(suite.T(), err, errors.BadRequest)
			assert.ErrorContains(suite.T(), err, "Contact DevOps")
			assert.False(suite.T(), created)
		})
	})
	suite.Run("delete suitable environment", func() {
		deprecated_db.Truncate(suite.T(), suite.db)
		suite.seedClusters(suite.T(), suite.db)

		suite.Run("unsuccessfully if not suitable", func() {
			_, err := suite.ClusterController.Delete(terraProdCluster.Name, generateUser(suite.T(), suite.db, false))
			assert.ErrorContains(suite.T(), err, errors.Forbidden)
		})
		suite.Run("successfully if suitable", func() {
			deleted, err := suite.ClusterController.Delete(terraProdCluster.Name, generateUser(suite.T(), suite.db, true))
			assert.NoError(suite.T(), err)
			assert.Equal(suite.T(), terraProdCluster.Name, deleted.Name)
		})
	})
}
