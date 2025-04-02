package database

import (
	"database/sql"
	"fmt"
	"log"

	"go-web-scraper/config"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("postgres", config.GetDBConnStr())
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Database is not reachable:", err)
	}

	fmt.Println("Connected to the database!")

	createTableSQL := `
	CREATE TABLE IF NOT EXISTS scraped_data (
		id SERIAL PRIMARY KEY,
		title TEXT,
		link TEXT UNIQUE -- Added UNIQUE constraint to avoid duplicate links
	);` // Using raw string literal for multi-line SQL

	// Execute the SQL statement
	_, err = DB.Exec(createTableSQL)
	if err != nil {
		DB.Close() // Close the connection if table creation fails
		log.Fatalf("Failed to create table 'scraped_data': %v", err)
	}

	fmt.Println("Table 'scraped_data' checked/created successfully.")
}

func CloseDB() {
	DB.Close()
}
