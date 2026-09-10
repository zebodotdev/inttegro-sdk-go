package schedule

import "time"

// Schedule describes a scheduled chime and its execution state.
type Schedule struct {
	ID         string     `json:"id,omitempty"`
	Recipients []string   `json:"recipients,omitempty"`
	Content    string     `json:"content,omitempty"`
	SenderID   string     `json:"sender_id,omitempty"`
	Purpose    *string    `json:"purpose,omitempty"`
	SendAfter  *time.Time `json:"send_after,omitempty"`
	CreatedAt  *time.Time `json:"created_at,omitempty"`
	ExecutedAt *time.Time `json:"executed_at,omitempty"`
	CanceledAt *time.Time `json:"canceled_at,omitempty"`
	Errors     []Error    `json:"errors,omitempty"`
	ChimeIDs   []string   `json:"chime_ids,omitempty"`
}

// ScheduleError reports a per-recipient scheduling error.
type Error struct {
	Recipient string `json:"recipient,omitempty"`
	FixCode   string `json:"fix_code,omitempty"`
	Type      string `json:"type,omitempty"`
}

// ScheduledChime is returned after scheduling a chime.
type ScheduledChime struct {
	ID          string     `json:"id,omitempty"`
	Recipients  []string   `json:"recipients,omitempty"`
	FullMessage string     `json:"full_message,omitempty"`
	SenderID    string     `json:"sender_id,omitempty"`
	Purpose     *string    `json:"purpose,omitempty"`
	SendAfter   *time.Time `json:"send_after,omitempty"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	ExecutedAt  *time.Time `json:"executed_at,omitempty"`
}
