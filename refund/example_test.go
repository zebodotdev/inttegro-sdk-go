package refund_test

import (
	"context"

	inttegro "github.com/zebodotdev/inttegro-sdk-go/v6"
	"github.com/zebodotdev/inttegro-sdk-go/v6/money"
	"github.com/zebodotdev/inttegro-sdk-go/v6/refund"
)

func ExampleService_Create() {
	client := inttegro.NewClient("sk_test_example")
	_, _ = client.Refunds.Create(context.Background(), refund.CreateParams{
		OrderID: "or_example",
		Reason:  refund.ReasonRequestedByCustomer,
		LineItems: []refund.CreateLineItem{{
			OrderLineItemID: "oli_example",
			RefundAmount: money.AmountParams{
				Currency: money.GHS,
				Value:    2500,
			},
		}},
	})
}
