package cookies

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func SetJwtCookie(ctx *fiber.Ctx, token string) {
	ctx.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    token,
		Expires:  time.Now().Add(time.Second * 30),
		HTTPOnly: true,
	})
}

func GetJwtFromCookie(ctx *fiber.Ctx) string {
	return ctx.Cookies("access_token")
}

func ClearJwtCookie(ctx *fiber.Ctx) {
	ctx.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	})
}

func SetRefreshTokenCookie(ctx *fiber.Ctx, token string) {
	ctx.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    token,
		Expires:  time.Now().Add(24 * time.Hour),
		HTTPOnly: true,
	})
}

func GetRefreshTokenFromCookie(ctx *fiber.Ctx) string {
	return ctx.Cookies("refresh_token")
}

func ClearRefreshTokenCookie(ctx *fiber.Ctx) {
	ctx.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	})
}
