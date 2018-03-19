package service

import (
	"fmt"
	"github.com/zhangyuyu/zy-trip-spider/models"
)

func Crawl(url string, depth int, fetcher Fetcher) {
	results := make(chan *models.Result)
	fetched := make(map[string]bool)
	fetch := func(url string, depth int) {
		result, ok := fetcher.Fetch(url)
		if ok {
			results <- &models.Result{url, result.Body, result.Urls, nil, depth}
		} else {
			results <- &models.Result{url, "", nil, fmt.Errorf("Deepth: %d, Not found: %s", depth, url), depth}
		}
	}

	go fetch(url, depth)
	fetched[url] = true

	for fetching := 1; fetching > 0; fetching-- {
		res := <-results

		if res.Err != nil {
			fmt.Println(res.Err)
			continue
		}

		fmt.Printf("Deepth: %d, Found: %s %q\n", res.Depth, res.Url, res.Body)

		// follow links if depth has not been exhausted
		if res.Depth > 0 {
			for _, u := range res.Urls {
				// don't attempt to re-fetch known url, decrement depth
				if !fetched[u] {
					fetching++
					go fetch(u, res.Depth-1)
					fetched[u] = true
				}
			}
		}
	}

	close(results)
}
