package v2controllers

import (
	"fmt"
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

func TestChartVersionControllerSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping functional test")
	}
	suite.Run(t, new(chartVersionControllerSuite))
}

type chartVersionControllerSuite struct {
	suite.Suite
	*ControllerSet
	db *gorm.DB
}

func (suite *chartVersionControllerSuite) SetupSuite() {
	config.LoadTestConfig()
	suite.db = deprecated_db.ConnectAndConfigureFromTest(suite.T())
	suite.ControllerSet = NewControllerSet(v2models.NewStoreSet(suite.db))
}

func (suite *chartVersionControllerSuite) TearDownSuite() {
	deprecated_db.Truncate(suite.T(), suite.db)
}

//
// Controller seeding
//

var (
	leonardo1ChartVersion = CreatableChartVersion{
		Chart:        leonardoChart.Name,
		ChartVersion: "1.2.3",
	}
	leonardo2ChartVersion = CreatableChartVersion{
		Chart:              leonardoChart.Name,
		ChartVersion:       "1.2.4",
		ParentChartVersion: fmt.Sprintf("%s/%s", leonardoChart.Name, leonardo1ChartVersion.ChartVersion),
	}
	leonardo3ChartVersion = CreatableChartVersion{
		Chart:              leonardoChart.Name,
		ChartVersion:       "1.2.5",
		ParentChartVersion: fmt.Sprintf("%s/%s", leonardoChart.Name, leonardo2ChartVersion.ChartVersion),
	}
	sam1ChartVersion = CreatableChartVersion{
		Chart:        samChart.Name,
		ChartVersion: "0.0.1",
	}
	sam2ChartVersion = CreatableChartVersion{
		Chart:              samChart.Name,
		ChartVersion:       "0.0.2",
		ParentChartVersion: fmt.Sprintf("%s/%s", samChart.Name, sam1ChartVersion.ChartVersion),
	}
	yale1ChartVersion = CreatableChartVersion{
		Chart:        yaleChart.Name,
		ChartVersion: "0.0.100",
	}
	terraClusterStorage1ChartVersion = CreatableChartVersion{
		Chart:        terraClusterStorageChart.Name,
		ChartVersion: "0.0.1",
	}
	datarepo1ChartVersion = CreatableChartVersion{
		Chart:        datarepoChart.Name,
		ChartVersion: "1.1.1",
	}
	honeycomb1ChartVersion = CreatableChartVersion{
		Chart:        honeycombChart.Name,
		ChartVersion: "0.0.0",
	}
	chartVersionSeedList = []CreatableChartVersion{
		leonardo1ChartVersion,
		leonardo2ChartVersion,
		leonardo3ChartVersion,
		sam1ChartVersion,
		sam2ChartVersion,
		yale1ChartVersion,
		terraClusterStorage1ChartVersion,
		datarepo1ChartVersion,
		honeycomb1ChartVersion,
	}
)

func (controllerSet *ControllerSet) seedChartVersions(t *testing.T, db *gorm.DB) {
	for _, creatable := range chartVersionSeedList {
		if _, _, err := controllerSet.ChartVersionController.Create(creatable, generateUser(t, db, false)); err != nil {
			t.Errorf("error seeding app version %s for chart %s: %v", creatable.ChartVersion, creatable.Chart, err)
		}
	}
}

//
// Controller tests
//

func (suite *chartVersionControllerSuite) TestChartVersionCreate() {
	suite.Run("can create a new app version", func() {
		deprecated_db.Truncate(suite.T(), suite.db)
		suite.seedCharts(suite.T(), suite.db)

		chartVersion, created, err := suite.ChartVersionController.Create(leonardo1ChartVersion, generateUser(suite.T(), suite.db, false))
		assert.NoError(suite.T(), err)
		assert.True(suite.T(), created)
		assert.Equal(suite.T(), leonardo1ChartVersion.ChartVersion, chartVersion.ChartVersion)
		assert.True(suite.T(), chartVersion.ID > 0)

		suite.Run("can accept duplicates", func() {
			secondChartVersion, created, err := suite.ChartVersionController.Create(leonardo1ChartVersion, generateUser(suite.T(), suite.db, false))
			assert.NoError(suite.T(), err)
			assert.False(suite.T(), created)
			assert.Equal(suite.T(), leonardo1ChartVersion.ChartVersion, secondChartVersion.ChartVersion)
			assert.True(suite.T(), secondChartVersion.ID > 0)
		})
	})
	suite.Run("validates incoming entries", func() {
		deprecated_db.Truncate(suite.T(), suite.db)

		_, created, err := suite.ChartVersionController.Create(CreatableChartVersion{}, generateUser(suite.T(), suite.db, false))
		assert.ErrorContains(suite.T(), err, errors.BadRequest)
		assert.False(suite.T(), created)
	})
	suite.Run("rejects mismatched duplicates", func() {
		deprecated_db.Truncate(suite.T(), suite.db)
		suite.seedCharts(suite.T(), suite.db)
		suite.seedChartVersions(suite.T(), suite.db)

		_, created, err := suite.ChartVersionController.Create(CreatableChartVersion{
			Chart:        leonardoChart.Name,
			ChartVersion: "1.2.5",
			// Mismatched parent
			ParentChartVersion: fmt.Sprintf("%s/%s", leonardoChart.Name, leonardo1ChartVersion.ChartVersion),
		}, generateUser(suite.T(), suite.db, false))
		assert.ErrorContains(suite.T(), err, errors.Conflict)
		assert.False(suite.T(), created)
	})
	suite.Run("accepts bad parents", func() {
		deprecated_db.Truncate(suite.T(), suite.db)
		suite.seedCharts(suite.T(), suite.db)
		suite.seedChartVersions(suite.T(), suite.db)

		chartVersion, created, err := suite.ChartVersionController.Create(CreatableChartVersion{
			Chart:        datarepoChart.Name,
			ChartVersion: "1.1.2",
			// Nonexistent parent
			ParentChartVersion: fmt.Sprintf("%s/%s", datarepoChart.Name, leonardo1ChartVersion.ChartVersion),
		}, generateUser(suite.T(), suite.db, false))
		assert.NoError(suite.T(), err)
		assert.True(suite.T(), created)
		assert.Empty(suite.T(), chartVersion.ParentChartVersion)
	})
}

