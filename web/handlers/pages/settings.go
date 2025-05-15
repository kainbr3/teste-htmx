package pages

import (
	repo "crypto-braza-tokens-admin/repositories/mongo"
	l "crypto-braza-tokens-admin/utils/logger"
	shrd "crypto-braza-tokens-admin/web/handlers/shared"
	sp "crypto-braza-tokens-admin/web/pages/settings"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.uber.org/zap"
)

type SettingsPagesHandler struct{}

func (h SettingsPagesHandler) KvsPage(ctx *fiber.Ctx) error {
	repo := repo.NewMongoRepository[repo.KeyValue]("keys_values_namespaces")

	namespacesResult, err := repo.Distinct(ctx.UserContext(), bson.M{}, "namespace")
	if err != nil {
		l.Logger.Error("Error fetching namespaces", zap.Error(err))
		return shrd.RenderError(ctx, fiber.StatusInternalServerError, err, false)
	}

	namespaces := make([]string, len(namespacesResult))
	for i, ns := range namespacesResult {
		namespaces[i] = ns.(string)
	}

	options := map[string]string{}
	for _, namespace := range namespaces {
		options[namespace] = namespace
	}

	htmx := templ.Attributes{
		"hx-get":       "/components/select-list",
		"hx-target":    "#container-variables-table",
		"hx-swap":      "innterHTML",
		"hx-trigger":   "change",
		"hx-target-4*": "#error-wrapper",
		"hx-target-5*": "#error-wrapper",
	}

	return shrd.Render(ctx, sp.KvsPage(options, htmx))
}
