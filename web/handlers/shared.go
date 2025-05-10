package handlers

import (
	l "crypto-braza-tokens-admin/utils/logger"
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func Render(ctx *fiber.Ctx, view templ.Component) error {
	ctx.Set("Content-Type", "text/html")
	return view.Render(ctx.UserContext(), ctx.Response().BodyWriter())
}

func BadRequestWrapper(ctx *fiber.Ctx, resource string, err error) error {
	msg := fmt.Sprintf("handler: error %s", resource)

	l.Logger.Error(msg, zap.Error(err))

	formattedError := fmt.Sprintf("%s with error: %v", msg, err)

	return ctx.Status(http.StatusBadRequest).JSON(ErrorMessage{Message: formattedError})
}
