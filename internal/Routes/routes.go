package routes

import (
	"net/http"
	controllers "url-shortner/internal/Controllers"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {
	app.Get("/favicon.ico", func(c *fiber.Ctx) error { return c.Status(http.StatusNoContent).Send(nil) })
	app.Post("/api/urls", controllers.AddUrlController)
	app.Get("/:shortURL", controllers.GetLongUrlController)
	app.Get("/api/urls/:shortURL", controllers.GetUrlInfo)

}
