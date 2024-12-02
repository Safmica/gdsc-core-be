package routes

import (
	controller "gdsc-core-be/controllers/request"

	"github.com/gofiber/fiber/v2"
)

func TokenRoutes(app *fiber.App) {
	app.Post("/token", controller.RefreshToken)
}
