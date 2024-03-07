package config

import (
	"net/http"
	"time"
)

type HttpClient interface {
	Do(r *http.Request) (*http.Response, error)
}

func NewHttpclient() HttpClient {
	return &http.Client{
		Timeout:   10 * time.Second,
		Transport: getTransport(),
	}
}

func getTransport() *http.Transport {
	t := http.DefaultTransport.(*http.Transport).Clone()
	t.MaxIdleConns = 100
	t.MaxConnsPerHost = 100
	t.MaxIdleConnsPerHost = 100
	return t
}
