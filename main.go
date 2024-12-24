package main

import (
	"log"
	"time"

	routes "url-shortner/internal/Routes"
	"url-shortner/internal/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
)

func main() {
	config.ConnectDB()
	defer config.DB.Close()

	// Initialize a new Fiber app
	app := fiber.New()
	app.Use(cache.New(cache.Config{
		Expiration: 30 * time.Minute, // 30 minutes
	}))

	routes.RegisterRoutes(app)
	// Start the server on port 3000
	log.Fatal(app.Listen(":3000"))
}
