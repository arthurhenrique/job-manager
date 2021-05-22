package main

import (
	"hasty-challenge-manager/cmd"
	"hasty-challenge-manager/repository"

	"github.com/sirupsen/logrus"
)

func main() {
	// Database
	err := repository.Setup()
	if err != nil {
		logrus.Fatalf("error getting db, err: %v", err)
	}

	cmd.Execute()
}
