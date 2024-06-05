// Code generated by mockery v2.32.4. DO NOT EDIT.

package github_mocks

import mock "github.com/stretchr/testify/mock"

// MockMockablePullRequestsClient is an autogenerated mock type for the mockablePullRequestsClient type
type MockMockablePullRequestsClient struct {
	mock.Mock
}

type MockMockablePullRequestsClient_Expecter struct {
	mock *mock.Mock
}

func (_m *MockMockablePullRequestsClient) EXPECT() *MockMockablePullRequestsClient_Expecter {
	return &MockMockablePullRequestsClient_Expecter{mock: &_m.Mock}
}

// NewMockMockablePullRequestsClient creates a new instance of MockMockablePullRequestsClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockMockablePullRequestsClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockMockablePullRequestsClient {
	mock := &MockMockablePullRequestsClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}