package repository

import (
	"database/sql"
	"hasty-challenge-manager/app"
	"sync/atomic"
	"time"

	_ "github.com/jackc/pgx/v4/stdlib"
)

var (
	// DB is a db instance
	DB *sql.DB

	dbReady = int32(0)
)

func dbIsReady() {
	atomic.StoreInt32(&dbReady, 1)
}

func isDBReady() bool {
	return atomic.LoadInt32(&dbReady) == 1
}

func setupDB() error {
	if isDBReady() {
		return nil
	}
	connectionString := app.GetEnv("DATASOURCE_NAME")
	db, err := sql.Open("pgx", connectionString)
	if err != nil {
		return err
	}
	db.SetMaxIdleConns(app.GetEnvInt("DATASOURCE_NAME"))
	db.SetMaxOpenConns(app.GetEnvInt("DB_MAX_OPEN_CONNS"))
	db.SetConnMaxLifetime(time.Duration(app.GetEnvInt("DB_CONN_MAX_LIFETIME")) * time.Second)

	DB = db
	err = db.Ping()
	if err != nil {
		return err
	}
	dbIsReady()

	return nil
}

func Setup() error {
	err := setupDB()
	if err != nil {
		return err
	}

	return nil
}
