package repository_test

import (
	"errors"
	"sql-tx-abstraction/repository"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func Test_TransactionProvider_RunInTx_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer func() {
		assert.NoError(t, mock.ExpectationsWereMet())
		db.Close()
	}()

	mock.ExpectBegin()
	mock.ExpectCommit()

	txProvider := repository.NewTransactionProvider(db)

	err = txProvider.RunInTx(func(atomicRepos repository.Repositories) error {
		return nil
	})

	assert.NoError(t, err)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func Test_TransactionProvider_RunInTx_BeginError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer func() {
		assert.NoError(t, mock.ExpectationsWereMet())
		db.Close()
	}()

	expectedErr := errors.New("begin error")

	mock.ExpectBegin().WillReturnError(expectedErr)

	txProvider := repository.NewTransactionProvider(db)

	err = txProvider.RunInTx(func(atomicRepos repository.Repositories) error {
		return nil
	})

	assert.Equal(t, expectedErr, err)
}

func Test_TransactionProvider_RunInTx_FnError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer func() {
		assert.NoError(t, mock.ExpectationsWereMet())
		db.Close()
	}()

	mock.ExpectBegin()
	mock.ExpectRollback()

	txProvider := repository.NewTransactionProvider(db)

	expectedErr := errors.New("fn error")

	err = txProvider.RunInTx(func(atomicRepos repository.Repositories) error {
		return expectedErr
	})

	assert.Equal(t, expectedErr, err)
}

func Test_TransactionProvider_RunInTx_RollbackError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer func() {
		assert.NoError(t, mock.ExpectationsWereMet())
		db.Close()
	}()

	expectedFnErr := errors.New("fn error")
	expectedRbErr := errors.New("rollback error")
	expectedJoinedErr := errors.Join(expectedFnErr, expectedRbErr)

	mock.ExpectBegin()
	mock.ExpectRollback().WillReturnError(expectedRbErr)

	txProvider := repository.NewTransactionProvider(db)

	err = txProvider.RunInTx(func(atomicRepos repository.Repositories) error {
		return expectedFnErr
	})

	assert.Equal(t, expectedJoinedErr, err)
}

func Test_TransactionProvider_RunInTx_CommitError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer func() {
		assert.NoError(t, mock.ExpectationsWereMet())
		db.Close()
	}()

	expectedErr := errors.New("commit error")

	mock.ExpectBegin()
	mock.ExpectCommit().WillReturnError(expectedErr)

	txProvider := repository.NewTransactionProvider(db)

	err = txProvider.RunInTx(func(atomicRepos repository.Repositories) error {
		return nil
	})

	assert.Equal(t, expectedErr, err)
}
