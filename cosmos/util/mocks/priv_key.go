// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	types "github.com/cosmos/cosmos-sdk/crypto/types"
	mock "github.com/stretchr/testify/mock"
)

// MockPrivKey is an autogenerated mock type for the PrivKey type
type MockPrivKey struct {
	mock.Mock
}

type MockPrivKey_Expecter struct {
	mock *mock.Mock
}

func (_m *MockPrivKey) EXPECT() *MockPrivKey_Expecter {
	return &MockPrivKey_Expecter{mock: &_m.Mock}
}

// Bytes provides a mock function with given fields:
func (_m *MockPrivKey) Bytes() []byte {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Bytes")
	}

	var r0 []byte
	if rf, ok := ret.Get(0).(func() []byte); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	return r0
}

// MockPrivKey_Bytes_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Bytes'
type MockPrivKey_Bytes_Call struct {
	*mock.Call
}

// Bytes is a helper method to define mock.On call
func (_e *MockPrivKey_Expecter) Bytes() *MockPrivKey_Bytes_Call {
	return &MockPrivKey_Bytes_Call{Call: _e.mock.On("Bytes")}
}

func (_c *MockPrivKey_Bytes_Call) Run(run func()) *MockPrivKey_Bytes_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPrivKey_Bytes_Call) Return(_a0 []byte) *MockPrivKey_Bytes_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPrivKey_Bytes_Call) RunAndReturn(run func() []byte) *MockPrivKey_Bytes_Call {
	_c.Call.Return(run)
	return _c
}

// Equals provides a mock function with given fields: _a0
func (_m *MockPrivKey) Equals(_a0 types.LedgerPrivKey) bool {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Equals")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(types.LedgerPrivKey) bool); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockPrivKey_Equals_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Equals'
type MockPrivKey_Equals_Call struct {
	*mock.Call
}

// Equals is a helper method to define mock.On call
//   - _a0 types.LedgerPrivKey
func (_e *MockPrivKey_Expecter) Equals(_a0 interface{}) *MockPrivKey_Equals_Call {
	return &MockPrivKey_Equals_Call{Call: _e.mock.On("Equals", _a0)}
}

func (_c *MockPrivKey_Equals_Call) Run(run func(_a0 types.LedgerPrivKey)) *MockPrivKey_Equals_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(types.LedgerPrivKey))
	})
	return _c
}

func (_c *MockPrivKey_Equals_Call) Return(_a0 bool) *MockPrivKey_Equals_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPrivKey_Equals_Call) RunAndReturn(run func(types.LedgerPrivKey) bool) *MockPrivKey_Equals_Call {
	_c.Call.Return(run)
	return _c
}

// ProtoMessage provides a mock function with given fields:
func (_m *MockPrivKey) ProtoMessage() {
	_m.Called()
}

// MockPrivKey_ProtoMessage_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ProtoMessage'
type MockPrivKey_ProtoMessage_Call struct {
	*mock.Call
}

// ProtoMessage is a helper method to define mock.On call
func (_e *MockPrivKey_Expecter) ProtoMessage() *MockPrivKey_ProtoMessage_Call {
	return &MockPrivKey_ProtoMessage_Call{Call: _e.mock.On("ProtoMessage")}
}

func (_c *MockPrivKey_ProtoMessage_Call) Run(run func()) *MockPrivKey_ProtoMessage_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPrivKey_ProtoMessage_Call) Return() *MockPrivKey_ProtoMessage_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockPrivKey_ProtoMessage_Call) RunAndReturn(run func()) *MockPrivKey_ProtoMessage_Call {
	_c.Call.Return(run)
	return _c
}

// PubKey provides a mock function with given fields:
func (_m *MockPrivKey) PubKey() types.PubKey {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for PubKey")
	}

	var r0 types.PubKey
	if rf, ok := ret.Get(0).(func() types.PubKey); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.PubKey)
		}
	}

	return r0
}

// MockPrivKey_PubKey_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PubKey'
type MockPrivKey_PubKey_Call struct {
	*mock.Call
}

// PubKey is a helper method to define mock.On call
func (_e *MockPrivKey_Expecter) PubKey() *MockPrivKey_PubKey_Call {
	return &MockPrivKey_PubKey_Call{Call: _e.mock.On("PubKey")}
}

