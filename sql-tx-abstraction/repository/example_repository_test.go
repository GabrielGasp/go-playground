package repository_test

import (
	"errors"
	"sql-playground/repository"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func Test_ExampleRepo_Do_Success(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer mockDB.Close()
	defer mock.ExpectationsWereMet()

	mock.ExpectQuery("SELECT 1").WillReturnRows(sqlmock.NewRows([]string{"result"}).AddRow(1))

	exampleRepo := repository.NewExampleRepo(mockDB)

	err = exampleRepo.Do()

	assert.NoError(t, err)
}

func Test_ExampleRepo_Do_QueryRowError(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer mockDB.Close()
	defer mock.ExpectationsWereMet()

	expectedErr := errors.New("QueryRow Error")

	mock.ExpectQuery("SELECT 1").WillReturnError(expectedErr)

	exampleRepo := repository.NewExampleRepo(mockDB)

	err = exampleRepo.Do()

	assert.Equal(t, expectedErr, err)
}
