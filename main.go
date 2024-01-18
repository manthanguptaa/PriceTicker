package main

import (
	"context"
	"fmt"
	"log"
)

func main() {
	svc := NewLoggingService(&priceFetcher{})

	price, err := svc.FetchPrice(context.Background(), "ETH")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(price)
}
