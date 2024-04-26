package main

import (
	"database/sql"
	"fmt"
)

type ExampleRepo interface {
	ExecExample() error
}

type exampleRepo struct {
	conn BasicDB
}

func NewExampleRepo(conn BasicDB) ExampleRepo {
	return &exampleRepo{conn: conn}
}

func (r *exampleRepo) ExecExample() error {
	switch r.conn.(type) {
	case *sql.DB:
		fmt.Println("Repository using *sql.DB")
	case *sql.Tx:
		fmt.Println("Repository using *sql.Tx")
	}

	return nil
}
