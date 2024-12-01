package controller

import (
	"fmt"
	"gdsc-core-be/controllers/validation"
	"gdsc-core-be/database"
	"gdsc-core-be/models"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetAllBatch(ctx *fiber.Ctx) error {
	var batch []models.Batch

	result := database.DB.Find(&batch)
	if err := validation.QueryResultValidation(result, "batch"); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"batch": batch,
	})
}

func GetBatch(ctx *fiber.Ctx) error {
	batchID := ctx.Params("idbatch")
	batch := models.Batch{}

	idUUID, err := uuid.Parse(batchID)
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

	result := database.DB.Where("id_batch = ?", idUUID).First(&batch)
	if err := validation.QueryResultValidation(result, "batch"); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"batch": batch,
	})
}

func CreateBatch(ctx *fiber.Ctx) error {
	batch := new(models.Batch)

	if err := ctx.BodyParser(batch); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	batch.IDBatch = uuid.New()

	if err := database.DB.Debug().Create(&batch).Error; err != nil {
		fmt.Printf("Error: %v\n", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create batch",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "created successfully",
		"batch":    batch,
	})
}

func DeleteBatch(ctx *fiber.Ctx) error {
	batchID := ctx.Params("idbatch")
	idUUID, err := uuid.Parse(batchID)
	if err != nil {
		fmt.Println("Invalid UUID format:", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid UUID format",
		})
	}

	if idUUID == uuid.Nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "id not found",
		})
	}
	batch := models.Batch{}
	
	result := database.DB.Unscoped().Where("id_batch = ?", idUUID).Delete(&batch)

	if result.Error != nil {
		fmt.Println("Failed to delete all records:", result.Error)
	} else {
		fmt.Printf("Successfully deleted %d records\n", result.RowsAffected)
	}

	return ctx.JSON(fiber.Map{
		"message": "delete successfully",
	})
}