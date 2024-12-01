package routes

import (
	controller "gdsc-core-be/controllers/request"
	"gdsc-core-be/middleware"

	"github.com/gofiber/fiber/v2"
)

func TokenRoutes(app *fiber.App) {
	app.Post("/token", middleware.Auth(), controller.RefreshToken)
}
