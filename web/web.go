package web

import (
	"fmt"
	"os"

	bta "crypto-braza-tokens-dashboard/clients/braza-tokens-api"
	"crypto-braza-tokens-dashboard/repositories"
	l "crypto-braza-tokens-dashboard/utils/logger"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func Start() {
	// creates a new fiber instance
	app := fiber.New()

	apiPort := "8080" // default port when not set
	if port := os.Getenv("SERVER_PORT"); port != "" {
		apiPort = port
	}

	// initializes required dependencies
	btCli, err := bta.NewBrazaTokensApiClient()
	if err != nil {
		l.Logger.Fatal("failed creating braza tokens api client", zap.Error(err))
	}

	repo := repositories.NewRepository()

	// creates all the endpoints to be served by the service
	buildRoutes(app, btCli, repo)

	// starts serving the api
	err = app.Listen(":" + apiPort)
	if err != nil {
		l.Logger.Fatal(fmt.Sprintf("failed serving api on port %s", apiPort), zap.Error(err))
	}
}
