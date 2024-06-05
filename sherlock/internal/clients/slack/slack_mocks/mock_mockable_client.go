// Code generated by mockery v2.32.4. DO NOT EDIT.

package slack_mocks

import (
	context "context"

	slack "github.com/slack-go/slack"
	mock "github.com/stretchr/testify/mock"
)

// MockMockableClient is an autogenerated mock type for the mockableClient type
type MockMockableClient struct {
	mock.Mock
}

type MockMockableClient_Expecter struct {
	mock *mock.Mock
}

func (_m *MockMockableClient) EXPECT() *MockMockableClient_Expecter {
	return &MockMockableClient_Expecter{mock: &_m.Mock}
}

// AddReaction provides a mock function with given fields: name, item
func (_m *MockMockableClient) AddReaction(name string, item slack.ItemRef) error {
	ret := _m.Called(name, item)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, slack.ItemRef) error); ok {
		r0 = rf(name, item)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockMockableClient_AddReaction_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddReaction'
type MockMockableClient_AddReaction_Call struct {
	*mock.Call
}

// AddReaction is a helper method to define mock.On call
//   - name string
//   - item slack.ItemRef
func (_e *MockMockableClient_Expecter) AddReaction(name interface{}, item interface{}) *MockMockableClient_AddReaction_Call {
	return &MockMockableClient_AddReaction_Call{Call: _e.mock.On("AddReaction", name, item)}
}

func (_c *MockMockableClient_AddReaction_Call) Run(run func(name string, item slack.ItemRef)) *MockMockableClient_AddReaction_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(slack.ItemRef))
	})
	return _c
}

func (_c *MockMockableClient_AddReaction_Call) Return(_a0 error) *MockMockableClient_AddReaction_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockMockableClient_AddReaction_Call) RunAndReturn(run func(string, slack.ItemRef) error) *MockMockableClient_AddReaction_Call {
	_c.Call.Return(run)
	return _c
}

// GetUserByEmailContext provides a mock function with given fields: ctx, email
func (_m *MockMockableClient) GetUserByEmailContext(ctx context.Context, email string) (*slack.User, error) {
	ret := _m.Called(ctx, email)

	var r0 *slack.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*slack.User, error)); ok {
		return rf(ctx, email)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *slack.User); ok {
		r0 = rf(ctx, email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*slack.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockMockableClient_GetUserByEmailContext_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUserByEmailContext'
type MockMockableClient_GetUserByEmailContext_Call struct {
	*mock.Call
}

// GetUserByEmailContext is a helper method to define mock.On call
//   - ctx context.Context
//   - email string
func (_e *MockMockableClient_Expecter) GetUserByEmailContext(ctx interface{}, email interface{}) *MockMockableClient_GetUserByEmailContext_Call {
	return &MockMockableClient_GetUserByEmailContext_Call{Call: _e.mock.On("GetUserByEmailContext", ctx, email)}
}

func (_c *MockMockableClient_GetUserByEmailContext_Call) Run(run func(ctx context.Context, email string)) *MockMockableClient_GetUserByEmailContext_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockMockableClient_GetUserByEmailContext_Call) Return(_a0 *slack.User, _a1 error) *MockMockableClient_GetUserByEmailContext_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockMockableClient_GetUserByEmailContext_Call) RunAndReturn(run func(context.Context, string) (*slack.User, error)) *MockMockableClient_GetUserByEmailContext_Call {
	_c.Call.Return(run)
	return _c
}

// SendMessageContext provides a mock function with given fields: ctx, channelID, options
func (_m *MockMockableClient) SendMessageContext(ctx context.Context, channelID string, options ...slack.MsgOption) (string, string, string, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, channelID)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 string
	var r1 string
	var r2 string
	var r3 error
	if rf, ok := ret.Get(0).(func(context.Context, string, ...slack.MsgOption) (string, string, string, error)); ok {
		return rf(ctx, channelID, options...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, ...slack.MsgOption) string); ok {
		r0 = rf(ctx, channelID, options...)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, ...slack.MsgOption) string); ok {
		r1 = rf(ctx, channelID, options...)
	} else {
		r1 = ret.Get(1).(string)
	}

	if rf, ok := ret.Get(2).(func(context.Context, string, ...slack.MsgOption) string); ok {
		r2 = rf(ctx, channelID, options...)
	} else {
		r2 = ret.Get(2).(string)
	}

	if rf, ok := ret.Get(3).(func(context.Context, string, ...slack.MsgOption) error); ok {
		r3 = rf(ctx, channelID, options...)
	} else {
		r3 = ret.Error(3)
	}

	return r0, r1, r2, r3
}

