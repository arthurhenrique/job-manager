package worker

import (
	"hasty-challenge-manager/domain"
	"hasty-challenge-manager/facade"
	"math/rand"
	"time"

	"github.com/sirupsen/logrus"
)

func getRandomSleep() int {
	rand.Seed(time.Now().UnixNano())
	sleep := rand.Intn(domain.Max-domain.Min) + domain.Min
	return sleep
}

func startSleep(jobID string, sleep int) error {
	logrus.Debugf("jobID: %s sleep: %d STARTED", jobID, sleep)
	time.Sleep(time.Duration(sleep) * time.Second)
	logrus.Debugf("jobID: %s sleep: %d FINISHED", jobID, sleep)

	return nil
}

func Trigger(jobID string) error {
	sleep := getRandomSleep()
	err := facade.Get().UpdateSleep(jobID, sleep)
	if err != nil {
		return err
	}
	logrus.Debugf("jobID: %s updated to '%d' sleep", jobID, sleep)

	err = startSleep(jobID, sleep)
	if err != nil {
		err = facade.Get().UpdateStatus(jobID, domain.Failed.String())
		logrus.Errorf("jobID: %s updated to '%s' Status", jobID, domain.Failed.String())
		return err
	}

	err = facade.Get().UpdateStatus(jobID, domain.Success.String())
	if err != nil {
		return err
	}
	logrus.Debugf("jobID: %s updated to '%s' status", jobID, domain.Success.String())

	return nil
}
