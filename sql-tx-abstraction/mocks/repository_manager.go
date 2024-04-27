package mocks

import (
	"sql-playground/repository"

	"github.com/stretchr/testify/mock"
)

type RepositoryManagerMock struct {
	mock.Mock
}

func (m *RepositoryManagerMock) RunAtomic(f func(repository.RepositoryManager) error) error {
	ret := m.Called(f)

	return ret.Error(0)
}

func (m *RepositoryManagerMock) ExampleRepo() repository.ExampleRepo {
	ret := m.Called()

	return ret.Get(0).(repository.ExampleRepo)
}
