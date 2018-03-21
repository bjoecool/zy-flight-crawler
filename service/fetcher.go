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
	result, err := http.Get(url)
	return result, err != nil
}
