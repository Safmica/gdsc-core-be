package controller

import (
	"fmt"
	"gdsc-core-be/controllers/validation"
	"gdsc-core-be/database"
	"gdsc-core-be/models"
	"gdsc-core-be/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetAllFinalProject(ctx *fiber.Ctx) error {
	var finalProjects []models.FinalProject

	result := database.DB.Find(&finalProjects)
	if err := validation.QueryResultValidation(result, "final_project"); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"final_projects": finalProjects,
	})
}

func GetFinalProject(ctx *fiber.Ctx) error {
	finalProjectID := ctx.Params("idfinalproject")
	finalProject := models.FinalProject{}

	idUUID, err := uuid.Parse(finalProjectID)
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

	result := database.DB.Where("id_final_project = ?", idUUID).First(&finalProject)
	if err := validation.QueryResultValidation(result, "final_project"); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"final_project": finalProject,
	})
}

func CreateFinalProject(ctx *fiber.Ctx) error {
	finalProject := new(models.FinalProject)

	if err := ctx.BodyParser(finalProject); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	accessToken := ctx.Cookies("access_token")
	claims, err := utils.DecodeJwtWithRole(accessToken)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "invalid token",
		})
	}

	member := models.Member{}
	idMember := claims["id_member"]

	database.DB.Where("id_member =?",idMember).First(&member)

	finalProject.IDFinalProject = uuid.New()
	finalProject.IDMember = member.IDMember

	if err := database.DB.Debug().Create(&finalProject).Error; err != nil {
		fmt.Printf("Error: %v\n", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create finalProject",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "created successfully",
		"final_project":    finalProject,
	})
}

func UpdateFinalProject(ctx *fiber.Ctx) error {
	finalProjectID := ctx.Params("idfinalProject")
	finalProjectNew := models.FinalProject{}
	finalProject := models.FinalProject{}

	idUUID, err := uuid.Parse(finalProjectID)
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

	if err := ctx.BodyParser(finalProjectNew); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	result := database.DB.Where("id_final_project =?", idUUID).First(&finalProject)
	if err = validation.EntityByIDValidation(result, "dosen"); err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if finalProjectNew.Title != "" {
		finalProject.Title = finalProjectNew.Title
	}

	if finalProjectNew.Description != "" {
		finalProject.Description = finalProjectNew.Description
	}

	if finalProjectNew.Url != "" {
		finalProject.Url = finalProjectNew.Url
	}

	if err := database.DB.Save(&finalProject).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"message":       "updated successfully",
		"final_project": finalProject,
	})
}

func DeleteFinalProject(ctx *fiber.Ctx) error {
	finalProjectID := ctx.Params("idfinalproject")
	idUUID, err := uuid.Parse(finalProjectID)
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
	finalProject := models.FinalProject{}
	result := database.DB.Unscoped().Where("id_final_project = ?", idUUID).Delete(&finalProject)
	if err = validation.EntityByIDValidation(result, "dosen"); err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if result.Error != nil {
		fmt.Println("Failed to delete all records:", result.Error)
	} else {
		fmt.Printf("Successfully deleted %d records\n", result.RowsAffected)
	}

	return ctx.JSON(fiber.Map{
		"message": "delete successfully",
	})
}