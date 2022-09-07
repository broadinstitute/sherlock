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

func TestEnvironmentControllerSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping functional test")
	}
	suite.Run(t, new(environmentControllerSuite))
}

type environmentControllerSuite struct {
	suite.Suite
	*ControllerSet
	db *gorm.DB
}

func (suite *environmentControllerSuite) SetupTest() {
	config.LoadTestConfig(suite.T())
	suite.db = db.ConnectAndConfigureFromTest(suite.T())
	suite.db.Begin()
	suite.ControllerSet = NewControllerSet(v2models.NewStoreSet(suite.db))
}

func (suite *environmentControllerSuite) TearDownTest() {
	suite.db.Rollback()
}

//
// Controller seeding
//

var (
	terraDevEnvironment = CreatableEnvironment{
		Name:      "terra-dev",
		Base:      "live",
		Lifecycle: "static",
		EditableEnvironment: EditableEnvironment{
			DefaultCluster:   &terraDevCluster.Name,
			DefaultNamespace: testutils.PointerTo("terra-dev"),
			Owner:            testutils.PointerTo("dsp-devops@broadinstitute.org"),
		},
	}
	terraProdEnvironment = CreatableEnvironment{
		Name:      "terra-prod",
		Base:      "live",
		Lifecycle: "static",
		EditableEnvironment: EditableEnvironment{
			DefaultCluster:      &terraProdCluster.Name,
			Owner:               testutils.PointerTo("dsp-devops@broadinstitute.org"),
			RequiresSuitability: testutils.PointerTo(true),
		},
	}
	swatomationEnvironment = CreatableEnvironment{
		Name:      "swatomation",
		Base:      "bee",
		Lifecycle: "template",
		EditableEnvironment: EditableEnvironment{
			DefaultCluster: &terraQaBeesCluster.Name,
			Owner:          testutils.PointerTo("dsp-devops@broadinstitute.org"),
		},
	}
	dynamicSwatomationEnvironment = CreatableEnvironment{
		Name:                "swatomation-instance-one",
		TemplateEnvironment: swatomationEnvironment.Name,
	}
	environmentSeedList = []CreatableEnvironment{terraDevEnvironment, terraProdEnvironment, swatomationEnvironment, dynamicSwatomationEnvironment}
)

func (controllerSet *ControllerSet) seedEnvironments(t *testing.T) {
	for _, creatable := range environmentSeedList {
		if _, _, err := controllerSet.EnvironmentController.Create(creatable, auth.GenerateUser(t, true)); err != nil {
			t.Errorf("error seeding environment %s: %v", creatable.Name, err)
		}
	}
}

//
// Controller tests
//

