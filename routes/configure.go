package routes

import (
	controller "gdsc-core-be/controllers/request"
	"gdsc-core-be/middleware"

	"github.com/gofiber/fiber/v2"
)

func ConfigureRoutes(app *fiber.App) {
	app.Patch("/configure/:idconfigure",middleware.BlockRoleAuth("member"), controller.UpdateConfigure)
}
