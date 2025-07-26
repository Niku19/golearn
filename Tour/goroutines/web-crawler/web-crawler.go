package main

import (
	"fmt"
	"sync" // Import the sync package for Mutex and WaitGroup
)

// Fetcher interface (provided in the exercise)
type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// --- Concurrency Safety for Visited URLs ---

// fetchedData holds the shared state for URLs that have been fetched.
// It uses a Mutex to protect concurrent access to the map.
var fetchedData = struct {
	mu   sync.Mutex      // Mutex to synchronize access to the 'urls' map
	urls map[string]bool // Map to store visited URLs (true if visited)
}{
	urls: make(map[string]bool), // Initialize the map
}

// --- Modified Crawl Function ---

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
// It now takes a WaitGroup pointer to coordinate goroutines.
func Crawl(url string, depth int, fetcher Fetcher, wg *sync.WaitGroup) {
	// Decrement the WaitGroup counter when this goroutine finishes.
	// This ensures that the main goroutine knows when all child goroutines are done.
	defer wg.Done()

	if depth <= 0 {
		return // Base case: stop crawling if depth limit reached
	}

	// Check if this URL has already been fetched or is being fetched.
	// Lock the mutex before accessing the shared map.
	fetchedData.mu.Lock()
	if fetchedData.urls[url] {
		fetchedData.mu.Unlock() // Release the lock before returning
		return                  // URL already processed, so skip
	}
	fetchedData.urls[url] = true // Mark this URL as being fetched
	fetchedData.mu.Unlock()      // Release the lock immediately after marking

	// Fetch the URL's content.
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err) // Print error but don't stop other crawls
		return
	}
	fmt.Printf("found: %s %q\n", url, body)

	// For each found URL, kick off a new goroutine to crawl it.
	for _, u := range urls {
		// Increment the WaitGroup counter for each new goroutine.
		wg.Add(1)
		// Start a new goroutine for the recursive crawl.
		go Crawl(u, depth-1, fetcher, wg)
	}
	// The function implicitly returns here after the loop.
}

// --- Main Function (Modified to use concurrency) ---

func main() {
	var wg sync.WaitGroup // Create a WaitGroup instance.

	// Initial call to Crawl.
	// Increment the WaitGroup counter for the very first crawl operation.
	wg.Add(1)
	go Crawl("https://golang.org/", 4, fetcher, &wg) // Start the initial crawl in a goroutine.

	// Wait for all goroutines (the initial one and all its children) to complete.
	wg.Wait()
	fmt.Println("\nWeb crawl finished.")
}

// --- Provided fakeFetcher (for testing purposes) ---

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

// fetcher is a populated fakeFetcher instance.
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
			"https://golang.org/", // This URL will be fetched again, but skipped by the cache
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
	"https://golang.org/cmd/": &fakeResult{
		"Commands",
		[]string{
			"https://golang.org/",
		},
	},
}
