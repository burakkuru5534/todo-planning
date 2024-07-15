package util

import (
	"net/http"
	"time"
)

func CreateHTTPClient() *http.Client {
	return &http.Client{
		Timeout: time.Second * 10,
	}
}
