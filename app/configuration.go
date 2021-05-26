package app

import (
	"io/ioutil"
	"os"
	"strconv"

	"github.com/sirupsen/logrus"
)

var config = map[string]string{
	"ENVIRONMENT": "dev",
	"LOG_LEVEL":   "DEBUG",
	"HTTP_PORT":   "9000",
	// Job timeout to set job as CANCELLED status
	"JOB_TIMEOUT": "300",
	// Job windows 5 minutes
	"JOB_WINDOW_UPDATE": "5",
	// Database configuration environment vars
	"DATASOURCE_NAME":      "host=localhost port=5432 user=master password=123456 dbname=job_manager sslmode=disable",
	"DB_MAX_IDLE_CONNS":    "5",
	"DB_MAX_OPEN_CONNS":    "10",
	"DB_CONN_MAX_LIFETIME": "300",
}

func init() {
	// Env vars
	for k := range config {
		v := os.Getenv(k)
		if v != "" {
			config[k] = v
		}
	}

	// Logging
	logLevel, err := logrus.ParseLevel(GetEnv("LOG_LEVEL"))
	if err != nil {
		logLevel = logrus.DebugLevel
	}
	logrus.SetLevel(logLevel)
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.RegisterExitHandler(func() {
		logrus.Info("Application will stop probably due to a OS signal")
	})
	if GetEnv("ENVIRONMENT") == "test" {
		logrus.SetOutput(ioutil.Discard)
	} else {
		logrus.SetOutput(os.Stdout)
	}
}

func GetEnv(configKey string) string {
	return config[configKey]
}

func GetEnvInt(configKey string) int {
	i := 0
	i, _ = strconv.Atoi(GetEnv(configKey))
	return i
}
