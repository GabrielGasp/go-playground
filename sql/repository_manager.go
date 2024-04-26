package main

import "database/sql"

type BasicDB interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Prepare(query string) (*sql.Stmt, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
}

type repositories struct {
	ExampleRepo ExampleRepo
	// Add more repositories here
}

func bootstrapRepositories(db BasicDB) repositories {
	return repositories{
		ExampleRepo: NewExampleRepo(db),
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
	return &repositoryManager{
		db:    db,
		repos: bootstrapRepositories(db),
	}
}

func (rm *repositoryManager) RunAtomic(fn func(atomicRepositoryManager RepositoryManager) error) error {
	tx, err := rm.db.Begin()
	if err != nil {
		return err
	}

	atomicRepositoryManager := &repositoryManager{repos: bootstrapRepositories(tx)}

	if err = fn(atomicRepositoryManager); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (rm *repositoryManager) ExampleRepo() ExampleRepo {
	return rm.repos.ExampleRepo
}
