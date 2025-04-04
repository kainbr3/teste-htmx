package brazaauth

import (
	"context"
	k "crypto-braza-tokens-dashboard/utils/keys-values"
	l "crypto-braza-tokens-dashboard/utils/logger"
	"crypto-braza-tokens-dashboard/utils/requests"

	"go.uber.org/zap"
)

type BrazaAuthClient struct {
	apiUrl string
}

func NewBrazaAuthClient() *BrazaAuthClient {
	url, err := k.Get("BRAZA_AUTH_API_URL")
	if err != nil {
		l.Logger.Fatal("braza-auth client: error getting braza auth api url", zap.Error(err))
	}

	return &BrazaAuthClient{url}
}

func (b *BrazaAuthClient) GenerateToken(ctx context.Context, clientId, clientSecret string) (*GenerateTokenResponse, error) {
	endpoint := b.apiUrl + "/v1/auth/authorization"

	request := &GenerateTokenRequest{clientId, clientSecret}
	parameters := map[string]any{"payload": request}

	result := &GenerateTokenResponse{}
	err := requests.Execute(ctx, "POST", endpoint, &result, parameters)
	if err != nil {
		l.Logger.Error("braza-auth client: error generating token", zap.Error(err))
		return nil, err
	}

	return result, nil
}

func (b *BrazaAuthClient) DecodeToken(ctx context.Context, token string) (*DecodeTokenResponse, error) {
	endpoint := b.apiUrl + "/v1/auth/valid-token"

	request := &DecodeTokenRequest{AccessToken: token}
	parameters := map[string]any{"payload": request}

	result := &DecodeTokenResponse{}
	err := requests.Execute(ctx, "POST", endpoint, &result, parameters)
	if err != nil {
		l.Logger.Error("braza-auth client: error decoding token", zap.Error(err))
		return nil, err
	}

	return result, nil
}
