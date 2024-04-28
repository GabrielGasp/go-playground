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

type AtomicFn func(atomicRM RepositoryManager) error

type RepositoryManager interface {
	RunAtomic(fn AtomicFn) error
	ExampleRepo() ExampleRepo
	// Add more repositories here
}

type repositoryManager struct {
	db    *sql.DB
	tx    *sql.Tx
	repos repositories
}

func NewRepositoryManager(db *sql.DB) RepositoryManager {
	return repositoryManager{
		db:    db,
		repos: bootstrapRepositories(db),
	}
}

func (rm repositoryManager) RunAtomic(fn AtomicFn) error {
	if rm.tx != nil { // Reuse the existing tx in case of nested RunAtomic calls
		return fn(rm)
	}

	tx, err := rm.db.Begin()
	if err != nil {
		return err
	}

	atomicRM := repositoryManager{
		tx:    tx,
		repos: bootstrapRepositories(tx),
	}

	if fnErr := fn(atomicRM); fnErr != nil {
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
