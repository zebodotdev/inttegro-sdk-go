package app

import (
	"context"

	"github.com/zebodotdev/inttegro-sdk-go/v6/internal/transport"
	"github.com/zebodotdev/inttegro-sdk-go/v6/request"
)

// AppsService manages applications associated with Inttegro API keys.
type Service struct {
	client transport.Client
}

// Create creates an Inttegro child app.
func (s *Service) Create(ctx context.Context, params CreateParams) (*App, error) {
	var resp struct {
		App App `json:"app"`
	}
	if err := s.client.Do(ctx, "POST", "/apps/create", params, &resp); err != nil {
		return nil, err
	}
	return &resp.App, nil
}

// Lookup retrieves the application associated with the API key used for the request.
func (s *Service) Lookup(ctx context.Context) (*App, error) {
	var resp struct {
		App App `json:"app"`
	}
	if err := s.client.Do(ctx, "POST", "/apps/lookup", map[string]any{}, &resp); err != nil {
		return nil, err
	}
	return &resp.App, nil
}

// Update changes attributes of the application associated with the API key used for the request.
func (s *Service) Update(ctx context.Context, params UpdateParams, opts ...request.Option) (*App, error) {
	var resp struct {
		App App `json:"app"`
	}
	if err := s.client.DoJSON(ctx, "/apps/update", params, request.Apply(opts), &resp); err != nil {
		return nil, err
	}
	return &resp.App, nil
}

// NewService constructs the resource service used by inttegro.Client.
func NewService(client transport.Client) *Service { return &Service{client: client} }
