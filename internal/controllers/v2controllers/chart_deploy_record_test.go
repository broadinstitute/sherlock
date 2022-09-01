package v2controllers

import (
	"fmt"
	"github.com/broadinstitute/sherlock/internal/auth"
	"github.com/broadinstitute/sherlock/internal/config"
	"github.com/broadinstitute/sherlock/internal/db"
	"github.com/broadinstitute/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/internal/models/v2models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
	"testing"
)

//
// Test suite configuration
//

func TestChartDeployRecordControllerSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping functional test")
	}
	suite.Run(t, new(chartDeployRecordControllerSuite))
}

type chartDeployRecordControllerSuite struct {
	suite.Suite
	*ControllerSet
	db *gorm.DB
}

func (suite *chartDeployRecordControllerSuite) SetupTest() {
	config.LoadTestConfig(suite.T())
	suite.db = db.ConnectAndConfigureFromTest(suite.T())
	suite.db.Begin()
	suite.ControllerSet = NewControllerSet(v2models.NewStoreSet(suite.db))
}

func (suite *chartDeployRecordControllerSuite) TearDownTest() {
	suite.db.Rollback()
}

//
// Controller seeding
//

var (
	datarepoDev1ChartDeployRecord = CreatableChartDeployRecord{
		ChartRelease:      datarepoDevChartRelease.Name,
		ExactChartVersion: "1.2.3",
		ExactAppVersion:   "0.200.0",
		HelmfileRef:       "a1b2c3",
	}
	datarepoDev2ChartDeployRecord = CreatableChartDeployRecord{
		ChartRelease:      datarepoDevChartRelease.Name,
		ExactChartVersion: "1.2.3",
		ExactAppVersion:   "0.200.1",
		HelmfileRef:       "a1b2c3",
	}
	datarepoProd1ChartDeployRecord = CreatableChartDeployRecord{
		ChartRelease:      datarepoProdChartRelease.Name,
		ExactChartVersion: "1.2.3",
		ExactAppVersion:   "0.200.0",
		HelmfileRef:       "a1b2c3",
	}
	datarepoProd2ChartDeployRecord = CreatableChartDeployRecord{
		ChartRelease:      datarepoProdChartRelease.Name,
		ExactChartVersion: "1.2.3",
		ExactAppVersion:   "0.200.1",
		HelmfileRef:       "a1b2c3",
	}
	chartDeployRecordSeedList = []CreatableChartDeployRecord{datarepoDev1ChartDeployRecord, datarepoDev2ChartDeployRecord, datarepoProd1ChartDeployRecord, datarepoProd2ChartDeployRecord}
)

func (controllerSet *ControllerSet) seedChartDeployRecords(t *testing.T) {
	for _, creatable := range chartDeployRecordSeedList {
		if _, err := controllerSet.ChartDeployRecordController.Create(creatable, auth.GenerateUser(t, true)); err != nil {
			t.Errorf("error seeding chart deploy version %s,%s,%s for chart release %s: %v", creatable.ExactChartVersion, creatable.ExactAppVersion, creatable.HelmfileRef, creatable.ChartRelease, err)
		}
	}
}

//
// Controller tests
//

