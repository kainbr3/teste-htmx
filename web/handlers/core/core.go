package core

import (
	"crypto-braza-tokens-admin/repositories/mongo"
	repo "crypto-braza-tokens-admin/repositories/mongo"
	l "crypto-braza-tokens-admin/utils/logger"
	cpt "crypto-braza-tokens-admin/web/components"
	shrd "crypto-braza-tokens-admin/web/handlers/shared"
	"fmt"
	"time"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

type CoreHandler struct{}

func (h CoreHandler) StatusOk(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).SendString("")
}

func (h CoreHandler) KvsTableRowAdd(ctx *fiber.Ctx) error {
	namespace := ctx.Query("namespace")
	if namespace == "" {
		l.Logger.Error("Error fetching namespace from URL")
		return shrd.RenderError(ctx, fiber.StatusBadRequest, fmt.Errorf("namespace is required"), false)
	}

	htmxSave := templ.Attributes{
		"hx-post":      "/core/kvs/add-save",
		"hx-target":    "closest tr",
		"hx-swap":      "outerHTML",
		"hx-trigger":   "click",
		"hx-target-4*": "#error-notifier",
		"hx-target-5*": "#error-notifier",
	}

	htmxCancel := templ.Attributes{
		"hx-get":     "/core/none",
		"hx-target":  "closest tr",
		"hx-swap":    "delete",
		"hx-trigger": "click",
	}

	return shrd.Render(ctx, cpt.TableInsertRow(namespace, htmxSave, htmxCancel))
}

func (h CoreHandler) KvsTableRowAddSave(ctx *fiber.Ctx) error {
	key := ctx.FormValue("key")
	if key == "" {
		return shrd.RenderError(ctx, fiber.StatusBadRequest, fmt.Errorf("key is required"), false)
	}

	value := ctx.FormValue("value")
	if value == "" {
		return shrd.RenderError(ctx, fiber.StatusBadRequest, fmt.Errorf("value is required"), false)
	}

	namespace := ctx.FormValue("namespace")
	if namespace == "" {
		return shrd.RenderError(ctx, fiber.StatusBadRequest, fmt.Errorf("namespace is required"), false)
	}

	repo := mongo.NewMongoRepository[mongo.KeyValue]("keys_values_namespaces")
	newKeyValue := &mongo.KeyValue{
		ID:        primitive.NewObjectID(),
		Key:       key,
		Value:     value,
		Namespace: namespace,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UpdatedBy: "ninja", //change to logged user later
	}

	id, err := repo.InsertOne(ctx.Context(), newKeyValue)
	if err != nil {
		l.Logger.Error("Error inserting new key-value pair into MongoDB", zap.Error(err))
		return shrd.RenderError(ctx, fiber.StatusInternalServerError, err, false)
	}

	parsedID := id.(primitive.ObjectID)

	htmxSave := templ.Attributes{
		"hx-get":       fmt.Sprintf(`/core/kvs/edit/%s`, parsedID.Hex()),
		"hx-target":    "closest tr",
		"hx-swap":      "outerHTML",
		"hx-trigger":   "click",
		"hx-target-4*": "#error-notifier",
		"hx-target-5*": "#error-notifier",
	}

	htmxDelete := templ.Attributes{
		"hx-delete":    fmt.Sprintf(`/core/kvs/delete/%s`, parsedID.Hex()),
		"hx-target":    "closest tr",
		"hx-swap":      "delete",
		"hx-trigger":   "click",
		"hx-target-4*": "#error-notifier",
		"hx-target-5*": "#error-notifier",
	}

	return shrd.Render(ctx, cpt.TableReadRow(parsedID.Hex(), key, value, htmxSave, htmxDelete))
}

func (h CoreHandler) KvsTableRowEdit(ctx *fiber.Ctx) error {
	kvsItemId := ctx.Params("variable_id")
	if kvsItemId == "" {
		l.Logger.Error("Error fetching variable ID from URL")
		return shrd.RenderError(ctx, fiber.StatusBadRequest, fmt.Errorf("variable ID is required"), false)
	}

	id, err := primitive.ObjectIDFromHex(kvsItemId)
	if err != nil {
		l.Logger.Error("Error converting variable ID to ObjectID"+kvsItemId, zap.Error(err))
		return shrd.RenderError(ctx, fiber.StatusInternalServerError, err, false)
	}

	repo := repo.NewMongoRepository[repo.KeyValue]("keys_values_namespaces")

	result, err := repo.FindOne(ctx.UserContext(), bson.M{"_id": id})
	if err != nil {
		l.Logger.Error("Error fetching variable: "+kvsItemId, zap.Error(err))
		return shrd.RenderError(ctx, fiber.StatusInternalServerError, err, false)
	}

	htmxUpdate := templ.Attributes{
		"hx-patch":     fmt.Sprintf(`/core/kvs/edit-save/%s`, result.ID.Hex()),
		"hx-target":    "closest tr",
		"hx-swap":      "outerHTML",
		"hx-trigger":   "click",
		"hx-include":   "closest form",
		"hx-target-4*": "#error-notifier",
		"hx-target-5*": "#error-notifier",
	}

	htmxCancel := templ.Attributes{
		"hx-get":       fmt.Sprintf(`/core/kvs/edit-cancel/%s`, result.ID.Hex()),
		"hx-target":    "closest tr",
		"hx-swap":      "outerHTML",
		"hx-trigger":   "click",
		"hx-target-4*": "#error-wrapper",
		"hx-target-5*": "#error-wrapper",
	}

	return shrd.Render(ctx, cpt.TableEditRow(result.ID.Hex(), result.Key, result.Value, htmxUpdate, htmxCancel))
}

