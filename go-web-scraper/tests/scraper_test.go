package tests

import (
	"go-web-scraper/database"
	"go-web-scraper/scraper"
	"testing"
)

func TestScraper(t *testing.T) {

	database.InitDB()
	defer database.CloseDB()

	url := "https://example.com"
	scraper.ScrapeMultiple([]string{url})

	//  check the database for inserted data
	var count int
	err := database.DB.QueryRow("SELECT COUNT(*) FROM scraped_data").Scan(&count)
	if err != nil {
		t.Fatalf("Error querying database: %v", err)
	}

	// Check if at least one row was inserted
	if count == 0 {
		t.Errorf("Expected at least one row to be inserted, but found none")
	}

	// Cleanup test data
	_, err = database.DB.Exec("DELETE FROM scraped_data")
	if err != nil {
		t.Fatalf("Error cleaning up database: %v", err)
	}
}
