package deploys

import (
	"database/sql"
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
		Name: faker.UUIDHyphenated(),
	}

	suite.goodServiceReq = services.CreateServiceRequest{
		Name:    faker.UUIDHyphenated(),
		RepoURL: faker.URL(),
	}
}

func (suite *ServiceInstanceIntegrationTestSuite) TearDownTest() {
	suite.app.db.Rollback()
}

type testApplication struct {
	serviceInstances *ServiceInstanceController
	db               *gorm.DB
}

func initTestApp(t *testing.T) *testApplication {
	dbConn := testutils.ConnectAndMigrate(t)
	dbConn = dbConn.Begin(&sql.TxOptions{Isolation: sql.LevelSerializable})
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
	// pretest setup
	// prepoulate an environment
	preExistingEnv, err := suite.app.serviceInstances.environments.CreateNew(suite.goodEnvironmentReq)
	suite.Require().NoError(err)

	// pre-populate an existing service
	preExistingService, err := suite.app.serviceInstances.services.CreateNew(suite.goodServiceReq)
	suite.Require().NoError(err)

	suite.Run("creates association between existing service and environment", func() {

		// attempt to create a service instance from the above
		newServiceInstanceReq := CreateServiceInstanceRequest{
			EnvironmentName: preExistingEnv.Name,
			ServiceName:     preExistingService.Name,
		}

		result, err := suite.app.serviceInstances.CreateNew(newServiceInstanceReq)
		suite.Require().NoError(err)

		suite.Assert().Equal(preExistingService.ID, result.ServiceID)
		suite.Assert().Equal(preExistingEnv.ID, result.EnvironmentID)
	})

	suite.Run("creates an environment if not exists", func() {

		newServiceInstanceReq := CreateServiceInstanceRequest{
			EnvironmentName: "does-not-exist",
			ServiceName:     preExistingService.Name,
		}

		result, err := suite.app.serviceInstances.CreateNew(newServiceInstanceReq)
		suite.Require().NoError(err)

		suite.Assert().NotEqual(0, result.EnvironmentID)
	})

	suite.Run("creates a service if not exists", func() {

		newServiceInstanceReq := CreateServiceInstanceRequest{
			EnvironmentName: preExistingEnv.Name,
			ServiceName:     "does-not-exist",
		}

		result, err := suite.app.serviceInstances.CreateNew(newServiceInstanceReq)
		suite.Require().NoError(err)

		suite.Assert().NotEqual(0, result.ServiceID)
	})

	suite.Run("cannot create the same service instance twice", func() {

		// attempt to create a service instance from the above
		newServiceInstanceReq := CreateServiceInstanceRequest{
			EnvironmentName: preExistingEnv.Name,
			ServiceName:     preExistingService.Name,
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

		// attempt to create a service instance from the above
		newServiceInstanceReq := CreateServiceInstanceRequest{
			EnvironmentName: preExistingEnv.Name,
			ServiceName:     preExistingService.Name,
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
