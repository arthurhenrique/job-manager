package repository

import (
	"database/sql"
	"time"
)

type Repository interface {
	CheckJobByStatus(tx *sql.Tx, objectId string) error
	InsertJobExecution(tx *sql.Tx, objectId string) (int, error)
	UpdateJobExecution(tx *sql.Tx, objectId string) error
}

type repositoryImpl struct{}

var (
	instance   = &repositoryImpl{}
	stmtInsert *sql.Stmt
	err        error
)

func Get() Repository {
	return instance
}

func (r *repositoryImpl) InsertJobExecution(tx *sql.Tx, objectId string) (int, error) {
	insert, values, err := Psq.Insert("job_manager.job_execution").Columns(`
	object_id,
	status,
	updated_at,
	created_at
	`).Values(
		objectId,
		"PROCESSING",
		time.Now(),
		time.Now(),
	).Suffix("RETURNING id").ToSql()
	if err != nil {
		return 0, err
	}

	ID := 0
	err = tx.QueryRow(insert, values...).Scan(&ID)
	if err != nil {
		return 0, err
	}

	return ID, err
}

func (r *repositoryImpl) UpdateJobExecution(tx *sql.Tx, objectId string) error {
	return nil
}

func (r *repositoryImpl) CheckJobByStatus(tx *sql.Tx, objectId string) error {
	_, err := stmtInsert.Exec()
	return err
}
