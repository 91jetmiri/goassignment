package server

import (
	"go-web-scraper/handlers"

	"github.com/go-chi/chi/v5"
)

func Routes() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/scrape", handlers.ScrapeHandler)
	r.Get("/data", handlers.GetDataHandler)
	return r
}
