package v2controllers

import (
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/deprecated_db"
	"github.com/broadinstitute/sherlock/sherlock/internal/deprecated_models/v2models"
	"github.com/google/uuid"
	"strings"
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

func (suite *environmentControllerSuite) SetupSuite() {
	config.LoadTestConfig()
	suite.db = deprecated_db.ConnectAndConfigureFromTest(suite.T())
	suite.ControllerSet = NewControllerSet(v2models.NewStoreSet(suite.db))
}

func (suite *environmentControllerSuite) TearDownSuite() {
	deprecated_db.Truncate(suite.T(), suite.db)
}

//
// Controller seeding
//

var (
	terraDevEnvironment = CreatableEnvironment{
		Name:             "terra-dev",
		Base:             "live",
		Lifecycle:        "static",
		DefaultNamespace: "terra-dev",
		EditableEnvironment: EditableEnvironment{
			DefaultCluster: &terraDevCluster.Name,
		},
	}
	terraStagingEnvironment = CreatableEnvironment{
		Name:      "terra-staging",
		Base:      "live",
		Lifecycle: "static",
		EditableEnvironment: EditableEnvironment{
			DefaultCluster:      &terraStagingCluster.Name,
			RequiresSuitability: utils.PointerTo(true),
		},
	}
	terraProdEnvironment = CreatableEnvironment{
		Name:      "terra-prod",
		Base:      "live",
		Lifecycle: "static",
		EditableEnvironment: EditableEnvironment{
			DefaultCluster:      &terraProdCluster.Name,
			RequiresSuitability: utils.PointerTo(true),
		},
	}
	swatomationEnvironment = CreatableEnvironment{
		Name:                      "swatomation",
		Base:                      "bee",
		Lifecycle:                 "template",
		AutoPopulateChartReleases: utils.PointerTo(false),
		EditableEnvironment: EditableEnvironment{
			DefaultCluster: &terraQaBeesCluster.Name,
		},
	}
	dynamicSwatomationEnvironment = CreatableEnvironment{
		Name:                "swatomation-instance-one",
		TemplateEnvironment: swatomationEnvironment.Name,
	}
	prodlikeTemplateEnvironment = CreatableEnvironment{
		Name:                      "prodlike",
		Base:                      "bee",
		Lifecycle:                 "template",
		AutoPopulateChartReleases: utils.PointerTo(false),
		EditableEnvironment: EditableEnvironment{
			DefaultCluster:             &terraQaBeesCluster.Name,
			DefaultFirecloudDevelopRef: utils.PointerTo("prod"),
		},
	}
	dynamicProdlikeEnvironment = CreatableEnvironment{
		Name:                "prodlike-one",
		TemplateEnvironment: prodlikeTemplateEnvironment.Name,
	}
	environmentSeedList = []CreatableEnvironment{
		terraDevEnvironment,
		terraStagingEnvironment,
		terraProdEnvironment,
		swatomationEnvironment,
		prodlikeTemplateEnvironment,
		dynamicSwatomationEnvironment,
		dynamicProdlikeEnvironment,
	}
)

func (controllerSet *ControllerSet) seedEnvironments(t *testing.T, db *gorm.DB) {
	for _, creatable := range environmentSeedList {
		if _, _, err := controllerSet.EnvironmentController.Create(creatable, generateUser(t, db, true)); err != nil {
			t.Errorf("error seeding environment %s: %v", creatable.Name, err)
		}
	}
}

//
// Controller tests
//

func (suite *environmentControllerSuite) TestEnvironmentCreate() {
	suite.Run("can create a new environment", func() {
		deprecated_db.Truncate(suite.T(), suite.db)
		suite.seedClusters(suite.T(), suite.db)

		suite.Run("static", func() {
			env, created, err := suite.EnvironmentController.Create(terraDevEnvironment, generateUser(suite.T(), suite.db, false))
			assert.NoError(suite.T(), err)
			assert.True(suite.T(), created)
			assert.Equal(suite.T(), terraDevEnvironment.Name, env.Name)
			assert.True(suite.T(), env.ID > 0)
			suite.Run("default non-suitable", func() {
				assert.False(suite.T(), *env.RequiresSuitability)
			})
			suite.Run("default terra-helmfile-ref", func() {
				suite.Assert().Equal("HEAD", *env.HelmfileRef)
			})
			suite.Run("default firecloud-develop ref", func() {
				suite.Assert().Equal("terra-dev", *env.DefaultFirecloudDevelopRef)
			})
		})
		suite.Run("template", func() {
			env, created, err := suite.EnvironmentController.Create(swatomationEnvironment, generateUser(suite.T(), suite.db, false))
			assert.NoError(suite.T(), err)
			assert.True(suite.T(), created)
			assert.Equal(suite.T(), swatomationEnvironment.Name, env.Name)
			assert.True(suite.T(), env.ID > 0)
			suite.Run("default non-suitable", func() {
				assert.False(suite.T(), *env.RequiresSuitability)
			})
			suite.Run("default terra-helmfile ref head", func() {
				suite.Assert().Equal("HEAD", *env.HelmfileRef)
			})
			suite.Run("default firecloud develop ref", func() {
				suite.Assert().Equal("dev", *env.DefaultFirecloudDevelopRef)
			})
		})
		suite.Run("dynamic", func() {
			user := generateUser(suite.T(), suite.db, false)
			env, created, err := suite.EnvironmentController.Create(dynamicSwatomationEnvironment, generateUser(suite.T(), suite.db, false))
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
				assert.Equal(suite.T(), user.Email, *env.Owner)
			})
			suite.Run("namespace of terra-$name", func() {
				assert.Equal(suite.T(), fmt.Sprintf("terra-%s", env.Name), env.DefaultNamespace)
			})
			suite.Run("default terra-helmfile ref head", func() {
				suite.Assert().Equal("HEAD", *env.HelmfileRef)
			})
		})
		suite.Run("dynamic with name prefix", func() {
			prefix := "boo"
			env, created, err := suite.EnvironmentController.Create(CreatableEnvironment{
				TemplateEnvironment: swatomationEnvironment.Name,
				NamePrefix:          prefix,
			}, generateUser(suite.T(), suite.db, false))
			assert.NoError(suite.T(), err)
			assert.True(suite.T(), created)
			assert.True(suite.T(), strings.HasPrefix(env.Name, prefix))
		})
		suite.Run("prodlike template", func() {
			env, created, err := suite.EnvironmentController.Create(prodlikeTemplateEnvironment, generateUser(suite.T(), suite.db, false))
			assert.NoError(suite.T(), err)
			assert.True(suite.T(), created)
			assert.Equal(suite.T(), prodlikeTemplateEnvironment.Name, env.Name)
			assert.True(suite.T(), env.ID > 0)
			suite.Run("default terra-helmfile ref head", func() {
				suite.Assert().Equal("HEAD", *env.HelmfileRef)
			})
			suite.Run("overrides default firecloud develop ref", func() {
				suite.Assert().Equal(*prodlikeTemplateEnvironment.DefaultFirecloudDevelopRef, *env.DefaultFirecloudDevelopRef)
			})
		})
		suite.Run("dynamic prodlike", func() {
			user := generateUser(suite.T(), suite.db, false)
			env, created, err := suite.EnvironmentController.Create(dynamicProdlikeEnvironment, generateUser(suite.T(), suite.db, false))
			assert.NoError(suite.T(), err)
			assert.True(suite.T(), created)
			assert.Equal(suite.T(), dynamicProdlikeEnvironment.Name, env.Name)
			assert.True(suite.T(), env.ID > 0)
			suite.Run("references template name in defaults", func() {
				assert.Equal(suite.T(), prodlikeTemplateEnvironment.Name, env.TemplateEnvironment)
				assert.Equal(suite.T(), prodlikeTemplateEnvironment.Name, env.ValuesName)
			})
			suite.Run("base of template", func() {
				assert.Equal(suite.T(), prodlikeTemplateEnvironment.Base, env.Base)
			})
			suite.Run("cluster of template", func() {
				assert.Equal(suite.T(), prodlikeTemplateEnvironment.DefaultCluster, env.DefaultCluster)
			})
			suite.Run("fills owner", func() {
				assert.Equal(suite.T(), user.Email, *env.Owner)
			})
			suite.Run("namespace of terra-$name", func() {
				assert.Equal(suite.T(), fmt.Sprintf("terra-%s", env.Name), env.DefaultNamespace)
			})
			suite.Run("default terra-helmfile ref head", func() {
				suite.Assert().Equal("HEAD", *env.HelmfileRef)
			})
			suite.Run("uses default firecloud develop ref from template", func() {
				suite.Assert().Equal(*prodlikeTemplateEnvironment.DefaultFirecloudDevelopRef, *env.DefaultFirecloudDevelopRef)
			})
		})
	})
	suite.Run("won't create duplicates", func() {
		deprecated_db.Truncate(suite.T(), suite.db)
		suite.seedClusters(suite.T(), suite.db)

		env, created, err := suite.EnvironmentController.Create(terraDevEnvironment, generateUser(suite.T(), suite.db, false))
		assert.NoError(suite.T(), err)
		assert.True(suite.T(), created)
		assert.True(suite.T(), env.ID > 0)
		_, created, err = suite.EnvironmentController.Create(terraDevEnvironment, generateUser(suite.T(), suite.db, false))
		assert.ErrorContains(suite.T(), err, errors.Conflict)
		assert.False(suite.T(), created)
	})
	suite.Run("validates incoming entries", func() {
		deprecated_db.Truncate(suite.T(), suite.db)

		_, created, err := suite.EnvironmentController.Create(CreatableEnvironment{}, generateUser(suite.T(), suite.db, false))
		assert.ErrorContains(suite.T(), err, errors.BadRequest)
		assert.False(suite.T(), created)
	})
	suite.Run("validates incoming user against DB", func() {
		deprecated_db.Truncate(suite.T(), suite.db)
		suite.seedClusters(suite.T(), suite.db)

		_, created, err := suite.EnvironmentController.Create(CreatableEnvironment{
			Name:             "terra-dev-2",
			Base:             "live",
			Lifecycle:        "static",
			DefaultNamespace: "terra-dev-2",
			EditableEnvironment: EditableEnvironment{
				DefaultCluster: &terraDevCluster.Name,
				Owner:          utils.PointerTo("some-email-not-in-the-database@broadinstitute.ogr"),
			},
		}, generateUser(suite.T(), suite.db, false))
		suite.Assert().ErrorContains(err, errors.NotFound)
		suite.Assert().False(created)

	})
	suite.Run("checks suitability", func() {
		deprecated_db.Truncate(suite.T(), suite.db)
		suite.seedClusters(suite.T(), suite.db)

		assert.True(suite.T(), *terraProdEnvironment.RequiresSuitability)
		suite.Run("blocks suitable creation for non-suitable", func() {
			_, created, err := suite.EnvironmentController.Create(terraProdEnvironment, generateUser(suite.T(), suite.db, false))
			assert.ErrorContains(suite.T(), err, errors.Forbidden)
			assert.False(suite.T(), created)
		})
		suite.Run("allows suitable creation for suitable", func() {
			env, created, err := suite.EnvironmentController.Create(terraProdEnvironment, generateUser(suite.T(), suite.db, true))
			assert.NoError(suite.T(), err)
			assert.True(suite.T(), created)
			assert.True(suite.T(), env.ID > 0)
		})
	})
	suite.Run("dynamic defaulting respects template edits", func() {
		deprecated_db.Truncate(suite.T(), suite.db)
		suite.seedClusters(suite.T(), suite.db)
		suite.seedEnvironments(suite.T(), suite.db)

		swat, err := suite.EnvironmentController.Get(swatomationEnvironment.Name)
		assert.NoError(suite.T(), err)
		assert.False(suite.T(), *swat.RequiresSuitability)

		env1, created, err := suite.EnvironmentController.Create(CreatableEnvironment{TemplateEnvironment: swatomationEnvironment.Name}, generateUser(suite.T(), suite.db, false))
		assert.NoError(suite.T(), err)
		assert.True(suite.T(), created)
		assert.False(suite.T(), *env1.RequiresSuitability)

		swat, err = suite.EnvironmentController.Edit(swatomationEnvironment.Name, EditableEnvironment{RequiresSuitability: utils.PointerTo(true)}, generateUser(suite.T(), suite.db, true))
		assert.NoError(suite.T(), err)
		assert.True(suite.T(), *swat.RequiresSuitability)
		assert.False(suite.T(), *env1.RequiresSuitability)

		env2, created, err := suite.EnvironmentController.Create(CreatableEnvironment{TemplateEnvironment: swatomationEnvironment.Name}, generateUser(suite.T(), suite.db, true))
		assert.NoError(suite.T(), err)
		assert.True(suite.T(), created)
		assert.True(suite.T(), *env2.RequiresSuitability)
		assert.False(suite.T(), *env1.RequiresSuitability)
	})
	suite.Run("copies template chart releases", func() {
		deprecated_db.Truncate(suite.T(), suite.db)
		suite.seedClusters(suite.T(), suite.db)
		suite.seedEnvironments(suite.T(), suite.db)
		suite.seedCharts(suite.T(), suite.db)
		suite.seedAppVersions(suite.T(), suite.db)
		suite.seedChartVersions(suite.T(), suite.db)
		suite.seedChartReleases(suite.T(), suite.db)

		swatReleases, err := suite.ChartReleaseController.ListAllMatching(
			ChartRelease{CreatableChartRelease: CreatableChartRelease{Environment: swatomationEnvironment.Name}}, 0)
		assert.NoError(suite.T(), err)
		assert.True(suite.T(), len(swatReleases) > 0)

		environment, created, err := suite.EnvironmentController.Create(CreatableEnvironment{TemplateEnvironment: swatomationEnvironment.Name}, generateUser(suite.T(), suite.db, false))
		assert.NoError(suite.T(), err)
		assert.True(suite.T(), created)
		suite.Run("copies base domain", func() {
			assert.Equal(suite.T(), "bee.envs-terra.bio", *environment.BaseDomain)
		})
		suite.Run("copies name prefixing", func() {
			assert.Equal(suite.T(), true, *environment.NamePrefixesDomain)
		})

		environmentReleases, err := suite.ChartReleaseController.ListAllMatching(
			ChartRelease{CreatableChartRelease: CreatableChartRelease{Environment: environment.Name}}, 0)
		assert.NoError(suite.T(), err)

		for _, swatRelease := range swatReleases {
			found := false
			for _, envRelease := range environmentReleases {
				if swatRelease.Chart == envRelease.Chart {
					found = true
					// The template gets this field set based on the Chart's app main branch, dynamic should copy
					assert.Equal(suite.T(), swatRelease.AppVersionBranch, envRelease.AppVersionBranch)
					assert.True(suite.T(), envRelease.ID > 0)
					assert.Equal(suite.T(), swatRelease.Subdomain, envRelease.Subdomain)
					if envRelease.ChartInfo.LegacyConfigsEnabled != nil && *envRelease.ChartInfo.LegacyConfigsEnabled {
						suite.Assert().Equal(*environment.DefaultFirecloudDevelopRef, *envRelease.FirecloudDevelopRef)
					}
					if envRelease.ChartInfo.LegacyConfigsEnabled == nil || !*envRelease.ChartInfo.LegacyConfigsEnabled {
						suite.Assert().Nil(envRelease.FirecloudDevelopRef, "firecloud dev ref should be nil when legacy configs are not enabled")
					}
				}
			}
			assert.True(suite.T(), found)
		}
	})
}

