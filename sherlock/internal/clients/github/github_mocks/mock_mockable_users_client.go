// Code generated by mockery v2.32.4. DO NOT EDIT.

package github_mocks

import (
	context "context"

	github "github.com/google/go-github/v58/github"
	mock "github.com/stretchr/testify/mock"
)

// MockMockableUsersClient is an autogenerated mock type for the mockableUsersClient type
type MockMockableUsersClient struct {
	mock.Mock
}

type MockMockableUsersClient_Expecter struct {
	mock *mock.Mock
}

func (_m *MockMockableUsersClient) EXPECT() *MockMockableUsersClient_Expecter {
	return &MockMockableUsersClient_Expecter{mock: &_m.Mock}
}

// Get provides a mock function with given fields: ctx, user
func (_m *MockMockableUsersClient) Get(ctx context.Context, user string) (*github.User, *github.Response, error) {
	ret := _m.Called(ctx, user)

	var r0 *github.User
	var r1 *github.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*github.User, *github.Response, error)); ok {
		return rf(ctx, user)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *github.User); ok {
		r0 = rf(ctx, user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*github.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) *github.Response); ok {
		r1 = rf(ctx, user)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*github.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, string) error); ok {
		r2 = rf(ctx, user)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MockMockableUsersClient_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockMockableUsersClient_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - user string
func (_e *MockMockableUsersClient_Expecter) Get(ctx interface{}, user interface{}) *MockMockableUsersClient_Get_Call {
	return &MockMockableUsersClient_Get_Call{Call: _e.mock.On("Get", ctx, user)}
}

func (_c *MockMockableUsersClient_Get_Call) Run(run func(ctx context.Context, user string)) *MockMockableUsersClient_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockMockableUsersClient_Get_Call) Return(_a0 *github.User, _a1 *github.Response, _a2 error) *MockMockableUsersClient_Get_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *MockMockableUsersClient_Get_Call) RunAndReturn(run func(context.Context, string) (*github.User, *github.Response, error)) *MockMockableUsersClient_Get_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockMockableUsersClient creates a new instance of MockMockableUsersClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockMockableUsersClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockMockableUsersClient {
	mock := &MockMockableUsersClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
