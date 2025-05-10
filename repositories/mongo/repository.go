package mongo

import (
	"context"

	r "crypto-braza-tokens-admin/repositories"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoRepository[T any] struct {
	collection *mongo.Collection
}

// NewMongoRepository cria uma nova instância do repositório para T
func NewMongoRepository[T any](collectionName string) r.IRepository[T] {
	collection := database.Collection(collectionName)
	return &mongoRepository[T]{collection: collection}
}

func (m *mongoRepository[T]) Ping(ctx context.Context) error {
	err := database.Client().Ping(ctx, nil)
	if err != nil {
		return err
	}
	return nil
}

func (r *mongoRepository[T]) Find(ctx context.Context, filter, sort, limit, skip any) ([]*T, error) {
	if filter == nil {
		filter = bson.M{}
	}

	options := options.Find()
	if sort != nil {
		options.SetSort(sort)
	}

	if limit != nil {
		options.SetLimit(limit.(int64))
	}

	if skip != nil {
		options.SetSkip(skip.(int64))
	}

	cur, err := r.collection.Find(ctx, filter, options)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var results []*T
	for cur.Next(ctx) {
		var elem *T
		if err := cur.Decode(&elem); err != nil {
			return nil, err
		}
		results = append(results, elem)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

func (r *mongoRepository[T]) FindOne(ctx context.Context, filter any) (*T, error) {
	if filter == nil {
		filter = bson.M{}
	}

	res := r.collection.FindOne(ctx, filter)
	if res.Err() != nil {
		return nil, res.Err()
	}

	var result T
	if err := res.Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *mongoRepository[T]) InsertOne(ctx context.Context, document *T) (any, error) {
	res, err := r.collection.InsertOne(ctx, document)
	if err != nil {
		return nil, err
	}

	return res.InsertedID, nil
}

func (r *mongoRepository[T]) UpdateOne(ctx context.Context, filter, fields any) (any, error) {
	res, err := r.collection.UpdateOne(ctx, filter, fields)
	if err != nil {
		return nil, err
	}

	return res.UpsertedID, nil
}

func (r *mongoRepository[T]) DeleteOne(ctx context.Context, filter any) error {
	_, err := r.collection.DeleteOne(ctx, filter)
	return err
}

func (r *mongoRepository[T]) Count(ctx context.Context, filter any) (int64, error) {
	count, err := r.collection.CountDocuments(ctx, filter)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *mongoRepository[T]) Distinct(ctx context.Context, filter, field any) ([]any, error) {
	results, err := r.collection.Distinct(ctx, field.(string), filter)
	if err != nil {
		return nil, err
	}

	return results, nil
}
