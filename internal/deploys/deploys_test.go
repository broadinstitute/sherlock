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
		existingServiceInstance, err := suite.app.deploys.serviceInstances.CreateNew(existingServiceInstanceReq)
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

	suite.Run("creates service instance if not exists", func() {
		testutils.Cleanup(suite.T(), suite.app.db)

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

		result, err := suite.app.deploys.CreateNew(newDeployReq)
		suite.Assert().NoError(err)

		// assert the deploy contains expected info from the pre-existing service instance and build
		suite.Assert().Equal(existingBuildReq.ServiceName, result.Build.Service.Name)
		// make sure both build and service instance reference the same service
		suite.Assert().Equal(result.Build.Service.ID, result.ServiceInstance.Service.ID)
		suite.Assert().Equal(existingBuildReq.VersionString, result.Build.VersionString)
		suite.Assert().Equal(newDeployReq.EnvironmentName, result.ServiceInstance.Environment.Name)
		suite.Assert().Equal(newDeployReq.ServiceName, result.ServiceInstance.Service.Name)
	})

	// there should never be a situation where sherlock tries to register a deploy
	// of a build that doesn't already exist, so this should error
	suite.Run("fails if build doesn't exist", func() {
		testutils.Cleanup(suite.T(), suite.app.db)

		newDeployReq := CreateDeployRequest{
			EnvironmentName:    faker.Word(),
			ServiceName:        faker.Word(),
			BuildVersionString: faker.URL(),
		}

		_, err := suite.app.deploys.CreateNew(newDeployReq)
		suite.Assert().ErrorIs(err, builds.ErrBuildNotFound)
	})
}

func (suite *DeployIntegrationTestSuite) TestGetDeploysByServiceAndEnvironment() {
	suite.Run("returns a single deploy", func() {
		testutils.Cleanup(suite.T(), suite.app.db)

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
		testutils.Cleanup(suite.T(), suite.app.db)

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

		result, err := suite.app.deploys.GetDeploysByEnvironmentAndService(serviceName, environmentName)
		suite.Assert().NoError(err)
		suite.Assert().Equal(5, len(result))

		// make sure all the results are from the same service instance
		for _, deploy := range result {
			suite.Assert().Equal(serviceName, deploy.ServiceInstance.Service.Name)
			suite.Assert().Equal(environmentName, deploy.ServiceInstance.Environment.Name)
		}
	})
}
