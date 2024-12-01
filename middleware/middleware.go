package middleware

import (
	"gdsc-core-be/cookies"
	// "gdsc-core-be/database"
	// "gdsc-core-be/models"
	"gdsc-core-be/utils"
	"time"

	"github.com/gofiber/fiber/v2"
)

func BlockRoleAuth(deniedRole string) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		AccToken := cookies.GetJwtFromCookie(ctx)
		if AccToken == "" {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "unauthenticated",
			})
		}

		claims, err := utils.DecodeJwtWithRole(AccToken)
		if err != nil {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "invalid token",
			})
		}

		// role := models.Role{}
		namaRole := claims["role"].(string)
		// result := database.DB.Where("id_role =?", idRole).First(&role)
		// if result.Error != nil {
		// 	return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		// 		"message": "invalid token",
		// 	})
		// }
		if namaRole == deniedRole {
			return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"message": "forbidden",
			})
		}

		exp, ok := claims["exp"].(float64)
		if !ok || time.Unix(int64(exp), 0).Before(time.Now()) {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "access token expired",
			})
		}
		return ctx.Next()
	}
}

func Auth() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		AccToken := cookies.GetJwtFromCookie(ctx)
		if AccToken == "" {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "unauthenticated",
			})
		}
		claims, err := utils.DecodeJwtWithRole(AccToken)
		if err != nil {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "unauthenticated",
			})
		}
		exp, ok := claims["exp"].(float64)
		if !ok || time.Unix(int64(exp), 0).Before(time.Now()) {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "access token expired",
			})
		}

		return ctx.Next()
	}
}