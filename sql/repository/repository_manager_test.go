package repository_test

import (
	"errors"
	"sql-playground/repository"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func Test_RunAtomic_CommitSuccessfully(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer mockDB.Close()
	defer mock.ExpectationsWereMet()

	mock.ExpectBegin()
	mock.ExpectCommit()

	rm := repository.NewRepositoryManager(mockDB)

	err = rm.RunAtomic(func(atomicRM repository.RepositoryManager) error {
		return nil
	})

	assert.NoError(t, err)
}

func Test_RunAtomic_RollbackSuccessfully(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer mockDB.Close()
	defer mock.ExpectationsWereMet()

	mock.ExpectBegin()
	mock.ExpectRollback()

	rm := repository.NewRepositoryManager(mockDB)

	expectedErr := errors.New("fn error")

	err = rm.RunAtomic(func(atomicRM repository.RepositoryManager) error {
		return expectedErr
	})

	assert.Equal(t, expectedErr, err)
}

func Test_RunAtomic_NestedTransaction(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer mockDB.Close()
	defer mock.ExpectationsWereMet()

	mock.ExpectBegin()
	mock.ExpectCommit()

	rm := repository.NewRepositoryManager(mockDB)

	err = rm.RunAtomic(func(atomicRM repository.RepositoryManager) error {
		return atomicRM.RunAtomic(func(nestedAtomicRM repository.RepositoryManager) error {
			assert.Equal(t, atomicRM, nestedAtomicRM)
			return nil
		})
	})

	assert.NoError(t, err)
}

func Test_RunAtomic_BeginError(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer mockDB.Close()
	defer mock.ExpectationsWereMet()

	expectedErr := errors.New("begin error")

	mock.ExpectBegin().WillReturnError(expectedErr)

	rm := repository.NewRepositoryManager(mockDB)

	err = rm.RunAtomic(func(atomicRM repository.RepositoryManager) error {
		return nil
	})

	assert.Equal(t, expectedErr, err)
}

func Test_ExampleRepo(t *testing.T) {
	mockDB, _, err := sqlmock.New()
	assert.NoError(t, err)
	defer mockDB.Close()

	rm := repository.NewRepositoryManager(mockDB)

	exampleRepo := rm.ExampleRepo()

	assert.NotNil(t, exampleRepo)
}
