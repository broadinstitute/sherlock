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

func TestPagerdutyIntegrationControllerSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping functional test")
	}
	suite.Run(t, new(pagerdutyIntegrationControllerSuite))
}

type pagerdutyIntegrationControllerSuite struct {
	suite.Suite
	*ControllerSet
	db *gorm.DB
}

func (suite *pagerdutyIntegrationControllerSuite) SetupSuite() {
	config.LoadTestConfig()
	suite.db = deprecated_db.ConnectAndConfigureFromTest(suite.T())
	suite.ControllerSet = NewControllerSet(v2models.NewStoreSet(suite.db))
}

func (suite *pagerdutyIntegrationControllerSuite) TearDownSuite() {
	deprecated_db.Truncate(suite.T(), suite.db)
}

//
// Controller seeding
//

var (
	pagerdutyIntegration1 = CreatablePagerdutyIntegration{
		PagerdutyID: "abcd1234",
		EditablePagerdutyIntegration: EditablePagerdutyIntegration{
			Name: utils.PointerTo("ABC service"),
			Key:  utils.PointerTo("a1b2c3d4"),
			Type: utils.PointerTo("service"),
		},
	}
	pagerdutyIntegration2 = CreatablePagerdutyIntegration{
		PagerdutyID: "abcde12345",
		EditablePagerdutyIntegration: EditablePagerdutyIntegration{
			Name: utils.PointerTo("ABC service 2"),
			Key:  utils.PointerTo("a1b2c3d4e5"),
			Type: utils.PointerTo("service"),
		},
	}
	pagerdutyIntgrationSeedList = []CreatablePagerdutyIntegration{
		pagerdutyIntegration1,
		pagerdutyIntegration2,
	}
)

func (controllerSet *ControllerSet) seedPagerdutyIntegrations(t *testing.T, db *gorm.DB) {
	for _, creatable := range pagerdutyIntgrationSeedList {
		if _, _, err := controllerSet.PagerdutyIntegrationController.Create(creatable, generateUser(t, db, true)); err != nil {
			t.Errorf("error seeding pagerduty integration %s: %w", creatable.PagerdutyID, err)
		}
	}
}

//
// Controller tests
//

func (suite *pagerdutyIntegrationControllerSuite) TestPagerdutyIntegrationCreate() {
	suite.Run("can create new", func() {
		deprecated_db.Truncate(suite.T(), suite.db)

		pi, created, err := suite.PagerdutyIntegrationController.Create(pagerdutyIntegration1, generateUser(suite.T(), suite.db, true))
		assert.NoError(suite.T(), err)
		assert.True(suite.T(), created)
		assert.Equal(suite.T(), pagerdutyIntegration1.Name, pi.Name)
		assert.True(suite.T(), pi.ID > 0)
	})
	suite.Run("checks suitability", func() {
		deprecated_db.Truncate(suite.T(), suite.db)

		_, created, err := suite.PagerdutyIntegrationController.Create(pagerdutyIntegration1, generateUser(suite.T(), suite.db, false))
		assert.ErrorContains(suite.T(), err, errors.Forbidden)
		assert.False(suite.T(), created)
	})
}