func (suite *environmentControllerSuite) TestEnvironmentCreate() {
	suite.Run("can create a new environment", func() {
		db.Truncate(suite.T(), suite.db)
		suite.seedClusters(suite.T())

		suite.Run("static", func() {
			env, created, err := suite.EnvironmentController.Create(terraDevEnvironment, auth.GenerateUser(suite.T(), false))
			assert.NoError(suite.T(), err)
			assert.True(suite.T(), created)
			assert.Equal(suite.T(), terraDevEnvironment.Name, env.Name)
			assert.True(suite.T(), env.ID > 0)
			suite.Run("default non-suitable", func() {
				assert.False(suite.T(), *env.RequiresSuitability)
			})
			suite.Run("default the default namespace to environment name", func() {
				assert.Equal(suite.T(), terraDevEnvironment.Name, *env.DefaultNamespace)
			})
		})
		suite.Run("template", func() {
			env, created, err := suite.EnvironmentController.Create(swatomationEnvironment, auth.GenerateUser(suite.T(), false))
			assert.NoError(suite.T(), err)
			assert.True(suite.T(), created)
			assert.Equal(suite.T(), swatomationEnvironment.Name, env.Name)
			assert.True(suite.T(), env.ID > 0)
			suite.Run("default non-suitable", func() {
				assert.False(suite.T(), *env.RequiresSuitability)
			})
		})
		suite.Run("dynamic", func() {
			user := auth.GenerateUser(suite.T(), false)
			env, created, err := suite.EnvironmentController.Create(dynamicSwatomationEnvironment, auth.GenerateUser(suite.T(), false))
			assert.NoError(suite.T(), err)
			assert.True(suite.T(), created)
			assert.Equal(suite.T(), dynamicSwatomationEnvironment.Name, env.Name)
			assert.True(suite.T(), env.ID > 0)
			suite.Run("references template name in defaults", func() {
				assert.Equal(suite.T(), swatomationEnvironment.Name, env.TemplateEnvironment)
				assert.Equal(suite.T(), swatomationEnvironment.Name, env.ValuesName)
			})
			suite.Run("base of template", func() {
				assert.Equal(suite.T(), swatomationEnvironment.Base, env.Base)
			})
			suite.Run("cluster of template", func() {
				assert.Equal(suite.T(), swatomationEnvironment.DefaultCluster, env.DefaultCluster)
			})
			suite.Run("fills owner", func() {
				assert.Equal(suite.T(), user.AuthenticatedEmail, *env.Owner)
			})
			suite.Run("namespace of name", func() {
				assert.Equal(suite.T(), env.Name, *env.DefaultNamespace)
			})
		})
	})
	suite.Run("won't create duplicates", func() {
		db.Truncate(suite.T(), suite.db)
		suite.seedClusters(suite.T())

		env, created, err := suite.EnvironmentController.Create(terraDevEnvironment, auth.GenerateUser(suite.T(), false))
		assert.NoError(suite.T(), err)
		assert.True(suite.T(), created)
		assert.True(suite.T(), env.ID > 0)
		_, created, err = suite.EnvironmentController.Create(terraDevEnvironment, auth.GenerateUser(suite.T(), false))
		assert.ErrorContains(suite.T(), err, errors.Conflict)
		assert.False(suite.T(), created)
	})
	suite.Run("validates incoming entries", func() {
		db.Truncate(suite.T(), suite.db)

		_, created, err := suite.EnvironmentController.Create(CreatableEnvironment{}, auth.GenerateUser(suite.T(), false))
		assert.ErrorContains(suite.T(), err, errors.BadRequest)
		assert.False(suite.T(), created)
	})
	suite.Run("checks suitability", func() {
		db.Truncate(suite.T(), suite.db)
		suite.seedClusters(suite.T())

		assert.True(suite.T(), *terraProdEnvironment.RequiresSuitability)
		suite.Run("blocks suitable creation for non-suitable", func() {
			_, created, err := suite.EnvironmentController.Create(terraProdEnvironment, auth.GenerateUser(suite.T(), false))
			assert.ErrorContains(suite.T(), err, errors.Forbidden)
			assert.False(suite.T(), created)
		})
		suite.Run("allows suitable creation for suitable", func() {
			env, created, err := suite.EnvironmentController.Create(terraProdEnvironment, auth.GenerateUser(suite.T(), true))
			assert.NoError(suite.T(), err)
			assert.True(suite.T(), created)
			assert.True(suite.T(), env.ID > 0)
		})
	})
	suite.Run("dynamic defaulting respects template edits", func() {
		db.Truncate(suite.T(), suite.db)
		suite.seedClusters(suite.T())
		suite.seedEnvironments(suite.T())

		swat, err := suite.EnvironmentController.Get(swatomationEnvironment.Name)
		assert.NoError(suite.T(), err)
		assert.False(suite.T(), *swat.RequiresSuitability)

		env1, created, err := suite.EnvironmentController.Create(CreatableEnvironment{TemplateEnvironment: swatomationEnvironment.Name}, auth.GenerateUser(suite.T(), false))
		assert.NoError(suite.T(), err)
		assert.True(suite.T(), created)
		assert.False(suite.T(), *env1.RequiresSuitability)

		swat, err = suite.EnvironmentController.Edit(swatomationEnvironment.Name, EditableEnvironment{RequiresSuitability: testutils.PointerTo(true)}, auth.GenerateUser(suite.T(), true))
		assert.NoError(suite.T(), err)
		assert.True(suite.T(), *swat.RequiresSuitability)
		assert.False(suite.T(), *env1.RequiresSuitability)

		env2, created, err := suite.EnvironmentController.Create(CreatableEnvironment{TemplateEnvironment: swatomationEnvironment.Name}, auth.GenerateUser(suite.T(), true))
		assert.NoError(suite.T(), err)
		assert.True(suite.T(), created)
		assert.True(suite.T(), *env2.RequiresSuitability)
		assert.False(suite.T(), *env1.RequiresSuitability)
	})
	suite.Run("copies template chart releases", func() {
		db.Truncate(suite.T(), suite.db)
		suite.seedClusters(suite.T())
		suite.seedEnvironments(suite.T())
		suite.seedCharts(suite.T())
		suite.seedChartReleases(suite.T())

		swatReleases, err := suite.ChartReleaseController.ListAllMatching(
			ChartRelease{CreatableChartRelease: CreatableChartRelease{Environment: swatomationEnvironment.Name}}, 0)
		assert.NoError(suite.T(), err)
		assert.True(suite.T(), len(swatReleases) > 0)

		environment, created, err := suite.EnvironmentController.Create(CreatableEnvironment{TemplateEnvironment: swatomationEnvironment.Name}, auth.GenerateUser(suite.T(), false))
		assert.NoError(suite.T(), err)
		assert.True(suite.T(), created)

		environmentReleases, err := suite.ChartReleaseController.ListAllMatching(
			ChartRelease{CreatableChartRelease: CreatableChartRelease{Environment: environment.Name}}, 0)
		assert.NoError(suite.T(), err)

		for _, swatRelease := range swatReleases {
			found := false
			for _, envRelease := range environmentReleases {
				if swatRelease.Chart == envRelease.Chart {
					found = true
					// The template gets this field set based on the Chart's app main branch, dynamic should copy
					assert.Equal(suite.T(), swatRelease.TargetAppVersionBranch, envRelease.TargetAppVersionBranch)
					assert.True(suite.T(), envRelease.ID > 0)
				}
			}
			assert.True(suite.T(), found)
		}
	})
}

