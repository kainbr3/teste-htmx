package web

import (
	bta "crypto-braza-tokens-dashboard/clients/braza-tokens-api"
	"crypto-braza-tokens-dashboard/repositories"
	"crypto-braza-tokens-dashboard/web/components"
	"crypto-braza-tokens-dashboard/web/handlers"
	"crypto-braza-tokens-dashboard/web/layout"
	evw "crypto-braza-tokens-dashboard/web/views/empty"
	kvsviews "crypto-braza-tokens-dashboard/web/views/keys-values"

	"github.com/gofiber/fiber/v2"
)

// buildRoutes - setup api and creates all routes
func buildRoutes(app *fiber.App, btCli *bta.BrazaTokensApiClient, repo *repositories.Repository) {
	// Default path validation and redirect
	app.Get("/", func(ctx *fiber.Ctx) error { return ctx.Redirect("/auth/login") })

	// Static files to serve the web application
	app.Static("/static", "./web/static/dist")

	app.Get("dashboard/overview", func(ctx *fiber.Ctx) error {
		return handlers.Render(ctx, layout.Base(evw.Blank()))
	})

	app.Get("wallets/balances", func(ctx *fiber.Ctx) error {
		return handlers.Render(ctx, layout.Base(evw.Blank()))
	})

	app.Get("/operations/execute", handlers.OperationsHandler{BtCli: btCli}.GetOperations)

	app.Get("operations/history", func(ctx *fiber.Ctx) error {
		return handlers.Render(ctx, layout.Base(evw.Blank()))
	})

	app.Get("operations/:id", func(ctx *fiber.Ctx) error {
		return handlers.Render(ctx, layout.Base(evw.Blank()))
	})

	app.Get("transactions/history", func(ctx *fiber.Ctx) error {
		return handlers.Render(ctx, layout.Base(evw.Blank()))
	})

	app.Get("transactions/:id", func(ctx *fiber.Ctx) error {
		return handlers.Render(ctx, layout.Base(evw.Blank()))
	})

	// app.Get("settings/kvs", func(ctx *fiber.Ctx) error {
	// 	context := ctx.UserContext()
	// 	namespaces, err := repo.FindAllNamespaces(context)
	// 	if err != nil {
	// 		return ctx.Status(fiber.StatusInternalServerError).SendString("Error fetching namespaces")
	// 	}

	// 	return handlers.Render(ctx, layout.Base(kvsviews.Kvs(namespaces)))
	// })

	app.Get("settings/kvs", func(ctx *fiber.Ctx) error {
		variables, err := repo.Find(ctx.UserContext())
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).SendString("Error fetching variables")
		}

		return handlers.Render(
			ctx,
			layout.Base(
				kvsviews.KvsV2(
					components.SelectListNamespaces(variables),
				),
			),
		)
	})

	app.Get("settings/kvs/findall", func(ctx *fiber.Ctx) error {
		variables, err := repo.Find(ctx.UserContext())
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).SendString("Error fetching variables")
		}

		return handlers.Render(ctx, components.SelectListNamespaces(variables))
	})

	app.Get("settings/kvs/details", func(ctx *fiber.Ctx) error {
		context := ctx.UserContext()
		namespace := ctx.Query("selectednamespace")
		if namespace == "" {
			return ctx.Status(fiber.StatusBadRequest).SendString("Namespace is required")
		}

		variables, err := repo.FindAllVariablesFromNamespace(context, namespace)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).SendString("Error fetching variables")
		}

		return handlers.Render(ctx, components.KvsTable(namespace, variables))
	})

	app.Delete("settings/kvs/delete", func(ctx *fiber.Ctx) error {
		context := ctx.UserContext()
		namespace := ctx.Query("selectednamespace")
		if namespace == "" {
			return ctx.Status(fiber.StatusBadRequest).SendString("Namespace is required")
		}

		err := repo.DeleteEntireNamespace(context, namespace)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).SendString("Error deleting namespace")
		}

		return ctx.Status(fiber.StatusOK).JSON(nil)
	})
}
