package routes

import (
	controller "gdsc-core-be/controllers/request"

	"github.com/gofiber/fiber/v2"
)

func FinalProjectRoutes(app *fiber.App) {
	app.Get("/all-final-project",controller.GetAllFinalProject)
	app.Get("/final-project/:idfinalproject",controller.GetFinalProject)
	app.Post("/final-project",controller.CreateFinalProject)
	app.Patch("/final-project/:idfinalproject",controller.UpdateFinalProject)
	app.Delete("/final-project/:idfinalproject",controller.DeleteFinalProject)
}
