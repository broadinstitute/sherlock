package v1controllers

import (
	"github.com/broadinstitute/sherlock/internal/config"
	"github.com/broadinstitute/sherlock/internal/db"
	"github.com/broadinstitute/sherlock/internal/models/v1models"
	"testing"
	"time"

	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/suite"
)

type BuildsFunctionalTestSuite struct {
	suite.Suite
	app                    *TestApplication
	goodCreateBuildRequest CreateBuildRequest
}

func TestBuildsFunctionalSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping functional test")
	}
	suite.Run(t, new(BuildsFunctionalTestSuite))
}

func (suite *BuildsFunctionalTestSuite) SetupTest() {
	suite.goodCreateBuildRequest = CreateBuildRequest{
		VersionString: faker.URL(),
		CommitSha:     faker.UUIDDigit(),
		BuildURL:      faker.URL(),
		BuiltAt:       time.Now(),
		ServiceName:   faker.UUIDHyphenated(),
		ServiceRepo:   "github.com/broadinstitute/rawls",
	}
	suite.app = initTestBuildApp(suite.T())
}

func (suite *BuildsFunctionalTestSuite) TearDownTest() {
	// each test runs in its own isolated transaction
	// this ensures we cleanup after each test as it completes
	suite.app.DB.Rollback()
}

func (suite *BuildsFunctionalTestSuite) TestCreateBuild() {
	suite.Run("creates a new build", func() {

		newBuild := suite.goodCreateBuildRequest

		build, err := suite.app.Builds.CreateNew(newBuild)
		suite.Assert().NoError(err)

		suite.Assert().Equal(newBuild.VersionString, build.VersionString)
	})
}

func (suite *BuildsFunctionalTestSuite) TestCreateBuildEmptyRequest() {
	newBuild := CreateBuildRequest{}

	_, err := suite.app.Builds.CreateNew(newBuild)

	suite.Assert().ErrorIs(err, v1models.ErrBadCreateRequest)
}

func (suite *BuildsFunctionalTestSuite) TestCreateNonUniqueVersion() {
	// create a valid build
	_, err := suite.app.Builds.CreateNew(suite.goodCreateBuildRequest)
	suite.Require().NoError(err)

	_, err = suite.app.Builds.CreateNew(suite.goodCreateBuildRequest)
	suite.Assert().ErrorIs(err, v1models.ErrDuplicateVersionString)
}

func (suite *BuildsFunctionalTestSuite) TestGetByID() {
	suite.Run("fails with non-existent id", func() {

		_, err := suite.app.Builds.GetByID(23)
		suite.Require().ErrorIs(err, v1models.ErrBuildNotFound)

	})

	suite.Run("retrives an existing build", func() {

		newBuild := suite.goodCreateBuildRequest
		build, err := suite.app.Builds.CreateNew(newBuild)
		suite.Require().NoError(err)

		result, err := suite.app.Builds.GetByID(build.ID)
		suite.Assert().NoError(err)

		suite.Assert().Equal(result.Service.Name, newBuild.ServiceName)
		suite.Assert().Equal(result.VersionString, newBuild.VersionString)
	})
}

func (suite *BuildsFunctionalTestSuite) TestGetByVersionString() {
	suite.Run("successful looks up existing build by version string", func() {

		// create a build instance to look up
		existingBuild, err := suite.app.Builds.CreateNew(suite.goodCreateBuildRequest)
		suite.Require().NoError(err)

		result, err := suite.app.Builds.GetByVersionString(suite.goodCreateBuildRequest.VersionString)
		suite.Assert().NoError(err)

		// make sure the ids match
		suite.Assert().Equal(existingBuild.ID, result.ID)
	})

	suite.Run("errors not found for non-existent version string", func() {

		_, err := suite.app.Builds.GetByVersionString("does-not-exist")
		suite.Assert().ErrorIs(err, v1models.ErrBuildNotFound)
	})
}

func initTestBuildApp(t *testing.T) *TestApplication {
	config.LoadTestConfig(t)
	dbConn := db.ConnectFromTest(t)
	// ensures each test will run in it's own isolated transaction
	// The transaction will be rolled back after each test
	// regardless of pass or fail
	dbConn = dbConn.Begin()
	return &TestApplication{
		Builds: NewBuildController(dbConn),
		DB:     dbConn,
	}
}
