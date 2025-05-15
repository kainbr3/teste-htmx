package web

import (
	brazatokens "crypto-braza-tokens-admin/clients/braza-tokens"
	c "crypto-braza-tokens-admin/constants"
	"crypto-braza-tokens-admin/web/router"
	"os"

	"github.com/gofiber/fiber/v2"
)

func Start() {
	brazaTokensCli, err := brazatokens.NewBrazaTokensApiClient()
	if err != nil {
		panic(err)
	}
	// Create a new Fiber instance
	app := fiber.New()

	router.BuildRoutes(app, brazaTokensCli)

	port := c.APPLICATION_DEFAULT_PORT
	if os.Getenv(c.LISTENING_PORT) != "" {
		port = os.Getenv(c.LISTENING_PORT)
	}

	err = app.Listen(":" + port)
	if err != nil {
		panic(err)
	}
}
