package builds

import (
	"fmt"
	"testing"
	"time"

	"github.com/broadinstitute/sherlock/internal/testutils"
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
		VersionString: "docker.io/broad/rawls:12.3.",
		CommitSha:     "asdfewrf",
		BuildURL:      "https://jenkins.job/1",
		BuiltAt:       time.Now(),
		ServiceName:   "rawls",
		ServiceRepo:   "github.com/broadinstitute/rawls",
	}
	suite.app = initTestApp(suite.T())
}

func (suite *BuildsIntegrationTestSuite) TearDownSuite() {
	// ensure we clean the db at the very end
	testutils.Cleanup(suite.T(), suite.app.db)
}

func (suite *BuildsIntegrationTestSuite) TestCreateBuild() {
	suite.Run("creates a new build", func() {
		testutils.Cleanup(suite.T(), suite.app.db)

		newBuild := suite.goodCreateBuildRequest

		build, err := suite.app.builds.CreateNew(newBuild)
		suite.Assert().NoError(err)

		suite.Assert().Equal(newBuild.VersionString, build.VersionString)
	})

	suite.Run("fails with empty create request", func() {
		testutils.Cleanup(suite.T(), suite.app.db)

		newBuild := CreateBuildRequest{}

		_, err := suite.app.builds.CreateNew(newBuild)

		suite.Require().Error(err)
	})

	suite.Run("fails on non-unique version string", func() {
		testutils.Cleanup(suite.T(), suite.app.db)

		// create a valid build
		_, err := suite.app.builds.CreateNew(suite.goodCreateBuildRequest)
		suite.Assert().NoError(err)

		// try to create another build with the same version string
		_, err = suite.app.builds.CreateNew(suite.goodCreateBuildRequest)
		suite.Assert().Error(err)
	})
}

func (suite *BuildsIntegrationTestSuite) TestGetByID() {
	suite.Run("fails with non-existent id", func() {
		testutils.Cleanup(suite.T(), suite.app.db)

		_, err := suite.app.builds.GetByID(23)
		suite.Require().ErrorIs(err, ErrBuildNotFound)

	})

	suite.Run("retrives an existing build", func() {
		testutils.Cleanup(suite.T(), suite.app.db)

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
		testutils.Cleanup(suite.T(), suite.app.db)

		// create a build instance to look up
		existingBuild, err := suite.app.builds.CreateNew(suite.goodCreateBuildRequest)
		suite.Require().NoError(err)

		result, err := suite.app.builds.GetByVersionString(suite.goodCreateBuildRequest.VersionString)
		suite.Assert().NoError(err)

		// make sure the id's match
		suite.Assert().Equal(existingBuild.ID, result.ID)
		fmt.Println(result)
	})

	suite.Run("errors not found for non-existent version string", func() {
		testutils.Cleanup(suite.T(), suite.app.db)

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
	return &testApplication{
		builds: NewController(dbConn),
		db:     dbConn,
	}
}
