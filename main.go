package main

import (
	"fmt"
	"log"
	"time"

	config "url-shortner/internal/Config"
	models "url-shortner/internal/Models"
	routes "url-shortner/internal/Routes"

	"github.com/gofiber/fiber/v2"
)

func main() {

	config.ConnectDB()
	defer config.DB.Close()

	app := fiber.New()
	go func() {
		for {
			// Call the cleanup function
			err := models.CleanUP()
			if err != nil {
				log.Printf("Cleanup failed: %v", err)
			} else {
				fmt.Println("Cleanup successful")
			}
			// Wait for 24 hours before running it again
			time.Sleep(24 * time.Hour)
		}
	}()
	routes.RegisterRoutes(app)
	log.Fatal(app.Listen(":3000"))
}
