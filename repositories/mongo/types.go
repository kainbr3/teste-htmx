package mongo

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// type Namespace struct {
// 	ID          primitive.ObjectID `bson:"_id"`
// 	Environment string             `bson:"environment"`
// 	Name        string             `bson:"namespace"`
// 	CreatedAt   time.Time          `bson:"created_at"`
// 	UpdatedAt   time.Time          `bson:"updated_at"`
// 	UpdatedBy   string             `bson:"updated_by"`
// }

// type Variable struct {
// 	ID          primitive.ObjectID `bson:"_id"`
// 	Environment string             `bson:"environment"`
// 	NamespaceID string             `bson:"namespace_id"`
// 	Key         string             `bson:"key"`
// 	Value       string             `bson:"value"`
// 	CreatedAt   time.Time          `bson:"created_at"`
// 	UpdatedAt   time.Time          `bson:"updated_at"`
// 	UpdatedBy   string             `bson:"updated_by"`
// }

type KeyValue struct {
	ID          primitive.ObjectID `bson:"_id"`
	Environment string             `bson:"environment,omitempty"`
	Namespace   string             `bson:"namespace"`
	Key         string             `bson:"key"`
	Value       string             `bson:"value"`
	CreatedAt   time.Time          `bson:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at"`
	UpdatedBy   string             `bson:"updated_by"`
}
