package scraper

import (
	"fmt"
	"go-web-scraper/database"
	"go-web-scraper/models"
	"net/http"
	"strings"
	"sync"

	"golang.org/x/net/html"
)

// ScrapeMultiple scrapes multiple URLs concurrently
func ScrapeMultiple(urls []string) {
	var wg sync.WaitGroup
	dataChannel := make(chan models.ScrapedData, len(urls)*10) // Buffered channel

	// Launch a goroutine for each URL
	for _, url := range urls {
		wg.Add(1)
		go func(u string) {
			defer wg.Done()
			scrape(u, dataChannel)
		}(url)
	}

	// Close the channel once all scraping is done
	go func() {
		wg.Wait()
		close(dataChannel)
	}()

	// Store data from channel into the database
	for data := range dataChannel {
		_, err := database.DB.Exec("INSERT INTO scraped_data (title, link) VALUES ($1, $2)", data.Title, data.Link)
		if err != nil {
			fmt.Println("Error inserting data:", err)
		}
	}
}

// scrape extracts data from a single page and sends it to the channel
func scrape(url string, dataChannel chan<- models.ScrapedData) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching the page:", err)
		return
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Println("Error parsing the page:", err)
		return
	}

	var extractLinks func(*html.Node)
	extractLinks = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			var title string
			var link string

			for _, attr := range n.Attr {
				if attr.Key == "href" {
					link = attr.Val
				}
			}

			if n.FirstChild != nil {
				title = strings.TrimSpace(n.FirstChild.Data)
			}

			if title != "" && link != "" {
				dataChannel <- models.ScrapedData{Title: title, Link: link}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			extractLinks(c)
		}
	}

	extractLinks(doc)
}
