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

func (suite *chartVersionControllerSuite) SetupTest() {
	config.LoadTestConfig(suite.T())
	suite.db = db.ConnectAndConfigureFromTest(suite.T())
	suite.db.Begin()
	suite.ControllerSet = NewControllerSet(v2models.NewStoreSet(suite.db))
}

func (suite *chartVersionControllerSuite) TearDownTest() {
	suite.db.Rollback()
}

//
// Controller seeding
//

var (
	leonardoMain1ChartVersion = CreatableChartVersion{
		Chart:        leonardoChart.Name,
		ChartVersion: "1.2.3",
	}
	leonardoMain2ChartVersion = CreatableChartVersion{
		Chart:              leonardoChart.Name,
		ChartVersion:       "1.2.4",
		ParentChartVersion: fmt.Sprintf("%s/%s", leonardoChart.Name, leonardoMain1ChartVersion.ChartVersion),
	}
	leonardoMain3ChartVersion = CreatableChartVersion{
		Chart:              leonardoChart.Name,
		ChartVersion:       "1.2.5",
		ParentChartVersion: fmt.Sprintf("%s/%s", leonardoChart.Name, leonardoMain2ChartVersion.ChartVersion),
	}
	leonardoBranch1ChartVersion = CreatableChartVersion{
		Chart:              leonardoChart.Name,
		ChartVersion:       "1.2.4-a1c1",
		ParentChartVersion: fmt.Sprintf("%s/%s", leonardoChart.Name, leonardoMain1ChartVersion.ChartVersion),
	}
	leonardoBranch2ChartVersion = CreatableChartVersion{
		Chart:              leonardoChart.Name,
		ChartVersion:       "1.2.4-a1c2",
		ParentChartVersion: fmt.Sprintf("%s/%s", leonardoChart.Name, leonardoBranch1ChartVersion.ChartVersion),
	}
	leonardoBranch3ChartVersion = CreatableChartVersion{
		Chart:              leonardoChart.Name,
		ChartVersion:       "1.2.4-a1c3",
		ParentChartVersion: fmt.Sprintf("%s/%s", leonardoChart.Name, leonardoBranch2ChartVersion.ChartVersion),
	}
	chartVersionSeedList = []CreatableChartVersion{leonardoMain1ChartVersion, leonardoMain2ChartVersion, leonardoMain3ChartVersion, leonardoBranch1ChartVersion, leonardoBranch2ChartVersion, leonardoBranch3ChartVersion}
)

func (controllerSet *ControllerSet) seedChartVersions(t *testing.T) {
	for _, creatable := range chartVersionSeedList {
		if _, _, err := controllerSet.ChartVersionController.Create(creatable, auth.GenerateUser(t, false)); err != nil {
			t.Errorf("error seeding app version %s for chart %s: %v", creatable.ChartVersion, creatable.Chart, err)
		}
	}
}

//
// Controller tests
//

func (suite *chartVersionControllerSuite) TestChartVersionCreate() {
	suite.Run("can create a new app version", func() {
		db.Truncate(suite.T(), suite.db)
		suite.seedCharts(suite.T())

		chartVersion, created, err := suite.ChartVersionController.Create(leonardoMain1ChartVersion, auth.GenerateUser(suite.T(), false))
		assert.NoError(suite.T(), err)
		assert.True(suite.T(), created)
		assert.Equal(suite.T(), leonardoMain1ChartVersion.ChartVersion, chartVersion.ChartVersion)
		assert.True(suite.T(), chartVersion.ID > 0)

		suite.Run("can accept duplicates", func() {
			secondChartVersion, created, err := suite.ChartVersionController.Create(leonardoMain1ChartVersion, auth.GenerateUser(suite.T(), false))
			assert.NoError(suite.T(), err)
			assert.False(suite.T(), created)
			assert.Equal(suite.T(), leonardoMain1ChartVersion.ChartVersion, secondChartVersion.ChartVersion)
			assert.True(suite.T(), secondChartVersion.ID > 0)
		})
	})
	suite.Run("validates incoming entries", func() {
		db.Truncate(suite.T(), suite.db)

		_, created, err := suite.ChartVersionController.Create(CreatableChartVersion{}, auth.GenerateUser(suite.T(), false))
		assert.ErrorContains(suite.T(), err, errors.BadRequest)
		assert.False(suite.T(), created)
	})
	suite.Run("rejects mismatched duplicates", func() {
		db.Truncate(suite.T(), suite.db)
		suite.seedCharts(suite.T())
		suite.seedChartVersions(suite.T())

		_, created, err := suite.ChartVersionController.Create(CreatableChartVersion{
			Chart:        leonardoChart.Name,
			ChartVersion: "1.2.5",
			// Mismatched parent
			ParentChartVersion: fmt.Sprintf("%s/%s", leonardoChart.Name, leonardoMain1ChartVersion.ChartVersion),
		}, auth.GenerateUser(suite.T(), false))
		assert.ErrorContains(suite.T(), err, errors.Conflict)
		assert.False(suite.T(), created)
	})
	suite.Run("accepts bad parents", func() {
		db.Truncate(suite.T(), suite.db)
		suite.seedCharts(suite.T())
		suite.seedChartVersions(suite.T())

		chartVersion, created, err := suite.ChartVersionController.Create(CreatableChartVersion{
			Chart:        datarepoChart.Name,
			ChartVersion: "1.1.1",
			// Nonexistent parent
			ParentChartVersion: fmt.Sprintf("%s/%s", datarepoChart.Name, leonardoMain1ChartVersion.ChartVersion),
		}, auth.GenerateUser(suite.T(), false))
		assert.NoError(suite.T(), err)
		assert.True(suite.T(), created)
		assert.Empty(suite.T(), chartVersion.ParentChartVersion)
	})
}

