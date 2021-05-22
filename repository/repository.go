package repository

import (
	"database/sql"
)

type Repository interface {
	CheckJobByStatus(base string) error
	InsertJobExecution(base string) error
	UpdateJobExecution(base string) error
}

type repositoryImpl struct{}

var (
	instance   = &repositoryImpl{}
	stmtInsert *sql.Stmt
)

func Get() Repository {
	return instance
}

func (r *repositoryImpl) CheckJobByStatus(base string) error {
	_, err := stmtInsert.Exec()
	return err
}

func (r *repositoryImpl) InsertJobExecution(base string) error {
	_, err := stmtInsert.Exec()
	return err
}

func (r *repositoryImpl) UpdateJobExecution(base string) error {
	_, err := stmtInsert.Exec()
	return err
}
