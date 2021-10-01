package builds

import (
	"testing"
	"time"

	"github.com/broadinstitute/sherlock/internal/testutils"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type BuildsIntegrationTestSuite struct {
	suite.Suite
	app *testApplication
}

func TestBuildsIntegrationSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	suite.Run(t, new(BuildsIntegrationTestSuite))
}

func (suite *BuildsIntegrationTestSuite) SetupSuite() {
	suite.app = initTestApp(suite.T())
	// ensure the db is clean before running suite
	testutils.Cleanup(suite.T(), suite.app.db)
}

func (suite *BuildsIntegrationTestSuite) TestCreateBuild() {
	suite.Run("creates a new build", func() {
		testutils.Cleanup(suite.T(), suite.app.db)

		newBuild := CreateBuildRequest{
			VersionString: "docker.io/broad/rawls:12.3.",
			CommitSha:     "asdfewrf",
			BuildURL:      "https://jenkins.job/1",
			BuiltAt:       time.Now(),
			ServiceName:   "rawls",
			ServiceRepo:   "github.com/broadinstitute/rawls",
		}

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
