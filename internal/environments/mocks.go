package environments

import (
	"github.com/broadinstitute/sherlock/internal/models"
	"github.com/stretchr/testify/mock"
)

// MockEnvironmentStore is a concrete type that
// implements environmentStore interface for use in unit tests
type MockEnvironmentStore struct {
	mock.Mock
}

func (m *MockEnvironmentStore) ListAll() ([]models.Environment, error) {
	retVal := m.Called()
	return retVal.Get(0).([]models.Environment), retVal.Error(1)
}

func (m *MockEnvironmentStore) CreateNew(newEnvironment models.CreateEnvironmentRequest) (models.Environment, error) {
	retEnv := newEnvironment.EnvironmentReq()
	retVal := m.Called(newEnvironment)
	return retEnv, retVal.Error(1)
}

func (m *MockEnvironmentStore) GetByID(id int) (models.Environment, error) {
	retVal := m.Called(id)
	return retVal.Get(0).(models.Environment), retVal.Error(1)
}

func (m *MockEnvironmentStore) GetByName(name string) (models.Environment, error) {
	retVal := m.Called(name)
	return retVal.Get(0).(models.Environment), retVal.Error(1)
}

// NewMockController returns an EnvironmentController instance with the provided mock
// of the storage layer for use in unit tests
func NewMockController(mockStore *MockEnvironmentStore) *EnvironmentController {
	return &EnvironmentController{
		store: mockStore,
	}
}
