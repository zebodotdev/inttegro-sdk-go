package commerce

import "context"

// BalanceTransactionsService provides access to balance transaction history.
//
// Balance transactions represent funds from completed payments. Each successful
// payment creates a balance transaction that ages for 7 days (or your configured
// aging period) before becoming eligible for payout.
//
// Use this service to:
//   - View available and pending balance
//   - Track balance transaction aging
//   - Reconcile payouts with source payments
//
// Example:
//
//	transactions, err := client.BalanceTransactions.Page(ctx, commerce.BalanceTransactionPageParams{
//	    PageSize: 100,
//	})
//	for _, tx := range transactions {
//	    fmt.Printf("%s: %s %d (available: %s)\n",
//	        tx.ID, tx.AmountAvailable.Currency,
//	        tx.AmountAvailable.Value, tx.AvailableAt)
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
