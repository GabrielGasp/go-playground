package mocks

import "github.com/stretchr/testify/mock"

type ExampleServiceMock struct {
	mock.Mock
}

func (m *ExampleServiceMock) ExecExample() error {
	ret := m.Called()

	return ret.Error(0)
}
