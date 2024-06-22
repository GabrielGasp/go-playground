// Code generated by mockery v2.43.2. DO NOT EDIT.

package io

import mock "github.com/stretchr/testify/mock"

// MockReadWriter is an autogenerated mock type for the ReadWriter type
type MockReadWriter struct {
	mock.Mock
}

type MockReadWriter_Expecter struct {
	mock *mock.Mock
}

func (_m *MockReadWriter) EXPECT() *MockReadWriter_Expecter {
	return &MockReadWriter_Expecter{mock: &_m.Mock}
}

// Read provides a mock function with given fields: p
func (_m *MockReadWriter) Read(p []byte) (int, error) {
	ret := _m.Called(p)

	if len(ret) == 0 {
		panic("no return value specified for Read")
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

// MockReadWriter_Read_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Read'
type MockReadWriter_Read_Call struct {
	*mock.Call
}

// Read is a helper method to define mock.On call
//   - p []byte
func (_e *MockReadWriter_Expecter) Read(p interface{}) *MockReadWriter_Read_Call {
	return &MockReadWriter_Read_Call{Call: _e.mock.On("Read", p)}
}

func (_c *MockReadWriter_Read_Call) Run(run func(p []byte)) *MockReadWriter_Read_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]byte))
	})
	return _c
}

func (_c *MockReadWriter_Read_Call) Return(n int, err error) *MockReadWriter_Read_Call {
	_c.Call.Return(n, err)
	return _c
}

func (_c *MockReadWriter_Read_Call) RunAndReturn(run func([]byte) (int, error)) *MockReadWriter_Read_Call {
	_c.Call.Return(run)
	return _c
}

// Write provides a mock function with given fields: p
func (_m *MockReadWriter) Write(p []byte) (int, error) {
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

// MockReadWriter_Write_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Write'
type MockReadWriter_Write_Call struct {
	*mock.Call
}

// Write is a helper method to define mock.On call
//   - p []byte
func (_e *MockReadWriter_Expecter) Write(p interface{}) *MockReadWriter_Write_Call {
	return &MockReadWriter_Write_Call{Call: _e.mock.On("Write", p)}
}

func (_c *MockReadWriter_Write_Call) Run(run func(p []byte)) *MockReadWriter_Write_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]byte))
	})
	return _c
}

func (_c *MockReadWriter_Write_Call) Return(n int, err error) *MockReadWriter_Write_Call {
	_c.Call.Return(n, err)
	return _c
}

func (_c *MockReadWriter_Write_Call) RunAndReturn(run func([]byte) (int, error)) *MockReadWriter_Write_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockReadWriter creates a new instance of MockReadWriter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockReadWriter(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockReadWriter {
	mock := &MockReadWriter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
