package routes

import (
	controller "gdsc-core-be/controllers/request"
	"gdsc-core-be/middleware"

	"github.com/gofiber/fiber/v2"
)

func MemberRoutes(app *fiber.App) {
	app.Get("/all-member", middleware.BlockRoleAuth("member"), controller.GetAllMember)
	app.Get("/member/:idmember",  middleware.Auth(), controller.GetMember)
	app.Post("/member", middleware.BlockRoleAuth("member"), controller.CreateMember)
	app.Delete("/member/:idmember", middleware.BlockRoleAuth("member"), controller.DeleteMember)
}
