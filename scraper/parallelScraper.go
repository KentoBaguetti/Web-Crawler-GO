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
	jobs := make(chan string, 2)
	done := make(chan bool)

	
	for i := range numWorkers {
		go worker(i, &q, &seen, maxCrawlPages, done)
		fmt.Printf("Worker %d created\n", i)
	}
	
	jobs <- initialUrl

	<-done

	fmt.Println(seen.Elements)

}

/**
	The reason why it is currently not working is because when the jobs channel is being added to, once it reaches its capactity
	it blocks the goroutine from moving on.
	Need to find a way to prevent this, such as store urls in a queue first, then populate the channel once it channel only has x
	urls left in it.
	Or make a goroutine for every single url (this will probably be too expensive to do)
	Or 
*/


/**
Design:
	worker arguments should be fed from a buffer, each worker should then run on its own goroutine
*/
func worker(id uint8, q* datastructures.Queue, seen* datastructures.Set, maxCrawlPages uint16, done chan bool) {

	url := q.Dequeue()

	fmt.Printf("Worker {%d} received url: %s\n", id, url)

	scrapePageInParallel(url, q, seen)

	for q.Length > 0 {
		newUrl := q.Dequeue()
		scrapePageInParallel(newUrl, q, seen)
	}

	if q.Length == 0 || seen.Length >= int(maxCrawlPages) {
		done <- true
	}

	

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

	fmt.Printf("Start parsing html: %s\n", url)
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
					fmt.Printf("Added url %s to the channel\n", url)
				}
			}

		}

		tokens++

	}

}

