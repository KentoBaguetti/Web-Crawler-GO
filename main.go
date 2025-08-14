package main

import (
	"fmt"
	"time"

	"github.com/KentoBaguetti/Web-Crawler-GO/scraper"
)

func main() {
	
	fmt.Println("Init main")
	start := time.Now()

	defer fmt.Println("Finished main")

	//testLink := "https://en.wikipedia.org/wiki/Japan"
	testLink := "https://en.wikipedia.org/wiki/Computer_science"
	// testLink := "https://www.cs.ubc.ca/"
	keywords := [2]string{"computer", "science"}

	scraper.Crawl(testLink, 100, 2500, keywords[:])

	t := time.Now()

	fmt.Println(t.Sub(start))

}




