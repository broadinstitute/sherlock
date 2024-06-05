// Code generated by mockery v2.32.4. DO NOT EDIT.

package github_mocks

import mock "github.com/stretchr/testify/mock"

// MockMockableIssueImportClient is an autogenerated mock type for the mockableIssueImportClient type
type MockMockableIssueImportClient struct {
	mock.Mock
}

type MockMockableIssueImportClient_Expecter struct {
	mock *mock.Mock
}

func (_m *MockMockableIssueImportClient) EXPECT() *MockMockableIssueImportClient_Expecter {
	return &MockMockableIssueImportClient_Expecter{mock: &_m.Mock}
}

// NewMockMockableIssueImportClient creates a new instance of MockMockableIssueImportClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockMockableIssueImportClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockMockableIssueImportClient {
	mock := &MockMockableIssueImportClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}