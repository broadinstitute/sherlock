package v2controllers

import (
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/broadinstitute/sherlock/sherlock/internal/deprecated_db"
	"github.com/broadinstitute/sherlock/sherlock/internal/deprecated_models/v2models"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
	"testing"
)

//
// Test suite configuration
//

func TestDatabaseInstanceControllerSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping functional test")
	}
	suite.Run(t, new(databaseInstanceControllerSuite))
}

type databaseInstanceControllerSuite struct {
	suite.Suite
	*ControllerSet
	db *gorm.DB
}

func (suite *databaseInstanceControllerSuite) SetupSuite() {
	config.LoadTestConfig()
	suite.db = deprecated_db.ConnectAndConfigureFromTest(suite.T())
	suite.ControllerSet = NewControllerSet(v2models.NewStoreSet(suite.db))
}

func (suite *databaseInstanceControllerSuite) TearDownSuite() {
	deprecated_db.Truncate(suite.T(), suite.db)
}

//
// Controller seeding
//

var (
	datarepoDevDatabaseInstance = CreatableDatabaseInstance{
		ChartRelease: datarepoDevChartRelease.Name,
		EditableDatabaseInstance: EditableDatabaseInstance{
			Platform:      utils.PointerTo("google"),
			GoogleProject: utils.PointerTo("broad-datarepo-dev"),
			InstanceName:  utils.PointerTo("datarepo-abcdef"),
		},
	}
	datarepoProdDatabaseInstance = CreatableDatabaseInstance{
		ChartRelease: datarepoProdChartRelease.Name,
		EditableDatabaseInstance: EditableDatabaseInstance{
			Platform:        utils.PointerTo("google"),
			GoogleProject:   utils.PointerTo("broad-datarepo-prod"),
			InstanceName:    utils.PointerTo("datarepo-abcdef"),
			DefaultDatabase: utils.PointerTo("datarepo-2"),
		},
	}
	datarepoSwatomationDatabaseInstance = CreatableDatabaseInstance{
		ChartRelease: fmt.Sprintf("%s-%s", datarepoChart.Name, swatomationEnvironment.Name),
		EditableDatabaseInstance: EditableDatabaseInstance{
			Platform: utils.PointerTo("kubernetes"),
		},
	}
	datarepoDynamicSwatomationDatabaseInstance = CreatableDatabaseInstance{
		ChartRelease: fmt.Sprintf("%s-%s", datarepoChart.Name, dynamicSwatomationEnvironment.Name),
		EditableDatabaseInstance: EditableDatabaseInstance{
			Platform: utils.PointerTo("kubernetes"),
		},
	}

	databaseInstanceSeedList = []CreatableDatabaseInstance{
		datarepoDevDatabaseInstance,
		datarepoProdDatabaseInstance,
		datarepoSwatomationDatabaseInstance,
		datarepoDynamicSwatomationDatabaseInstance,
	}
)

func (controllerSet *ControllerSet) seedDatabaseInstances(t *testing.T, db *gorm.DB) {
	for _, creatable := range databaseInstanceSeedList {
		if _, _, err := controllerSet.DatabaseInstanceController.Create(creatable, generateUser(t, db, true)); err != nil {
			t.Errorf("error seeding database instance for %s: %w", creatable.ChartRelease, err)
		}
	}
}

//
// Controller tests
//

