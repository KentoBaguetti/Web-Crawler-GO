package scraper

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/KentoBaguetti/Web-Crawler-GO/datastructures"
	"golang.org/x/net/html"
)

func ParallelCrawl(initialUrl string, numWorkers uint8, maxCrawlPages uint16, maxTokensPerPage uint16, keywords *[]string) {

	fmt.Println("Start ParallelCrawl")
	defer fmt.Println("Finished ParallelCrawl")

	jobs := make(chan string, 100)
	pq := datastructures.CreatePriorityQueue(false) // max-heap
	seen := datastructures.Set{Elements: make(map[string]bool), Length: 0}
	var wg sync.WaitGroup
	var qMux sync.Mutex

	searchedUrls := datastructures.Queue{Elements: make([]string, 0), Length: 0}

	inFlight := 0 // number of urls currently being processed, protected by qMux, like a semaphore
	initialUrl = strings.TrimSpace(initialUrl)
	if initialUrl == "" {
		fmt.Println("Empty URL provided.")
		return
	}
	if maxCrawlPages == 0 {
		fmt.Println("maxCrawlPages is 0. Nothing to crawl.")
		return
	}

	seen.Add(initialUrl)
	searchedUrls.Enqueue(initialUrl)
	pq.Append(initialUrl, CalculateKeywordScore(initialUrl, *keywords))

	// create workers
	for i := uint8(0); i < numWorkers; i++ {
		id := i
		wg.Add(1)
		go func() {
			fmt.Printf("Made worker {%d}\n", id)
			defer wg.Done()
			worker(maxTokensPerPage, maxCrawlPages, jobs, pq, &qMux, &seen, &inFlight, &searchedUrls, keywords)
		}()
	}

	// this goroutine feeds the jobs channel from the queue
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			qMux.Lock()
			// fmt.Println(1)
			if pq.Length > 0 && len(jobs) < cap(jobs) { // check if the channel has space so it doesn't block immediately
				scoreValueObj, err := pq.Pop()
				if err != nil {
					// fmt.Println(2)
					fmt.Println("Error with popping from pq")
					qMux.Unlock()
					close(jobs)
					break
				}
				// fmt.Println(3)
				inFlight++ // incrament sempahore thing
				qMux.Unlock()
				jobs <- scoreValueObj.Value // feed the job queue/channel
			} else {
				// fmt.Println(4)
				shouldClose := pq.Length == 0 && inFlight == 0 // break condition
				qMux.Unlock()

				// fmt.Println(5)
				if shouldClose {
					// fmt.Println(6)
					close(jobs)
					break
				}
				// fmt.Println(7)

				time.Sleep(100 * time.Millisecond)

			}

		}
	}()

	wg.Wait()
	qMux.Lock()

	for i, url := range searchedUrls.Elements {
		fmt.Printf("\n%d: %s\n", i+1, url)
	}

	fmt.Println("seen length: ", seen.Length)
	fmt.Println("pq length: ", pq.Length)
	qMux.Unlock()

}

/*
*
Design:

	worker arguments should be fed from a buffer, each worker should then run on its own goroutine.
	Workers feed into a separate queue, this way there is no send-receive blocking
*/
func worker(maxTokensPerPage uint16, maxCrawlPages uint16, jobs chan string, pq *datastructures.PriorityQueue, qMux *sync.Mutex, seen *datastructures.Set, inFlight *int, searchedUrls *datastructures.Queue, keywords *[]string) {

	// worker constanly feeds off the jobs channel until its empty
	for url := range jobs {
		fmt.Printf("Received job: %s\n", url)
		scrapePageInParallel(url, maxTokensPerPage, maxCrawlPages, pq, qMux, seen, searchedUrls, keywords)
		qMux.Lock()
		*inFlight-- // finsished parsing a page so decrement the semaphore
		qMux.Unlock()
	}

}

// can and should remove this function later
func scrapePageInParallel(url string, maxTokensPerPage uint16, maxCrawlPages uint16, pq *datastructures.PriorityQueue, qMux *sync.Mutex, seen *datastructures.Set, searchedUrls *datastructures.Queue, keywords *[]string) {

	res, err := http.Get(url)

	if err != nil {
		fmt.Println("Error fetching data:", err)
		return
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return
	}

	parseHtmlInsideWorker(body, maxTokensPerPage, maxCrawlPages, pq, qMux, seen, searchedUrls, keywords)

}

func parseHtmlInsideWorker(content []byte, maxTokensPerPage uint16, maxCrawlPages uint16, pq *datastructures.PriorityQueue, qMux *sync.Mutex, seen *datastructures.Set, searchedUrls *datastructures.Queue, keywords *[]string) {

	z := html.NewTokenizer(bytes.NewReader(content))
	var tokens uint16 = 0

	// check all tokens
	for tokens < maxTokensPerPage {

		tt := z.Next()

		if tt == html.ErrorToken {
			// fmt.Println("Error processing token")
			return
		}

		token := z.Token()

		// if the token is an anchor tag, add the url to search
		if token.Type == html.StartTagToken && token.Data == "a" {

			ok, url := getLink(token)

			if ok {
				qMux.Lock()
				if seen.Length < int(maxCrawlPages) {
					if _, exists := seen.Elements[url]; !exists {
						searchedUrls.Enqueue(url)
						seen.Add(url)
						pq.Append(url, CalculateKeywordScore(url, *keywords))
					}
				}
				qMux.Unlock()
			}

		}

		tokens++

	}

}
