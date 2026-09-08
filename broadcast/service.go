package broadcast

import (
	"context"

	"github.com/zebodotdev/inttegro-sdk-go/v6/internal/transport"
)

// BroadcastsService manages broadcast chime operations.
type Service struct {
	client transport.Client
}

// Lookup retrieves broadcast details by broadcast ID.
func (s *Service) Lookup(ctx context.Context, broadcastID string) (*Broadcast, error) {
	var resp struct {
		Broadcast Broadcast `json:"broadcast"`
	}
	if err := s.client.Do(ctx, "POST", "/broadcasts/lookup", LookupParams{BroadcastID: broadcastID}, &resp); err != nil {
		return nil, err
	}
	return &resp.Broadcast, nil
}

// Cancel cancels a broadcast by broadcast ID.
func (s *Service) Cancel(ctx context.Context, broadcastID string) (*Broadcast, error) {
	var resp struct {
		Broadcast Broadcast `json:"broadcast"`
	}
	if err := s.client.Do(ctx, "POST", "/broadcasts/cancel", CancelParams{BroadcastID: broadcastID}, &resp); err != nil {
		return nil, err
	}
	return &resp.Broadcast, nil
}

// NewService constructs the resource service used by inttegro.Client.
func NewService(client transport.Client) *Service { return &Service{client: client} }
