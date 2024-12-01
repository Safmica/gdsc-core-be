package routes

import (
	controller "gdsc-core-be/controllers/request"
	"gdsc-core-be/middleware"

	"github.com/gofiber/fiber/v2"
)

func FinalProjectRoutes(app *fiber.App) {
	app.Get("/all-final-project", middleware.BlockRoleAuth("member"), controller.GetAllFinalProject)
	app.Get("/final-project/:idfinalproject", middleware.Auth(), controller.GetFinalProject)
	app.Post("/final-project", middleware.Auth(),  controller.CreateFinalProject)
	app.Patch("/final-project/:idfinalproject", middleware.Auth(),  controller.UpdateFinalProject)
	app.Delete("/final-project/:idfinalproject", middleware.Auth(),  controller.DeleteFinalProject)
}