func (suite *chartVersionControllerSuite) TestChartVersionListAllMatching() {
	deprecated_db.Truncate(suite.T(), suite.db)
	suite.seedCharts(suite.T(), suite.db)
	suite.seedChartVersions(suite.T(), suite.db)

	suite.Run("lists all chartVersions", func() {
		matching, err := suite.ChartVersionController.ListAllMatching(ChartVersion{}, 0)
		assert.NoError(suite.T(), err)
		assert.Equal(suite.T(), len(chartVersionSeedList), len(matching))
		suite.Run("orders by latest updated", func() {
			latestUpdated := matching[0].UpdatedAt
			for _, chartVersion := range matching {
				assert.GreaterOrEqual(suite.T(), latestUpdated, chartVersion.UpdatedAt)
			}
		})
	})
	suite.Run("limits", func() {
		limit := 2
		matching, err := suite.ChartVersionController.ListAllMatching(ChartVersion{}, limit)
		assert.NoError(suite.T(), err)
		assert.Equal(suite.T(), limit, len(matching))
	})
	suite.Run("filters exactly", func() {
		matching, err := suite.ChartVersionController.ListAllMatching(ChartVersion{CreatableChartVersion: leonardo3ChartVersion}, 0)
		assert.NoError(suite.T(), err)
		assert.Equal(suite.T(), 1, len(matching))
		assert.Equal(suite.T(), leonardo3ChartVersion.ChartVersion, matching[0].ChartVersion)
	})
	suite.Run("filters multiple", func() {
		matching, err := suite.ChartVersionController.ListAllMatching(ChartVersion{CreatableChartVersion: CreatableChartVersion{Chart: leonardoChart.Name}}, 0)
		assert.NoError(suite.T(), err)
		assert.True(suite.T(), len(matching) > 1)
		for _, chartVersion := range matching {
			assert.Equal(suite.T(), "leonardo", chartVersion.Chart)
		}
	})
	suite.Run("none is an empty list, not null", func() {
		matching, err := suite.ChartVersionController.ListAllMatching(
			ChartVersion{CreatableChartVersion: CreatableChartVersion{ChartVersion: "blah"}}, 0)
		assert.NoError(suite.T(), err)
		assert.NotNil(suite.T(), matching)
		assert.Empty(suite.T(), matching)
	})
}

func (suite *chartVersionControllerSuite) TestChartVersionGetChildrenPathToParent() {
	deprecated_db.Truncate(suite.T(), suite.db)
	suite.seedCharts(suite.T(), suite.db)
	suite.seedChartVersions(suite.T(), suite.db)

	childSelector := fmt.Sprintf("%s/%s", leonardoChart.Name, leonardo3ChartVersion.ChartVersion)
	parentSelector := fmt.Sprintf("%s/%s", leonardoChart.Name, leonardo1ChartVersion.ChartVersion)

	suite.Run("handles same", func() {
		path, connected, err := suite.ChartVersionController.GetChildrenPathToParent(childSelector, childSelector)
		suite.Assert().NoError(err)
		suite.Assert().True(connected)
		suite.Assert().Len(path, 0)
		path, connected, err = suite.ChartVersionController.GetChildrenPathToParent(parentSelector, parentSelector)
		suite.Assert().NoError(err)
		suite.Assert().True(connected)
		suite.Assert().Len(path, 0)
	})
	suite.Run("handles different but connected", func() {
		path, connected, err := suite.ChartVersionController.GetChildrenPathToParent(childSelector, parentSelector)
		suite.Assert().NoError(err)
		suite.Assert().True(connected)
		suite.Assert().Len(path, 2)
		suite.Assert().Equal(leonardo2ChartVersion.ChartVersion, path[1].ChartVersion)
	})
	suite.Run("handles different but disconnected", func() {
		path, connected, err := suite.ChartVersionController.GetChildrenPathToParent(parentSelector, childSelector)
		suite.Assert().NoError(err)
		suite.Assert().False(connected)
		suite.Assert().Len(path, 1)
	})
}

