package config

import (
	"net/http"
	"time"
)

func NewHttpClient() *http.Client {
	return &http.Client{
		Timeout: 10 * time.Second,
	}
}
