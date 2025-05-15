package pages

import (
	shrd "crypto-braza-tokens-admin/web/handlers/shared"
	ip "crypto-braza-tokens-admin/web/pages/info"

	"github.com/gofiber/fiber/v2"
)

type InfoPagesHandler struct{}

func (h InfoPagesHandler) Dashboard(ctx *fiber.Ctx) error {
	return shrd.Render(ctx, ip.DashboardPage())
}

func (h InfoPagesHandler) TreasuryManagement(ctx *fiber.Ctx) error {
	return shrd.Render(ctx, ip.TreasuryManagementPage())
}
