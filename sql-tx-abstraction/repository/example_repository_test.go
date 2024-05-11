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

	exampleRepo := repository.NewExampleRepo(db)

	err = exampleRepo.Do()

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

	exampleRepo := repository.NewExampleRepo(db)

	err = exampleRepo.Do()

	assert.Equal(t, expectedErr, err)
}
