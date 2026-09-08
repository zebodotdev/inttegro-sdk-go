// Package order provides order resources and operations.
package order

import (
	"context"
	"fmt"

	"github.com/zebodotdev/inttegro-sdk-go/v5/bankaccount"
	"github.com/zebodotdev/inttegro-sdk-go/v5/checkout"
	"github.com/zebodotdev/inttegro-sdk-go/v5/customer"
	"github.com/zebodotdev/inttegro-sdk-go/v5/financialaccount"
	"github.com/zebodotdev/inttegro-sdk-go/v5/internal/transport"
	"github.com/zebodotdev/inttegro-sdk-go/v5/invoice"
	"github.com/zebodotdev/inttegro-sdk-go/v5/money"
	"github.com/zebodotdev/inttegro-sdk-go/v5/payment"
	"github.com/zebodotdev/inttegro-sdk-go/v5/paymentmethod"
	"github.com/zebodotdev/inttegro-sdk-go/v5/price"
	"github.com/zebodotdev/inttegro-sdk-go/v5/product"
	"github.com/zebodotdev/inttegro-sdk-go/v5/refund"
	"github.com/zebodotdev/inttegro-sdk-go/v5/request"
	"github.com/zebodotdev/inttegro-sdk-go/v5/wallet"
)

type LineItemType string

const (
	LineItemTypeProduct  LineItemType = "product"
	LineItemTypeFee      LineItemType = "fee"
	LineItemTypeShipping LineItemType = "shipping"
)

type Status string

const (
	StatusPreparing       Status = "preparing"
	StatusRequiresPayment Status = "requires_payment"
	StatusPaid            Status = "paid"
	StatusCompleted       Status = "completed"
	StatusCanceled        Status = "canceled"
	StatusExpired         Status = "expired"
	StatusUnknown         Status = "unknown"
)

type CreatedFromResourceType string

const CreatedFromResourceTypePurchaseIntent CreatedFromResourceType = "purchase_intent"

// OrdersService provides access to order creation, payment, and lifecycle management.
//
// Orders represent purchase transactions with line items, customer information,
// and payment details. Use this service to:
//
//   - Create draft orders or charge immediately
//   - Initiate and confirm payments
//   - Finalize orders for hosted checkout
//   - Complete, cancel, or refund orders
//   - List recent orders
//
// Example:
//
//	// Create and charge an order
//	order, err := client.Orders.Create(ctx, order.CreateParams{
//	    CustomerData: &customer.Data{
//	        Name:        "Jane Doe",
//	        Email:       "jane@example.com",
//	        PhoneNumber: "+233244123456",
//	    },
//	    LineItems: []order.LineItemParams{
//	        {
//	            Type: order.LineItemTypeProduct,
//	            Product: &order.ProductLineItemParams{
//	                Type:     "digital",
//	                Name:     "Premium Plan",
//	                Quantity: 1,
//	                Price: price.InlineParams{AmountParams: money.AmountParams{
//	                    Currency: money.USD,
//	                    Value:    999,
//	                }},
//	            },
//	        },
//	    },
//	    BillingDetails: order.BillingDetails{...},
//	    ExecutePayment: inttegro.Bool(true),
//	    RequestMeta: &request.Meta{IdempotencyKey: "order_20231215_jane_001"},
//	})
type Service struct {
	client transport.Client
}

// Create creates a new order.
//
// This is the primary method for creating orders. Supports three main flows:
//
// 1. Draft order: Create without payment method, finalize later
// 2. Immediate charge: Create with payment method and execute_payment: true
// 3. Hosted checkout: Create with finalize: true, redirect customer to invoice page
//
// Parameters:
//   - CustomerData OR CustomerID: Specify customer (exactly one required)
//   - LineItems: Products, fees, and shipping charges (required)
//   - BillingDetails: Billing contact and address (optional)
//   - PaymentMethodID OR PaymentMethodData: Payment method (optional)
//   - ExecutePayment: Charge immediately (optional, requires payment method)
//   - Finalize: Generate checkout page (optional)
//   - RequestMeta.IdempotencyKey: Prevent duplicates (optional but recommended)
//
// Returns the created order with all details, including:
//   - Order ID and status
//   - Customer information
//   - Line items and totals
//   - Payment details (if payment was initiated)
//   - Invoice links (if finalized)
//
// Example (immediate charge with inline payment method):
//
//	order, err := client.Orders.Create(ctx, order.CreateParams{
//	    CustomerData: &customer.Data{
//	        Name:        "Jane Doe",
//	        Email:       "jane@example.com",
//	        PhoneNumber: "+233244123456",
//	    },
//	    PaymentMethodData: &paymentmethod.Data{
//	        Type: paymentmethod.TypeMobileMoney,
//	        MobileMoney: &paymentmethod.MobileMoneyParams{
//	            Network: "mtn",
//	            AccountNumber: "+233244123456",
//	        },
//	    },
//	    LineItems: []order.LineItemParams{...},
//	    BillingDetails: order.BillingDetails{...},
//	    ExecutePayment: inttegro.Bool(true),
//	    RequestMeta: &request.Meta{IdempotencyKey: "order_20231215_001"},
//	})
//	if err != nil {
//	    return err
//	}
//	// Check if payment requires confirmation (OTP)
//	if order.Payment != nil && order.Payment.NextAction != nil {
//	    if order.Payment.NextAction.Type == "confirm_payment" {
//	        // Prompt customer for OTP and call ConfirmPayment()
//	    }
//	}
//
// Learn more: https://studio.inttegro.com/create-order
func (s *Service) Create(ctx context.Context, params CreateParams) (*Resource, error) {
	return s.createWithPath(ctx, "/orders/create", params)
}

