package validation

import (
	"errors"
	"fmt"
	"gdsc-core-be/database"
	"gdsc-core-be/models"
)

func RefreshTokenValidation(id uint, refreshToken string, token models.Token) error {
	result := database.DB.Debug().Where("id_user =?", id).First(&token)
	if result == nil {
		fmt.Println("Errorr")
		return errors.New("invalid refresh token")
	}
	return nil
}