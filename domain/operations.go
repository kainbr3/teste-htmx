package domain

import (
	"context"
	bta "crypto-braza-tokens-admin/clients/braza-tokens-api"
)

type OperationsDomain struct {
	BtaCli *bta.BrazaTokensApiClient
}

func NewOperationsDomain(btaCli *bta.BrazaTokensApiClient) *OperationsDomain {
	return &OperationsDomain{BtaCli: btaCli}
}

func (i *OperationsDomain) GetOperationsTypes(ctx context.Context) ([]*bta.OperationTypesResponse, error) {
	result, err := i.BtaCli.GetOperationsTypes(ctx)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (i *OperationsDomain) GetOperationsDomains(ctx context.Context) ([]*bta.OperationDomainsResponse, error) {
	result, err := i.BtaCli.GetOperationsDomains(ctx)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (i *OperationsDomain) GetBlockchains(ctx context.Context) ([]*bta.BlockchainsResponse, error) {
	result, err := i.BtaCli.GetBlockchains(ctx)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (i *OperationsDomain) GetTokens(ctx context.Context) ([]*bta.TokensResponse, error) {
	result, err := i.BtaCli.GetTokens(ctx)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (i *OperationsDomain) GetOperation(ctx context.Context, id string) (any, error) {
	result, err := i.BtaCli.GetOperation(ctx, id)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (i *OperationsDomain) GetOperations(ctx context.Context, id string) (*bta.PaginatedOperationsResponse, error) {
	result, err := i.BtaCli.GetOperations(ctx)
	if err != nil {
		return nil, err
	}

	return result, nil
}
