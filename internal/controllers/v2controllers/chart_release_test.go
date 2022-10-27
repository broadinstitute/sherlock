package v2controllers

import (
	"fmt"
	"testing"

	"github.com/broadinstitute/sherlock/internal/auth"
	"github.com/broadinstitute/sherlock/internal/config"
	"github.com/broadinstitute/sherlock/internal/db"
	"github.com/broadinstitute/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/internal/models/v2models"
	"github.com/broadinstitute/sherlock/internal/testutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

//
// Test suite configuration
//

func TestChartReleaseControllerSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping functional test")
	}
	suite.Run(t, new(chartReleaseControllerSuite))
}

type chartReleaseControllerSuite struct {
	suite.Suite
	*ControllerSet
	db *gorm.DB
}

func (suite *chartReleaseControllerSuite) SetupTest() {
	config.LoadTestConfig(suite.T())
	suite.db = db.ConnectAndConfigureFromTest(suite.T())
	suite.db.Begin()
	suite.ControllerSet = NewControllerSet(v2models.NewStoreSet(suite.db))
}

func (suite *chartReleaseControllerSuite) TearDownTest() {
	suite.db.Rollback()
}

//
// Controller seeding
//

var (
	leonardoDevChartRelease = CreatableChartRelease{
		Chart:       leonardoChart.Name,
		Environment: terraDevEnvironment.Name,
	}
	datarepoDevChartRelease = CreatableChartRelease{
		Name:              "datarepo-dev",
		Chart:             datarepoChart.Name,
		Environment:       terraDevEnvironment.Name,
		Cluster:           datarepoDevCluster.Name,
		AppVersionExact:   testutils.PointerTo("a1b2c3d4"),
		ChartVersionExact: testutils.PointerTo("1.2.3"),
	}
	samDevChartRelease = CreatableChartRelease{
		Chart:       samChart.Name,
		Environment: terraDevEnvironment.Name,
	}
	yaleDevChartRelease = CreatableChartRelease{
		Chart:       yaleChart.Name,
		Environment: terraDevEnvironment.Name,
		Namespace:   "yale",
	}
	storageDevChartRelease = CreatableChartRelease{
		Chart:     terraClusterStorageChart.Name,
		Cluster:   terraDevCluster.Name,
		Namespace: "default",
	}

	leonardoProdChartRelease = CreatableChartRelease{
		Chart:             leonardoChart.Name,
		Environment:       terraProdEnvironment.Name,
		AppVersionExact:   testutils.PointerTo("a1b2c3d4"),
		ChartVersionExact: testutils.PointerTo("1.2.3"),
	}
	datarepoProdChartRelease = CreatableChartRelease{
		Name:              "datarepo-prod",
		Chart:             datarepoChart.Name,
		Environment:       terraProdEnvironment.Name,
		Cluster:           datarepoProdCluster.Name,
		AppVersionExact:   testutils.PointerTo("2.200.1"),
		ChartVersionExact: testutils.PointerTo("0.10.0"),
	}
	samProdChartRelease = CreatableChartRelease{
		Chart:             samChart.Name,
		Environment:       terraProdEnvironment.Name,
		AppVersionExact:   &samMain2AppVersion.AppVersion,
		ChartVersionExact: &sam2ChartVersion.ChartVersion,
	}
	yaleProdChartRelease = CreatableChartRelease{
		Chart:             yaleChart.Name,
		Environment:       terraProdEnvironment.Name,
		Namespace:         "yale",
		AppVersionExact:   testutils.PointerTo("0.0.18"),
		ChartVersionExact: testutils.PointerTo("1.15.0"),
	}
	storageProdChartRelease = CreatableChartRelease{
		Chart:     terraClusterStorageChart.Name,
		Cluster:   terraProdCluster.Name,
		Namespace: "default",
	}

	leonardoStagingChartRelease = CreatableChartRelease{
		Chart:             leonardoChart.Name,
		Environment:       terraStagingEnvironment.Name,
		AppVersionExact:   testutils.PointerTo("a1b2c3d4"),
		ChartVersionExact: testutils.PointerTo("1.2.3"),
	}
	samStagingChartRelease = CreatableChartRelease{
		Chart:             samChart.Name,
		Environment:       terraStagingEnvironment.Name,
		AppVersionExact:   &samMain2AppVersion.AppVersion,
		ChartVersionExact: &sam2ChartVersion.ChartVersion,
	}
	yaleStagingChartRelease = CreatableChartRelease{
		Chart:             yaleChart.Name,
		Environment:       terraStagingEnvironment.Name,
		Namespace:         "yale",
		AppVersionExact:   testutils.PointerTo("0.0.18"),
		ChartVersionExact: testutils.PointerTo("1.15.0"),
	}
	storageStagingChartRelease = CreatableChartRelease{
		Chart:     terraClusterStorageChart.Name,
		Cluster:   terraStagingCluster.Name,
		Namespace: "default",
	}

	storageBeeChartRelease = CreatableChartRelease{
		Chart:     terraClusterStorageChart.Name,
		Cluster:   terraQaBeesCluster.Name,
		Namespace: "default",
	}

	leonardoSwatomationChartRelease = CreatableChartRelease{
		Chart:       leonardoChart.Name,
		Environment: swatomationEnvironment.Name,
	}
	datarepoSwatomationChartRelease = CreatableChartRelease{
		Chart:       datarepoChart.Name,
		Environment: swatomationEnvironment.Name,
	}
	samSwatomationChartRelease = CreatableChartRelease{
		Chart:       samChart.Name,
		Environment: swatomationEnvironment.Name,
	}
	honeycombSwatomationChartRelease = CreatableChartRelease{
		Chart:       honeycombChart.Name,
		Environment: swatomationEnvironment.Name,
	}

	leonardoDynamicSwatomationChartRelease = CreatableChartRelease{
		Chart:       leonardoChart.Name,
		Environment: dynamicSwatomationEnvironment.Name,
	}
	datarepoDynamicSwatomationChartRelease = CreatableChartRelease{
		Chart:       datarepoChart.Name,
		Environment: dynamicSwatomationEnvironment.Name,
	}
	honeycombDynamicSwatomationChartRelease = CreatableChartRelease{
		Chart:       honeycombChart.Name,
		Environment: dynamicSwatomationEnvironment.Name,
	}

	chartReleaseSeedList = []CreatableChartRelease{
		leonardoDevChartRelease,
		datarepoDevChartRelease,
		samDevChartRelease,
		yaleDevChartRelease,
		storageDevChartRelease,
		leonardoProdChartRelease,
		datarepoProdChartRelease,
		samProdChartRelease,
		yaleProdChartRelease,
		storageProdChartRelease,
		leonardoStagingChartRelease,
		samStagingChartRelease,
		yaleStagingChartRelease,
		storageStagingChartRelease,
		storageBeeChartRelease,
		leonardoSwatomationChartRelease,
		datarepoSwatomationChartRelease,
		samSwatomationChartRelease,
		honeycombSwatomationChartRelease,
		leonardoDynamicSwatomationChartRelease,
		datarepoDynamicSwatomationChartRelease,
		honeycombDynamicSwatomationChartRelease,
	}
)

