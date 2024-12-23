package controllers

import (
	"fmt"
	models "url-shortner/internal/Models"

	"github.com/gofiber/fiber/v2"
)

func urlController(c *fiber.Ctx) error {
	longUrl := c.FormValue("longUrl")
	shortUrl := c.FormValue("shortUrl")
	url, err := models.AddUrl(longUrl, shortUrl)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return c.JSON(url)

}
