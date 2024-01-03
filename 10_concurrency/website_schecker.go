package crawler

type WebsitesChecker func(string) bool

func CheckWebsites(wc WebsitesChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	for _, url := range urls {
		go func() {
			results[url] = wc(url)
		}()
	}
	return results
}