func (_c *MockPrivKey_PubKey_Call) Run(run func()) *MockPrivKey_PubKey_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPrivKey_PubKey_Call) Return(_a0 types.PubKey) *MockPrivKey_PubKey_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPrivKey_PubKey_Call) RunAndReturn(run func() types.PubKey) *MockPrivKey_PubKey_Call {
	_c.Call.Return(run)
	return _c
}

// Reset provides a mock function with given fields:
func (_m *MockPrivKey) Reset() {
	_m.Called()
}

// MockPrivKey_Reset_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Reset'
type MockPrivKey_Reset_Call struct {
	*mock.Call
}

// Reset is a helper method to define mock.On call
func (_e *MockPrivKey_Expecter) Reset() *MockPrivKey_Reset_Call {
	return &MockPrivKey_Reset_Call{Call: _e.mock.On("Reset")}
}

func (_c *MockPrivKey_Reset_Call) Run(run func()) *MockPrivKey_Reset_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPrivKey_Reset_Call) Return() *MockPrivKey_Reset_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockPrivKey_Reset_Call) RunAndReturn(run func()) *MockPrivKey_Reset_Call {
	_c.Call.Return(run)
	return _c
}

// Sign provides a mock function with given fields: msg
func (_m *MockPrivKey) Sign(msg []byte) ([]byte, error) {
	ret := _m.Called(msg)

	if len(ret) == 0 {
		panic("no return value specified for Sign")
	}

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func([]byte) ([]byte, error)); ok {
		return rf(msg)
	}
	if rf, ok := ret.Get(0).(func([]byte) []byte); ok {
		r0 = rf(msg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(msg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockPrivKey_Sign_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Sign'
type MockPrivKey_Sign_Call struct {
	*mock.Call
}

// Sign is a helper method to define mock.On call
//   - msg []byte
func (_e *MockPrivKey_Expecter) Sign(msg interface{}) *MockPrivKey_Sign_Call {
	return &MockPrivKey_Sign_Call{Call: _e.mock.On("Sign", msg)}
}

func (_c *MockPrivKey_Sign_Call) Run(run func(msg []byte)) *MockPrivKey_Sign_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]byte))
	})
	return _c
}

func (_c *MockPrivKey_Sign_Call) Return(_a0 []byte, _a1 error) *MockPrivKey_Sign_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockPrivKey_Sign_Call) RunAndReturn(run func([]byte) ([]byte, error)) *MockPrivKey_Sign_Call {
	_c.Call.Return(run)
	return _c
}

// String provides a mock function with given fields:
func (_m *MockPrivKey) String() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for String")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockPrivKey_String_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'String'
type MockPrivKey_String_Call struct {
	*mock.Call
}

// String is a helper method to define mock.On call
func (_e *MockPrivKey_Expecter) String() *MockPrivKey_String_Call {
	return &MockPrivKey_String_Call{Call: _e.mock.On("String")}
}

func (_c *MockPrivKey_String_Call) Run(run func()) *MockPrivKey_String_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPrivKey_String_Call) Return(_a0 string) *MockPrivKey_String_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPrivKey_String_Call) RunAndReturn(run func() string) *MockPrivKey_String_Call {
	_c.Call.Return(run)
	return _c
}

// Type provides a mock function with given fields:
func (_m *MockPrivKey) Type() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Type")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockPrivKey_Type_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Type'
type MockPrivKey_Type_Call struct {
	*mock.Call
}

// Type is a helper method to define mock.On call
func (_e *MockPrivKey_Expecter) Type() *MockPrivKey_Type_Call {
	return &MockPrivKey_Type_Call{Call: _e.mock.On("Type")}
}

func (_c *MockPrivKey_Type_Call) Run(run func()) *MockPrivKey_Type_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPrivKey_Type_Call) Return(_a0 string) *MockPrivKey_Type_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPrivKey_Type_Call) RunAndReturn(run func() string) *MockPrivKey_Type_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockPrivKey creates a new instance of MockPrivKey. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockPrivKey(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockPrivKey {
	mock := &MockPrivKey{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}