// MockMockableClient_SendMessageContext_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SendMessageContext'
type MockMockableClient_SendMessageContext_Call struct {
	*mock.Call
}

// SendMessageContext is a helper method to define mock.On call
//   - ctx context.Context
//   - channelID string
//   - options ...slack.MsgOption
func (_e *MockMockableClient_Expecter) SendMessageContext(ctx interface{}, channelID interface{}, options ...interface{}) *MockMockableClient_SendMessageContext_Call {
	return &MockMockableClient_SendMessageContext_Call{Call: _e.mock.On("SendMessageContext",
		append([]interface{}{ctx, channelID}, options...)...)}
}

func (_c *MockMockableClient_SendMessageContext_Call) Run(run func(ctx context.Context, channelID string, options ...slack.MsgOption)) *MockMockableClient_SendMessageContext_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]slack.MsgOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(slack.MsgOption)
			}
		}
		run(args[0].(context.Context), args[1].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockMockableClient_SendMessageContext_Call) Return(_channel string, _timestamp string, _text string, err error) *MockMockableClient_SendMessageContext_Call {
	_c.Call.Return(_channel, _timestamp, _text, err)
	return _c
}

func (_c *MockMockableClient_SendMessageContext_Call) RunAndReturn(run func(context.Context, string, ...slack.MsgOption) (string, string, string, error)) *MockMockableClient_SendMessageContext_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateMessageContext provides a mock function with given fields: ctx, channelID, timestamp, options
func (_m *MockMockableClient) UpdateMessageContext(ctx context.Context, channelID string, timestamp string, options ...slack.MsgOption) (string, string, string, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, channelID, timestamp)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 string
	var r1 string
	var r2 string
	var r3 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, ...slack.MsgOption) (string, string, string, error)); ok {
		return rf(ctx, channelID, timestamp, options...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, ...slack.MsgOption) string); ok {
		r0 = rf(ctx, channelID, timestamp, options...)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, ...slack.MsgOption) string); ok {
		r1 = rf(ctx, channelID, timestamp, options...)
	} else {
		r1 = ret.Get(1).(string)
	}

	if rf, ok := ret.Get(2).(func(context.Context, string, string, ...slack.MsgOption) string); ok {
		r2 = rf(ctx, channelID, timestamp, options...)
	} else {
		r2 = ret.Get(2).(string)
	}

	if rf, ok := ret.Get(3).(func(context.Context, string, string, ...slack.MsgOption) error); ok {
		r3 = rf(ctx, channelID, timestamp, options...)
	} else {
		r3 = ret.Error(3)
	}

	return r0, r1, r2, r3
}

// MockMockableClient_UpdateMessageContext_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateMessageContext'
type MockMockableClient_UpdateMessageContext_Call struct {
	*mock.Call
}

// UpdateMessageContext is a helper method to define mock.On call
//   - ctx context.Context
//   - channelID string
//   - timestamp string
//   - options ...slack.MsgOption
func (_e *MockMockableClient_Expecter) UpdateMessageContext(ctx interface{}, channelID interface{}, timestamp interface{}, options ...interface{}) *MockMockableClient_UpdateMessageContext_Call {
	return &MockMockableClient_UpdateMessageContext_Call{Call: _e.mock.On("UpdateMessageContext",
		append([]interface{}{ctx, channelID, timestamp}, options...)...)}
}

func (_c *MockMockableClient_UpdateMessageContext_Call) Run(run func(ctx context.Context, channelID string, timestamp string, options ...slack.MsgOption)) *MockMockableClient_UpdateMessageContext_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]slack.MsgOption, len(args)-3)
		for i, a := range args[3:] {
			if a != nil {
				variadicArgs[i] = a.(slack.MsgOption)
			}
		}
		run(args[0].(context.Context), args[1].(string), args[2].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockMockableClient_UpdateMessageContext_Call) Return(_channel string, _timestamp string, _text string, err error) *MockMockableClient_UpdateMessageContext_Call {
	_c.Call.Return(_channel, _timestamp, _text, err)
	return _c
}

func (_c *MockMockableClient_UpdateMessageContext_Call) RunAndReturn(run func(context.Context, string, string, ...slack.MsgOption) (string, string, string, error)) *MockMockableClient_UpdateMessageContext_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockMockableClient creates a new instance of MockMockableClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockMockableClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockMockableClient {
	mock := &MockMockableClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
