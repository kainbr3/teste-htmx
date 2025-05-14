package web

import (
	c "crypto-braza-tokens-admin/constants"
	"crypto-braza-tokens-admin/web/router"
	"os"

	"github.com/gofiber/fiber/v2"
)

func Start() {
	// Create a new Fiber instance
	app := fiber.New()

	router.BuildRoutes(app)

	port := c.APPLICATION_DEFAULT_PORT
	if os.Getenv(c.LISTENING_PORT) != "" {
		port = os.Getenv(c.LISTENING_PORT)
	}

	err := app.Listen(":" + port)
	if err != nil {
		panic(err)
	}
}
