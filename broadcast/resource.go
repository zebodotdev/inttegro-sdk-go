package broadcast

// BroadcastCreation summarizes a broadcast enqueue request.
type Creation struct {
	BroadcastID     string `json:"broadcast_id,omitempty"`
	Status          string `json:"status,omitempty"`
	RecipientsCount int    `json:"recipients_count,omitempty"`
	QueuedAt        string `json:"queued_at,omitempty"`
}

// BroadcastDetail describes a broadcast chime and its execution state.
type Resource struct {
	ID         string   `json:"id,omitempty"`
	Recipients []string `json:"recipients,omitempty"`
	Content    string   `json:"content,omitempty"`
	SenderID   string   `json:"sender_id,omitempty"`
	Purpose    *string  `json:"purpose,omitempty"`
	SendAfter  string   `json:"send_after,omitempty"`
	CreatedAt  string   `json:"created_at,omitempty"`
	ExecutedAt *string  `json:"executed_at,omitempty"`
	CanceledAt *string  `json:"canceled_at,omitempty"`
	Errors     []Error  `json:"errors,omitempty"`
	ChimeIDs   []string `json:"chime_ids,omitempty"`
}

// BroadcastError reports a per-recipient broadcast error.
type Error struct {
	Recipient string `json:"recipient,omitempty"`
	FixCode   string `json:"fix_code,omitempty"`
	Type      string `json:"type,omitempty"`
}
