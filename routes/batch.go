package routes

import (
	controller "gdsc-core-be/controllers/request"

	"github.com/gofiber/fiber/v2"
)

func BatchRoutes(app *fiber.App) {
	app.Get("/all-batch",controller.GetAllBatch)
	app.Get("/batch/:idbatch",controller.GetBatch)
	app.Post("/batch",controller.CreateBatch)
	app.Delete("/batch/:idbatch",controller.DeleteBatch)
}
