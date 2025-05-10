package repositories

import "context"

type IRepository[T any] interface {
	Ping(ctx context.Context) error
	Find(ctx context.Context, filter, sort, limit, skip any) ([]*T, error)
	FindOne(ctx context.Context, filter any) (*T, error)
	InsertOne(ctx context.Context, entity *T) (any, error)
	UpdateOne(ctx context.Context, filter, fields any) (any, error)
	DeleteOne(ctx context.Context, filter any) error
	Count(ctx context.Context, filter any) (int64, error)
	Distinct(ctx context.Context, filter, field any) ([]any, error)
}
