package v1controllers

import (
	"errors"
	"github.com/broadinstitute/sherlock/internal/models/v1models"
	"testing"

	"github.com/broadinstitute/sherlock/internal/testutils"
	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/suite"
)

type ServiceInstanceFunctionalTestSuite struct {
	suite.Suite
	app                *TestApplication
	goodEnvironmentReq v1models.CreateEnvironmentRequest
	goodServiceReq     v1models.CreateServiceRequest
	goodClusterReq     v1models.CreateClusterRequest
}

func TestServiceInstanceFunctionalSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping functional test")
	}
	suite.Run(t, new(ServiceInstanceFunctionalTestSuite))
}

func (suite *ServiceInstanceFunctionalTestSuite) SetupTest() {
	suite.app = initTestServiceInstanceApp(suite.T())
	suite.goodEnvironmentReq = v1models.CreateEnvironmentRequest{
		Name: faker.UUIDHyphenated(),
	}

	suite.goodServiceReq = v1models.CreateServiceRequest{
		Name:    faker.UUIDHyphenated(),
		RepoURL: faker.URL(),
	}

	suite.goodClusterReq = v1models.CreateClusterRequest{
		Name: faker.Word(),
	}
}

func (suite *ServiceInstanceFunctionalTestSuite) TearDownTest() {
	// each test runs in its own isolated transaction
	// this ensures we cleanup after each test as it completes
	suite.app.DB.Rollback()
}

// Runs before every test (but not sub-test)
func initTestServiceInstanceApp(t *testing.T) *TestApplication {
	dbConn := testutils.ConnectAndMigrate(t)
	// ensures each test will run in it's own isolated transaction
	// The transaction will be rolled back after each test
	// regardless of pass or fail
	dbConn = dbConn.Begin()
	return &TestApplication{
		ServiceInstances: NewServiceInstanceController(dbConn),
		Clusters:         NewClusterController(dbConn),
		DB:               dbConn,
	}
}

func (suite *ServiceInstanceFunctionalTestSuite) TestListServiceInstancesError() {
	targetError := errors.New("some internal error")
	controller := setupMockController(suite.T(), []v1models.ServiceInstance{}, targetError, "ListAll")
	_, err := controller.ListAll()
	suite.Assert().ErrorIs(err, targetError, "expected an internal error from DB layer, received some other error")
}

func (suite *ServiceInstanceFunctionalTestSuite) TestCreateServiceInstance() {
	suite.Run("creates association between existing service, environment, and cluster", func() {
		testutils.Cleanup(suite.T(), suite.app.DB)
		preExistingCluster, preExistingService, preExistingEnv := suite.preProvisionDependentObjects(true, true, true)

		// attempt to create a service instance from the above
		newServiceInstanceRequest := CreateServiceInstanceRequest{
			EnvironmentName: preExistingEnv.Name,
			ServiceName:     preExistingService.Name,
			ClusterName:     preExistingCluster.Name,
		}
		result, err := suite.app.ServiceInstances.CreateNew(newServiceInstanceRequest)
		suite.Require().NoError(err)

		suite.Assert().Equal(preExistingService.ID, result.ServiceID)
		suite.Assert().Equal(preExistingEnv.ID, result.EnvironmentID)
		suite.Assert().Equal(preExistingCluster.ID, result.ClusterID)
	})

	suite.Run("creates an environment if not exists", func() {
		testutils.Cleanup(suite.T(), suite.app.DB)
		preExistingCluster, preExistingService, _ := suite.preProvisionDependentObjects(true, true, false)

		newServiceInstanceReq := CreateServiceInstanceRequest{
			EnvironmentName: "does-not-exist",
			ServiceName:     preExistingService.Name,
			ClusterName:     preExistingCluster.Name,
		}

		result, err := suite.app.ServiceInstances.CreateNew(newServiceInstanceReq)
		suite.Require().NoError(err)

		suite.Assert().NotEqual(0, result.EnvironmentID)
	})

	suite.Run("creates a service if not exists", func() {
		testutils.Cleanup(suite.T(), suite.app.DB)
		preExistingCluster, _, preExistingEnv := suite.preProvisionDependentObjects(true, false, true)

		newServiceInstanceReq := CreateServiceInstanceRequest{
			EnvironmentName: preExistingEnv.Name,
			ServiceName:     "does-not-exist",
			ClusterName:     preExistingCluster.Name,
		}

		result, err := suite.app.ServiceInstances.CreateNew(newServiceInstanceReq)
		suite.Require().NoError(err)

		suite.Assert().NotEqual(0, result.ServiceID)
	})

	suite.Run("cannot create the same service instance twice", func() {
		testutils.Cleanup(suite.T(), suite.app.DB)
		preExistingCluster, preExistingService, preExistingEnv := suite.preProvisionDependentObjects(true, true, true)

		// attempt to create a service instance from the above
		newServiceInstanceReq := CreateServiceInstanceRequest{
			EnvironmentName: preExistingEnv.Name,
			ServiceName:     preExistingService.Name,
			ClusterName:     preExistingCluster.Name,
		}

		_, err := suite.app.ServiceInstances.CreateNew(newServiceInstanceReq)
		suite.Require().NoError(err)

		// trying to create the same service instance again should error
		_, err = suite.app.ServiceInstances.CreateNew(newServiceInstanceReq)
		suite.Assert().Error(err)
	})
}

