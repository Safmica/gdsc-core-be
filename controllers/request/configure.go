package controller

import (
	"fmt"
	"gdsc-core-be/controllers/validation"
	"gdsc-core-be/database"
	"gdsc-core-be/models"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func UpdateConfigure(ctx *fiber.Ctx) error {
	configureID := ctx.Params("idconfigure")
	configureNew := models.Configure{}
	configure := models.Configure{}

	idUUID, err := uuid.Parse(configureID)
	if err != nil {
		fmt.Println("Invalid UUID format:", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid UUID format",
		})
	}

	if idUUID == uuid.Nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "id not valid",
		})
	}

	if err := ctx.BodyParser(configureNew); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	result := database.DB.Where("id_configure =?", idUUID).First(&configure)
	if err = validation.EntityByIDValidation(result, "dosen"); err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if configureNew.CurrentBatch != 0 {
		configure.CurrentBatch = configureNew.CurrentBatch
	}

	if err := database.DB.Save(&configure).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"message":       "updated successfully",
		"configure": configure,
	})
}