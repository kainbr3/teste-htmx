package repositories

import (
	"context"
	l "crypto-braza-tokens-dashboard/utils/logger"
	"errors"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

var environment = os.Getenv("ENVIRONMENT")

func (r *Repository) Find(ctx context.Context) ([]*ClientNamespace, error) {
	filter := bson.M{}
	findOptions := options.Find().SetSort(bson.M{"namespace": 1})

	cursor, err := r.keysvaluesCollection.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, fmt.Errorf("error finding keys values: %v", err)
	}
	defer cursor.Close(ctx)

	var result []*ClientNamespace
	if err = cursor.All(ctx, &result); err != nil {
		return nil, fmt.Errorf("error parsing keys values result: %v", err)
	}

	if len(result) == 0 {
		log.Fatal("no key values found")
	}

	return result, nil
}

func (r *Repository) FindAllNamespaces(ctx context.Context) ([]string, error) {
	filter := bson.M{"environment": environment}
	opts := options.Find().SetSort(bson.D{{Key: "namespace", Value: -1}})

	cursor, err := r.keysvaluesCollection.Find(ctx, filter, opts)
	if err != nil {
		l.Logger.Error("repository: error finding Namespace", zap.Error(err))
		return nil, err
	}
	defer cursor.Close(ctx)

	namespaceSet := make(map[string]struct{})
	for cursor.Next(ctx) {
		var record ClientNamespace
		if err := cursor.Decode(&record); err != nil {
			l.Logger.Error("repository: error decoding Namespace record", zap.Error(err))
			return nil, err
		}
		namespaceSet[record.Namespace] = struct{}{}
	}

	if err := cursor.Err(); err != nil {
		l.Logger.Error("repository: cursor error", zap.Error(err))
		return nil, err
	}

	var namespaces []string
	for namespace := range namespaceSet {
		namespaces = append(namespaces, namespace)
	}

	return namespaces, nil
}

func (r *Repository) FindAllVariablesFromNamespace(ctx context.Context, namespace string) ([]*ClientNamespace, error) {
	filter := bson.M{"namespace": namespace, "environment": environment}

	cursor, err := r.keysvaluesCollection.Find(ctx, filter)
	if err != nil {
		l.Logger.Error("repository: error finding Namespace", zap.Error(err))
		return nil, err
	}
	defer cursor.Close(ctx)
	var result []*ClientNamespace
	for cursor.Next(ctx) {
		var record ClientNamespace
		if err := cursor.Decode(&record); err != nil {
			l.Logger.Error("repository: error decoding Namespace record", zap.Error(err))
			return nil, err
		}
		result = append(result, &record)
	}
	if err := cursor.Err(); err != nil {
		l.Logger.Error("repository: cursor error", zap.Error(err))
		return nil, err
	}
	if len(result) == 0 {
		l.Logger.Error("repository: no Namespace found")
		return nil, errors.New("no Namespace found")
	}

	return result, nil
}

func (r *Repository) FindOneNamespace(ctx context.Context, namespace string) (*ClientNamespace, error) {
	filter := bson.M{"namespace": namespace, "environment": environment}

	var result ClientNamespace
	err := r.keysvaluesCollection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		l.Logger.Error("repository: error finding Namespace", zap.Error(err))
		return nil, err
	}

	return &result, nil
}

func (r *Repository) FindOneVariableFromNamespace(ctx context.Context, namespace, key, value string) (*ClientNamespace, error) {
	filter := bson.M{
		"environment": environment,
		"namespace":   namespace,
		"key":         key,
		"value":       value,
	}

	var result ClientNamespace
	err := r.keysvaluesCollection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		l.Logger.Error("repository: error finding variable by fields", zap.Error(err))
		return nil, err
	}

	return &result, nil
}

func (r *Repository) CreateNewVariableOnNamespace(ctx context.Context, namespace, key, value string) (primitive.ObjectID, error) {
	data := ClientNamespace{
		ID:        primitive.NewObjectID(),
		Namespace: namespace,
		Key:       key,
		Value:     value,
	}

	result, err := r.keysvaluesCollection.InsertOne(ctx, data)
	if err != nil {
		l.Logger.Error("repository: error saving variable", zap.Error(err))
		return primitive.NilObjectID, err
	}

	id, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		l.Logger.Error("repository: error converting inserted ID to ObjectID")
		return primitive.NilObjectID, errors.New("failed to retrieve inserted ID")
	}

	return id, nil
}

func (r *Repository) DeleteEntireNamespace(ctx context.Context, namespace string) error {
	filter := bson.M{
		"environment": environment,
		"namespace":   namespace,
	}

	_, err := r.keysvaluesCollection.DeleteMany(ctx, filter)
	if err != nil {
		l.Logger.Error("repository: error deleting Namespace", zap.Error(err))
		return err
	}

	return nil
}

func (r *Repository) DeleteVariable(ctx context.Context, namespace, key, value string) error {
	filter := bson.M{
		"environment": environment,
		"namespace":   namespace,
		"key":         key,
		"value":       value,
	}

	_, err := r.keysvaluesCollection.DeleteOne(ctx, filter)
	if err != nil {
		l.Logger.Error("repository: error deleting variable", zap.Error(err))
		return err
	}

	return nil
}

func (r *Repository) EditVariable(ctx context.Context, variableInfoFilter *ClientNamespace, newKey, newValue string) error {
	filter := bson.M{
		"environment": environment,
		"namespace":   variableInfoFilter.Namespace,
		"key":         variableInfoFilter.Key,
		"value":       variableInfoFilter.Value,
	}

	update := bson.M{"$set": bson.M{
		"key":   newKey,
		"value": newValue,
	}}

	_, err := r.keysvaluesCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		l.Logger.Error("repository: error updating variable", zap.Error(err))
		return err
	}

	return nil
}
