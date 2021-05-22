package v1

import (
	"fmt"
	"hasty-challenge-manager/common"
	"hasty-challenge-manager/facade"
	"net/http"

	"github.com/gorilla/mux"
)

type JobResponse struct {
	JobId    string `json:"job_id"`
	ObjectId string `json:"object_id"`
}

func TriggerPostAPIHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	vars := mux.Vars(r)
	objectId := vars["id"]

	ID, err := facade.Get().Insert(objectId)
	if err != nil {
		return
	}

	common.Write(w, JobResponse{JobId: fmt.Sprint(ID), ObjectId: objectId}, http.StatusAccepted)
}

func TriggerPutAPIHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	vars := mux.Vars(r)
	objectId := vars["id"]

	ID, err := facade.Get().Update(objectId)
	if err != nil {
		return
	}

	common.Write(w, JobResponse{JobId: fmt.Sprint(ID), ObjectId: objectId}, http.StatusAccepted)
}
