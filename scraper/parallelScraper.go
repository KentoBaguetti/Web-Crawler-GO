package scraper

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/KentoBaguetti/Web-Crawler-GO/datastructures"
	"golang.org/x/net/html"
)

func ParallelCrawl(initialUrl string, numWorkers uint8, maxCrawlPages uint16, maxTokensPerPage uint16) {

	fmt.Println("Start ParallelCrawl")
	defer fmt.Println("Finished ParallelCrawl")

	q := datastructures.Queue{Elements: make([]string, 0), Length: 0}
	seen := datastructures.Set{Elements: make(map[string]bool), Length: 0}

	q.Enqueue(initialUrl)

	for q.Length > 0 && seen.Length < int(maxCrawlPages) {

		

	}

}


/**
Design:
	worker arguments should be fed from a buffer, each worker should then run on its own goroutine
*/
func worker(url string, q* datastructures.Queue, seen* datastructures.Set) {


	scrapePageInParallel(url, q, seen)


}

func scrapePageInParallel(url string, q* datastructures.Queue, seen* datastructures.Set) {

	res, err := http.Get(url)

	if err != nil {
		fmt.Println("Error fetching data")
		return
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		fmt.Println("Error fetching data")
	}

	parseHtmlInsideWorker(url, body, q, seen)

}

func parseHtmlInsideWorker(url string, content []byte, q* datastructures.Queue, seen* datastructures.Set) {

	z := html.NewTokenizer(bytes.NewReader(content))
	var tokens uint16 = 0

	for tokens < 2500 {

		tt := z.Next()

		if tt == html.ErrorToken {
			fmt.Println("Error processing token")
			return
		}

		token := z.Token()

		if token.Type == html.StartTagToken {

			if token.Data == "a" {
				ok, url := getLink(token)

				if ok && !seen.Contains((url)) {
					q.Enqueue(url)
				}
			}

		}

		tokens++

	}

}