func (suite *ServiceInstanceFunctionalTestSuite) TestGetByEnvironmentAndServiceName() {
	suite.Run("returns an existing service instance", func() {
		testutils.Cleanup(suite.T(), suite.app.DB)
		preExistingCluster, preExistingService, preExistingEnv := suite.preProvisionDependentObjects(true, true, true)

		// attempt to create a service instance from the above
		newServiceInstanceReq := CreateServiceInstanceRequest{
			EnvironmentName: preExistingEnv.Name,
			ServiceName:     preExistingService.Name,
			ClusterName:     preExistingCluster.Name,
		}

		existingServiceInstance, err := suite.app.ServiceInstances.CreateNew(newServiceInstanceReq)
		suite.Require().NoError(err)

		result, err := suite.app.ServiceInstances.GetByEnvironmentAndServiceName(preExistingEnv.Name, preExistingService.Name)
		suite.Require().NoError(err)

		suite.Assert().Equal(existingServiceInstance.EnvironmentID, result.EnvironmentID)
	})

	suite.Run("it returns error not found for non-existent record", func() {

		_, err := suite.app.ServiceInstances.GetByEnvironmentAndServiceName("non-existent-env", "non-existent-service")
		suite.Assert().ErrorIs(err, v1models.ErrServiceInstanceNotFound)
	})

	suite.Run("it returns error not found for non-existent record", func() {
		_, err := suite.app.ServiceInstances.GetByEnvironmentAndServiceName("", "")
		suite.Assert().ErrorIs(err, v1models.ErrServiceInstanceNotFound)
	})
}

//
// Helper Methods
//

func setupMockController(
	t *testing.T,
	expectedServiceInstances []v1models.ServiceInstance,
	expectedError error, methodName string) *ServiceInstanceController {

	t.Helper()
	mockStore := &MockServiceInstanceStore{}
	mockStore.On(methodName).Return(expectedServiceInstances, expectedError)
	return NewServiceInstanceMockController(mockStore)
}

// helper method on suite to pre-provision all the required objects for ServiceInstance to exist,
// takes a bool for each of Cluster/Service/Environment on whether to create the object or not.
func (suite *ServiceInstanceFunctionalTestSuite) preProvisionDependentObjects(makeCluster, makeService, makeEnv bool) (v1models.Cluster, v1models.Service, v1models.Environment) {
	var preExistingEnv v1models.Environment
	var preExistingService v1models.Service
	var preExistingCluster v1models.Cluster
	var err error

	if makeEnv {
		preExistingEnv, err = suite.app.ServiceInstances.environments.CreateNew(suite.goodEnvironmentReq)
		suite.Require().NoError(err)
	}

	if makeService {
		preExistingService, err = suite.app.ServiceInstances.services.CreateNew(suite.goodServiceReq)
		suite.Require().NoError(err)
	}

	if makeCluster {
		preExistingCluster, err = suite.app.Clusters.CreateNew(suite.goodClusterReq)
		suite.Require().NoError(err)
	}

	return preExistingCluster, preExistingService, preExistingEnv
}
