package repository

import (
	"fmt"
)

type ExampleRepo interface {
	Do() error
}

type exampleRepo struct {
	conn dbOps
}

func NewExampleRepo(conn dbOps) ExampleRepo {
	return exampleRepo{conn: conn}
}

// func (r exampleRepo) Do() error {
// 	switch r.conn.(type) {
// 	case *sql.DB:
// 		fmt.Println("Repository using *sql.DB")
// 	case *sql.Tx:
// 		fmt.Println("Repository using *sql.Tx")
// 	}

// 	return nil
// }

func (r exampleRepo) Do() error {
	var result int
	if err := r.conn.QueryRow("SELECT 1").Scan(&result); err != nil {
		return err
	}

	fmt.Println("Result:", result)

	return nil
}
