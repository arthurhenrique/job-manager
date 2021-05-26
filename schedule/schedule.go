package schedule

import (
	"fmt"
	"hasty-challenge-manager/domain"
	"hasty-challenge-manager/facade"
	"hasty-challenge-manager/worker"

	"github.com/sirupsen/logrus"
)

func Run() error {
	err := cancelJobsByTimeout()
	if err != nil {
		logrus.Error("cancelJobsByTimeout error %v", err)
		return err
	}

	err = retryJobsFailed()
	if err != nil {
		logrus.Error("retryJobsFailed error %v", err)
		return err
	}

	return nil
}

func retryJobsFailed() error {
	jobsFailed, _ := facade.Get().SelectByStatus(domain.Failed.String())
	for _, v := range jobsFailed {
		worker.Trigger(fmt.Sprint(v.ID))
	}
	return nil
}

func cancelJobsByTimeout() error {
	jobsProcessing, _ := facade.Get().SelectByStatus(domain.Processing.String())
	for _, v := range jobsProcessing {
		if ok, _ := facade.Get().CheckTimeoutProcessing(fmt.Sprint(v.ObjectID)); ok {
			err := facade.Get().UpdateStatus(fmt.Sprint(v.ID), domain.Cancelled.String())
			if err != nil {
				return err
			}
		}
	}
	return nil
}
