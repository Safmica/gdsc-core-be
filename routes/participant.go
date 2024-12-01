package routes

import (
	controller "gdsc-core-be/controllers/request"

	"github.com/gofiber/fiber/v2"
)

func ParticipantRoutes(app *fiber.App) {
	app.Get("/all-participant",controller.GetAllParticipant)
	app.Post("/participant",controller.CreateParticipant)
	app.Delete("/participant/:idparticipant",controller.DeleteParticipant)
}