func (controllerSet *ControllerSet) seedChartReleases(t *testing.T) {
	for _, creatable := range chartReleaseSeedList {
		if _, _, err := controllerSet.ChartReleaseController.Create(creatable, auth.GenerateUser(t, true)); err != nil {
			t.Errorf("error seeding chart release for %s in %s/%s: %v", creatable.Chart, creatable.Environment, creatable.Cluster, err)
		}
	}
}

//
// Controller tests
//

func (suite *chartReleaseControllerSuite) TestChartReleaseCreate() {
	suite.Run("can create a new chart release", func() {
		db.Truncate(suite.T(), suite.db)
		suite.seedClusters(suite.T())
		suite.seedEnvironments(suite.T())
		suite.seedCharts(suite.T())
		suite.seedAppVersions(suite.T())
		suite.seedChartVersions(suite.T())

		suite.Run("simple app release", func() {
			release, created, err := suite.ChartReleaseController.Create(leonardoDevChartRelease, auth.GenerateUser(suite.T(), false))
			assert.NoError(suite.T(), err)
			assert.True(suite.T(), created)
			assert.True(suite.T(), release.ID > 0)
			suite.Run("defaults target app version branch to that of chart", func() {
				assert.Equal(suite.T(), leonardoChart.AppImageGitMainBranch, release.AppVersionBranch)
			})
			suite.Run("defaults target app version use to branch when nothing else set", func() {
				assert.Equal(suite.T(), testutils.PointerTo("branch"), release.AppVersionResolver)
				suite.Run("automatically derives exact version", func() {
					assert.Equal(suite.T(), leonardoMain3AppVersion.AppVersion, *release.AppVersionExact)
				})
			})
			suite.Run("defaults target chart version use to latest", func() {
				assert.Equal(suite.T(), testutils.PointerTo("latest"), release.ChartVersionResolver)
				suite.Run("automatically derives exact version", func() {
					assert.Equal(suite.T(), leonardo3ChartVersion.ChartVersion, *release.ChartVersionExact)
				})
			})
			suite.Run("defaults name to chart-env if possible", func() {
				assert.Equal(suite.T(), fmt.Sprintf("%s-%s", leonardoChart.Name, terraDevEnvironment.Name), release.Name)
			})
			suite.Run("defaults cluster to that of environment", func() {
				assert.Equal(suite.T(), terraDevCluster.Name, release.Cluster)
			})
			suite.Run("defaults namespace to that of environment", func() {
				assert.Equal(suite.T(), terraDevEnvironment.DefaultNamespace, release.Namespace)
			})
			suite.Run("sets destination type to environment", func() {
				assert.Equal(suite.T(), "environment", release.DestinationType)
			})
			suite.Run("default to chart subdomain", func() {
				assert.Equal(suite.T(), "leonardo", *release.Subdomain)
			})
			suite.Run("default to chart protocol", func() {
				assert.Equal(suite.T(), "https", *release.Protocol)
			})
			suite.Run("default to chart port", func() {
				assert.Equal(suite.T(), uint(443), *release.Port)
			})
			suite.Run("has legacy configs enabeld and a defaultFirecloudDevelopRef on environment", func() {
				suite.Assert().Equal("terra-dev", *release.FirecloudDevelopRef)
			})
		})
		suite.Run("custom cluster app release", func() {
			release, created, err := suite.ChartReleaseController.Create(datarepoDevChartRelease, auth.GenerateUser(suite.T(), false))
			assert.NoError(suite.T(), err)
			assert.True(suite.T(), created)
			assert.True(suite.T(), release.ID > 0)
			suite.Run("name doesn't get overridden when it is provided", func() {
				assert.Equal(suite.T(), datarepoDevChartRelease.Name, release.Name)
			})
			suite.Run("cluster doesn't get overridden when it is provided", func() {
				assert.Equal(suite.T(), datarepoDevChartRelease.Cluster, release.Cluster)
			})
			suite.Run("namespace still gets defaulted", func() {
				assert.Equal(suite.T(), terraDevEnvironment.DefaultNamespace, release.Namespace)
			})
		})
		suite.Run("release in an env but with a set namespace", func() {
			release, created, err := suite.ChartReleaseController.Create(yaleDevChartRelease, auth.GenerateUser(suite.T(), false))
			assert.NoError(suite.T(), err)
			assert.True(suite.T(), created)
			assert.True(suite.T(), release.ID > 0)
			suite.Run("namespace doesn't get overridden when it is provided", func() {
				assert.Equal(suite.T(), yaleDevChartRelease.Namespace, release.Namespace)
			})
		})
		suite.Run("cluster release", func() {
			release, created, err := suite.ChartReleaseController.Create(storageDevChartRelease, auth.GenerateUser(suite.T(), false))
			assert.NoError(suite.T(), err)
			assert.True(suite.T(), created)
			assert.True(suite.T(), release.ID > 0)
			suite.Run("environment stays empty when not provided", func() {
				assert.Equal(suite.T(), "", release.Environment)
			})
			suite.Run("defaults name to chart-namespace-cluster if env omitted", func() {
				assert.Equal(suite.T(), fmt.Sprintf("%s-%s-%s", terraClusterStorageChart.Name, storageDevChartRelease.Namespace, storageDevChartRelease.Cluster), release.Name)
			})
			suite.Run("sets destination type to cluster", func() {
				assert.Equal(suite.T(), "cluster", release.DestinationType)
			})
			suite.Run("chart without an endpoint doesn't fill endpoint fields in chart release", func() {
				assert.Empty(suite.T(), release.Subdomain)
				assert.Empty(suite.T(), release.Protocol)
				assert.Empty(suite.T(), release.Port)
			})
		})
		suite.Run("won't create duplicates", func() {
			db.Truncate(suite.T(), suite.db)
			suite.seedClusters(suite.T())
			suite.seedEnvironments(suite.T())
			suite.seedCharts(suite.T())
			suite.seedAppVersions(suite.T())
			suite.seedChartVersions(suite.T())
			suite.seedChartReleases(suite.T())

			suite.Run("exact duplicate", func() {
				_, created, err := suite.ChartReleaseController.Create(leonardoDevChartRelease, auth.GenerateUser(suite.T(), false))
				assert.ErrorContains(suite.T(), err, errors.Conflict)
				assert.False(suite.T(), created)
			})
			suite.Run("duplicate chart/env", func() {
				_, created, err := suite.ChartReleaseController.Create(CreatableChartRelease{
					Chart:       leonardoChart.Name,
					Environment: terraDevEnvironment.Name,
					Namespace:   "abc",
					Name:        "def",
				}, auth.GenerateUser(suite.T(), false))
				assert.ErrorContains(suite.T(), err, errors.Conflict)
				assert.False(suite.T(), created)
			})
			suite.Run("duplicate chart/namespace/cluster", func() {
				release, err := suite.ChartReleaseController.Get(fmt.Sprintf("%s-%s", leonardoChart.Name, terraDevEnvironment.Name))
				assert.NoError(suite.T(), err)
				_, created, err := suite.ChartReleaseController.Create(CreatableChartRelease{
					Chart:     leonardoChart.Name,
					Cluster:   release.Cluster,
					Namespace: release.Namespace,
					Name:      "abc",
				}, auth.GenerateUser(suite.T(), false))
				assert.ErrorContains(suite.T(), err, errors.Conflict)
				assert.False(suite.T(), created)
			})
			suite.Run("duplicate name", func() {
				_, created, err := suite.ChartReleaseController.Create(CreatableChartRelease{
					Chart:       yaleChart.Name,
					Environment: dynamicSwatomationEnvironment.Name,
					// This name exists associated to another environment, namespace, and cluster
					Name: datarepoDevChartRelease.Name,
				}, auth.GenerateUser(suite.T(), false))
				assert.ErrorContains(suite.T(), err, errors.Conflict)
				assert.False(suite.T(), created)
			})
		})
		suite.Run("validates incoming entries", func() {
			db.Truncate(suite.T(), suite.db)
			suite.seedClusters(suite.T())
			suite.seedEnvironments(suite.T())
			suite.seedCharts(suite.T())
			suite.seedAppVersions(suite.T())
			suite.seedChartVersions(suite.T())

			suite.Run("no associations", func() {
				_, created, err := suite.ChartReleaseController.Create(CreatableChartRelease{}, auth.GenerateUser(suite.T(), false))
				assert.ErrorContains(suite.T(), err, errors.BadRequest)
				assert.False(suite.T(), created)
			})
			suite.Run("good associations but bad values", func() {
				_, created, err := suite.ChartReleaseController.Create(CreatableChartRelease{
					Chart:                leonardoChart.Name,
					Environment:          terraDevEnvironment.Name,
					ChartVersionResolver: testutils.PointerTo("something obviously incorrect"),
				}, auth.GenerateUser(suite.T(), false))
				assert.ErrorContains(suite.T(), err, errors.BadRequest)
				assert.False(suite.T(), created)
			})
		})
		suite.Run("checks suitability", func() {
			db.Truncate(suite.T(), suite.db)
			suite.seedClusters(suite.T())
			suite.seedEnvironments(suite.T())
			suite.seedCharts(suite.T())
			suite.seedAppVersions(suite.T())
			suite.seedChartVersions(suite.T())

			suite.Run("blocks suitable creation for non-suitable", func() {
				_, created, err := suite.ChartReleaseController.Create(leonardoProdChartRelease, auth.GenerateUser(suite.T(), false))
				assert.ErrorContains(suite.T(), err, errors.Forbidden)
				assert.False(suite.T(), created)
			})
			suite.Run("allows suitable creation for suitable", func() {
				release, created, err := suite.ChartReleaseController.Create(leonardoProdChartRelease, auth.GenerateUser(suite.T(), true))
				assert.NoError(suite.T(), err)
				assert.True(suite.T(), created)
				assert.True(suite.T(), release.ID > 0)
			})
		})
	})
}

