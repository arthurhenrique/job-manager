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
	"JOB_TIMEOUT": "300",
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
