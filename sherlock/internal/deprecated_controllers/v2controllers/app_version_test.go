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

func (suite *appVersionControllerSuite) SetupSuite() {
	config.LoadTestConfig()
	suite.db = deprecated_db.ConnectAndConfigureFromTest(suite.T())
	suite.ControllerSet = NewControllerSet(v2models.NewStoreSet(suite.db))
}

func (suite *appVersionControllerSuite) TearDownSuite() {
	deprecated_db.Truncate(suite.T(), suite.db)
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
		Chart:            leonardoChart.Name,
		AppVersion:       "1.2.4",
		GitCommit:        "a1b2",
		GitBranch:        *leonardoChart.AppImageGitMainBranch,
		ParentAppVersion: fmt.Sprintf("%s/%s", leonardoChart.Name, leonardoMain1AppVersion.AppVersion),
	}
	leonardoMain3AppVersion = CreatableAppVersion{
		Chart:            leonardoChart.Name,
		AppVersion:       "1.2.5",
		GitCommit:        "a1b3",
		GitBranch:        *leonardoChart.AppImageGitMainBranch,
		ParentAppVersion: fmt.Sprintf("%s/%s", leonardoChart.Name, leonardoMain2AppVersion.AppVersion),
	}
	leonardoBranch1AppVersion = CreatableAppVersion{
		Chart:            leonardoChart.Name,
		AppVersion:       "1.2.4-a1c1",
		GitCommit:        "a1c1",
		GitBranch:        "branchy-branch",
		ParentAppVersion: fmt.Sprintf("%s/%s", leonardoChart.Name, leonardoMain1AppVersion.AppVersion),
	}
	leonardoBranch2AppVersion = CreatableAppVersion{
		Chart:            leonardoChart.Name,
		AppVersion:       "1.2.4-a1c2",
		GitCommit:        "a1c2",
		GitBranch:        "branchy-branch",
		ParentAppVersion: fmt.Sprintf("%s/%s", leonardoChart.Name, leonardoBranch1AppVersion.AppVersion),
	}
	leonardoBranch3AppVersion = CreatableAppVersion{
		Chart:            leonardoChart.Name,
		AppVersion:       "1.2.4-a1c3",
		GitCommit:        "a1c3",
		GitBranch:        "branchy-branch",
		ParentAppVersion: fmt.Sprintf("%s/%s", leonardoChart.Name, leonardoBranch2AppVersion.AppVersion),
	}
	samMain1AppVersion = CreatableAppVersion{
		Chart:      samChart.Name,
		AppVersion: "0.1.0",
		GitCommit:  "a1b2",
		GitBranch:  *samChart.AppImageGitMainBranch,
	}
	samMain2AppVersion = CreatableAppVersion{
		Chart:            samChart.Name,
		AppVersion:       "0.2.0",
		GitCommit:        "c3d4",
		GitBranch:        *samChart.AppImageGitMainBranch,
		ParentAppVersion: fmt.Sprintf("%s/%s", samChart.Name, samMain1AppVersion.AppVersion),
	}
	yaleMain1AppVersion = CreatableAppVersion{
		Chart:      yaleChart.Name,
		AppVersion: "someversion",
		GitCommit:  "somecommit",
		GitBranch:  "main",
	}
	datarepoMain1AppVerison = CreatableAppVersion{
		Chart:      datarepoChart.Name,
		AppVersion: "1.1.1",
		GitCommit:  "abcd",
		GitBranch:  "develop",
	}
	appVersionSeedList = []CreatableAppVersion{
		leonardoMain1AppVersion,
		leonardoMain2AppVersion,
		leonardoMain3AppVersion,
		leonardoBranch1AppVersion,
		leonardoBranch2AppVersion,
		leonardoBranch3AppVersion,
		samMain1AppVersion,
		samMain2AppVersion,
		yaleMain1AppVersion,
		datarepoMain1AppVerison,
	}
)

func (controllerSet *ControllerSet) seedAppVersions(t *testing.T, db *gorm.DB) {
	for _, creatable := range appVersionSeedList {
		if _, _, err := controllerSet.AppVersionController.Create(creatable, generateUser(t, db, false)); err != nil {
			t.Errorf("error seeding app version %s for chart %s: %w", creatable.AppVersion, creatable.Chart, err)
		}
	}
}

//
// Controller tests
//

