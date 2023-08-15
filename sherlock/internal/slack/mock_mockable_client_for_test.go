// Code generated by mockery v2.32.4. DO NOT EDIT.

package slack

import (
	context "context"

	slack_goslack "github.com/slack-go/slack"
	mock "github.com/stretchr/testify/mock"
)

// mockMockableClient is an autogenerated mock type for the mockableClient type
type mockMockableClient struct {
	mock.Mock
}

type mockMockableClient_Expecter struct {
	mock *mock.Mock
}

func (_m *mockMockableClient) EXPECT() *mockMockableClient_Expecter {
	return &mockMockableClient_Expecter{mock: &_m.Mock}
}

// AddReaction provides a mock function with given fields: name, item
func (_m *mockMockableClient) AddReaction(name string, item slack_goslack.ItemRef) error {
	ret := _m.Called(name, item)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, slack_goslack.ItemRef) error); ok {
		r0 = rf(name, item)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// mockMockableClient_AddReaction_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddReaction'
type mockMockableClient_AddReaction_Call struct {
	*mock.Call
}

// AddReaction is a helper method to define mock.On call
//   - name string
//   - item slack_goslack.ItemRef
func (_e *mockMockableClient_Expecter) AddReaction(name interface{}, item interface{}) *mockMockableClient_AddReaction_Call {
	return &mockMockableClient_AddReaction_Call{Call: _e.mock.On("AddReaction", name, item)}
}

func (_c *mockMockableClient_AddReaction_Call) Run(run func(name string, item slack_goslack.ItemRef)) *mockMockableClient_AddReaction_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(slack_goslack.ItemRef))
	})
	return _c
}

func (_c *mockMockableClient_AddReaction_Call) Return(_a0 error) *mockMockableClient_AddReaction_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mockMockableClient_AddReaction_Call) RunAndReturn(run func(string, slack_goslack.ItemRef) error) *mockMockableClient_AddReaction_Call {
	_c.Call.Return(run)
	return _c
}

// SendMessageContext provides a mock function with given fields: ctx, channelID, options
func (_m *mockMockableClient) SendMessageContext(ctx context.Context, channelID string, options ...slack_goslack.MsgOption) (string, string, string, error) {
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
	if rf, ok := ret.Get(0).(func(context.Context, string, ...slack_goslack.MsgOption) (string, string, string, error)); ok {
		return rf(ctx, channelID, options...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, ...slack_goslack.MsgOption) string); ok {
		r0 = rf(ctx, channelID, options...)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, ...slack_goslack.MsgOption) string); ok {
		r1 = rf(ctx, channelID, options...)
	} else {
		r1 = ret.Get(1).(string)
	}

	if rf, ok := ret.Get(2).(func(context.Context, string, ...slack_goslack.MsgOption) string); ok {
		r2 = rf(ctx, channelID, options...)
	} else {
		r2 = ret.Get(2).(string)
	}

	if rf, ok := ret.Get(3).(func(context.Context, string, ...slack_goslack.MsgOption) error); ok {
		r3 = rf(ctx, channelID, options...)
	} else {
		r3 = ret.Error(3)
	}

	return r0, r1, r2, r3
}

// mockMockableClient_SendMessageContext_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SendMessageContext'
type mockMockableClient_SendMessageContext_Call struct {
	*mock.Call
}

// SendMessageContext is a helper method to define mock.On call
//   - ctx context.Context
//   - channelID string
//   - options ...slack_goslack.MsgOption
func (_e *mockMockableClient_Expecter) SendMessageContext(ctx interface{}, channelID interface{}, options ...interface{}) *mockMockableClient_SendMessageContext_Call {
	return &mockMockableClient_SendMessageContext_Call{Call: _e.mock.On("SendMessageContext",
		append([]interface{}{ctx, channelID}, options...)...)}
}

func (_c *mockMockableClient_SendMessageContext_Call) Run(run func(ctx context.Context, channelID string, options ...slack_goslack.MsgOption)) *mockMockableClient_SendMessageContext_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]slack_goslack.MsgOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(slack_goslack.MsgOption)
			}
		}
		run(args[0].(context.Context), args[1].(string), variadicArgs...)
	})
	return _c
}

func (_c *mockMockableClient_SendMessageContext_Call) Return(_channel string, _timestamp string, _text string, err error) *mockMockableClient_SendMessageContext_Call {
	_c.Call.Return(_channel, _timestamp, _text, err)
	return _c
}

func (_c *mockMockableClient_SendMessageContext_Call) RunAndReturn(run func(context.Context, string, ...slack_goslack.MsgOption) (string, string, string, error)) *mockMockableClient_SendMessageContext_Call {
	_c.Call.Return(run)
	return _c
}

// newMockMockableClient creates a new instance of mockMockableClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockMockableClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *mockMockableClient {
	mock := &mockMockableClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}