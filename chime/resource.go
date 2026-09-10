package chime

import "time"

type Page struct {
	Number int     `json:"number,omitempty"`
	Size   int     `json:"size,omitempty"`
	Chimes []Chime `json:"chimes,omitempty"`
}

// Chime represents a notification message.
//
// Chimes track delivery status, transmission details, and any errors
// that occurred during sending.
type Chime struct {
	// ID is the unique chime identifier (read-only).
	// Starts with "chm_". Example: "chm_abc123def456"
	ID string `json:"id,omitempty"`

	// CreatedAt is when the chime was created (ISO 8601, read-only).
	CreatedAt *time.Time `json:"created_at,omitempty"`

	// FullMessage is the message content that was sent.
	FullMessage string `json:"full_message,omitempty"`

	// Recipient contains the recipient details.
	Recipient *Recipient `json:"recipient,omitempty"`

	// SenderID is the sender name shown to the recipient.
	SenderID string `json:"sender_id,omitempty"`

	// Purpose is the chime category.
	Purpose string `json:"purpose,omitempty"`

	// CustomData contains attached custom data.
	CustomData map[string]string `json:"custom_data,omitempty"`

	// Delivery contains delivery status and timestamps.
	// Tracks whether the message was successfully delivered.
	Delivery any `json:"delivery,omitempty"`

	// Transmission contains low-level transmission details.
	// Includes carrier responses, error codes, etc.
	Transmission any `json:"transmission,omitempty"`
}