func (suite *chartDeployRecordControllerSuite) TestChartDeployRecordCreate() {
	suite.Run("can create a new chart deploy record", func() {
		db.Truncate(suite.T(), suite.db)
		suite.seedClusters(suite.T())
		suite.seedEnvironments(suite.T())
		suite.seedCharts(suite.T())
		suite.seedChartReleases(suite.T())

		chartDeployRecord, err := suite.ChartDeployRecordController.Create(datarepoDev1ChartDeployRecord, auth.GenerateUser(suite.T(), false))
		assert.NoError(suite.T(), err)
		assert.Equal(suite.T(), datarepoDev1ChartDeployRecord.ExactAppVersion, chartDeployRecord.ExactAppVersion)
		assert.True(suite.T(), chartDeployRecord.ID > 0)

		suite.Run("can create duplicates", func() {
			secondChartDeployRecord, err := suite.ChartDeployRecordController.Create(datarepoDev1ChartDeployRecord, auth.GenerateUser(suite.T(), false))
			assert.NoError(suite.T(), err)
			assert.Equal(suite.T(), datarepoDev1ChartDeployRecord.ExactAppVersion, secondChartDeployRecord.ExactAppVersion)
			assert.True(suite.T(), secondChartDeployRecord.ID > 0)
		})
		suite.Run("will pull defaults from chart release", func() {
			thirdChartDeployRecord, err := suite.ChartDeployRecordController.Create(CreatableChartDeployRecord{
				ChartRelease: datarepoDevChartRelease.Name,
			}, auth.GenerateUser(suite.T(), false))
			assert.NoError(suite.T(), err)
			assert.Equal(suite.T(), *datarepoDevChartRelease.CurrentChartVersionExact, thirdChartDeployRecord.ExactChartVersion)
			assert.Equal(suite.T(), *datarepoDevChartRelease.CurrentAppVersionExact, thirdChartDeployRecord.ExactAppVersion)
			assert.Equal(suite.T(), "HEAD", thirdChartDeployRecord.HelmfileRef)
		})
	})
	suite.Run("validates incoming entries", func() {
		db.Truncate(suite.T(), suite.db)

		_, err := suite.ChartDeployRecordController.Create(CreatableChartDeployRecord{}, auth.GenerateUser(suite.T(), false))
		assert.ErrorContains(suite.T(), err, errors.BadRequest)
	})
	suite.Run("checks suitability", func() {
		db.Truncate(suite.T(), suite.db)
		suite.seedClusters(suite.T())
		suite.seedEnvironments(suite.T())
		suite.seedCharts(suite.T())
		suite.seedChartReleases(suite.T())

		suite.Run("blocks suitable creation for non-suitable", func() {
			_, err := suite.ChartDeployRecordController.Create(datarepoProd1ChartDeployRecord, auth.GenerateUser(suite.T(), false))
			assert.ErrorContains(suite.T(), err, errors.Forbidden)
		})
		suite.Run("allows suitable creation for suitable", func() {
			chartDeployRecord, err := suite.ChartDeployRecordController.Create(datarepoProd1ChartDeployRecord, auth.GenerateUser(suite.T(), true))
			assert.NoError(suite.T(), err)
			assert.True(suite.T(), chartDeployRecord.ID > 0)
		})
	})
}

func (suite *chartDeployRecordControllerSuite) TestChartDeployRecordListAllMatching() {
	db.Truncate(suite.T(), suite.db)
	suite.seedClusters(suite.T())
	suite.seedEnvironments(suite.T())
	suite.seedCharts(suite.T())
	suite.seedChartReleases(suite.T())
	suite.seedChartDeployRecords(suite.T())

	suite.Run("lists all chartDeployRecords", func() {
		matching, err := suite.ChartDeployRecordController.ListAllMatching(ChartDeployRecord{}, 0)
		assert.NoError(suite.T(), err)
		assert.Equal(suite.T(), len(chartDeployRecordSeedList), len(matching))
		suite.Run("orders by latest updated", func() {
			latestUpdated := matching[0].UpdatedAt
			for _, chartDeployRecord := range matching {
				assert.GreaterOrEqual(suite.T(), latestUpdated, chartDeployRecord.UpdatedAt)
			}
		})
	})
	suite.Run("limits", func() {
		limit := 2
		matching, err := suite.ChartDeployRecordController.ListAllMatching(ChartDeployRecord{}, limit)
		assert.NoError(suite.T(), err)
		assert.Equal(suite.T(), limit, len(matching))
	})
	suite.Run("filters exactly", func() {
		matching, err := suite.ChartDeployRecordController.ListAllMatching(ChartDeployRecord{CreatableChartDeployRecord: datarepoDev1ChartDeployRecord}, 0)
		assert.NoError(suite.T(), err)
		assert.Equal(suite.T(), 1, len(matching))
		assert.Equal(suite.T(), datarepoDev1ChartDeployRecord.ExactAppVersion, matching[0].ExactAppVersion)
	})
	suite.Run("filters multiple", func() {
		matching, err := suite.ChartDeployRecordController.ListAllMatching(ChartDeployRecord{CreatableChartDeployRecord: CreatableChartDeployRecord{ChartRelease: datarepoDevChartRelease.Name}}, 0)
		assert.NoError(suite.T(), err)
		assert.True(suite.T(), len(matching) > 1)
		for _, chartDeployRecord := range matching {
			assert.Equal(suite.T(), datarepoDevChartRelease.Name, chartDeployRecord.ChartRelease)
		}
	})
	suite.Run("none is an empty list, not null", func() {
		matching, err := suite.ChartDeployRecordController.ListAllMatching(
			ChartDeployRecord{CreatableChartDeployRecord: CreatableChartDeployRecord{ExactChartVersion: "blah"}}, 0)
		assert.NoError(suite.T(), err)
		assert.NotNil(suite.T(), matching)
		assert.Empty(suite.T(), matching)
	})
}

