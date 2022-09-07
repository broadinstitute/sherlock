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
	config.LoadTestConfig(suite.T())
	suite.db = db.ConnectAndConfigureFromTest(suite.T())
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
		},
	}
	datarepoChart = CreatableChart{
		Name: "datarepo",
		EditableChart: EditableChart{
			ChartRepo:             testutils.PointerTo("datarepo-helm"),
			AppImageGitRepo:       testutils.PointerTo("DataBiosphere/jade-data-repo"),
			AppImageGitMainBranch: testutils.PointerTo("develop"),
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
	chartSeedList = []CreatableChart{leonardoChart, datarepoChart, honeycombChart, terraClusterStorageChart, yaleChart, ingressChart, yaleLibChart}
)

func (controllerSet *ControllerSet) seedCharts(t *testing.T) {
	for _, creatable := range chartSeedList {
		if _, _, err := controllerSet.ChartController.Create(creatable, auth.GenerateUser(t, false)); err != nil {
			t.Errorf("error seeding chart %s: %v", creatable.Name, err)
		}
	}
}

//
// Controller tests
//

func (suite *chartControllerSuite) TestChartCreate() {
	suite.Run("can create a new chart", func() {
		db.Truncate(suite.T(), suite.db)

		chart, created, err := suite.ChartController.Create(leonardoChart, auth.GenerateUser(suite.T(), false))
		assert.NoError(suite.T(), err)
		assert.True(suite.T(), created)
		assert.Equal(suite.T(), leonardoChart.Name, chart.Name)
		assert.True(suite.T(), chart.ID > 0)
		suite.Run("default chart repo terra-helm", func() {
			assert.Equal(suite.T(), "terra-helm", *chart.ChartRepo)
		})
	})
	suite.Run("chart repo can be customized", func() {
		db.Truncate(suite.T(), suite.db)

		chart, created, err := suite.ChartController.Create(datarepoChart, auth.GenerateUser(suite.T(), false))
		assert.NoError(suite.T(), err)
		assert.True(suite.T(), created)
		assert.Equal(suite.T(), datarepoChart.ChartRepo, chart.ChartRepo)
	})
	suite.Run("won't create duplicates", func() {
		db.Truncate(suite.T(), suite.db)

		chart, created, err := suite.ChartController.Create(leonardoChart, auth.GenerateUser(suite.T(), false))
		assert.NoError(suite.T(), err)
		assert.True(suite.T(), created)
		assert.True(suite.T(), chart.ID > 0)
		_, created, err = suite.ChartController.Create(leonardoChart, auth.GenerateUser(suite.T(), false))
		assert.ErrorContains(suite.T(), err, errors.Conflict)
		assert.False(suite.T(), created)
	})
	suite.Run("validates incoming entries", func() {
		db.Truncate(suite.T(), suite.db)

		_, created, err := suite.ChartController.Create(CreatableChart{}, auth.GenerateUser(suite.T(), false))
		assert.ErrorContains(suite.T(), err, errors.BadRequest)
		assert.False(suite.T(), created)
	})
}

func (suite *chartControllerSuite) TestChartListAllMatching() {
	db.Truncate(suite.T(), suite.db)
	suite.seedCharts(suite.T())

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
	db.Truncate(suite.T(), suite.db)
	suite.seedCharts(suite.T())

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
	db.Truncate(suite.T(), suite.db)
	suite.seedCharts(suite.T())

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
		db.Truncate(suite.T(), suite.db)
		suite.seedCharts(suite.T())

		before, err := suite.ChartController.Get(yaleChart.Name)
		assert.NoError(suite.T(), err)
		assert.Equal(suite.T(), yaleChart.AppImageGitMainBranch, before.AppImageGitMainBranch)
		newBranch := testutils.PointerTo("new")
		edited, err := suite.ChartController.Edit(yaleChart.Name, EditableChart{AppImageGitMainBranch: newBranch}, auth.GenerateUser(suite.T(), false))
		assert.NoError(suite.T(), err)
		assert.Equal(suite.T(), newBranch, edited.AppImageGitMainBranch)
		after, err := suite.ChartController.Get(yaleChart.Name)
		assert.NoError(suite.T(), err)
		assert.Equal(suite.T(), newBranch, after.AppImageGitMainBranch)
	})
	suite.Run("unsuccessfully if invalid", func() {
		db.Truncate(suite.T(), suite.db)
		suite.seedCharts(suite.T())

		_, err := suite.ChartController.Edit(yaleLibChart.Name, EditableChart{ChartRepo: testutils.PointerTo("")}, auth.GenerateUser(suite.T(), false))
		assert.ErrorContains(suite.T(), err, errors.BadRequest)
	})
}

func (suite *chartControllerSuite) TestChartDelete() {
	suite.Run("successfully", func() {
		db.Truncate(suite.T(), suite.db)
		suite.seedCharts(suite.T())

		deleted, err := suite.ChartController.Delete(leonardoChart.Name, auth.GenerateUser(suite.T(), false))
		assert.NoError(suite.T(), err)
		assert.Equal(suite.T(), leonardoChart.Name, deleted.Name)
		_, err = suite.ChartController.Get(leonardoChart.Name)
		assert.ErrorContains(suite.T(), err, errors.NotFound)
		suite.Run("sql constraints ignore soft deletion", func() {
			_, created, err := suite.ChartController.Create(leonardoChart, auth.GenerateUser(suite.T(), false))
			assert.ErrorContains(suite.T(), err, errors.BadRequest)
			assert.ErrorContains(suite.T(), err, "Contact DevOps")
			assert.False(suite.T(), created)
		})
	})
}
