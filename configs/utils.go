package configs

import (
	"log"
	"os"
)

func validateRequiredEnvs() {
	requiredEnvs := []string{
		"LISTENING_PORT",
		"NAMESPACE",
		"MONGO_URI",
		"MONGO_DATABASE",
		"KVS_MONGO_URI",
		"KVS_DATABASE",
		"KVS_COLLECTION",
	}

	for _, env := range requiredEnvs {
		if os.Getenv(env) == "" {
			log.Fatalf("Required environment variable %s is not set", env)
		}
	}
}
