package routes

import (
	controller "gdsc-core-be/controllers/request"
	"gdsc-core-be/middleware"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {
	app.Get("/all-user", middleware.BlockRoleAuth("member"),controller.GetAllUser)
	app.Get("/user/:iduser", middleware.Auth(), controller.GetUser)
	app.Post("/user", middleware.BlockRoleAuth("member"), controller.CreateUser)
	app.Post("/user/login",controller.UserLogin)
	app.Post("/user/logout",controller.UserLogout)
	app.Patch("/user/:iduser", middleware.Auth(), controller.UpdateUser)
	app.Delete("/user/:iduser", middleware.BlockRoleAuth("member"), controller.DeleteUser)
}
