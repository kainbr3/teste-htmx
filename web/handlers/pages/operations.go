package pages

import (
	btcli "crypto-braza-tokens-admin/clients/braza-tokens"
	shrd "crypto-braza-tokens-admin/web/handlers/shared"
	op "crypto-braza-tokens-admin/web/pages/operations"
	"fmt"

	"golang.org/x/sync/errgroup"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
)

type OperationsPagesHandler struct {
	BtCli *btcli.BrazaTokensApiClient
}

func (h OperationsPagesHandler) Execute(ctx *fiber.Ctx) error {
	var types []*btcli.OperationTypesResponse
	// var domains []*btcli.OperationDomainsResponse
	// var blockchains []*btcli.BlockchainsResponse
	// var tokens []*btcli.TokensResponse
	var eg errgroup.Group

	eg.Go(func() error {
		var err error
		types, err = h.BtCli.GetOperationsTypes(ctx.UserContext())
		if len(types) == 0 {
			return shrd.RenderError(ctx, fiber.StatusInternalServerError, fmt.Errorf("any types found"), false)
		}
		return err
	})

	// eg.Go(func() error {
	// 	var err error
	// 	domains, err = h.BtCli.GetOperationsDomains(ctx.UserContext())
	// 	if len(domains) == 0 {
	// 		return shrd.RenderError(ctx, fiber.StatusInternalServerError, fmt.Errorf("any domains found"), false)
	// 	}
	// 	return err
	// })

	// eg.Go(func() error {
	// 	var err error
	// 	blockchains, err = h.BtCli.GetBlockchains(ctx.UserContext())
	// 	if len(blockchains) == 0 {
	// 		return shrd.RenderError(ctx, fiber.StatusInternalServerError, fmt.Errorf("any blockchains found"), false)
	// 	}
	// 	return err
	// })

	// eg.Go(func() error {
	// 	var err error
	// 	tokens, err = h.BtCli.GetTokens(ctx.UserContext())
	// 	if len(tokens) == 0 {
	// 		return shrd.RenderError(ctx, fiber.StatusInternalServerError, fmt.Errorf("any tokens found"), false)
	// 	}
	// 	return err
	// })

	if err := eg.Wait(); err != nil {
		return shrd.RenderError(ctx, fiber.StatusInternalServerError, err, false)
	}

	typesOptions := map[string]string{}
	for _, item := range types {
		typesOptions[item.ID] = item.Name
	}

	// domainsOptions := map[string]string{}
	// for _, item := range domains {
	// 	domainsOptions[item.ID] = item.Name
	// }

	// blockchainsOptions := map[string]string{}
	// for _, item := range blockchains {
	// 	blockchainsOptions[item.ID] = item.Name
	// }

	// tokensOptions := map[string]string{}
	// for _, item := range tokens {
	// 	tokensOptions[item.ID] = item.Name
	// }

	htmx := templ.Attributes{
		"hx-get":       "/components/select-list-operations-domains",
		"hx-target":    "#container-domains-selector",
		"hx-swap":      "outerHTML",
		"hx-trigger":   "change",
		"hx-target-4*": "#error-wrapper",
		"hx-target-5*": "#error-wrapper",
	}

	// return shrd.Render(ctx, op.ExecuteOperation(typesOptions, domainsOptions, blockchainsOptions, tokensOptions))
	return shrd.Render(ctx, op.ExecuteOperation(typesOptions, nil, nil, nil, htmx))
}

func (h OperationsPagesHandler) History(ctx *fiber.Ctx) error {
	return shrd.Render(ctx, op.HistoryPage())
}
