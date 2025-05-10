package web

import (
	bta "crypto-braza-tokens-admin/clients/braza-tokens-api"
	cfg "crypto-braza-tokens-admin/configs"
	c "crypto-braza-tokens-admin/constants"
	d "crypto-braza-tokens-admin/domain"
	"crypto-braza-tokens-admin/utils/logger"
	"os"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func Start() {
	// Create a new Fiber instance
	app := fiber.New()

	btaCli, err := bta.NewBrazaTokensApiClient()
	if err != nil {
		logger.Logger.Fatal("Failed to create BrazaTokensApiClient", zap.Error(err))
	}

	settings := &cfg.Configs{
		SettingsDomain:   d.NewSettingsDomain(),
		OperationsDomain: d.NewOperationsDomain(btaCli),
	}

	// register the routes and static files to serve the web application
	router(app, settings)

	// Set the port to listen on, defaulting to 8080 if not set in the environment
	port := c.APPLICATION_DEFAULT_PORT
	if os.Getenv(c.LISTENING_PORT) != "" {
		port = os.Getenv(c.LISTENING_PORT)
	}

	// Start the server on the specified port
	err = app.Listen(":" + port)
	if err != nil {
		panic(err)
	}
}