func (suite *appVersionControllerSuite) TestAppVersionCreate() {
	suite.Run("can create a new app version", func() {
		deprecated_db.Truncate(suite.T(), suite.db)
		suite.seedCharts(suite.T(), suite.db)

		appVersion, created, err := suite.AppVersionController.Create(leonardoMain1AppVersion, generateUser(suite.T(), suite.db, false))
		assert.NoError(suite.T(), err)
		assert.True(suite.T(), created)
		assert.Equal(suite.T(), leonardoMain1AppVersion.AppVersion, appVersion.AppVersion)
		assert.True(suite.T(), appVersion.ID > 0)

		suite.Run("can accept duplicates", func() {
			secondAppVersion, created, err := suite.AppVersionController.Create(leonardoMain1AppVersion, generateUser(suite.T(), suite.db, false))
			assert.NoError(suite.T(), err)
			assert.False(suite.T(), created)
			assert.Equal(suite.T(), leonardoMain1AppVersion.AppVersion, secondAppVersion.AppVersion)
			assert.True(suite.T(), secondAppVersion.ID > 0)
		})
	})
	suite.Run("validates incoming entries", func() {
		deprecated_db.Truncate(suite.T(), suite.db)

		_, created, err := suite.AppVersionController.Create(CreatableAppVersion{}, generateUser(suite.T(), suite.db, false))
		assert.ErrorContains(suite.T(), err, errors.BadRequest)
		assert.False(suite.T(), created)
	})
	suite.Run("rejects mismatched duplicates", func() {
		deprecated_db.Truncate(suite.T(), suite.db)
		suite.seedCharts(suite.T(), suite.db)
		suite.seedAppVersions(suite.T(), suite.db)

		_, created, err := suite.AppVersionController.Create(CreatableAppVersion{
			Chart:      leonardoChart.Name,
			AppVersion: "1.2.5",
			GitCommit:  "a1b3",
			GitBranch:  *leonardoChart.AppImageGitMainBranch,
			// Mismatched parent
			ParentAppVersion: fmt.Sprintf("%s/%s", leonardoChart.Name, leonardoMain1AppVersion.AppVersion),
		}, generateUser(suite.T(), suite.db, false))
		assert.ErrorContains(suite.T(), err, errors.Conflict)
		assert.False(suite.T(), created)
	})
	suite.Run("accepts bad parents", func() {
		deprecated_db.Truncate(suite.T(), suite.db)
		suite.seedCharts(suite.T(), suite.db)
		suite.seedAppVersions(suite.T(), suite.db)

		appVersion, created, err := suite.AppVersionController.Create(CreatableAppVersion{
			Chart:      datarepoChart.Name,
			AppVersion: "1.1.2",
			// Nonexistent parent
			ParentAppVersion: fmt.Sprintf("%s/%s", datarepoChart.Name, leonardoMain1AppVersion.AppVersion),
		}, generateUser(suite.T(), suite.db, false))
		assert.NoError(suite.T(), err)
		assert.True(suite.T(), created)
		assert.Empty(suite.T(), appVersion.ParentAppVersion)
	})
}

