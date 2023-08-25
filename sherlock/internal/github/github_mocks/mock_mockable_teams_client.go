// Code generated by mockery v2.32.4. DO NOT EDIT.

package github_mocks

import mock "github.com/stretchr/testify/mock"

// MockMockableTeamsClient is an autogenerated mock type for the mockableTeamsClient type
type MockMockableTeamsClient struct {
	mock.Mock
}

type MockMockableTeamsClient_Expecter struct {
	mock *mock.Mock
}

func (_m *MockMockableTeamsClient) EXPECT() *MockMockableTeamsClient_Expecter {
	return &MockMockableTeamsClient_Expecter{mock: &_m.Mock}
}

// NewMockMockableTeamsClient creates a new instance of MockMockableTeamsClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockMockableTeamsClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockMockableTeamsClient {
	mock := &MockMockableTeamsClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
