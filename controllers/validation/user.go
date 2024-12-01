package validation

import (
	"errors"
	"gdsc-core-be/database"
	"gdsc-core-be/models"
	"gdsc-core-be/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func UserLoginValidation(user *models.User, ctx *fiber.Ctx) (uuid.UUID, string, string, error) {
	var existingUser models.User
	var token models.Token

	if user.IDUser != uuid.Nil {
		return uuid.UUID{}, "", "", errors.New("id_user is not allowed to be input manually")
	}

	if user.Email == "" {
		return uuid.UUID{}, "", "", errors.New("NIDN is required")
	}

	if user.Password == "" {
		return uuid.UUID{}, "", "", errors.New("password is required")
	}

	result := database.DB.Where("email = ?", user.Email).First(&existingUser)
	if result.Error != nil {
		return uuid.UUID{}, "", "", errors.New("wrong credentials")
	}

	isValid := utils.VerifyPassword(existingUser.Password, user.Password)
	if !isValid {
		return uuid.UUID{}, "", "", errors.New("wrong credentials")
	}

	accessToken, refreshToken, errGenerateToken := utils.GenerateTokens(existingUser)
	if errGenerateToken != nil {
		return uuid.UUID{}, "", "", errGenerateToken
	}

	tokenResult := database.DB.Where("id_user = ?", existingUser.IDUser).First(&token)
	if tokenResult.Error != nil && !errors.Is(tokenResult.Error, gorm.ErrRecordNotFound) {
		return uuid.UUID{}, "", "", tokenResult.Error
	}

	if tokenResult.RowsAffected > 0 {
		updateResult := database.DB.Model(&token).Where("id_user = ?", existingUser.IDUser).Update("refresh_token", refreshToken)
		if updateResult.Error != nil {
			return uuid.UUID{}, "", "", updateResult.Error
		}
	} else {
		token = models.Token{
			IDUser:       existingUser.IDUser,
			RefreshToken: refreshToken,
		}
		saveResult := database.DB.Create(&token)
		if saveResult.Error != nil {
			return uuid.UUID{}, "", "", saveResult.Error
		}
	}

	return existingUser.IDUser, accessToken, refreshToken, nil
}


func QueryResultValidation(result *gorm.DB, entityName string) error {
	if result.Error != nil {
		return errors.New("Failed to fetch " + entityName)
	}

	return nil
}

func ParseAndIDValidation(ctx *fiber.Ctx, param string, entityName string) (uint64, error) {
	idParam := ctx.Params(param)

	entityID, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil || entityID == 0 {
		return 0, errors.New("Invalid " + entityName + "_id")
	}

	return entityID, nil
}

func EntityByIDValidation(result *gorm.DB, entityName string) error {
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return errors.New(entityName + " not found")
	}

	QueryResultValidation(result, entityName)

	return nil
}