package deploys

import (
	"encoding/json"
	"net/http"
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
	})

	t.Run("test get /service_instances handler", func(t *testing.T) {
		defer testutils.Cleanup(t)

		app := initTestApp(t)
		expectedServiceInstances := app.seedServiceInstanceControllerTestData(t)
		context, response := testutils.SetupTestContext()

		app.serviceInstances.getServiceInstances(context)
		assert.Equal(t, http.StatusOK, response.Code, "expected response code : %d", http.StatusOK)

		expectedServiceInstanceResponse := app.serviceInstances.Serialize(expectedServiceInstances...)

		var gotResponse []ServiceInstanceResponse
		err := json.NewDecoder(response.Body).Decode(&gotResponse)
		assert.NoError(t, err, "unexpected error decoding response body")

		assert.ElementsMatch(t, expectedServiceInstanceResponse, gotResponse, "expectation and response didn't contain same elements")
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
