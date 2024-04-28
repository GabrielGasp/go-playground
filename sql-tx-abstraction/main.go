package main

import (
	"database/sql"
	"sql-tx-abstraction/repository"
	"sql-tx-abstraction/service"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, _ := sql.Open("sqlite3", ":memory:")

	rm := repository.NewRepositoryManager(db)

	svc := service.NewExampleService(rm)

	_ = svc.Do()
}
