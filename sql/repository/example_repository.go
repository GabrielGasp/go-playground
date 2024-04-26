package repository

import (
	"database/sql"
	"fmt"
)

type ExampleRepo interface {
	Do() error
}

type exampleRepo struct {
	conn DBOps
}

func NewExampleRepo(conn DBOps) ExampleRepo {
	return exampleRepo{conn: conn}
}

// Use this one when running the main.go
func (r exampleRepo) Do() error {
	switch r.conn.(type) {
	case *sql.DB:
		fmt.Println("Repository using *sql.DB")
	case *sql.Tx:
		fmt.Println("Repository using *sql.Tx")
	}

	return nil
}

// // Use this one when running tests
// func (r exampleRepo) Do() error {
// 	var result int
// 	if err := r.conn.QueryRow("SELECT 1").Scan(&result); err != nil {
// 		return err
// 	}

// 	fmt.Println("Result:", result)

// 	return nil
// }
