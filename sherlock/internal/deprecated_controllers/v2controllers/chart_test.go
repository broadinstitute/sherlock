package v2controllers

import (
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/testutils"
	"github.com/broadinstitute/sherlock/sherlock/internal/deprecated_db"
	"github.com/broadinstitute/sherlock/sherlock/internal/deprecated_models/v2models"
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

func TestChartControllerSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping functional test")
	}
	suite.Run(t, new(chartControllerSuite))
}

type chartControllerSuite struct {
	suite.Suite
	*ControllerSet
	db *gorm.DB
}

func (suite *chartControllerSuite) SetupTest() {
	config.LoadTestConfig()
	suite.db = deprecated_db.ConnectAndConfigureFromTest(suite.T())
	suite.db.Begin()
	suite.ControllerSet = NewControllerSet(v2models.NewStoreSet(suite.db))
}

func (suite *chartControllerSuite) TearDownTest() {
	suite.db.Rollback()
}

//
// Controller seeding
//

var (
	// "app release" (per environment), normal app chart
	leonardoChart = CreatableChart{
		Name: "leonardo",
		EditableChart: EditableChart{
			AppImageGitRepo:       testutils.PointerTo("DataBiosphere/leonardo"),
			AppImageGitMainBranch: testutils.PointerTo("main"),
			ChartExposesEndpoint:  testutils.PointerTo(true),
			LegacyConfigsEnabled:  testutils.PointerTo(true),
		},
	}
	samChart = CreatableChart{
		Name: "sam",
		EditableChart: EditableChart{
			AppImageGitRepo:       testutils.PointerTo("broadinstitute/sam"),
			AppImageGitMainBranch: testutils.PointerTo("develop"),
			ChartExposesEndpoint:  testutils.PointerTo(true),
			LegacyConfigsEnabled:  testutils.PointerTo(true),
		},
	}
	datarepoChart = CreatableChart{
		Name: "datarepo",
		EditableChart: EditableChart{
			ChartRepo:             testutils.PointerTo("datarepo-helm"),
			AppImageGitRepo:       testutils.PointerTo("DataBiosphere/jade-data-repo"),
			AppImageGitMainBranch: testutils.PointerTo("develop"),
			ChartExposesEndpoint:  testutils.PointerTo(true),
			LegacyConfigsEnabled:  testutils.PointerTo(false),
		},
	}
	// "app release" (per environment), but doesn't actually have an application
	honeycombChart = CreatableChart{
		Name: "honeycomb",
	}
	// "cluster release" (per cluster), doesn't have an application
	terraClusterStorageChart = CreatableChart{
		Name: "terra-cluster-storage",
	}
	// "cluster release" (per cluster), but still deploys an application
	yaleChart = CreatableChart{
		Name: "yale",
		EditableChart: EditableChart{
			AppImageGitRepo:       testutils.PointerTo("broadinstitute/yale"),
			AppImageGitMainBranch: testutils.PointerTo("main"),
		},
	}
	// library sub-chart
	ingressChart = CreatableChart{
		Name: "ingress",
	}
	// library template chart
	yaleLibChart = CreatableChart{
		Name: "yalelib",
	}
	chartSeedList = []CreatableChart{
		leonardoChart,
		samChart,
		datarepoChart,
		honeycombChart,
		terraClusterStorageChart,
		yaleChart,
		ingressChart,
		yaleLibChart,
	}
)

func (controllerSet *ControllerSet) seedCharts(t *testing.T, db *gorm.DB) {
	for _, creatable := range chartSeedList {
		if _, _, err := controllerSet.ChartController.Create(creatable, generateUser(t, db, false)); err != nil {
			t.Errorf("error seeding chart %s: %v", creatable.Name, err)
		}
	}
}

//
// Controller tests
//

