package api

import (
	"fmt"
	"net/http"
	"time"

	"hasty-challenge-manager/app"

	v1 "hasty-challenge-manager/api/v1"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/sirupsen/logrus"
)

var (
	RootHandlerPath        = "/"
	HealthCheckHandlerPath = "/healthcheck"

	// v1
	JobHandlerV1Path = "/v1/job/{id:[0-9]+}"

	methodNotAllowedErrorMessage = "Invalid request method"
)

func Setup() error {
	port := app.GetEnv("HTTP_PORT")

	r := mux.NewRouter()

	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"*"},
		AllowedMethods: []string{},
		MaxAge:         600,
	})

	r.HandleFunc(RootHandlerPath, RootHandler)
	r.HandleFunc(HealthCheckHandlerPath, HealthCheckHandler)

	r.HandleFunc(JobHandlerV1Path, JobHandlerV1)

	srv := &http.Server{
		ReadTimeout:  time.Duration(30) * time.Second,
		WriteTimeout: time.Duration(30) * time.Second,
		Addr:         ":" + port,
		Handler:      corsHandler.Handler(r),
	}

	logrus.Infof("listening on %s", port)
	logrus.Fatal(srv.ListenAndServe())

	return nil
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, methodNotAllowedErrorMessage, http.StatusMethodNotAllowed)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hasty - Job Manager")
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, methodNotAllowedErrorMessage, http.StatusMethodNotAllowed)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "OK")
}

func JobHandlerV1(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		v1.JobPostAPIHandler(w, r)
	case http.MethodPut:
		v1.JobPutAPIHandler(w, r)
	case http.MethodGet:
		v1.JobGetAPIHandler(w, r)
	default:
		http.Error(w, methodNotAllowedErrorMessage, http.StatusMethodNotAllowed)
	}
}
