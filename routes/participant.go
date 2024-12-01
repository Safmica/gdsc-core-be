package routes

import (
	controller "gdsc-core-be/controllers/request"
	"gdsc-core-be/middleware"

	"github.com/gofiber/fiber/v2"
)

func ParticipantRoutes(app *fiber.App) {
	app.Get("/all-participant", middleware.BlockRoleAuth("member"), controller.GetAllParticipant)
	app.Post("/participant/:idactivity", middleware.Auth(),  controller.CreateParticipant)
	app.Delete("/participant/:idparticipant", middleware.Auth(),  controller.DeleteParticipant)
}
