package controller

import (
	"fmt"
	"gdsc-core-be/controllers/validation"
	"gdsc-core-be/database"
	"gdsc-core-be/models"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetAllActivity(ctx *fiber.Ctx) error {
	var activities []models.Activity

	result := database.DB.Find(&activities)
	if err := validation.QueryResultValidation(result, "activity"); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"activities": activities,
	})
}

func GetActivity(ctx *fiber.Ctx) error {
	activityID := ctx.Params("idactivity")
	activity := models.Activity{}

	idUUID, err := uuid.Parse(activityID)
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

	result := database.DB.Where("id_activity = ?", idUUID).First(&activity)
	if err := validation.QueryResultValidation(result, "activity"); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"activity": activity,
	})
}

func CreateActivity(ctx *fiber.Ctx) error {
	batchID := ctx.Params("idbatch")
	activity := new(models.Activity)

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

	if err := ctx.BodyParser(activity); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	activity.IDActivity = uuid.New()
	activity.IDBatch = idUUID

	if err := database.DB.Debug().Create(&activity).Error; err != nil {
		fmt.Printf("Error: %v\n", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create activity",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "created successfully",
		"activity":    activity,
	})
}

func UpdateActivity(ctx *fiber.Ctx) error {
	activityID := ctx.Params("idactivity")
	activityNew := models.Activity{}
	activity := models.Activity{}

	idUUID, err := uuid.Parse(activityID)
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

	if err := ctx.BodyParser(activityNew); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	result := database.DB.Where("id_activity =?", idUUID).First(&activity)
	if err = validation.EntityByIDValidation(result, "dosen"); err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if activityNew.Name != "" {
		activity.Name = activityNew.Name
	}

	if !activityNew.Date.IsZero() {
		activity.Date = activityNew.Date
	}

	if activityNew.Description != "" {
		activity.Description = activityNew.Description
	}

	if err := database.DB.Save(&activity).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"message":       "updated successfully",
		"activity": activity,
	})
}

func DeleteActivity(ctx *fiber.Ctx) error {
	activityID := ctx.Params("idactivity")
	idUUID, err := uuid.Parse(activityID)
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
	activity := models.Activity{}
	
	result := database.DB.Unscoped().Where("id_activity = ?", idUUID).Delete(&activity)

	if result.Error != nil {
		fmt.Println("Failed to delete all records:", result.Error)
	} else {
		fmt.Printf("Successfully deleted %d records\n", result.RowsAffected)
	}

	return ctx.JSON(fiber.Map{
		"message": "delete successfully",
	})
}