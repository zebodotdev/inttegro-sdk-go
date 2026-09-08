package product_test

import (
	"context"

	inttegro "github.com/zebodotdev/inttegro-sdk-go/v4"
	"github.com/zebodotdev/inttegro-sdk-go/v4/product"
)

func ExampleService_Create() {
	client := inttegro.NewClient("sk_test_example")
	_, _ = client.Products.Create(context.Background(), product.CreateParams{
		Type: product.TypeDigital,
		Name: "Monthly subscription",
	})
}
