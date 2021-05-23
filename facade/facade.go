package facade

import (
	"database/sql"
	"hasty-challenge-manager/repository"
)

// Tx handles tx lifecycle at facade layer
type Tx interface {
	// Begin begins a database tx
	Begin() (*sql.Tx, error)
	// Resolve ensures that the given tx will be resolved
	Resolve(tx *sql.Tx, err *error)
}

type TxImpl struct{}

var tx = &TxImpl{}

// GetTx returns the facade tx instance
func GetTx() Tx {
	return tx
}

func (t *TxImpl) Resolve(tx *sql.Tx, err *error) {
	if p := recover(); p != nil {
		tx.Rollback()
		panic(p)
	} else if *err != nil {
		tx.Rollback()
	} else {
		err := tx.Commit()
		if err != nil {
			panic(err)
		}
	}
}

func (t *TxImpl) Begin() (*sql.Tx, error) {
	return repository.DB.Begin()
}

// WithTx execute the given func within a database tx
func WithTx(txm Tx, fn func(tx *sql.Tx) error) error {
	tx, err := txm.Begin()
	if err != nil {
		return err
	}
	defer txm.Resolve(tx, &err)

	err = fn(tx)
	if err != nil {
		return err
	}

	return nil
}
