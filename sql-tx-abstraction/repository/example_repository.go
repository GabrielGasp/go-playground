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

func newExampleRepo(conn dbOps) ExampleRepo {
	return exampleRepo{conn: conn}
}

func (r exampleRepo) Do() error {
	fmt.Printf("Connection type is: %T\n", r.conn)

	var result int
	if err := r.conn.QueryRow("SELECT 1").Scan(&result); err != nil {
		return err
	}

	fmt.Println("Result:", result)

	return nil
}
