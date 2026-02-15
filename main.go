package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/KentoBaguetti/Web-Crawler-GO/scraper"
)

func main() {

	fmt.Println("Init main")
	defer fmt.Println("Finished main")

	r := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a link you want to search: ")
	givenUrl, err := r.ReadString('\n')
	givenUrl = strings.TrimSpace(givenUrl)
	if err != nil {
		fmt.Println("Error reading url from terminal.")
		return
	}
	fmt.Println(givenUrl)

	//testLink := "https://en.wikipedia.org/wiki/Japan"
	// testLink := "https://www.cs.ubc.ca/"
	// testLink := "https://en.wikipedia.org/wiki/Computer_science"
	// keywords := [2]string{"computer", "science"}

	// scraper.Crawl(testLink, 50, 2500, keywords[:])

	// initialUrl, numWorkers, numLinks to Crawl , maxTokensToSearchPerPage
	start := time.Now()
	keywords := []string{"tv"}
	scraper.ParallelCrawl(givenUrl, 16, 100, 1000, &keywords)

	// pq := datastructures.CreatePriorityQueue(false)
	// pq.Append("Kentaro", 55)
	// pq.Append("Barnes", 10)
	// for i, _ := range pq.Elements {
	// 	fmt.Printf("Size of PQ: %d\n", pq.Size())
	// 	item, err := pq.Pop()
	// 	if err != nil {
	// 		continue
	// 	}
	// 	fmt.Printf("%d, %s\n", i, item.Value)
	// }

	t := time.Now()

	fmt.Println(t.Sub(start))

}
