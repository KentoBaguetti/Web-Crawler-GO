package main

import (
	"fmt"

	"github.com/KentoBaguetti/Web-Crawler-GO/scraper"
)

func main() {
	
	fmt.Println("Init main")
	defer fmt.Println("Finished main")

	//testLink := "https://en.wikipedia.org/wiki/Japan"
	testLink := "https://en.wikipedia.org/wiki/Computer_science"
	// testLink := "https://www.cs.ubc.ca/"
	keywords := [2]string{"computer", "science"}

	scraper.Crawl(testLink, 5, 2500, keywords[:])

}




