package routes

import (
	controller "gdsc-core-be/controllers/request"

	"github.com/gofiber/fiber/v2"
)

func RoleRoutes(app *fiber.App) {
	app.Post("/role",controller.CreateRole)
	app.Patch("/role/:idrole",controller.UpdateRole)
	app.Delete("/role/:idrole",controller.DeleteRole)
}
