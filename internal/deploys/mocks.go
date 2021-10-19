package deploys

import "github.com/stretchr/testify/mock"

// mockServiceInstanceStore is a concrete type that
// implements serviceInstanceStore interface for use in unit tests
type mockServiceInstanceStore struct {
	mock.Mock
}

func (m *mockServiceInstanceStore) listAll() ([]ServiceInstance, error) {
	retVal := m.Called()
	return retVal.Get(0).([]ServiceInstance), retVal.Error(1)
}

func (m *mockServiceInstanceStore) createNew(clusterID, serviceID, environmentID int) (ServiceInstance, error) {
	retVal := m.Called(clusterID, serviceID, environmentID)
	return retVal.Get(0).(ServiceInstance), retVal.Error(1)
}

func (m *mockServiceInstanceStore) getByEnvironmentAndServiceID(environmentID, serviceID int) (ServiceInstance, error) {
	retVal := m.Called(environmentID, serviceID)
	return retVal.Get(0).(ServiceInstance), retVal.Error(1)
}

// no idea what this does, just added it cuz I messed w/ the store.
func (m *mockServiceInstanceStore) Reload(serviceInstance ServiceInstance, reloadCluster bool, reloadEnvironment bool, reloadService bool) (ServiceInstance, error) {
	retVal := m.Called(serviceInstance, reloadCluster, reloadEnvironment, reloadService)
	return retVal.Get(0).(ServiceInstance), retVal.Error(1)
}

// NewMockController returns an EnvironmentController instance with the provided mock
// of the storage layer for use in unit tests
func NewMockController(mockStore *mockServiceInstanceStore) *ServiceInstanceController {
	return &ServiceInstanceController{
		store: mockStore,
	}
}
