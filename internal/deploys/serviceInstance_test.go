package deploys

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/broadinstitute/sherlock/internal/controllers"
	"github.com/broadinstitute/sherlock/internal/environments"
	"github.com/broadinstitute/sherlock/internal/models"
	"github.com/broadinstitute/sherlock/internal/services"
	"github.com/broadinstitute/sherlock/internal/testutils"
	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type ServiceInstanceIntegrationTestSuite struct {
	suite.Suite
	app                *testApplication
	goodEnvironmentReq environments.CreateEnvironmentRequest
	goodServiceReq     services.CreateServiceRequest
	goodClusterReq     models.CreateClusterRequest
}

func TestServiceInstanceIntegrationSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	suite.Run(t, new(ServiceInstanceIntegrationTestSuite))
}

func (suite *ServiceInstanceIntegrationTestSuite) SetupTest() {
	suite.app = initTestApp(suite.T())
	suite.goodEnvironmentReq = environments.CreateEnvironmentRequest{
		Name: faker.UUIDHyphenated(),
	}

	suite.goodServiceReq = services.CreateServiceRequest{
		Name:    faker.UUIDHyphenated(),
		RepoURL: faker.URL(),
	}

	suite.goodClusterReq = models.CreateClusterRequest{
		Name: faker.Word(),
	}
}

func (suite *ServiceInstanceIntegrationTestSuite) TearDownTest() {
	// each test runs in its own isolated transaction
	// this ensures we cleanup after each test as it completes
	suite.app.db.Rollback()
}

type testApplication struct {
	serviceInstances *ServiceInstanceController
	clusterInstances *controllers.ClusterController
	db               *gorm.DB
}

func initTestApp(t *testing.T) *testApplication {
	dbConn := testutils.ConnectAndMigrate(t)
	// ensures each test will run in it's own isolated transaction
	// The transaction will be rolled back after each test
	// regardless of pass or fail
	dbConn = dbConn.Begin()
	return &testApplication{
		serviceInstances: NewServiceInstanceController(dbConn),
		clusterInstances: controllers.NewClusterController(dbConn),
		db:               dbConn,
	}
}

func (suite *ServiceInstanceIntegrationTestSuite) TestListServiceInstancesError() {
	targetError := errors.New("some internal error")
	controller := setupMockController(suite.T(), []ServiceInstance{}, targetError, "listAll")
	_, err := controller.ListAll()
	suite.Assert().ErrorIs(err, targetError, "expected an internal error from DB layer, received some other error")
}

