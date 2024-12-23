package controllers

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
	models "url-shortner/internal/Models"

	"github.com/gofiber/fiber/v2"
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
func genShortUrl() string {
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789")

	b := make([]rune, 6)
	for i := range b {
		b[i] = letterRunes[rng.Intn(len(letterRunes))]
	}
	return string(b)

}
