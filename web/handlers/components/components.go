package components

import (
	"context"
	"crypto-braza-tokens-admin/repositories/mongo"
	cpt "crypto-braza-tokens-admin/web/components"
	shrd "crypto-braza-tokens-admin/web/handlers/shared"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

type ComponentsHandler struct{}

func (h ComponentsHandler) SelectList(ctx *fiber.Ctx) error {
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
