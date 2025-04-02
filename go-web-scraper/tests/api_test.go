package tests

import (
	"go-web-scraper/server"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestScrapeHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/scrape", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(server.Routes().ServeHTTP)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status code %v, got %v", http.StatusOK, status)
	}
}
