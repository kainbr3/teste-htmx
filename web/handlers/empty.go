package handlers

import (
	"github.com/gofiber/fiber/v2"
)

type EmptyHandler struct{}

func (h EmptyHandler) GetNone(ctx *fiber.Ctx) error {
	return nil
}
