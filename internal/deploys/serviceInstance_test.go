package deploys

import (
	"context"
	"errors"
	"testing"

	"github.com/broadinstitute/sherlock/internal/environments"
	"github.com/broadinstitute/sherlock/internal/services"
	"github.com/broadinstitute/sherlock/internal/testutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

func TestRunServiceInstancesIntegrationSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	suite.Run(t, new(ServiceInstancesIntegrationSuite))
}

type ServiceInstancesIntegrationSuite struct {
	suite.Suite
	app *testApplication
	ctx context.Context
}

func (suite *ServiceInstancesIntegrationSuite) SetupSuite() {
	// connect to the the db and create a test application instance to be used in the suite
	t := suite.T()

	suite.ctx = context.Background()
	app := initTestApp(suite.ctx, t)
	suite.app = app
}

func (suite *ServiceInstancesIntegrationSuite) TearDownSuite() {
	testutils.Cleanup(suite.T(), suite.app.db)
}

// func (suite *ServiceInstancesIntegrationSuite) TestListServiceInstances() {
// 	assert := suite.Assert()

// 	serviceInstances, err := suite.app.serviceInstances.ListAll()

// 	fmt.Println(serviceInstances)
// 	assert.NoError(err)

// 	assert.ElementsMatch(suite.expectedServiceInstances, serviceInstances)

// 	// check serialzied responses
// 	serializedExpectations := suite.app.serviceInstances.Serialize(suite.expectedServiceInstances...)
// 	serializedResult := suite.app.serviceInstances.Serialize(serviceInstances...)

// 	assert.ElementsMatch(serializedExpectations, serializedResult)
// }

func Test_ListServiceInstancesError(t *testing.T) {
	targetError := errors.New("some internal error")
	controller := setupMockController(t, []ServiceInstance{}, targetError, "listAll")
	_, err := controller.ListAll()
	assert.ErrorIs(t, err, targetError, "expected an internal error from DB layer, received some other error")
}

func (suite *ServiceInstancesIntegrationSuite) TestCreateServiceInstance() {
	assert := suite.Assert()
	require := suite.Require()

	// make a service entity
	testService := services.CreateServiceRequest{
		Name:    "buffer",
		RepoURL: "https://github.com/databiosphere/buffer",
	}
	createdService, err := suite.app.serviceInstances.services.CreateNew(testService)
	require.NoError(err)

	testEnv := environments.CreateEnvironmentRequest{
		Name: "dev",
	}
	createdEnv, err := suite.app.serviceInstances.environments.CreateNew(testEnv)

	result, err := suite.app.serviceInstances.CreateNew(createdService.Name, createdEnv.Name)
	assert.NoError(err)

	assert.Equal(testEnv.Name, result.Environment.Name)
	assert.Equal(testService.Name, result.Service.Name)
}

type testApplication struct {
	serviceInstances *ServiceInstanceController
	db               *gorm.DB
}

func initTestApp(ctx context.Context, t *testing.T) *testApplication {
	dbConn := testutils.ConnectAndMigrate(t)
	// This is to associate a specific context with all db operations performed in this
	// test suite. This creates a gorm session which is useful for grouping db transactions associated
	// with this test suite and keeping them isolated from other db operations
	// https://gorm.io/docs/context.html#Continuous-session-mode
	dbConn = dbConn.WithContext(ctx)
	return &testApplication{
		serviceInstances: NewServiceInstanceController(dbConn),
		db:               dbConn,
	}
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
