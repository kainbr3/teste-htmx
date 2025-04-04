package repositories

import (
	"context"
	kvs "crypto-braza-tokens-dashboard/utils/keys-values"
	l "crypto-braza-tokens-dashboard/utils/logger"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

var repo *Repository

type Repository struct {
	database             *mongo.Database
	whitelistCollection  *mongo.Collection
	keysvaluesCollection *mongo.Collection
}

func NewRepository() *Repository {
	if repo != nil {
		return repo
	}

	connectionString := os.Getenv("MONGO_URI")
	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		l.Logger.Fatal("repository: failed to connect to mongo instance", zap.Error(err))
	}

	brazaTokensDatabase, err := kvs.Get("MONGO_BRAZA_TOKENS_DATABASE")
	if err != nil {
		l.Logger.Fatal("repository: " + err.Error())
	}
	database := client.Database(brazaTokensDatabase)

	whitelistCollection, err := kvs.Get("MONGO_WHITELIST_COLLECTION")
	if err != nil {
		l.Logger.Fatal("repository: " + err.Error())
	}
	whitelist := database.Collection(whitelistCollection)

	keysvaluesCollection := os.Getenv("KVS_COLLECTION")
	keysvalues := database.Collection(keysvaluesCollection)

	repo = &Repository{
		database,
		whitelist,
		keysvalues,
	}

	return repo
}

func (r *Repository) CheckHealth(ctx context.Context) error {
	return r.database.Client().Ping(ctx, nil)
}
