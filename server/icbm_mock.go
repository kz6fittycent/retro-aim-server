// Code generated by mockery v2.34.2. DO NOT EDIT.

package server

import (
	context "context"

	oscar "github.com/mkaminski/goaim/oscar"
	mock "github.com/stretchr/testify/mock"
)

// MockICBMHandler is an autogenerated mock type for the ICBMHandler type
type MockICBMHandler struct {
	mock.Mock
}

type MockICBMHandler_Expecter struct {
	mock *mock.Mock
}

func (_m *MockICBMHandler) EXPECT() *MockICBMHandler_Expecter {
	return &MockICBMHandler_Expecter{mock: &_m.Mock}
}

// ChannelMsgToHostHandler provides a mock function with given fields: ctx, sess, snacPayloadIn
func (_m *MockICBMHandler) ChannelMsgToHostHandler(ctx context.Context, sess *Session, snacPayloadIn oscar.SNAC_0x04_0x06_ICBMChannelMsgToHost) (*oscar.XMessage, error) {
	ret := _m.Called(ctx, sess, snacPayloadIn)

	var r0 *oscar.XMessage
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *Session, oscar.SNAC_0x04_0x06_ICBMChannelMsgToHost) (*oscar.XMessage, error)); ok {
		return rf(ctx, sess, snacPayloadIn)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *Session, oscar.SNAC_0x04_0x06_ICBMChannelMsgToHost) *oscar.XMessage); ok {
		r0 = rf(ctx, sess, snacPayloadIn)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*oscar.XMessage)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *Session, oscar.SNAC_0x04_0x06_ICBMChannelMsgToHost) error); ok {
		r1 = rf(ctx, sess, snacPayloadIn)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockICBMHandler_ChannelMsgToHostHandler_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ChannelMsgToHostHandler'
type MockICBMHandler_ChannelMsgToHostHandler_Call struct {
	*mock.Call
}

// ChannelMsgToHostHandler is a helper method to define mock.On call
//   - ctx context.Context
//   - sess *Session
//   - snacPayloadIn oscar.SNAC_0x04_0x06_ICBMChannelMsgToHost
func (_e *MockICBMHandler_Expecter) ChannelMsgToHostHandler(ctx interface{}, sess interface{}, snacPayloadIn interface{}) *MockICBMHandler_ChannelMsgToHostHandler_Call {
	return &MockICBMHandler_ChannelMsgToHostHandler_Call{Call: _e.mock.On("ChannelMsgToHostHandler", ctx, sess, snacPayloadIn)}
}

func (_c *MockICBMHandler_ChannelMsgToHostHandler_Call) Run(run func(ctx context.Context, sess *Session, snacPayloadIn oscar.SNAC_0x04_0x06_ICBMChannelMsgToHost)) *MockICBMHandler_ChannelMsgToHostHandler_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*Session), args[2].(oscar.SNAC_0x04_0x06_ICBMChannelMsgToHost))
	})
	return _c
}

func (_c *MockICBMHandler_ChannelMsgToHostHandler_Call) Return(_a0 *oscar.XMessage, _a1 error) *MockICBMHandler_ChannelMsgToHostHandler_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockICBMHandler_ChannelMsgToHostHandler_Call) RunAndReturn(run func(context.Context, *Session, oscar.SNAC_0x04_0x06_ICBMChannelMsgToHost) (*oscar.XMessage, error)) *MockICBMHandler_ChannelMsgToHostHandler_Call {
	_c.Call.Return(run)
	return _c
}

// ClientEventHandler provides a mock function with given fields: ctx, sess, snacPayloadIn
func (_m *MockICBMHandler) ClientEventHandler(ctx context.Context, sess *Session, snacPayloadIn oscar.SNAC_0x04_0x14_ICBMClientEvent) error {
	ret := _m.Called(ctx, sess, snacPayloadIn)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *Session, oscar.SNAC_0x04_0x14_ICBMClientEvent) error); ok {
		r0 = rf(ctx, sess, snacPayloadIn)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockICBMHandler_ClientEventHandler_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ClientEventHandler'
type MockICBMHandler_ClientEventHandler_Call struct {
	*mock.Call
}

// ClientEventHandler is a helper method to define mock.On call
//   - ctx context.Context
//   - sess *Session
//   - snacPayloadIn oscar.SNAC_0x04_0x14_ICBMClientEvent
func (_e *MockICBMHandler_Expecter) ClientEventHandler(ctx interface{}, sess interface{}, snacPayloadIn interface{}) *MockICBMHandler_ClientEventHandler_Call {
	return &MockICBMHandler_ClientEventHandler_Call{Call: _e.mock.On("ClientEventHandler", ctx, sess, snacPayloadIn)}
}

func (_c *MockICBMHandler_ClientEventHandler_Call) Run(run func(ctx context.Context, sess *Session, snacPayloadIn oscar.SNAC_0x04_0x14_ICBMClientEvent)) *MockICBMHandler_ClientEventHandler_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*Session), args[2].(oscar.SNAC_0x04_0x14_ICBMClientEvent))
	})
	return _c
}

