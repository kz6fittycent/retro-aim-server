// Code generated by mockery v2.40.1. DO NOT EDIT.

package handler

import (
	context "context"

	state "github.com/mk6i/retro-aim-server/state"
	mock "github.com/stretchr/testify/mock"

	wire "github.com/mk6i/retro-aim-server/wire"
)

// mockOServiceService is an autogenerated mock type for the OServiceService type
type mockOServiceService struct {
	mock.Mock
}

type mockOServiceService_Expecter struct {
	mock *mock.Mock
}

func (_m *mockOServiceService) EXPECT() *mockOServiceService_Expecter {
	return &mockOServiceService_Expecter{mock: &_m.Mock}
}

// ClientVersions provides a mock function with given fields: ctx, frame, bodyIn
func (_m *mockOServiceService) ClientVersions(ctx context.Context, frame wire.SNACFrame, bodyIn wire.SNAC_0x01_0x17_OServiceClientVersions) wire.SNACMessage {
	ret := _m.Called(ctx, frame, bodyIn)

	if len(ret) == 0 {
		panic("no return value specified for ClientVersions")
	}

	var r0 wire.SNACMessage
	if rf, ok := ret.Get(0).(func(context.Context, wire.SNACFrame, wire.SNAC_0x01_0x17_OServiceClientVersions) wire.SNACMessage); ok {
		r0 = rf(ctx, frame, bodyIn)
	} else {
		r0 = ret.Get(0).(wire.SNACMessage)
	}

	return r0
}

// mockOServiceService_ClientVersions_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ClientVersions'
type mockOServiceService_ClientVersions_Call struct {
	*mock.Call
}

// ClientVersions is a helper method to define mock.On call
//   - ctx context.Context
//   - frame wire.SNACFrame
//   - bodyIn wire.SNAC_0x01_0x17_OServiceClientVersions
func (_e *mockOServiceService_Expecter) ClientVersions(ctx interface{}, frame interface{}, bodyIn interface{}) *mockOServiceService_ClientVersions_Call {
	return &mockOServiceService_ClientVersions_Call{Call: _e.mock.On("ClientVersions", ctx, frame, bodyIn)}
}

func (_c *mockOServiceService_ClientVersions_Call) Run(run func(ctx context.Context, frame wire.SNACFrame, bodyIn wire.SNAC_0x01_0x17_OServiceClientVersions)) *mockOServiceService_ClientVersions_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(wire.SNACFrame), args[2].(wire.SNAC_0x01_0x17_OServiceClientVersions))
	})
	return _c
}

func (_c *mockOServiceService_ClientVersions_Call) Return(_a0 wire.SNACMessage) *mockOServiceService_ClientVersions_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mockOServiceService_ClientVersions_Call) RunAndReturn(run func(context.Context, wire.SNACFrame, wire.SNAC_0x01_0x17_OServiceClientVersions) wire.SNACMessage) *mockOServiceService_ClientVersions_Call {
	_c.Call.Return(run)
	return _c
}

// IdleNotification provides a mock function with given fields: ctx, sess, bodyIn
func (_m *mockOServiceService) IdleNotification(ctx context.Context, sess *state.Session, bodyIn wire.SNAC_0x01_0x11_OServiceIdleNotification) error {
	ret := _m.Called(ctx, sess, bodyIn)

	if len(ret) == 0 {
		panic("no return value specified for IdleNotification")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *state.Session, wire.SNAC_0x01_0x11_OServiceIdleNotification) error); ok {
		r0 = rf(ctx, sess, bodyIn)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// mockOServiceService_IdleNotification_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IdleNotification'
type mockOServiceService_IdleNotification_Call struct {
	*mock.Call
}

// IdleNotification is a helper method to define mock.On call
//   - ctx context.Context
//   - sess *state.Session
//   - bodyIn wire.SNAC_0x01_0x11_OServiceIdleNotification
func (_e *mockOServiceService_Expecter) IdleNotification(ctx interface{}, sess interface{}, bodyIn interface{}) *mockOServiceService_IdleNotification_Call {
	return &mockOServiceService_IdleNotification_Call{Call: _e.mock.On("IdleNotification", ctx, sess, bodyIn)}
}

func (_c *mockOServiceService_IdleNotification_Call) Run(run func(ctx context.Context, sess *state.Session, bodyIn wire.SNAC_0x01_0x11_OServiceIdleNotification)) *mockOServiceService_IdleNotification_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*state.Session), args[2].(wire.SNAC_0x01_0x11_OServiceIdleNotification))
	})
	return _c
}

