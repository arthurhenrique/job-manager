package http

import (
	"net/http"
	"time"
)

var Instance *http.Client

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

func CreateClient(readTimeout time.Duration) *http.Client {
	var customHTTPTransport = &http.Transport{
		IdleConnTimeout: 30 * time.Second,
	}

	return &http.Client{
		Transport: customHTTPTransport,
	}
}

func init() {
	readTimeout := time.Duration(10)
	Instance = CreateClient(readTimeout * time.Second)
}
