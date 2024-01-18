package main

import (
	"flag"
)

func main() {
	listenAddr := flag.String("listenaddr", ":3000", "Service is running")
	flag.Parse()
	svc := NewLoggingService(NewMetricService(&priceFetcher{}))

	server := NewJSONAPIServer(*listenAddr, svc)
	server.Run()
}