func (suite *chartDeployRecordControllerSuite) TestChartDeployRecordGet() {
	db.Truncate(suite.T(), suite.db)
	suite.seedClusters(suite.T())
	suite.seedEnvironments(suite.T())
	suite.seedCharts(suite.T())
	suite.seedChartReleases(suite.T())
	suite.seedChartDeployRecords(suite.T())
	suite.seedAppVersions(suite.T())
	suite.seedChartVersions(suite.T())
	chartDeployRecords, _ := suite.ChartDeployRecordController.ListAllMatching(ChartDeployRecord{
		CreatableChartDeployRecord: CreatableChartDeployRecord{
			ChartRelease: datarepoDevChartRelease.Name,
		},
	}, 1)
	anID := chartDeployRecords[0].ID

	suite.Run("successfully", func() {
		byID, err := suite.ChartDeployRecordController.Get(fmt.Sprintf("%d", anID))
		assert.NoError(suite.T(), err)
		assert.Equal(suite.T(), chartDeployRecords[0].ExactAppVersion, byID.ExactAppVersion)
	})
	suite.Run("unsuccessfully for non-present", func() {
		// Can't predict IDs so we just delete one that we know existed
		_, _ = suite.ChartDeployRecordController.Delete(fmt.Sprintf("%d", anID), auth.GenerateUser(suite.T(), false))
		_, err := suite.ChartDeployRecordController.Get(fmt.Sprintf("%d", anID))
		assert.ErrorContains(suite.T(), err, errors.NotFound)
	})
	suite.Run("unsuccessfully for invalid selector", func() {
		_, err := suite.ChartDeployRecordController.Get("something obviously invalid")
		assert.ErrorContains(suite.T(), err, errors.BadRequest)
	})
}

func (suite *chartDeployRecordControllerSuite) TestChartDeployRecordGetOtherValidSelectors() {
	db.Truncate(suite.T(), suite.db)
	suite.seedClusters(suite.T())
	suite.seedEnvironments(suite.T())
	suite.seedCharts(suite.T())
	suite.seedChartReleases(suite.T())
	suite.seedChartDeployRecords(suite.T())
	chartDeployRecords, _ := suite.ChartDeployRecordController.ListAllMatching(ChartDeployRecord{
		CreatableChartDeployRecord: CreatableChartDeployRecord{
			ChartRelease: datarepoDevChartRelease.Name,
		},
	}, 1)
	anID := fmt.Sprintf("%d", chartDeployRecords[0].ID)

	suite.Run("successfully", func() {
		selectors, err := suite.ChartDeployRecordController.GetOtherValidSelectors(anID)
		assert.NoError(suite.T(), err)
		assert.Equal(suite.T(), 1, len(selectors))
		assert.Equal(suite.T(), anID, selectors[0])
	})
	suite.Run("unsuccessfully for not found", func() {
		// Can't predict IDs so we just delete one that we know existed
		_, _ = suite.ChartDeployRecordController.Delete(anID, auth.GenerateUser(suite.T(), false))
		_, err := suite.ChartDeployRecordController.GetOtherValidSelectors(anID)
		assert.ErrorContains(suite.T(), err, errors.NotFound)
	})
	suite.Run("unsuccessfully for invalid", func() {
		_, err := suite.ChartDeployRecordController.GetOtherValidSelectors("something obviously invalid")
		assert.ErrorContains(suite.T(), err, errors.BadRequest)
	})
}

