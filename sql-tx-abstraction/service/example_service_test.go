package service_test

import (
	"errors"
	"sql-tx-abstraction/repository"
	"sql-tx-abstraction/repository/mocks"
	"sql-tx-abstraction/service"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_ExampleService_Do_Success(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	repoManagerMock := mocks.NewMockRepositoryManager(mockCtrl)
	atomicRepoManagerMock := mocks.NewMockRepositoryManager(mockCtrl)
	exampleRepoMock := mocks.NewMockExampleRepo(mockCtrl)

	repoManagerMock.
		EXPECT().
		RunAtomic(gomock.Any()).
		DoAndReturn(func(fn repository.AtomicFn) error {
			atomicRepoManagerMock.EXPECT().ExampleRepo().Return(exampleRepoMock)
			exampleRepoMock.EXPECT().Do().Return(nil)

			return fn(atomicRepoManagerMock)
		})

	service := service.NewExampleService(repoManagerMock)

	err := service.Do()

	assert.Nil(t, err)
}

func Test_ExampleService_Do_Failure(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	repoManagerMock := mocks.NewMockRepositoryManager(mockCtrl)
	atomicRepoManagerMock := mocks.NewMockRepositoryManager(mockCtrl)
	exampleRepoMock := mocks.NewMockExampleRepo(mockCtrl)

	expectedErr := errors.New("repo Do error")

	repoManagerMock.
		EXPECT().
		RunAtomic(gomock.Any()).
		DoAndReturn(func(fn repository.AtomicFn) error {
			atomicRepoManagerMock.EXPECT().ExampleRepo().Return(exampleRepoMock)
			exampleRepoMock.EXPECT().Do().Return(expectedErr)

			return fn(atomicRepoManagerMock)
		})

	service := service.NewExampleService(repoManagerMock)

	err := service.Do()

	assert.Equal(t, expectedErr, err)
}
