package service_test

import (
	"errors"
	"sql-playground/mocks"
	"sql-playground/repository"
	"sql-playground/service"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_ExampleService_Do_Success(t *testing.T) {
	repoManagerMock := &mocks.RepositoryManagerMock{}
	defer repoManagerMock.AssertExpectations(t)
	atomicRepoManagerMock := &mocks.RepositoryManagerMock{}
	defer atomicRepoManagerMock.AssertExpectations(t)
	exampleRepoMock := &mocks.ExampleRepoMock{}
	defer exampleRepoMock.AssertExpectations(t)

	repoManagerMock.
		On("RunAtomic", mock.Anything).
		Run(func(args mock.Arguments) {
			fn := args.Get(0).(func(atomicRM repository.RepositoryManager) error)

			atomicRepoManagerMock.On("ExampleRepo").Return(exampleRepoMock).Once()
			exampleRepoMock.On("Do").Return(nil).Once()

			err := fn(atomicRepoManagerMock)

			assert.Nil(t, err)
		}).
		Return(nil).
		Once()

	service := service.NewExampleService(repoManagerMock)

	err := service.Do()

	assert.Nil(t, err)
}

func Test_ExampleService_Do_Failure(t *testing.T) {
	repoManagerMock := &mocks.RepositoryManagerMock{}
	defer repoManagerMock.AssertExpectations(t)
	atomicRepoManagerMock := &mocks.RepositoryManagerMock{}
	defer atomicRepoManagerMock.AssertExpectations(t)
	exampleRepoMock := &mocks.ExampleRepoMock{}
	defer exampleRepoMock.AssertExpectations(t)

	expectedErr := errors.New("repo Do error")

	repoManagerMock.
		On("RunAtomic", mock.Anything).
		Run(func(args mock.Arguments) {
			fn := args.Get(0).(func(atomicRM repository.RepositoryManager) error)

			atomicRepoManagerMock.On("ExampleRepo").Return(exampleRepoMock).Once()
			exampleRepoMock.On("Do").Return(expectedErr).Once()

			err := fn(atomicRepoManagerMock)

			assert.Equal(t, expectedErr, err)
		}).
		Return(expectedErr).
		Once()

	service := service.NewExampleService(repoManagerMock)

	err := service.Do()

	assert.Equal(t, expectedErr, err)
}
