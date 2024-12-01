package middleware

// import (
// 	"be-merged/cookies"
// 	"be-merged/database"
// 	"be-merged/models"
// 	"be-merged/utils"
// 	"time"

// 	"github.com/gofiber/fiber/v2"
// )

// func RoleBasedAuth(requiredRole string) fiber.Handler {
// 	return func(ctx *fiber.Ctx) error {
// 		AccToken := cookies.GetJwtFromCookie(ctx)
// 		if AccToken == "" {
// 			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
// 				"message": "unauthenticated",
// 			})
// 		}

// 		claims, err := utils.DecodeJwtWithRole(AccToken)
// 		if err != nil {
// 			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
// 				"message": "invalid token",
// 			})
// 		}

// 		role := models.Role{}
// 		idRole := claims["role"].(string)
// 		result := database.DB.Where("id_role =?", idRole).First(&role)
// 		if result.Error != nil {
// 			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
// 				"message": "invalid token",
// 			})
// 		}
// 		if role.Nama != requiredRole {
// 			return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
// 				"message": "forbidden",
// 			})
// 		}

// 		exp, ok := claims["exp"].(float64)
// 		if !ok || time.Unix(int64(exp), 0).Before(time.Now()) {
// 			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
// 				"message": "access token expired",
// 			})
// 		}
// 		return ctx.Next()
// 	}
// }

// func Auth() fiber.Handler {
// 	return func(ctx *fiber.Ctx) error {
// 		AccToken := cookies.GetJwtFromCookie(ctx)
// 		if AccToken == "" {
// 			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
// 				"message": "unauthenticated",
// 			})
// 		}
// 		claims, err := utils.DecodeJwtWithRole(AccToken)
// 		if err != nil {
// 			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
// 				"message": "unauthenticated",
// 			})
// 		}
// 		exp, ok := claims["exp"].(float64)
// 		if !ok || time.Unix(int64(exp), 0).Before(time.Now()) {
// 			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
// 				"message": "access token expired",
// 			})
// 		}

// 		return ctx.Next()
// 	}
// }