func (suite *chartReleaseControllerSuite) TestChartReleaseListAllMatching() {
	db.Truncate(suite.T(), suite.db)
	suite.seedClusters(suite.T())
	suite.seedEnvironments(suite.T())
	suite.seedCharts(suite.T())
	suite.seedAppVersions(suite.T())
	suite.seedChartVersions(suite.T())
	suite.seedChartReleases(suite.T())

	suite.Run("lists all chart releases", func() {
		matching, err := suite.ChartReleaseController.ListAllMatching(ChartRelease{}, 0)
		assert.NoError(suite.T(), err)
		assert.Equal(suite.T(), len(chartReleaseSeedList), len(matching))
		suite.Run("orders by latest updated", func() {
			latestUpdated := matching[0].UpdatedAt
			for _, chartRelease := range matching {
				assert.GreaterOrEqual(suite.T(), latestUpdated, chartRelease.UpdatedAt)
			}
		})
	})
	suite.Run("limits", func() {
		limit := 2
		matching, err := suite.ChartReleaseController.ListAllMatching(ChartRelease{}, limit)
		assert.NoError(suite.T(), err)
		assert.Equal(suite.T(), limit, len(matching))
	})
	suite.Run("filters exactly", func() {
		matching, err := suite.ChartReleaseController.ListAllMatching(ChartRelease{CreatableChartRelease: leonardoDevChartRelease}, 0)
		assert.NoError(suite.T(), err)
		assert.Equal(suite.T(), 1, len(matching))
		assert.Equal(suite.T(), leonardoDevChartRelease.Chart, matching[0].Chart)
		assert.Equal(suite.T(), leonardoDevChartRelease.Environment, matching[0].Environment)
	})
	suite.Run("filters multiple", func() {
		matching, err := suite.ChartReleaseController.ListAllMatching(
			ChartRelease{CreatableChartRelease: CreatableChartRelease{Cluster: terraDevCluster.Name}}, 0)
		assert.NoError(suite.T(), err)
		assert.True(suite.T(), len(matching) > 1)
		for _, chartRelease := range matching {
			assert.Equal(suite.T(), terraDevCluster.Name, chartRelease.Cluster)
		}
	})
	suite.Run("none is an empty list, not null", func() {
		matching, err := suite.ChartReleaseController.ListAllMatching(
			ChartRelease{CreatableChartRelease: CreatableChartRelease{Name: "blah"}}, 0)
		assert.NoError(suite.T(), err)
		assert.NotNil(suite.T(), matching)
		assert.Empty(suite.T(), matching)
	})
}

