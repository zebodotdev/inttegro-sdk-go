package schedule

// ScheduleDetail describes a scheduled chime and its execution state.
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

// ScheduleError reports a per-recipient scheduling error.
type Error struct {
	Recipient string `json:"recipient,omitempty"`
	FixCode   string `json:"fix_code,omitempty"`
	Type      string `json:"type,omitempty"`
}

// ScheduledChime is returned after scheduling a chime.
type ScheduledChime struct {
	ID          string   `json:"id,omitempty"`
	Recipients  []string `json:"recipients,omitempty"`
	FullMessage string   `json:"full_message,omitempty"`
	SenderID    string   `json:"sender_id,omitempty"`
	Purpose     *string  `json:"purpose,omitempty"`
	SendAfter   string   `json:"send_after,omitempty"`
	CreatedAt   string   `json:"created_at,omitempty"`
	ExecutedAt  *string  `json:"executed_at,omitempty"`
}
