package main

import (
	"gdsc-core-be/database"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.EnvInit()
	database.Databaseinit()
	database.DBMigration()
	app := fiber.New()

	app.Listen(":3000")
}