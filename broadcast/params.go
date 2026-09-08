package broadcast

// BroadcastChimeParams sends a broadcast chime to multiple recipients.
type CreateParams struct {
	Recipients       []string `json:"recipients,omitempty"`
	MessageTemplate  string   `json:"message_template,omitempty"`
	ServiceName      string   `json:"service_name,omitempty"`
	Sender           string   `json:"sender,omitempty"`
	Purpose          string   `json:"purpose,omitempty"`
	PreferredGateway string   `json:"preferred_gateway,omitempty"`
	IdempotencyKey   string   `json:"idempotency_key,omitempty"`
}

// LookupBroadcastParams specifies which broadcast to retrieve.
type LookupParams struct {
	BroadcastID string `json:"broadcast_id"`
}

// CancelBroadcastParams specifies which broadcast to cancel.
type CancelParams struct {
	BroadcastID string `json:"broadcast_id"`
}
