package facade

import (
	"database/sql"
	"hasty-challenge-manager/repository"
)

var instance = &Facade{
	Tx:   GetTx(),
	Jobs: repository.Get(),
}

type Facade struct {
	Tx   Tx
	Jobs repository.Repository
}

func Get() *Facade {
	return instance
}

// Insert a job execution
func (f *Facade) Insert(objectId string) (ID int, err error) {
	err = WithTx(f.Tx, func(tx *sql.Tx) error {
		var err error

		ID, err = f.Jobs.InsertJobExecution(tx, objectId)
		if err != nil {
			return err
		}

		return nil
	})

	return
}

// Update a job execution
func (f *Facade) Update(objectId string) (ID int, err error) {
	err = WithTx(f.Tx, func(tx *sql.Tx) error {
		var err error

		err = f.Jobs.UpdateJobExecution(tx, objectId)
		if err != nil {
			return err
		}

		return nil
	})

	return
}
