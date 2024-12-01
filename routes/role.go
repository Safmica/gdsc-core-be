package routes

import (
	controller "gdsc-core-be/controllers/request"
	"gdsc-core-be/middleware"

	"github.com/gofiber/fiber/v2"
)

func RoleRoutes(app *fiber.App) {
	app.Post("/role", middleware.BlockRoleAuth("member"), controller.CreateRole)
	app.Patch("/role/:idrole", middleware.BlockRoleAuth("member"), controller.UpdateRole)
	app.Delete("/role/:idrole", middleware.BlockRoleAuth("member"), controller.DeleteRole)
}
