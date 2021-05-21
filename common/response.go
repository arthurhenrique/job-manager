package common

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

type Error struct {
	Code  string `json:"code"`
	Error string `json:"error"`
}

func Write(w http.ResponseWriter, body interface{}, status int) {
	if body == nil {
		w.WriteHeader(status)
		return
	}
	bytes, err := json.Marshal(body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logrus.Errorf("error marshaling json body %s", err)
		return
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(status)
	w.Write(bytes)
}

func WriteValidationError(w http.ResponseWriter, err error) {
	Write(w, Error{
		Code:  "VALIDATION",
		Error: err.Error(),
	}, http.StatusBadRequest)
}

func WriteServerError(w http.ResponseWriter, err error, message string) {
	Write(w, Error{
		Code:  "GENERAL",
		Error: err.Error(),
	}, http.StatusInternalServerError)
}
