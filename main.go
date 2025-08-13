package main

import (
	"fmt"

	"github.com/KentoBaguetti/Web-Crawler-GO/scraper"
)

func main() {
	
	fmt.Println("Init")
	defer fmt.Println("Finished")

	testLink := "https://en.wikipedia.org/wiki/Japan"
	// testLink := "https://www.kentarobarnes.com/"

	scraper.Crawl(testLink, 500, 50)

}