func (suite *pagerdutyIntegrationControllerSuite) TestPagerdutyIntegrationListAllMatching() {
	deprecated_db.Truncate(suite.T(), suite.db)
	suite.seedPagerdutyIntegrations(suite.T(), suite.db)

	suite.Run("lists all", func() {
		matching, err := suite.PagerdutyIntegrationController.ListAllMatching(PagerdutyIntegration{}, 0)
		assert.NoError(suite.T(), err)
		assert.Equal(suite.T(), len(pagerdutyIntgrationSeedList), len(matching))
		suite.Run("orders by latest updated", func() {
			latestUpdated := matching[0].UpdatedAt
			for _, match := range matching {
				assert.GreaterOrEqual(suite.T(), latestUpdated, match.UpdatedAt)
			}
		})
	})
	suite.Run("limits", func() {
		limit := 1
		matching, err := suite.PagerdutyIntegrationController.ListAllMatching(PagerdutyIntegration{}, limit)
		assert.NoError(suite.T(), err)
		assert.Equal(suite.T(), limit, len(matching))
	})
	suite.Run("filters exactly", func() {
		matching, err := suite.PagerdutyIntegrationController.ListAllMatching(PagerdutyIntegration{PagerdutyID: pagerdutyIntegration1.PagerdutyID}, 0)
		assert.NoError(suite.T(), err)
		assert.Equal(suite.T(), 1, len(matching))
		assert.Equal(suite.T(), pagerdutyIntegration1.PagerdutyID, matching[0].PagerdutyID)
	})
	suite.Run("filters multiple", func() {
		matching, err := suite.PagerdutyIntegrationController.ListAllMatching(
			PagerdutyIntegration{Type: utils.PointerTo("service")}, 0)
		assert.NoError(suite.T(), err)
		assert.True(suite.T(), len(matching) > 1)
		for _, match := range matching {
			assert.Equal(suite.T(), utils.PointerTo("service"), match.Type)
		}
	})
	suite.Run("none is an empty list, not null", func() {
		matching, err := suite.PagerdutyIntegrationController.ListAllMatching(
			PagerdutyIntegration{PagerdutyID: "foo"}, 0)
		assert.NoError(suite.T(), err)
		assert.NotNil(suite.T(), matching)
		assert.Empty(suite.T(), matching)
	})
}

func (suite *pagerdutyIntegrationControllerSuite) TestPagerdutyIntegrationGet() {
	deprecated_db.Truncate(suite.T(), suite.db)
	suite.seedPagerdutyIntegrations(suite.T(), suite.db)

	suite.Run("successfully", func() {
		byPagerdutyID, err := suite.PagerdutyIntegrationController.Get(fmt.Sprintf("pd-id/%s", pagerdutyIntegration1.PagerdutyID))
		assert.NoError(suite.T(), err)
		assert.True(suite.T(), byPagerdutyID.ID > 0)
		byID, err := suite.PagerdutyIntegrationController.Get(fmt.Sprintf("%d", byPagerdutyID.ID))
		assert.NoError(suite.T(), err)
		assert.Equal(suite.T(), pagerdutyIntegration1.Name, byID.Name)
	})
	suite.Run("unsuccessfully for non-present", func() {
		_, err := suite.PagerdutyIntegrationController.Get("pd-id/abc")
		assert.ErrorContains(suite.T(), err, errors.NotFound)
	})
	suite.Run("unsuccessfully for invalid selector", func() {
		_, err := suite.PagerdutyIntegrationController.Get("something obviously invalid")
		assert.ErrorContains(suite.T(), err, errors.BadRequest)
	})
}

func (suite *pagerdutyIntegrationControllerSuite) TestPagerdutyIntegrationGetOtherValidSelectors() {
	deprecated_db.Truncate(suite.T(), suite.db)
	suite.seedPagerdutyIntegrations(suite.T(), suite.db)

	suite.Run("successfully", func() {
		selectors, err := suite.PagerdutyIntegrationController.GetOtherValidSelectors(fmt.Sprintf("pd-id/%s", pagerdutyIntegration1.PagerdutyID))
		assert.NoError(suite.T(), err)
		assert.Equal(suite.T(), 2, len(selectors))
		assert.Equal(suite.T(), fmt.Sprintf("pd-id/%s", pagerdutyIntegration1.PagerdutyID), selectors[1])
	})
	suite.Run("unsuccessfully for not found", func() {
		_, err := suite.PagerdutyIntegrationController.GetOtherValidSelectors("pd-id/abc")
		assert.ErrorContains(suite.T(), err, errors.NotFound)
	})
	suite.Run("unsuccessfully for invalid", func() {
		_, err := suite.PagerdutyIntegrationController.GetOtherValidSelectors("something obviously invalid")
		assert.ErrorContains(suite.T(), err, errors.BadRequest)
	})
}

