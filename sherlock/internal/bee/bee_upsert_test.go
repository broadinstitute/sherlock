package bee

import (
	"testing"

	"github.com/broadinstitute/sherlock/sherlock/internal/authentication/test_users"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/stretchr/testify/suite"
)

type BeeTestSuite struct {
	suite.Suite
	test_users.TestUserHelper
	models.TestSuiteHelper
}

func TestBeeSuite(t *testing.T) {
	suite.Run(t, new(BeeTestSuite))
}

func (suite *BeeTestSuite) TestGetEnvByName() {
	suite.Run("should load get existing environment by name", func() {
		suite.TestData.Environment_Swatomation_TestBee()

		resultEnvModel, err := getEnvByName("swatomation", suite.DB)

		suite.Equal(nil, err)
		suite.Equal("swatomation", resultEnvModel.Name)
	})
}

func (suite *BeeTestSuite) TestGetBee() {
	suite.Run("should load get existing bee by name", func() {
		//suite.TestData.Environment_Swatomation()
		beeModel := suite.TestData.Environment_Swatomation_TestBee()

		resultEnvModel, err := getBee(beeModel, suite.DB)

		suite.Equal(nil, err)
		suite.Equal("swatomation-test-bee", resultEnvModel.Name)
	})
}

func (suite *BeeTestSuite) TestGetBee_nopanic() {
	suite.Run("should not panic on not-bees", func() {
		beeModel := suite.TestData.Environment_Swatomation()

		resultEnvModel, err := getBee(beeModel, suite.DB)

		suite.Equal(nil, err)
		suite.Equal("swatomation", resultEnvModel.Name)
	})
}

func (suite *BeeTestSuite) TestGetEnvByName_err() {
	suite.Run("should return empty Model and an error if no match", func() {
		resultEnvModel, err := getEnvByName("swatomation", suite.DB)

		suite.Equal(gorm.ErrRecordNotFound, err)
		suite.Equal("", resultEnvModel.Name)
	})
}

func (suite *BeeTestSuite) TestUpdateBee() {
	suite.Run("should return an existing Bee and update it", func() {
		myBee := suite.TestData.Environment_Swatomation_TestBee()
		toAppVersion := suite.TestData.AppVersion_Leonardo_V1()
		suite.TestData.AppVersion_Leonardo_V3()
		toChartVersion := suite.TestData.ChartVersion_Leonardo_V1()
		fromChartVersion := suite.TestData.ChartVersion_Leonardo_V3()
		suite.TestData.ChartRelease_LeonardoSwatomation_TestBee(myBee.ID)

		var leoChart models.Chart
		_ = suite.DB.Find(&leoChart, fromChartVersion.ChartID)

		var myQuery models.ChartRelease
		myQuery.EnvironmentID = &myBee.ID
		myQuery.ChartID = leoChart.ID

		var myChartRelease models.ChartRelease
		suite.DB.Where(&myQuery).First(&myChartRelease)
		suite.Equal("0.3.0", *myChartRelease.ChartReleaseVersion.ChartVersionExact)
		suite.Equal("v0.0.3", *myChartRelease.ChartReleaseVersion.AppVersionExact)

		// manually build Changeset test data
		var incomingChangesets []models.Changeset
		myChangeSet := suite.TestData.Changeset_LeonardoSwatomation_TestBee_V3toV1_factory(myChartRelease.ID)
		incomingChangesets = append(incomingChangesets, myChangeSet)

		// the actual update
		err := updateBee(incomingChangesets, suite.DB)

		// grab the release from db again
		var finalChartRelease models.ChartRelease
		_ = suite.DB.Preload(clause.Associations).Where(&myQuery).First(&finalChartRelease).Error

		//assert
		suite.Equal(nil, err)
		suite.NotEqual(fromChartVersion.ChartVersion, toChartVersion.ChartVersion)
		suite.Equal(toAppVersion.AppVersion, *finalChartRelease.ChartReleaseVersion.AppVersionExact)
		suite.Equal(toChartVersion.ChartVersion, *finalChartRelease.ChartReleaseVersion.ChartVersionExact)
	})
}

// Common Happy Path E2E Test (of the main package)
func (suite *BeeTestSuite) TestBeeUpsert() {
	suite.Run("should return an existing Bee and update it", func() {
		myBee := suite.TestData.Environment_Swatomation_TestBee()
		toAppVersion := suite.TestData.AppVersion_Leonardo_V1()
		suite.TestData.AppVersion_Leonardo_V3()
		toChartVersion := suite.TestData.ChartVersion_Leonardo_V1()
		fromChartVersion := suite.TestData.ChartVersion_Leonardo_V3()
		suite.TestData.ChartRelease_LeonardoSwatomation_TestBee(myBee.ID)

		var leoChart models.Chart
		_ = suite.DB.Find(&leoChart, fromChartVersion.ChartID)

		var myQuery models.ChartRelease
		myQuery.EnvironmentID = &myBee.ID
		myQuery.ChartID = leoChart.ID

		var myChartRelease models.ChartRelease
		_ = suite.DB.Where(&myQuery).First(&myChartRelease)
		suite.Equal(fromChartVersion.ChartVersion, *myChartRelease.ChartReleaseVersion.ChartVersionExact)

		// manually build Changset test data
		var incomingChangesets []models.Changeset
		myChangeSet := suite.TestData.Changeset_LeonardoSwatomation_TestBee_V3toV1_factory(myChartRelease.ID)
		incomingChangesets = append(incomingChangesets, myChangeSet)

		// the actual update
		resultBee, _ := BeeUpsert(myBee, incomingChangesets, suite.DB)

		// grab the release from db again
		var finalChartRelease models.ChartRelease
		_ = suite.DB.Where(&myQuery).First(&finalChartRelease).Error

		//assert
		suite.Equal(resultBee.ID, myBee.ID)
		suite.NotEqual(fromChartVersion.ChartVersion, toChartVersion.ChartVersion)
		suite.Equal(toAppVersion.AppVersion, *finalChartRelease.ChartReleaseVersion.AppVersionExact)
		suite.Equal(toChartVersion.ChartVersion, *finalChartRelease.ChartReleaseVersion.ChartVersionExact)
	})
}
