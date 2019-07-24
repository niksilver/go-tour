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

type OneResult struct {
    body string
    err error
}

type Results struct {
    r map[string]OneResult
    mux sync.Mutex
}

var results = Results{
    r: make(map[string]OneResult),
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	if depth <= 0 {
		return
	}
    if _, ok := results.r[url]; ok {
        return
    }
    body, urls, err := fetcher.Fetch(url)
    recordResult(url, body, err)
    if err == nil {
        parallelCrawl(urls, depth-1, fetcher)
    }
}

func parallelCrawl(urls []string, depth int, fetcher Fetcher) {
    done := make(chan string, len(urls))
    for _, u := range urls {
        go func(u string) {
            Crawl(u, depth, fetcher)
            done <- u
        }(u)
    }
    for i := 0; i < len(urls); i++ {
        <-done
    }
}

func recordResult(url string, body string, err error) {
    results.mux.Lock()
    if err != nil {
        results.r[url] = OneResult{err: err}
    } else {
        results.r[url] = OneResult{body: body}
    }
    results.mux.Unlock()
 }

func main() {
	Crawl("https://golang.org/", 4, fetcher)
    for url, result := range results.r {
        fmt.Printf("%v: ", url)
        if result.err == nil {
            fmt.Printf("Body %v\n", result.body)
        } else {
            fmt.Printf("Error %v\n", result.err.Error())
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