func (_c *mockOServiceService_IdleNotification_Call) Return(_a0 error) *mockOServiceService_IdleNotification_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mockOServiceService_IdleNotification_Call) RunAndReturn(run func(context.Context, *state.Session, wire.SNAC_0x01_0x11_OServiceIdleNotification) error) *mockOServiceService_IdleNotification_Call {
	_c.Call.Return(run)
	return _c
}

// RateParamsQuery provides a mock function with given fields: ctx, frame
func (_m *mockOServiceService) RateParamsQuery(ctx context.Context, frame wire.SNACFrame) wire.SNACMessage {
	ret := _m.Called(ctx, frame)

	if len(ret) == 0 {
		panic("no return value specified for RateParamsQuery")
	}

	var r0 wire.SNACMessage
	if rf, ok := ret.Get(0).(func(context.Context, wire.SNACFrame) wire.SNACMessage); ok {
		r0 = rf(ctx, frame)
	} else {
		r0 = ret.Get(0).(wire.SNACMessage)
	}

	return r0
}

// mockOServiceService_RateParamsQuery_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RateParamsQuery'
type mockOServiceService_RateParamsQuery_Call struct {
	*mock.Call
}

// RateParamsQuery is a helper method to define mock.On call
//   - ctx context.Context
//   - frame wire.SNACFrame
func (_e *mockOServiceService_Expecter) RateParamsQuery(ctx interface{}, frame interface{}) *mockOServiceService_RateParamsQuery_Call {
	return &mockOServiceService_RateParamsQuery_Call{Call: _e.mock.On("RateParamsQuery", ctx, frame)}
}

func (_c *mockOServiceService_RateParamsQuery_Call) Run(run func(ctx context.Context, frame wire.SNACFrame)) *mockOServiceService_RateParamsQuery_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(wire.SNACFrame))
	})
	return _c
}

func (_c *mockOServiceService_RateParamsQuery_Call) Return(_a0 wire.SNACMessage) *mockOServiceService_RateParamsQuery_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mockOServiceService_RateParamsQuery_Call) RunAndReturn(run func(context.Context, wire.SNACFrame) wire.SNACMessage) *mockOServiceService_RateParamsQuery_Call {
	_c.Call.Return(run)
	return _c
}

// RateParamsSubAdd provides a mock function with given fields: _a0, _a1
func (_m *mockOServiceService) RateParamsSubAdd(_a0 context.Context, _a1 wire.SNAC_0x01_0x08_OServiceRateParamsSubAdd) {
	_m.Called(_a0, _a1)
}

// mockOServiceService_RateParamsSubAdd_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RateParamsSubAdd'
type mockOServiceService_RateParamsSubAdd_Call struct {
	*mock.Call
}

// RateParamsSubAdd is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 wire.SNAC_0x01_0x08_OServiceRateParamsSubAdd
func (_e *mockOServiceService_Expecter) RateParamsSubAdd(_a0 interface{}, _a1 interface{}) *mockOServiceService_RateParamsSubAdd_Call {
	return &mockOServiceService_RateParamsSubAdd_Call{Call: _e.mock.On("RateParamsSubAdd", _a0, _a1)}
}

func (_c *mockOServiceService_RateParamsSubAdd_Call) Run(run func(_a0 context.Context, _a1 wire.SNAC_0x01_0x08_OServiceRateParamsSubAdd)) *mockOServiceService_RateParamsSubAdd_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(wire.SNAC_0x01_0x08_OServiceRateParamsSubAdd))
	})
	return _c
}

func (_c *mockOServiceService_RateParamsSubAdd_Call) Return() *mockOServiceService_RateParamsSubAdd_Call {
	_c.Call.Return()
	return _c
}

func (_c *mockOServiceService_RateParamsSubAdd_Call) RunAndReturn(run func(context.Context, wire.SNAC_0x01_0x08_OServiceRateParamsSubAdd)) *mockOServiceService_RateParamsSubAdd_Call {
	_c.Call.Return(run)
	return _c
}

