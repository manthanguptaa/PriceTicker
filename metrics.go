package main

import (
	"context"
	"fmt"
)

type metricService struct {
	next PriceFetcher
}

func NewMetricService(next PriceFetcher) PriceFetcher {
	return &metricService{
		next: next,
	}
}

func (m *metricService) FetchPrice(ctx context.Context, ticker string) (price float64, err error) {
	fmt.Println("pushing metrics to prometheus")
	// push metrics to prometheus. metric storage
	return m.next.FetchPrice(ctx, ticker)
}
