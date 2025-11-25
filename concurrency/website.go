package concurrency

type WebsiteChecker func(string) bool
type checkResult struct {
	url     string
	success bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan checkResult)

	for _, url := range urls {
		go func() {
			resultChannel <- checkResult{url, wc(url)}
		}()
	}

	for range urls {
		r := <-resultChannel
		results[r.url] = r.success
	}

	return results
}
