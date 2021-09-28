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

func TestRunListServiceInstancesIntegrationSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	suite.Run(t, new(ListServiceInstancesIntegrationSuite))
}

type ListServiceInstancesIntegrationSuite struct {
	suite.Suite
	app                      *testApplication
	expectedServiceInstances []ServiceInstance
	ctx                      context.Context
}

func (suite *ListServiceInstancesIntegrationSuite) SetupSuite() {
	// connect to the the db and create a test application instance to be used in the suite
	t := suite.T()

	suite.ctx = context.Background()
	app := initTestApp(suite.ctx, t)
	suite.app = app
	suite.expectedServiceInstances = SeedServiceInstances(t, suite.app.db)
}

func (suite *ListServiceInstancesIntegrationSuite) TearDownSuite() {
	testutils.Cleanup(suite.T(), suite.app.db)
}

func (suite *ListServiceInstancesIntegrationSuite) Test_Integration_ListServiceInstances() {
	assert := suite.Assert()

	serviceInstances, err := suite.app.serviceInstances.ListAll()
	assert.NoError(err)

	assert.ElementsMatch(suite.expectedServiceInstances, serviceInstances)

	// check serialzied responses
	assert.ElementsMatch(suite.app.serviceInstances.Serialize(suite.expectedServiceInstances...), suite.app.serviceInstances.Serialize(serviceInstances...))
}

func Test_ListServiceInstancesError(t *testing.T) {
	targetError := errors.New("some internal error")
	controller := setupMockController(t, []ServiceInstance{}, targetError, "listAll")
	_, err := controller.ListAll()
	assert.ErrorIs(t, err, targetError, "expected an internal error from DB layer, received some other error")
}

// func Test_Integration_CreateServiceInstance(t *testing.T) {
// 	if testing.Short() {
// 		t.Skip("skipping integration test")
// 	}
// 	app := initTestApp(t)
// 	defer testutils.Cleanup(t, app.db)

// 	t.Run("successful create new service instance with pre-existing service and environment", func(t *testing.T) {

// 		app.seedServicesAndEnvironments(t)

// 		result, err := app.serviceInstances.CreateNew("cromwell", "dev")
// 		assert.NoError(t, err)

// 		assert.Equal(t, "dev", result.Environment.Name)
// 		assert.Equal(t, "cromwell", result.Service.Name)
// 	})
// }

type testApplication struct {
	serviceInstances *ServiceInstanceController
	services         *services.ServiceController
	environments     *environments.EnvironmentController
	db               *gorm.DB
}

func initTestApp(ctx context.Context, t *testing.T) *testApplication {
	dbConn := testutils.ConnectAndMigrate(t)
	dbConn = dbConn.WithContext(ctx)
	return &testApplication{
		serviceInstances: NewServiceInstanceController(dbConn),
		services:         services.NewController(dbConn),
		environments:     environments.NewController(dbConn),
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
