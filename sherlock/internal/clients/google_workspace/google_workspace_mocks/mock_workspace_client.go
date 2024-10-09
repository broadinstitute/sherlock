// Code generated by mockery v2.32.4. DO NOT EDIT.

package google_workspace_mocks

import (
	context "context"

	admin "google.golang.org/api/admin/directory/v1"

	mock "github.com/stretchr/testify/mock"
)

// MockWorkspaceClient is an autogenerated mock type for the WorkspaceClient type
type MockWorkspaceClient struct {
	mock.Mock
}

type MockWorkspaceClient_Expecter struct {
	mock *mock.Mock
}

func (_m *MockWorkspaceClient) EXPECT() *MockWorkspaceClient_Expecter {
	return &MockWorkspaceClient_Expecter{mock: &_m.Mock}
}

// GetCurrentUsers provides a mock function with given fields: ctx, domain
func (_m *MockWorkspaceClient) GetCurrentUsers(ctx context.Context, domain string) ([]*admin.User, error) {
	ret := _m.Called(ctx, domain)

	var r0 []*admin.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]*admin.User, error)); ok {
		return rf(ctx, domain)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []*admin.User); ok {
		r0 = rf(ctx, domain)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*admin.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockWorkspaceClient_GetCurrentUsers_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetCurrentUsers'
type MockWorkspaceClient_GetCurrentUsers_Call struct {
	*mock.Call
}

// GetCurrentUsers is a helper method to define mock.On call
//   - ctx context.Context
//   - domain string
func (_e *MockWorkspaceClient_Expecter) GetCurrentUsers(ctx interface{}, domain interface{}) *MockWorkspaceClient_GetCurrentUsers_Call {
	return &MockWorkspaceClient_GetCurrentUsers_Call{Call: _e.mock.On("GetCurrentUsers", ctx, domain)}
}

func (_c *MockWorkspaceClient_GetCurrentUsers_Call) Run(run func(ctx context.Context, domain string)) *MockWorkspaceClient_GetCurrentUsers_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockWorkspaceClient_GetCurrentUsers_Call) Return(_a0 []*admin.User, _a1 error) *MockWorkspaceClient_GetCurrentUsers_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockWorkspaceClient_GetCurrentUsers_Call) RunAndReturn(run func(context.Context, string) ([]*admin.User, error)) *MockWorkspaceClient_GetCurrentUsers_Call {
	_c.Call.Return(run)
	return _c
}

// SuspendUser provides a mock function with given fields: ctx, email
func (_m *MockWorkspaceClient) SuspendUser(ctx context.Context, email string) error {
	ret := _m.Called(ctx, email)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, email)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockWorkspaceClient_SuspendUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SuspendUser'
type MockWorkspaceClient_SuspendUser_Call struct {
	*mock.Call
}

// SuspendUser is a helper method to define mock.On call
//   - ctx context.Context
//   - email string
func (_e *MockWorkspaceClient_Expecter) SuspendUser(ctx interface{}, email interface{}) *MockWorkspaceClient_SuspendUser_Call {
	return &MockWorkspaceClient_SuspendUser_Call{Call: _e.mock.On("SuspendUser", ctx, email)}
}

func (_c *MockWorkspaceClient_SuspendUser_Call) Run(run func(ctx context.Context, email string)) *MockWorkspaceClient_SuspendUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockWorkspaceClient_SuspendUser_Call) Return(_a0 error) *MockWorkspaceClient_SuspendUser_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockWorkspaceClient_SuspendUser_Call) RunAndReturn(run func(context.Context, string) error) *MockWorkspaceClient_SuspendUser_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockWorkspaceClient creates a new instance of MockWorkspaceClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockWorkspaceClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockWorkspaceClient {
	mock := &MockWorkspaceClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
