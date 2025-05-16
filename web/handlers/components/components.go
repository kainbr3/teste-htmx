package components

import (
	"context"
	btcli "crypto-braza-tokens-admin/clients/braza-tokens"
	"crypto-braza-tokens-admin/repositories/mongo"
	cpt "crypto-braza-tokens-admin/web/components"
	shrd "crypto-braza-tokens-admin/web/handlers/shared"
	"fmt"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

type ComponentsHandler struct {
	BtCli *btcli.BrazaTokensApiClient
}

func (h ComponentsHandler) SelectListKvs(ctx *fiber.Ctx) error {
	namespace := ctx.FormValue("selected-option")
	if namespace == "" {
		return shrd.RenderError(ctx, fiber.StatusBadRequest, fmt.Errorf("Namespace is required"), false)
	}

	repo := mongo.NewMongoRepository[mongo.KeyValue]("keys_values_namespaces")

	result, err := repo.Find(context.Background(), bson.M{"namespace": namespace}, bson.M{"key": 1}, nil, nil)
	if err != nil {
		return shrd.RenderError(ctx, fiber.StatusInternalServerError, err, false)
	}

	return shrd.Render(ctx, cpt.Table("variables-table", namespace, result))
}

func (h ComponentsHandler) SelectListTypes(ctx *fiber.Ctx) error {
	result, err := h.BtCli.GetOperationsTypes(ctx.UserContext())
	if err != nil {
		return shrd.RenderError(ctx, fiber.StatusInternalServerError, err, false)
	}

	if len(result) == 0 {
		return shrd.RenderError(ctx, fiber.StatusInternalServerError, fmt.Errorf("any operation types found"), false)
	}

	options := map[string]string{}
	for _, item := range result {
		options[item.ID] = item.Name
	}

	htmx := templ.Attributes{
		"hx-get":       "/components/select-list-operations-domains",
		"hx-target":    "#container-domains-selector",
		"hx-swap":      "outerHTML",
		"hx-trigger":   "change",
		"hx-target-4*": "#error-wrapper",
		"hx-target-5*": "#error-wrapper",
	}

	return shrd.Render(ctx, cpt.SelectList("types-selector", "Type", "0", "Select a type...", options, htmx))
}

func (h ComponentsHandler) SelectListDomains(ctx *fiber.Ctx) error {
	result, err := h.BtCli.GetOperationsDomains(ctx.UserContext())
	if err != nil {
		return shrd.RenderError(ctx, fiber.StatusInternalServerError, err, false)
	}

	if len(result) == 0 {
		return shrd.RenderError(ctx, fiber.StatusInternalServerError, fmt.Errorf("any operation domains found"), false)
	}

	options := map[string]string{}
	for _, item := range result {
		options[item.ID] = item.Name
	}

	htmx := templ.Attributes{
		"hx-get":       "/components/select-list-operations-blockchains",
		"hx-target":    "#container-blockchains-selector",
		"hx-swap":      "outerHTML",
		"hx-trigger":   "change",
		"hx-target-4*": "#error-wrapper",
		"hx-target-5*": "#error-wrapper",
	}

	return shrd.Render(ctx, cpt.SelectList("domains-selector", "Domain", "0", "Select a domain...", options, htmx))
}

func (h ComponentsHandler) SelectListBlockchains(ctx *fiber.Ctx) error {
	result, err := h.BtCli.GetBlockchains(ctx.UserContext())
	if err != nil {
		return shrd.RenderError(ctx, fiber.StatusInternalServerError, err, false)
	}

	if len(result) == 0 {
		return shrd.RenderError(ctx, fiber.StatusInternalServerError, fmt.Errorf("any blockchain found"), false)
	}

	options := map[string]string{}
	for _, item := range result {
		options[item.ID] = item.Name
	}

	htmx := templ.Attributes{
		"hx-get":       "/components/select-list-operations-tokens",
		"hx-target":    "#container-tokens-selector",
		"hx-swap":      "outerHTML",
		"hx-trigger":   "change",
		"hx-target-4*": "#error-wrapper",
		"hx-target-5*": "#error-wrapper",
	}

	return shrd.Render(ctx, cpt.SelectList("blockchains-selector", "Blockchain", "0", "Select a blockchain...", options, htmx))
}

func (h ComponentsHandler) SelectListTokens(ctx *fiber.Ctx) error {
	result, err := h.BtCli.GetTokens(ctx.UserContext())
	if err != nil {
		return shrd.RenderError(ctx, fiber.StatusInternalServerError, err, false)
	}

	if len(result) == 0 {
		return shrd.RenderError(ctx, fiber.StatusInternalServerError, fmt.Errorf("any token found"), false)
	}

	options := map[string]string{}
	for _, item := range result {
		options[item.ID] = item.Name
	}

	htmx := templ.Attributes{
		"hx-get":       "/components/input-decimal",
		"hx-target":    "#input-amount",
		"hx-swap":      "outerHTML",
		"hx-trigger":   "change",
		"hx-target-4*": "#error-wrapper",
		"hx-target-5*": "#error-wrapper",
	}

	return shrd.Render(ctx, cpt.SelectList("tokens-selector", "Token", "0", "Select a token...", options, htmx))
}
