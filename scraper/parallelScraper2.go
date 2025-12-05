package scraper

import (
	"fmt"
	"strings"
)

// returns an array of all text within specific closing tags
// example: string
func KeywordPriorityCrawler(initialUrl string, keywords []string) (urls []string) {

	fmt.Println("Start Crawl")
	defer fmt.Println("Finished Crawl")

	return

}

func calculateKeywordScore(s string, keywords []string) (score int) {

	for _, word := range keywords {
		if strings.Contains(s, word) {
			score++
		}
	}

	return

}
