package inttegro

import "context"

// BroadcastsService manages broadcast chime operations.
type BroadcastsService struct {
	client *Client
}

// Lookup retrieves broadcast details by broadcast ID.
func (s *BroadcastsService) Lookup(ctx context.Context, broadcastID string) (*BroadcastDetail, error) {
	var resp struct {
		Broadcast BroadcastDetail `json:"broadcast"`
	}
	if err := s.client.do(ctx, "POST", "/broadcasts/lookup", LookupBroadcastParams{BroadcastID: broadcastID}, &resp); err != nil {
		return nil, err
	}
	return &resp.Broadcast, nil
}

// Cancel cancels a broadcast by broadcast ID.
func (s *BroadcastsService) Cancel(ctx context.Context, broadcastID string) (*BroadcastDetail, error) {
	var resp struct {
		Broadcast BroadcastDetail `json:"broadcast"`
	}
	if err := s.client.do(ctx, "POST", "/broadcasts/cancel", CancelBroadcastParams{BroadcastID: broadcastID}, &resp); err != nil {
		return nil, err
	}
	return &resp.Broadcast, nil
}
