package routes

import (
	controller "gdsc-core-be/controllers/request"

	"github.com/gofiber/fiber/v2"
)

func DivisionRoutes(app *fiber.App) {
	app.Get("/all-division",controller.GetAllDivision)
	app.Get("/division/:iddivision",controller.GetDivision)
	app.Post("/division",controller.CreateDivision)
	app.Patch("/division/:iddivision",controller.UpdateDivision)
	app.Delete("/division/:iddivision",controller.DeleteDivision)
}
