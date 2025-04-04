package handlers

import (
	brazatokensapi "crypto-braza-tokens-dashboard/clients/braza-tokens-api"
	bta "crypto-braza-tokens-dashboard/clients/braza-tokens-api"
	"crypto-braza-tokens-dashboard/web/layout"
	ovw "crypto-braza-tokens-dashboard/web/views/operations"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/sync/errgroup"
)

type OperationsHandler struct {
	BtCli *bta.BrazaTokensApiClient
}

func (h OperationsHandler) GetOperations(ctx *fiber.Ctx) error {
	// divide the operations into multiple goroutines to fetch data concurrently
	// using errgroup to handle errors from multiple goroutines
	// and wait for all of them to finish
	// and return the result
	// if any of the goroutines return an error, it will be returned to the caller
	// and the other goroutines will be cancelled
	// and the result will be nil
	eg := new(errgroup.Group)

	var types []*bta.OperationTypesResponse
	var domains []*bta.OperationDomainsResponse
	var blockchains []*bta.BlockchainsResponse
	var tokens []*bta.TokensResponse

	eg.Go(func() error {
		var err error
		types, err = h.BtCli.GetOperationsTypes(ctx.UserContext())
		return err
	})

	eg.Go(func() error {
		var err error
		domains, err = h.BtCli.GetOperationsDomains(ctx.UserContext())
		return err
	})

	eg.Go(func() error {
		var err error
		blockchains, err = h.BtCli.GetBlockchains(ctx.UserContext())
		return err
	})

	eg.Go(func() error {
		var err error
		tokens, err = h.BtCli.GetTokens(ctx.UserContext())
		return err
	})

	if err := eg.Wait(); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err)
	}

	// check if any of the data is nil
	// if any of the data is nil, return an error
	if types == nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON("failed to fetch data: types is nil")
	}
	if domains == nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON("failed to fetch data: domains is nil")
	}
	if blockchains == nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON("failed to fetch data: blockchains is nil")
	}
	if tokens == nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON("failed to fetch data: tokens is nil")
	}

	// check if any of the data is empty
	// if any of the data is empty, return an error
	if len(types) == 0 {
		return ctx.Status(fiber.StatusInternalServerError).JSON("no data found: types is empty")
	}
	if len(domains) == 0 {
		return ctx.Status(fiber.StatusInternalServerError).JSON("no data found: domains is empty")
	}
	if len(blockchains) == 0 {
		return ctx.Status(fiber.StatusInternalServerError).JSON("no data found: blockchains is empty")
	}
	if len(tokens) == 0 {
		return ctx.Status(fiber.StatusInternalServerError).JSON("no data found: tokens is empty")
	}

	// create the input data for the operation
	// and return the data to the caller
	data := &brazatokensapi.InputOperation{
		OperationTypes:   types,
		OperationDomains: domains,
		Blockchain:       blockchains,
		Tokens:           tokens,
	}

	return Render(ctx, layout.Base(ovw.Operations(data)))
}
