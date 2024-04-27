package repository

import (
	"database/sql"
	"errors"
)

type DBOps interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Prepare(query string) (*sql.Stmt, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
}

type repositories struct {
	ExampleRepo ExampleRepo
	// Add more repositories here
}

func bootstrapRepositories(conn DBOps) repositories {
	return repositories{
		ExampleRepo: NewExampleRepo(conn),
		// Add more repositories here
	}
}

// ----------------------------------------------------------

type RepositoryManager interface {
	RunAtomic(fn func(atomicRepositoryManager RepositoryManager) error) error
	ExampleRepo() ExampleRepo
	// Add more repositories here
}

type repositoryManager struct {
	db    *sql.DB
	repos repositories
}

func NewRepositoryManager(db *sql.DB) RepositoryManager {
	return repositoryManager{
		db:    db,
		repos: bootstrapRepositories(db),
	}
}

func (rm repositoryManager) RunAtomic(fn func(atomicRepositoryManager RepositoryManager) error) error {
	if rm.db == nil { // this means we are already in a transaction
		return fn(rm)
	}

	tx, err := rm.db.Begin()
	if err != nil {
		return err
	}

	atomicRepositoryManager := repositoryManager{repos: bootstrapRepositories(tx)}

	if fnErr := fn(atomicRepositoryManager); fnErr != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return errors.Join(fnErr, rbErr)
		}

		return fnErr
	}

	return tx.Commit()
}

func (rm repositoryManager) ExampleRepo() ExampleRepo {
	return rm.repos.ExampleRepo
}