func (suite *chartControllerSuite) TestChartCreate() {
	suite.Run("can create a new chart", func() {
		deprecated_db.Truncate(suite.T(), suite.db)

		chart, created, err := suite.ChartController.Create(leonardoChart, generateUser(suite.T(), suite.db, false))
		assert.NoError(suite.T(), err)
		assert.True(suite.T(), created)
		assert.Equal(suite.T(), leonardoChart.Name, chart.Name)
		assert.True(suite.T(), chart.ID > 0)
		suite.Run("default chart repo terra-helm", func() {
			assert.Equal(suite.T(), "terra-helm", *chart.ChartRepo)
		})
		suite.Run("default subdomain is name", func() {
			assert.Equal(suite.T(), "leonardo", *chart.DefaultSubdomain)
		})
		suite.Run("default protocol is https", func() {
			assert.Equal(suite.T(), "https", *chart.DefaultProtocol)
		})
		suite.Run("default port is 443", func() {
			assert.Equal(suite.T(), uint(443), *chart.DefaultPort)
		})
		suite.Run("legacy configs are enabled", func() {
			suite.Assert().Equal(true, *chart.LegacyConfigsEnabled)
		})
	})
	suite.Run("chart repo can be customized", func() {
		deprecated_db.Truncate(suite.T(), suite.db)

		chart, created, err := suite.ChartController.Create(datarepoChart, generateUser(suite.T(), suite.db, false))
		assert.NoError(suite.T(), err)
		assert.True(suite.T(), created)
		assert.Equal(suite.T(), datarepoChart.ChartRepo, chart.ChartRepo)
	})
	suite.Run("won't create duplicates", func() {
		deprecated_db.Truncate(suite.T(), suite.db)

		chart, created, err := suite.ChartController.Create(leonardoChart, generateUser(suite.T(), suite.db, false))
		assert.NoError(suite.T(), err)
		assert.True(suite.T(), created)
		assert.True(suite.T(), chart.ID > 0)
		_, created, err = suite.ChartController.Create(leonardoChart, generateUser(suite.T(), suite.db, false))
		assert.ErrorContains(suite.T(), err, errors.Conflict)
		assert.False(suite.T(), created)
	})
	suite.Run("validates incoming entries", func() {
		deprecated_db.Truncate(suite.T(), suite.db)

		_, created, err := suite.ChartController.Create(CreatableChart{}, generateUser(suite.T(), suite.db, false))
		assert.ErrorContains(suite.T(), err, errors.BadRequest)
		assert.False(suite.T(), created)
	})
}

func (suite *chartControllerSuite) TestChartListAllMatching() {
	deprecated_db.Truncate(suite.T(), suite.db)
	suite.seedCharts(suite.T(), suite.db)

	suite.Run("lists all charts", func() {
		matching, err := suite.ChartController.ListAllMatching(Chart{}, 0)
		assert.NoError(suite.T(), err)
		assert.Equal(suite.T(), len(chartSeedList), len(matching))
		suite.Run("orders by latest updated", func() {
			latestUpdated := matching[0].UpdatedAt
			for _, chart := range matching {
				assert.GreaterOrEqual(suite.T(), latestUpdated, chart.UpdatedAt)
			}
		})
	})
	suite.Run("limits", func() {
		limit := 2
		matching, err := suite.ChartController.ListAllMatching(Chart{}, limit)
		assert.NoError(suite.T(), err)
		assert.Equal(suite.T(), limit, len(matching))
	})
	suite.Run("filters exactly", func() {
		matching, err := suite.ChartController.ListAllMatching(Chart{CreatableChart: datarepoChart}, 0)
		assert.NoError(suite.T(), err)
		assert.Equal(suite.T(), 1, len(matching))
		assert.Equal(suite.T(), datarepoChart.Name, matching[0].Name)
	})
	suite.Run("filters multiple", func() {
		matching, err := suite.ChartController.ListAllMatching(
			Chart{CreatableChart: CreatableChart{EditableChart: EditableChart{ChartRepo: testutils.PointerTo("terra-helm")}}}, 0)
		assert.NoError(suite.T(), err)
		assert.True(suite.T(), len(matching) > 1)
		for _, chart := range matching {
			assert.Equal(suite.T(), testutils.PointerTo("terra-helm"), chart.ChartRepo)
		}
	})
	suite.Run("none is an empty list, not null", func() {
		matching, err := suite.ChartController.ListAllMatching(
			Chart{CreatableChart: CreatableChart{Name: "blah"}}, 0)
		assert.NoError(suite.T(), err)
		assert.NotNil(suite.T(), matching)
		assert.Empty(suite.T(), matching)
	})
}