func (s *Service) createWithPath(ctx context.Context, path string, params CreateParams) (*Resource, error) {
	var resp struct {
		Order       Resource `json:"order"`
		RedirectURL *string  `json:"redirect_url,omitempty"`
	}
	if err := s.client.Do(ctx, "POST", path, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Order, nil
}

// Lookup retrieves an order by ID.
//
// Use this to fetch current order state, payment status, and any required
// customer actions (like OTP confirmation).
//
// Parameters:
//   - orderID: The order's unique identifier (starts with "or_")
//
// Returns the complete order object including:
//   - Current status and timestamps
//   - Payment details and status
//   - Customer and billing information
//   - Line items and totals
//   - Invoice links (if finalized)
//
// Example:
//
//	order, err := client.Orders.Lookup(ctx, "or_abc123def456")
//	if err != nil {
//	    return err
//	}
//	fmt.Printf("Order status: %s\n", order.Status)
//	if order.Payment != nil {
//	    fmt.Printf("Payment status: %s\n", order.Payment.Status)
//	}
//
// Common use cases:
//   - Polling for payment status after initiating charge
//   - Checking if OTP confirmation is still required
//   - Displaying order details to customer
//   - Syncing order state with your system
//
// Learn more: https://studio.inttegro.com/retrieve-order
func (s *Service) Lookup(ctx context.Context, orderID string) (*Resource, error) {
	var resp struct {
		Order Resource `json:"order"`
	}
	if err := s.client.Do(ctx, "POST", "/orders/lookup", LookupParams{OrderID: orderID}, &resp); err != nil {
		return nil, err
	}
	return &resp.Order, nil
}

// Update modifies mutable fields on an existing order.
func (s *Service) Update(ctx context.Context, payload any) (*Resource, error) {
	var resp struct {
		Order Resource `json:"order"`
	}
	if err := s.client.Do(ctx, "POST", "/orders/update", payload, &resp); err != nil {
		return nil, err
	}
	return &resp.Order, nil
}

// Pay initiates payment for an existing order.
//
// Use this to charge an order after creation, or to retry a failed payment.
// Provide either PaymentMethodID (for saved methods) or PaymentMethodData
// (for one-time use).
//
// Parameters:
//   - OrderID: The order to charge (required)
//   - PaymentMethodID: Saved payment method ID (optional, mutually exclusive with PaymentMethodData)
//   - PaymentMethodData: Inline payment details (optional, mutually exclusive with PaymentMethodID)
//   - PaidOutOfBand: Mark as paid externally (optional)
//
// Returns the updated order with its payment and next-action state.
//
// Example (charging with saved payment method):
//
//	order, err := client.Orders.Pay(ctx, order.PayParams{
//	    OrderID:         "or_abc123",
//	    PaymentMethodID: "pm_def456",
//	})
//	if err != nil {
//	    return err
//	}
//	if order.Payment != nil && order.Payment.NextAction != nil {
//	    // Prompt for OTP and call ConfirmPayment()
//	}
//
// Example (charging with inline payment method):
//
//	order, err := client.Orders.Pay(ctx, order.PayParams{
//	    OrderID: "or_abc123",
//	    PaymentMethodData: &paymentmethod.Data{
//	        Type: paymentmethod.TypeMobileMoney,
//	        MobileMoney: &paymentmethod.MobileMoneyParams{
//	            Network: "mtn",
//	            AccountNumber: "+233244123456",
//	        },
//	    },
//	})
//
// Learn more: https://studio.inttegro.com/orders#pay-for-an-order
func (s *Service) Pay(ctx context.Context, params PayParams) (*Resource, error) {
	var resp struct {
		Order Resource `json:"order"`
	}
	if err := s.client.Do(ctx, "POST", "/orders/pay", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Order, nil
}

// ConfirmPayment confirms a payment using an OTP token.
//
// Call this after initiating payment when confirms_use is enabled.
// The customer receives an OTP via SMS—submit it here to complete payment.
//
// Parameters:
//   - OrderID: The order awaiting confirmation (required)
//   - Token: The OTP from the customer (required, typically 4-6 digits)
//
// Returns the updated order with payment status.
//
// Example:
//
//	// Customer enters OTP: "123456"
//	order, err := client.Orders.ConfirmPayment(ctx, order.ConfirmParams{
//	    OrderID: "or_abc123",
//	    Token:   "123456",
//	})
//	if err != nil {
//	    // Invalid or expired token
//	    return err
//	}
//	// Payment confirmed, check order.Payment.Status
//
// Learn more: https://studio.inttegro.com/confirm-payment
func (s *Service) ConfirmPayment(ctx context.Context, params ConfirmParams) (*Resource, error) {
	var resp struct {
		Order Resource `json:"order"`
	}
	if err := s.client.Do(ctx, "POST", "/orders/confirm_payment", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Order, nil
}

// RequestConfirmation requests a new OTP for payment confirmation.
//
// Use this if the customer didn't receive the original OTP or if it expired.
// Triggers a new OTP to be sent to the customer's phone number.
//
// Parameters:
//   - orderID: The order needing a new confirmation token (required)
//
// Returns the updated order with the new confirmation request.
//
// Example:
//
//	// Customer says they didn't receive OTP
//	order, err := client.Orders.RequestConfirmation(ctx, "or_abc123")
//	if err != nil {
//	    return err
//	}
//	// New OTP sent, prompt customer again
func (s *Service) RequestConfirmation(ctx context.Context, orderID string) (*Resource, error) {
	return s.RequestConfirmationWithParams(ctx, RequestConfirmationParams{
		OrderID:     orderID,
		RequestMeta: stableOrderRequestMeta("request_confirmation", orderID),
	})
}

func (s *Service) RequestConfirmationWithParams(ctx context.Context, params RequestConfirmationParams) (*Resource, error) {
	var resp struct {
		Order Resource `json:"order"`
	}
	if err := s.client.Do(ctx, "POST", "/orders/request_confirmation", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Order, nil
}

// Finalize seals an order and generates hosted checkout page and invoice.
//
// Finalizing an order:
//   - Locks the order (no more modifications)
//   - Calculates final totals
//   - Generates invoice documents (PDF and web page)
//   - Creates hosted checkout URL
//
// After finalization, the order cannot be edited. Required for hosted checkout flow.
//
// Parameters:
//   - orderID: The order to finalize (required)
//
// Returns the finalized order with invoice links.
//
// Example:
//
//	order, err := client.Orders.Finalize(ctx, "or_abc123")
//	if err != nil {
//	    return err
//	}
//	// Redirect customer to order.Invoice.Format.Web.URL
//	checkoutURL := order.Invoice.Format.Web.URL
func (s *Service) Finalize(ctx context.Context, orderID string) (*Resource, error) {
	return s.FinalizeWithParams(ctx, FinalizeParams{
		OrderID:     orderID,
		RequestMeta: stableOrderRequestMeta("finalize", orderID),
	})
}

func (s *Service) FinalizeWithParams(ctx context.Context, params FinalizeParams) (*Resource, error) {
	var resp struct {
		Order Resource `json:"order"`
	}
	if err := s.client.Do(ctx, "POST", "/orders/finalize", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Order, nil
}

// SendInvoice sends the hosted invoice link for an existing order.
//
// Inttegro delivers the invoice link to every contact method available on the
// order customer.
func (s *Service) SendInvoice(ctx context.Context, params SendInvoiceParams) (*DocumentDeliveryResult, error) {
	var resp DocumentDeliveryResult
	if err := s.client.Do(ctx, "POST", "/orders/send_invoice", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SendReceipt sends the hosted receipt link for a paid order.
//
// Receipt delivery is only valid after the order has been paid.
func (s *Service) SendReceipt(ctx context.Context, params SendReceiptParams) (*DocumentDeliveryResult, error) {
	var resp DocumentDeliveryResult
	if err := s.client.Do(ctx, "POST", "/orders/send_receipt", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Complete marks an order as completed and fulfilled.
//
// Completing an order:
//   - Marks order as fulfilled
//   - Creates balance transaction (makes funds available for payout after aging)
//   - Triggers payout eligibility countdown
//
// Only valid for paid orders. Use PaidOutOfBand if payment happened outside
// Inttegro (cash, bank transfer, etc).
//
// Parameters:
//   - params.OrderID: The order to complete (required)
//   - params.PaidOutOfBand: Mark as paid externally (optional)
//
// Returns the completed order.
//
// Example:
//
//	order, err := client.Orders.Complete(ctx, order.CompleteParams{
//	    OrderID: "or_abc123",
//	})
func (s *Service) Complete(ctx context.Context, params CompleteParams) (*Resource, error) {
	var resp struct {
		Order Resource `json:"order"`
	}
	if err := s.client.Do(ctx, "POST", "/orders/complete", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Order, nil
}

// Cancel cancels an order.
//
// Canceling an order:
//   - Prevents future payment attempts
//   - Marks order as permanently closed
//
// This operation does not move funds or create a refund. Use Refunds.Create
// when funds must be returned for paid line items.
//
// Parameters:
//   - orderID: The order to cancel (required)
//
// Returns the canceled order.
//
// Example:
//
//	order, err := client.Orders.Cancel(ctx, "or_abc123")
func (s *Service) Cancel(ctx context.Context, orderID string) (*Resource, error) {
	return s.CancelWithParams(ctx, CancelParams{
		OrderID:     orderID,
		RequestMeta: stableOrderRequestMeta("cancel", orderID),
	})
}

func (s *Service) CancelWithParams(ctx context.Context, params CancelParams) (*Resource, error) {
	var resp struct {
		Order Resource `json:"order"`
	}
	if err := s.client.Do(ctx, "POST", "/orders/cancel", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Order, nil
}

func stableOrderRequestMeta(action, orderID string) *request.Meta {
	return &request.Meta{IdempotencyKey: fmt.Sprintf("orders_%s_%s", action, orderID)}
}

// Page returns a paginated list of recent orders.
//
// Use this to display order history, search for orders, or sync order state
// with your system. Results are sorted by creation date (newest first).
//
// Parameters:
//   - params.PageNumber: Page to retrieve (optional, default: 1)
//   - params.PageSize: Orders per page (optional, default: 20, max: 100)
//
// Returns a slice of orders for the requested page.
//
// Example:
//
//	orders, err := client.Orders.Page(ctx, order.PageParams{
//	    PageNumber: 1,
//	    PageSize:   50,
//	})
//	for _, order := range orders {
//	    fmt.Printf("Order %s: %s\n", order.ID, order.Status)
//	}
func (s *Service) Page(ctx context.Context, params PageParams) ([]Resource, error) {
	var resp struct {
		Page struct {
			Orders []Resource `json:"orders"`
		} `json:"page"`
	}
	if err := s.client.Do(ctx, "POST", "/orders/page", params, &resp); err != nil {
		return nil, err
	}
	return resp.Page.Orders, nil
}

// Exec wraps do for shorter tests.
func (s *Service) Exec(ctx context.Context, method, path string, body any, out any) error {
	return s.client.Do(ctx, method, path, body, out)
}

// Validate basic required fields for create.
func (p CreateParams) Validate() error {
	if len(p.LineItems) == 0 {
		return fmt.Errorf("line_items is required")
	}
	if p.BillingDetails != (BillingDetails{}) && (p.BillingDetails.Name == "" || p.BillingDetails.Email == "" || p.BillingDetails.PhoneNumber == "") {
		return fmt.Errorf("billing_details.name, email_address, and phone_number are required")
	}
	if p.CustomerData == nil && p.CustomerID == "" {
		return fmt.Errorf("either customer_data or customer_id is required")
	}
	if p.StatementDescriptor != "" && p.StatementDescriptorPrefix != "" {
		return fmt.Errorf("statement_descriptor and statement_descriptor_prefix are mutually exclusive")
	}
	return nil
}

// BillingDetails captures billing contact information for an order.
//
// Required for all orders. This information appears on invoices and
// is used for payment authorization and dispute resolution.
//
// The billing phone number and email are used for payment notifications
// and confirmation codes.
type BillingDetails struct {
	// Name is the billing contact's full name (required).
	// Typically matches the customer name.
	Name string `json:"name"`

	// Email is the billing email address (required).
	// Used for invoice delivery and payment receipts.
	// Validated as RFC 5322 email format.
	Email string `json:"email_address"`

	// PhoneNumber is the billing contact phone with country code (required).
	// Used for payment confirmations and delivery of OTP codes.
	// Example: "+233244123456"
	PhoneNumber string `json:"phone_number"`

	// Address is the billing postal address (required).
	// Must be a complete, valid address.
	Address customer.Address `json:"address"`
}

// Shipping holds shipping information for physical goods.
//
// Only required when selling physical products that need delivery.
// Digital goods and services don't require shipping information.
type Shipping struct {
	// Address is the delivery postal address (required for physical goods).
	// Must be a complete, deliverable address.
	Address customer.Address `json:"address"`
}

// ProductLineItemParams represents a product supplied in an order request.
//
// Products are goods or services sold to customers. Each product has a unit
// price, quantity, and optional descriptive information.
//
// Example:
//
//	item := &order.ProductLineItemParams{
//	    Type:     "physical",
//	    Name:     "Wireless Headphones",
//	    About:    "Bluetooth 5.0, 30-hour battery",
//	    Quantity: 2,
//	    Price:    price.InlineParams{AmountParams: money.AmountParams{Currency: money.USD, Value: 7999}}, // $79.99 each
//	    Reference: "SKU-12345",
//	}
type ProductLineItemParams struct {
	// Type indicates whether the product is physical or digital (required).
	// Values: "physical" or "digital"
	// Physical products require shipping address.
	// Digital products can be delivered electronically.
	Type product.Type `json:"type"`

	// Name is the product name (required).
	// Displayed on invoices and checkout pages.
	// Maximum 255 characters.
	Name string `json:"name"`

	// About is a short product description (optional).
	// Additional details about the product.
	// Maximum 1000 characters.
	About string `json:"about,omitempty"`

	// Quantity is the number of units (required).
	// Must be positive. Total line item amount = Price * Quantity.
	Quantity int64 `json:"quantity"`

	// Price is the unit price per item (required).
	// In minor units. The line item total is Price.Value * Quantity.
	Price price.InlineParams `json:"price"`

	// Reference is your internal product identifier (optional).
	// Link to your inventory system's SKU or product ID.
	// Maximum 255 characters.
	Reference string `json:"reference,omitempty"`

	// TaxCode specifies the tax treatment (optional).
	// Used for tax calculation when tax integration is enabled.
	// Format depends on your tax provider.
	TaxCode string `json:"tax_code,omitempty"`

	// CustomData holds arbitrary key-value custom data (optional).
	// Both keys and values must be strings.
	// Maximum 25KB when serialized.
	// Learn more: https://studio.inttegro.com/custom-data
	CustomData map[string]string `json:"custom_data,omitempty"`
}

// ProductLineItem is a product returned in an order.
type ProductLineItem struct {
	ID         string            `json:"id"`
	Type       product.Type      `json:"type"`
	Name       string            `json:"name"`
	About      string            `json:"about,omitempty"`
	Quantity   int64             `json:"quantity"`
	Price      price.Inline      `json:"price"`
	Reference  string            `json:"reference,omitempty"`
	TaxCode    string            `json:"tax_code,omitempty"`
	CustomData map[string]string `json:"custom_data,omitempty"`
}

// FeeLineItemParams represents an additional charge supplied in a request.
//
// Fees are one-time charges added to the order subtotal. Unlike products,
// fees don't have quantities—they're always a fixed amount.
//
// Example:
//
//	fee := &order.FeeLineItemParams{
//	    Label:       "Service Fee",
//	    Description: "Platform usage fee",
//	    Amount:      money.AmountParams{Currency: money.USD, Value: 299}, // $2.99
//	}
type FeeLineItemParams struct {
	// Label is the fee name (optional but recommended).
	// Displayed on invoices. Example: "Service Fee", "Processing Fee"
	// Maximum 255 characters.
	Label string `json:"label,omitempty"`

	// Description explains the fee (optional).
	// Additional context about why this fee is charged.
	// Maximum 1000 characters.
	Description string `json:"description,omitempty"`

	// TaxCode specifies the tax treatment (optional).
	// Used for tax calculation when tax integration is enabled.
	TaxCode string `json:"tax_code,omitempty"`

	// CustomData holds arbitrary key-value custom data (optional).
	// Both keys and values must be strings.
	// Maximum 25KB when serialized.
	CustomData map[string]string `json:"custom_data,omitempty"`

	// Amount is the total fee charge (required).
	// In minor units. Not multiplied by any quantity.
	Amount money.AmountParams `json:"amount"`
}

// FeeLineItem is an additional charge returned in an order.
type FeeLineItem struct {
	ID          string            `json:"id"`
	Label       string            `json:"label"`
	Description string            `json:"description,omitempty"`
	TaxCode     string            `json:"tax_code,omitempty"`
	CustomData  map[string]string `json:"custom_data,omitempty"`
	Amount      money.Amount      `json:"amount"`
}

// ShippingLineItemParams represents a delivery charge supplied in a request.
//
// Only needed when selling physical products. Automatically omitted for
// orders containing only digital products.
//
// Example:
//
//	shipping := &order.ShippingLineItemParams{
//	    Fee: money.AmountParams{Currency: money.USD, Value: 500}, // $5.00
//	}
type ShippingLineItemParams struct {
	// Fee is the total shipping charge (required).
	// In minor units. Not multiplied by any quantity.
	Fee money.AmountParams `json:"fee"`

	// TaxCode specifies the tax treatment (optional).
	// Used for tax calculation when tax integration is enabled.
	TaxCode string `json:"tax_code,omitempty"`

	// CustomData holds arbitrary key-value custom data (optional).
	// Both keys and values must be strings.
	// Maximum 25KB when serialized.
	CustomData map[string]string `json:"custom_data,omitempty"`
}

// ShippingLineItem is a delivery charge returned in an order.
type ShippingLineItem struct {
	ID         string            `json:"id"`
	Fee        money.Amount      `json:"fee"`
	TaxCode    string            `json:"tax_code,omitempty"`
	CustomData map[string]string `json:"custom_data,omitempty"`
}

// OrderLineItemParams is a discriminated union supplied in an order request.
//
// Each order line item is one of three types: product, fee, or shipping.
// Set Type and the corresponding field (Product, Fee, or Shipping).
// Leave the other fields nil.
//
// Example (product):
//
//	lineItem := order.LineItemParams{
//	    Type: order.LineItemTypeProduct,
//	    Product: &order.ProductLineItemParams{
//	        Type:     "digital",
//	        Name:     "Premium Subscription",
//	        Quantity: 1,
//	        Price:    price.InlineParams{AmountParams: money.AmountParams{Currency: money.USD, Value: 999}},
//	    },
//	}
//
// Example (fee):
//
//	lineItem := order.LineItemParams{
//	    Type: order.LineItemTypeFee,
//	    Fee: &order.FeeLineItemParams{
//	        Label:  "Platform Fee",
//	        Amount: money.AmountParams{Currency: money.USD, Value: 299},
//	    },
//	}
type LineItemParams struct {
	// Type specifies which variant is active (required).
	Type LineItemType `json:"type"`

	// Product is populated when Type is LineItemTypeProduct.
	// Nil for other types.
	Product *ProductLineItemParams `json:"product,omitempty"`

	// Fee is populated when Type is LineItemTypeFee.
	// Nil for other types.
	Fee *FeeLineItemParams `json:"fee,omitempty"`

	// Shipping is populated when Type is LineItemTypeShipping.
	// Nil for other types.
	Shipping *ShippingLineItemParams `json:"shipping,omitempty"`
}

// OrderLineItem is a discriminated union returned by the API.
type LineItem struct {
	Type     LineItemType      `json:"type"`
	Product  *ProductLineItem  `json:"product,omitempty"`
	Fee      *FeeLineItem      `json:"fee,omitempty"`
	Shipping *ShippingLineItem `json:"shipping,omitempty"`
}

// OrderPayoutSettings overrides payout configuration for a single order.
type PayoutSettings struct {
	// Destination specifies where payout funds should be sent (optional).
	Destination *PayoutDestination `json:"destination,omitempty"`

	// EnableFX controls foreign exchange conversion for the payout (optional).
	EnableFX *bool `json:"enable_fx,omitempty"`
}

// OrderPayoutDestination sets the payout destination for the order.
type PayoutDestination struct {
	// FinancialAccountID references an existing financial account (optional).
	FinancialAccountID string `json:"financial_account_id,omitempty"`

	// FinancialAccountData provides inline financial account details (optional).
	FinancialAccountData *PayoutFinancialAccount `json:"financial_account_data,omitempty"`
}

// OrderPayoutFinancialAccount defines inline payout destination details.
type PayoutFinancialAccount struct {
	// Type specifies the account type (wallet, bank_account, dosh_account).
	Type financialaccount.Type `json:"type"`

	// Wallet contains mobile money details when Type is "wallet".
	Wallet *wallet.Config `json:"wallet,omitempty"`

	// BankAccount contains bank account details when Type is "bank_account".
	BankAccount *bankaccount.Config `json:"bank_account,omitempty"`

	// DoshAccount contains Dosh wallet details when Type is "dosh_account".
	DoshAccount map[string]any `json:"dosh_account,omitempty"`
}

// OrderCreateParams contains all parameters for creating an order.
//
// Orders represent a purchase transaction with line items, customer info,
// and payment details. The order creation flow is flexible:
//
// 1. Create draft order (no payment method) → finalize → customer pays from hosted page
// 2. Create order with payment method → execute payment immediately
// 3. Create order with checkout settings → finalize → redirect customer to hosted page
//
// Customer specification (exactly one required):
//   - CustomerData: for new customers (API creates customer record)
//   - CustomerID: for existing customers
//
// Payment method specification (optional):
//   - PaymentMethodID: reference saved payment method
//   - PaymentMethodData: inline payment method for one-time use
//
// Example (new customer, immediate payment):
//
//	params := order.CreateParams{
//	    CustomerData: &customer.Data{
//	        Name:        "Jane Doe",
//	        Email:       "jane@example.com",
//	        PhoneNumber: "+233244123456",
//	    },
//	    PaymentMethodData: &paymentmethod.Data{
//	        Type: paymentmethod.TypeMobileMoney,
//	        MobileMoney: &paymentmethod.MobileMoneyParams{
//	            Network: "mtn",
//	            AccountNumber: "+233244123456",
//	        },
//	    },
//	    LineItems: []order.LineItemParams{
//	        {
//	            Type: order.LineItemTypeProduct,
//	            Product: &order.ProductLineItemParams{
//	                Type:     "digital",
//	                Name:     "Premium Plan",
//	                Quantity: 1,
//	                Price:    price.InlineParams{AmountParams: money.AmountParams{Currency: money.GHS, Value: 10000}},
//	            },
//	        },
//	    },
//	    BillingDetails: order.BillingDetails{
//	        Name:        "Jane Doe",
//	        Email:       "jane@example.com",
//	        PhoneNumber: "+233244123456",
//	        Address: customer.Address{
//	            Name:        "Jane Doe",
//	            PhoneNumber: "+233244123456",
//	            Line1:       "123 Main St",
//	            Town:        "Accra",
//	            Country:     "GH",
//	        },
//	    },
//	    ExecutePayment: inttegro.Bool(true),
//	    RequestMeta: &request.Meta{IdempotencyKey: "order_20231215_jane_001"},
//	}
type CreateParams struct {
	// RequestMeta carries per-request controls such as idempotency.
	// Prefer RequestMeta.IdempotencyKey over the legacy top-level IdempotencyKey.
	RequestMeta *request.Meta `json:"request_meta,omitempty"`

	// CustomerData provides inline customer information for new customers (optional).
	// Use this when the customer doesn't have a Inttegro customer ID yet.
	// The API creates a customer record and returns it in the response.
	// Mutually exclusive with CustomerID (provide exactly one).
	CustomerData *customer.Data `json:"customer_data,omitempty"`

	// CustomerID references an existing customer by ID (optional).
	// Use this for repeat customers who already have a Inttegro customer record.
	// Mutually exclusive with CustomerData (provide exactly one).
	CustomerID string `json:"customer_id,omitempty"`

	// PaymentMethodID references a saved payment method by ID (optional).
	// Used for charging repeat customers with tokenized payment methods.
	// Mutually exclusive with PaymentMethodData.
	// If provided, you can set ExecutePayment: true to charge immediately.
	PaymentMethodID string `json:"payment_method_id,omitempty"`

	// PaymentMethodData provides inline payment details for one-time use (optional).
	// Use this when you don't want to save the payment method.
	// Mutually exclusive with PaymentMethodID.
	// If provided, you can set ExecutePayment: true to charge immediately.
	PaymentMethodData *paymentmethod.Data `json:"payment_method_data,omitempty"`

	// StatementDescriptor appears on the customer's payment statement (optional).
	// Maximum 22 characters. Only letters, numbers, and spaces allowed.
	// Example: "ACME INC SUBSCRIPTION"
	// If omitted, uses your business name from settings.
	StatementDescriptor string `json:"statement_descriptor,omitempty"`

	// StatementDescriptorPrefix builds a descriptor from a 2-10 character prefix and generated order ID.
	// The API formats it as prefix*order_id and truncates the order ID to fit.
	// Mutually exclusive with StatementDescriptor.
	StatementDescriptorPrefix string `json:"statement_descriptor_prefix,omitempty"`

	// ExecutePayment triggers immediate payment attempt (optional, default: false).
	// Only valid when PaymentMethodID or PaymentMethodData is provided.
	// If false (default), order is created but payment must be initiated separately.
	// If true, attempts to charge the payment method immediately.
	ExecutePayment *bool `json:"execute_payment,omitempty"`

	// Finalize seals the order and generates checkout page (optional, default: false).
	// If true, order is finalized and checkout URL is returned in response.
	// Finalized orders cannot be modified—line items and amounts are locked.
	// Required for hosted checkout flow.
	Finalize *bool `json:"finalize,omitempty"`

	// CheckoutSettings configures hosted checkout page redirects (optional).
	// Only relevant when Finalize is true.
	// Specifies where customers are redirected after payment or cancellation.
	CheckoutSettings *checkout.Settings `json:"checkout_settings,omitempty"`

	// PayoutSettings overrides payout configuration for this order (optional).
	PayoutSettings *PayoutSettings `json:"payout_settings,omitempty"`

	// Number is a custom order number for your records (optional).
	// If omitted, Inttegro generates a unique order number automatically.
	// Maximum 255 characters. Must be unique across your orders.
	// Example: "ORD-2023-00123"
	Number string `json:"number,omitempty"`

	// LineItems is the list of products, fees, and shipping charges (required).
	// Must have at least one line item. Orders without line items are invalid.
	// The order total is calculated from all line items.
	LineItems []LineItemParams `json:"line_items"`

	// CustomData holds arbitrary key-value metadata for the order (optional).
	// Both keys and values must be strings.
	CustomData map[string]string `json:"custom_data,omitempty"`

	// BillingDetails optionally supplies billing contact and address details.
	// Hosted checkout can collect payment-specific details when this is zero.
	BillingDetails BillingDetails `json:"billing_details,omitzero"`

	// Shipping provides delivery address for physical goods (optional).
	// Required only when line items include physical products.
	// Omit for digital-only orders.
	Shipping *Shipping `json:"shipping,omitempty"`
}

// OrderLookupParams specifies which order to retrieve.
type LookupParams struct {
	// OrderID is the unique order identifier (required).
	// Starts with "or_". Example: "or_abc123def456"
	OrderID string `json:"order_id"`
}

// OrderPayParams initiates payment for an existing order.
//
// Use this to charge an order after creation, or to retry payment
// on an order with a failed payment attempt.
//
// Provide either PaymentMethodID (for saved payment methods) or
// PaymentMethodData (for one-time payment methods).
//
// Example (with saved payment method):
//
//	params := order.PayParams{
//	    OrderID:         createdOrder.ID,
//	    PaymentMethodID: "pm_abc123",
//	}
//	response, err := client.Orders.Pay(ctx, params)
type PayParams struct {
	// OrderID is the order to charge (required).
	OrderID string `json:"order_id"`

	// RequestMeta carries per-request controls such as idempotency.
	RequestMeta *request.Meta `json:"request_meta,omitempty"`

	// PaymentMethodID references a saved payment method (optional).
	// Mutually exclusive with PaymentMethodData.
	// Use this for repeat customers with tokenized payment methods.
	PaymentMethodID string `json:"payment_method_id,omitempty"`

	// PaymentMethodData provides inline payment details (optional).
	// Mutually exclusive with PaymentMethodID.
	// Use this for one-time payments without saving the method.
	PaymentMethodData *paymentmethod.Data `json:"payment_method_data,omitempty"`

	// PaidOutOfBand marks payment as completed outside Inttegro (optional).
	// Set to true if customer paid via cash, bank transfer, or other method.
	// The order is marked paid without actually charging the payment method.
	// Use carefully—this bypasses payment processing entirely.
	PaidOutOfBand *bool `json:"paid_out_of_band,omitempty"`
}

// OrderConfirmParams confirms a payment using a verification token.
//
// After initiating payment on an order with confirms_use enabled,
// the customer receives an OTP. Use this to submit their OTP and
// complete the payment.
//
// Example:
//
//	params := order.ConfirmParams{
//	    OrderID: createdOrder.ID,
//	    Token:   "123456", // OTP from customer
//	}
//	order, err := client.Orders.ConfirmPayment(ctx, params)
type ConfirmParams struct {
	// OrderID is the order awaiting confirmation (required).
	OrderID string `json:"order_id"`

	// RequestMeta carries per-request controls such as idempotency.
	RequestMeta *request.Meta `json:"request_meta,omitempty"`

	// Token is the OTP or verification code (required).
	// Typically 4-6 digits sent to the customer's phone via SMS.
	Token string `json:"token"`
}

// OrderRequestConfirmationParams requests a new OTP for payment confirmation.
//
// Use this if the customer didn't receive the original OTP or if it expired.
// Triggers a new OTP to be sent to the customer's phone number.
type RequestConfirmationParams struct {
	// OrderID is the order needing a new confirmation token (required).
	OrderID string `json:"order_id"`

	// RequestMeta carries per-request controls such as idempotency.
	RequestMeta *request.Meta `json:"request_meta,omitempty"`
}

// OrderFinalizeParams seals an order and generates checkout assets.
//
// Finalizing an order:
// - Locks the order (no more modifications)
// - Calculates final totals
// - Generates hosted checkout page
// - Creates invoice documents
//
// After finalization, the order cannot be edited. Use this when you're
// ready for the customer to pay.
type FinalizeParams struct {
	// OrderID is the order to finalize (required).
	OrderID string `json:"order_id"`

	// RequestMeta carries per-request controls such as idempotency.
	RequestMeta *request.Meta `json:"request_meta,omitempty"`
}

// OrderSendInvoiceParams sends an invoice link for an order.
type SendInvoiceParams struct {
	// OrderID is the order whose invoice should be sent (required).
	OrderID string `json:"order_id"`

	// RequestMeta carries per-request controls such as idempotency.
	RequestMeta *request.Meta `json:"request_meta,omitempty"`
}

// OrderSendReceiptParams sends a receipt link for a paid order.
type SendReceiptParams struct {
	// OrderID is the paid order whose receipt should be sent (required).
	OrderID string `json:"order_id"`

	// RequestMeta carries per-request controls such as idempotency.
	RequestMeta *request.Meta `json:"request_meta,omitempty"`
}

// OrderCompleteParams marks an order as completed.
//
// Completing an order:
// - Marks order as fulfilled
// - Triggers payout eligibility (after aging period)
// - Creates balance transaction
//
// Only valid for paid orders. Use PaidOutOfBand if payment happened
// outside the Inttegro platform.
type CompleteParams struct {
	// OrderID is the order to complete (required).
	OrderID string `json:"order_id"`

	// PaidOutOfBand indicates payment happened externally (optional).
	// If true, marks order as paid without charging the payment method.
	// Use for cash, bank transfer, or other offline payment methods.
	PaidOutOfBand *bool `json:"paid_out_of_band,omitempty"`
}

// OrderCancelParams cancels an order.
//
// Canceling an order:
// - Prevents any future payment attempts
// - Marks order as permanently closed
//
// This operation does not move funds or create a refund.
type CancelParams struct {
	// OrderID is the order to cancel (required).
	OrderID string `json:"order_id"`

	// RequestMeta carries per-request controls such as idempotency.
	RequestMeta *request.Meta `json:"request_meta,omitempty"`
}

// OrderPageParams specifies pagination for listing orders.
//
// Use this to fetch recent orders with pagination support.
//
// Example:
//
//	params := order.PageParams{
//	    PageNumber: 1,
//	    PageSize:   50,
//	}
//	orders, err := client.Orders.Page(ctx, params)
type PageParams struct {
	// PageNumber is the page to retrieve (optional, default: 1).
	// Pages are 1-indexed. First page is 1, not 0.
	PageNumber int `json:"page_number,omitempty"`

	// PageSize is the number of orders per page (optional, default: 20).
	// Maximum 100. Minimum 1.
	PageSize int `json:"page_size,omitempty"`
}

// LineItemGroup contains line items grouped by type with totals.
//
// The API returns this grouped structure in order responses to make
// it easier to display cart breakdowns.
type LineItemGroup struct {
	// LineItems is the list of cart items.
	LineItems []LineItem `json:"line_items"`

	// Total is the sum of all line items.
	// This is the order amount before any fees or discounts.
	Total money.Amount `json:"total"`
}

// Order represents a complete order object.
//
// Orders are the central resource in Inttegro, representing a purchase
// transaction from cart to fulfillment. Orders go through several states:
//
// 1. preparing: Created but not finalized
// 2. requires_payment: Finalized and ready for payment
// 3. paid: Payment succeeded
// 4. completed: Fulfilled and settled
// 5. canceled: Permanently canceled
// 6. expired: Payment window elapsed
type Resource struct {
	// ID is the unique order identifier (read-only).
	// Starts with "or_". Example: "or_abc123def456"
	ID string `json:"id"`

	// Status is the order's current state (read-only).
	// Values: "preparing", "requires_payment", "paid", "completed", "canceled", "expired", "unknown"
	Status Status `json:"status"`

	// Number is the human-readable order number.
	// Auto-generated if not provided during creation.
	// Example: "ORD-2023-00123"
	Number string `json:"number,omitempty"`

	// CustomerID is the owning customer's ID.
	CustomerID string `json:"customer_id,omitempty"`

	// Customer contains full customer details.
	// Only populated in responses when customer exists.
	Customer *customer.Data `json:"customer,omitempty"`

	// BillingDetails holds billing contact and address.
	BillingDetails *BillingDetails `json:"billing_details,omitempty"`

	// Shipping holds delivery address for physical goods.
	// Nil for digital-only orders.
	Shipping *Shipping `json:"shipping,omitempty"`

	// LineItems is the list of products, fees, and shipping charges.
	LineItems []LineItem `json:"line_items,omitempty"`

	// LineItemGroup contains line items grouped with totals.
	// Useful for displaying cart breakdowns.
	LineItemGroup *LineItemGroup `json:"line_item_group,omitempty"`

	// Payment contains payment details and status.
	// Nil if order hasn't been charged yet.
	Payment *payment.Resource `json:"payment,omitempty"`

	// PaymentStatus is a summary of payment state (read-only).
	// Values: "unpaid", "requires_action", "processing", "paid", "failed"
	PaymentStatus string `json:"payment_status,omitempty"`

	// PaymentMethodID is the attached payment method's ID.
	// May be set but not yet charged.
	PaymentMethodID string `json:"payment_method_id,omitempty"`

	// StatementDescriptor appears on payment statements.
	StatementDescriptor string `json:"statement_descriptor,omitempty"`

	// CheckoutSettings contains hosted checkout configuration.
	CheckoutSettings *checkout.Settings `json:"checkout_settings,omitempty"`

	// InitiatedAt is when the order was created (ISO 8601, read-only).
	InitiatedAt string `json:"initiated_at,omitempty"`

	// SealedAt is when the order was finalized (ISO 8601, read-only).
	// Nil if not yet finalized.
	SealedAt *string `json:"sealed_at,omitempty"`

	// CompletedAt is when the order was marked complete (ISO 8601, read-only).
	// Nil if not yet completed.
	CompletedAt *string `json:"completed_at,omitempty"`

	// ExpiresAt is when the order expires if unpaid (ISO 8601, read-only).
	// Nil for completed orders.
	ExpiresAt *string `json:"expires_at,omitempty"`

	// CreatedAt is the creation timestamp (ISO 8601, read-only).
	CreatedAt *string `json:"created_at,omitempty"`

	// UpdatedAt is the last modification timestamp (ISO 8601, read-only).
	UpdatedAt *string `json:"updated_at,omitempty"`

	// PaidAt is when payment was confirmed (ISO 8601, read-only).
	// Nil if not yet paid.
	PaidAt *string `json:"paid_at,omitempty"`

	// CancelledAt is when the order was canceled (ISO 8601, read-only).
	// Nil if not canceled.
	CancelledAt *string `json:"cancelled_at,omitempty"`

	// Invoice contains invoice document links and delivery status.
	// Nil if order not finalized.
	Invoice *invoice.Resource `json:"invoice,omitempty"`

	// Refunds contains every refund issued for this order, newest first.
	// It is omitted when no refunds exist.
	Refunds []refund.Resource `json:"refunds,omitempty"`
}

// NewService constructs the resource service used by inttegro.Client.
func NewService(client transport.Client) *Service { return &Service{client: client} }
