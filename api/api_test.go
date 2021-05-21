package api_test

import (
	"net/http"
	"os"
	"testing"

	"hasty-challenge-manager/api"
	"hasty-challenge-manager/test"
)

func TestMain(m *testing.M) {
	go api.Setup()
	os.Exit(m.Run())
}

var fixture string

func TestHandlers(t *testing.T) {

	test.MockHTTP(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/v1/trigger" && r.Method == "POST" {
			body := []byte(fixture)
			w.WriteHeader(http.StatusAccepted)
			w.Write(body)
			return
		}
		w.WriteHeader(http.StatusNotImplemented)
	})

	testCases := []test.APITestCase{
		{
			Name:   "should return OK status with body - POST (TriggerAPIHandler)",
			Method: http.MethodGet,
			Route:  "http://localhost:9001/v1/trigger",
			Body:   `{""}`,
			Status: http.StatusAccepted,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, tc.Run)
	}
}
