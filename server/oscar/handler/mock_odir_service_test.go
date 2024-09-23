// Code generated by mockery v2.45.0. DO NOT EDIT.

package handler

import (
	context "context"

	wire "github.com/mk6i/retro-aim-server/wire"
	mock "github.com/stretchr/testify/mock"
)

// mockODirService is an autogenerated mock type for the ODirService type
type mockODirService struct {
	mock.Mock
}

type mockODirService_Expecter struct {
	mock *mock.Mock
}

func (_m *mockODirService) EXPECT() *mockODirService_Expecter {
	return &mockODirService_Expecter{mock: &_m.Mock}
}

// InfoQuery provides a mock function with given fields: ctx, inFrame, inBody
func (_m *mockODirService) InfoQuery(ctx context.Context, inFrame wire.SNACFrame, inBody wire.SNAC_0x0F_0x02_InfoQuery) (wire.SNACMessage, error) {
	ret := _m.Called(ctx, inFrame, inBody)

	if len(ret) == 0 {
		panic("no return value specified for InfoQuery")
	}

	var r0 wire.SNACMessage
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, wire.SNACFrame, wire.SNAC_0x0F_0x02_InfoQuery) (wire.SNACMessage, error)); ok {
		return rf(ctx, inFrame, inBody)
	}
	if rf, ok := ret.Get(0).(func(context.Context, wire.SNACFrame, wire.SNAC_0x0F_0x02_InfoQuery) wire.SNACMessage); ok {
		r0 = rf(ctx, inFrame, inBody)
	} else {
		r0 = ret.Get(0).(wire.SNACMessage)
	}

	if rf, ok := ret.Get(1).(func(context.Context, wire.SNACFrame, wire.SNAC_0x0F_0x02_InfoQuery) error); ok {
		r1 = rf(ctx, inFrame, inBody)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mockODirService_InfoQuery_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'InfoQuery'
type mockODirService_InfoQuery_Call struct {
	*mock.Call
}

// InfoQuery is a helper method to define mock.On call
//   - ctx context.Context
//   - inFrame wire.SNACFrame
//   - inBody wire.SNAC_0x0F_0x02_InfoQuery
func (_e *mockODirService_Expecter) InfoQuery(ctx interface{}, inFrame interface{}, inBody interface{}) *mockODirService_InfoQuery_Call {
	return &mockODirService_InfoQuery_Call{Call: _e.mock.On("InfoQuery", ctx, inFrame, inBody)}
}

func (_c *mockODirService_InfoQuery_Call) Run(run func(ctx context.Context, inFrame wire.SNACFrame, inBody wire.SNAC_0x0F_0x02_InfoQuery)) *mockODirService_InfoQuery_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(wire.SNACFrame), args[2].(wire.SNAC_0x0F_0x02_InfoQuery))
	})
	return _c
}

func (_c *mockODirService_InfoQuery_Call) Return(_a0 wire.SNACMessage, _a1 error) *mockODirService_InfoQuery_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *mockODirService_InfoQuery_Call) RunAndReturn(run func(context.Context, wire.SNACFrame, wire.SNAC_0x0F_0x02_InfoQuery) (wire.SNACMessage, error)) *mockODirService_InfoQuery_Call {
	_c.Call.Return(run)
	return _c
}

// KeywordListQuery provides a mock function with given fields: _a0, _a1
func (_m *mockODirService) KeywordListQuery(_a0 context.Context, _a1 wire.SNACFrame) (wire.SNACMessage, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for KeywordListQuery")
	}

	var r0 wire.SNACMessage
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, wire.SNACFrame) (wire.SNACMessage, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, wire.SNACFrame) wire.SNACMessage); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(wire.SNACMessage)
	}

	if rf, ok := ret.Get(1).(func(context.Context, wire.SNACFrame) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mockODirService_KeywordListQuery_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'KeywordListQuery'
type mockODirService_KeywordListQuery_Call struct {
	*mock.Call
}

// KeywordListQuery is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 wire.SNACFrame
func (_e *mockODirService_Expecter) KeywordListQuery(_a0 interface{}, _a1 interface{}) *mockODirService_KeywordListQuery_Call {
	return &mockODirService_KeywordListQuery_Call{Call: _e.mock.On("KeywordListQuery", _a0, _a1)}
}

func (_c *mockODirService_KeywordListQuery_Call) Run(run func(_a0 context.Context, _a1 wire.SNACFrame)) *mockODirService_KeywordListQuery_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(wire.SNACFrame))
	})
	return _c
}

func (_c *mockODirService_KeywordListQuery_Call) Return(_a0 wire.SNACMessage, _a1 error) *mockODirService_KeywordListQuery_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *mockODirService_KeywordListQuery_Call) RunAndReturn(run func(context.Context, wire.SNACFrame) (wire.SNACMessage, error)) *mockODirService_KeywordListQuery_Call {
	_c.Call.Return(run)
	return _c
}

// newMockODirService creates a new instance of mockODirService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockODirService(t interface {
	mock.TestingT
	Cleanup(func())
}) *mockODirService {
	mock := &mockODirService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
