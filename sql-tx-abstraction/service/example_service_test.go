package service_test

import (
	"errors"
	"sql-tx-abstraction/repository"
	"sql-tx-abstraction/service"
	"sql-tx-abstraction/service/testdata"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_ExampleService_Do_Success(t *testing.T) {
	mockExampleRepo := testdata.NewMockExampleRepo(t)
	mockTxProvider := testdata.NewMockTransactionProvider(t)

	atomicRepos := repository.Repositories{
		ExampleRepo: mockExampleRepo,
	}

	mockTxProvider.
		EXPECT().
		RunInTx(mock.Anything).
		RunAndReturn(func(fn repository.TxFn) error {
			mockExampleRepo.EXPECT().Do().Return(nil)
			return fn(atomicRepos)
		})

	service := service.NewExampleService(mockExampleRepo, mockTxProvider)

	err := service.Do()

	assert.Nil(t, err)
}

func Test_ExampleService_Do_Failure(t *testing.T) {
	mockExampleRepo := testdata.NewMockExampleRepo(t)
	mockTxProvider := testdata.NewMockTransactionProvider(t)

	atomicRepos := repository.Repositories{
		ExampleRepo: mockExampleRepo,
	}

	expectedErr := errors.New("repo Do error")

	mockTxProvider.
		EXPECT().
		RunInTx(mock.Anything).
		RunAndReturn(func(fn repository.TxFn) error {
			mockExampleRepo.EXPECT().Do().Return(expectedErr)
			return fn(atomicRepos)
		})

	service := service.NewExampleService(mockExampleRepo, mockTxProvider)

	err := service.Do()

	assert.Equal(t, expectedErr, err)
}