func (suite *appVersionControllerSuite) TestAppVersionListAllMatching() {
	deprecated_db.Truncate(suite.T(), suite.db)
	suite.seedCharts(suite.T(), suite.db)
	suite.seedAppVersions(suite.T(), suite.db)

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

func (suite *appVersionControllerSuite) TestAppVersionGetChildrenPathToParent() {
	deprecated_db.Truncate(suite.T(), suite.db)
	suite.seedCharts(suite.T(), suite.db)
	suite.seedAppVersions(suite.T(), suite.db)

	childSelector := fmt.Sprintf("%s/%s", leonardoChart.Name, leonardoMain3AppVersion.AppVersion)
	parentSelector := fmt.Sprintf("%s/%s", leonardoChart.Name, leonardoMain1AppVersion.AppVersion)

	suite.Run("handles same", func() {
		path, connected, err := suite.AppVersionController.GetChildrenPathToParent(childSelector, childSelector)
		suite.Assert().NoError(err)
		suite.Assert().True(connected)
		suite.Assert().Len(path, 0)
		path, connected, err = suite.AppVersionController.GetChildrenPathToParent(parentSelector, parentSelector)
		suite.Assert().NoError(err)
		suite.Assert().True(connected)
		suite.Assert().Len(path, 0)
	})
	suite.Run("handles different but connected", func() {
		path, connected, err := suite.AppVersionController.GetChildrenPathToParent(childSelector, parentSelector)
		suite.Assert().NoError(err)
		suite.Assert().True(connected)
		suite.Assert().Len(path, 2)
		suite.Assert().Equal(leonardoMain2AppVersion.AppVersion, path[1].AppVersion)
	})
	suite.Run("handles different but disconnected", func() {
		path, connected, err := suite.AppVersionController.GetChildrenPathToParent(parentSelector, childSelector)
		suite.Assert().NoError(err)
		suite.Assert().False(connected)
		suite.Assert().Len(path, 1)
	})
}

func (suite *appVersionControllerSuite) TestAppVersionGet() {
	deprecated_db.Truncate(suite.T(), suite.db)
	suite.seedCharts(suite.T(), suite.db)
	suite.seedAppVersions(suite.T(), suite.db)
	appVersions, _ := suite.AppVersionController.ListAllMatching(AppVersion{}, 1)
	anID := appVersions[0].ID

	suite.Run("successfully", func() {
		byID, err := suite.AppVersionController.Get(fmt.Sprintf("%d", anID))
		assert.NoError(suite.T(), err)
		assert.Equal(suite.T(), appVersions[0].AppVersion, byID.AppVersion)
	})
	suite.Run("unsuccessfully for non-present", func() {
		// Can't predict IDs so we just delete one that we know existed
		_, _ = suite.AppVersionController.Delete(fmt.Sprintf("%d", anID), generateUser(suite.T(), suite.db, false))
		_, err := suite.AppVersionController.Get(fmt.Sprintf("%d", anID))
		assert.ErrorContains(suite.T(), err, errors.NotFound)
	})
	suite.Run("unsuccessfully for invalid selector", func() {
		_, err := suite.AppVersionController.Get("something obviously invalid")
		assert.ErrorContains(suite.T(), err, errors.BadRequest)
	})
}

func (suite *appVersionControllerSuite) TestAppVersionGetOtherValidSelectors() {
	deprecated_db.Truncate(suite.T(), suite.db)
	suite.seedCharts(suite.T(), suite.db)
	suite.seedAppVersions(suite.T(), suite.db)
	appVersions, _ := suite.AppVersionController.ListAllMatching(AppVersion{}, 1)
	anID := fmt.Sprintf("%d", appVersions[0].ID)

	suite.Run("successfully", func() {
		selectors, err := suite.AppVersionController.GetOtherValidSelectors(anID)
		assert.NoError(suite.T(), err)
		assert.Equal(suite.T(), 3, len(selectors))
		assert.Equal(suite.T(), anID, selectors[0])
	})
	suite.Run("unsuccessfully for not found", func() {
		// Can't predict IDs so we just delete one that we know existed
		_, _ = suite.AppVersionController.Delete(anID, generateUser(suite.T(), suite.db, false))
		_, err := suite.AppVersionController.GetOtherValidSelectors(anID)
		assert.ErrorContains(suite.T(), err, errors.NotFound)
	})
	suite.Run("unsuccessfully for invalid", func() {
		_, err := suite.AppVersionController.GetOtherValidSelectors("something obviously invalid")
		assert.ErrorContains(suite.T(), err, errors.BadRequest)
	})
}

func (suite *appVersionControllerSuite) TestAppVersionEdit() {
	suite.Run("successfully", func() {
		deprecated_db.Truncate(suite.T(), suite.db)
		suite.seedCharts(suite.T(), suite.db)
		suite.seedAppVersions(suite.T(), suite.db)
		appVersions, _ := suite.AppVersionController.ListAllMatching(AppVersion{}, 1)
		anID := fmt.Sprintf("%d", appVersions[0].ID)

		response, err := suite.AppVersionController.Edit(anID, EditableAppVersion{Description: "blah"}, generateUser(suite.T(), suite.db, false))
		assert.NoError(suite.T(), err)
		assert.Equal(suite.T(), "blah", response.Description)
	})
}

func (suite *appVersionControllerSuite) TestAppVersionDelete() {
	suite.Run("successfully", func() {
		deprecated_db.Truncate(suite.T(), suite.db)
		suite.seedCharts(suite.T(), suite.db)
		suite.seedAppVersions(suite.T(), suite.db)
		appVersions, _ := suite.AppVersionController.ListAllMatching(AppVersion{}, 1)
		anID := fmt.Sprintf("%d", appVersions[0].ID)

		deleted, err := suite.AppVersionController.Delete(anID, generateUser(suite.T(), suite.db, false))
		assert.NoError(suite.T(), err)
		assert.Equal(suite.T(), appVersions[0].AppVersion, deleted.AppVersion)
		_, err = suite.AppVersionController.Get(anID)
		assert.ErrorContains(suite.T(), err, errors.NotFound)
	})
}
