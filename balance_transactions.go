package inttegro

import "context"

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
//	transactions, err := client.BalanceTransactions.Page(ctx, inttegro.BalanceTransactionPageParams{
//	    PageSize: 100,
//	})
//	for _, tx := range transactions {
//	    sourceID, _ := tx.SourceID()
//	    fmt.Printf("%s (%s %s): %s %d\n",
//	        tx.ID, tx.Type, sourceID, tx.Amount.Currency, tx.Amount.Value)
//	}
type BalanceTransactionsService struct {
	client *Client
}

// Lookup retrieves a balance transaction by ID.
func (s *BalanceTransactionsService) Lookup(ctx context.Context, transactionID string) (*BalanceTransaction, error) {
	var resp struct {
		Transaction BalanceTransaction `json:"transaction"`
	}
	if err := s.client.do(ctx, "POST", "/balance_transactions/lookup", map[string]string{"transaction_id": transactionID}, &resp); err != nil {
		return nil, err
	}
	return &resp.Transaction, nil
}

// Page returns a paginated list of balance transactions.
//
// Results are sorted by creation date (newest first). Use this to view
// available balance, track aging, or reconcile payouts.
func (s *BalanceTransactionsService) Page(ctx context.Context, params BalanceTransactionPageParams) ([]BalanceTransaction, error) {
	var resp struct {
		Page struct {
			Transactions []BalanceTransaction `json:"transactions"`
		} `json:"page"`
	}
	if err := s.client.do(ctx, "POST", "/balance_transactions/page", params, &resp); err != nil {
		return nil, err
	}
	return resp.Page.Transactions, nil
}