func (suite *environmentControllerSuite) TestEnvironmentListAllMatching() {
	db.Truncate(suite.T(), suite.db)
	suite.seedClusters(suite.T())
	suite.seedEnvironments(suite.T())

	suite.Run("lists all environments", func() {
		matching, err := suite.EnvironmentController.ListAllMatching(Environment{}, 0)
		assert.NoError(suite.T(), err)
		assert.Equal(suite.T(), len(environmentSeedList), len(matching))
		suite.Run("orders by latest updated", func() {
			latestUpdated := matching[0].UpdatedAt
			for _, environment := range matching {
				assert.GreaterOrEqual(suite.T(), latestUpdated, environment.UpdatedAt)
			}
		})
	})
	suite.Run("limits", func() {
		limit := 2
		matching, err := suite.EnvironmentController.ListAllMatching(Environment{}, limit)
		assert.NoError(suite.T(), err)
		assert.Equal(suite.T(), limit, len(matching))
	})
	suite.Run("filters exactly", func() {
		matching, err := suite.EnvironmentController.ListAllMatching(Environment{CreatableEnvironment: terraDevEnvironment}, 0)
		assert.NoError(suite.T(), err)
		assert.Equal(suite.T(), 1, len(matching))
		assert.Equal(suite.T(), terraDevEnvironment.Name, matching[0].Name)
	})
	suite.Run("filters multiple", func() {
		matching, err := suite.EnvironmentController.ListAllMatching(
			Environment{CreatableEnvironment: CreatableEnvironment{Base: "live"}}, 0)
		assert.NoError(suite.T(), err)
		assert.True(suite.T(), len(matching) > 1)
		for _, environment := range matching {
			assert.Equal(suite.T(), "live", environment.Base)
		}
	})
	suite.Run("none is an empty list, not null", func() {
		matching, err := suite.EnvironmentController.ListAllMatching(
			Environment{CreatableEnvironment: CreatableEnvironment{Name: "blah"}}, 0)
		assert.NoError(suite.T(), err)
		assert.NotNil(suite.T(), matching)
		assert.Empty(suite.T(), matching)
	})
}