func (suite *chartVersionControllerSuite) TestChartVersionListAllMatching() {
	db.Truncate(suite.T(), suite.db)
	suite.seedCharts(suite.T())
	suite.seedChartVersions(suite.T())

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
		matching, err := suite.ChartVersionController.ListAllMatching(ChartVersion{CreatableChartVersion: leonardoBranch3ChartVersion}, 0)
		assert.NoError(suite.T(), err)
		assert.Equal(suite.T(), 1, len(matching))
		assert.Equal(suite.T(), leonardoBranch3ChartVersion.ChartVersion, matching[0].ChartVersion)
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

func (suite *chartVersionControllerSuite) TestChartVersionGet() {
	db.Truncate(suite.T(), suite.db)
	suite.seedCharts(suite.T())
	suite.seedChartVersions(suite.T())
	chartVersions, _ := suite.ChartVersionController.ListAllMatching(ChartVersion{}, 1)
	anID := chartVersions[0].ID

	suite.Run("successfully", func() {
		byID, err := suite.ChartVersionController.Get(fmt.Sprintf("%d", anID))
		assert.NoError(suite.T(), err)
		assert.Equal(suite.T(), chartVersions[0].ChartVersion, byID.ChartVersion)
	})
	suite.Run("unsuccessfully for non-present", func() {
		// Can't predict IDs so we just delete one that we know existed
		_, _ = suite.ChartVersionController.Delete(fmt.Sprintf("%d", anID), auth.GenerateUser(suite.T(), false))
		_, err := suite.ChartVersionController.Get(fmt.Sprintf("%d", anID))
		assert.ErrorContains(suite.T(), err, errors.NotFound)
	})
	suite.Run("unsuccessfully for invalid selector", func() {
		_, err := suite.ChartVersionController.Get("something obviously invalid")
		assert.ErrorContains(suite.T(), err, errors.BadRequest)
	})
}

func (suite *chartVersionControllerSuite) TestChartVersionGetOtherValidSelectors() {
	db.Truncate(suite.T(), suite.db)
	suite.seedCharts(suite.T())
	suite.seedChartVersions(suite.T())
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
		_, _ = suite.ChartVersionController.Delete(anID, auth.GenerateUser(suite.T(), false))
		_, err := suite.ChartVersionController.GetOtherValidSelectors(anID)
		assert.ErrorContains(suite.T(), err, errors.NotFound)
	})
	suite.Run("unsuccessfully for invalid", func() {
		_, err := suite.ChartVersionController.GetOtherValidSelectors("something obviously invalid")
		assert.ErrorContains(suite.T(), err, errors.BadRequest)
	})
}

func (suite *chartVersionControllerSuite) TestChartVersionEdit() {
	suite.Run("'successfully' but there's no fields", func() {
		db.Truncate(suite.T(), suite.db)
		suite.seedCharts(suite.T())
		suite.seedChartVersions(suite.T())
		chartVersions, _ := suite.ChartVersionController.ListAllMatching(ChartVersion{}, 1)
		anID := fmt.Sprintf("%d", chartVersions[0].ID)

		_, err := suite.ChartVersionController.Edit(anID, EditableChartVersion{}, auth.GenerateUser(suite.T(), false))
		assert.NoError(suite.T(), err)
	})
}

func (suite *chartVersionControllerSuite) TestChartVersionDelete() {
	suite.Run("successfully", func() {
		db.Truncate(suite.T(), suite.db)
		suite.seedCharts(suite.T())
		suite.seedChartVersions(suite.T())
		chartVersions, _ := suite.ChartVersionController.ListAllMatching(ChartVersion{}, 1)
		anID := fmt.Sprintf("%d", chartVersions[0].ID)

		deleted, err := suite.ChartVersionController.Delete(anID, auth.GenerateUser(suite.T(), false))
		assert.NoError(suite.T(), err)
		assert.Equal(suite.T(), chartVersions[0].ChartVersion, deleted.ChartVersion)
		_, err = suite.ChartVersionController.Get(anID)
		assert.ErrorContains(suite.T(), err, errors.NotFound)
	})
}
