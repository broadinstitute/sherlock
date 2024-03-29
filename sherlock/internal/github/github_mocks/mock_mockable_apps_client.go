// Code generated by mockery v2.32.4. DO NOT EDIT.

package github_mocks

import mock "github.com/stretchr/testify/mock"

// MockMockableAppsClient is an autogenerated mock type for the mockableAppsClient type
type MockMockableAppsClient struct {
	mock.Mock
}

type MockMockableAppsClient_Expecter struct {
	mock *mock.Mock
}

func (_m *MockMockableAppsClient) EXPECT() *MockMockableAppsClient_Expecter {
	return &MockMockableAppsClient_Expecter{mock: &_m.Mock}
}

// NewMockMockableAppsClient creates a new instance of MockMockableAppsClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockMockableAppsClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockMockableAppsClient {
	mock := &MockMockableAppsClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
