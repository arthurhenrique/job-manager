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
	FindByJobID(tx *sql.Tx, jobID string) (domain.JobExecution, error)
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

func CloseRows(rows *sql.Rows) {
	if rows != nil {
		rows.Close()
	}
}

func (r *repositoryImpl) scanRow(rows *sql.Rows) (domain.JobExecution, error) {
	result := domain.JobExecution{}
	err := rows.Scan(
		&result.ID,
		&result.ObjectID,
		&result.Sleep,
		&result.Status,
		&result.CreatedAt,
		&result.UpdatedAt,
	)
	if err != nil {
		return result, err
	}

	return result, nil
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

func (r *repositoryImpl) FindByJobID(tx *sql.Tx, jobID string) (domain.JobExecution, error) {
	query, values, err := Psq.Select("id,object_id,sleep,status,created_at,updated_at").From("job_manager.job_execution").Where(sq.Eq{"id": jobID}).ToSql()
	if err != nil {
		return domain.JobExecution{}, err
	}

	var rows *sql.Rows
	if tx == nil {
		rows, err = DB.Query(query, values...)
	} else {
		rows, err = tx.Query(query, values...)
	}
	if err != nil {
		return domain.JobExecution{}, err
	}
	defer CloseRows(rows)

	if rows.Next() {
		result, err := r.scanRow(rows)
		if err != nil {
			return domain.JobExecution{}, err
		}
		return result, nil
	}

	return domain.JobExecution{}, sql.ErrNoRows
}
