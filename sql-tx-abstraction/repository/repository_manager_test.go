package repository_test

import (
	"errors"
	"sql-tx-abstraction/repository"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func Test_RepositoryManager_RunAtomic_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer func() {
		assert.NoError(t, mock.ExpectationsWereMet())
		db.Close()
	}()

	mock.ExpectBegin()
	mock.ExpectCommit()

	rm := repository.NewRepositoryManager(db)

	err = rm.RunAtomic(func(atomicRM repository.RepositoryManager) error {
		return nil
	})

	assert.NoError(t, err)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func Test_RepositoryManager_RunAtomic_BeginError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer func() {
		assert.NoError(t, mock.ExpectationsWereMet())
		db.Close()
	}()

	expectedErr := errors.New("begin error")

	mock.ExpectBegin().WillReturnError(expectedErr)

	rm := repository.NewRepositoryManager(db)

	err = rm.RunAtomic(func(atomicRM repository.RepositoryManager) error {
		return nil
	})

	assert.Equal(t, expectedErr, err)
}

func Test_RepositoryManager_RunAtomic_FnError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer func() {
		assert.NoError(t, mock.ExpectationsWereMet())
		db.Close()
	}()

	mock.ExpectBegin()
	mock.ExpectRollback()

	rm := repository.NewRepositoryManager(db)

	expectedErr := errors.New("fn error")

	err = rm.RunAtomic(func(atomicRM repository.RepositoryManager) error {
		return expectedErr
	})

	assert.Equal(t, expectedErr, err)
}

func Test_RepositoryManager_RunAtomic_RollbackError(t *testing.T) {
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

	rm := repository.NewRepositoryManager(db)

	err = rm.RunAtomic(func(atomicRM repository.RepositoryManager) error {
		return expectedFnErr
	})

	assert.Equal(t, expectedJoinedErr, err)
}

func Test_RepositoryManager_RunAtomic_CommitError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer func() {
		assert.NoError(t, mock.ExpectationsWereMet())
		db.Close()
	}()

	expectedErr := errors.New("commit error")

	mock.ExpectBegin()
	mock.ExpectCommit().WillReturnError(expectedErr)

	rm := repository.NewRepositoryManager(db)

	err = rm.RunAtomic(func(atomicRM repository.RepositoryManager) error {
		return nil
	})

	assert.Equal(t, expectedErr, err)
}

func Test_RepositoryManager_RunAtomic_NestedRunAtomic(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer func() {
		assert.NoError(t, mock.ExpectationsWereMet())
		db.Close()
	}()

	mock.ExpectBegin()
	mock.ExpectCommit()

	rm := repository.NewRepositoryManager(db)

	err = rm.RunAtomic(func(atomicRM repository.RepositoryManager) error {
		return atomicRM.RunAtomic(func(nestedAtomicRM repository.RepositoryManager) error {
			assert.Equal(t, atomicRM, nestedAtomicRM)
			return nil
		})
	})

	assert.NoError(t, err)
}

func Test_RepositoryManager_Getters(t *testing.T) {
	db, _, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	rm := repository.NewRepositoryManager(db)

	repos := map[string]any{
		"Example Repo": rm.ExampleRepo(),
		// Add more repositories here
	}

	for name, repository := range repos {
		t.Run(name, func(t *testing.T) {
			assert.NotNil(t, repository)
		})
	}
}
