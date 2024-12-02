package routes

import (
	controller "gdsc-core-be/controllers/request"
	"gdsc-core-be/middleware"

	"github.com/gofiber/fiber/v2"
)

func DivisionRoutes(app *fiber.App) {
	app.Get("/all-division", middleware.BlockRoleAuth("member"), controller.GetAllDivision)
	app.Get("/division/:iddivision", middleware.Auth(), controller.GetDivision)
	app.Post("/division/:idbatch", middleware.BlockRoleAuth("member"), controller.CreateDivision)
	app.Patch("/division/:iddivision",middleware.BlockRoleAuth("member"), controller.UpdateDivision)
	app.Delete("/division/:iddivision",middleware.BlockRoleAuth("member"), controller.DeleteDivision)
}