func (suite *environmentControllerSuite) TestEnvironmentGet() {
	db.Truncate(suite.T(), suite.db)
	suite.seedClusters(suite.T())
	suite.seedEnvironments(suite.T())

	suite.Run("successfully", func() {
		byName, err := suite.EnvironmentController.Get(terraDevEnvironment.Name)
		assert.NoError(suite.T(), err)
		assert.True(suite.T(), byName.ID > 0)
		byID, err := suite.EnvironmentController.Get(fmt.Sprintf("%d", byName.ID))
		assert.NoError(suite.T(), err)
		assert.Equal(suite.T(), terraDevEnvironment.Name, byID.Name)
	})
	suite.Run("unsuccessfully for non-present", func() {
		_, err := suite.EnvironmentController.Get("foobar")
		assert.ErrorContains(suite.T(), err, errors.NotFound)
	})
	suite.Run("unsuccessfully for invalid selector", func() {
		_, err := suite.EnvironmentController.Get("something obviously invalid")
		assert.ErrorContains(suite.T(), err, errors.BadRequest)
	})
}

func (suite *environmentControllerSuite) TestEnvironmentGetOtherValidSelectors() {
	db.Truncate(suite.T(), suite.db)
	suite.seedClusters(suite.T())
	suite.seedEnvironments(suite.T())

	suite.Run("successfully", func() {
		selectors, err := suite.EnvironmentController.GetOtherValidSelectors(terraDevEnvironment.Name)
		assert.NoError(suite.T(), err)
		assert.Equal(suite.T(), 2, len(selectors))
		assert.Equal(suite.T(), terraDevEnvironment.Name, selectors[0])
	})
	suite.Run("unsuccessfully for not found", func() {
		_, err := suite.EnvironmentController.GetOtherValidSelectors("foobar")
		assert.ErrorContains(suite.T(), err, errors.NotFound)
	})
	suite.Run("unsuccessfully for invalid", func() {
		_, err := suite.EnvironmentController.GetOtherValidSelectors("something obviously invalid")
		assert.ErrorContains(suite.T(), err, errors.BadRequest)
	})
}

