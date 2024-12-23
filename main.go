package main

import (
	"log"

	routes "url-shortner/internal/Routes"
	"url-shortner/internal/config"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.ConnectDB()
	defer config.DB.Close()

	// Initialize a new Fiber app
	app := fiber.New()

	routes.RegisterRoutes(app)
	// Start the server on port 3000
	log.Fatal(app.Listen(":3000"))
}
