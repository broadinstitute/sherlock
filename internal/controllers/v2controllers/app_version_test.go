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

func TestAppVersionControllerSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping functional test")
	}
	suite.Run(t, new(appVersionControllerSuite))
}

type appVersionControllerSuite struct {
	suite.Suite
	*ControllerSet
	db *gorm.DB
}

func (suite *appVersionControllerSuite) SetupTest() {
	config.LoadTestConfig(suite.T())
	suite.db = db.ConnectAndConfigureFromTest(suite.T())
	suite.db.Begin()
	suite.ControllerSet = NewControllerSet(v2models.NewStoreSet(suite.db))
}

func (suite *appVersionControllerSuite) TearDownTest() {
	suite.db.Rollback()
}

//
// Controller seeding
//

var (
	leonardoMain1AppVersion = CreatableAppVersion{
		Chart:      leonardoChart.Name,
		AppVersion: "1.2.3",
		GitCommit:  "a1b1",
		GitBranch:  *leonardoChart.AppImageGitMainBranch,
	}
	leonardoMain2AppVersion = CreatableAppVersion{
		Chart:      leonardoChart.Name,
		AppVersion: "1.2.4",
		GitCommit:  "a1b2",
		GitBranch:  *leonardoChart.AppImageGitMainBranch,
	}
	leonardoMain3AppVersion = CreatableAppVersion{
		Chart:      leonardoChart.Name,
		AppVersion: "1.2.5",
		GitCommit:  "a1b3",
		GitBranch:  *leonardoChart.AppImageGitMainBranch,
	}
	leonardoBranch1AppVersion = CreatableAppVersion{
		Chart:      leonardoChart.Name,
		AppVersion: "1.2.3-a1c1",
		GitCommit:  "a1c1",
		GitBranch:  "branchy-branch",
	}
	leonardoBranch2AppVersion = CreatableAppVersion{
		Chart:      leonardoChart.Name,
		AppVersion: "1.2.3-a1c2",
		GitCommit:  "a1c2",
		GitBranch:  "branchy-branch",
	}
	leonardoBranch3AppVersion = CreatableAppVersion{
		Chart:      leonardoChart.Name,
		AppVersion: "1.2.3-a1c3",
		GitCommit:  "a1c3",
		GitBranch:  "branchy-branch",
	}
	appVersionSeedList = []CreatableAppVersion{leonardoMain1AppVersion, leonardoMain2AppVersion, leonardoMain3AppVersion, leonardoBranch1AppVersion, leonardoBranch2AppVersion, leonardoBranch3AppVersion}
)

func (controllerSet *ControllerSet) seedAppVersions(t *testing.T) {
	for _, creatable := range appVersionSeedList {
		if _, _, err := controllerSet.AppVersionController.Create(creatable, auth.GenerateUser(t, false)); err != nil {
			t.Errorf("error seeding app version %s for chart %s: %v", creatable.AppVersion, creatable.Chart, err)
		}
	}
}

//
// Controller tests
//

func (suite *appVersionControllerSuite) TestAppVersionCreate() {
	suite.Run("can create a new app version", func() {
		db.Truncate(suite.T(), suite.db)
		suite.seedCharts(suite.T())

		appVersion, created, err := suite.AppVersionController.Create(leonardoBranch3AppVersion, auth.GenerateUser(suite.T(), false))
		assert.NoError(suite.T(), err)
		assert.True(suite.T(), created)
		assert.Equal(suite.T(), leonardoBranch3AppVersion.AppVersion, appVersion.AppVersion)
		assert.True(suite.T(), appVersion.ID > 0)

		suite.Run("can create duplicates", func() {
			secondAppVersion, created, err := suite.AppVersionController.Create(leonardoBranch3AppVersion, auth.GenerateUser(suite.T(), false))
			assert.NoError(suite.T(), err)
			assert.True(suite.T(), created)
			assert.Equal(suite.T(), leonardoBranch3AppVersion.AppVersion, secondAppVersion.AppVersion)
			assert.True(suite.T(), secondAppVersion.ID > 0)
		})
	})
	suite.Run("validates incoming entries", func() {
		db.Truncate(suite.T(), suite.db)

		_, created, err := suite.AppVersionController.Create(CreatableAppVersion{}, auth.GenerateUser(suite.T(), false))
		assert.ErrorContains(suite.T(), err, errors.BadRequest)
		assert.False(suite.T(), created)
	})
}