func (suite *environmentControllerSuite) TestEnvironmentAutoPopulateChartReleases() {
	deprecated_db.Truncate(suite.T(), suite.db)
	suite.seedClusters(suite.T(), suite.db)
	suite.seedEnvironments(suite.T(), suite.db)
	suite.seedCharts(suite.T(), suite.db)
	suite.seedAppVersions(suite.T(), suite.db)
	suite.seedChartVersions(suite.T(), suite.db)
	suite.seedChartReleases(suite.T(), suite.db)
	suite.seedDatabaseInstances(suite.T(), suite.db)

	suite.Run("check that honeycomb is still in the config", func() {
		autoPopulateCharts := config.Config.Slices("model.environments.templates.autoPopulateCharts")
		assert.NotNil(suite.T(), autoPopulateCharts)
		honeycombPresent := false
		for _, entry := range autoPopulateCharts {
			if entry.String("name") == "honeycomb" {
				honeycombPresent = true
			}
		}
		assert.True(suite.T(), honeycombPresent)
	})

	suite.Run("template includes honeycomb by default", func() {
		template, created, err := suite.EnvironmentController.Create(CreatableEnvironment{
			Name:      "some-template",
			Base:      "bee",
			Lifecycle: "template",
			EditableEnvironment: EditableEnvironment{
				DefaultCluster: &terraQaBeesCluster.Name,
			},
		}, generateUser(suite.T(), suite.db, false))
		assert.NoError(suite.T(), err)
		assert.True(suite.T(), created)
		templateChartReleases, err := suite.ChartReleaseController.ListAllMatching(ChartRelease{
			CreatableChartRelease: CreatableChartRelease{
				Environment: template.Name,
			},
		}, 0)
		defaultTemplateChartReleaseCount := len(templateChartReleases)
		assert.NoError(suite.T(), err)
		honeycombPresent := false
		for _, chartRelease := range templateChartReleases {
			if chartRelease.Chart == "honeycomb" {
				honeycombPresent = true
			}
		}
		assert.True(suite.T(), honeycombPresent)

		suite.Run("can add to template", func() {
			_, created, err = suite.ChartReleaseController.Create(CreatableChartRelease{
				Chart:       datarepoChart.Name,
				Environment: template.Name,
			}, generateUser(suite.T(), suite.db, false))
			assert.NoError(suite.T(), err)
			assert.True(suite.T(), created)
			_, created, err = suite.DatabaseInstanceController.Create(CreatableDatabaseInstance{
				ChartRelease: fmt.Sprintf("%s/%s", template.Name, datarepoChart.Name),
			}, generateUser(suite.T(), suite.db, false))
			assert.NoError(suite.T(), err)
			assert.True(suite.T(), created)

			templateChartReleases, err = suite.ChartReleaseController.ListAllMatching(ChartRelease{
				CreatableChartRelease: CreatableChartRelease{
					Environment: template.Name,
				},
			}, 0)
			assert.NoError(suite.T(), err)
			assert.Equal(suite.T(), defaultTemplateChartReleaseCount+1, len(templateChartReleases))
			honeycombPresent = false
			for _, chartRelease := range templateChartReleases {
				if chartRelease.Chart == "honeycomb" {
					honeycombPresent = true
				}
			}
			assert.True(suite.T(), honeycombPresent)
		})

		suite.Run("dynamic environments copy template chart releases", func() {
			bee, created, err := suite.EnvironmentController.Create(CreatableEnvironment{
				TemplateEnvironment: template.Name,
			}, generateUser(suite.T(), suite.db, false))
			assert.NoError(suite.T(), err)
			assert.True(suite.T(), created)
			beeChartReleases, err := suite.ChartReleaseController.ListAllMatching(ChartRelease{
				CreatableChartRelease: CreatableChartRelease{
					Environment: bee.Name,
				},
			}, 0)
			assert.NoError(suite.T(), err)
			assert.Equal(suite.T(), len(templateChartReleases), len(beeChartReleases))
			honeycombPresent = false
			for _, chartRelease := range beeChartReleases {
				if chartRelease.Chart == "honeycomb" {
					honeycombPresent = true
				}
			}
			assert.True(suite.T(), honeycombPresent)
			suite.Run("database instances copied", func() {
				databaseInstanceCopied := false
				for _, templateChartRelease := range templateChartReleases {
					templateDatabaseInstance, err := suite.DatabaseInstanceController.Get(fmt.Sprintf("chart-release/%s", templateChartRelease.Name))
					if err != nil {
						assert.ErrorContains(suite.T(), err, errors.NotFound)
						continue
					} else {
						beeDatabaseInstance, err := suite.DatabaseInstanceController.Get(fmt.Sprintf("chart-release/%s/%s", bee.Name, templateChartRelease.Chart))
						assert.NoError(suite.T(), err)
						assert.NotEqual(suite.T(), templateDatabaseInstance.ID, beeDatabaseInstance.ID)
						databaseInstanceCopied = true
					}
				}
				assert.True(suite.T(), databaseInstanceCopied)
			})
		})
	})
}

