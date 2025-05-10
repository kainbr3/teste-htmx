package handlers

import (
	cfg "crypto-braza-tokens-admin/configs"
	repo "crypto-braza-tokens-admin/repositories/mongo"
	l "crypto-braza-tokens-admin/utils/logger"
	cpt "crypto-braza-tokens-admin/web/components"
	sp "crypto-braza-tokens-admin/web/pages/settings"
	"fmt"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

type SettingsHandler struct {
	Configs *cfg.Configs
}

func (h SettingsHandler) KvsPage(ctx *fiber.Ctx) error {
	repo := repo.NewMongoRepository[repo.KeyValue]("keys_values_namespaces_old")

	namespacesResult, err := repo.Distinct(ctx.UserContext(), bson.M{}, "namespace")
	if err != nil {
		l.Logger.Error("Error fetching namespaces: "+err.Error(), zap.Error(err))
		return ctx.Status(fiber.StatusInternalServerError).SendString("Error fetching namespaces")
	}

	// Convert namespacesResult to a slice of strings
	namespaces := make([]string, len(namespacesResult))
	for i, ns := range namespacesResult {
		namespaces[i] = ns.(string)
	}

	kvsOptions := map[string]string{}
	for _,namespace := range namespaces {
		kvsOptions[namespace] = namespace
	}

	// fmt.Println("Namespaces:", namespaces)

	// result, err := repo.Find(ctx.Context(), nil, nil, nil, nil)
	// if err != nil {
	// 	panic(err)
	// }

	// filteredOptions := map[string]string{}
	// for _, item := range result {
	// 	_, exists := filteredOptions[item.Namespace]
	// 	if !exists {
	// 		filteredOptions[item.Namespace] = item.ID.Hex()
	// 	}
	// }

	// kvsOptions := map[string]string{}
	// for key, value := range filteredOptions {
	// 	kvsOptions[value] = key
	// }

	htmx := templ.Attributes{
		"hx-get":     "/test",
		"hx-target":  "#container-variables-table",
		"hx-swap":    "innerHTML",
		"hx-trigger": "change",
	}

	return Render(ctx, sp.Kvs(kvsOptions, htmx))
}

func (h SettingsHandler) KvsTableAdd(ctx *fiber.Ctx) error {

	htmxSave := templ.Attributes{
		"hx-post":    "/settings/kvs/add-save",
		"hx-target":  "closest tr",
		"hx-swap":    "outerHTML",
		"hx-trigger": "click",
	}

	htmxCancel := templ.Attributes{
		"hx-get":     "/teste2",
		"hx-target":  "closest tr",
		"hx-swap":    "delete",
		"hx-trigger": "click",
	}

	return Render(ctx, cpt.TableInsertRow(htmxSave, htmxCancel))
}

func (h SettingsHandler) KvsTableEdit(ctx *fiber.Ctx) error {
	variableID := ctx.Params("variable_id")

	id, err := primitive.ObjectIDFromHex(variableID)
	if err != nil {
		l.Logger.Error("Error converting variable ID to ObjectID: "+variableID, zap.Error(err))
		return ctx.Status(fiber.StatusBadRequest).SendString("Invalid variable ID")
	}

	repo := repo.NewMongoRepository[repo.KeyValue]("keys_values_namespaces_old")

	result, err := repo.FindOne(ctx.UserContext(), bson.M{"_id": id})
	if err != nil {
		l.Logger.Error("Error fetching variable: "+variableID, zap.Error(err))
		return ctx.Status(fiber.StatusInternalServerError).SendString("Error fetching variable")
	}

	htmxUpdate := templ.Attributes{
		"hx-patch":   fmt.Sprintf(`/settings/kvs/edit-save/%s`, result.ID.Hex()),
		"hx-target":  "closest tr",
		"hx-swap":    "outerHTML",
		"hx-trigger": "click",
		"hx-include": "closest form",
	}

	htmxCancel := templ.Attributes{
		"hx-get":     fmt.Sprintf(`/settings/kvs/edit-cancel/%s`, result.ID.Hex()),
		"hx-target":  "closest tr",
		"hx-swap":    "outerHTML",
		"hx-trigger": "click",
	}

	return Render(ctx, cpt.TableEditRow(result.ID.Hex(), result.Key, result.Value, htmxUpdate, htmxCancel))
}

func (h SettingsHandler) KvsTableEditCancel(ctx *fiber.Ctx) error {
	variableID := ctx.Params("variable_id")

	id, err := primitive.ObjectIDFromHex(variableID)
	if err != nil {
		l.Logger.Error("Error converting variable ID to ObjectID: "+variableID, zap.Error(err))
		return ctx.Status(fiber.StatusBadRequest).SendString("Invalid variable ID")
	}

	repo := repo.NewMongoRepository[repo.KeyValue]("keys_values_namespaces_old")

	result, err := repo.FindOne(ctx.UserContext(), bson.M{"_id": id})
	if err != nil {
		l.Logger.Error("Error fetching variable: "+variableID, zap.Error(err))
		return ctx.Status(fiber.StatusInternalServerError).SendString("Error fetching variable")
	}

	htmxEdit := templ.Attributes{
		"hx-get":     fmt.Sprintf(`/settings/kvs/edit/%s`, result.ID.Hex()),
		"hx-target":  "closest tr",
		"hx-swap":    "outerHTML",
		"hx-trigger": "click",
	}

	htmxDelete := templ.Attributes{
		"hx-delete":  fmt.Sprintf(`/settings/kvs/delete/%s`, result.ID.Hex()),
		"hx-target":  "closest tr",
		"hx-swap":    "delete",
		"hx-trigger": "click",
	}

	return Render(ctx, cpt.TableReadRow(result.ID.Hex(), result.Key, result.Value, htmxEdit, htmxDelete))
}

func (h SettingsHandler) KvsTableDelete(ctx *fiber.Ctx) error {
	variableID := ctx.Params("variable_id")

	id, err := primitive.ObjectIDFromHex(variableID)
	if err != nil {
		l.Logger.Error("Error converting variable ID to ObjectID: "+variableID, zap.Error(err))
		return ctx.Status(fiber.StatusBadRequest).SendString("Invalid variable ID")
	}

	repo := repo.NewMongoRepository[repo.KeyValue]("keys_values_namespaces_old")

	err = repo.DeleteOne(ctx.UserContext(), bson.M{"_id": id})
	if err != nil {
		l.Logger.Error("Error deleting variable: "+variableID, zap.Error(err))
		return ctx.Status(fiber.StatusInternalServerError).SendString("Error deleting variable")
	}

	return ctx.SendStatus(fiber.StatusOK)
}
