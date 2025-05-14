package mongo

import (
	"context"
	c "crypto-braza-tokens-admin/constants"
	"fmt"
	"os"
	"sync"
	"time"

	l "crypto-braza-tokens-admin/utils/logger"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

var (
	clientInstance *mongo.Client
	connectOnce    sync.Once
	databaseOnce   sync.Once
	database       *mongo.Database
)

// initializes and returns a singleton instance of the mongo.Client.
// It uses environment variables to configure the connection URI and ensures the connection is tested before use.
func connect() *mongo.Client {
	connectOnce.Do(func() {
		timeout := 10 * time.Second
		if os.Getenv(c.MONGO_TIMEOUT) != "" {
			var err error
			timeout, err = time.ParseDuration(os.Getenv(c.MONGO_TIMEOUT))
			if err != nil {
				l.Logger.Fatal(fmt.Sprintf("Failed to parse MONGO_TIMEOUT: %v", err))
			}
		}

		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		// fmt.Println(ctx)

		var err error
		clientInstance, err = mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv(c.MONGO_URI)))
		//clientInstance, err = mongo.Connect(context.Background(), options.Client().ApplyURI(os.Getenv(c.MONGO_URI)))
		if err != nil {
			l.Logger.Fatal(fmt.Sprintf("Failed to connect to MongoDB host %s with error: %v", c.MONGO_URI, err))
		}

		// // Testa a conexão
		// if err = clientInstance.Ping(ctx, nil); err != nil {
		// 	l.Logger.Fatal(fmt.Sprintf("Failed to ping MongoDB host %s with error: %v", c.MONGO_URI, err))
		// }

		l.Logger.Info("MongoDB connection established successfully")
	})

	return clientInstance
}

func setDatabase(dbName string) {
	databaseOnce.Do(func() {
		timeout := 10 * time.Second
		if os.Getenv(c.MONGO_TIMEOUT) != "" {
			var err error
			timeout, err = time.ParseDuration(os.Getenv(c.MONGO_TIMEOUT))
			if err != nil {
				l.Logger.Fatal(fmt.Sprintf("Failed to parse MONGO_TIMEOUT: %v", err))
			}
		}

		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		// fmt.Println(ctx)

		if dbName == "" {
			l.Logger.Fatal("MongoDB database name is not set")
		}

		if clientInstance == nil {
			l.Logger.Fatal("MongoDB client is not initialized")
		}

		database = clientInstance.Database(dbName)

		if database == nil {
			l.Logger.Fatal("Failed to set MongoDB database")
		}

		if err := database.Client().Ping(ctx, nil); err != nil {
			// if err := database.Client().Ping(context.Background(), nil); err != nil {
			l.Logger.Fatal(fmt.Sprintf("Failed to ping MongoDB database %s with error: %v", os.Getenv(c.MONGO_DATABASE), err))
		}

		l.Logger.Info("MongoDB database set successfully", zap.String("database", os.Getenv(c.MONGO_DATABASE)))
	})
}

func Start(dbName string) {
	connect()
	setDatabase(dbName)
}
