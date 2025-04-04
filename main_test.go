package main

import (
	"context"
	"crypto-braza-tokens-dashboard/repositories"
	kvs "crypto-braza-tokens-dashboard/utils/keys-values"
	l "crypto-braza-tokens-dashboard/utils/logger"

	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	repo *repositories.Repository
)

func TestMain(m *testing.M) {
	l.NewLogger()
	kvs.Start()

	repo = repositories.NewRepository()

	os.Exit(m.Run())
}

func TestFindAllNamespaces(t *testing.T) {
	ctx := context.Background()
	namespaces, err := repo.FindAllNamespaces(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, namespaces)
}

func TestFindAllVariablesFromNamespace(t *testing.T) {
	ctx := context.Background()

	variable, err := repo.FindAllVariablesFromNamespace(ctx, "braza-tokens-dashboard")
	assert.NoError(t, err)
	assert.NotNil(t, variable)

}

func TestFindOneVariablesFromNamespace(t *testing.T) {
	ctx := context.Background()

	variables, err := repo.FindOneVariableFromNamespace(ctx, "test", "test-key", "test-value")
	assert.NoError(t, err)
	assert.NotNil(t, variables)
}

func TestCreateNewVariable(t *testing.T) {
	ctx := context.Background()

	variable := &repositories.ClientNamespace{
		Namespace: "test",
		Key:       "test-key",
		Value:     "test-value",
	}

	ID, err := repo.CreateNewVariableOnNamespace(ctx, variable.Namespace, variable.Key, variable.Value)
	assert.NoError(t, err)
	assert.NotNil(t, ID)
}

func TestDeleteNamespace(t *testing.T) {
	ctx := context.Background()

	err := repo.DeleteEntireNamespace(ctx, "test")
	assert.NoError(t, err)
}

func TestDeleteVariable(t *testing.T) {
	ctx := context.Background()

	err := repo.DeleteVariable(ctx, "test", "test-key1", "test-value1")
	assert.NoError(t, err)
}

func TestEditVariable(t *testing.T) {
	ctx := context.Background()

	filter := &repositories.ClientNamespace{
		Namespace: "test",
		Key:       "test-key",
		Value:     "test-value",
	}

	err := repo.EditVariable(ctx, filter, "test-key1", "test-value1")
	assert.NoError(t, err)
}
