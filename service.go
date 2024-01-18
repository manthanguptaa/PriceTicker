package main

import (
	"context"
	"fmt"
)

type PriceFetcher interface {
	FetchPrice(context.Context, string) (float64, error)
}

type priceFetcher struct{}

func (s *priceFetcher) FetchPrice(ctx context.Context, ticker string) (float64, error) {
	return MockPriceFetcher(ctx, ticker)
}

var cryptoPriceMocks = map[string]float64{
	"BTC":  3552871.41,
	"ETH":  211553.44,
	"DOGE": 6.71,
}

func MockPriceFetcher(ctx context.Context, ticker string) (float64, error) {
	price, ok := cryptoPriceMocks[ticker]

	if !ok {
		return price, fmt.Errorf("the given ticker (%s) isn't supported", ticker)
	}

	return price, nil
}
