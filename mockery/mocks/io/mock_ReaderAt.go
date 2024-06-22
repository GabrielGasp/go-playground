// Code generated by mockery v2.43.2. DO NOT EDIT.

package io

import mock "github.com/stretchr/testify/mock"

// MockReaderAt is an autogenerated mock type for the ReaderAt type
type MockReaderAt struct {
	mock.Mock
}

type MockReaderAt_Expecter struct {
	mock *mock.Mock
}

func (_m *MockReaderAt) EXPECT() *MockReaderAt_Expecter {
	return &MockReaderAt_Expecter{mock: &_m.Mock}
}

// ReadAt provides a mock function with given fields: p, off
func (_m *MockReaderAt) ReadAt(p []byte, off int64) (int, error) {
	ret := _m.Called(p, off)

	if len(ret) == 0 {
		panic("no return value specified for ReadAt")
	}

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func([]byte, int64) (int, error)); ok {
		return rf(p, off)
	}
	if rf, ok := ret.Get(0).(func([]byte, int64) int); ok {
		r0 = rf(p, off)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func([]byte, int64) error); ok {
		r1 = rf(p, off)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockReaderAt_ReadAt_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ReadAt'
type MockReaderAt_ReadAt_Call struct {
	*mock.Call
}

// ReadAt is a helper method to define mock.On call
//   - p []byte
//   - off int64
func (_e *MockReaderAt_Expecter) ReadAt(p interface{}, off interface{}) *MockReaderAt_ReadAt_Call {
	return &MockReaderAt_ReadAt_Call{Call: _e.mock.On("ReadAt", p, off)}
}

func (_c *MockReaderAt_ReadAt_Call) Run(run func(p []byte, off int64)) *MockReaderAt_ReadAt_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]byte), args[1].(int64))
	})
	return _c
}

func (_c *MockReaderAt_ReadAt_Call) Return(n int, err error) *MockReaderAt_ReadAt_Call {
	_c.Call.Return(n, err)
	return _c
}

func (_c *MockReaderAt_ReadAt_Call) RunAndReturn(run func([]byte, int64) (int, error)) *MockReaderAt_ReadAt_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockReaderAt creates a new instance of MockReaderAt. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockReaderAt(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockReaderAt {
	mock := &MockReaderAt{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
