package brazatokensapi_test

import (
	"context"
	brazatokensapi "crypto-braza-tokens-admin/clients/braza-tokens-api"
	kvs "crypto-braza-tokens-admin/utils/keys-values"
	l "crypto-braza-tokens-admin/utils/logger"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBrazaTokensApiClient(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, _ := brazatokensapi.NewBrazaTokensApiClient()

	// Test GetOperation
	operationID := "some-operation-id"
	operationResponse, err := client.GetOperation(ctx, operationID)
	require.NoError(t, err)
	assert.NotNil(t, operationResponse)
	assert.Equal(t, operationID, operationResponse.ID)

	// Test ListOperations
	paginatedOperationsResponse, err := client.GetOperations(ctx)
	require.NoError(t, err)
	assert.NotNil(t, paginatedOperationsResponse)
	assert.Greater(t, len(paginatedOperationsResponse.Data), 0)
}

var client *brazatokensapi.BrazaTokensApiClient

func TestMain(m *testing.M) {
	// Load environment variables
	l.NewLogger("")
	kvs.Start()

	client, _ = brazatokensapi.NewBrazaTokensApiClient()

	os.Exit(m.Run())
}

func Test_Get_Operation(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Test GetOperation
	operationID := "674dfda936158e475705ccb3"
	operationResponse, err := client.GetOperation(ctx, operationID)
	require.NoError(t, err)
	assert.NotNil(t, operationResponse)
	assert.Equal(t, operationID, operationResponse.ID)
}

func Test_List_Operations(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Test ListOperations
	paginatedOperationsResponse, err := client.GetOperations(ctx)
	require.NoError(t, err)
	assert.NotNil(t, paginatedOperationsResponse)
	assert.Greater(t, len(paginatedOperationsResponse.Data), 0)

	// Check if the first operation has a valid ID
	if len(paginatedOperationsResponse.Data) > 0 {
		assert.NotEmpty(t, paginatedOperationsResponse.Data[0].ID)
	}
}

func Test_Get_Operations_Types(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Test GetOperationsTypes
	operationTypesResponse, err := client.GetOperationsTypes(ctx)
	require.NoError(t, err)
	assert.NotNil(t, operationTypesResponse)
}

func Test_Get_Operations_Domains(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Test GetOperationsTypes
	operationTypesResponse, err := client.GetOperationsDomains(ctx)
	require.NoError(t, err)
	assert.NotNil(t, operationTypesResponse)
}

func Test_Get_Blockchains(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Test GetOperationsTypes
	blockchainsResponse, err := client.GetBlockchains(ctx)
	require.NoError(t, err)
	assert.NotNil(t, blockchainsResponse)
}

func Test_Get_Tokens(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Test GetOperationsTypes
	tokensResponse, err := client.GetTokens(ctx)
	require.NoError(t, err)
	assert.NotNil(t, tokensResponse)
}