func (suite *chartReleaseControllerSuite) TestChartReleaseGet() {
	db.Truncate(suite.T(), suite.db)
	suite.seedClusters(suite.T())
	suite.seedEnvironments(suite.T())
	suite.seedCharts(suite.T())
	suite.seedAppVersions(suite.T())
	suite.seedChartVersions(suite.T())
	suite.seedChartReleases(suite.T())

	suite.Run("successfully", func() {
		var chartReleaseID, chartID, environmentID, clusterID uint
		var namespace string
		suite.Run("by name", func() {
			release, err := suite.ChartReleaseController.Get(datarepoDevChartRelease.Name)
			assert.NoError(suite.T(), err)
			assert.Equal(suite.T(), datarepoDevChartRelease.Name, release.Name)
			assert.True(suite.T(), release.ID > 0)
			chartReleaseID = release.ID
			chartID = release.ChartInfo.ID
			environmentID = release.EnvironmentInfo.ID
			clusterID = release.ClusterInfo.ID
			namespace = release.Namespace
		})
		suite.Run("by ID", func() {
			release, err := suite.ChartReleaseController.Get(fmt.Sprintf("%d", chartReleaseID))
			assert.NoError(suite.T(), err)
			assert.Equal(suite.T(), chartReleaseID, release.ID)
		})
		suite.Run("by environment/chart", func() {
			suite.Run("by name", func() {
				release, err := suite.ChartReleaseController.Get(fmt.Sprintf("%s/%s", terraDevEnvironment.Name, datarepoChart.Name))
				assert.NoError(suite.T(), err)
				assert.Equal(suite.T(), chartReleaseID, release.ID)
			})
			suite.Run("by ID", func() {
				release, err := suite.ChartReleaseController.Get(fmt.Sprintf("%d/%d", environmentID, chartID))
				assert.NoError(suite.T(), err)
				assert.Equal(suite.T(), chartReleaseID, release.ID)
			})
		})
		suite.Run("by cluster/namespace/chart", func() {
			suite.Run("by name", func() {
				release, err := suite.ChartReleaseController.Get(fmt.Sprintf("%s/%s/%s", datarepoDevCluster.Name, namespace, datarepoChart.Name))
				assert.NoError(suite.T(), err)
				assert.Equal(suite.T(), chartReleaseID, release.ID)
			})
			suite.Run("by ID", func() {
				release, err := suite.ChartReleaseController.Get(fmt.Sprintf("%d/%s/%d", clusterID, namespace, chartID))
				assert.NoError(suite.T(), err)
				assert.Equal(suite.T(), chartReleaseID, release.ID)
			})
		})
	})
	suite.Run("unsuccessfully for non-present", func() {
		_, err := suite.ChartReleaseController.Get("foobar")
		assert.ErrorContains(suite.T(), err, errors.NotFound)
	})
	suite.Run("unsuccessfully for invalid selector", func() {
		_, err := suite.ChartReleaseController.Get("something obviously invalid")
		assert.ErrorContains(suite.T(), err, errors.BadRequest)
	})
}

