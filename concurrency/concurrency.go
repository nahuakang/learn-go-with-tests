package concurrency

// WebsiteChecker is a function
type WebsiteChecker func(string) bool

// CheckWebsites checks if a URL is valid
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	for _, url := range urls {
		results[url] = wc(url)
	}

	return results
}
