package service

import (
	"github.com/zhangyuyu/zy-trip-spider/models"
)

type Fetcher interface {
	Fetch(url string) (*models.FakeResult, bool)
}

// fakeFetcher is Fetcher that returns canned results.
type FakeFetcher map[string]*models.FakeResult

func (f *FakeFetcher) Fetch(url string) (*models.FakeResult, bool) {
	result, ok := (*f)[url]
	return result, ok
}
