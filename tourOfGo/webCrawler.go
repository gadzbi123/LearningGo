package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

type urls map[string]string

type CachedUrls struct {
	fetchedUrls urls
	mtx         sync.Mutex
}

var cachedUrls = &CachedUrls{fetchedUrls: make(urls)}

func (cu *CachedUrls) add(url, body string) {
	cu.mtx.Lock()
	cu.fetchedUrls[url] = body
	cu.mtx.Unlock()
}

func (cu *CachedUrls) read() []string {
	arr := []string{}
	cu.mtx.Lock()
	for k, _ := range cu.fetchedUrls {
		arr = append(arr, k)
	}
	cu.mtx.Unlock()
	return arr
}

func (cu *CachedUrls) exists(url string) bool {
	res := false
	cu.mtx.Lock()
	if _, ok := cu.fetchedUrls[url]; ok {
		res = true
	}
	cu.mtx.Unlock()
	return res
}

func Crawl(url string, depth int, fetcher Fetcher, wg *sync.WaitGroup) {
	if depth <= 0 {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	cachedUrls.add(url, body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %v\n", url, urls)

	// fmt.Printf("cached: %s\n", cachedUrls.read())
	for _, u := range urls {
		if !cachedUrls.exists(u) {
			wg.Add(1)
			go func() {
				Crawl(u, depth-1, fetcher, wg)
				wg.Done()
			}()
		}
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	Crawl("https://golang.org/", 4, fetcher, &wg)
	wg.Done()

	wg.Wait()
	cachedUrls.mtx.Lock()
	fmt.Printf("RESULT: %v", cachedUrls.fetchedUrls)
	cachedUrls.mtx.Unlock()
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
