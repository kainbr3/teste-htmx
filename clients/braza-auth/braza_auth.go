package brazaauth

import (
	k "crypto-braza-tokens-admin/utils/keys-values"
	l "crypto-braza-tokens-admin/utils/logger"
	"crypto-braza-tokens-admin/utils/requests"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

func GetAuthorizations(ctx *fiber.Ctx, code string) error {
	url := fmt.Sprintf("%s/api/auth/callback?code=%s&url=%s", os.Getenv("URL_GET_TOKEN"), code, os.Getenv("AWS_COGNITO_URL_REDIRECT"))
	var requestResponse map[string]any

	err := requests.Execute(ctx.UserContext(), "GET", url, &requestResponse, nil)
	if err != nil {
		return fmt.Errorf("error getting auth token: %w", err)
	}

	url, err = k.Get("ACCESS_CONTROL_UAT")
	if err != nil {
		l.Logger.Sugar().Errorf("Failed to get ACCESS_CONTROL_UAT: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).SendString("Failed to get ACCESS_CONTROL_UAT")
	}

	authToken := requestResponse["accessToken"].(string)
	headers := map[string]any{
		"headers": map[string]string{
			"Authorization": fmt.Sprintf("Bearer %s", authToken),
		},
	}

	var response map[string]any

	err = requests.Execute(ctx.UserContext(), "GET", url, &response, headers)
	if err != nil {
		l.Logger.Sugar().Errorf("Failed to get authorizations: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).SendString("Failed to get authorizations")
	}

	perfis, ok := response["perfis"].([]any)
	if !ok {
		l.Logger.Sugar().Error("Failed to parse perfis from response")
		return ctx.Status(fiber.StatusInternalServerError).SendString("Failed to parse perfis from response")
	}

	expTimeStr, err := k.Get("AUTH_EXP_TIME")
	if err != nil {
		l.Logger.Sugar().Errorf("Failed to get AUTH_EXP_TIME: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).SendString("Failed to get AUTH_EXP_TIME")
	}

	expTime, err := strconv.Atoi(expTimeStr)
	if err != nil {
		l.Logger.Sugar().Errorf("Failed to convert AUTH_EXP_TIME to int: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).SendString("Failed to convert AUTH_EXP_TIME to int")
	}

	perfisStr := fmt.Sprintf("%v", perfis)
	ctx.Cookie(&fiber.Cookie{
		Name:     "authorizations",
		Value:    perfisStr,
		HTTPOnly: true,
		Secure:   true,
		SameSite: "Strict",
		Expires:  time.Now().Add(time.Duration(expTime) * time.Minute),
	})
	return err
}

func BuildProfileMap(perfisStr string) map[string]bool {
	perfisStr = strings.Trim(perfisStr, "[]")
	perfis := strings.Fields(perfisStr)

	profileMap := make(map[string]bool)
	for _, p := range perfis {
		profileMap[p] = true
	}
	return profileMap
}

func HasProfileUsingMap(profileMap map[string]bool, perfil string) bool {
	return profileMap[perfil]
}
