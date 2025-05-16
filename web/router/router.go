package router

import (
	btcli "crypto-braza-tokens-admin/clients/braza-tokens"
	cpt "crypto-braza-tokens-admin/web/handlers/components"
	ch "crypto-braza-tokens-admin/web/handlers/core"
	ph "crypto-braza-tokens-admin/web/handlers/pages"

	"github.com/gofiber/fiber/v2"
)

func BuildRoutes(app *fiber.App, btCli *btcli.BrazaTokensApiClient) {
	app.Static("/static", "./web/static")
	app.Get("/", ph.HomePagesHandler{}.Index) // if has session, redirect to /home/index else redirect to /auth/login

	pagesRoutes(app, btCli)
	componentesRoutes(app, btCli)
	coreRoutes(app)
}

func pagesRoutes(app *fiber.App, btCli *btcli.BrazaTokensApiClient) {
	pages := app.Group("/pages")
	pages.Get("/auth/login", ph.AuthPagesHandler{}.LoginPage)
	pages.Get("/home/index", ph.HomePagesHandler{}.Index)
	pages.Get("/info/dashboard", ph.InfoPagesHandler{}.Dashboard)
	pages.Get("/info/treasury-management", ph.InfoPagesHandler{}.TreasuryManagement)
	pages.Get("/operations/execute", ph.OperationsPagesHandler{BtCli: btCli}.Execute)
	pages.Get("/operations/history", ph.OperationsPagesHandler{BtCli: btCli}.History)
	pages.Get("/transactions/history", ph.TransactionsPagesHandler{}.History)
	pages.Get("/settings/kvs", ph.SettingsPagesHandler{}.KvsPage)
}

func componentesRoutes(app *fiber.App, btCli *btcli.BrazaTokensApiClient) {
	components := app.Group("/components")
	components.Get("/select-list-kvs", cpt.ComponentsHandler{}.SelectListKvs)
	components.Get("/select-list-operations-types", cpt.ComponentsHandler{BtCli: btCli}.SelectListTypes)
	components.Get("/select-list-operations-domains", cpt.ComponentsHandler{BtCli: btCli}.SelectListDomains)
	components.Get("/select-list-operations-blockchains", cpt.ComponentsHandler{BtCli: btCli}.SelectListBlockchains)
	components.Get("/select-list-operations-tokens", cpt.ComponentsHandler{BtCli: btCli}.SelectListTokens)
}

func coreRoutes(app *fiber.App) {
	core := app.Group("/core")
	core.Get("/none", ch.CoreHandler{}.StatusOk)
	core.Get("/kvs/add", ch.CoreHandler{}.KvsTableRowAdd)
	core.Post("/kvs/add-save", ch.CoreHandler{}.KvsTableRowAddSave)
	core.Get("/kvs/edit/:variable_id", ch.CoreHandler{}.KvsTableRowEdit)
	core.Get("/kvs/edit-cancel/:variable_id", ch.CoreHandler{}.KvsTableRowEditCancel)
	core.Patch("/kvs/edit-save/:variable_id", ch.CoreHandler{}.KvsTableRowEditSave)
	core.Delete("/kvs/delete/:variable_id", ch.CoreHandler{}.KvsTableRowDelete)
}