func (_c *MockICBMHandler_ClientEventHandler_Call) Return(_a0 error) *MockICBMHandler_ClientEventHandler_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockICBMHandler_ClientEventHandler_Call) RunAndReturn(run func(context.Context, *Session, oscar.SNAC_0x04_0x14_ICBMClientEvent) error) *MockICBMHandler_ClientEventHandler_Call {
	_c.Call.Return(run)
	return _c
}

// EvilRequestHandler provides a mock function with given fields: ctx, sess, snacPayloadIn
func (_m *MockICBMHandler) EvilRequestHandler(ctx context.Context, sess *Session, snacPayloadIn oscar.SNAC_0x04_0x08_ICBMEvilRequest) (oscar.XMessage, error) {
	ret := _m.Called(ctx, sess, snacPayloadIn)

	var r0 oscar.XMessage
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *Session, oscar.SNAC_0x04_0x08_ICBMEvilRequest) (oscar.XMessage, error)); ok {
		return rf(ctx, sess, snacPayloadIn)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *Session, oscar.SNAC_0x04_0x08_ICBMEvilRequest) oscar.XMessage); ok {
		r0 = rf(ctx, sess, snacPayloadIn)
	} else {
		r0 = ret.Get(0).(oscar.XMessage)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *Session, oscar.SNAC_0x04_0x08_ICBMEvilRequest) error); ok {
		r1 = rf(ctx, sess, snacPayloadIn)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockICBMHandler_EvilRequestHandler_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'EvilRequestHandler'
type MockICBMHandler_EvilRequestHandler_Call struct {
	*mock.Call
}

// EvilRequestHandler is a helper method to define mock.On call
//   - ctx context.Context
//   - sess *Session
//   - snacPayloadIn oscar.SNAC_0x04_0x08_ICBMEvilRequest
func (_e *MockICBMHandler_Expecter) EvilRequestHandler(ctx interface{}, sess interface{}, snacPayloadIn interface{}) *MockICBMHandler_EvilRequestHandler_Call {
	return &MockICBMHandler_EvilRequestHandler_Call{Call: _e.mock.On("EvilRequestHandler", ctx, sess, snacPayloadIn)}
}

func (_c *MockICBMHandler_EvilRequestHandler_Call) Run(run func(ctx context.Context, sess *Session, snacPayloadIn oscar.SNAC_0x04_0x08_ICBMEvilRequest)) *MockICBMHandler_EvilRequestHandler_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*Session), args[2].(oscar.SNAC_0x04_0x08_ICBMEvilRequest))
	})
	return _c
}

func (_c *MockICBMHandler_EvilRequestHandler_Call) Return(_a0 oscar.XMessage, _a1 error) *MockICBMHandler_EvilRequestHandler_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockICBMHandler_EvilRequestHandler_Call) RunAndReturn(run func(context.Context, *Session, oscar.SNAC_0x04_0x08_ICBMEvilRequest) (oscar.XMessage, error)) *MockICBMHandler_EvilRequestHandler_Call {
	_c.Call.Return(run)
	return _c
}

// ParameterQueryHandler provides a mock function with given fields: _a0
func (_m *MockICBMHandler) ParameterQueryHandler(_a0 context.Context) oscar.XMessage {
	ret := _m.Called(_a0)

	var r0 oscar.XMessage
	if rf, ok := ret.Get(0).(func(context.Context) oscar.XMessage); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(oscar.XMessage)
	}

	return r0
}

// MockICBMHandler_ParameterQueryHandler_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ParameterQueryHandler'
type MockICBMHandler_ParameterQueryHandler_Call struct {
	*mock.Call
}

// ParameterQueryHandler is a helper method to define mock.On call
//   - _a0 context.Context
func (_e *MockICBMHandler_Expecter) ParameterQueryHandler(_a0 interface{}) *MockICBMHandler_ParameterQueryHandler_Call {
	return &MockICBMHandler_ParameterQueryHandler_Call{Call: _e.mock.On("ParameterQueryHandler", _a0)}
}

func (_c *MockICBMHandler_ParameterQueryHandler_Call) Run(run func(_a0 context.Context)) *MockICBMHandler_ParameterQueryHandler_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockICBMHandler_ParameterQueryHandler_Call) Return(_a0 oscar.XMessage) *MockICBMHandler_ParameterQueryHandler_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockICBMHandler_ParameterQueryHandler_Call) RunAndReturn(run func(context.Context) oscar.XMessage) *MockICBMHandler_ParameterQueryHandler_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockICBMHandler creates a new instance of MockICBMHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockICBMHandler(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockICBMHandler {
	mock := &MockICBMHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
