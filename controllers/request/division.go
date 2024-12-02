package controller

import (
	"fmt"
	"gdsc-core-be/controllers/validation"
	"gdsc-core-be/database"
	"gdsc-core-be/models"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetAllDivision(ctx *fiber.Ctx) error {
	var divisions []models.Division

	result := database.DB.Find(&divisions)
	if err := validation.QueryResultValidation(result, "division"); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"divisions": divisions,
	})
}

func GetDivision(ctx *fiber.Ctx) error {
	divisionID := ctx.Params("iddivision")
	division := models.Division{}

	idUUID, err := uuid.Parse(divisionID)
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

	result := database.DB.Where("id_division = ?", idUUID).First(&division)
	if err := validation.QueryResultValidation(result, "division"); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"division": division,
	})
}

func CreateDivision(ctx *fiber.Ctx) error {
	batchID := ctx.Params("idbatch")
	division := new(models.Division)

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

	if err := ctx.BodyParser(division); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	division.IDDivision = uuid.New()
	division.IDBatch = idUUID

	if err := database.DB.Debug().Create(&division).Error; err != nil {
		fmt.Printf("Error: %v\n", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create division",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "created successfully",
		"division":    division,
	})
}

func UpdateDivision(ctx *fiber.Ctx) error {
	divisionID := ctx.Params("iddivision")
	divisionNew := new(models.Division)
	division := models.Division{}

	idUUID, err := uuid.Parse(divisionID)
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

	if err := ctx.BodyParser(divisionNew); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	result := database.DB.Where("id_division =?", idUUID).First(&division)
	if err = validation.EntityByIDValidation(result, "dosen"); err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if divisionNew.Name != "" {
		division.Name = divisionNew.Name
	}

	if err := database.DB.Save(&division).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"message":       "updated successfully",
		"division": division,
	})
}

func DeleteDivision(ctx *fiber.Ctx) error {
	divisionID := ctx.Params("iddivision")
	idUUID, err := uuid.Parse(divisionID)
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
	division := models.Division{}
	
	result := database.DB.Unscoped().Where("id_division = ?", idUUID).Delete(&division)

	if result.Error != nil {
		fmt.Println("Failed to delete all records:", result.Error)
	} else {
		fmt.Printf("Successfully deleted %d records\n", result.RowsAffected)
	}

	return ctx.JSON(fiber.Map{
		"message": "delete successfully",
	})
}