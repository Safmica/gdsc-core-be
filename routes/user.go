package routes

import (
	controller "gdsc-core-be/controllers/request"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {
	app.Get("/all-user",controller.GetAllUser)
	app.Get("/user/:iduser",controller.GetUser)
	app.Post("/user",controller.CreateUser)
	app.Patch("/user/:iduser",controller.UpdateUser)
	app.Delete("/user/:iduser",controller.DeleteUser)
}
