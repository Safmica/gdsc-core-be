package main

import (
	"gdsc-core-be/database"
	"gdsc-core-be/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.EnvInit()
	database.Databaseinit()
	database.DBMigration()
	app := fiber.New()

	routes.UserRoutes(app)
	routes.RoleRoutes(app)
	routes.ParticipantRoutes(app)
	routes.MemberRoutes(app)
	routes.FinalProjectRoutes(app)
	routes.DivisionRoutes(app)
	routes.ConfigureRoutes(app)
	routes.BatchRoutes(app)
	routes.ActivityRoutes(app)

	app.Listen(":3000")
}