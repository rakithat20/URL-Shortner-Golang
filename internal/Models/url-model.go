package models

import (
	"log"
	"time"
	"url-shortner/internal/config"
)

type URL struct {
	ID          uint      `json:"ID"`
	ShortURL    string    `json:"ShortURL"`
	LongURL     string    `json:"LongURL"`
	CreatedAt   time.Time `json:"CreatedAt"`
	TotalClicks uint      `json:"TotalClicks"`
}

func AddUrl(longUrl string, shortUrl string) (URL, error) {
	stmt := `insert into urls (short_url, long_url) values($1, $2) RETURNING id, short_url, long_url, created_at,total_clicks`
	var url URL

	err := config.DB.QueryRow(stmt, shortUrl, longUrl).Scan(&url.ID, &url.ShortURL, &url.LongURL, &url.CreatedAt, &url.TotalClicks)
	if err != nil {
		log.Printf("Error creating URL: %v", err)
		return url, err
	}
	return url, nil
}
func GetUrlInfo(shortUrl string) (URL, error) {
	var url URL
	stmt := "SELECT * FROM urls WHERE short_url = $1"
	err := config.DB.QueryRow(stmt, shortUrl).Scan(&url.ID, &url.ShortURL, &url.LongURL, &url.CreatedAt, &url.TotalClicks)
	if err != nil {
		return url, err
	}
	return url, nil

}
func GetUrl(shortUrl string) (URL, error) {

	stmt := `UPDATE urls SET total_clicks = total_clicks + 1 WHERE short_url = $1 RETURNING id, short_url, long_url, created_at, total_clicks`
	var url URL

	err := config.DB.QueryRow(stmt, shortUrl).Scan(&url.ID, &url.ShortURL, &url.LongURL, &url.CreatedAt, &url.TotalClicks)
	if err != nil {
		log.Printf("Error getting URL: %v", err)
		return url, err
	}
	return url, nil
}
