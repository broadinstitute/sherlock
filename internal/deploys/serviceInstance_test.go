package deploys

import (
	"errors"
	"testing"

	"github.com/broadinstitute/sherlock/internal/environments"
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
		Name: faker.Word(),
	}

	suite.goodServiceReq = services.CreateServiceRequest{
		Name:    faker.Word(),
		RepoURL: faker.URL(),
	}
}

func (suite *ServiceInstanceIntegrationTestSuite) TearDownSuite() {
	testutils.Cleanup(suite.T(), suite.app.db)
}

type testApplication struct {
	serviceInstances *ServiceInstanceController
	db               *gorm.DB
}

func initTestApp(t *testing.T) *testApplication {
	dbConn := testutils.ConnectAndMigrate(t)
	return &testApplication{
		serviceInstances: NewServiceInstanceController(dbConn),
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
	suite.Run("creates association between existing service and environment", func() {
		testutils.Cleanup(suite.T(), suite.app.db)

		// prepoulate an environment
		preExistingEnv, err := suite.app.serviceInstances.environments.CreateNew(suite.goodEnvironmentReq)
		suite.Require().NoError(err)

		// pre-populate an existing service
		preExistingService, err := suite.app.serviceInstances.services.CreateNew(suite.goodServiceReq)
		suite.Require().NoError(err)

		// attempt to create a service instance from the above
		newServiceInstanceReq := CreateServiceInstanceRequest{
			EnvironmentName: preExistingEnv.Name,
			ServiceName:     preExistingService.Name,
		}

		result, err := suite.app.serviceInstances.CreateNew(newServiceInstanceReq)
		suite.Require().NoError(err)

		suite.Assert().Equal(preExistingService.Name, result.Service.Name)
		suite.Assert().Equal(preExistingEnv.Name, result.Environment.Name)
	})

	suite.Run("creates an environment if not exists", func() {
		testutils.Cleanup(suite.T(), suite.app.db)

		// pre-poulate an existing service
		preExistingService, err := suite.app.serviceInstances.services.CreateNew(suite.goodServiceReq)
		suite.Require().NoError(err)

		newServiceInstanceReq := CreateServiceInstanceRequest{
			EnvironmentName: "does-not-exist",
			ServiceName:     preExistingService.Name,
		}

		result, err := suite.app.serviceInstances.CreateNew(newServiceInstanceReq)
		suite.Require().NoError(err)

		suite.Assert().Equal(newServiceInstanceReq.EnvironmentName, result.Environment.Name)
	})

	suite.Run("creates a service if not exists", func() {
		testutils.Cleanup(suite.T(), suite.app.db)

		// pre-populate an existing environment
		preExistingEnv, err := suite.app.serviceInstances.environments.CreateNew(suite.goodEnvironmentReq)
		suite.Require().NoError(err)

		newServiceInstanceReq := CreateServiceInstanceRequest{
			EnvironmentName: preExistingEnv.Name,
			ServiceName:     "does-not-exist",
		}

		result, err := suite.app.serviceInstances.CreateNew(newServiceInstanceReq)
		suite.Require().NoError(err)

		suite.Assert().Equal(newServiceInstanceReq.ServiceName, result.Service.Name)
	})

	suite.Run("cannot create the same service instance twice", func() {
		testutils.Cleanup(suite.T(), suite.app.db)

		// prepoulate an environment
		preExistingEnv, err := suite.app.serviceInstances.environments.CreateNew(suite.goodEnvironmentReq)
		suite.Require().NoError(err)

		// pre-populate an existing service
		preExistingService, err := suite.app.serviceInstances.services.CreateNew(suite.goodServiceReq)
		suite.Require().NoError(err)

		// attempt to create a service instance from the above
		newServiceInstanceReq := CreateServiceInstanceRequest{
			EnvironmentName: preExistingEnv.Name,
			ServiceName:     preExistingService.Name,
		}

		_, err = suite.app.serviceInstances.CreateNew(newServiceInstanceReq)
		suite.Require().NoError(err)

		// trying to create the same service instance again should error
		_, err = suite.app.serviceInstances.CreateNew(newServiceInstanceReq)
		suite.Assert().Error(err)
	})
}

func (suite *ServiceInstanceIntegrationTestSuite) TestGetByEnvironmentAndServiceName() {
	suite.Run("returns an existing service instance", func() {
		testutils.Cleanup(suite.T(), suite.app.db)

		// prepoulate an environment
		preExistingEnv, err := suite.app.serviceInstances.environments.CreateNew(suite.goodEnvironmentReq)
		suite.Require().NoError(err)

		// pre-populate an existing service
		preExistingService, err := suite.app.serviceInstances.services.CreateNew(suite.goodServiceReq)
		suite.Require().NoError(err)

		// attempt to create a service instance from the above
		newServiceInstanceReq := CreateServiceInstanceRequest{
			EnvironmentName: preExistingEnv.Name,
			ServiceName:     preExistingService.Name,
		}

		existingServiceInstance, err := suite.app.serviceInstances.CreateNew(newServiceInstanceReq)
		suite.Require().NoError(err)

		result, err := suite.app.serviceInstances.GetByEnvironmentAndServiceName(preExistingEnv.Name, preExistingService.Name)
		suite.Require().NoError(err)

		suite.Assert().Equal(existingServiceInstance.Environment.Name, result.Environment.Name)
	})

	suite.Run("it returns error not found for non-existent record", func() {
		testutils.Cleanup(suite.T(), suite.app.db)

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
