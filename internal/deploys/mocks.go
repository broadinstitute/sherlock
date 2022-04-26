package deploys

import (
	"github.com/broadinstitute/sherlock/internal/models/v1models"
	"github.com/stretchr/testify/mock"
)

// mockServiceInstanceStore is a concrete type that
// implements serviceInstanceStore interface for use in unit tests
type mockServiceInstanceStore struct {
	mock.Mock
}

func (m *mockServiceInstanceStore) ListAll() ([]v1models.ServiceInstance, error) {
	retVal := m.Called()
	return retVal.Get(0).([]v1models.ServiceInstance), retVal.Error(1)
}

func (m *mockServiceInstanceStore) CreateNew(clusterID, serviceID, environmentID int) (v1models.ServiceInstance, error) {
	retVal := m.Called(clusterID, serviceID, environmentID)
	return retVal.Get(0).(v1models.ServiceInstance), retVal.Error(1)
}

func (m *mockServiceInstanceStore) GetByEnvironmentAndServiceID(environmentID, serviceID int) (v1models.ServiceInstance, error) {
	retVal := m.Called(environmentID, serviceID)
	return retVal.Get(0).(v1models.ServiceInstance), retVal.Error(1)
}

func (m *mockServiceInstanceStore) Reload(serviceInstance v1models.ServiceInstance, reloadCluster bool, reloadEnvironment bool, reloadService bool) (v1models.ServiceInstance, error) {
	retVal := m.Called(serviceInstance, reloadCluster, reloadEnvironment, reloadService)
	return retVal.Get(0).(v1models.ServiceInstance), retVal.Error(1)
}

// NewMockController returns an EnvironmentController instance with the provided mock
// of the storage layer for use in unit tests
func NewMockController(mockStore *mockServiceInstanceStore) *ServiceInstanceController {
	return &ServiceInstanceController{
		store: mockStore,
	}
}