func (suite *chartDeployRecordControllerSuite) TestChartDeployRecordEdit() {
	suite.Run("'successfully' but there's no fields", func() {
		db.Truncate(suite.T(), suite.db)
		suite.seedClusters(suite.T())
		suite.seedEnvironments(suite.T())
		suite.seedCharts(suite.T())
		suite.seedChartReleases(suite.T())
		suite.seedChartDeployRecords(suite.T())
		chartDeployRecords, _ := suite.ChartDeployRecordController.ListAllMatching(ChartDeployRecord{
			CreatableChartDeployRecord: CreatableChartDeployRecord{
				ChartRelease: datarepoDevChartRelease.Name,
			},
		}, 1)
		anID := fmt.Sprintf("%d", chartDeployRecords[0].ID)

		_, err := suite.ChartDeployRecordController.Edit(anID, EditableChartDeployRecord{}, auth.GenerateUser(suite.T(), false))
		assert.NoError(suite.T(), err)
	})
	suite.Run("checks suitability despite there being no fields", func() {
		db.Truncate(suite.T(), suite.db)
		suite.seedClusters(suite.T())
		suite.seedEnvironments(suite.T())
		suite.seedCharts(suite.T())
		suite.seedChartReleases(suite.T())
		suite.seedChartDeployRecords(suite.T())
		chartDeployRecords, _ := suite.ChartDeployRecordController.ListAllMatching(ChartDeployRecord{
			CreatableChartDeployRecord: CreatableChartDeployRecord{
				// Filter for prod ones to find something requiring suitability
				ChartRelease: datarepoProdChartRelease.Name,
			},
		}, 1)
		anID := fmt.Sprintf("%d", chartDeployRecords[0].ID)

		suite.Run("blocks non-suitable callers", func() {
			_, err := suite.ChartDeployRecordController.Edit(anID, EditableChartDeployRecord{}, auth.GenerateUser(suite.T(), false))
			assert.ErrorContains(suite.T(), err, errors.Forbidden)

		})
		suite.Run("allows suitable callers", func() {
			_, err := suite.ChartDeployRecordController.Edit(anID, EditableChartDeployRecord{}, auth.GenerateUser(suite.T(), true))
			assert.NoError(suite.T(), err)
		})
	})
}

func (suite *chartDeployRecordControllerSuite) TestChartDeployRecordDelete() {
	suite.Run("successfully", func() {
		db.Truncate(suite.T(), suite.db)
		suite.seedClusters(suite.T())
		suite.seedEnvironments(suite.T())
		suite.seedCharts(suite.T())
		suite.seedChartReleases(suite.T())
		suite.seedChartDeployRecords(suite.T())
		chartDeployRecords, _ := suite.ChartDeployRecordController.ListAllMatching(ChartDeployRecord{
			CreatableChartDeployRecord: CreatableChartDeployRecord{
				// Filter for prod ones to find something requiring suitability
				ChartRelease: datarepoProdChartRelease.Name,
			},
		}, 1)
		anID := fmt.Sprintf("%d", chartDeployRecords[0].ID)

		suite.Run("checks suitability", func() {
			_, err := suite.ChartDeployRecordController.Delete(anID, auth.GenerateUser(suite.T(), false))
			assert.ErrorContains(suite.T(), err, errors.Forbidden)
		})

		deleted, err := suite.ChartDeployRecordController.Delete(anID, auth.GenerateUser(suite.T(), true))
		assert.NoError(suite.T(), err)
		assert.Equal(suite.T(), chartDeployRecords[0].ExactAppVersion, deleted.ExactAppVersion)
		_, err = suite.ChartDeployRecordController.Get(anID)
		assert.ErrorContains(suite.T(), err, errors.NotFound)
	})
}
