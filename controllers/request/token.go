package controller

import (
	"gdsc-core-be/cookies"
	"gdsc-core-be/database"
	"gdsc-core-be/models"
	"gdsc-core-be/utils"
	"time"

	"github.com/gofiber/fiber/v2"
)

func RefreshToken(ctx *fiber.Ctx) error {
	refreshToken := cookies.GetRefreshTokenFromCookie(ctx)
	if refreshToken == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	var token models.Token
	claims, err := utils.DecodeJwtWithRole(refreshToken)
	if err != nil || claims["type"] != "refresh" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "invalid refresh token",
		})
	}
	exp, ok := claims["exp"].(float64)
	if !ok || time.Unix(int64(exp), 0).Before(time.Now()) {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "refresh token expired",
		})
	}

	id := claims["id"]
	role := claims["role"].(string)

	result := database.DB.Debug().Where("id_user = ?", id).First(&token)
	if result.Error != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "user not found or token invalid",
		})
	}

	if refreshToken != token.RefreshToken {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "invalid refresh token",
		})
	}

	user := models.User{}
	batch := models.Batch{}
	member := models.Member{}

	database.DB.Select("current_batch").Where("id_user = ?", id).First(&user)
	database.DB.Select("id_batch").Where("year = ?", user.CurrentBatch).First(&batch)
	database.DB.Select("id_role").Where("id_user = ? AND id_batch = ?", user.IDUser, batch.IDBatch).First(&member)
	database.DB.Select("nama").Where("id_role = ?", member.IDRole).First(&role)

	newAccessToken, err := utils.GenerateNewJwt(token.IDUser, member.IDMember, role, user.CurrentBatch, time.Now().Add(30*time.Second), "access")
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "could not generate new access token",
		})
	}

	cookies.SetJwtCookie(ctx, newAccessToken)

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "access token updated",
	})
}