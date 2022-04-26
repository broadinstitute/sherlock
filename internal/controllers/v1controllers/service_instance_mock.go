package v1controllers

import (
	"github.com/broadinstitute/sherlock/internal/models/v1models"
	"github.com/stretchr/testify/mock"
)

// MockServiceInstanceStore is a concrete type that
// implements serviceInstanceStore interface for use in unit tests
type MockServiceInstanceStore struct {
	mock.Mock
}

func (m *MockServiceInstanceStore) ListAll() ([]v1models.ServiceInstance, error) {
	retVal := m.Called()
	return retVal.Get(0).([]v1models.ServiceInstance), retVal.Error(1)
}

func (m *MockServiceInstanceStore) CreateNew(clusterID, serviceID, environmentID int) (v1models.ServiceInstance, error) {
	retVal := m.Called(clusterID, serviceID, environmentID)
	return retVal.Get(0).(v1models.ServiceInstance), retVal.Error(1)
}

func (m *MockServiceInstanceStore) GetByEnvironmentAndServiceID(environmentID, serviceID int) (v1models.ServiceInstance, error) {
	retVal := m.Called(environmentID, serviceID)
	return retVal.Get(0).(v1models.ServiceInstance), retVal.Error(1)
}

func (m *MockServiceInstanceStore) Reload(serviceInstance v1models.ServiceInstance, reloadCluster bool, reloadEnvironment bool, reloadService bool) (v1models.ServiceInstance, error) {
	retVal := m.Called(serviceInstance, reloadCluster, reloadEnvironment, reloadService)
	return retVal.Get(0).(v1models.ServiceInstance), retVal.Error(1)
}

// NewServiceInstanceMockController returns an EnvironmentController instance with the provided mock
// of the storage layer for use in unit tests
func NewServiceInstanceMockController(mockStore *MockServiceInstanceStore) *ServiceInstanceController {
	return &ServiceInstanceController{
		store: mockStore,
	}
}