func (h CoreHandler) KvsTableRowEditCancel(ctx *fiber.Ctx) error {
	kvsItemId := ctx.Params("variable_id")
	if kvsItemId == "" {
		l.Logger.Error("Error fetching variable ID from URL")
		return shrd.RenderError(ctx, fiber.StatusBadRequest, fmt.Errorf("variable ID is required"), false)
	}

	id, err := primitive.ObjectIDFromHex(kvsItemId)
	if err != nil {
		l.Logger.Error("Error converting variable ID to ObjectID"+kvsItemId, zap.Error(err))
		return shrd.RenderError(ctx, fiber.StatusInternalServerError, err, false)
	}

	repo := repo.NewMongoRepository[repo.KeyValue]("keys_values_namespaces")

	result, err := repo.FindOne(ctx.UserContext(), bson.M{"_id": id})
	if err != nil {
		l.Logger.Error("Error fetching variable: "+kvsItemId, zap.Error(err))
		return shrd.RenderError(ctx, fiber.StatusInternalServerError, err, false)
	}

	htmxEdit := templ.Attributes{
		"hx-get":       fmt.Sprintf(`/core/kvs/edit/%s`, result.ID.Hex()),
		"hx-target":    "closest tr",
		"hx-swap":      "outerHTML",
		"hx-trigger":   "click",
		"hx-target-4*": "#error-wrapper",
		"hx-target-5*": "#error-wrapper",
	}

	htmxDelete := templ.Attributes{
		"hx-delete":    fmt.Sprintf(`/core/kvs/delete/%s`, result.ID.Hex()),
		"hx-target":    "closest tr",
		"hx-swap":      "delete",
		"hx-trigger":   "click",
		"hx-target-4*": "#error-wrapper",
		"hx-target-5*": "#error-wrapper",
	}

	return shrd.Render(ctx, cpt.TableReadRow(result.ID.Hex(), result.Key, result.Value, htmxEdit, htmxDelete))
}

func (h CoreHandler) KvsTableRowEditSave(ctx *fiber.Ctx) error {
	variableID := ctx.Params("variable_id")
	if variableID == "" {
		return shrd.RenderError(ctx, fiber.StatusBadRequest, fmt.Errorf("variable ID is required"), false)
	}

	id, err := primitive.ObjectIDFromHex(variableID)
	if err != nil {
		l.Logger.Error("Error converting variable ID to ObjectID: "+variableID, zap.Error(err))
		shrd.RenderError(ctx, fiber.StatusInternalServerError, err, false)
	}

	key := ctx.FormValue("key")
	if key == "" {
		return shrd.RenderError(ctx, fiber.StatusBadRequest, fmt.Errorf("Key is required"), false)
	}

	value := ctx.FormValue("value")
	if value == "" {
		return shrd.RenderError(ctx, fiber.StatusBadRequest, fmt.Errorf("Value is required"), false)
	}

	repo := mongo.NewMongoRepository[mongo.KeyValue]("keys_values_namespaces")

	_, err = repo.UpdateOne(ctx.Context(), bson.M{"_id": id}, bson.M{"$set": bson.M{"key": key, "value": value}})
	if err != nil {
		l.Logger.Error("Error updating variable in MongoDB", zap.Error(err))
		return shrd.RenderError(ctx, fiber.StatusInternalServerError, err, false)
	}

	htmxEdit := templ.Attributes{
		"hx-get":       fmt.Sprintf(`/core/kvs/edit/%s`, variableID),
		"hx-target":    "closest tr",
		"hx-swap":      "outerHTML",
		"hx-trigger":   "click",
		"hx-target-4*": "#error-wrapper",
		"hx-target-5*": "#error-wrapper",
	}

	htmxDelete := templ.Attributes{
		"hx-delete":    fmt.Sprintf(`/core/kvs/delete/%s`, variableID),
		"hx-target":    "closest tr",
		"hx-swap":      "delete",
		"hx-trigger":   "click",
		"hx-target-4*": "#error-wrapper",
		"hx-target-5*": "#error-wrapper",
	}

	return shrd.Render(ctx, cpt.TableReadRow(variableID, key, value, htmxEdit, htmxDelete))
}

func (h CoreHandler) KvsTableRowDelete(ctx *fiber.Ctx) error {
	variableID := ctx.Params("variable_id")
	if variableID == "" {
		l.Logger.Error("Error fetching variable ID from URL")
		return shrd.RenderError(ctx, fiber.StatusBadRequest, fmt.Errorf("variable ID is required"), true)
	}

	id, err := primitive.ObjectIDFromHex(variableID)
	if err != nil {
		l.Logger.Error("Error converting variable ID to ObjectID: "+variableID, zap.Error(err))
		return shrd.RenderError(ctx, fiber.StatusInternalServerError, err, true)
	}

	repo := repo.NewMongoRepository[repo.KeyValue]("keys_values_namespaces")

	err = repo.DeleteOne(ctx.UserContext(), bson.M{"_id": id})
	if err != nil {
		l.Logger.Error("Error deleting variable: "+variableID, zap.Error(err))
		return shrd.RenderError(ctx, fiber.StatusInternalServerError, err, true)
	}

	return ctx.SendStatus(fiber.StatusOK)
}