func (suite *chartReleaseControllerSuite) TestChartReleaseGetOtherValidSelectors() {
	db.Truncate(suite.T(), suite.db)
	suite.seedClusters(suite.T())
	suite.seedEnvironments(suite.T())
	suite.seedCharts(suite.T())
	suite.seedAppVersions(suite.T())
	suite.seedChartVersions(suite.T())
	suite.seedChartReleases(suite.T())
	suite.Run("successfully", func() {
		selectors, err := suite.ChartReleaseController.GetOtherValidSelectors(datarepoDevChartRelease.Name)
		assert.NoError(suite.T(), err)
		assert.Equal(suite.T(), 12, len(selectors))
		assert.Contains(suite.T(), selectors, datarepoDevChartRelease.Name)
	})
	suite.Run("unsuccessfully for non-present", func() {
		_, err := suite.ChartReleaseController.GetOtherValidSelectors("foobar")
		assert.ErrorContains(suite.T(), err, errors.NotFound)
	})
	suite.Run("unsuccessfully for invalid selector", func() {
		_, err := suite.ChartReleaseController.GetOtherValidSelectors("something obviously invalid")
		assert.ErrorContains(suite.T(), err, errors.BadRequest)
	})
}

func (suite *chartReleaseControllerSuite) TestChartReleaseEdit() {
	// Right now chart releases don't have any edits but I'm not deleting this because it's a (slightly out of date)
	// skeleton for testing edits

	//suite.Run("successfully", func() {
	//	db.Truncate(suite.T(), suite.db)
	//	suite.seedClusters(suite.T())
	//	suite.seedEnvironments(suite.T())
	//	suite.seedCharts(suite.T())
	//	suite.seedChartReleases(suite.T())
	//
	//	before, err := suite.ChartReleaseController.Get(datarepoDevChartRelease.Name)
	//	assert.NoError(suite.T(), err)
	//	assert.Equal(suite.T(), "latest", *before.TargetChartVersionUse)
	//	edited, err := suite.ChartReleaseController.Edit(datarepoDevChartRelease.Name, EditableChartRelease{
	//		TargetChartVersionUse: testutils.PointerTo("exact"), TargetChartVersionExact: testutils.PointerTo("11.22.33"),
	//	}, auth.GenerateUser(suite.T(), false))
	//	assert.NoError(suite.T(), err)
	//	assert.Equal(suite.T(), "exact", *edited.TargetChartVersionUse)
	//	assert.Equal(suite.T(), "11.22.33", *edited.TargetChartVersionExact)
	//	after, err := suite.ChartReleaseController.Get(datarepoDevChartRelease.Name)
	//	assert.NoError(suite.T(), err)
	//	assert.Equal(suite.T(), "exact", *after.TargetChartVersionUse)
	//	assert.Equal(suite.T(), "11.22.33", *after.TargetChartVersionExact)
	//})
	//suite.Run("edit to suitable chart release", func() {
	//	db.Truncate(suite.T(), suite.db)
	//	suite.seedClusters(suite.T())
	//	suite.seedEnvironments(suite.T())
	//	suite.seedCharts(suite.T())
	//	suite.seedChartReleases(suite.T())
	//
	//	suite.Run("unsuccessfully if not suitable", func() {
	//		before, err := suite.ChartReleaseController.Get(datarepoProdChartRelease.Name)
	//		assert.NoError(suite.T(), err)
	//		assert.Nil(suite.T(), before.TargetAppVersionCommit)
	//		_, err = suite.ChartReleaseController.Edit(datarepoProdChartRelease.Name, EditableChartRelease{
	//			TargetAppVersionCommit: testutils.PointerTo("abc"),
	//		}, auth.GenerateUser(suite.T(), false))
	//		assert.ErrorContains(suite.T(), err, errors.Forbidden)
	//		notEdited, err := suite.ChartReleaseController.Get(datarepoProdChartRelease.Name)
	//		assert.NoError(suite.T(), err)
	//		assert.Nil(suite.T(), notEdited.TargetAppVersionCommit)
	//	})
	//	suite.Run("successfully if suitable", func() {
	//		edited, err := suite.ChartReleaseController.Edit(datarepoProdChartRelease.Name, EditableChartRelease{
	//			TargetAppVersionCommit: testutils.PointerTo("abc"),
	//		}, auth.GenerateUser(suite.T(), true))
	//		assert.NoError(suite.T(), err)
	//		assert.Equal(suite.T(), "abc", *edited.TargetAppVersionCommit)
	//	})
	//})
	//suite.Run("unsuccessfully if invalid", func() {
	//	db.Truncate(suite.T(), suite.db)
	//	suite.seedClusters(suite.T())
	//	suite.seedEnvironments(suite.T())
	//	suite.seedCharts(suite.T())
	//	suite.seedChartReleases(suite.T())
	//
	//	_, err := suite.ChartReleaseController.Edit(datarepoProdChartRelease.Name, EditableChartRelease{
	//		TargetAppVersionUse: testutils.PointerTo("something obviously incorrect"),
	//	}, auth.GenerateUser(suite.T(), true))
	//	assert.ErrorContains(suite.T(), err, errors.BadRequest)
	//})
}

