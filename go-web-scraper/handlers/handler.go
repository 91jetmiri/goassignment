package handlers

import (
	"encoding/json"
	"fmt"
	"go-web-scraper/database"
	"go-web-scraper/models"
	"go-web-scraper/scraper"
	"net/http"
)

// ScrapeHandler triggers the web scraper
func ScrapeHandler(w http.ResponseWriter, r *http.Request) {
	url := "https://example.com"
	go scraper.Scrape(url) // Runs scraper in a goroutine
	fmt.Fprintln(w, "Scraping started!")
}

// GetDataHandler retrieves scraped data from the database
func GetDataHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query("SELECT id, title, link FROM scraped_data")
	if err != nil {
		http.Error(w, "Database query failed", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var results []models.ScrapedData
	for rows.Next() {
		var data models.ScrapedData
		if err := rows.Scan(&data.ID, &data.Title, &data.Link); err != nil {
			http.Error(w, "Data scan failed", http.StatusInternalServerError)
			return
		}
		results = append(results, data)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}
