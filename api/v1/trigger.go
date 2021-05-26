package v1

import (
	"fmt"
	"hasty-challenge-manager/common"
	"hasty-challenge-manager/facade"
	"hasty-challenge-manager/worker"
	"net/http"

	"github.com/gorilla/mux"
)

type TriggerResponse struct {
	JobID string `json:"job_id"`
}

func TriggerPostAPIHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	vars := mux.Vars(r)
	objectID := vars["id"]

	jobID, err := facade.Get().Insert(objectID)
	if err != nil {
		common.WriteServerError(w, err, "error trigger this job")
		return
	}

	go worker.Trigger(jobID)

	common.Write(w, TriggerResponse{JobID: fmt.Sprint(jobID)}, http.StatusOK)
}

func TriggerPutAPIHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	vars := mux.Vars(r)
	objectID := vars["id"]

	err := facade.Get().CheckTimeWindow(objectID)
	if err != nil {
		common.WriteValidationError(w, err)
		return
	}

	jobID, err := facade.Get().Update(objectID)
	if err != nil {
		common.WriteServerError(w, err, "error trigger this job")
		return
	}

	go worker.Trigger(jobID)

	common.Write(w, TriggerResponse{JobID: fmt.Sprint(jobID)}, http.StatusOK)
}
