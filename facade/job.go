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
func (f *Facade) Insert(objectID string) (jobID string, err error) {
	err = WithTx(f.Tx, func(tx *sql.Tx) error {
		var err error

		jobID, err = f.Jobs.InsertJobExecution(tx, objectID)
		if err != nil {
			return err
		}

		return nil
	})

	return
}

// UpdateStatus update the status given a job execution
func (f *Facade) UpdateStatus(objectID, status string) (err error) {
	err = WithTx(f.Tx, func(tx *sql.Tx) error {
		var err error

		err = f.Jobs.UpdateJobStatus(tx, objectID, status)
		if err != nil {
			return err
		}

		return nil
	})

	return
}

// UpdateSleep update the sleep given a job execution
func (f *Facade) UpdateSleep(objectID string, sleep int) (err error) {
	err = WithTx(f.Tx, func(tx *sql.Tx) error {
		var err error

		err = f.Jobs.UpdateJobSleep(tx, objectID, sleep)
		if err != nil {
			return err
		}

		return nil
	})

	return
}
