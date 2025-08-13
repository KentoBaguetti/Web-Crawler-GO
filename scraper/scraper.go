package scraper

import (
	"fmt"
	"io"
	"net/http"
)

func SimpleScrape(url string) {

	resp, err := http.Get(url)
	
	if err != nil {
		fmt.Println("Error fetching data")
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Error fetching data")
	}

	fmt.Println(string(body))

}