func (suite *pagerdutyIntegrationControllerSuite) TestPagerdutyIntegrationEdit() {
	suite.Run("successfully", func() {
		deprecated_db.Truncate(suite.T(), suite.db)
		suite.seedPagerdutyIntegrations(suite.T(), suite.db)

		before, err := suite.PagerdutyIntegrationController.Get(fmt.Sprintf("pd-id/%s", pagerdutyIntegration1.PagerdutyID))
		assert.NoError(suite.T(), err)
		assert.Equal(suite.T(), pagerdutyIntegration1.Name, before.Name)
		newName := utils.PointerTo("new")
		edited, err := suite.PagerdutyIntegrationController.Edit(fmt.Sprintf("pd-id/%s", pagerdutyIntegration1.PagerdutyID), EditablePagerdutyIntegration{Name: newName}, generateUser(suite.T(), suite.db, true))
		assert.NoError(suite.T(), err)
		assert.Equal(suite.T(), newName, edited.Name)
		after, err := suite.PagerdutyIntegrationController.Get(fmt.Sprintf("pd-id/%s", pagerdutyIntegration1.PagerdutyID))
		assert.NoError(suite.T(), err)
		assert.Equal(suite.T(), newName, after.Name)
	})
	suite.Run("unsuccessfully if invalid", func() {
		deprecated_db.Truncate(suite.T(), suite.db)
		suite.seedPagerdutyIntegrations(suite.T(), suite.db)

		_, err := suite.PagerdutyIntegrationController.Edit(fmt.Sprintf("pd-id/%s", pagerdutyIntegration1.PagerdutyID), EditablePagerdutyIntegration{Name: utils.PointerTo("")}, generateUser(suite.T(), suite.db, true))
		assert.ErrorContains(suite.T(), err, errors.BadRequest)
	})
	suite.Run("unsuccessfully if forbidden", func() {
		deprecated_db.Truncate(suite.T(), suite.db)
		suite.seedPagerdutyIntegrations(suite.T(), suite.db)

		_, err := suite.PagerdutyIntegrationController.Edit(fmt.Sprintf("pd-id/%s", pagerdutyIntegration1.PagerdutyID), EditablePagerdutyIntegration{Name: utils.PointerTo("foo")}, generateUser(suite.T(), suite.db, false))
		assert.ErrorContains(suite.T(), err, errors.Forbidden)
	})
}

