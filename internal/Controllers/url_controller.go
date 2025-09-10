package Controllers

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"net/http"
	config "url-shortner/internal/Config"
	models "url-shortner/internal/Models"

	"github.com/gofiber/fiber/v2"
	"github.com/patrickmn/go-cache"
)

func AddUrlController(c *fiber.Ctx) error {
	var longUrl = c.FormValue("longUrl")
	if longUrl != "" {
		url, err := models.AddUrl(longUrl, genShortUrl())
		if err != nil {
			fmt.Println(err)
			return c.Status(http.StatusBadRequest).SendString("Failed to add URL")
		}
		return c.JSON(url)
	}

	return c.Status(http.StatusBadRequest).SendString("empty url")

}
func GetLongUrlController(c *fiber.Ctx) error {
	shortUrl := c.Params("shortURL")
	if shortUrl != "" {
		url, err := models.GetUrl(shortUrl)

		if err != nil {
			fmt.Println(err)
			return c.Status(http.StatusBadRequest).SendString("Failed to get URL")
		}
		config.CACHE.Set(url.ShortURL, &url, cache.DefaultExpiration)
		return c.Redirect(url.LongURL)
	}
	return c.Status(http.StatusBadRequest).SendString("empty url")

}
func genShortUrl() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const urlLength = 6

	result := make([]byte, urlLength)
	charsetLen := big.NewInt(int64(len(charset)))

	for i := range result {
		randomIndex, _ := rand.Int(rand.Reader, charsetLen)

		result[i] = charset[randomIndex.Int64()]
	}

	return string(result)
}
func GetUrlInfo(c *fiber.Ctx) error {

	shortUrl := c.Params("shortURL")
	if shortUrl != "" {
		url, err := models.GetUrlInfo(shortUrl)

		if err != nil {
			return c.Status(http.StatusBadRequest).SendString("Failed to get URL")
		}
		config.CACHE.Set(url.ShortURL, &url, cache.DefaultExpiration)
		return c.JSON(url)
	}
	return c.Status(http.StatusBadRequest).SendString("ShortURL was empty")

}
