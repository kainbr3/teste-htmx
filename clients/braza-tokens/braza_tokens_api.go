package brazatokens

import (
	"context"
	k "crypto-braza-tokens-admin/utils/keys-values"
	l "crypto-braza-tokens-admin/utils/logger"
	r "crypto-braza-tokens-admin/utils/requests"
	"fmt"
	"os"

	"go.uber.org/zap"
)

type BrazaTokensApiClient struct {
	apiUrl    string
	apiSecret string
}

func NewBrazaTokensApiClient() (*BrazaTokensApiClient, error) {
	apiUrl, err := k.Get("BRAZA_TOKENS_API_URL")
	if err != nil {
		return nil, err
	}

	apiSecret := os.Getenv("BRAZA_TOKENS_API_SECRET")
	if apiSecret == "" {
		return nil, fmt.Errorf("BRAZA_TOKENS_API_SECRET environment variable is not set")
	}

	return &BrazaTokensApiClient{
		apiUrl:    apiUrl,
		apiSecret: apiSecret,
	}, nil
}

func (b *BrazaTokensApiClient) GetOperationsTypes(ctx context.Context) ([]*OperationTypesResponse, error) {
	endpoint := b.apiUrl + "/v1/operations-types"

	parameters := map[string]any{"headers": map[string]string{"trusted-client": b.apiSecret}}

	result := []*OperationTypesResponse{}
	err := r.Execute(ctx, "GET", endpoint, &result, parameters)
	if err != nil {
		l.Logger.Error("braza-tokens-api client: error generating token", zap.Error(err))
		return nil, err
	}

	return result, nil
}

func (b *BrazaTokensApiClient) GetOperationsDomains(ctx context.Context) ([]*OperationDomainsResponse, error) {
	endpoint := b.apiUrl + "/v1/operations-domains"

	parameters := map[string]any{"headers": map[string]string{"trusted-client": b.apiSecret}}

	result := []*OperationDomainsResponse{}
	err := r.Execute(ctx, "GET", endpoint, &result, parameters)
	if err != nil {
		l.Logger.Error("braza-tokens-api client: error generating token", zap.Error(err))
		return nil, err
	}

	return result, nil
}

func (b *BrazaTokensApiClient) GetBlockchains(ctx context.Context) ([]*BlockchainsResponse, error) {
	endpoint := b.apiUrl + "/v1/blockchains"

	parameters := map[string]any{"headers": map[string]string{"trusted-client": b.apiSecret}}

	result := []*BlockchainsResponse{}
	err := r.Execute(ctx, "GET", endpoint, &result, parameters)
	if err != nil {
		l.Logger.Error("braza-tokens-api client: error generating token", zap.Error(err))
		return nil, err
	}

	return result, nil
}

func (b *BrazaTokensApiClient) GetTokens(ctx context.Context) ([]*TokensResponse, error) {
	endpoint := b.apiUrl + "/v1/tokens"

	parameters := map[string]any{"headers": map[string]string{"trusted-client": b.apiSecret}}

	result := []*TokensResponse{}
	err := r.Execute(ctx, "GET", endpoint, &result, parameters)
	if err != nil {
		l.Logger.Error("braza-tokens-api client: error generating token", zap.Error(err))
		return nil, err
	}

	return result, nil
}

func (b *BrazaTokensApiClient) GetOperation(ctx context.Context, operationID string) (*OperationResponse, error) {
	endpoint := fmt.Sprintf("%s/v1/operations/%s", b.apiUrl, operationID)

	parameters := map[string]any{"headers": map[string]string{"trusted-client": b.apiSecret}}

	result := &OperationResponse{}
	err := r.Execute(ctx, "GET", endpoint, &result, parameters)
	if err != nil {
		l.Logger.Error("braza-tokens-api client: error generating token", zap.Error(err))
		return nil, err
	}

	return result, nil
}

func (b *BrazaTokensApiClient) GetOperations(ctx context.Context) (*PaginatedOperationsResponse, error) {
	endpoint := b.apiUrl + "/v1/operations"

	parameters := map[string]any{"headers": map[string]string{"trusted-client": b.apiSecret}}

	result := &PaginatedOperationsResponse{}
	err := r.Execute(ctx, "GET", endpoint, &result, parameters)
	if err != nil {
		l.Logger.Error("braza-tokens-api client: error generating token", zap.Error(err))
		return nil, err
	}

	return result, nil
}

func (b *BrazaTokensApiClient) PostOperation(ctx context.Context, clientId, clientSecret string) (*OperationResponse, error) {
	endpoint := b.apiUrl + "/v1/operations"

	parameters := map[string]any{"headers": map[string]string{"trusted-client": b.apiSecret}}

	result := &OperationResponse{}
	err := r.Execute(ctx, "POST", endpoint, &result, parameters)
	if err != nil {
		l.Logger.Error("braza-tokens-api client: error generating token", zap.Error(err))
		return nil, err
	}

	return result, nil
}

func (b *BrazaTokensApiClient) GetWalletBalances(ctx context.Context) (*WalletBalancesResponse, error) {
	endpoint := fmt.Sprintf("%s/v1/wallets-balances", b.apiUrl)

	parameters := map[string]any{"headers": map[string]string{"trusted-client": b.apiSecret}}

	result := &WalletBalancesResponse{}
	err := r.Execute(ctx, "GET", endpoint, &result, parameters)
	if err != nil {
		l.Logger.Error("braza-tokens-api client: error generating token", zap.Error(err))
		return nil, err
	}

	return result, nil
}
