# Inttegro Go SDK

[![OpenSSF Scorecard](https://api.scorecard.dev/projects/github.com/zebodotdev/inttegro-sdk-go/badge)](https://scorecard.dev/viewer/?uri=github.com/zebodotdev/inttegro-sdk-go)

The official Go client for building server-side Inttegro integrations.

> **Fastest, most modern path:** connect an agent to [Inttegro MCP](https://studio.inttegro.com/inttegro-mcp) at `https://mcp.inttegro.com`, then ask it to run `design_integration`. It will produce an implementation and test plan for your application. Use this SDK when you are ready to connect that plan to your Go service.

All official Inttegro SDKs expose the same API capabilities. This module adds Go-specific types, concurrency, and transport control.

## Install

```bash
go get github.com/zebodotdev/inttegro-sdk-go/v4
```

Store your secret key in the server environment:

```bash
export INTTEGRO_API_KEY="your_secret_key"
```

Never put the key in browser code, a mobile app, or source control. The client uses `https://api.inttegro.com` by default.

## Create a hosted checkout

Create and finalize an order, then send the customer to its hosted invoice URL:

```go
package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	inttegro "github.com/zebodotdev/inttegro-sdk-go/v4"
	"github.com/zebodotdev/inttegro-sdk-go/v4/money"
)

func main() {
	client := inttegro.NewClient(os.Getenv("INTTEGRO_API_KEY"))

	order, err := client.Orders.Create(context.Background(), inttegro.OrderCreateParams{
		RequestMeta: &inttegro.RequestMeta{IdempotencyKey: "checkout-cart-123"},
		CustomerData: &inttegro.CustomerData{
			Name: "Akua Mensah", Email: "akua@example.com", PhoneNumber: "+233544998605",
		},
		Finalize: inttegro.Bool(true),
		CheckoutSettings: &inttegro.CheckoutSettings{
			RedirectURL: "https://example.com/orders/complete",
			CancelURL:   "https://example.com/cart",
		},
		LineItems: []inttegro.OrderLineItemParams{{
            Type: inttegro.LineItemTypeProduct,
            Product: &inttegro.ProductLineItemParams{
                Type: inttegro.ProductTypeDigital, Name: "Monthly subscription", Quantity: 1,
                Price: inttegro.PriceParams{AmountParams: money.AmountParams{
                    Currency: money.GHS, Value: 5000,
                }},
			},
		}},
		BillingDetails: inttegro.BillingDetails{
			Name: "Akua Mensah", Email: "akua@example.com", PhoneNumber: "+233544998605",
			Address: inttegro.Address{
				Name: "Akua Mensah", PhoneNumber: "+233544998605",
				Line1: "23 High Street", Town: "Accra", Country: "GH",
			},
		},
	})
	if err != nil {
		var apiErr *inttegro.APIError
		if errors.As(err, &apiErr) {
			log.Fatalf("%s: %s", apiErr.Code, apiErr.Detail)
		}
		log.Fatal(err)
	}

	if order.Invoice == nil || order.Invoice.Format == nil || order.Invoice.Format.Web == nil {
		log.Fatal("order did not include a checkout URL")
	}
	fmt.Println(order.ID, order.Invoice.Format.Web.URL)
}
```

Amounts use integer minor units: `5000` GHS is GHS 50.00. Reuse the same idempotency key when retrying the same logical write. If you omit one, the SDK generates a UUIDv7 key for mutating calls.

## Refund paid line items

Create a partial or full refund by selecting the paid order line items and the
amount to return from each one:

```go
lineReason := inttegro.RefundReasonItemDamaged
refund, err := client.Refunds.Create(context.Background(), inttegro.CreateRefundRequest{
	OrderID: "or_0123456789abcdefghijklmnopqrstuvwxyzABCD",
	Reason:  inttegro.RefundReasonRequestedByCustomer,
	LineItems: []inttegro.CreateRefundLineItem{{
		OrderLineItemID: "oli_abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMN",
		RefundAmount:    money.AmountParams{Currency: money.GHS, Value: 2500},
		Reason:          &lineReason,
		ReasonDetails:   "damaged in transit",
	}},
	Reference: "RETURN-123",
})
if err != nil {
	log.Fatal(err)
}
fmt.Println(refund.ID, refund.Status)
```

Refunds return funds to the original payment method and are processed
asynchronously. Use `Refunds.Lookup` to retrieve the current status.

## Work with the API

The SDK covers orders and checkout, customers, products and prices, purchase intents, payment methods, balances, payouts and refunds, notifications, files, application settings, keys, and country specifications. Services use exported fields such as `PurchaseIntents` and `PaymentMethods`.

Go-specific features:

- Typed request and domain structs with exported constants for public enum values.
- `context.Context` on every operation for deadlines and cancellation.
- Safe sharing across goroutines.
- Standard-library HTTP with no runtime dependencies.
- `WithHTTPClient` and `WithBaseURL` options for proxies, connection pools, tests, and custom timeouts.
- Package-level godoc for resource methods and models.

See the [API reference](https://studio.inttegro.com/api-reference) for request fields and lifecycle rules, [errors](https://studio.inttegro.com/errors) for recovery guidance, and [idempotency](https://studio.inttegro.com/idempotency) for safe retries.

## Verify a release

Go installs the module from the tagged Git repository, directly or through a module proxy, and records module checksums in `go.sum`. The corresponding GitHub release contains an archive of the exact tagged commit, SHA-256 checksums, and a Sigstore attestation tied to the release workflow.

```bash
sha256sum --check SHA256SUMS
gh attestation verify inttegro-sdk-go-4.2.0.tar.gz \
  --repo zebodotdev/inttegro-sdk-go
```

## Develop

```bash
GOWORK=off go test ./...
```
