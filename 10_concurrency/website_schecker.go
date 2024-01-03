package crawler

import "time"

type WebsitesChecker func(string) bool

func CheckWebsites(wc WebsitesChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	for _, url := range urls {
		go func(u string) {
			results[u] = wc(u)
		}(url)
	}
	time.Sleep(2 * time.Second)
	return results
}
