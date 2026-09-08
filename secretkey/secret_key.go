// Package secretkey provides secretkey resources and operations.
package secretkey

import (
	"context"

	"github.com/zebodotdev/inttegro-sdk-go/v5/internal/transport"
)

type TokenType string

const TokenTypeBearer TokenType = "bearer"

type Status string

const (
	StatusActive Status = "active"
)

const (
	StatusRevoked Status = "revoked"
)

const (
	StatusExpired Status = "expired"
)

type AuthResult string

const (
	AuthResultSucceeded AuthResult = "succeeded"
)

const (
	AuthResultFailed AuthResult = "failed"
)

// KeysService manages secret keys for the authenticated application.
type Service struct {
	client transport.Client
}

type GenerateParams struct {
	Label string `json:"label,omitempty"`
}

type LookupParams struct {
	SecretKeyID string `json:"secret_key_id,omitempty"`
}

type UpdateParams struct {
	SecretKeyID string `json:"secret_key_id,omitempty"`
	Label       string `json:"label"`
}

type DestroyParams struct {
	SecretKeyID string `json:"secret_key_id,omitempty"`
}

type PageParams struct {
	Page   int `json:"page,omitempty"`
	Number int `json:"number,omitempty"`
	Size   int `json:"size,omitempty"`
}

type UsageParams struct {
	SecretKeyID string `json:"secret_key_id,omitempty"`
	Page        int    `json:"page,omitempty"`
	Number      int    `json:"number,omitempty"`
	Size        int    `json:"size,omitempty"`
}

type Generated struct {
	ID        string    `json:"id"`
	Label     string    `json:"label,omitempty"`
	TokenType TokenType `json:"token_type"`
	IssuedAt  string    `json:"issued_at"`
	Token     string    `json:"token"`
}

type Resource struct {
	ID         string    `json:"id"`
	Label      string    `json:"label,omitempty"`
	TokenType  TokenType `json:"token_type"`
	IssuedAt   string    `json:"issued_at"`
	UpdatedAt  string    `json:"updated_at,omitempty"`
	ExpiresAt  string    `json:"expires_at,omitempty"`
	Status     Status    `json:"status"`
	Active     bool      `json:"active"`
	RevokedAt  string    `json:"revoked_at,omitempty"`
	LastUsedAt string    `json:"last_used_at,omitempty"`
	UsageCount int       `json:"usage_count,omitempty"`
}

type Page struct {
	Number  int        `json:"number"`
	Size    int        `json:"size"`
	Count   int        `json:"count"`
	Total   int        `json:"total"`
	HasMore bool       `json:"has_more"`
	Keys    []Resource `json:"keys"`
}

type UsageRow struct {
	SecretKeyID string     `json:"secret_key_id"`
	OccurredAt  string     `json:"occurred_at"`
	AuthResult  AuthResult `json:"auth_result"`
}

type UsagePage struct {
	Number  int        `json:"number"`
	Size    int        `json:"size"`
	Count   int        `json:"count"`
	Total   int        `json:"total"`
	HasMore bool       `json:"has_more"`
	Rows    []UsageRow `json:"rows"`
}

// SecretKeyUsage contains key metadata and its recent authentication outcomes.
type Usage struct {
	Key   Resource  `json:"key"`
	Usage UsagePage `json:"usage"`
}

// Generate creates a new active secret key. The token is returned only once.
func (s *Service) Generate(ctx context.Context, params GenerateParams) (*Generated, error) {
	var resp struct {
		Key Generated `json:"key"`
	}
	if err := s.client.Do(ctx, "POST", "/keys/generate", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Key, nil
}

// Page lists safe secret key metadata.
func (s *Service) Page(ctx context.Context, params PageParams) (*Page, error) {
	var resp struct {
		Page Page `json:"page"`
	}
	if err := s.client.Do(ctx, "POST", "/keys/page", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Page, nil
}

// Lookup retrieves safe metadata for a secret key by ID.
func (s *Service) Lookup(ctx context.Context, secretKeyID string) (*Resource, error) {
	return s.LookupWithParams(ctx, LookupParams{SecretKeyID: secretKeyID})
}

// LookupWithParams retrieves safe metadata using the canonical secret key ID field.
func (s *Service) LookupWithParams(ctx context.Context, params LookupParams) (*Resource, error) {
	var resp struct {
		Key Resource `json:"key"`
	}
	if err := s.client.Do(ctx, "POST", "/keys/lookup", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Key, nil
}

// Update changes safe mutable metadata for a secret key.
func (s *Service) Update(ctx context.Context, params UpdateParams) (*Resource, error) {
	var resp struct {
		Key Resource `json:"key"`
	}
	if err := s.client.Do(ctx, "POST", "/keys/update", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Key, nil
}

// Destroy revokes a secret key.
func (s *Service) Destroy(ctx context.Context, secretKeyID string) (*Resource, error) {
	return s.DestroyWithParams(ctx, DestroyParams{SecretKeyID: secretKeyID})
}

// DestroyWithParams revokes a secret key using the canonical secret key ID field.
func (s *Service) DestroyWithParams(ctx context.Context, params DestroyParams) (*Resource, error) {
	var resp struct {
		Key Resource `json:"key"`
	}
	if err := s.client.Do(ctx, "POST", "/keys/destroy", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Key, nil
}

// Usage retrieves recent public authentication outcomes for a secret key.
func (s *Service) Usage(ctx context.Context, params UsageParams) (*Usage, error) {
	var resp Usage
	if err := s.client.Do(ctx, "POST", "/keys/usage", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// NewService constructs the resource service used by inttegro.Client.
func NewService(client transport.Client) *Service { return &Service{client: client} }
