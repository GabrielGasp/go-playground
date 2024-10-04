package repository

import (
	"database/sql"
	"errors"
)

type TxFn func(repositories Repositories) error

type TransactionProvider interface {
	RunInTx(fn TxFn) error
}

type transactionProvider struct {
	db *sql.DB
}

func NewTransactionProvider(db *sql.DB) TransactionProvider {
	return transactionProvider{
		db: db,
	}
}

func (tp transactionProvider) RunInTx(fn TxFn) error {
	tx, err := tp.db.Begin()
	if err != nil {
		return err
	}

	txRepos := BootstrapRepositories(tx)

	fnErr := fn(txRepos)
	if fnErr == nil {
		return tx.Commit()
	}

	rbErr := tx.Rollback()
	if rbErr != nil {
		return errors.Join(fnErr, rbErr)
	}

	return fnErr
}
