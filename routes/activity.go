package routes

import (
	controller "gdsc-core-be/controllers/request"

	"github.com/gofiber/fiber/v2"
)

func ActivityRoutes(app *fiber.App) {
	app.Get("/all-activity",controller.GetAllActivity)
	app.Get("/activity/:idactivity",controller.GetActivity)
	app.Post("/activity/:idbatch",controller.CreateActivity)
	app.Patch("/activity/:idactivity",controller.UpdateActivity)
	app.Delete("/activity/:idactivity",controller.DeleteActivity)
}
