// Package broadcast provides broadcast resources and operations.
package broadcast

import (
	"context"

	"github.com/zebodotdev/inttegro-sdk-go/v5/internal/transport"
)

// BroadcastsService manages broadcast chime operations.
type Service struct {
	client transport.Client
}

// Lookup retrieves broadcast details by broadcast ID.
func (s *Service) Lookup(ctx context.Context, broadcastID string) (*Resource, error) {
	var resp struct {
		Broadcast Resource `json:"broadcast"`
	}
	if err := s.client.Do(ctx, "POST", "/broadcasts/lookup", LookupParams{BroadcastID: broadcastID}, &resp); err != nil {
		return nil, err
	}
	return &resp.Broadcast, nil
}

// Cancel cancels a broadcast by broadcast ID.
func (s *Service) Cancel(ctx context.Context, broadcastID string) (*Resource, error) {
	var resp struct {
		Broadcast Resource `json:"broadcast"`
	}
	if err := s.client.Do(ctx, "POST", "/broadcasts/cancel", CancelParams{BroadcastID: broadcastID}, &resp); err != nil {
		return nil, err
	}
	return &resp.Broadcast, nil
}

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

// LookupBroadcastParams specifies which broadcast to retrieve.
type LookupParams struct {
	BroadcastID string `json:"broadcast_id"`
}

// CancelBroadcastParams specifies which broadcast to cancel.
type CancelParams struct {
	BroadcastID string `json:"broadcast_id"`
}

// NewService constructs the resource service used by inttegro.Client.
func NewService(client transport.Client) *Service { return &Service{client: client} }
