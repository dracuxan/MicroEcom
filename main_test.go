package main

import (
	"testing"

	"MicroEcom/sdk/client"
	"MicroEcom/sdk/client/products"
)

func TestOurClient(t *testing.T) {
	cfg := client.DefaultTransportConfig().WithHost("localhost:9090")
	c := client.NewHTTPClientWithConfig(nil, cfg)

	params := products.NewGetProductsParams()

	_, err := c.Products.GetProducts(params)
	if err != nil {
		t.Fatalf("Error getting products: %v", err)
	}
}
