// Code generated by mockery v2.43.2. DO NOT EDIT.

package io

import mock "github.com/stretchr/testify/mock"

// MockSeeker is an autogenerated mock type for the Seeker type
type MockSeeker struct {
	mock.Mock
}

type MockSeeker_Expecter struct {
	mock *mock.Mock
}

func (_m *MockSeeker) EXPECT() *MockSeeker_Expecter {
	return &MockSeeker_Expecter{mock: &_m.Mock}
}

// Seek provides a mock function with given fields: offset, whence
func (_m *MockSeeker) Seek(offset int64, whence int) (int64, error) {
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

// MockSeeker_Seek_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Seek'
type MockSeeker_Seek_Call struct {
	*mock.Call
}

// Seek is a helper method to define mock.On call
//   - offset int64
//   - whence int
func (_e *MockSeeker_Expecter) Seek(offset interface{}, whence interface{}) *MockSeeker_Seek_Call {
	return &MockSeeker_Seek_Call{Call: _e.mock.On("Seek", offset, whence)}
}

func (_c *MockSeeker_Seek_Call) Run(run func(offset int64, whence int)) *MockSeeker_Seek_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int64), args[1].(int))
	})
	return _c
}

func (_c *MockSeeker_Seek_Call) Return(_a0 int64, _a1 error) *MockSeeker_Seek_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSeeker_Seek_Call) RunAndReturn(run func(int64, int) (int64, error)) *MockSeeker_Seek_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockSeeker creates a new instance of MockSeeker. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockSeeker(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockSeeker {
	mock := &MockSeeker{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
