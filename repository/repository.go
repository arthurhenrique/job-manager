package repository

import (
	"database/sql"
	"hasty-challenge-manager/domain"
	"time"

	sq "github.com/Masterminds/squirrel"
)

type Repository interface {
	InsertJobExecution(tx *sql.Tx, objectID string) (string, error)
	UpdateJobStatus(tx *sql.Tx, objectID, status string) error
	UpdateJobSleep(tx *sql.Tx, objectID string, sleep int) error
	CheckJobByStatus(tx *sql.Tx, objectID string) error
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

func (r *repositoryImpl) InsertJobExecution(tx *sql.Tx, objectID string) (string, error) {
	insert, values, err := Psq.Insert("job_manager.job_execution").Columns(`
	object_id,
	status,
	updated_at,
	created_at
	`).Values(
		objectID,
		domain.Processing.String(),
		time.Now(),
		time.Now(),
	).Suffix("RETURNING id").ToSql()
	if err != nil {
		return "", err
	}

	ID := ""
	err = tx.QueryRow(insert, values...).Scan(&ID)
	if err != nil {
		return "", err
	}

	return ID, err
}

func (r *repositoryImpl) UpdateJobStatus(tx *sql.Tx, objectID, status string) error {
	set, values, err := Psq.Update("job_manager.job_execution").
		Set("status", status).
		Set("updated_at", time.Now()).
		Where(sq.Eq{"id": objectID}).ToSql()
	if err != nil {
		return nil
	}

	_, err = tx.Exec(set, values...)
	if err != nil {
		return err
	}

	return nil
}

func (r *repositoryImpl) UpdateJobSleep(tx *sql.Tx, objectID string, sleep int) error {
	set, values, err := Psq.Update("job_manager.job_execution").
		Set("sleep", sleep).
		Set("updated_at", time.Now()).
		Where(sq.Eq{"id": objectID}).ToSql()
	if err != nil {
		return nil
	}

	_, err = tx.Exec(set, values...)
	if err != nil {
		return err
	}

	return nil
}

func (r *repositoryImpl) CheckJobByStatus(tx *sql.Tx, objectID string) error {
	_, err := stmtInsert.Exec()
	return err
}
