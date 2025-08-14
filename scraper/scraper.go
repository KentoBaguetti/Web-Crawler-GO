package scraper

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"

	"github.com/KentoBaguetti/Web-Crawler-GO/datastructures"
	"golang.org/x/net/html"
)

func Crawl(initialUrl string, maxCrawlPages uint16, maxTokensPerPage uint16, keywords []string) {

	q := datastructures.Queue{Elements: make([]string, 0), Length: 0}
	seen := datastructures.Set{Elements: make(map[string]bool), Length: 0}
	ch1 := make(chan []byte) // chnage this to a buffered channel to allow for workers

	parseComplete := make(chan bool)
	var mux sync.Mutex

	urlLinkFreq := make(map[string]int)

	q.Enqueue(initialUrl)

	// bfs loop
	for q.Size() > 0 && seen.Size() < int(maxCrawlPages) {
		
		currUrl := q.Dequeue()
		
		if seen.Contains(currUrl) {
			continue
		}

		seen.Add(currUrl)

		// use a goroutine to scrape the first page
		go ScrapeOnePage(currUrl, ch1, &q)

		x := <- ch1

		// once a page is scraped, start parsing it while a new page is being scraped
		go func (url string, content []byte) {
			mux.Lock()
			ParseHTML(url, content, maxTokensPerPage, &q, &urlLinkFreq, keywords)
			mux.Unlock()
			parseComplete <- true
		}(currUrl, x)

		<-parseComplete

	}

	defer fmt.Println("Pages Scraped:", seen.Size())
	defer fmt.Println(urlLinkFreq)
	fmt.Println("Finished Crawling")

}

func ScrapeOnePage(url string, c chan []byte, q* datastructures.Queue) {

	resp, err := http.Get(url)
	
	if err != nil {
		fmt.Println("Error fetching data")
	}

	defer resp.Body.Close()

	// read the body of the html
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Error fetching data")
	}

	c <- body

}

func ParseHTML(url string, content []byte, maxTokens uint16, q* datastructures.Queue, m* map[string]int, keywords []string) {

	// tokenize the html pge based on bytes
	z := html.NewTokenizer(bytes.NewReader(content))
	var tokens uint16 = 0
	var linksNo int = 0

	// only search a set number of tokens per html page
	for tokens < maxTokens {

		tt := z.Next()

		if tt == html.ErrorToken {
			fmt.Println("Error processing token")
			return
		}

		token := z.Token()

		// only queue links (anchor tags)
		if token.Type == html.StartTagToken {

			if token.Data == "a" {

				ok, url := getLink(token)

				if ok && HasKeyWords(url, keywords) {
					q.Enqueue(url)
					linksNo++
				}

			} 

		} 

		tokens++

	}

	(*m)[url] = linksNo

}

func getLink(token html.Token) (ok bool, url string) {

	for _, t := range token.Attr {

		if t.Key == "href" {

			// if the link value in the anchor tag isnt valid
			if len(t.Val) == 0 || !strings.HasPrefix(t.Val, "http") {
				ok = false
				url = t.Val
				return ok, url
			}

			ok = true
			url = t.Val
					
		}

	}

	return ok, url

}

// return true if the url contains keywords from the keywords string
// should probably modify this function to return an integer based on the number of matching keywords
// eg more keywords == beter score
func HasKeyWords (url string, keywords []string) bool {

	hasKeywords := false

	for _, word := range keywords {
		if strings.Contains(url, word) {
			hasKeywords = true
			break
		}
	}

	return hasKeywords

}