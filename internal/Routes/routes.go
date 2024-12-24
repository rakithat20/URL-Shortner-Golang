package routes

import (
	controllers "url-shortner/internal/Controllers"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {
	app.Post("/url", controllers.AddUrlController)
	app.Get("/:shortURL", controllers.GetLongUrlController)
}