func (suite *appVersionControllerSuite) TestAppVersionListAllMatching() {
	db.Truncate(suite.T(), suite.db)
	suite.seedCharts(suite.T())
	suite.seedAppVersions(suite.T())

	suite.Run("lists all appVersions", func() {
		matching, err := suite.AppVersionController.ListAllMatching(AppVersion{}, 0)
		assert.NoError(suite.T(), err)
		assert.Equal(suite.T(), len(appVersionSeedList), len(matching))
		suite.Run("orders by latest updated", func() {
			latestUpdated := matching[0].UpdatedAt
			for _, appVersion := range matching {
				assert.GreaterOrEqual(suite.T(), latestUpdated, appVersion.UpdatedAt)
			}
		})
	})
	suite.Run("limits", func() {
		limit := 2
		matching, err := suite.AppVersionController.ListAllMatching(AppVersion{}, limit)
		assert.NoError(suite.T(), err)
		assert.Equal(suite.T(), limit, len(matching))
	})
	suite.Run("filters exactly", func() {
		matching, err := suite.AppVersionController.ListAllMatching(AppVersion{CreatableAppVersion: leonardoBranch3AppVersion}, 0)
		assert.NoError(suite.T(), err)
		assert.Equal(suite.T(), 1, len(matching))
		assert.Equal(suite.T(), leonardoBranch3AppVersion.AppVersion, matching[0].AppVersion)
	})
	suite.Run("filters multiple", func() {
		matching, err := suite.AppVersionController.ListAllMatching(AppVersion{CreatableAppVersion: CreatableAppVersion{Chart: leonardoChart.Name}}, 0)
		assert.NoError(suite.T(), err)
		assert.True(suite.T(), len(matching) > 1)
		for _, appVersion := range matching {
			assert.Equal(suite.T(), "leonardo", appVersion.Chart)
		}
	})
	suite.Run("none is an empty list, not null", func() {
		matching, err := suite.AppVersionController.ListAllMatching(
			AppVersion{CreatableAppVersion: CreatableAppVersion{AppVersion: "blah"}}, 0)
		assert.NoError(suite.T(), err)
		assert.NotNil(suite.T(), matching)
		assert.Empty(suite.T(), matching)
	})
}

func (suite *appVersionControllerSuite) TestAppVersionGet() {
	db.Truncate(suite.T(), suite.db)
	suite.seedCharts(suite.T())
	suite.seedAppVersions(suite.T())
	appVersions, _ := suite.AppVersionController.ListAllMatching(AppVersion{}, 1)
	anID := appVersions[0].ID

	suite.Run("successfully", func() {
		byID, err := suite.AppVersionController.Get(fmt.Sprintf("%d", anID))
		assert.NoError(suite.T(), err)
		assert.Equal(suite.T(), appVersions[0].AppVersion, byID.AppVersion)
	})
	suite.Run("unsuccessfully for non-present", func() {
		// Can't predict IDs so we just delete one that we know existed
		_, _ = suite.AppVersionController.Delete(fmt.Sprintf("%d", anID), auth.GenerateUser(suite.T(), false))
		_, err := suite.AppVersionController.Get(fmt.Sprintf("%d", anID))
		assert.ErrorContains(suite.T(), err, errors.NotFound)
	})
	suite.Run("unsuccessfully for invalid selector", func() {
		_, err := suite.AppVersionController.Get("something obviously invalid")
		assert.ErrorContains(suite.T(), err, errors.BadRequest)
	})
}

func (suite *appVersionControllerSuite) TestAppVersionGetOtherValidSelectors() {
	db.Truncate(suite.T(), suite.db)
	suite.seedCharts(suite.T())
	suite.seedAppVersions(suite.T())
	appVersions, _ := suite.AppVersionController.ListAllMatching(AppVersion{}, 1)
	anID := fmt.Sprintf("%d", appVersions[0].ID)

	suite.Run("successfully", func() {
		selectors, err := suite.AppVersionController.GetOtherValidSelectors(anID)
		assert.NoError(suite.T(), err)
		assert.Equal(suite.T(), 1, len(selectors))
		assert.Equal(suite.T(), anID, selectors[0])
	})
	suite.Run("unsuccessfully for not found", func() {
		// Can't predict IDs so we just delete one that we know existed
		_, _ = suite.AppVersionController.Delete(anID, auth.GenerateUser(suite.T(), false))
		_, err := suite.AppVersionController.GetOtherValidSelectors(anID)
		assert.ErrorContains(suite.T(), err, errors.NotFound)
	})
	suite.Run("unsuccessfully for invalid", func() {
		_, err := suite.AppVersionController.GetOtherValidSelectors("something obviously invalid")
		assert.ErrorContains(suite.T(), err, errors.BadRequest)
	})
}

func (suite *appVersionControllerSuite) TestAppVersionEdit() {
	suite.Run("'successfully' but there's no fields", func() {
		db.Truncate(suite.T(), suite.db)
		suite.seedCharts(suite.T())
		suite.seedAppVersions(suite.T())
		appVersions, _ := suite.AppVersionController.ListAllMatching(AppVersion{}, 1)
		anID := fmt.Sprintf("%d", appVersions[0].ID)

		_, err := suite.AppVersionController.Edit(anID, EditableAppVersion{}, auth.GenerateUser(suite.T(), false))
		assert.NoError(suite.T(), err)
	})
}

func (suite *appVersionControllerSuite) TestAppVersionDelete() {
	suite.Run("successfully", func() {
		db.Truncate(suite.T(), suite.db)
		suite.seedCharts(suite.T())
		suite.seedAppVersions(suite.T())
		appVersions, _ := suite.AppVersionController.ListAllMatching(AppVersion{}, 1)
		anID := fmt.Sprintf("%d", appVersions[0].ID)

		deleted, err := suite.AppVersionController.Delete(anID, auth.GenerateUser(suite.T(), false))
		assert.NoError(suite.T(), err)
		assert.Equal(suite.T(), appVersions[0].AppVersion, deleted.AppVersion)
		_, err = suite.AppVersionController.Get(anID)
		assert.ErrorContains(suite.T(), err, errors.NotFound)
	})
}
