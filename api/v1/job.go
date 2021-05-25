package v1

import (
	"fmt"
	"hasty-challenge-manager/common"
	"hasty-challenge-manager/facade"
	"hasty-challenge-manager/worker"
	"net/http"

	"github.com/gorilla/mux"
)

type JobResponse struct {
	JobID    string `json:"job_id"`
	ObjectID string `json:"object_id"`
}

func JobPostAPIHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	vars := mux.Vars(r)
	objectID := vars["id"]

	jobID, err := facade.Get().Insert(objectID)
	if err != nil {
		return
	}

	go worker.Trigger(jobID)

	common.Write(w, JobResponse{JobID: fmt.Sprint(jobID), ObjectID: objectID}, http.StatusOK)
}

func JobPutAPIHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
}
