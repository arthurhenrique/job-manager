package v1

import (
	"hasty-challenge-manager/common"
	"net/http"
)

func TriggerAPIHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	fixture := `{"job_id": "1"}`
	body := []byte(fixture)

	common.Write(w, body, http.StatusAccepted)
}
