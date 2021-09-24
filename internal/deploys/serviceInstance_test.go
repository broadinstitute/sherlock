package deploys

import (
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
		defer testutils.Cleanup(t)

		app := initTestApp(t)
		expectedServiceInstances := app.seedServiceInstanceControllerTestData(t)

		serviceInstances, err := app.serviceInstances.ListAll()
		assert.NoError(t, err)

		assert.ElementsMatch(t, expectedServiceInstances, serviceInstances)

		// check serialzied responses
		assert.ElementsMatch(t, app.serviceInstances.Serialize(expectedServiceInstances...), app.serviceInstances.Serialize(serviceInstances...))
	})
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
