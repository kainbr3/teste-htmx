package handlers

import (
	k "crypto-braza-tokens-admin/utils/keys-values"
	l "crypto-braza-tokens-admin/utils/logger"

	"strconv"
	"strings"
	"time"

	"slices"

	"github.com/gofiber/fiber/v2"
)

func PermissionsMiddleware(c *fiber.Ctx) error {
	bypassRoutes := map[string]bool{
		"/sso":    true,
		"/":       true,
		"/login":  true,
		"/logout": true,
	}

	if bypassRoutes[c.Path()] {
		return c.Next()
	}

	authCookie := c.Cookies("authorizations")
	if authCookie == "" {
		l.Logger.Error("Authorization cookie not found")
		return c.Redirect("/login", fiber.StatusFound)
	}

	permissions := strings.Split(authCookie, " ")
	if len(permissions) == 0 {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "No permissions found",
		})
	}

	hasPermission := slices.Contains(permissions, "CRYPTO.DASHBOARD.OPERATIONS")

	if !hasPermission {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "Insufficient permissions",
		})
	}

	expTimeStr, err := k.Get("AUTH_EXP_TIME")
	if err != nil {
		l.Logger.Sugar().Errorf("Failed to get AUTH_EXP_TIME: %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to get AUTH_EXP_TIME")
	}

	expTime, err := strconv.Atoi(expTimeStr)
	if err != nil {
		l.Logger.Sugar().Errorf("Failed to convert AUTH_EXP_TIME to int: %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to convert AUTH_EXP_TIME to int")
	}

	c.Cookie(&fiber.Cookie{
		Name:     "authorizations",
		Value:    authCookie,
		HTTPOnly: true,
		Secure:   true,
		SameSite: "Strict",
		Expires:  time.Now().Add(time.Duration(expTime) * time.Minute),
	})

	return c.Next()
}
