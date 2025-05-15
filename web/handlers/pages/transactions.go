package pages

import (
	shrd "crypto-braza-tokens-admin/web/handlers/shared"
	tp "crypto-braza-tokens-admin/web/pages/transactions"

	"github.com/gofiber/fiber/v2"
)

type TransactionsPagesHandler struct{}

func (h TransactionsPagesHandler) History(ctx *fiber.Ctx) error {
	return shrd.Render(ctx, tp.HistoryPage())
}
