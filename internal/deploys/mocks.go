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

func (m *mockServiceInstanceStore) createNew(serviceID, environmentID int) (ServiceInstance, error) {
	retVal := m.Called(serviceID, environmentID)
	return retVal.Get(0).(ServiceInstance), retVal.Error(1)
}

func (m *mockServiceInstanceStore) getByEnvironmentAndServiceID(environmentID, serviceID int) (ServiceInstance, error) {
	retVal := m.Called(environmentID, serviceID)
	return retVal.Get(0).(ServiceInstance), retVal.Error(1)
}

// NewMockController returns an EnvironmentController instance with the provided mock
// of the storage layer for use in unit tests
func NewMockController(mockStore *mockServiceInstanceStore) *ServiceInstanceController {
	return &ServiceInstanceController{
		store: mockStore,
	}
}
