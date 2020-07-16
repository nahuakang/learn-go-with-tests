package concurrency

import "time"

// WebsiteChecker is a function
type WebsiteChecker func(string) bool

// CheckWebsites checks if a URL is valid
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	for _, url := range urls {
		go func(u string) {
			results[u] = wc(u)
		}(url)
	}

	time.Sleep(2 * time.Second) // let the goroutines do work

	return results
}