func (suite *environmentControllerSuite) TestEnvironmentEdit() {
	suite.Run("successfully", func() {
		db.Truncate(suite.T(), suite.db)
		suite.seedClusters(suite.T())
		suite.seedEnvironments(suite.T())

		before, err := suite.EnvironmentController.Get(terraDevEnvironment.Name)
		assert.NoError(suite.T(), err)
		assert.Equal(suite.T(), terraDevEnvironment.DefaultNamespace, before.DefaultNamespace)
		newNamespace := testutils.PointerTo("new")
		edited, err := suite.EnvironmentController.Edit(terraDevEnvironment.Name, EditableEnvironment{DefaultNamespace: newNamespace}, auth.GenerateUser(suite.T(), false))
		assert.NoError(suite.T(), err)
		assert.Equal(suite.T(), newNamespace, edited.DefaultNamespace)
		after, err := suite.EnvironmentController.Get(terraDevEnvironment.Name)
		assert.NoError(suite.T(), err)
		assert.Equal(suite.T(), newNamespace, after.DefaultNamespace)
	})
	suite.Run("edit to suitable environment", func() {
		db.Truncate(suite.T(), suite.db)
		suite.seedClusters(suite.T())
		suite.seedEnvironments(suite.T())
		newNamespace := testutils.PointerTo("new")

		suite.Run("unsuccessfully if not suitable", func() {
			_, err := suite.EnvironmentController.Edit(terraProdEnvironment.Name, EditableEnvironment{DefaultNamespace: newNamespace}, auth.GenerateUser(suite.T(), false))
			assert.ErrorContains(suite.T(), err, errors.Forbidden)
			notEdited, err := suite.EnvironmentController.Get(terraProdEnvironment.Name)
			assert.NoError(suite.T(), err)
			assert.NotEqual(suite.T(), newNamespace, notEdited.DefaultNamespace)
		})
		suite.Run("successfully if suitable", func() {
			edited, err := suite.EnvironmentController.Edit(terraProdEnvironment.Name, EditableEnvironment{DefaultNamespace: newNamespace}, auth.GenerateUser(suite.T(), true))
			assert.NoError(suite.T(), err)
			assert.Equal(suite.T(), newNamespace, edited.DefaultNamespace)
		})
	})
	suite.Run("edit that would make environment suitable", func() {
		db.Truncate(suite.T(), suite.db)
		suite.seedClusters(suite.T())
		suite.seedEnvironments(suite.T())

		suite.Run("unsuccessfully if not suitable", func() {
			_, err := suite.EnvironmentController.Edit(terraDevEnvironment.Name, EditableEnvironment{RequiresSuitability: testutils.PointerTo(true)}, auth.GenerateUser(suite.T(), false))
			assert.ErrorContains(suite.T(), err, errors.Forbidden)
			notEdited, err := suite.EnvironmentController.Get(terraDevEnvironment.Name)
			assert.NoError(suite.T(), err)
			assert.False(suite.T(), *notEdited.RequiresSuitability)
		})
		suite.Run("successfully if suitable", func() {
			edited, err := suite.EnvironmentController.Edit(terraDevEnvironment.Name, EditableEnvironment{RequiresSuitability: testutils.PointerTo(true)}, auth.GenerateUser(suite.T(), true))
			assert.NoError(suite.T(), err)
			assert.True(suite.T(), *edited.RequiresSuitability)
		})
	})
	suite.Run("unsuccessfully if invalid", func() {
		db.Truncate(suite.T(), suite.db)
		suite.seedClusters(suite.T())
		suite.seedEnvironments(suite.T())

		_, err := suite.EnvironmentController.Edit(terraDevEnvironment.Name, EditableEnvironment{DefaultNamespace: testutils.PointerTo("")}, auth.GenerateUser(suite.T(), false))
		assert.ErrorContains(suite.T(), err, errors.BadRequest)
	})
}

func (suite *environmentControllerSuite) TestEnvironmentDelete() {
	suite.Run("successfully", func() {
		db.Truncate(suite.T(), suite.db)
		suite.seedClusters(suite.T())
		suite.seedEnvironments(suite.T())

		deleted, err := suite.EnvironmentController.Delete(terraDevEnvironment.Name, auth.GenerateUser(suite.T(), false))
		assert.NoError(suite.T(), err)
		assert.Equal(suite.T(), terraDevEnvironment.Name, deleted.Name)
		_, err = suite.EnvironmentController.Get(terraDevEnvironment.Name)
		assert.ErrorContains(suite.T(), err, errors.NotFound)
		suite.Run("sql constraints ignore soft deletion", func() {
			_, created, err := suite.EnvironmentController.Create(terraDevEnvironment, auth.GenerateUser(suite.T(), false))
			assert.ErrorContains(suite.T(), err, errors.BadRequest)
			assert.ErrorContains(suite.T(), err, "Contact DevOps")
			assert.False(suite.T(), created)
		})
	})
	suite.Run("delete suitable environment", func() {
		db.Truncate(suite.T(), suite.db)
		suite.seedClusters(suite.T())
		suite.seedEnvironments(suite.T())

		suite.Run("unsuccessfully if not suitable", func() {
			_, err := suite.EnvironmentController.Delete(terraProdEnvironment.Name, auth.GenerateUser(suite.T(), false))
			assert.ErrorContains(suite.T(), err, errors.Forbidden)
		})
		suite.Run("successfully if suitable", func() {
			deleted, err := suite.EnvironmentController.Delete(terraProdEnvironment.Name, auth.GenerateUser(suite.T(), true))
			assert.NoError(suite.T(), err)
			assert.Equal(suite.T(), terraProdEnvironment.Name, deleted.Name)
		})
	})
}
