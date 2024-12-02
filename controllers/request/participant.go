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

func GetAllParticipant(ctx *fiber.Ctx) error {
	var participants []models.Participant

	result := database.DB.Find(&participants)
	if err := validation.QueryResultValidation(result, "participant"); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"participants": participants,
	})
}

func CreateParticipant(ctx *fiber.Ctx) error {
	activityID := ctx.Params("idactivity")
	participant := new(models.Participant)

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

	participant.IDParticipant = uuid.New()

	accessToken := ctx.Cookies("access_token")
	claims, err := utils.DecodeJwtWithRole(accessToken)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "invalid token",
		})
	}

	member := models.Member{}
	idMember := claims["id_member"].(string)

	idUUID, err = uuid.Parse(idMember)
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

	database.DB.Where("id_member =?",idMember).First(&member)

	participant.IDMember = member.IDMember
	participant.IDActivity = idUUID

	if err := database.DB.Debug().Create(&participant).Error; err != nil {
		fmt.Printf("Error: %v\n", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create participant",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "created successfully",
		"participant":    participant,
	})
}

func DeleteParticipant(ctx *fiber.Ctx) error {
	participantID := ctx.Params("idparticipant")
	idUUID, err := uuid.Parse(participantID)
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
	participant := models.Participant{}
	
	result := database.DB.Unscoped().Where("id_participant = ?", idUUID).Delete(&participant)

	if result.Error != nil {
		fmt.Println("Failed to delete all records:", result.Error)
	} else {
		fmt.Printf("Successfully deleted %d records\n", result.RowsAffected)
	}

	return ctx.JSON(fiber.Map{
		"message": "delete successfully",
	})
}