package order_test

import (
	"context"

	inttegro "github.com/zebodotdev/inttegro-sdk-go/v4"
	"github.com/zebodotdev/inttegro-sdk-go/v4/money"
	"github.com/zebodotdev/inttegro-sdk-go/v4/order"
	"github.com/zebodotdev/inttegro-sdk-go/v4/price"
	"github.com/zebodotdev/inttegro-sdk-go/v4/product"
)

func ExampleService_Create() {
	client := inttegro.NewClient("sk_test_example")
	_, _ = client.Orders.Create(context.Background(), order.CreateParams{
		LineItems: []order.LineItemParams{{
			Type: order.LineItemTypeProduct,
			Product: &order.ProductLineItemParams{
				Type:     product.TypeDigital,
				Name:     "Monthly subscription",
				Quantity: 1,
				Price: price.InlineParams{AmountParams: money.AmountParams{
					Currency: money.GHS,
					Value:    5000,
				}},
			},
		}},
	})
}
