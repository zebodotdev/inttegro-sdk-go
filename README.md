# Commerce Go SDK

Idiomatic Go client for the Commerce API. Lightweight, stdlib-only HTTP with clear resource services and typed request/response models.

**Comprehensive godoc documentation:** Every type, method, and field is fully documented with examples. Use `go doc` or your editor's intellisense for inline help.

## Installation

```bash
go get github.com/zebodotdev/commerce-sdk-go
```

## Quickstart

```go
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	commerce "github.com/zebodotdev/commerce-sdk-go"
)

func main() {
	client := commerce.NewClient(os.Getenv("COMMERCE_API_KEY"))

	order, err := client.Orders.Create(context.Background(), commerce.OrderCreateParams{
		CustomerData: &commerce.CustomerData{
			Name:        "Akua Mensah",
			PhoneNumber: "+233544998605",
			Email:       "akua@example.com",
		},
		PaymentMethodData: &commerce.PaymentMethodData{
			Type: commerce.PaymentMethodTypeMobileMoney,
			MobileMoney: &commerce.MobileMoneyParams{
				Issuer: "mtn",
				Number: "0544998605",
			},
		},
		LineItems: []commerce.OrderLineItem{
			{
				Type: commerce.LineItemTypeProduct,
				Product: &commerce.ProductLineItem{
					Name:     "Monthly Subscription",
					Type:     "digital",
					Price:    commerce.Money{Currency: "ghs", Value: 5000},
					Quantity: 1,
				},
			},
		},
		BillingDetails: commerce.BillingDetails{
			Name:        "Akua Mensah",
			Email:       "akua@example.com",
			PhoneNumber: "+233544998605",
			Address: commerce.Address{
				Name:        "Akua Mensah",
				PhoneNumber: "+233544998605",
				Line1:       "23 Adenta High Street",
				Town:        "Accra",
				Country:     "GH",
			},
		},
		PayoutSettings: &commerce.OrderPayoutSettings{
			Destination: &commerce.OrderPayoutDestination{
				FinancialAccountID: "fa_1234567890abcdef",
			},
			EnableFX: commerce.Bool(false),
		},
		ExecutePayment: commerce.Bool(true),
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Order created:", order.ID)
}
```

## Resource snippets

### Apps

```go
created, _ := client.Apps.Create(ctx, map[string]any{"name": "My App"})
current, _ := client.Apps.Lookup(ctx)
updated, _ := client.Apps.Update(ctx, map[string]any{"alias": "my-app"})
_, _, _ = created, current, updated
```

### Orders

```go
// Lookup
ord, _ := client.Orders.Lookup(ctx, "or_123")

// Pay with saved method
payResp, _ := client.Orders.Pay(ctx, commerce.OrderPayParams{
	OrderID: "or_123",
})

// Finalize
finalized, _ := client.Orders.Finalize(ctx, "or_123")
```

### Payment methods

```go
pm, _ := client.PaymentMethods.Tokenize(ctx, commerce.TokenizePaymentMethodParams{
	CustomerID: "cu_123",
	PaymentMethodData: commerce.PaymentMethodData{
		Type: commerce.PaymentMethodTypeMobileMoney,
		MobileMoney: &commerce.MobileMoneyParams{
			Issuer: "mtn",
			Number: "0544998605",
		},
	},
})

// Send verification OTP
_, _ = client.PaymentMethods.Verify(ctx, pm.ID)

// Confirm verification
verified, _ := client.PaymentMethods.ConfirmVerification(ctx, commerce.ConfirmPaymentMethodVerificationParams{
	PaymentMethodID: pm.ID,
	Token:           "000000",
})
_ = verified
```

### Chimes

```go
// Send now
sent, _ := client.Chimes.Send(ctx, commerce.SendChimeParams{
	Recipient: commerce.ChimeRecipient{
		Type: commerce.ChimeRecipientTypePhone,
		Phone: &struct{ Number string `json:"number"` }{
			Number: "+233544998605",
		},
	},
	FullMessage: "Your code is 123456",
})

// Schedule later
scheduled, _ := client.Chimes.Schedule(ctx, commerce.ScheduleChimeParams{
	Recipients:  []string{"+233544998605"},
	FullMessage: "Reminder: payment due tomorrow",
	SendAfter:   "2025-01-01T10:00:00Z",
	SenderID:    "YourBrand",
})
_ = scheduled

// Broadcast to many recipients
broadcast, _ := client.Chimes.Broadcast(ctx, commerce.BroadcastChimeParams{
	Recipients:      []string{"+233544998605", "user@example.com"},
	MessageTemplate: "Hello! Check out our new product launch.",
	ServiceName:     "MarketingCampaign",
	Sender:          "YourBrand",
})
_ = broadcast

// Look up and cancel scheduled/broadcast chimes
scheduleInfo, _ := client.Schedules.Lookup(ctx, "sch_abc123def456ghi789")
canceledSchedule, _ := client.Schedules.Cancel(ctx, "sch_abc123def456ghi789")
broadcastInfo, _ := client.Broadcasts.Lookup(ctx, "brc_abc123def456ghi789")
canceledBroadcast, _ := client.Broadcasts.Cancel(ctx, "brc_abc123def456ghi789")
_, _, _, _ = scheduleInfo, canceledSchedule, broadcastInfo, canceledBroadcast
```

### Payouts

