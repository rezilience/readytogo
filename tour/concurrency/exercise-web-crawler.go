package main

import (
	"fmt"
	"sync"
	"time"
)

// Set A thread safe Set data structure using a map and Mutex
type Set struct {
	m   map[string]int
	mux sync.Mutex
}

// Contains returns true if key is present else false
func (u *Set) Contains(key string) bool {
	u.mux.Lock()
	defer u.mux.Unlock()
	_, ok := u.m[key]
	return ok
}

// Add the key to the set
func (u *Set) Add(key string) {
	u.mux.Lock()
	defer u.mux.Unlock()
	u.m[key] = 1
}

// Fetcher returns the body of URL and
// a slice of URLs found on that page.
type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl crawls the given url recursively for given depth and fetcher
func Crawl(url string, depth int, fetcher Fetcher) {
	crawled := &Set{m: make(map[string]int)}
	DeepCrawl(url, depth, fetcher, crawled)
}

// DeepCrawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func DeepCrawl(url string, depth int, fetcher Fetcher, crawled *Set) {
	// TODO: Instead of time.Sleep, use a better way to be notified when all goroutines are complete
	if crawled.Contains(url) {
		return
	} else {
		crawled.Add(url)
	}
	if depth <= 0 {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		go DeepCrawl(u, depth-1, fetcher, crawled)
	}
	time.Sleep(time.Second)
	return
}

func main() {
	Crawl("https://golang.org/", 4, fetcher)
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
