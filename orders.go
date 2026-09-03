package inttegro

import (
	"context"
	"fmt"
)

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
//	order, err := client.Orders.Create(ctx, inttegro.OrderCreateParams{
//	    CustomerData: &inttegro.CustomerData{
//	        Name:        "Jane Doe",
//	        Email:       "jane@example.com",
//	        PhoneNumber: "+233244123456",
//	    },
//	    LineItems: []inttegro.OrderLineItem{
//	        {
//	            Type: inttegro.LineItemTypeProduct,
//	            Product: &inttegro.ProductLineItem{
//	                Type:     "digital",
//	                Name:     "Premium Plan",
//	                Quantity: 1,
//	                Price:    inttegro.Money{Currency: "usd", Value: 999},
//	            },
//	        },
//	    },
//	    BillingDetails: inttegro.BillingDetails{...},
//	    ExecutePayment: inttegro.Bool(true),
//	    RequestMeta: &inttegro.RequestMeta{IdempotencyKey: "order_20231215_jane_001"},
//	})
type OrdersService struct {
	client *Client
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
//   - BillingDetails: Billing contact and address (required)
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
//	order, err := client.Orders.Create(ctx, inttegro.OrderCreateParams{
//	    CustomerData: &inttegro.CustomerData{
//	        Name:        "Jane Doe",
//	        Email:       "jane@example.com",
//	        PhoneNumber: "+233244123456",
//	    },
//	    PaymentMethodData: &inttegro.PaymentMethodData{
//	        Type: inttegro.PaymentMethodTypeMobileMoney,
//	        MobileMoney: &inttegro.MobileMoneyParams{
//	            Network: "mtn",
//	            AccountNumber: "+233244123456",
//	        },
//	    },
//	    LineItems: []inttegro.OrderLineItem{...},
//	    BillingDetails: inttegro.BillingDetails{...},
//	    ExecutePayment: inttegro.Bool(true),
//	    RequestMeta: &inttegro.RequestMeta{IdempotencyKey: "order_20231215_001"},
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
func (s *OrdersService) Create(ctx context.Context, params OrderCreateParams) (*Order, error) {
	return s.createWithPath(ctx, "/orders/create", params)
}

// New creates an order through the legacy /orders/new compatibility endpoint.
func (s *OrdersService) New(ctx context.Context, params OrderCreateParams) (*Order, error) {
	return s.createWithPath(ctx, "/orders/new", params)
}

func (s *OrdersService) createWithPath(ctx context.Context, path string, params OrderCreateParams) (*Order, error) {
	var resp struct {
		Order       Order   `json:"order"`
		RedirectURL *string `json:"redirect_url,omitempty"`
	}
	if err := s.client.do(ctx, "POST", path, params, &resp); err != nil {
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
func (s *OrdersService) Lookup(ctx context.Context, orderID string) (*Order, error) {
	var resp struct {
		Order Order `json:"order"`
	}
	if err := s.client.do(ctx, "POST", "/orders/lookup", OrderLookupParams{OrderID: orderID}, &resp); err != nil {
		return nil, err
	}
	return &resp.Order, nil
}

// Update modifies mutable fields on an existing order.
func (s *OrdersService) Update(ctx context.Context, payload any) (*Order, error) {
	var resp struct {
		Order Order `json:"order"`
	}
	if err := s.client.do(ctx, "POST", "/orders/update", payload, &resp); err != nil {
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
//	order, err := client.Orders.Pay(ctx, inttegro.OrderPayParams{
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
//	order, err := client.Orders.Pay(ctx, inttegro.OrderPayParams{
//	    OrderID: "or_abc123",
//	    PaymentMethodData: &inttegro.PaymentMethodData{
//	        Type: inttegro.PaymentMethodTypeMobileMoney,
//	        MobileMoney: &inttegro.MobileMoneyParams{
//	            Network: "mtn",
//	            AccountNumber: "+233244123456",
//	        },
//	    },
//	})
//
// Learn more: https://studio.inttegro.com/orders#pay-for-an-order
func (s *OrdersService) Pay(ctx context.Context, params OrderPayParams) (*Order, error) {
	var resp struct {
		Order Order `json:"order"`
	}
	if err := s.client.do(ctx, "POST", "/orders/pay", params, &resp); err != nil {
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
//	order, err := client.Orders.ConfirmPayment(ctx, inttegro.OrderConfirmParams{
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
func (s *OrdersService) ConfirmPayment(ctx context.Context, params OrderConfirmParams) (*Order, error) {
	var resp struct {
		Order Order `json:"order"`
	}
	if err := s.client.do(ctx, "POST", "/orders/confirm_payment", params, &resp); err != nil {
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
func (s *OrdersService) RequestConfirmation(ctx context.Context, orderID string) (*Order, error) {
	return s.RequestConfirmationWithParams(ctx, OrderRequestConfirmationParams{
		OrderID:     orderID,
		RequestMeta: stableOrderRequestMeta("request_confirmation", orderID),
	})
}

func (s *OrdersService) RequestConfirmationWithParams(ctx context.Context, params OrderRequestConfirmationParams) (*Order, error) {
	var resp struct {
		Order Order `json:"order"`
	}
	if err := s.client.do(ctx, "POST", "/orders/request_confirmation", params, &resp); err != nil {
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
func (s *OrdersService) Finalize(ctx context.Context, orderID string) (*Order, error) {
	return s.FinalizeWithParams(ctx, OrderFinalizeParams{
		OrderID:     orderID,
		RequestMeta: stableOrderRequestMeta("finalize", orderID),
	})
}

func (s *OrdersService) FinalizeWithParams(ctx context.Context, params OrderFinalizeParams) (*Order, error) {
	var resp struct {
		Order Order `json:"order"`
	}
	if err := s.client.do(ctx, "POST", "/orders/finalize", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Order, nil
}

// SendInvoice sends the hosted invoice link for an existing order.
//
// Inttegro delivers the invoice link to every contact method available on the
// order customer.
func (s *OrdersService) SendInvoice(ctx context.Context, params OrderSendInvoiceParams) (*OrderDocumentDeliveryResult, error) {
	var resp OrderDocumentDeliveryResult
	if err := s.client.do(ctx, "POST", "/orders/send_invoice", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SendReceipt sends the hosted receipt link for a paid order.
//
// Receipt delivery is only valid after the order has been paid.
func (s *OrdersService) SendReceipt(ctx context.Context, params OrderSendReceiptParams) (*OrderDocumentDeliveryResult, error) {
	var resp OrderDocumentDeliveryResult
	if err := s.client.do(ctx, "POST", "/orders/send_receipt", params, &resp); err != nil {
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
//	order, err := client.Orders.Complete(ctx, inttegro.OrderCompleteParams{
//	    OrderID: "or_abc123",
//	})
func (s *OrdersService) Complete(ctx context.Context, params OrderCompleteParams) (*Order, error) {
	var resp struct {
		Order Order `json:"order"`
	}
	if err := s.client.do(ctx, "POST", "/orders/complete", params, &resp); err != nil {
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
func (s *OrdersService) Cancel(ctx context.Context, orderID string) (*Order, error) {
	return s.CancelWithParams(ctx, OrderCancelParams{
		OrderID:     orderID,
		RequestMeta: stableOrderRequestMeta("cancel", orderID),
	})
}

func (s *OrdersService) CancelWithParams(ctx context.Context, params OrderCancelParams) (*Order, error) {
	var resp struct {
		Order Order `json:"order"`
	}
	if err := s.client.do(ctx, "POST", "/orders/cancel", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Order, nil
}

// Refund creates a refund through the compatibility /orders/refund route.
// It accepts and returns the same contract as Refunds.Create.
//
// Deprecated: use Refunds.Create for new integrations.
func (s *OrdersService) Refund(
	ctx context.Context,
	request CreateRefundRequest,
) (*Refund, error) {
	return createRefund(ctx, s.client, "/orders/refund", request)
}

func stableOrderRequestMeta(action, orderID string) *RequestMeta {
	return &RequestMeta{IdempotencyKey: fmt.Sprintf("orders_%s_%s", action, orderID)}
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
//	orders, err := client.Orders.Page(ctx, inttegro.OrderPageParams{
//	    PageNumber: 1,
//	    PageSize:   50,
//	})
//	for _, order := range orders {
//	    fmt.Printf("Order %s: %s\n", order.ID, order.Status)
//	}
func (s *OrdersService) Page(ctx context.Context, params OrderPageParams) ([]Order, error) {
	var resp struct {
		Page struct {
			Orders []Order `json:"orders"`
		} `json:"page"`
	}
	if err := s.client.do(ctx, "POST", "/orders/page", params, &resp); err != nil {
		return nil, err
	}
	return resp.Page.Orders, nil
}

// Exec wraps do for shorter tests.
func (s *OrdersService) Exec(ctx context.Context, method, path string, body any, out any) error {
	return s.client.do(ctx, method, path, body, out)
}

// Validate basic required fields for create.
func (p OrderCreateParams) Validate() error {
	if len(p.LineItems) == 0 {
		return fmt.Errorf("line_items is required")
	}
	if p.BillingDetails.Name == "" || p.BillingDetails.Email == "" || p.BillingDetails.PhoneNumber == "" {
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
