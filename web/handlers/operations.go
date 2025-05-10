package handlers

import (
	cfg "crypto-braza-tokens-admin/configs"
	op "crypto-braza-tokens-admin/web/pages/operations"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type OperationsHandler struct {
	Configs *cfg.Configs
}

func (h OperationsHandler) ExecuteOperationPage(ctx *fiber.Ctx) error {
	types, err := h.Configs.OperationsDomain.GetOperationsTypes(ctx.UserContext())
	if err != nil {
		return BadRequestWrapper(ctx, "KVS Page", fmt.Errorf("Error getting namespaces: %v", err))
	}

	if types == nil || len(types) == 0 {
		return BadRequestWrapper(ctx, "KVS Page", fmt.Errorf("No namespaces found"))
	}

	domains, err := h.Configs.OperationsDomain.GetOperationsDomains(ctx.UserContext())
	if err != nil {
		return BadRequestWrapper(ctx, "KVS Page", fmt.Errorf("Error getting namespaces: %v", err))
	}

	if domains == nil || len(domains) == 0 {
		return BadRequestWrapper(ctx, "KVS Page", fmt.Errorf("No namespaces found"))
	}

	blockchains, err := h.Configs.OperationsDomain.GetBlockchains(ctx.UserContext())
	if err != nil {
		return BadRequestWrapper(ctx, "KVS Page", fmt.Errorf("Error getting namespaces: %v", err))
	}

	if blockchains == nil || len(blockchains) == 0 {
		return BadRequestWrapper(ctx, "KVS Page", fmt.Errorf("No namespaces found"))
	}

	tokens, err := h.Configs.OperationsDomain.GetTokens(ctx.UserContext())
	if err != nil {
		return BadRequestWrapper(ctx, "KVS Page", fmt.Errorf("Error getting namespaces: %v", err))
	}

	if tokens == nil || len(tokens) == 0 {
		return BadRequestWrapper(ctx, "KVS Page", fmt.Errorf("No namespaces found"))
	}

	typesOptions := map[string]string{}
	for _, item := range types {
		typesOptions[item.ID] = item.Name
	}

	domainsOptions := map[string]string{}
	for _, item := range domains {
		domainsOptions[item.ID] = item.Name
	}

	blockchainsOptions := map[string]string{}
	for _, item := range blockchains {
		blockchainsOptions[item.ID] = item.Name
	}

	tokensOptions := map[string]string{}
	for _, item := range tokens {
		tokensOptions[item.ID] = item.Name
	}

	return Render(ctx, op.ExecuteOperation(typesOptions, domainsOptions, blockchainsOptions, tokensOptions))
}
