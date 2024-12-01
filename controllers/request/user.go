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

func GetAllUser(ctx *fiber.Ctx) error {
	var users []models.User

	result := database.DB.Find(&users)
	if err := validation.QueryResultValidation(result, "user"); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"users": users,
	})
}

func GetUser(ctx *fiber.Ctx) error {
	userID := ctx.Params("iduser")
	user := models.User{}

	idUUID, err := uuid.Parse(userID)
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

	result := database.DB.Where("id_user = ?", idUUID).First(&user)
	if err := validation.QueryResultValidation(result, "user"); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"user": user,
	})
}

func CreateUser(ctx *fiber.Ctx) error {
	user := new(models.User)
	configure := new(models.Configure)

	if err := ctx.BodyParser(user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	user.IDUser = uuid.New()
	user.CurrentBatch = configure.CurrentBatch
	hashedPassword, err := utils.HashingPassword(user.Password)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error hashing password",
		})
	}
	user.Password = hashedPassword


	if err := database.DB.Debug().Create(&user).Error; err != nil {
		fmt.Printf("Error: %v\n", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create user",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "created successfully",
		"user":    user,
	})
}

func UpdateUser(ctx *fiber.Ctx) error {
	userID := ctx.Params("iduser")
	userNew := models.User{}
	user := models.User{}

	idUUID, err := uuid.Parse(userID)
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

	if err := ctx.BodyParser(userNew); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	result := database.DB.Where("id_user =?", idUUID).First(&user)
	if err = validation.EntityByIDValidation(result, "dosen"); err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if userNew.Name != "" {
		user.Name = userNew.Name
	}

	if userNew.Email != "" {
		user.Email = userNew.Email
	}

	if userNew.Major != "" {
		user.Major = userNew.Major
	}

	if userNew.NIM != "" {
		user.NIM = userNew.NIM
	}

	if userNew.Year != 0 {
		user.Year = userNew.Year
	}

	if userNew.University != "" {
		user.University = userNew.University
	}

	if userNew.NewPassword != "" {
		hashedPassword, err := utils.HashingPassword(userNew.NewPassword)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Error hashing password",
			})
		}
		user.Password = hashedPassword
	}


	if err := database.DB.Save(&user).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"message":       "updated successfully",
		"user": user,
	})
}

func DeleteUser(ctx *fiber.Ctx) error {
	userID := ctx.Params("iduser")
	idUUID, err := uuid.Parse(userID)
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
	user := models.User{}
	
	result := database.DB.Unscoped().Where("id_user = ?", idUUID).Delete(&user)

	if result.Error != nil {
		fmt.Println("Failed to delete all records:", result.Error)
	} else {
		fmt.Printf("Successfully deleted %d records\n", result.RowsAffected)
	}

	return ctx.JSON(fiber.Map{
		"message": "delete successfully",
	})
}