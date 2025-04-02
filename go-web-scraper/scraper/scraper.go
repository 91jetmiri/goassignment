package scraper

import (
	"fmt"
	"go-web-scraper/database"
	"go-web-scraper/models"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

// Scrape function to fetch and parse HTML
func Scrape(url string) {
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

	var titles []models.ScrapedData
	var scrape func(*html.Node)
	scrape = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					title := strings.TrimSpace(n.FirstChild.Data)
					if title != "" {
						titles = append(titles, models.ScrapedData{Title: title, Link: attr.Val})
					}
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			scrape(c)
		}
	}
	scrape(doc)

	// Store results in DB
	for _, item := range titles {
		_, err := database.DB.Exec("INSERT INTO scraped_data (title, link) VALUES ($1, $2)", item.Title, item.Link)
		if err != nil {
			fmt.Println("Error inserting data:", err)
		}
	}
}
