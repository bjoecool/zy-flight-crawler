package service

import (
	"net/http"
)

type Fetcher interface {
	Fetch(url string) (*http.Response, bool)
}

type TripFetcher struct {
}

func (fetcher *TripFetcher) Fetch(url string) (*http.Response, bool) {
	httpResponse, err := http.Get(url)
	return httpResponse, err != nil
}
