package tests

import (
	"go-web-scraper/database"
	"testing"

	_ "github.com/lib/pq"
)

func TestDBConnection(t *testing.T) {
	database.InitDB()
	defer database.CloseDB()

	err := database.DB.Ping()
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
}
