package services

import (
	"github.com/broadinstitute/sherlock/internal/models/v1models"
	"github.com/stretchr/testify/mock"
)

// MockServiceStore is used for mocking underlying database operations for
// services in unit tests
type MockServiceStore struct {
	mock.Mock
}

// this is boilerplate code for the testify mock library
func (m *MockServiceStore) ListAll() ([]v1models.Service, error) {
	retVal := m.Called()
	return retVal.Get(0).([]v1models.Service), retVal.Error(1)
}

func (m *MockServiceStore) CreateNew(newService v1models.CreateServiceRequest) (v1models.Service, error) {
	retService := newService.Service()
	retVal := m.Called(newService)
	return retService, retVal.Error(1)
}

func (m *MockServiceStore) GetByName(name string) (v1models.Service, error) {
	retVal := m.Called(name)
	return retVal.Get(0).(v1models.Service), retVal.Error(1)
}

// NewMockController returns a service controller that will use a customizable mock
// store for use in tests in other packages
func NewMockController(mockStore *MockServiceStore) *ServiceController {
	return &ServiceController{store: mockStore}
}
