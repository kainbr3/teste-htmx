package handlers

import (
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
)

func Render(ctx *fiber.Ctx, view templ.Component) error {
	ctx.Set("Content-Type", "text/html")
	return view.Render(ctx.UserContext(), ctx.Response().BodyWriter())
}

func RenderParams(ctx *fiber.Ctx, view templ.Component, params any) error {
	ctx.Set("Content-Type", "text/html")
	return view.Render(ctx.UserContext(), ctx.Response().BodyWriter())
}
