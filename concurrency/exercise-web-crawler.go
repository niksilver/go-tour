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
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	if depth <= 0 {
		return
	}
    if _, ok := results.r[url]; ok {
        return
    }
    body, urls, err := fetcher.Fetch(url)
    recordResult(url, body, err)
    if err == nil {
        for _, u := range urls {
            Crawl(u, depth-1, fetcher)
        }
    }
}

func recordResult(url string, body string, err error) {
    if err != nil {
        results.r[url] = OneResult{err: err}
    } else {
        results.r[url] = OneResult{body: body}
    }
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