func (suite *databaseInstanceControllerSuite) TestDatabaseInstanceCreate() {
	suite.Run("can create a new database instance", func() {
		deprecated_db.Truncate(suite.T(), suite.db)
		suite.seedClusters(suite.T(), suite.db)
		suite.seedEnvironments(suite.T(), suite.db)
		suite.seedCharts(suite.T(), suite.db)
		suite.seedAppVersions(suite.T(), suite.db)
		suite.seedChartVersions(suite.T(), suite.db)
		suite.seedChartReleases(suite.T(), suite.db)

		suite.Run("simple kubernetes", func() {
			instance, created, err := suite.DatabaseInstanceController.Create(datarepoSwatomationDatabaseInstance, generateUser(suite.T(), suite.db, false))
			assert.NoError(suite.T(), err)
			assert.True(suite.T(), created)
			assert.True(suite.T(), instance.ID > 0)
			suite.Run("default database is chart name", func() {
				assert.Equal(suite.T(), datarepoChart.Name, *instance.DefaultDatabase)
			})
		})
		suite.Run("google", func() {
			instance, created, err := suite.DatabaseInstanceController.Create(datarepoDevDatabaseInstance, generateUser(suite.T(), suite.db, false))
			assert.NoError(suite.T(), err)
			assert.True(suite.T(), created)
			assert.True(suite.T(), instance.ID > 0)
		})
		suite.Run("azure", func() {
			instance, created, err := suite.DatabaseInstanceController.Create(CreatableDatabaseInstance{
				ChartRelease: fmt.Sprintf("%s-%s", yaleChart.Name, terraDevEnvironment.Name),
				EditableDatabaseInstance: EditableDatabaseInstance{
					Platform:     utils.PointerTo("azure"),
					InstanceName: utils.PointerTo("ghi"),
				},
			}, generateUser(suite.T(), suite.db, false))
			assert.NoError(suite.T(), err)
			assert.True(suite.T(), created)
			assert.True(suite.T(), instance.ID > 0)

		})
	})
	suite.Run("won't create duplicates", func() {
		deprecated_db.Truncate(suite.T(), suite.db)
		suite.seedClusters(suite.T(), suite.db)
		suite.seedEnvironments(suite.T(), suite.db)
		suite.seedCharts(suite.T(), suite.db)
		suite.seedAppVersions(suite.T(), suite.db)
		suite.seedChartVersions(suite.T(), suite.db)
		suite.seedChartReleases(suite.T(), suite.db)
		suite.seedDatabaseInstances(suite.T(), suite.db)

		suite.Run("exact duplicate", func() {
			_, created, err := suite.DatabaseInstanceController.Create(datarepoSwatomationDatabaseInstance, generateUser(suite.T(), suite.db, false))
			assert.ErrorContains(suite.T(), err, errors.Conflict)
			assert.False(suite.T(), created)
		})
	})
	suite.Run("can create a new database instance", func() {
		deprecated_db.Truncate(suite.T(), suite.db)
		suite.seedClusters(suite.T(), suite.db)
		suite.seedEnvironments(suite.T(), suite.db)
		suite.seedCharts(suite.T(), suite.db)
		suite.seedAppVersions(suite.T(), suite.db)
		suite.seedChartVersions(suite.T(), suite.db)
		suite.seedChartReleases(suite.T(), suite.db)

		suite.Run("no associations", func() {
			_, created, err := suite.DatabaseInstanceController.Create(CreatableDatabaseInstance{
				EditableDatabaseInstance: EditableDatabaseInstance{
					Platform:      utils.PointerTo("google"),
					GoogleProject: utils.PointerTo("broad-datarepo-dev"),
					InstanceName:  utils.PointerTo("datarepo-abcdef"),
				},
			}, generateUser(suite.T(), suite.db, false))
			assert.ErrorContains(suite.T(), err, errors.BadRequest)
			assert.False(suite.T(), created)
		})
		suite.Run("good associations but bad values", func() {
			_, created, err := suite.DatabaseInstanceController.Create(CreatableDatabaseInstance{
				ChartRelease: datarepoDevChartRelease.Name,
				EditableDatabaseInstance: EditableDatabaseInstance{
					Platform: utils.PointerTo("google"),
				},
			}, generateUser(suite.T(), suite.db, false))
			assert.ErrorContains(suite.T(), err, errors.BadRequest)
			assert.False(suite.T(), created)
		})
	})
	suite.Run("checks suitability", func() {
		deprecated_db.Truncate(suite.T(), suite.db)
		suite.seedClusters(suite.T(), suite.db)
		suite.seedEnvironments(suite.T(), suite.db)
		suite.seedCharts(suite.T(), suite.db)
		suite.seedAppVersions(suite.T(), suite.db)
		suite.seedChartVersions(suite.T(), suite.db)
		suite.seedChartReleases(suite.T(), suite.db)
		suite.Run("blocks suitable creation for non-suitable", func() {
			_, created, err := suite.DatabaseInstanceController.Create(datarepoProdDatabaseInstance, generateUser(suite.T(), suite.db, false))
			assert.ErrorContains(suite.T(), err, errors.Forbidden)
			assert.False(suite.T(), created)
		})
		suite.Run("allows suitable creation for suitable", func() {
			instance, created, err := suite.DatabaseInstanceController.Create(datarepoProdDatabaseInstance, generateUser(suite.T(), suite.db, true))
			assert.NoError(suite.T(), err)
			assert.True(suite.T(), created)
			assert.True(suite.T(), instance.ID > 0)
		})
	})
}

