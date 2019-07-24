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
    prefix := fmt.Sprintf("%v, %v:", depth, url)
    fmt.Printf("%v Entering...\n", prefix)
	if depth <= 0 {
        fmt.Printf("%v Returning early (zero depth)\n", prefix)
		return
	}
    if _, ok := results.r[url]; ok {
        fmt.Printf("%v Returning early (fetched previously)\n", prefix)
        return
    }
    fmt.Printf("%v Fetching\n", prefix)
    body, urls, err := fetcher.Fetch(url)
    fmt.Printf("%v Fetched\n", prefix)
    recordResult(url, body, err)
    if err == nil {
        fmt.Printf("%v Got URLs to crawl\n", prefix)
        done := make(chan string, len(urls))
        for i, u := range urls {
            fmt.Printf("%v %v: About to go crawl for %v\n", prefix, i, u)
            go func(u string) {
                fmt.Printf("%v In go routine, about to crawl for %v-%v\n", prefix, i, u)
                Crawl(u, depth-1, fetcher)
                fmt.Printf("%v In go routine, done crawl for %v-%v\n", prefix, i, u)
                done <- u
            }(u)
            fmt.Printf("%v Set off go routine for %v-%v\n", prefix, i, u)
        }
        for i := 0; i < len(urls); i++ {
            fmt.Printf("%v Waiting for done %v/%v\n", prefix, i+1, len(urls))
            u := <-done
            fmt.Printf("%v Received done %v/%v: %v\n", prefix, i+1, len(urls), u)
        }
        fmt.Printf("%v Returning after crawls\n", prefix)
    } else {
        fmt.Printf("%v Got error, returning without crawls\n", prefix)
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
