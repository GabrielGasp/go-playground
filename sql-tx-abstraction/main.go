package main

import (
	"database/sql"
	"sql-tx-abstraction/repository"
	"sql-tx-abstraction/service"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, _ := sql.Open("sqlite3", ":memory:")

	repos := repository.BootstrapRepositories(db)
	txProvider := repository.NewTransactionProvider(db)

	svc := service.NewExampleService(repos.ExampleRepo, txProvider)

	_ = svc.Do()
}