func (suite *chartReleaseControllerSuite) TestDatabaseInstanceListAllMatching() {
	deprecated_db.Truncate(suite.T(), suite.db)
	suite.seedClusters(suite.T(), suite.db)
	suite.seedEnvironments(suite.T(), suite.db)
	suite.seedCharts(suite.T(), suite.db)
	suite.seedAppVersions(suite.T(), suite.db)
	suite.seedChartVersions(suite.T(), suite.db)
	suite.seedChartReleases(suite.T(), suite.db)
	suite.seedDatabaseInstances(suite.T(), suite.db)

	suite.Run("lists all database instances", func() {
		matching, err := suite.DatabaseInstanceController.ListAllMatching(DatabaseInstance{}, 0)
		assert.NoError(suite.T(), err)
		assert.Equal(suite.T(), len(databaseInstanceSeedList), len(matching))
		suite.Run("orders by latest updated", func() {
			latestUpdated := matching[0].UpdatedAt
			for _, databaseInstance := range matching {
				assert.GreaterOrEqual(suite.T(), latestUpdated, databaseInstance.UpdatedAt)
			}
		})
	})
	suite.Run("limits", func() {
		limit := 2
		matching, err := suite.DatabaseInstanceController.ListAllMatching(DatabaseInstance{}, limit)
		assert.NoError(suite.T(), err)
		assert.Equal(suite.T(), limit, len(matching))
	})
	suite.Run("filters exactly", func() {
		matching, err := suite.DatabaseInstanceController.ListAllMatching(
			DatabaseInstance{CreatableDatabaseInstance: CreatableDatabaseInstance{
				ChartRelease: datarepoDevChartRelease.Name,
			}}, 0)
		assert.NoError(suite.T(), err)
		assert.Equal(suite.T(), 1, len(matching))
		assert.Equal(suite.T(), datarepoDevDatabaseInstance.Platform, matching[0].Platform)
		assert.Equal(suite.T(), datarepoDevDatabaseInstance.InstanceName, matching[0].InstanceName)
	})
	suite.Run("filters multiple", func() {
		matching, err := suite.DatabaseInstanceController.ListAllMatching(
			DatabaseInstance{CreatableDatabaseInstance: CreatableDatabaseInstance{
				EditableDatabaseInstance: EditableDatabaseInstance{
					Platform: utils.PointerTo("google"),
				},
			}}, 0)
		assert.NoError(suite.T(), err)
		assert.True(suite.T(), len(matching) > 1)
		for _, databaseInstance := range matching {
			assert.Equal(suite.T(), "google", *databaseInstance.Platform)
		}
	})
	suite.Run("none is an empty list, not null", func() {
		matching, err := suite.DatabaseInstanceController.ListAllMatching(
			DatabaseInstance{CreatableDatabaseInstance: CreatableDatabaseInstance{
				EditableDatabaseInstance: EditableDatabaseInstance{
					InstanceName: utils.PointerTo("blah"),
				},
			}}, 0)
		assert.NoError(suite.T(), err)
		assert.NotNil(suite.T(), matching)
		assert.Empty(suite.T(), matching)
	})
}

func (suite *databaseInstanceControllerSuite) TestDatabaseInstanceGet() {
	deprecated_db.Truncate(suite.T(), suite.db)
	suite.seedClusters(suite.T(), suite.db)
	suite.seedEnvironments(suite.T(), suite.db)
	suite.seedCharts(suite.T(), suite.db)
	suite.seedAppVersions(suite.T(), suite.db)
	suite.seedChartVersions(suite.T(), suite.db)
	suite.seedChartReleases(suite.T(), suite.db)
	suite.seedDatabaseInstances(suite.T(), suite.db)

	suite.Run("successfully", func() {
		byChartRelease, err := suite.DatabaseInstanceController.Get(fmt.Sprintf("chart-release/%s", datarepoDevChartRelease.Name))
		assert.NoError(suite.T(), err)
		assert.True(suite.T(), byChartRelease.ID > 0)
		byID, err := suite.DatabaseInstanceController.Get(fmt.Sprintf("%d", byChartRelease.ID))
		assert.NoError(suite.T(), err)
		assert.Equal(suite.T(), datarepoDevChartRelease.Name, byID.ChartRelease)
	})
	suite.Run("unsuccessfully for non-present", func() {
		_, err := suite.DatabaseInstanceController.Get("chart-release/foobar")
		assert.ErrorContains(suite.T(), err, errors.NotFound)
	})
	suite.Run("unsuccessfully for invalid selector", func() {
		_, err := suite.DatabaseInstanceController.Get("foobar")
		assert.ErrorContains(suite.T(), err, errors.BadRequest)
	})
}

