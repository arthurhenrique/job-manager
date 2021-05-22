package api_test

import (
	"net/http"
	"os"
	"testing"

	"hasty-challenge-manager/api"
	"hasty-challenge-manager/repository"
	"hasty-challenge-manager/test"

	"github.com/sirupsen/logrus"
)

func TestMain(m *testing.M) {
	go api.Setup()
	os.Exit(m.Run())
}

func TestHandlers(t *testing.T) {
	err := repository.Setup()
	if err != nil {
		logrus.Fatalf("error getting db, err: %v", err)
	}
	test.MockHTTP(t, func(w http.ResponseWriter, r *http.Request) {
		fixture := `{""}`
		if r.URL.Path == "/v1/trigger/999" && r.Method == "POST" {
			body := []byte(fixture)
			w.WriteHeader(http.StatusOK)
			w.Write(body)
			return
		}
		w.WriteHeader(http.StatusNotImplemented)
	})

	testCases := []test.APITestCase{
		{
			Name:   "should return OK status with body POST",
			Method: http.MethodPost,
			Route:  "http://localhost:9000/v1/trigger/999",
			Status: http.StatusOK,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, tc.Run)
	}
}
