package controller

import (
	"gdsc-core-be/controllers/validation"
	"gdsc-core-be/database"
	"gdsc-core-be/models"

	"github.com/gofiber/fiber/v2"
)

func UpdateConfigure(ctx *fiber.Ctx) error {
	configureID := ctx.Params("idconfigure")
	configureNew := new(models.Configure)
	configure := models.Configure{}

	if err := ctx.BodyParser(configureNew); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	result := database.DB.Where("id_configure =?", configureID).First(&configure)
	if err := validation.EntityByIDValidation(result, "dosen"); err != nil {
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