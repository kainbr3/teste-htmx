package pages

import (
	shrd "crypto-braza-tokens-admin/web/handlers/shared"
	hp "crypto-braza-tokens-admin/web/pages/home"

	"github.com/gofiber/fiber/v2"
)

type HomePagesHandler struct{}

func (h HomePagesHandler) Index(ctx *fiber.Ctx) error {
	return shrd.Render(ctx, hp.HomePage())
}
