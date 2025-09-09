package main

import (
	"fmt"
	"log"
	"time"
	Models "url-shortner/internal/Models"

	config "url-shortner/internal/Config"
	routes "url-shortner/internal/Routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	l := log.New(log.Writer(), "URL Shortner: ", log.LstdFlags|log.Lshortfile)

	config.ConnectDB(l)
	defer config.DB.Close()

	app := fiber.New()
	app.Use(logger.New(logger.Config{
		Format:     "[${time}] ${status} - ${method} ${path} (${latency})\n",
		TimeFormat: "02-Jan-2006 15:04:05",
		TimeZone:   "Asia/Colombo",
	}))
	routes.RegisterRoutes(app)

	go func() {
		for {
			// Call the cleanup function
			err := Models.CleanUP()
			if err != nil {
				log.Printf("Cleanup failed: %v", err)
			} else {
				fmt.Println("Cleanup successful")
			}
			// Wait for 24 hours before running it again
			time.Sleep(24 * time.Hour)
		}
	}()
	log.Println("Server starting on port 3000")
	if err := app.Listen(":3000"); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
