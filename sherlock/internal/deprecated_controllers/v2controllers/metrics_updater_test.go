package v2controllers

import (
	"context"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/broadinstitute/sherlock/sherlock/internal/deprecated_db"
	"github.com/broadinstitute/sherlock/sherlock/internal/deprecated_models/v2models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
	"testing"
)

// v2's metrics code is entirely in the model, but we aren't able to run functional tests from the model because
// we get circular imports trying to connect to the database from there (database connection code is in another package
// that itself requires the model for migrations). That's not a problem anywhere else because we run functional tests
// via controller methods anyway. There isn't a controller component to v2's metrics, but our best option to run
// a basic "does the SQL validate" test is to run it from here similar to other functional tests.
//
// In other words, I (Jack) thinks having this rudimentary test not co-located is okay for now because literally all the
// other testing is on the happy path and I don't feel like fighting Go imports here. Maybe we'll add higher-level
// metrics code down the line and this will make sense as is, or once v1 is gone maybe the database connection code
// can live in the model itself and then we can move this there.

func TestMetricsUpdaterSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping functional test")
	}
	suite.Run(t, new(metricsUpdaterSuite))
}

type metricsUpdaterSuite struct {
	suite.Suite
	db *gorm.DB
}

func (suite *metricsUpdaterSuite) SetupSuite() {
	config.LoadTestConfig()
	suite.db = deprecated_db.ConnectAndConfigureFromTest(suite.T())
}

func (suite *metricsUpdaterSuite) TearDownSuite() {
	deprecated_db.Truncate(suite.T(), suite.db)
}

func (suite *metricsUpdaterSuite) TestUpdateMetrics() {
	suite.Run("doesn't error when running custom SQL", func() {
		deprecated_db.Truncate(suite.T(), suite.db)
		err := v2models.UpdateMetrics(context.Background(), suite.db)
		assert.NoError(suite.T(), err)
	})
}
