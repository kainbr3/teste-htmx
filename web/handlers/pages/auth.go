package pages

import (
	shrd "crypto-braza-tokens-admin/web/handlers/shared"
	ap "crypto-braza-tokens-admin/web/pages/auth"

	"github.com/gofiber/fiber/v2"
)

type AuthPagesHandler struct{}

func (h AuthPagesHandler) LoginPage(ctx *fiber.Ctx) error {
	return shrd.Render(ctx, ap.LoginPage())
}
