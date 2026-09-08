// Package balancetransaction provides balancetransaction resources and operations.
package balancetransaction

import (
	"context"

	"github.com/zebodotdev/inttegro-sdk-go/v5/internal/transport"
	"github.com/zebodotdev/inttegro-sdk-go/v5/money"
	"github.com/zebodotdev/inttegro-sdk-go/v5/payout"
)

// BalanceTransactionsService provides access to balance transaction history.
//
// Balance transactions are merchant balance entries caused by payments or refunds.
// Type identifies the semantic source, and the matching PaymentID or RefundID
// provides the strong source reference.
//
// Use this service to:
//   - View available and pending balance
//   - Track balance transaction aging
//   - Reconcile payouts with source payments and refunds
//
// Example:
//
//	transactions, err := client.BalanceTransactions.Page(ctx, balancetransaction.PageParams{
//	    PageSize: 100,
//	})
//	for _, tx := range transactions {
//	    sourceID, _ := tx.SourceID()
//	    fmt.Printf("%s (%s %s): %s %d\n",
//	        tx.ID, tx.Type, sourceID, tx.Amount.Currency, tx.Amount.Value)
//	}
type Service struct {
	client transport.Client
}

// Lookup retrieves a balance transaction by ID.
func (s *Service) Lookup(ctx context.Context, transactionID string) (*Resource, error) {
	var resp struct {
		Transaction Resource `json:"transaction"`
	}
	if err := s.client.Do(ctx, "POST", "/balance_transactions/lookup", map[string]string{"transaction_id": transactionID}, &resp); err != nil {
		return nil, err
	}
	return &resp.Transaction, nil
}

// Page returns a paginated list of balance transactions.
//
// Results are sorted by creation date (newest first). Use this to view
// available balance, track aging, or reconcile payouts.
func (s *Service) Page(ctx context.Context, params PageParams) ([]Resource, error) {
	var resp struct {
		Page struct {
			Transactions []Resource `json:"transactions"`
		} `json:"page"`
	}
	if err := s.client.Do(ctx, "POST", "/balance_transactions/page", params, &resp); err != nil {
		return nil, err
	}
	return resp.Page.Transactions, nil
}

type Type string

const (
	TypePayment Type = "payment"
)

const (
	TypeRefund Type = "refund"
)

// BalanceTransaction represents a merchant balance entry caused by a payment or
// refund. Exactly one of PaymentID and RefundID is present, matching Type.
type Resource struct {
	// ID is the unique balance transaction identifier (read-only).
	// Starts with "bt_". Example: "bt_abc123def456"
	ID string `json:"id"`

	// Type identifies the semantic source, not the accounting direction.
	Type Type `json:"type"`

	// PaymentID is present only when Type is payment.
	PaymentID string `json:"payment_id,omitempty"`

	// RefundID is present only when Type is refund.
	RefundID string `json:"refund_id,omitempty"`

	// PayoutID identifies the payout that claimed this transaction, when present.
	PayoutID string `json:"payout_id,omitempty"`

	// OrderID is the strongly referenced source order ID.
	OrderID string `json:"order_id"`

	// Amount is the transaction amount in the public money shape.
	Amount money.Amount `json:"amount"`

	// Deprecated: the reviewed API does not return amount_expected. Use Amount.
	AmountExpected *money.Amount `json:"amount_expected,omitempty"`

	// Deprecated: the reviewed API does not return amount_available. Use Amount.
	AmountAvailable *money.Amount `json:"amount_available,omitempty"`

	// AvailableAt is when funds become eligible for payout (ISO 8601, read-only).
	AvailableAt *string `json:"available_at,omitempty"`

	// ClaimedAt is when the transaction was claimed for payout.
	ClaimedAt *string `json:"claimed_at,omitempty"`

	// PaidAt is when the transaction was paid out or otherwise settled.
	PaidAt *string `json:"paid_at,omitempty"`

	// CreatedAt is when the balance transaction was created (ISO 8601).
	CreatedAt string `json:"created_at"`

	// Deprecated: the reviewed API does not return payout_configuration on balance transactions.
	PayoutConfiguration *payout.Configuration `json:"payout_configuration,omitempty"`
}

// SourceID returns the matching strong source reference. It returns false for
// incomplete or contradictory transaction values.
func (t Resource) SourceID() (string, bool) {
	switch t.Type {
	case TypePayment:
		if t.PaymentID != "" && t.RefundID == "" {
			return t.PaymentID, true
		}
	case TypeRefund:
		if t.RefundID != "" && t.PaymentID == "" {
			return t.RefundID, true
		}
	}
	return "", false
}

// BalanceTransactionPageParams specifies pagination for listing balance transactions.
type PageParams struct {
	// PageNumber is the page to retrieve (optional, default: 1).
	// Pages are 1-indexed.
	PageNumber int `json:"page_number,omitempty"`

	// PageSize is the number of transactions per page (optional, default: 20).
	// Maximum 100.
	PageSize int `json:"page_size,omitempty"`
}

// NewService constructs the resource service used by inttegro.Client.
func NewService(client transport.Client) *Service { return &Service{client: client} }
