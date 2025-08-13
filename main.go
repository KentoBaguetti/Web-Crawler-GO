package main

import (
	"fmt"

	"github.com/KentoBaguetti/Web-Crawler-GO/scraper"
)

func main() {
	
	fmt.Println("Init")
	defer fmt.Println("Finished")

	// testLink := "https://github.com/KentoBaguetti/Web-Crawler-GO"
	testLink := "https://www.kentarobarnes.com/"

	scraper.Crawl(testLink, 1, 0)

}




