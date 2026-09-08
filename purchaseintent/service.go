package purchaseintent

import (
	"context"

	"github.com/zebodotdev/inttegro-sdk-go/v5/internal/transport"
)

// PurchaseIntentsService manages Buy link purchase intents.
type Service struct {
	client transport.Client
}

// Create creates a Buy link purchase intent.
func (s *Service) Create(ctx context.Context, params CreateParams) (*Resource, error) {
	var resp struct {
		PurchaseIntent Resource `json:"purchase_intent"`
	}
	if err := s.client.Do(ctx, "POST", "/purchase_intents/create", params, &resp); err != nil {
		return nil, err
	}
	return &resp.PurchaseIntent, nil
}

// Update modifies mutable Buy link purchase intent fields.
func (s *Service) Update(ctx context.Context, params UpdateParams) (*Resource, error) {
	var resp struct {
		PurchaseIntent Resource `json:"purchase_intent"`
	}
	if err := s.client.Do(ctx, "POST", "/purchase_intents/update", params, &resp); err != nil {
		return nil, err
	}
	return &resp.PurchaseIntent, nil
}

// Cancel cancels a Buy link purchase intent.
func (s *Service) Cancel(ctx context.Context, id string) (*Resource, error) {
	var resp struct {
		PurchaseIntent Resource `json:"purchase_intent"`
	}
	if err := s.client.Do(ctx, "POST", "/purchase_intents/cancel", map[string]string{"id": id}, &resp); err != nil {
		return nil, err
	}
	return &resp.PurchaseIntent, nil
}

// Lookup retrieves a Buy link purchase intent by ID.
func (s *Service) Lookup(ctx context.Context, id string) (*Resource, error) {
	var resp struct {
		PurchaseIntent Resource `json:"purchase_intent"`
	}
	if err := s.client.Do(ctx, "POST", "/purchase_intents/lookup", map[string]string{"id": id}, &resp); err != nil {
		return nil, err
	}
	return &resp.PurchaseIntent, nil
}

// Page lists Buy link purchase intents.
func (s *Service) Page(ctx context.Context, params PageParams) (*Page, error) {
	var resp struct {
		Page Page `json:"page"`
	}
	if err := s.client.Do(ctx, "POST", "/purchase_intents/page", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Page, nil
}

// NewService constructs the resource service used by inttegro.Client.
func NewService(client transport.Client) *Service { return &Service{client: client} }
