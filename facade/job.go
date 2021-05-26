package facade

import (
	"database/sql"
	"errors"
	"hasty-challenge-manager/app"
	"hasty-challenge-manager/domain"
	"hasty-challenge-manager/repository"
	"time"
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

// Update updates a job execution
func (f *Facade) Update(objectID string) (jobID string, err error) {
	err = WithTx(f.Tx, func(tx *sql.Tx) error {
		var err error

		jobID, err = f.Jobs.UpdateJobExecution(tx, objectID)
		if err != nil {
			return err
		}

		return nil
	})

	return
}

// UpdateStatus update the status given a job execution
func (f *Facade) UpdateStatus(jobID, status string) (err error) {
	err = WithTx(f.Tx, func(tx *sql.Tx) error {
		var err error

		err = f.Jobs.UpdateJobStatus(tx, jobID, status)
		if err != nil {
			return err
		}

		return nil
	})

	return
}

// UpdateSleep update the sleep given a job execution
func (f *Facade) UpdateSleep(jobID string, sleep int) (err error) {
	err = WithTx(f.Tx, func(tx *sql.Tx) error {
		var err error

		err = f.Jobs.UpdateJobSleep(tx, jobID, sleep)
		if err != nil {
			return err
		}

		return nil
	})

	return
}

// Select return job execution by job ID
func (f *Facade) Select(jobID string) (result domain.JobExecution, err error) {
	err = WithTx(f.Tx, func(tx *sql.Tx) error {
		var err error

		result, err = f.Jobs.FindByJobID(tx, jobID)
		if err != nil {
			return err
		}

		return nil
	})

	return
}

// CheckTimeWindow window
func (f *Facade) CheckTimeWindow(objectID string) (isOk bool, err error) {
	err = WithTx(f.Tx, func(tx *sql.Tx) error {
		isOk = true
		result, err := f.Jobs.FindByObjectID(tx, objectID)
		if err != nil {
			isOk = false
			return err
		}

		jobWindowUpdate := time.Minute * time.Duration(app.GetEnvInt("JOB_WINDOW_UPDATE"))
		if time.Now().Sub(result.UpdatedAt) <= jobWindowUpdate {
			isOk = false
			return errors.New("This should be executed only once in a time window of 5 minutes")
		}

		return nil
	})

	return
}

// CheckTimeoutProcessing should verify using JOB_TIMEOUT to cancel
func (f *Facade) CheckTimeoutProcessing(objectID string) (isOk bool, err error) {
	err = WithTx(f.Tx, func(tx *sql.Tx) error {
		isOk = true
		result, err := f.Jobs.FindByObjectID(tx, objectID)
		if err != nil {
			isOk = false
			return err
		}

		jobWindowUpdate := time.Minute * time.Duration(app.GetEnvInt("JOB_TIMEOUT"))
		if time.Now().Sub(result.UpdatedAt) <= jobWindowUpdate {
			isOk = false
		}

		return nil
	})

	return
}

// SelectByStatus return jobs execution by status
func (f *Facade) SelectByStatus(status string) (result []domain.JobExecution, err error) {
	err = WithTx(f.Tx, func(tx *sql.Tx) error {
		var err error

		result, err = f.Jobs.FindByStatus(tx, status)
		if err != nil {
			return err
		}

		return nil
	})

	return
}
