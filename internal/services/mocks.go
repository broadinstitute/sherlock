package services

import "github.com/stretchr/testify/mock"

type MockServiceStore struct {
	mock.Mock
}

// this is boilerplate code for the testify mock library
func (m *MockServiceStore) ListAll() ([]*Service, error) {
	retVal := m.Called()
	return retVal.Get(0).([]*Service), retVal.Error(1)
}

func (m *MockServiceStore) CreateNew(newService CreateServiceRequest) (*Service, error) {
	retService := newService.service()
	retVal := m.Called(newService)
	return retService, retVal.Error(1)
}

func (m *MockServiceStore) GetByName(name string) (*Service, error) {
	retVal := m.Called(name)
	return retVal.Get(0).(*Service), retVal.Error(1)
}
