// Code generated by mockery v2.32.4. DO NOT EDIT.

package github_mocks

import mock "github.com/stretchr/testify/mock"

// MockMockableRepositoriesClient is an autogenerated mock type for the mockableRepositoriesClient type
type MockMockableRepositoriesClient struct {
	mock.Mock
}

type MockMockableRepositoriesClient_Expecter struct {
	mock *mock.Mock
}

func (_m *MockMockableRepositoriesClient) EXPECT() *MockMockableRepositoriesClient_Expecter {
	return &MockMockableRepositoriesClient_Expecter{mock: &_m.Mock}
}

// NewMockMockableRepositoriesClient creates a new instance of MockMockableRepositoriesClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockMockableRepositoriesClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockMockableRepositoriesClient {
	mock := &MockMockableRepositoriesClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