func (suite *chartReleaseControllerSuite) TestChartReleaseDelete() {
	suite.Run("successfully", func() {
		db.Truncate(suite.T(), suite.db)
		suite.seedClusters(suite.T())
		suite.seedEnvironments(suite.T())
		suite.seedCharts(suite.T())
		suite.seedAppVersions(suite.T())
		suite.seedChartVersions(suite.T())
		suite.seedChartReleases(suite.T())

		deleted, err := suite.ChartReleaseController.Delete(datarepoDevChartRelease.Name, auth.GenerateUser(suite.T(), false))
		assert.NoError(suite.T(), err)
		assert.Equal(suite.T(), datarepoDevChartRelease.Name, deleted.Name)
		_, err = suite.ChartReleaseController.Get(datarepoDevChartRelease.Name)
		assert.ErrorContains(suite.T(), err, errors.NotFound)
		suite.Run("sql constraints ignore soft deletion", func() {
			_, created, err := suite.ChartReleaseController.Create(datarepoDevChartRelease, auth.GenerateUser(suite.T(), false))
			assert.ErrorContains(suite.T(), err, errors.BadRequest)
			assert.ErrorContains(suite.T(), err, "Contact DevOps")
			assert.False(suite.T(), created)
		})
	})
	suite.Run("delete suitable chart release", func() {
		db.Truncate(suite.T(), suite.db)
		suite.seedClusters(suite.T())
		suite.seedEnvironments(suite.T())
		suite.seedCharts(suite.T())
		suite.seedAppVersions(suite.T())
		suite.seedChartVersions(suite.T())
		suite.seedChartReleases(suite.T())

		suite.Run("unsuccessfully if not suitable", func() {
			_, err := suite.ChartReleaseController.Delete(datarepoProdChartRelease.Name, auth.GenerateUser(suite.T(), false))
			assert.ErrorContains(suite.T(), err, errors.Forbidden)
		})
		suite.Run("successfully if suitable", func() {
			deleted, err := suite.ChartReleaseController.Delete(datarepoProdChartRelease.Name, auth.GenerateUser(suite.T(), true))
			assert.NoError(suite.T(), err)
			assert.Equal(suite.T(), datarepoProdChartRelease.Name, deleted.Name)
		})
	})
}
