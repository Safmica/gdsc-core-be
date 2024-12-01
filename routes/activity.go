package routes

import (
	controller "gdsc-core-be/controllers/request"
	"gdsc-core-be/middleware"

	"github.com/gofiber/fiber/v2"
)

func ActivityRoutes(app *fiber.App) {
	app.Get("/all-activity",middleware.BlockRoleAuth("member"), controller.GetAllActivity)
	app.Get("/activity/:idactivity", middleware.Auth(), controller.GetActivity)
	app.Post("/activity/:idbatch", middleware.BlockRoleAuth("member"), controller.CreateActivity)
	app.Patch("/activity/:idactivity",middleware.BlockRoleAuth("member"), controller.UpdateActivity)
	app.Delete("/activity/:idactivity", middleware.BlockRoleAuth("member"), controller.DeleteActivity)
}
