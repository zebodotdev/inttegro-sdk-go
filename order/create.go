package order

import (
	"context"
	"fmt"

	"github.com/zebodotdev/inttegro-sdk-go/v6/checkout"
	"github.com/zebodotdev/inttegro-sdk-go/v6/customer"
	"github.com/zebodotdev/inttegro-sdk-go/v6/paymentmethod"
	"github.com/zebodotdev/inttegro-sdk-go/v6/request"
)

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
func (s *Service) Create(ctx context.Context, params CreateParams) (*Order, error) {
	return s.createWithPath(ctx, "/orders/create", params)
}

func (s *Service) createWithPath(ctx context.Context, path string, params CreateParams) (*Order, error) {
	var resp struct {
		Order       Order   `json:"order"`
		RedirectURL *string `json:"redirect_url,omitempty"`
	}
	if err := s.client.Do(ctx, "POST", path, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Order, nil
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