func (suite *databaseInstanceControllerSuite) TestDatabaseInstanceGetOtherValidSelectors() {
	deprecated_db.Truncate(suite.T(), suite.db)
	suite.seedClusters(suite.T(), suite.db)
	suite.seedEnvironments(suite.T(), suite.db)
	suite.seedCharts(suite.T(), suite.db)
	suite.seedAppVersions(suite.T(), suite.db)
	suite.seedChartVersions(suite.T(), suite.db)
	suite.seedChartReleases(suite.T(), suite.db)
	suite.seedDatabaseInstances(suite.T(), suite.db)

	suite.Run("successfully", func() {
		selectors, err := suite.DatabaseInstanceController.GetOtherValidSelectors(fmt.Sprintf("chart-release/%s", datarepoDevChartRelease.Name))
		assert.NoError(suite.T(), err)
		assert.Equal(suite.T(), 5, len(selectors))
		assert.Equal(suite.T(), fmt.Sprintf("chart-release/%s", datarepoDevChartRelease.Name), selectors[0])
	})
	suite.Run("unsuccessfully for non-present", func() {
		_, err := suite.DatabaseInstanceController.GetOtherValidSelectors("chart-release/foobar")
		assert.ErrorContains(suite.T(), err, errors.NotFound)
	})
	suite.Run("unsuccessfully for invalid selector", func() {
		_, err := suite.DatabaseInstanceController.GetOtherValidSelectors("foobar")
		assert.ErrorContains(suite.T(), err, errors.BadRequest)
	})
}

func (suite *databaseInstanceControllerSuite) TestDatabaseInstanceEdit() {
	suite.Run("successfully", func() {
		deprecated_db.Truncate(suite.T(), suite.db)
		suite.seedClusters(suite.T(), suite.db)
		suite.seedEnvironments(suite.T(), suite.db)
		suite.seedCharts(suite.T(), suite.db)
		suite.seedAppVersions(suite.T(), suite.db)
		suite.seedChartVersions(suite.T(), suite.db)
		suite.seedChartReleases(suite.T(), suite.db)
		suite.seedDatabaseInstances(suite.T(), suite.db)

		before, err := suite.DatabaseInstanceController.Get(fmt.Sprintf("chart-release/%s", datarepoDevDatabaseInstance.ChartRelease))
		assert.NoError(suite.T(), err)
		assert.Equal(suite.T(), datarepoDevDatabaseInstance.InstanceName, before.InstanceName)
		newInstanceName := utils.PointerTo("new")
		edited, err := suite.DatabaseInstanceController.Edit(fmt.Sprintf("chart-release/%s", datarepoDevDatabaseInstance.ChartRelease), EditableDatabaseInstance{
			InstanceName: newInstanceName,
		}, generateUser(suite.T(), suite.db, false))
		assert.NoError(suite.T(), err)
		assert.Equal(suite.T(), newInstanceName, edited.InstanceName)
		after, err := suite.DatabaseInstanceController.Get(fmt.Sprintf("chart-release/%s", datarepoDevDatabaseInstance.ChartRelease))
		assert.NoError(suite.T(), err)
		assert.Equal(suite.T(), newInstanceName, after.InstanceName)
	})
	suite.Run("edit to suitable database instance", func() {
		deprecated_db.Truncate(suite.T(), suite.db)
		suite.seedClusters(suite.T(), suite.db)
		suite.seedEnvironments(suite.T(), suite.db)
		suite.seedCharts(suite.T(), suite.db)
		suite.seedAppVersions(suite.T(), suite.db)
		suite.seedChartVersions(suite.T(), suite.db)
		suite.seedChartReleases(suite.T(), suite.db)
		suite.seedDatabaseInstances(suite.T(), suite.db)

		newInstanceName := utils.PointerTo("new")
		suite.Run("unsuccessfully if not suitable", func() {
			_, err := suite.DatabaseInstanceController.Edit(fmt.Sprintf("chart-release/%s", datarepoProdDatabaseInstance.ChartRelease), EditableDatabaseInstance{
				InstanceName: newInstanceName,
			}, generateUser(suite.T(), suite.db, false))
			assert.ErrorContains(suite.T(), err, errors.Forbidden)
			notEdited, err := suite.DatabaseInstanceController.Get(fmt.Sprintf("chart-release/%s", datarepoProdDatabaseInstance.ChartRelease))
			assert.NoError(suite.T(), err)
			assert.NotEqual(suite.T(), newInstanceName, notEdited.InstanceName)
		})
		suite.Run("successfully if suitable", func() {
			edited, err := suite.DatabaseInstanceController.Edit(fmt.Sprintf("chart-release/%s", datarepoProdDatabaseInstance.ChartRelease), EditableDatabaseInstance{
				InstanceName: newInstanceName,
			}, generateUser(suite.T(), suite.db, true))
			assert.NoError(suite.T(), err)
			assert.Equal(suite.T(), newInstanceName, edited.InstanceName)
		})
	})
}

