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

	exampleRepoMock := mocks.NewMockExampleRepo(mockCtrl)
	txProviderMock := mocks.NewMockTransactionProvider(mockCtrl)

	atomicRepos := repository.Repositories{
		ExampleRepo: exampleRepoMock,
	}

	txProviderMock.
		EXPECT().
		RunInTx(gomock.Any()).
		DoAndReturn(func(fn repository.TxFn) error {
			exampleRepoMock.EXPECT().Do().Return(nil)
			return fn(atomicRepos)
		})

	service := service.NewExampleService(exampleRepoMock, txProviderMock)

	err := service.Do()

	assert.Nil(t, err)
}

func Test_ExampleService_Do_Failure(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	exampleRepoMock := mocks.NewMockExampleRepo(mockCtrl)
	txProviderMock := mocks.NewMockTransactionProvider(mockCtrl)

	atomicRepos := repository.Repositories{
		ExampleRepo: exampleRepoMock,
	}

	expectedErr := errors.New("repo Do error")

	txProviderMock.
		EXPECT().
		RunInTx(gomock.Any()).
		DoAndReturn(func(fn repository.TxFn) error {
			exampleRepoMock.EXPECT().Do().Return(expectedErr)

			return fn(atomicRepos)
		})

	service := service.NewExampleService(exampleRepoMock, txProviderMock)

	err := service.Do()

	assert.Equal(t, expectedErr, err)
}