```go
settings, _ := client.Payouts.SetDestinations(ctx, map[string]string{
	"ghs": "fa_123",
})

// Enable FX
fx, _ := client.Payouts.EnableFX(ctx)
_ = fx

// Cancel a scheduled payout
canceled, _ := client.Payouts.Cancel(ctx, "po_123")
_ = canceled
```

### Balance transactions

```go
txs, _ := client.BalanceTransactions.Page(ctx, commerce.BalanceTransactionPageParams{
	PageNumber: 1,
	PageSize:   20,
})
_ = txs
```

### Financial accounts

```go
fa, _ := client.FinancialAccounts.Create(ctx, commerce.FinancialAccountCreateParams{
	Label:     "Primary Momo",
	Type:      commerce.FinancialAccountTypeWallet,
	Reference: "REF-2024-001",
	Currency:  "ghs",
	CustomData: map[string]string{
		"merchant_id": "merch_123",
	},
	PullConfiguration: &commerce.PullPushConfig{
		Enabled: commerce.Bool(true),
		Mandate: map[string]any{},
	},
	Owner: &commerce.BankAccountOwner{
		Name: "Jane Smith",
		Address: commerce.BankAccountOwnerAddress{
			Name:    "Business Address",
			Line1:   "456 Business Road",
			City:    "Accra",
			Region:  "Greater Accra",
			Country: "Ghana",
		},
	},
	Wallet: &commerce.WalletConfig{
		Type: "mobile_money",
		MobileMoney: &struct {
			ID            string `json:"id,omitempty"`
			AccountNumber string `json:"account_number"`
			Network       string `json:"network"`
		}{
			AccountNumber: "0241234567",
			Network:       "mtn",
		},
	},
})
_ = fa

page, _ := client.FinancialAccounts.Page(ctx, commerce.PageFinancialAccountsParams{
	PageNumber: 1,
	PageSize:   50,
})
_ = page

_, _ = client.FinancialAccounts.DisablePushWithOptions(ctx, commerce.FinancialAccountDisablePushParams{
	AccountID:                "fa_123",
	UnsetAsPayoutDestination: commerce.Bool(true),
})

_, _ = client.FinancialAccounts.Disconnect(ctx, commerce.FinancialAccountDisconnectParams{
	AccountID:                "fa_123",
	UnsetAsPayoutDestination: commerce.Bool(true),
})
```

### Customers

```go
customer, _ := client.Customers.Create(ctx, commerce.CreateCustomerParams{
	Name:        "Jane Doe",
	Email:       "jane@example.com",
	PhoneNumber: "+233501234567",
})
_ = customer

existing, _ := client.Customers.Lookup(ctx, "cu_123")
_ = existing

customersPage, _ := client.Customers.Page(ctx, commerce.PageCustomersParams{PageNumber: 1, PageSize: 50})
_ = customersPage
```

### Products

```go
product, _ := client.Products.Create(ctx, commerce.CreateProductParams{
	Type: "physical",
	Name: "Premium Cotton T-Shirt",
})
_ = product

price, _ := client.Products.AddPrice(ctx, commerce.AddProductPriceParams{
	ProductID:    "prod_123",
	Amount:       commerce.ProductPriceAmount{Currency: "ghs", Value: 5000},
	SetAsDefault: true,
})
_ = price

productsPage, _ := client.Products.Page(ctx, commerce.PageProductsParams{PageNumber: 1, PageSize: 50})
_ = productsPage

published, _ := client.Products.Publish(ctx, "prod_123")
_ = published
```

### Prices

```go
price, _ := client.Prices.Create(ctx, commerce.CreatePriceParams{
	Currency: "USD",
	Amount:   1999,
	Label:    "Standard pricing",
})
_ = price

updated, _ := client.Prices.Update(ctx, commerce.UpdatePriceParams{
	PriceID: "pr_123",
	Label:   "Premium pricing",
})
_ = updated
```

### Country specs

```go
countries, _ := client.Spec.Countries(ctx)
if gh, ok := countries["gh"]; ok {
	fmt.Println("Ghana currencies:", gh.Currencies)
}
```

## Documentation

All types, methods, and fields have comprehensive godoc comments. View documentation in your editor via intellisense/hover, or use:

```bash
go doc github.com/zebodotdev/commerce-sdk-go
go doc github.com/zebodotdev/commerce-sdk-go.Client
go doc github.com/zebodotdev/commerce-sdk-go.OrdersService.Create
go doc github.com/zebodotdev/commerce-sdk-go.OrderCreateParams
```

For complete API documentation, see [https://commerce.zebo.dev/api-reference](https://commerce.zebo.dev/api-reference)

## Error handling

Errors from the API return an `*commerce.APIError` with `StatusCode`, `Type`, `Code`, `Message`, and `URL`. Check with errors.As:

```go
if err != nil {
	var apiErr *commerce.APIError
	if errors.As(err, &apiErr) {
		log.Printf("API error (%d): %s", apiErr.StatusCode, apiErr.Message)
		log.Printf("Error code: %s", apiErr.Code)
		if apiErr.URL != "" {
			log.Printf("Docs: %s", apiErr.URL)
		}
	}
}
```

## Testing

No external deps. To run the package tests (none yet), disable the repo go.work and use a temp cache:

```bash
cd sdks/go
GOWORK=off GOCACHE=$(mktemp -d) go test ./...
```
