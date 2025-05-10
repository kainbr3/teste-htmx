package handlers

import (
	auth "crypto-braza-tokens-admin/clients/braza-auth"
	keysvalues "crypto-braza-tokens-admin/utils/keys-values"
	l "crypto-braza-tokens-admin/utils/logger"
	ap "crypto-braza-tokens-admin/web/pages/auth"

	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
)

func Login(ctx *fiber.Ctx) error {
	return Render(ctx, ap.Login())
}

func SSO(ctx *fiber.Ctx) error {
	urlSSOLogin, err := keysvalues.Get("URL_SSO_LOGIN")
	if err != nil {
		l.Logger.Sugar().Errorf("Failed to get URL_SSO_LOGIN: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).SendString("Failed to get SSO login URL")
	}

	url := fmt.Sprintf("%s?client_id=%s&response_type=code&scope=email+openid+phone&redirect_uri=%s", urlSSOLogin, os.Getenv("AWS_COGNITO_CLIENT_ID_INT"), os.Getenv("AWS_COGNITO_URL_REDIRECT"))

	ctx.Set("HX-Redirect", url)
	return ctx.SendStatus(fiber.StatusOK)
}

func GetAuthorizations(ctx *fiber.Ctx) error {
	code := ctx.Query("code")
	if code == "" {
		l.Logger.Sugar().Error("Code not provided in the query parameters")
		return ctx.Status(fiber.StatusBadRequest).SendString("Code not provided")
	}

	err := auth.GetAuthorizations(ctx, code)
	if err != nil {
		l.Logger.Sugar().Errorf("Failed to get authorizations: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).SendString("Failed to get authorizations")
	}

	return ctx.Redirect("/home", fiber.StatusFound)
}

func Logout(ctx *fiber.Ctx) error {
	token := ctx.Cookies("authorizations")
	if token == "" {
		l.Logger.Sugar().Error("No auth token found in cookies")
		return ctx.Redirect("/login", fiber.StatusFound)
	}

	ctx.ClearCookie("authorizations")

	return ctx.Redirect("/login", fiber.StatusFound)
}
