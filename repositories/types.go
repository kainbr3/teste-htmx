package repositories

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ClientNamespace struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	Environment string             `bson:"environment" json:"environment"`
	Namespace   string             `bson:"namespace" json:"namespace"`
	Key         string             `bson:"key" json:"key"`
	Value       string             `bson:"value" json:"value"`
}