func (suite *environmentControllerSuite) TestEnvironmentListAllMatching() {
	deprecated_db.Truncate(suite.T(), suite.db)
	suite.seedClusters(suite.T(), suite.db)
	suite.seedEnvironments(suite.T(), suite.db)

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
	deprecated_db.Truncate(suite.T(), suite.db)
	suite.seedClusters(suite.T(), suite.db)
	suite.seedEnvironments(suite.T(), suite.db)

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
	deprecated_db.Truncate(suite.T(), suite.db)
	suite.seedClusters(suite.T(), suite.db)
	suite.seedEnvironments(suite.T(), suite.db)

	suite.Run("successfully", func() {
		selectors, err := suite.EnvironmentController.GetOtherValidSelectors(terraDevEnvironment.Name)
		assert.NoError(suite.T(), err)
		assert.Equal(suite.T(), 3, len(selectors))
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
		deprecated_db.Truncate(suite.T(), suite.db)
		suite.seedClusters(suite.T(), suite.db)
		suite.seedEnvironments(suite.T(), suite.db)

		before, err := suite.EnvironmentController.Get(terraDevEnvironment.Name)
		assert.NoError(suite.T(), err)
		assert.Equal(suite.T(), terraDevEnvironment.Description, before.Description)
		assert.Equal(suite.T(), terraDevEnvironment.PactIdentifier, before.PactIdentifier)
		newDescription := utils.PointerTo("some description")
		pactUuid := utils.PointerTo(uuid.MustParse("7b95b0d1-407a-4e35-a95e-e867139d49b4"))
		assert.NoError(suite.T(), err)

		edited, err := suite.EnvironmentController.Edit(terraDevEnvironment.Name, EditableEnvironment{
			Description:    newDescription,
			PactIdentifier: pactUuid},
			generateUser(suite.T(), suite.db, false))
		assert.NoError(suite.T(), err)
		assert.Equal(suite.T(), newDescription, edited.Description)
		assert.Equal(suite.T(), pactUuid, edited.PactIdentifier)

		after, err := suite.EnvironmentController.Get(terraDevEnvironment.Name)
		assert.NoError(suite.T(), err)
		assert.Equal(suite.T(), newDescription, after.Description)
		assert.Equal(suite.T(), pactUuid, after.PactIdentifier)
	})
	suite.Run("edit to suitable environment", func() {
		deprecated_db.Truncate(suite.T(), suite.db)
		suite.seedClusters(suite.T(), suite.db)
		suite.seedEnvironments(suite.T(), suite.db)
		newDescription := utils.PointerTo("some description")

		suite.Run("unsuccessfully if not suitable", func() {
			_, err := suite.EnvironmentController.Edit(terraProdEnvironment.Name, EditableEnvironment{Description: newDescription}, generateUser(suite.T(), suite.db, false))
			assert.ErrorContains(suite.T(), err, errors.Forbidden)
			notEdited, err := suite.EnvironmentController.Get(terraProdEnvironment.Name)
			assert.NoError(suite.T(), err)
			assert.NotEqual(suite.T(), newDescription, notEdited.Description)
		})
		suite.Run("successfully if suitable", func() {
			edited, err := suite.EnvironmentController.Edit(terraProdEnvironment.Name, EditableEnvironment{Description: newDescription}, generateUser(suite.T(), suite.db, true))
			assert.NoError(suite.T(), err)
			assert.Equal(suite.T(), newDescription, edited.Description)
		})
	})
	suite.Run("edit that would make environment suitable", func() {
		deprecated_db.Truncate(suite.T(), suite.db)
		suite.seedClusters(suite.T(), suite.db)
		suite.seedEnvironments(suite.T(), suite.db)

		suite.Run("unsuccessfully if not suitable", func() {
			_, err := suite.EnvironmentController.Edit(terraDevEnvironment.Name, EditableEnvironment{RequiresSuitability: utils.PointerTo(true)}, generateUser(suite.T(), suite.db, false))
			assert.ErrorContains(suite.T(), err, errors.Forbidden)
			notEdited, err := suite.EnvironmentController.Get(terraDevEnvironment.Name)
			assert.NoError(suite.T(), err)
			assert.False(suite.T(), *notEdited.RequiresSuitability)
		})
		suite.Run("successfully if suitable", func() {
			edited, err := suite.EnvironmentController.Edit(terraDevEnvironment.Name, EditableEnvironment{RequiresSuitability: utils.PointerTo(true)}, generateUser(suite.T(), suite.db, true))
			assert.NoError(suite.T(), err)
			assert.True(suite.T(), *edited.RequiresSuitability)
		})
	})
	suite.Run("unsuccessfully if invalid", func() {
		deprecated_db.Truncate(suite.T(), suite.db)
		suite.seedClusters(suite.T(), suite.db)
		suite.seedEnvironments(suite.T(), suite.db)

		_, err := suite.EnvironmentController.Edit(terraDevEnvironment.Name, EditableEnvironment{Owner: utils.PointerTo("")}, generateUser(suite.T(), suite.db, false))
		assert.ErrorContains(suite.T(), err, errors.BadRequest)
	})
}

