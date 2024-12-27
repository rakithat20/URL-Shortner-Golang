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

	app := fiber.New()

	routes.RegisterRoutes(app)
	log.Fatal(app.Listen(":3000"))
}