// SetUserInfoFields provides a mock function with given fields: ctx, sess, frame, bodyIn
func (_m *mockOServiceService) SetUserInfoFields(ctx context.Context, sess *state.Session, frame wire.SNACFrame, bodyIn wire.SNAC_0x01_0x1E_OServiceSetUserInfoFields) (wire.SNACMessage, error) {
	ret := _m.Called(ctx, sess, frame, bodyIn)

	if len(ret) == 0 {
		panic("no return value specified for SetUserInfoFields")
	}

	var r0 wire.SNACMessage
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *state.Session, wire.SNACFrame, wire.SNAC_0x01_0x1E_OServiceSetUserInfoFields) (wire.SNACMessage, error)); ok {
		return rf(ctx, sess, frame, bodyIn)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *state.Session, wire.SNACFrame, wire.SNAC_0x01_0x1E_OServiceSetUserInfoFields) wire.SNACMessage); ok {
		r0 = rf(ctx, sess, frame, bodyIn)
	} else {
		r0 = ret.Get(0).(wire.SNACMessage)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *state.Session, wire.SNACFrame, wire.SNAC_0x01_0x1E_OServiceSetUserInfoFields) error); ok {
		r1 = rf(ctx, sess, frame, bodyIn)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mockOServiceService_SetUserInfoFields_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetUserInfoFields'
type mockOServiceService_SetUserInfoFields_Call struct {
	*mock.Call
}

// SetUserInfoFields is a helper method to define mock.On call
//   - ctx context.Context
//   - sess *state.Session
//   - frame wire.SNACFrame
//   - bodyIn wire.SNAC_0x01_0x1E_OServiceSetUserInfoFields
func (_e *mockOServiceService_Expecter) SetUserInfoFields(ctx interface{}, sess interface{}, frame interface{}, bodyIn interface{}) *mockOServiceService_SetUserInfoFields_Call {
	return &mockOServiceService_SetUserInfoFields_Call{Call: _e.mock.On("SetUserInfoFields", ctx, sess, frame, bodyIn)}
}

func (_c *mockOServiceService_SetUserInfoFields_Call) Run(run func(ctx context.Context, sess *state.Session, frame wire.SNACFrame, bodyIn wire.SNAC_0x01_0x1E_OServiceSetUserInfoFields)) *mockOServiceService_SetUserInfoFields_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*state.Session), args[2].(wire.SNACFrame), args[3].(wire.SNAC_0x01_0x1E_OServiceSetUserInfoFields))
	})
	return _c
}

func (_c *mockOServiceService_SetUserInfoFields_Call) Return(_a0 wire.SNACMessage, _a1 error) *mockOServiceService_SetUserInfoFields_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *mockOServiceService_SetUserInfoFields_Call) RunAndReturn(run func(context.Context, *state.Session, wire.SNACFrame, wire.SNAC_0x01_0x1E_OServiceSetUserInfoFields) (wire.SNACMessage, error)) *mockOServiceService_SetUserInfoFields_Call {
	_c.Call.Return(run)
	return _c
}

// UserInfoQuery provides a mock function with given fields: ctx, sess, frame
func (_m *mockOServiceService) UserInfoQuery(ctx context.Context, sess *state.Session, frame wire.SNACFrame) wire.SNACMessage {
	ret := _m.Called(ctx, sess, frame)

	if len(ret) == 0 {
		panic("no return value specified for UserInfoQuery")
	}

	var r0 wire.SNACMessage
	if rf, ok := ret.Get(0).(func(context.Context, *state.Session, wire.SNACFrame) wire.SNACMessage); ok {
		r0 = rf(ctx, sess, frame)
	} else {
		r0 = ret.Get(0).(wire.SNACMessage)
	}

	return r0
}

// mockOServiceService_UserInfoQuery_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UserInfoQuery'
type mockOServiceService_UserInfoQuery_Call struct {
	*mock.Call
}

// UserInfoQuery is a helper method to define mock.On call
//   - ctx context.Context
//   - sess *state.Session
//   - frame wire.SNACFrame
func (_e *mockOServiceService_Expecter) UserInfoQuery(ctx interface{}, sess interface{}, frame interface{}) *mockOServiceService_UserInfoQuery_Call {
	return &mockOServiceService_UserInfoQuery_Call{Call: _e.mock.On("UserInfoQuery", ctx, sess, frame)}
}

func (_c *mockOServiceService_UserInfoQuery_Call) Run(run func(ctx context.Context, sess *state.Session, frame wire.SNACFrame)) *mockOServiceService_UserInfoQuery_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*state.Session), args[2].(wire.SNACFrame))
	})
	return _c
}

func (_c *mockOServiceService_UserInfoQuery_Call) Return(_a0 wire.SNACMessage) *mockOServiceService_UserInfoQuery_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mockOServiceService_UserInfoQuery_Call) RunAndReturn(run func(context.Context, *state.Session, wire.SNACFrame) wire.SNACMessage) *mockOServiceService_UserInfoQuery_Call {
	_c.Call.Return(run)
	return _c
}

// newMockOServiceService creates a new instance of mockOServiceService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockOServiceService(t interface {
	mock.TestingT
	Cleanup(func())
}) *mockOServiceService {
	mock := &mockOServiceService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