func (suite *environmentControllerSuite) TestEnvironmentDelete() {
	suite.Run("successfully", func() {
		deprecated_db.Truncate(suite.T(), suite.db)
		suite.seedClusters(suite.T(), suite.db)
		suite.seedEnvironments(suite.T(), suite.db)
		suite.seedCharts(suite.T(), suite.db)
		suite.seedChartVersions(suite.T(), suite.db)
		suite.seedAppVersions(suite.T(), suite.db)
		suite.seedChartReleases(suite.T(), suite.db)

		chartReleases, err := suite.ChartReleaseController.ListAllMatching(ChartRelease{CreatableChartRelease: CreatableChartRelease{Environment: terraDevEnvironment.Name}}, 0)
		assert.NoError(suite.T(), err)
		assert.True(suite.T(), len(chartReleases) > 0)

		deleted, err := suite.EnvironmentController.Delete(terraDevEnvironment.Name, generateUser(suite.T(), suite.db, false))
		assert.NoError(suite.T(), err)
		assert.Equal(suite.T(), terraDevEnvironment.Name, deleted.Name)
		_, err = suite.EnvironmentController.Get(terraDevEnvironment.Name)
		assert.ErrorContains(suite.T(), err, errors.NotFound)
		suite.Run("deletions cascaded", func() {
			for _, chartRelease := range chartReleases {
				shouldBeEmpty, err := suite.ChartReleaseController.ListAllMatching(ChartRelease{ReadableBaseType: ReadableBaseType{ID: chartRelease.ID}}, 1)
				assert.NoError(suite.T(), err)
				assert.Len(suite.T(), shouldBeEmpty, 0)
			}
		})
		suite.Run("allows re-creation", func() {
			environment, created, err := suite.EnvironmentController.Create(terraDevEnvironment, generateUser(suite.T(), suite.db, false))
			assert.NoError(suite.T(), err)
			assert.True(suite.T(), created)
			assert.NotEqual(suite.T(), deleted.ID, environment.ID)
		})
	})
	suite.Run("delete suitable environment", func() {
		deprecated_db.Truncate(suite.T(), suite.db)
		suite.seedClusters(suite.T(), suite.db)
		suite.seedEnvironments(suite.T(), suite.db)

		suite.Run("unsuccessfully if not suitable", func() {
			_, err := suite.EnvironmentController.Delete(terraProdEnvironment.Name, generateUser(suite.T(), suite.db, false))
			assert.ErrorContains(suite.T(), err, errors.Forbidden)
		})
		suite.Run("successfully if suitable", func() {
			deleted, err := suite.EnvironmentController.Delete(terraProdEnvironment.Name, generateUser(suite.T(), suite.db, true))
			assert.NoError(suite.T(), err)
			assert.Equal(suite.T(), terraProdEnvironment.Name, deleted.Name)
		})
	})
}