func (suite *chartControllerSuite) TestChartGet() {
	deprecated_db.Truncate(suite.T(), suite.db)
	suite.seedCharts(suite.T(), suite.db)

	suite.Run("successfully", func() {
		byName, err := suite.ChartController.Get(yaleChart.Name)
		assert.NoError(suite.T(), err)
		assert.True(suite.T(), byName.ID > 0)
		byID, err := suite.ChartController.Get(fmt.Sprintf("%d", byName.ID))
		assert.NoError(suite.T(), err)
		assert.Equal(suite.T(), yaleChart.Name, byID.Name)
	})
	suite.Run("unsuccessfully for non-present", func() {
		_, err := suite.ChartController.Get("foobar")
		assert.ErrorContains(suite.T(), err, errors.NotFound)
	})
	suite.Run("unsuccessfully for invalid selector", func() {
		_, err := suite.ChartController.Get("something obviously invalid")
		assert.ErrorContains(suite.T(), err, errors.BadRequest)
	})
}

func (suite *chartControllerSuite) TestChartGetOtherValidSelectors() {
	deprecated_db.Truncate(suite.T(), suite.db)
	suite.seedCharts(suite.T(), suite.db)

	suite.Run("successfully", func() {
		selectors, err := suite.ChartController.GetOtherValidSelectors(yaleChart.Name)
		assert.NoError(suite.T(), err)
		assert.Equal(suite.T(), 2, len(selectors))
		assert.Equal(suite.T(), yaleChart.Name, selectors[0])
	})
	suite.Run("unsuccessfully for not found", func() {
		_, err := suite.ChartController.GetOtherValidSelectors("foobar")
		assert.ErrorContains(suite.T(), err, errors.NotFound)
	})
	suite.Run("unsuccessfully for invalid", func() {
		_, err := suite.ChartController.GetOtherValidSelectors("something obviously invalid")
		assert.ErrorContains(suite.T(), err, errors.BadRequest)
	})
}

func (suite *chartControllerSuite) TestChartEdit() {
	suite.Run("successfully", func() {
		deprecated_db.Truncate(suite.T(), suite.db)
		suite.seedCharts(suite.T(), suite.db)

		before, err := suite.ChartController.Get(yaleChart.Name)
		assert.NoError(suite.T(), err)
		assert.Equal(suite.T(), yaleChart.AppImageGitMainBranch, before.AppImageGitMainBranch)
		assert.Empty(suite.T(), before.Description)
		assert.Empty(suite.T(), before.PlaybookURL)
		newValue := testutils.PointerTo("new")
		edited, err := suite.ChartController.Edit(yaleChart.Name, EditableChart{AppImageGitMainBranch: newValue, Description: newValue, PlaybookURL: newValue}, generateUser(suite.T(), suite.db, false))
		assert.NoError(suite.T(), err)
		assert.Equal(suite.T(), newValue, edited.AppImageGitMainBranch)
		assert.Equal(suite.T(), newValue, edited.Description)
		assert.Equal(suite.T(), newValue, edited.PlaybookURL)
		after, err := suite.ChartController.Get(yaleChart.Name)
		assert.NoError(suite.T(), err)
		assert.Equal(suite.T(), newValue, after.AppImageGitMainBranch)
		assert.Equal(suite.T(), newValue, after.Description)
		assert.Equal(suite.T(), newValue, after.PlaybookURL)
	})
	suite.Run("unsuccessfully if invalid", func() {
		deprecated_db.Truncate(suite.T(), suite.db)
		suite.seedCharts(suite.T(), suite.db)

		_, err := suite.ChartController.Edit(yaleLibChart.Name, EditableChart{ChartRepo: testutils.PointerTo("")}, generateUser(suite.T(), suite.db, false))
		assert.ErrorContains(suite.T(), err, errors.BadRequest)
	})
}

func (suite *chartControllerSuite) TestChartDelete() {
	suite.Run("successfully", func() {
		deprecated_db.Truncate(suite.T(), suite.db)
		suite.seedCharts(suite.T(), suite.db)

		deleted, err := suite.ChartController.Delete(leonardoChart.Name, generateUser(suite.T(), suite.db, false))
		assert.NoError(suite.T(), err)
		assert.Equal(suite.T(), leonardoChart.Name, deleted.Name)
		_, err = suite.ChartController.Get(leonardoChart.Name)
		assert.ErrorContains(suite.T(), err, errors.NotFound)
		suite.Run("sql constraints ignore soft deletion", func() {
			_, created, err := suite.ChartController.Create(leonardoChart, generateUser(suite.T(), suite.db, false))
			assert.ErrorContains(suite.T(), err, errors.BadRequest)
			assert.ErrorContains(suite.T(), err, "Contact DevOps")
			assert.False(suite.T(), created)
		})
	})
}
