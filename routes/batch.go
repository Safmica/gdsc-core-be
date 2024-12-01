package routes

import (
	controller "gdsc-core-be/controllers/request"
	"gdsc-core-be/middleware"

	"github.com/gofiber/fiber/v2"
)

func BatchRoutes(app *fiber.App) {
	app.Get("/all-batch", middleware.BlockRoleAuth("member"), controller.GetAllBatch)
	app.Get("/batch/:idbatch",middleware.BlockRoleAuth("member"), controller.GetBatch)
	app.Post("/batch",middleware.BlockRoleAuth("member"), controller.CreateBatch)
	app.Delete("/batch/:idbatch",middleware.BlockRoleAuth("member"), controller.DeleteBatch)
}
