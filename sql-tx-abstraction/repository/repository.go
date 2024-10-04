package repository

import "database/sql"

type dbOps interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Prepare(query string) (*sql.Stmt, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
}

type Repositories struct {
	ExampleRepo ExampleRepo
	// Add more repositories here
}

func BootstrapRepositories(conn dbOps) Repositories {
	return Repositories{
		ExampleRepo: newExampleRepo(conn),
		// Add more repositories here
	}
}
