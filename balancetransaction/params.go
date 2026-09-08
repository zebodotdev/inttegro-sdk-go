package balancetransaction

// BalanceTransactionPageParams specifies pagination for listing balance transactions.
type PageParams struct {
	// PageNumber is the page to retrieve (optional, default: 1).
	// Pages are 1-indexed.
	PageNumber int `json:"page_number,omitempty"`

	// PageSize is the number of transactions per page (optional, default: 20).
	// Maximum 100.
	PageSize int `json:"page_size,omitempty"`
}
