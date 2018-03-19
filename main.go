package main

import (
	"github.com/zhangyuyu/zy-trip-spider/service"
	"github.com/zhangyuyu/zy-trip-spider/models"
)

func main() {
	service.Crawl("http://golang.org/", 4, fetcher)
}

// fetcher is a populated fakeFetcher.
var fetcher = &service.FakeFetcher{
	"http://golang.org/": &models.FakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &models.FakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &models.FakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &models.FakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}
