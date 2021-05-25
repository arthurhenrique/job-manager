package v1

import (
	"hasty-challenge-manager/common"
	"hasty-challenge-manager/domain"
	"hasty-challenge-manager/facade"
	"net/http"

	"github.com/gorilla/mux"
)

type JobExecutionResponse struct {
	JobExecution domain.JobExecution `json:"job_execution"`
}

func JobGetAPIHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	vars := mux.Vars(r)
	jobID := vars["id"]

	jobExecution, err := facade.Get().Select(jobID)
	if err != nil {
		return
	}

	common.Write(w, JobExecutionResponse{jobExecution}, http.StatusOK)

}
