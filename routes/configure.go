package routes

import (
	controller "gdsc-core-be/controllers/request"

	"github.com/gofiber/fiber/v2"
)

func ConfigureRoutes(app *fiber.App) {
	app.Patch("/configure/:idconfigure",controller.UpdateConfigure)
}
