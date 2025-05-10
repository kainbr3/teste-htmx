package web

import (
	"context"
	cfg "crypto-braza-tokens-admin/configs"
	"crypto-braza-tokens-admin/repositories/mongo"
	l "crypto-braza-tokens-admin/utils/logger"
	cpt "crypto-braza-tokens-admin/web/components"
	h "crypto-braza-tokens-admin/web/handlers"
	"crypto-braza-tokens-admin/web/pages/home"
	"fmt"
	"time"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func router(app *fiber.App, configs *cfg.Configs) {
	// Static files to serve the web application
	app.Static("/static", "./web/static")

	// Authorization middleware
	app.Use(func(ctx *fiber.Ctx) error {
		return h.PermissionsMiddleware(ctx)
	})

	app.Get("/home", func(ctx *fiber.Ctx) error {
		return h.Render(ctx, home.Home())
	})

	app.Get("/settings/kvs", h.SettingsHandler{Configs: configs}.KvsPage)

	app.Get("/settings/kvs/add", h.SettingsHandler{Configs: configs}.KvsTableAdd)

	app.Post("/settings/kvs/add-save", func(ctx *fiber.Ctx) error {
		key := ctx.FormValue("key")
		if key == "" {
			return ctx.Status(fiber.StatusBadRequest).SendString("Key is required")
		}

		value := ctx.FormValue("value")
		if value == "" {
			return ctx.Status(fiber.StatusBadRequest).SendString("Value is required")
		}

		namespace := ctx.FormValue("namespace")
		if namespace == "" {
			namespace = "braza-tokens-dashboard"
			// return ctx.Status(fiber.StatusBadRequest).SendString("Namespace is required")
		}

		repo := mongo.NewMongoRepository[mongo.KeyValue]("keys_values_namespaces_old")
		newKeyValue := &mongo.KeyValue{
			ID:        primitive.NewObjectID(),
			Key:       key,
			Value:     value,
			Namespace: namespace,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			UpdatedBy: "ninja",
		}

		id, err := repo.InsertOne(ctx.Context(), newKeyValue)
		if err != nil {
			l.Logger.Error("Error inserting new key-value pair into MongoDB", zap.Error(err))
			return ctx.Status(fiber.StatusInternalServerError).SendString("Error inserting new key-value pair")
		}

		parsedID := id.(primitive.ObjectID)

		htmxSave := templ.Attributes{
			"hx-get":     fmt.Sprintf(`/settings/kvs/edit/%s`, parsedID.Hex()),
			"hx-target":  "closest tr",
			"hx-swap":    "outerHTML",
			"hx-trigger": "click",
		}

		htmxDelete := templ.Attributes{
			"hx-delete":  fmt.Sprintf(`/settings/kvs/delete/%s`, parsedID.Hex()),
			"hx-target":  "closest tr",
			"hx-swap":    "delete",
			"hx-trigger": "click",
		}

		return h.Render(ctx, cpt.TableReadRow(parsedID.Hex(), key, value, htmxSave, htmxDelete))
	})

	app.Get("/settings/kvs/edit/:variable_id", h.SettingsHandler{Configs: configs}.KvsTableEdit)

	app.Patch("/settings/kvs/edit-save/:variable_id", func(ctx *fiber.Ctx) error {
		variableID := ctx.Params("variable_id")
		if variableID == "" {
			return ctx.Status(fiber.StatusBadRequest).SendString("Variable ID is required")
		}

		id, err := primitive.ObjectIDFromHex(variableID)
		if err != nil {
			l.Logger.Error("Error converting variable ID to ObjectID: "+variableID, zap.Error(err))
			return ctx.Status(fiber.StatusBadRequest).SendString("Invalid variable ID")
		}

		key := ctx.FormValue("key")
		if key == "" {
			return ctx.Status(fiber.StatusBadRequest).SendString("Key is required")
		}

		value := ctx.FormValue("value")
		if value == "" {
			return ctx.Status(fiber.StatusBadRequest).SendString("Value is required")
		}

		repo := mongo.NewMongoRepository[mongo.KeyValue]("keys_values_namespaces_old")
		_, err = repo.UpdateOne(ctx.Context(), bson.M{"_id": id}, bson.M{"$set": bson.M{"key": key, "value": value}})
		if err != nil {
			l.Logger.Error("Error updating variable in MongoDB", zap.Error(err))
			return ctx.Status(fiber.StatusInternalServerError).SendString("Error updating variable")
		}

		htmxEdit := templ.Attributes{
			"hx-get":     fmt.Sprintf(`/settings/kvs/edit/%s`, variableID),
			"hx-target":  "closest tr",
			"hx-swap":    "outerHTML",
			"hx-trigger": "click",
		}

		htmxDelete := templ.Attributes{
			"hx-delete":  fmt.Sprintf(`/settings/kvs/delete/%s`, variableID),
			"hx-target":  "closest tr",
			"hx-swap":    "delete",
			"hx-trigger": "click",
		}

		return h.Render(ctx, cpt.TableReadRow(variableID, key, value, htmxEdit, htmxDelete))
	})

	app.Get("/settings/kvs/edit-cancel/:variable_id", h.SettingsHandler{Configs: configs}.KvsTableEditCancel)

	app.Delete("/settings/kvs/delete/:variable_id", h.SettingsHandler{Configs: configs}.KvsTableDelete)

	app.Get("/operations/execute", h.OperationsHandler{Configs: configs}.ExecuteOperationPage)

	app.Get("/test", func(ctx *fiber.Ctx) error {
		value := ctx.FormValue("teste")
		if value == "" {
			return ctx.Status(fiber.StatusBadRequest).SendString("Value is required")
		}

		repo := mongo.NewMongoRepository[mongo.KeyValue]("keys_values_namespaces_old")

		result, err := repo.Find(context.Background(), bson.M{"namespace": value}, bson.M{"key": 1}, nil, nil)
		if err != nil {
			fmt.Println(err)
		}

		return h.Render(ctx, cpt.Table("variables-table", "braza-tokens-dashboard", result))
	})

	app.Get("/teste2", func(ctx *fiber.Ctx) error {
		return ctx.SendStatus(fiber.StatusOK)
	})

	// Login Routes
	app.Get("/login", h.Login)        // Tela de Login
	app.Get("/sso", h.SSO)            // Rota que chama a autenticação SSO
	app.Get("/", h.GetAuthorizations) // Rota de callback do AWS Cognito para obter o code do Query *(para alterar necessário infra)*
	app.Get("/logout", h.Logout)      // Logout
}
