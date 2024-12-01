package controller

import (
	"fmt"
	"gdsc-core-be/controllers/validation"
	"gdsc-core-be/database"
	"gdsc-core-be/models"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetAllMember(ctx *fiber.Ctx) error {
	var members []models.Member

	result := database.DB.Find(&members)
	if err := validation.QueryResultValidation(result, "member"); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"members": members,
	})
}

func GetMember(ctx *fiber.Ctx) error {
	memberID := ctx.Params("idmember")
	member := models.Member{}

	idUUID, err := uuid.Parse(memberID)
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

	result := database.DB.Where("id_member = ?", idUUID).First(&member)
	if err := validation.QueryResultValidation(result, "member"); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"member": member,
	})
}

func CreateMember(ctx *fiber.Ctx) error {
	member := new(models.Member)

	if err := ctx.BodyParser(member); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	user := models.User{}

	database.DB.Where("email = ?", member.Email).First(&user)

	member.IDMember = uuid.New()
	member.IDUser = user.IDUser

	if err := database.DB.Debug().Create(&member).Error; err != nil {
		fmt.Printf("Error: %v\n", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create member",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "created successfully",
		"member":    member,
	})
}

func DeleteMember(ctx *fiber.Ctx) error {
	memberID := ctx.Params("idmember")
	idUUID, err := uuid.Parse(memberID)
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
	member := models.Member{}
	
	result := database.DB.Unscoped().Where("id_member = ?", idUUID).Delete(&member)

	if result.Error != nil {
		fmt.Println("Failed to delete all records:", result.Error)
	} else {
		fmt.Printf("Successfully deleted %d records\n", result.RowsAffected)
	}

	return ctx.JSON(fiber.Map{
		"message": "delete successfully",
	})
}