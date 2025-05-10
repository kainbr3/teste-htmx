package domain

import (
	"context"
	r "crypto-braza-tokens-admin/repositories"
	m "crypto-braza-tokens-admin/repositories/mongo"
	l "crypto-braza-tokens-admin/utils/logger"

	"go.uber.org/zap"
)

type SettingsDomain struct {
	repo r.IRepository[m.KeyValue]
}

func NewSettingsDomain() *SettingsDomain {
	repoKvs := m.NewMongoRepository[m.KeyValue]("keys_values_namespaces_old")

	return &SettingsDomain{
		repo: repoKvs,
	}
}

func (i *SettingsDomain) GetNamespaces(ctx context.Context, filters, params map[string]any) ([]*m.KeyValue, error) {
	result, err := i.repo.Find(ctx, nil, nil, nil, nil)
	if err != nil {
		l.Logger.Error("settings domain: error getting namespaces", zap.Error(err))
		return nil, err
	}

	return result, nil
}
