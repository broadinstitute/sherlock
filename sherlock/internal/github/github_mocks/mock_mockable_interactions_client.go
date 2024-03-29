// Code generated by mockery v2.32.4. DO NOT EDIT.

package github_mocks

import mock "github.com/stretchr/testify/mock"

// MockMockableInteractionsClient is an autogenerated mock type for the mockableInteractionsClient type
type MockMockableInteractionsClient struct {
	mock.Mock
}

type MockMockableInteractionsClient_Expecter struct {
	mock *mock.Mock
}

func (_m *MockMockableInteractionsClient) EXPECT() *MockMockableInteractionsClient_Expecter {
	return &MockMockableInteractionsClient_Expecter{mock: &_m.Mock}
}

// NewMockMockableInteractionsClient creates a new instance of MockMockableInteractionsClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockMockableInteractionsClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockMockableInteractionsClient {
	mock := &MockMockableInteractionsClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
