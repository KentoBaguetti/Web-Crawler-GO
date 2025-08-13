package scraper

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/KentoBaguetti/Web-Crawler-GO/datastructures"
	"golang.org/x/net/html"
)

func Crawl(initialUrl string, maxCrawlPages uint16, maxTokensPerPage uint16) {

	q := datastructures.Queue{Elements: make([]string, 0), Length: 0}
	seen := datastructures.Set{Elements: make(map[string]bool), Length: 0}
	ch1 := make(chan []byte)

	q.Enqueue(initialUrl)

	for q.Size() > 0 && seen.Size() < int(maxCrawlPages) {
		
		currUrl := q.Dequeue()
		
		if seen.Contains(currUrl) {
			continue
		}

		seen.Add(currUrl)

		go ScrapeOnePage(currUrl, ch1, &q)

		x := <- ch1

		// fmt.Println(string(x))
		ParseHTML(currUrl, x, maxTokensPerPage, &q)


	}

	defer fmt.Println(seen.Size())
	fmt.Println("Finished Crawling")

}

func ScrapeOnePage(url string, c chan []byte, q* datastructures.Queue) {

	resp, err := http.Get(url)
	
	if err != nil {
		fmt.Println("Error fetching data")
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Error fetching data")
	}

	c <- body

}

func ParseHTML(url string, content []byte, maxTokens uint16, q* datastructures.Queue) {

	z := html.NewTokenizer(bytes.NewReader(content))
	var tokens uint16 = 0

	for tokens < maxTokens {

		tt := z.Next()

		if tt == html.ErrorToken {
			fmt.Println("Error processing token")
			return
		}

		token := z.Token()

		if token.Type == html.StartTagToken {

			if token.Data == "a" {

				ok, url := getLink(token)

				if ok {
					fmt.Println(ok, url)
					q.Enqueue(url)
				}

				// fmt.Println(ok, tokens, url)

			} 

		} 

		tokens++

	}

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