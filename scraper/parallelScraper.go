package scraper

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"sync"

	"github.com/KentoBaguetti/Web-Crawler-GO/datastructures"
	"golang.org/x/net/html"
)

func ParallelCrawl(initialUrl string, numWorkers uint8, maxCrawlPages uint16, maxTokensPerPage uint16) {

	fmt.Println("Start ParallelCrawl")
	defer fmt.Println("Finished ParallelCrawl")

	var wg sync.WaitGroup

	jobs := make(chan string, 10)
	seen := datastructures.Set{Elements: make(map[string]bool), Length: 0}
	
	jobs <- initialUrl

	
	for url := range jobs {
		wg.Add(1)
		fmt.Println("Making worker")
		go worker(&wg, url, jobs, &seen)
	}


	go func() {
		wg.Wait()
		close(jobs)
	}()

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
func worker(wg* sync.WaitGroup, url string, jobs chan string, seen* datastructures.Set) {

	defer wg.Done()



}

func scrapePageInParallel(url string, jobs chan string, seen* datastructures.Set) {

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

	// fmt.Printf("Start parsing html: %s\n", url)
	parseHtmlInsideWorker(body, jobs, seen)

}

func parseHtmlInsideWorker(content []byte, jobs chan string, seen* datastructures.Set) {

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
					jobs <- url
					// fmt.Printf("Added url %s to the channel\n", url)
				}
			}

		}

		tokens++

	}

}

