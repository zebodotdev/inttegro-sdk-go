package payout

type ScheduleParams struct {
	DestinationID string `json:"destination_id"`
	ExecuteAfter  string `json:"execute_after,omitempty"`
	MaxAmount     int64  `json:"max_amount"`
	Reference     string `json:"reference"`
}

// PayoutPageParams specifies pagination for listing payouts.
type PageParams struct {
	// PageNumber is the page to retrieve (optional, default: 1).
	// Pages are 1-indexed.
	PageNumber int `json:"page_number,omitempty"`

	// PageSize is the number of payouts per page (optional, default: 20).
	// Maximum 100.
	PageSize int `json:"page_size,omitempty"`
}
