package router

import (
	cpt "crypto-braza-tokens-admin/web/handlers/components"
	ch "crypto-braza-tokens-admin/web/handlers/core"
	ph "crypto-braza-tokens-admin/web/handlers/pages"

	"github.com/gofiber/fiber/v2"
)

func BuildRoutes(app *fiber.App) {
	app.Static("/static", "./web/static")
	app.Get("/", ph.HomePagesHandler{}.Index) // if has session, redirect to /home/index else redirect to /auth/login

	pagesRoutes(app)
	componentesRoutes(app)
	coreRoutes(app)
}

func pagesRoutes(app *fiber.App) {
	pages := app.Group("/pages")
	pages.Get("/auth/login", ph.AuthPagesHandler{}.LoginPage)
	pages.Get("/home/index", ph.HomePagesHandler{}.Index)
	pages.Get("/settings/kvs", ph.SettingsPagesHandler{}.KvsPage)
}

func componentesRoutes(app *fiber.App) {
	components := app.Group("/components")
	components.Get("/select-list", cpt.ComponentsHandler{}.SelectList)
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
