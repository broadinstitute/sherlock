// Code generated by mockery v2.32.4. DO NOT EDIT.

package intermediary_user_mocks

import (
	intermediary_user "github.com/broadinstitute/sherlock/sherlock/internal/role_propagation/intermediary_user"
	mock "github.com/stretchr/testify/mock"
)

// MockIdentifier is an autogenerated mock type for the Identifier type
type MockIdentifier struct {
	mock.Mock
}

type MockIdentifier_Expecter struct {
	mock *mock.Mock
}

func (_m *MockIdentifier) EXPECT() *MockIdentifier_Expecter {
	return &MockIdentifier_Expecter{mock: &_m.Mock}
}

// EqualTo provides a mock function with given fields: other
func (_m *MockIdentifier) EqualTo(other intermediary_user.Identifier) bool {
	ret := _m.Called(other)

	var r0 bool
	if rf, ok := ret.Get(0).(func(intermediary_user.Identifier) bool); ok {
		r0 = rf(other)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockIdentifier_EqualTo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'EqualTo'
type MockIdentifier_EqualTo_Call struct {
	*mock.Call
}

// EqualTo is a helper method to define mock.On call
//   - other intermediary_user.Identifier
func (_e *MockIdentifier_Expecter) EqualTo(other interface{}) *MockIdentifier_EqualTo_Call {
	return &MockIdentifier_EqualTo_Call{Call: _e.mock.On("EqualTo", other)}
}

func (_c *MockIdentifier_EqualTo_Call) Run(run func(other intermediary_user.Identifier)) *MockIdentifier_EqualTo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(intermediary_user.Identifier))
	})
	return _c
}

func (_c *MockIdentifier_EqualTo_Call) Return(_a0 bool) *MockIdentifier_EqualTo_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockIdentifier_EqualTo_Call) RunAndReturn(run func(intermediary_user.Identifier) bool) *MockIdentifier_EqualTo_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockIdentifier creates a new instance of MockIdentifier. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockIdentifier(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockIdentifier {
	mock := &MockIdentifier{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}