func (suite *ServiceInstanceIntegrationTestSuite) TestCreateServiceInstance() {
	// pretest setup
	// prepoulate an environment
	preExistingEnv, err := suite.app.serviceInstances.environments.CreateNew(suite.goodEnvironmentReq)
	suite.Require().NoError(err)

	// pre-populate an existing service
	preExistingService, err := suite.app.serviceInstances.services.CreateNew(suite.goodServiceReq)
	suite.Require().NoError(err)

	suite.Run("creates association between existing service, environment, and cluster", func() {
		testutils.Cleanup(suite.T(), suite.app.db)

		// pre-populate an environment
		preExistingEnv, err := suite.app.serviceInstances.environments.CreateNew(suite.goodEnvironmentReq)
		suite.Require().NoError(err)

		// pre-populate an existing cluster
		preExistingCluster, err := suite.app.clusterInstances.CreateNew(suite.goodClusterReq)
		suite.Require().NoError(err)

		// attempt to create a service instance from the above
		newServiceInstanceRequest := CreateServiceInstanceRequest{
			EnvironmentName: preExistingEnv.Name,
			ServiceName:     preExistingService.Name,
			ClusterName:     preExistingCluster.Name,
		}
		requestJson, _ := json.Marshal(newServiceInstanceRequest)
		suite.T().Log(string(requestJson))
		result, err := suite.app.serviceInstances.CreateNew(newServiceInstanceRequest)
		suite.Require().NoError(err)

		suite.Assert().Equal(preExistingService.ID, result.ServiceID)
		suite.Assert().Equal(preExistingEnv.ID, result.EnvironmentID)
		suite.Assert().Equal(preExistingCluster.ID, result.Cluster.ID)
	})

	suite.Run("creates an environment if not exists", func() {

		// pre-populate an existing cluster
		preExistingCluster, err := suite.app.clusterInstances.CreateNew(suite.goodClusterReq)
		suite.Require().NoError(err)

		newServiceInstanceReq := CreateServiceInstanceRequest{
			EnvironmentName: "does-not-exist",
			ServiceName:     preExistingService.Name,
			ClusterName:     preExistingCluster.Name,
		}

		result, err := suite.app.serviceInstances.CreateNew(newServiceInstanceReq)
		suite.Require().NoError(err)

		suite.Assert().NotEqual(0, result.EnvironmentID)
	})

	suite.Run("creates a service if not exists", func() {

		// pre-populate an existing cluster
		preExistingCluster, err := suite.app.clusterInstances.CreateNew(suite.goodClusterReq)
		suite.Require().NoError(err)

		newServiceInstanceReq := CreateServiceInstanceRequest{
			EnvironmentName: preExistingEnv.Name,
			ServiceName:     "does-not-exist",
			ClusterName:     preExistingCluster.Name,
		}

		result, err := suite.app.serviceInstances.CreateNew(newServiceInstanceReq)
		suite.Require().NoError(err)

		suite.Assert().NotEqual(0, result.ServiceID)
	})

	suite.Run("cannot create the same service instance twice", func() {

		// pre-populate an existing cluster
		preExistingCluster, err := suite.app.clusterInstances.CreateNew(suite.goodClusterReq)
		suite.Require().NoError(err)

		// attempt to create a service instance from the above
		newServiceInstanceReq := CreateServiceInstanceRequest{
			EnvironmentName: preExistingEnv.Name,
			ServiceName:     preExistingService.Name,
			ClusterName:     preExistingCluster.Name,
		}

		// trying to create the same service instance again should error
		_, err = suite.app.serviceInstances.CreateNew(newServiceInstanceReq)
		suite.Assert().Error(err)
	})
}

func (suite *ServiceInstanceIntegrationTestSuite) TestGetByEnvironmentAndServiceName() {
	suite.Run("returns an existing service instance", func() {

		// prepoulate an environment
		preExistingEnv, err := suite.app.serviceInstances.environments.CreateNew(suite.goodEnvironmentReq)
		suite.Require().NoError(err)

		// pre-populate an existing service
		preExistingService, err := suite.app.serviceInstances.services.CreateNew(suite.goodServiceReq)
		suite.Require().NoError(err)

		// pre-populate an existing cluster
		preExistingCluster, err := suite.app.clusterInstances.CreateNew(suite.goodClusterReq)
		suite.Require().NoError(err)

		// attempt to create a service instance from the above
		newServiceInstanceReq := CreateServiceInstanceRequest{
			EnvironmentName: preExistingEnv.Name,
			ServiceName:     preExistingService.Name,
			ClusterName:     preExistingCluster.Name,
		}

		existingServiceInstance, err := suite.app.serviceInstances.CreateNew(newServiceInstanceReq)
		suite.Require().NoError(err)

		result, err := suite.app.serviceInstances.GetByEnvironmentAndServiceName(preExistingEnv.Name, preExistingService.Name)
		suite.Require().NoError(err)

		suite.Assert().Equal(existingServiceInstance.EnvironmentID, result.EnvironmentID)
	})

	suite.Run("it returns error not found for non-existent record", func() {

		_, err := suite.app.serviceInstances.GetByEnvironmentAndServiceName("non-existent-env", "non-existent-service")
		suite.Assert().ErrorIs(err, ErrServiceInstanceNotFound)
	})
}

func setupMockController(
	t *testing.T,
	expectedServiceInstances []ServiceInstance,
	expectedError error, methodName string) *ServiceInstanceController {

	t.Helper()
	mockStore := &mockServiceInstanceStore{}
	mockStore.On(methodName).Return(expectedServiceInstances, expectedError)
	return NewMockController(mockStore)
}
