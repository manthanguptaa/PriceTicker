package main

import (
	"context"
	"fmt"
	"time"
)

// PriceFetcher is an interface to fetch the price
type PriceFetcher interface {
	FetchPrice(context.Context, string) (float64, error)
}

// priceFetcher implements the PriceFetcher interface
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
	// mimic the HTTP roundtrip
	time.Sleep(100 * time.Millisecond)

	price, ok := cryptoPriceMocks[ticker]

	if !ok {
		return price, fmt.Errorf("the given ticker (%s) isn't supported", ticker)
	}

	return price, nil
}