func (suite *chartVersionControllerSuite) TestChartVersionGet() {
	deprecated_db.Truncate(suite.T(), suite.db)
	suite.seedCharts(suite.T(), suite.db)
	suite.seedChartVersions(suite.T(), suite.db)
	chartVersions, _ := suite.ChartVersionController.ListAllMatching(ChartVersion{}, 1)
	anID := chartVersions[0].ID

	suite.Run("successfully", func() {
		byID, err := suite.ChartVersionController.Get(fmt.Sprintf("%d", anID))
		assert.NoError(suite.T(), err)
		assert.Equal(suite.T(), chartVersions[0].ChartVersion, byID.ChartVersion)
	})
	suite.Run("unsuccessfully for non-present", func() {
		// Can't predict IDs so we just delete one that we know existed
		_, _ = suite.ChartVersionController.Delete(fmt.Sprintf("%d", anID), generateUser(suite.T(), suite.db, false))
		_, err := suite.ChartVersionController.Get(fmt.Sprintf("%d", anID))
		assert.ErrorContains(suite.T(), err, errors.NotFound)
	})
	suite.Run("unsuccessfully for invalid selector", func() {
		_, err := suite.ChartVersionController.Get("something obviously invalid")
		assert.ErrorContains(suite.T(), err, errors.BadRequest)
	})
}

func (suite *chartVersionControllerSuite) TestChartVersionGetOtherValidSelectors() {
	deprecated_db.Truncate(suite.T(), suite.db)
	suite.seedCharts(suite.T(), suite.db)
	suite.seedChartVersions(suite.T(), suite.db)
	chartVersions, _ := suite.ChartVersionController.ListAllMatching(ChartVersion{}, 1)
	anID := fmt.Sprintf("%d", chartVersions[0].ID)

	suite.Run("successfully", func() {
		selectors, err := suite.ChartVersionController.GetOtherValidSelectors(anID)
		assert.NoError(suite.T(), err)
		assert.Equal(suite.T(), 3, len(selectors))
		assert.Equal(suite.T(), anID, selectors[0])
	})
	suite.Run("unsuccessfully for not found", func() {
		// Can't predict IDs so we just delete one that we know existed
		_, _ = suite.ChartVersionController.Delete(anID, generateUser(suite.T(), suite.db, false))
		_, err := suite.ChartVersionController.GetOtherValidSelectors(anID)
		assert.ErrorContains(suite.T(), err, errors.NotFound)
	})
	suite.Run("unsuccessfully for invalid", func() {
		_, err := suite.ChartVersionController.GetOtherValidSelectors("something obviously invalid")
		assert.ErrorContains(suite.T(), err, errors.BadRequest)
	})
}

func (suite *chartVersionControllerSuite) TestChartVersionEdit() {
	suite.Run("successfully", func() {
		deprecated_db.Truncate(suite.T(), suite.db)
		suite.seedCharts(suite.T(), suite.db)
		suite.seedChartVersions(suite.T(), suite.db)
		chartVersions, _ := suite.ChartVersionController.ListAllMatching(ChartVersion{}, 1)
		anID := fmt.Sprintf("%d", chartVersions[0].ID)

		response, err := suite.ChartVersionController.Edit(anID, EditableChartVersion{Description: "blah"}, generateUser(suite.T(), suite.db, false))
		assert.NoError(suite.T(), err)
		assert.Equal(suite.T(), "blah", response.Description)
	})
}

func (suite *chartVersionControllerSuite) TestChartVersionDelete() {
	suite.Run("successfully", func() {
		deprecated_db.Truncate(suite.T(), suite.db)
		suite.seedCharts(suite.T(), suite.db)
		suite.seedChartVersions(suite.T(), suite.db)
		chartVersions, _ := suite.ChartVersionController.ListAllMatching(ChartVersion{}, 1)
		anID := fmt.Sprintf("%d", chartVersions[0].ID)

		deleted, err := suite.ChartVersionController.Delete(anID, generateUser(suite.T(), suite.db, false))
		assert.NoError(suite.T(), err)
		assert.Equal(suite.T(), chartVersions[0].ChartVersion, deleted.ChartVersion)
		_, err = suite.ChartVersionController.Get(anID)
		assert.ErrorContains(suite.T(), err, errors.NotFound)
	})
}
