// Code generated by mockery v2.32.4. DO NOT EDIT.

package github_mocks

import mock "github.com/stretchr/testify/mock"

// MockMockableEnterpriseClient is an autogenerated mock type for the mockableEnterpriseClient type
type MockMockableEnterpriseClient struct {
	mock.Mock
}

type MockMockableEnterpriseClient_Expecter struct {
	mock *mock.Mock
}

func (_m *MockMockableEnterpriseClient) EXPECT() *MockMockableEnterpriseClient_Expecter {
	return &MockMockableEnterpriseClient_Expecter{mock: &_m.Mock}
}

// NewMockMockableEnterpriseClient creates a new instance of MockMockableEnterpriseClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockMockableEnterpriseClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockMockableEnterpriseClient {
	mock := &MockMockableEnterpriseClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}