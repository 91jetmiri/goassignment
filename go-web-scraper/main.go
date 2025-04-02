package main

import (
	"fmt"
	"go-web-scraper/database"
	"go-web-scraper/server"
	"go-web-scraper/utils"
	"net/http"
)

func main() {
	fmt.Println("Starting Go Web Scraper...")

	// Initialize database
	database.InitDB()
	defer database.CloseDB()

	// Initialize logger
	utils.InitLogger()

	// Setup routes
	r := server.Routes()

	// Start web server
	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
