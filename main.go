package main

import (
	"fmt"

	"github.com/KentoBaguetti/Web-Crawler-GO/scraper"
)

func main() {
	
	fmt.Println("Init main")
	defer fmt.Println("Finished main")

	testLink := "https://en.wikipedia.org/wiki/Japan"
	// testLink := "https://www.kentarobarnes.com/"

	scraper.Crawl(testLink, 5, 2500)

}




