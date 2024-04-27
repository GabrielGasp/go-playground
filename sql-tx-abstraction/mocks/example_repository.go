package mocks

import "github.com/stretchr/testify/mock"

type ExampleRepoMock struct {
	mock.Mock
}

func (m *ExampleRepoMock) Do() error {
	ret := m.Called()

	return ret.Error(0)
}