func (suite *pagerdutyIntegrationControllerSuite) TestPagerdutyIntegrationUpsert() {
	suite.Run("successfully", func() {
		deprecated_db.Truncate(suite.T(), suite.db)

		matches, err := suite.PagerdutyIntegrationController.ListAllMatching(PagerdutyIntegration{}, 0)
		assert.NoError(suite.T(), err)
		assert.Empty(suite.T(), matches)
		put, created, err := suite.PagerdutyIntegrationController.Upsert(fmt.Sprintf("pd-id/%s", pagerdutyIntegration1.PagerdutyID), pagerdutyIntegration1, pagerdutyIntegration1.EditablePagerdutyIntegration, generateUser(suite.T(), suite.db, true))
		assert.NoError(suite.T(), err)
		assert.True(suite.T(), created)
		assert.Equal(suite.T(), pagerdutyIntegration1.Name, put.Name)
		matches, err = suite.PagerdutyIntegrationController.ListAllMatching(PagerdutyIntegration{}, 0)
		assert.NoError(suite.T(), err)
		assert.Len(suite.T(), matches, 1)
		put, created, err = suite.PagerdutyIntegrationController.Upsert(fmt.Sprintf("pd-id/%s", pagerdutyIntegration1.PagerdutyID), pagerdutyIntegration1, pagerdutyIntegration1.EditablePagerdutyIntegration, generateUser(suite.T(), suite.db, true))
		assert.NoError(suite.T(), err)
		assert.False(suite.T(), created)
		assert.Equal(suite.T(), pagerdutyIntegration1.Name, put.Name)
		matches, err = suite.PagerdutyIntegrationController.ListAllMatching(PagerdutyIntegration{}, 0)
		assert.NoError(suite.T(), err)
		assert.Len(suite.T(), matches, 1)
		newName := utils.PointerTo("new")
		edited := CreatablePagerdutyIntegration{
			PagerdutyID: pagerdutyIntegration1.PagerdutyID,
			EditablePagerdutyIntegration: EditablePagerdutyIntegration{
				Name: newName,
				Key:  pagerdutyIntegration1.Key,
				Type: pagerdutyIntegration1.Type,
			},
		}
		put, created, err = suite.PagerdutyIntegrationController.Upsert(fmt.Sprintf("pd-id/%s", pagerdutyIntegration1.PagerdutyID), edited, edited.EditablePagerdutyIntegration, generateUser(suite.T(), suite.db, true))
		assert.NoError(suite.T(), err)
		assert.False(suite.T(), created)
		assert.Equal(suite.T(), newName, put.Name)
		suite.Run("edits without all fields being set", func() {
			newName = utils.PointerTo("new again")
			editedAgain := CreatablePagerdutyIntegration{
				PagerdutyID: pagerdutyIntegration1.PagerdutyID,
				EditablePagerdutyIntegration: EditablePagerdutyIntegration{
					Name: newName,
				},
			}
			put, created, err = suite.PagerdutyIntegrationController.Upsert(fmt.Sprintf("pd-id/%s", pagerdutyIntegration1.PagerdutyID), editedAgain, editedAgain.EditablePagerdutyIntegration, generateUser(suite.T(), suite.db, true))
			assert.NoError(suite.T(), err)
			assert.False(suite.T(), created)
			assert.Equal(suite.T(), newName, put.Name)
			assert.Equal(suite.T(), pagerdutyIntegration1.Type, put.Type)
		})
	})
	suite.Run("unsuccessfully if forbidden", func() {
		deprecated_db.Truncate(suite.T(), suite.db)

		matches, err := suite.PagerdutyIntegrationController.ListAllMatching(PagerdutyIntegration{}, 0)
		assert.NoError(suite.T(), err)
		assert.Empty(suite.T(), matches)
		_, created, err := suite.PagerdutyIntegrationController.Upsert(fmt.Sprintf("pd-id/%s", pagerdutyIntegration1.PagerdutyID), pagerdutyIntegration1, pagerdutyIntegration1.EditablePagerdutyIntegration, generateUser(suite.T(), suite.db, false))
		assert.ErrorContains(suite.T(), err, errors.Forbidden)
		assert.False(suite.T(), created)
	})
}

func (suite *pagerdutyIntegrationControllerSuite) TestPagerdutyIntegrationDelete() {
	suite.Run("successfully", func() {
		deprecated_db.Truncate(suite.T(), suite.db)
		suite.seedPagerdutyIntegrations(suite.T(), suite.db)

		deleted, err := suite.PagerdutyIntegrationController.Delete(fmt.Sprintf("pd-id/%s", pagerdutyIntegration1.PagerdutyID), generateUser(suite.T(), suite.db, true))
		assert.NoError(suite.T(), err)
		assert.Equal(suite.T(), pagerdutyIntegration1.PagerdutyID, deleted.PagerdutyID)
		_, err = suite.PagerdutyIntegrationController.Get(fmt.Sprintf("pd-id/%s", pagerdutyIntegration1.PagerdutyID))
		assert.ErrorContains(suite.T(), err, errors.NotFound)
	})
	suite.Run("unsuccessfully if forbidden", func() {
		deprecated_db.Truncate(suite.T(), suite.db)
		suite.seedPagerdutyIntegrations(suite.T(), suite.db)

		_, err := suite.PagerdutyIntegrationController.Delete(fmt.Sprintf("pd-id/%s", pagerdutyIntegration1.PagerdutyID), generateUser(suite.T(), suite.db, false))
		assert.ErrorContains(suite.T(), err, errors.Forbidden)
	})
}
