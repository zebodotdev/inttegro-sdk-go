package order

import (
	"context"
)

// Update modifies mutable fields on an existing order.
func (s *Service) Update(ctx context.Context, payload any) (*Resource, error) {
	var resp struct {
		Order Resource `json:"order"`
	}
	if err := s.client.Do(ctx, "POST", "/orders/update", payload, &resp); err != nil {
		return nil, err
	}
	return &resp.Order, nil
}
