package deploys

import (
	"github.com/broadinstitute/sherlock/internal/models/v1_models"
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
	// ensures each test will run in it's own isolated transaction
	// The transaction will be rolled back after each test
	// regardless of pass or fail
	dbConn = dbConn.Begin()
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

func (suite *DeployIntegrationTestSuite) SetupTest() {
	// start a new db transaction for each test
	suite.app = initTestDeployController(suite.T())
}

func (suite *DeployIntegrationTestSuite) TearDownTest() {
	// each test runs in its own isolated transaction
	// this ensures we cleanup after each test as it completes
	suite.app.db.Rollback()
}

func (suite *DeployIntegrationTestSuite) TestCreateDeploy() {
	suite.Run("fails to create deploy if missing reference data", func() {
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
			EnvironmentName: faker.UUIDHyphenated(),
			ServiceName:     existingBuildReq.ServiceName,
		}
		_, err = suite.app.deploys.serviceInstances.CreateNew(existingServiceInstanceReq)
		suite.Require().NoError(err)

		// actually create the deploy
		newDeployReq := CreateDeployRequest{
			EnvironmentName:    "",
			ServiceName:        "",
			BuildVersionString: existingBuild.VersionString,
		}

		// get the current buildcount before attempting to make a new deploy
		buildCount := suite.app.db.Find(&[]v1_models.Build{}).RowsAffected

		newDeploy, err := suite.app.deploys.CreateNew(newDeployReq)
		suite.Require().Error(err)

		// make sure we didn't create any new objects and everything returned is zero-valued
		suite.Assert().Equal(0, newDeploy.BuildID)
		suite.Assert().Equal(0, newDeploy.ServiceInstanceID)
		suite.Assert().Equal(buildCount, suite.app.db.Find(&[]v1_models.Build{}).RowsAffected)
	})

	suite.Run("creates deploy from pre-existing service instance and build", func() {
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
			EnvironmentName: faker.UUIDHyphenated(),
			ServiceName:     existingBuildReq.ServiceName,
		}
		existingServiceInstance, err := suite.app.deploys.serviceInstances.CreateNew(existingServiceInstanceReq)
		suite.Require().NoError(err)

		// Reload the ServiceInstance so it has the appropriate related objects populated to reference
		existingServiceInstance, err = suite.app.deploys.serviceInstances.Reload(existingServiceInstance, true, true, true)
		suite.Require().NoError(err)

		// actually create the deploy
		newDeployReq := CreateDeployRequest{
			EnvironmentName:    existingServiceInstance.Environment.Name,
			ServiceName:        existingServiceInstance.Service.Name,
			BuildVersionString: existingBuild.VersionString,
		}

		newDeploy, err := suite.app.deploys.CreateNew(newDeployReq)
		suite.Assert().NoError(err)

		// i think it's not finding the build versionstring and making a new one.
		// make sure both build and service instance reference the same object
		suite.Assert().Equal(existingBuild.ID, newDeploy.BuildID)
		suite.Assert().Equal(existingServiceInstance.ID, newDeploy.ServiceInstanceID)
	})

	suite.Run("creates service instance if not exists", func() {

		// populate a build to deploy
		existingBuildReq := builds.CreateBuildRequest{
			VersionString: faker.URL(),
			CommitSha:     faker.UUIDDigit(),
			ServiceName:   "rawls",
		}
		existingBuild, err := suite.app.deploys.builds.CreateNew(existingBuildReq)
		suite.Require().NoError(err)

		newDeployReq := CreateDeployRequest{
			EnvironmentName:    "terra-prod",
			ServiceName:        existingBuildReq.ServiceName,
			BuildVersionString: existingBuild.VersionString,
		}

		result, err := suite.app.deploys.CreateNew(newDeployReq)
		suite.Assert().NoError(err)

		// make sure both build and service instance reference the same service
		suite.Assert().Equal(existingBuild.ID, result.BuildID)
		suite.Assert().NotEqual(0, result.ServiceInstanceID)
	})

	// there should never be a situation where sherlock tries to register a deploy
	// of a build that doesn't already exist, so this should error
	suite.Run("creates build if doesn't exist", func() {

		newDeployReq := CreateDeployRequest{
			EnvironmentName:    faker.Word(),
			ServiceName:        faker.Word(),
			BuildVersionString: faker.URL(),
		}

		deploy, err := suite.app.deploys.CreateNew(newDeployReq)
		suite.Assert().NoError(err)

		suite.Assert().Equal(newDeployReq.EnvironmentName, deploy.ServiceInstance.Environment.Name)
		suite.Assert().Equal(newDeployReq.ServiceName, deploy.ServiceInstance.Service.Name)
		suite.Assert().NotZero(deploy.ID)
		// make sure the build was created
		suite.Assert().NotZero(deploy.Build.ID)
	})
}

func (suite *DeployIntegrationTestSuite) TestGetDeploysByServiceAndEnvironment() {
	suite.Run("returns a single deploy", func() {

		// populate a build to deploy
		existingBuildReq := builds.CreateBuildRequest{
			VersionString: faker.URL(),
			CommitSha:     faker.UUIDDigit(),
			ServiceName:   faker.Word(),
		}
		existingBuild, err := suite.app.deploys.builds.CreateNew(existingBuildReq)
		suite.Require().NoError(err)

		newDeployReq := CreateDeployRequest{
			EnvironmentName:    faker.Word(),
			ServiceName:        existingBuildReq.ServiceName,
			BuildVersionString: existingBuild.VersionString,
		}

		_, err = suite.app.deploys.CreateNew(newDeployReq)
		suite.Require().NoError(err)

		result, err := suite.app.deploys.GetDeploysByEnvironmentAndService(newDeployReq.EnvironmentName, newDeployReq.ServiceName)
		suite.Assert().NoError(err)

		// expect to get one deploy back
		suite.Assert().Equal(1, len(result))
	})

	suite.Run("returns multiple deploys", func() {

		// populate multiple deploys to search for
		serviceName := faker.Word()
		environmentName := faker.Word()
		for i := 0; i < 5; i++ {

			existingBuildReq := builds.CreateBuildRequest{
				VersionString: faker.URL(),
				CommitSha:     faker.UUIDDigit(),
				ServiceName:   serviceName,
			}
			existingBuild, err := suite.app.deploys.builds.CreateNew(existingBuildReq)
			suite.Require().NoError(err)

			newDeployReq := CreateDeployRequest{
				EnvironmentName:    environmentName,
				ServiceName:        existingBuildReq.ServiceName,
				BuildVersionString: existingBuild.VersionString,
			}

			_, err = suite.app.deploys.CreateNew(newDeployReq)
			suite.Require().NoError(err)
		}

		result, err := suite.app.deploys.GetDeploysByEnvironmentAndService(environmentName, serviceName)
		suite.Assert().NoError(err)
		suite.Assert().Equal(5, len(result))

		// make sure all the results are from the same service instance
		for _, deploy := range result {
			suite.Assert().Equal(serviceName, deploy.ServiceInstance.Service.Name)
			suite.Assert().Equal(environmentName, deploy.ServiceInstance.Environment.Name)
		}
	})
}
