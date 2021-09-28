package deploys

import (
	"errors"
	"fmt"
	"testing"

	"github.com/broadinstitute/sherlock/internal/environments"
	"github.com/broadinstitute/sherlock/internal/services"
	"github.com/broadinstitute/sherlock/internal/testutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func Test_Integration_ListServiceInstances(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	app := initTestApp(t)

	expectedServiceInstances, err := SeedServiceInstances(t, app.db)

	require.NoError(t, err, "unable to seed test service instance data, failing")

	t.Run("test ListAll controller method directlty", func(t *testing.T) {

		serviceInstances, err := app.serviceInstances.ListAll()
		assert.NoError(t, err)

		assert.ElementsMatch(t, expectedServiceInstances, serviceInstances)

		fmt.Printf("%#v\n", expectedServiceInstances)

		// check serialzied responses
		assert.ElementsMatch(t, app.serviceInstances.Serialize(expectedServiceInstances...), app.serviceInstances.Serialize(serviceInstances...))
	})
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

func initTestApp(t *testing.T) *testApplication {
	dbConn := testutils.ConnectAndMigrate(t)
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
