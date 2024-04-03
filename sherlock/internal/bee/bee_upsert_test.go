package bee

import (
	"testing"

	"github.com/broadinstitute/sherlock/sherlock/internal/authentication/test_users"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"gorm.io/gorm"

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
		suite.TestData.Environment_Swatomation()
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

func (suite *BeeTestSuite) TestUpdateBee() {
	suite.Run("should correctly update a Bee Environment", func() {
		beeModel := suite.TestData.Environment_Swatomation_TestBee()
		var beeEdits []models.Changeset
		beeEdits = append(beeEdits, suite.TestData.Changeset_LeonardoSwatomation_V1toV3())

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

// Common Happy Path E2E Test (of the main package)
func (suite *BeeTestSuite) TestBeeUpsert() {
	suite.Run("should return an existing Bee and update it", func() {
		var incomingChangesets []models.Changeset
		suite.TestData.Environment_Swatomation_TestBee()
		suite.TestData.AppVersion_Leonardo_V1()
		suite.TestData.AppVersion_Leonardo_V3()
		suite.TestData.ChartVersion_Leonardo_V1()
		toChartVersion := suite.TestData.ChartVersion_Leonardo_V3()
		myChartRelease := suite.TestData.ChartRelease_LeonardoSwatomation()

		incomingChangesets = append(incomingChangesets, suite.TestData.Changeset_LeonardoSwatomation_V1toV3())

		resultEnvModel, _ := getEnvByName("swatomation-test-bee", suite.DB)

		modifiedBee, _ := BeeUpsert(resultEnvModel, incomingChangesets, suite.DB)

		suite.Equal(resultEnvModel.ID, modifiedBee.ID)
		suite.Equal(myChartRelease.AppVersion, toChartVersion)
	})
}
