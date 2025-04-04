package main

import (
	k "crypto-braza-tokens-dashboard/utils/keys-values"
	l "crypto-braza-tokens-dashboard/utils/logger"
	"crypto-braza-tokens-dashboard/web"
	"os"
)

func main() {
	l.NewLogger()
	validateRequiredEnvs()
	k.Start()
	web.Start()
}

func validateRequiredEnvs() {
	requiredEnvs := []string{
		"NAMESPACE",
		"SERVER_PORT",
		"KVS_MONGO_URI",
		"KVS_DATABASE",
		"KVS_COLLECTION",
		"MONGO_URI",
	}

	for _, env := range requiredEnvs {
		if os.Getenv(env) == "" {
			l.Logger.Sugar().Fatalf("Required environment variable %s is not set", env)
		}
	}
}
