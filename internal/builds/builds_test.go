package builds

import (
	"database/sql"
	"testing"
	"time"

	"github.com/broadinstitute/sherlock/internal/testutils"
	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type BuildsIntegrationTestSuite struct {
	suite.Suite
	app                    *testApplication
	goodCreateBuildRequest CreateBuildRequest
}

func TestBuildsIntegrationSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	suite.Run(t, new(BuildsIntegrationTestSuite))
}

func (suite *BuildsIntegrationTestSuite) SetupTest() {
	// start a new db transaction for each test
	suite.goodCreateBuildRequest = CreateBuildRequest{
		VersionString: faker.URL(),
		CommitSha:     faker.UUIDDigit(),
		BuildURL:      faker.URL(),
		BuiltAt:       time.Now(),
		ServiceName:   faker.UUIDHyphenated(),
		ServiceRepo:   "github.com/broadinstitute/rawls",
	}
	suite.app = initTestApp(suite.T())
}

func (suite *BuildsIntegrationTestSuite) TearDownTest() {
	suite.app.db.Rollback()
}

func (suite *BuildsIntegrationTestSuite) TestCreateBuild() {
	suite.Run("creates a new build", func() {

		newBuild := suite.goodCreateBuildRequest

		build, err := suite.app.builds.CreateNew(newBuild)
		suite.Assert().NoError(err)

		suite.Assert().Equal(newBuild.VersionString, build.VersionString)
	})
}

func (suite *BuildsIntegrationTestSuite) TestCreateBuildEmptyRequest() {
	newBuild := CreateBuildRequest{}

	_, err := suite.app.builds.CreateNew(newBuild)

	suite.Require().Error(err)
}

func (suite *BuildsIntegrationTestSuite) TestCreateNonUniqueVersion() {
	// create a valid build
	_, err := suite.app.builds.CreateNew(suite.goodCreateBuildRequest)
	suite.Require().NoError(err)

	_, err = suite.app.builds.CreateNew(suite.goodCreateBuildRequest)
	suite.Assert().ErrorIs(err, ErrDuplicateVersionString)
}

func (suite *BuildsIntegrationTestSuite) TestGetByID() {
	suite.Run("fails with non-existent id", func() {

		_, err := suite.app.builds.GetByID(23)
		suite.Require().ErrorIs(err, ErrBuildNotFound)

	})

	suite.Run("retrives an existing build", func() {

		newBuild := suite.goodCreateBuildRequest
		build, err := suite.app.builds.CreateNew(newBuild)
		suite.Require().NoError(err)

		result, err := suite.app.builds.GetByID(build.ID)
		suite.Assert().NoError(err)

		suite.Assert().Equal(result.Service.Name, newBuild.ServiceName)
		suite.Assert().Equal(result.VersionString, newBuild.VersionString)
	})
}

func (suite *BuildsIntegrationTestSuite) TestGetByVersionString() {
	suite.Run("successful looks up existing build by version string", func() {

		// create a build instance to look up
		existingBuild, err := suite.app.builds.CreateNew(suite.goodCreateBuildRequest)
		suite.Require().NoError(err)

		result, err := suite.app.builds.GetByVersionString(suite.goodCreateBuildRequest.VersionString)
		suite.Assert().NoError(err)

		// make sure the id's match
		suite.Assert().Equal(existingBuild.ID, result.ID)
	})

	suite.Run("errors not found for non-existent version string", func() {

		_, err := suite.app.builds.GetByVersionString("does-not-exist")
		suite.Assert().ErrorIs(err, ErrBuildNotFound)
	})
}

type testApplication struct {
	builds *BuildController
	db     *gorm.DB
}

func initTestApp(t *testing.T) *testApplication {
	dbConn := testutils.ConnectAndMigrate(t)
	dbConn = dbConn.Begin(&sql.TxOptions{Isolation: sql.LevelSerializable})
	return &testApplication{
		builds: NewController(dbConn),
		db:     dbConn,
	}
}
