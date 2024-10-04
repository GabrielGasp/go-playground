package repository_test

import (
	"errors"
	"sql-tx-abstraction/repository"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func Test_ExampleRepo_Do_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer func() {
		assert.NoError(t, mock.ExpectationsWereMet())
		db.Close()
	}()

	mock.ExpectQuery("SELECT 1").WillReturnRows(sqlmock.NewRows([]string{"result"}).AddRow(1))

	repos := repository.BootstrapRepositories(db)

	err = repos.ExampleRepo.Do()

	assert.NoError(t, err)
}

func Test_ExampleRepo_Do_QueryRowError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer func() {
		assert.NoError(t, mock.ExpectationsWereMet())
		db.Close()
	}()

	expectedErr := errors.New("QueryRow Error")

	mock.ExpectQuery("SELECT 1").WillReturnError(expectedErr)

	repos := repository.BootstrapRepositories(db)

	err = repos.ExampleRepo.Do()

	assert.Equal(t, expectedErr, err)
}
