package models

import (
	"log"
	"time"
	"url-shortner/internal/config"
)

type URL struct {
	ID        uint      `json:"ID"`
	ShortURL  string    `json:"ShortURL"`
	LongURL   string    `json:"LongURL"`
	CreatedAt time.Time `json:"CreatedAt"`
}

func AddUrl(longUrl string, shortUrl string) (URL, error) {
	stmt := `insert into urls (short_url, long_url) values($1, $2) RETURNING id, short_url, long_url, created_at`
	var url URL

	err := config.DB.QueryRow(stmt, shortUrl, longUrl).Scan(&url.ID, &url.ShortURL, &url.LongURL, &url.CreatedAt)
	if err != nil {
		log.Printf("Error creating URL: %v", err)
		return url, err
	}
	return url, nil
}

func GetUrl(shortUrl string) (URL, error) {
	stmt := `SELECT * FROM urls WHERE shorturl = $1`
	var url URL

	err := config.DB.QueryRow(stmt, shortUrl).Scan(&url.ID, &url.ShortURL, &url.LongURL, &url.CreatedAt)
	if err != nil {
		log.Printf("Error getting URL: %v", err)
		return url, err
	}
	return url, nil
}
