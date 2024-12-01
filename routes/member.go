package routes

import (
	controller "gdsc-core-be/controllers/request"

	"github.com/gofiber/fiber/v2"
)

func MemberRoutes(app *fiber.App) {
	app.Get("/all-member",controller.GetAllMember)
	app.Get("/member/:idmember",controller.GetMember)
	app.Post("/member",controller.CreateMember)
	app.Delete("/member/:idmember",controller.DeleteMember)
}
