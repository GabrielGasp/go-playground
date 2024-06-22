// Code generated by mockery v2.43.2. DO NOT EDIT.

package io

import mock "github.com/stretchr/testify/mock"

// MockWriteSeeker is an autogenerated mock type for the WriteSeeker type
type MockWriteSeeker struct {
	mock.Mock
}

type MockWriteSeeker_Expecter struct {
	mock *mock.Mock
}

func (_m *MockWriteSeeker) EXPECT() *MockWriteSeeker_Expecter {
	return &MockWriteSeeker_Expecter{mock: &_m.Mock}
}

// Seek provides a mock function with given fields: offset, whence
func (_m *MockWriteSeeker) Seek(offset int64, whence int) (int64, error) {
	ret := _m.Called(offset, whence)

	if len(ret) == 0 {
		panic("no return value specified for Seek")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(int64, int) (int64, error)); ok {
		return rf(offset, whence)
	}
	if rf, ok := ret.Get(0).(func(int64, int) int64); ok {
		r0 = rf(offset, whence)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(int64, int) error); ok {
		r1 = rf(offset, whence)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockWriteSeeker_Seek_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Seek'
type MockWriteSeeker_Seek_Call struct {
	*mock.Call
}

// Seek is a helper method to define mock.On call
//   - offset int64
//   - whence int
func (_e *MockWriteSeeker_Expecter) Seek(offset interface{}, whence interface{}) *MockWriteSeeker_Seek_Call {
	return &MockWriteSeeker_Seek_Call{Call: _e.mock.On("Seek", offset, whence)}
}

func (_c *MockWriteSeeker_Seek_Call) Run(run func(offset int64, whence int)) *MockWriteSeeker_Seek_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int64), args[1].(int))
	})
	return _c
}

func (_c *MockWriteSeeker_Seek_Call) Return(_a0 int64, _a1 error) *MockWriteSeeker_Seek_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockWriteSeeker_Seek_Call) RunAndReturn(run func(int64, int) (int64, error)) *MockWriteSeeker_Seek_Call {
	_c.Call.Return(run)
	return _c
}

// Write provides a mock function with given fields: p
func (_m *MockWriteSeeker) Write(p []byte) (int, error) {
	ret := _m.Called(p)

	if len(ret) == 0 {
		panic("no return value specified for Write")
	}

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func([]byte) (int, error)); ok {
		return rf(p)
	}
	if rf, ok := ret.Get(0).(func([]byte) int); ok {
		r0 = rf(p)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(p)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockWriteSeeker_Write_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Write'
type MockWriteSeeker_Write_Call struct {
	*mock.Call
}

// Write is a helper method to define mock.On call
//   - p []byte
func (_e *MockWriteSeeker_Expecter) Write(p interface{}) *MockWriteSeeker_Write_Call {
	return &MockWriteSeeker_Write_Call{Call: _e.mock.On("Write", p)}
}

func (_c *MockWriteSeeker_Write_Call) Run(run func(p []byte)) *MockWriteSeeker_Write_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]byte))
	})
	return _c
}

func (_c *MockWriteSeeker_Write_Call) Return(n int, err error) *MockWriteSeeker_Write_Call {
	_c.Call.Return(n, err)
	return _c
}

func (_c *MockWriteSeeker_Write_Call) RunAndReturn(run func([]byte) (int, error)) *MockWriteSeeker_Write_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockWriteSeeker creates a new instance of MockWriteSeeker. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockWriteSeeker(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockWriteSeeker {
	mock := &MockWriteSeeker{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
