package models

type ScrapedData struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Link  string `json:"link"`
}
