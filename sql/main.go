package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, _ := sql.Open("sqlite3", ":memory:")

	rm := NewRepositoryManager(db)

	svc := NewExampleService(rm)

	_ = svc.ExecExample()
}
