package schedule

import (
	"fmt"
	"hasty-challenge-manager/domain"
	"hasty-challenge-manager/facade"
)

func Run() error {
	err := cancelJobsByTimeout()
	if err != nil {
		return err
	}

	return nil
}

func cancelJobsByTimeout() error {
	jobsProcessing, _ := facade.Get().SelectByStatus("PROCESSING")
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
