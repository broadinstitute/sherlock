package deploys

import (
	"errors"
	"testing"

	"github.com/broadinstitute/sherlock/internal/environments"
	"github.com/broadinstitute/sherlock/internal/services"
	"github.com/broadinstitute/sherlock/internal/testutils"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func Test_Integration_ListServiceInstances(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	t.Run("test ListAll controller method directlty", func(t *testing.T) {
		app := initTestApp(t)
		defer testutils.Cleanup(t, app.db)

		expectedServiceInstances := app.seedServiceInstanceControllerTestData(t)

		serviceInstances, err := app.serviceInstances.ListAll()
		assert.NoError(t, err)

		assert.ElementsMatch(t, expectedServiceInstances, serviceInstances)

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

type testApplication struct {
	serviceInstances *ServiceInstanceController
	db               *gorm.DB
}

func (app *testApplication) seedServiceInstanceControllerTestData(t *testing.T) []ServiceInstance {
	// seed db with data needed to construct service instances
	t.Helper()

	_, err := services.Seed(app.db)
	assert.NoError(t, err)
	_, err = environments.Seed(app.db)
	assert.NoError(t, err)
	expectedServiceInstances, err := SeedServiceInstances(app.db)
	assert.NoError(t, err)

	return expectedServiceInstances
}

func initTestApp(t *testing.T) *testApplication {
	dbConn := testutils.ConnectAndMigrate(t)
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
