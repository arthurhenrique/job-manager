package v1

import (
	"hasty-challenge-manager/common"
	"net/http"

	"github.com/gorilla/mux"
)

type JobResponse struct {
	JobId    string `json:"job_id"`
	ObjectId string `json:"object_id"`
}

func TriggerAPIHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	vars := mux.Vars(r)
	objectId := vars["id"]

	common.Write(w, JobResponse{JobId: "1", ObjectId: objectId}, http.StatusAccepted)

}