func (suite *databaseInstanceControllerSuite) TestDatabaseInstanceDelete() {
	suite.Run("successfully", func() {
		deprecated_db.Truncate(suite.T(), suite.db)
		suite.seedClusters(suite.T(), suite.db)
		suite.seedEnvironments(suite.T(), suite.db)
		suite.seedCharts(suite.T(), suite.db)
		suite.seedAppVersions(suite.T(), suite.db)
		suite.seedChartVersions(suite.T(), suite.db)
		suite.seedChartReleases(suite.T(), suite.db)
		suite.seedDatabaseInstances(suite.T(), suite.db)

		deleted, err := suite.DatabaseInstanceController.Delete(fmt.Sprintf("chart-release/%s", datarepoDevChartRelease.Name), generateUser(suite.T(), suite.db, false))
		assert.NoError(suite.T(), err)
		assert.Equal(suite.T(), datarepoDevChartRelease.Name, deleted.ChartRelease)
		_, err = suite.DatabaseInstanceController.Get(fmt.Sprintf("chart-release/%s", datarepoDevChartRelease.Name))
		assert.ErrorContains(suite.T(), err, errors.NotFound)
		suite.Run("allow re-creation", func() {
			instance, created, err := suite.DatabaseInstanceController.Create(datarepoDevDatabaseInstance, generateUser(suite.T(), suite.db, false))
			assert.NoError(suite.T(), err)
			assert.True(suite.T(), created)
			assert.NotEqual(suite.T(), deleted.ID, instance.ID)
		})
	})
	suite.Run("delete suitable database instance", func() {
		deprecated_db.Truncate(suite.T(), suite.db)
		suite.seedClusters(suite.T(), suite.db)
		suite.seedEnvironments(suite.T(), suite.db)
		suite.seedCharts(suite.T(), suite.db)
		suite.seedAppVersions(suite.T(), suite.db)
		suite.seedChartVersions(suite.T(), suite.db)
		suite.seedChartReleases(suite.T(), suite.db)
		suite.seedDatabaseInstances(suite.T(), suite.db)

		suite.Run("unsuccessfully if not suitable", func() {
			_, err := suite.DatabaseInstanceController.Delete(fmt.Sprintf("chart-release/%s", datarepoProdChartRelease.Name), generateUser(suite.T(), suite.db, false))
			assert.ErrorContains(suite.T(), err, errors.Forbidden)
		})
		suite.Run("successfully if suitable", func() {
			deleted, err := suite.DatabaseInstanceController.Delete(fmt.Sprintf("chart-release/%s", datarepoProdChartRelease.Name), generateUser(suite.T(), suite.db, true))
			assert.NoError(suite.T(), err)
			assert.Equal(suite.T(), datarepoProdChartRelease.Name, deleted.ChartRelease)
		})
	})
}
