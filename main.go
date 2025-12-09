package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/KentoBaguetti/Web-Crawler-GO/scraper"
)

func main() {

	fmt.Println("Init main")
	start := time.Now()
	defer fmt.Println("Finished main")

	r := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a link you want to search: ")
	givenUrl, err := r.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading url from terminal.")
		return
	}

	//testLink := "https://en.wikipedia.org/wiki/Japan"
	// testLink := "https://www.cs.ubc.ca/"
	// testLink := "https://en.wikipedia.org/wiki/Computer_science"
	// keywords := [2]string{"computer", "science"}

	// scraper.Crawl(testLink, 50, 2500, keywords[:])

	// initialUrl, numWorkers, numLinks to Crawl , maxTokensToSearchPerPage
	scraper.ParallelCrawl(givenUrl, 16, 100, 2500)

	t := time.Now()

	fmt.Println(t.Sub(start))

}
