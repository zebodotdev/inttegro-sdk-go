package commerce

import "context"

// PlatformService manages platform-level application, key, and session endpoints.
type PlatformService struct {
	client *Client
}

// CreateApp creates a Commerce application.
func (s *PlatformService) CreateApp(ctx context.Context, payload any) (map[string]any, error) {
	var resp map[string]any
	if err := s.client.do(ctx, "POST", "/apps/create", payload, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// GenerateKey creates an API key for an application.
func (s *PlatformService) GenerateKey(ctx context.Context, payload any) (map[string]any, error) {
	var resp map[string]any
	if err := s.client.do(ctx, "POST", "/keys/generate", payload, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// NewSession creates a new platform session.
func (s *PlatformService) NewSession(ctx context.Context, payload any) (map[string]any, error) {
	var resp map[string]any
	if err := s.client.do(ctx, "POST", "/sessions/new", payload, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}
