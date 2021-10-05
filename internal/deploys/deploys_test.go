package deploys

import (
	"testing"

	"github.com/broadinstitute/sherlock/internal/builds"
	"github.com/broadinstitute/sherlock/internal/testutils"
	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type DeployIntegrationTestSuite struct {
	suite.Suite
	app *testDeployController
}

type testDeployController struct {
	deploys *DeployController
	db      *gorm.DB
}

func initTestDeployController(t *testing.T) *testDeployController {
	dbConn := testutils.ConnectAndMigrate(t)
	return &testDeployController{
		deploys: NewDeployController(dbConn),
		db:      dbConn,
	}
}

func TestDeployIntegrationSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	suite.Run(t, new(DeployIntegrationTestSuite))
}

func (suite *DeployIntegrationTestSuite) SetupSuite() {
	suite.app = initTestDeployController(suite.T())
	// ensure we start the suite with a clean db
	testutils.Cleanup(suite.T(), suite.app.db)
}

func (suite *DeployIntegrationTestSuite) TestCreateDeploy() {
	suite.Run("creates deploy from pre-existing service instance and build", func() {
		testutils.Cleanup(suite.T(), suite.app.db)

		// populate a build to deploy
		existingBuildReq := builds.CreateBuildRequest{
			VersionString: faker.URL(),
			CommitSha:     faker.UUIDDigit(),
			ServiceName:   faker.Word(),
		}
		existingBuild, err := suite.app.deploys.builds.CreateNew(existingBuildReq)
		suite.Require().NoError(err)

		// populate a service instance to deploy to
		existingServiceInstanceReq := CreateServiceInstanceRequest{
			EnvironmentName: faker.Word(),
			ServiceName:     existingBuildReq.ServiceName,
		}
		existingServiceInstance, err := suite.app.deploys.serviceInsances.CreateNew(existingServiceInstanceReq)
		suite.Require().NoError(err)

		// actually create the deploy
		newDeployReq := CreateDeployRequest{
			EnvironmentName:    existingServiceInstance.Environment.Name,
			ServiceName:        existingServiceInstance.Service.Name,
			BuildVersionString: existingBuild.VersionString,
		}

		result, err := suite.app.deploys.CreateNew(newDeployReq)
		suite.Assert().NoError(err)

		// assert the deploy contains expected info from the pre-existing service instance and build
		suite.Assert().Equal(existingBuildReq.ServiceName, result.Build.Service.Name)
		// make sure both build and service instance reference the same service
		suite.Assert().Equal(result.Build.Service.ID, result.ServiceInstance.Service.ID)
		suite.Assert().Equal(existingBuildReq.VersionString, result.Build.VersionString)
		suite.Assert().Equal(existingServiceInstanceReq.EnvironmentName, result.ServiceInstance.Environment.Name)
	})
}
