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

func (suite *BeeTestSuite) TestUpdateBee_brokentest() {
	suite.Run("should correctly update a Bee Environment", func() {
		beeModel := suite.TestData.Environment_Swatomation_TestBee()
		var beeEdits []models.Changeset
		beeEdits = append(beeEdits, suite.TestData.Changeset_LeonardoSwatomation_TestBee_V1toV3(beeModel.ID))

		updateBee(beeEdits, suite.DB)

		suite.Equal(beeModel, beeEdits[0].NewAppVersions[0].AppVersion)
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
		var incomingChangesets []models.Changeset
		myBee := suite.TestData.Environment_Swatomation_TestBee()
		suite.TestData.AppVersion_Leonardo_V1()
		suite.TestData.AppVersion_Leonardo_V3()
		fromChartVersion := suite.TestData.ChartVersion_Leonardo_V1()
		toChartVersion := suite.TestData.ChartVersion_Leonardo_V3()
		UNUSED(toChartVersion)
		suite.TestData.ChartRelease_LeonardoSwatomation_TestBee(myBee.ID)

		var leoChart models.Chart
		_ = suite.DB.Find(&leoChart, fromChartVersion.ChartID)

		var myQuery models.ChartRelease
		myQuery.Environment = &myBee
		myQuery.Chart = &leoChart
		myQuery.EnvironmentID = &myBee.ID
		myQuery.ChartID = leoChart.ID
		//UNUSED(leoChart)

		var myChartRelease models.ChartRelease
		var myChartReleases []models.ChartRelease
		//_ = suite.DB.Preload(clause.Associations).Where(&myQuery).First(&myChartRelease).Error
		_ = suite.DB.Preload(clause.Associations).Where(&models.ChartRelease{
			EnvironmentID: &myBee.ID,
			//ChartID:       leoChart.ID,
		}).Find(&myChartReleases).Error
		ogChartReleaseVersion := myChartRelease.ChartReleaseVersion

		myChangeSet := suite.TestData.Changeset_LeonardoSwatomation_TestBee_V1toV3_factory(myChartRelease.ID)

		incomingChangesets = append(incomingChangesets, myChangeSet)

		resultEnvModel, _ := getEnvByName("swatomation-test-bee", suite.DB)

		// the actual update
		_ = updateBee(incomingChangesets, suite.DB)

		// grab the release from db again
		_ = suite.DB.Where(&myQuery).First(&myChartRelease).Error

		suite.Equal(resultEnvModel.ID, myBee.ID)
		suite.NotEqual(ogChartReleaseVersion, myChartRelease.ChartVersion)
		suite.Equal(toChartVersion.ChartVersion, *myChartRelease.ChartReleaseVersion.ChartVersionExact)
	})
}

// Common Happy Path E2E Test (of the main package)
func (suite *BeeTestSuite) TestBeeUpsert() {
	suite.Run("should return an existing Bee and update it", func() {
		var incomingChangesets []models.Changeset
		myBee := suite.TestData.Environment_Swatomation_TestBee()
		suite.TestData.AppVersion_Leonardo_V1()
		suite.TestData.AppVersion_Leonardo_V3()
		leoChart := suite.TestData.ChartVersion_Leonardo_V1().Chart
		toChartVersion := suite.TestData.ChartVersion_Leonardo_V3()

		var myQuery models.ChartRelease
		myQuery.Environment = &myBee
		myQuery.Chart = leoChart

		var myChartRelease models.ChartRelease
		_ = suite.DB.Preload(clause.Associations).Where(&myQuery).First(&myChartRelease).Error

		myChangeSet := suite.TestData.Changeset_LeonardoSwatomation_TestBee_V1toV3_factory(myChartRelease.ID)

		incomingChangesets = append(incomingChangesets, myChangeSet)

		resultEnvModel, _ := getEnvByName("swatomation-test-bee", suite.DB)

		// the actual update
		modifiedBee, _ := BeeUpsert(resultEnvModel, incomingChangesets, suite.DB)

		// grab the release from db again
		_ = suite.DB.Preload(clause.Associations).Where(&myQuery).First(&myChartRelease).Error

		suite.Equal(resultEnvModel.ID, modifiedBee.ID)
		suite.Equal(toChartVersion.ChartVersion, myChartRelease.ChartVersion.ChartVersion)
	})
}
