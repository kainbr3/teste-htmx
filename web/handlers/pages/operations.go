package pages

import (
	shrd "crypto-braza-tokens-admin/web/handlers/shared"
	op "crypto-braza-tokens-admin/web/pages/operations"

	"github.com/gofiber/fiber/v2"
)

type OperationsPagesHandler struct{}

func (h OperationsPagesHandler) Execute(ctx *fiber.Ctx) error {
	return shrd.Render(ctx, op.ExecutePage())
}

func (h OperationsPagesHandler) History(ctx *fiber.Ctx) error {
	return shrd.Render(ctx, op.HistoryPage())
}
