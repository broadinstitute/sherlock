// Code generated by mockery v2.32.4. DO NOT EDIT.

package github_mocks

import mock "github.com/stretchr/testify/mock"

// MockMockableSecretScanningClient is an autogenerated mock type for the mockableSecretScanningClient type
type MockMockableSecretScanningClient struct {
	mock.Mock
}

type MockMockableSecretScanningClient_Expecter struct {
	mock *mock.Mock
}

func (_m *MockMockableSecretScanningClient) EXPECT() *MockMockableSecretScanningClient_Expecter {
	return &MockMockableSecretScanningClient_Expecter{mock: &_m.Mock}
}

// NewMockMockableSecretScanningClient creates a new instance of MockMockableSecretScanningClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockMockableSecretScanningClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockMockableSecretScanningClient {
	mock := &MockMockableSecretScanningClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
