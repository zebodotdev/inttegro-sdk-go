package inttegro

import "context"

// AppsService manages applications associated with Inttegro API keys.
type AppsService struct {
	client *Client
}

// Create creates a Inttegro application.
func (s *AppsService) Create(ctx context.Context, payload any) (map[string]any, error) {
	var resp map[string]any
	if err := s.client.do(ctx, "POST", "/apps/create", payload, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// Lookup retrieves the application associated with the API key used for the request.
func (s *AppsService) Lookup(ctx context.Context) (map[string]any, error) {
	var resp map[string]any
	if err := s.client.do(ctx, "POST", "/apps/lookup", map[string]any{}, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// Update changes attributes of the application associated with the API key used for the request.
func (s *AppsService) Update(ctx context.Context, payload any) (map[string]any, error) {
	var resp map[string]any
	if err := s.client.do(ctx, "POST", "/apps/update", payload, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}
