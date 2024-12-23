package main

import (
	"fmt"
	"log"

	models "url-shortner/internal/Models"
	"url-shortner/internal/config"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.ConnectDB()
	defer config.DB.Close()

	// Initialize a new Fiber app
	app := fiber.New()

	// Define a route for the GET method on the root path '/'
	app.Post("/", func(c *fiber.Ctx) error {
		longUrl := c.FormValue("longUrl")
		shortUrl := c.FormValue("shortUrl")
		url, err := models.AddUrl(longUrl, shortUrl)
		if err != nil {
			fmt.Println(err)
			return c.SendStatus(400)
		}
		fmt.Println(url)
		return c.SendString("Hello, World 👋!")
	})
	// Start the server on port 3000
	log.Fatal(app.Listen(":3000"))
}
