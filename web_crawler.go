/*
In this exercise you'll use Go's concurrency features to parallelize a web crawler.

Modify the Crawl function to fetch URLs in parallel without fetching the same URL twice.

Hint: you can keep a cache of the URLs that have been fetched on a map, but maps alone are not safe for concurrent use!
*/

package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

type fetchResponse struct {
	url string
	body string
	err error
}

type fetchedResults struct {
	responses map[string]fetchResponse
	mux sync.Mutex
	wg sync.WaitGroup
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, results *fetchedResults) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:

	defer results.wg.Done()

	if depth <= 0 {
		return
	}

	results.mux.Lock()
	if _, ok := results.responses[url]; ok {
		results.mux.Unlock()
		// fmt.Println("Already fetched: ", x.url)
		return
	}
	results.mux.Unlock()

	body, urls, err := fetcher.Fetch(url)

	results.mux.Lock()
	if err != nil {
		results.responses[url] = fetchResponse{url, body, err}
		results.mux.Unlock()
		return
	} else {
		results.responses[url] = fetchResponse{url, body, nil}
	}
	results.mux.Unlock()

	for _, u := range urls {
		results.wg.Add(1)
		go Crawl(u, depth-1, fetcher, results)
	}

	return
}

func main() {
	var results fetchedResults
	results.responses = make(map[string]fetchResponse)
	results.wg.Add(1)

	Crawl("https://golang.org/", 4, fetcher, &results)

	results.wg.Wait()

	for url, res := range results.responses {
		if res.err != nil {
			fmt.Println(res.err)
		} else {
			fmt.Printf("found: %s %q\n", url, res.body)
		}
	}
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
