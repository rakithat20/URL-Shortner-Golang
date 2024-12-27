package main

import (
	"log"

	config "url-shortner/internal/Config"
	routes "url-shortner/internal/Routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.ConnectDB()
	defer config.DB.Close()

	app := fiber.New()

	routes.RegisterRoutes(app)
	log.Fatal(app.Listen(":3000"))
}
