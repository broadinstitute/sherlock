package environments

import "github.com/stretchr/testify/mock"

// MockEnvironmentStore is a concrete type that
// implements environmentStore interface for use in unit tests
type MockEnvironmentStore struct {
	mock.Mock
}

func (m *MockEnvironmentStore) listAll() ([]Environment, error) {
	retVal := m.Called()
	return retVal.Get(0).([]Environment), retVal.Error(1)
}

func (m *MockEnvironmentStore) createNew(newEnvironment CreateEnvironmentRequest) (Environment, error) {
	retEnv := newEnvironment.environmentReq()
	retVal := m.Called(newEnvironment)
	return retEnv, retVal.Error(1)
}

func (m *MockEnvironmentStore) getByID(id int) (Environment, error) {
	retVal := m.Called(id)
	return retVal.Get(0).(Environment), retVal.Error(1)
}

func (m *MockEnvironmentStore) getByName(name string) (Environment, error) {
	retVal := m.Called(name)
	return retVal.Get(0).(Environment), retVal.Error(1)
}

// NewMockController returns an EnvironmentController instance with the provided mock
// of the storage layer for use in unit tests
func NewMockController(mockStore *MockEnvironmentStore) *EnvironmentController {
	return &EnvironmentController{
		store: mockStore,
	}
}
