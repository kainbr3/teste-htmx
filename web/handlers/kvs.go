package handlers

// import (
// 	"crypto-braza-tokens-dashboard/repositories"

// 	"github.com/gofiber/fiber/v2"
// )

// type KvsHandler struct {
// 	repo *repositories.Repository
// }

// func NewKvsHandler(repo *repositories.Repository) *KvsHandler {
// 	return &KvsHandler{repo: repo}
// }

// func (h KvsHandler) GetVariables(ctx *fiber.Ctx) error {
// 	environment := ctx.Params("environment")
// 	namespace := ctx.Params("namespace")
// 	variables, err := h.repo.FindAllVariablesFromNamespace(ctx.UserContext(), environment, namespace)
// 	if err != nil {
// 		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
// 	}

// 	return ctx.JSON(variables)
// }

// func (h KvsHandler) GetVariable(ctx *fiber.Ctx) error {
// 	environment := ctx.Params("environment")
// 	namespace := ctx.Params("namespace")
// 	key := ctx.Params("key")
// 	value := ctx.Params("value")

// 	variable, err := h.repo.FindOneVariableFromNamespace(ctx.UserContext(), environment, namespace, key, value)
// 	if err != nil {
// 		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
// 	}

// 	return ctx.JSON(variable)
// }

// func (h KvsHandler) CreateNamespace(ctx *fiber.Ctx) error {
// 	var body repositories.ClientNamespace
// 	if err := ctx.BodyParser(&body); err != nil {
// 		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
// 	}

// 	_, err := h.repo.CreateNewVariableOnNamespace(ctx.UserContext(), body.Environment, body.Namespace, body.Key, body.Value)
// 	if err != nil {
// 		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
// 	}

// 	return ctx.SendStatus(fiber.StatusCreated)
// }

// func (h KvsHandler) UpdateVariable(ctx *fiber.Ctx) error {
// 	environment := ctx.Params("environment")
// 	namespace := ctx.Params("namespace")
// 	oldKey := ctx.Params("oldkey")
// 	oldValue := ctx.Params("oldvalue")

// 	variableInfoFilter := repositories.ClientNamespace{
// 		Environment: environment,
// 		Namespace:   namespace,
// 		Key:         oldKey,
// 		Value:       oldValue,
// 	}

// 	newKey := ctx.Params("newkey")
// 	newValue := ctx.Params("newvalue")

// 	err := h.repo.EditVariable(ctx.UserContext(), &variableInfoFilter, newKey, newValue)
// 	if err != nil {
// 		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
// 	}

// 	return ctx.SendStatus(fiber.StatusOK)
// }

// func (h KvsHandler) DeleteNamespace(ctx *fiber.Ctx) error {
// 	environment := ctx.Params("environment")
// 	namespace := ctx.Params("namespace")

// 	err := h.repo.DeleteEntireNamespace(ctx.UserContext(), environment, namespace)
// 	if err != nil {
// 		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
// 	}

// 	return ctx.SendStatus(fiber.StatusOK)
// }

// func (h KvsHandler) DeleteVariable(ctx *fiber.Ctx) error {
// 	environment := ctx.Params("environment")
// 	namespace := ctx.Params("namespace")
// 	key := ctx.Params("key")
// 	value := ctx.Params("value")

// 	err := h.repo.DeleteVariable(ctx.UserContext(), environment, namespace, key, value)
// 	if err != nil {
// 		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
// 	}

// 	return ctx.SendStatus(fiber.StatusOK)
// }
