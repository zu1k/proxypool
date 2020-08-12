package tool

import (
	"net/http"
	"time"
)

var httpClient = http.DefaultClient

func init() {
	httpClient.Timeout = time.Second * 10
	http.DefaultClient.Timeout = time.Second * 10
}

func GetHttpClient() *http.Client {
	c := *httpClient
	return &c
}
