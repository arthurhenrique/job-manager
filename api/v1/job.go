package v1

import (
	"fmt"
	"hasty-challenge-manager/common"
	"hasty-challenge-manager/domain"
	"hasty-challenge-manager/facade"
	"hasty-challenge-manager/worker"
	"net/http"

	"github.com/gorilla/mux"
)

type JobResponse struct {
	JobID    string `json:"job_id"`
	ObjectID string `json:"object_id"`
}

type JobExecutionResponse struct {
	JobExecution domain.JobExecution `json:"job_execution"`
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

func JobGetAPIHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	vars := mux.Vars(r)
	objectID := vars["id"]

	jobExecution, err := facade.Get().Select(objectID)
	if err != nil {
		return
	}

	common.Write(w, JobExecutionResponse{jobExecution}, http.StatusOK)

}
