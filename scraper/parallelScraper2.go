package scraper

import (
	"fmt"
	"strings"
	"sync"

	"github.com/KentoBaguetti/Web-Crawler-GO/datastructures"
)

// returns an array of all text within specific closing tags
// example: string
func KeywordPriorityCrawler(initialUrl string, keywords []string, numWorkers, maxCrawlPages, maxTokensPerPage int) (urls []string) {

	fmt.Println("Start Crawl")
	defer fmt.Println("Finished Crawl")

	jobs := make(chan string, 500)
	pq := datastructures.CreatePriorityQueue(false)

	fmt.Printf("jobs: %v\n", jobs)
	fmt.Printf("pq: %v\n", pq)

	return

}

func CalculateKeywordScore(s string, keywords []string) (score int) {

	s = strings.ToLower(s)

	for _, word := range keywords {

		word = strings.ToLower(word)
		if strings.Contains(s, word) {
			score++
		}
	}

	return

}

func parseAndScoreHtml(content []byte, maxTokensPerPage int, keywords []string, pq *datastructures.PriorityQueue, qMux *sync.Mutex, seen *datastructures.Set) {

	// z := html.NewTokenizer(bytes.NewReader(content))
	currTokens := 0
	fmt.Printf("currTokens: %v\n", currTokens)

	// for tokens <= maxTokensPerPage {

	// 	tt := z.Next()

	// 	if tt == html.ErrorToken {
	// 		return
	// 	}

	// 	token = z.Token()

	// 	if token.Type == html.StartTagToken && token.Data == "a" {

	// 		ok, url := getLink(token)

	// 		if ok && !seen.Contains(url) {
	// 			qMux.Lock()

	// 			urlScore = calculateKeywordScore(url, keywords)

	// 			qMux.Unlock()
	// 		}

	// 	}

	// }

}
