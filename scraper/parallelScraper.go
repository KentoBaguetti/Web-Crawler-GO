package scraper

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/KentoBaguetti/Web-Crawler-GO/datastructures"
	"golang.org/x/net/html"
)

func ParallelCrawl(initialUrl string, numWorkers uint8, maxCrawlPages uint16, maxTokensPerPage uint16) {

	fmt.Println("Start ParallelCrawl")
	defer fmt.Println("Finished ParallelCrawl")

	jobs := make(chan string, 100)
	q := datastructures.Queue{Elements: make([]string, 0), Length: 0}
	seen := datastructures.Set{Elements: make(map[string]bool), Length: 0}
	var wg sync.WaitGroup
	var qMux sync.Mutex

	q.Enqueue(initialUrl)

	for i := uint8(0); i < numWorkers; i++ {
		id := i
		wg.Add(1)
		go func() {
			fmt.Printf("Made worker {%d}\n", id)
			defer wg.Done()
			worker(maxTokensPerPage, jobs, &q, &qMux, &seen)
		}()
	}

	// this goroutine feeds the jobs channel from the queue
	wg.Add(1)
	go func() {
		// fmt.Println(1)
		defer wg.Done()
		for {
			// fmt.Println(2)
			qMux.Lock()
			if !q.IsEmpty() && len(jobs) < 100 && seen.Length < int(maxCrawlPages) {
				// fmt.Println(3)
				job := q.Dequeue()
				qMux.Unlock()
				jobs <- job
				// fmt.Println(4)
			} else {
				// fmt.Println(5)
				shouldClose := seen.Length >= int(maxCrawlPages)
				qMux.Unlock()

				// fmt.Println(6)
				if shouldClose {
					// fmt.Println(7)
					close(jobs)
					break
				}
				// fmt.Println(8)


				time.Sleep(100 * time.Millisecond)

			}
			
		}
	}()

	wg.Wait()
	fmt.Println(seen.Elements)
	fmt.Println("seen length: ", seen.Length)
	fmt.Println("queue length: ", q.Length)
	
	}

/**
Design:
	worker arguments should be fed from a buffer, each worker should then run on its own goroutine.
	Workers feed into a separate queue, this way there is no send-receive blocking
*/
func worker(maxTokensPerPage uint16, jobs chan string, q *datastructures.Queue, qMux *sync.Mutex, seen *datastructures.Set) {

	for url := range jobs {
		fmt.Printf("Received job: %s\ns", url)
		scrapePageInParallel(url, maxTokensPerPage, q, qMux, seen)
		qMux.Lock()
		seen.Add(url)
		qMux.Unlock()
	}

}

func scrapePageInParallel(url string, maxTokensPerPage uint16, q *datastructures.Queue, qMux *sync.Mutex, seen *datastructures.Set) {

	res, err := http.Get(url)

	if err != nil {
		// fmt.Println("Error fetching data")
		return
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		// fmt.Println("Error fetching data")
	}

	// fmt.Printf("Start parsing html: %s\n", url)
	parseHtmlInsideWorker(body, maxTokensPerPage, q, qMux, seen)

}

func parseHtmlInsideWorker(content []byte, maxTokensPerPage uint16, q *datastructures.Queue, qMux *sync.Mutex, seen *datastructures.Set) {

	z := html.NewTokenizer(bytes.NewReader(content))
	var tokens uint16 = 0

	for tokens < maxTokensPerPage {

		tt := z.Next()

		if tt == html.ErrorToken {
			// fmt.Println("Error processing token")
			return
		}

		token := z.Token()

		if token.Type == html.StartTagToken && token.Data == "a" {

				ok, url := getLink(token)

				if ok && !seen.Contains(url) {
						qMux.Lock()
						q.Enqueue(url)
						qMux.Unlock()
					// fmt.Printf("Added url %s to the channel\n", url)
				}	

		}

		tokens++

	}

}

