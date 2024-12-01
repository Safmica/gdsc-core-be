package validation

import (
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

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