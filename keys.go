package inttegro

import "context"

// KeysService manages secret keys for the authenticated application.
type KeysService struct {
	client *Client
}

type GenerateSecretKeyParams struct {
	Label string `json:"label,omitempty"`
}

type SecretKeyLookupParams struct {
	SecretKeyID string `json:"secret_key_id,omitempty"`
}

type UpdateSecretKeyParams struct {
	SecretKeyID string `json:"secret_key_id,omitempty"`
	Label       string `json:"label"`
}

type DestroySecretKeyParams struct {
	SecretKeyID string `json:"secret_key_id,omitempty"`
}

type PageSecretKeysParams struct {
	Page   int `json:"page,omitempty"`
	Number int `json:"number,omitempty"`
	Size   int `json:"size,omitempty"`
}

type SecretKeyUsageParams struct {
	SecretKeyID string `json:"secret_key_id,omitempty"`
	Page        int    `json:"page,omitempty"`
	Number      int    `json:"number,omitempty"`
	Size        int    `json:"size,omitempty"`
}

type GeneratedSecretKey struct {
	ID        string             `json:"id"`
	Label     string             `json:"label,omitempty"`
	TokenType SecretKeyTokenType `json:"token_type"`
	IssuedAt  string             `json:"issued_at"`
	Token     string             `json:"token"`
}

type SecretKey struct {
	ID         string             `json:"id"`
	Label      string             `json:"label,omitempty"`
	TokenType  SecretKeyTokenType `json:"token_type"`
	IssuedAt   string             `json:"issued_at"`
	UpdatedAt  string             `json:"updated_at,omitempty"`
	ExpiresAt  string             `json:"expires_at,omitempty"`
	Status     SecretKeyStatus    `json:"status"`
	Active     bool               `json:"active"`
	RevokedAt  string             `json:"revoked_at,omitempty"`
	LastUsedAt string             `json:"last_used_at,omitempty"`
	UsageCount int                `json:"usage_count,omitempty"`
}

type SecretKeyPage struct {
	Number  int         `json:"number"`
	Size    int         `json:"size"`
	Count   int         `json:"count"`
	Total   int         `json:"total"`
	HasMore bool        `json:"has_more"`
	Keys    []SecretKey `json:"keys"`
}

type SecretKeyUsageRow struct {
	SecretKeyID string              `json:"secret_key_id"`
	OccurredAt  string              `json:"occurred_at"`
	AuthResult  SecretKeyAuthResult `json:"auth_result"`
}

type SecretKeyUsagePage struct {
	Number  int                 `json:"number"`
	Size    int                 `json:"size"`
	Count   int                 `json:"count"`
	Total   int                 `json:"total"`
	HasMore bool                `json:"has_more"`
	Rows    []SecretKeyUsageRow `json:"rows"`
}

type SecretKeyUsageResponse struct {
	Key   SecretKey          `json:"key"`
	Usage SecretKeyUsagePage `json:"usage"`
}

// Generate creates a new active secret key. The token is returned only once.
func (s *KeysService) Generate(ctx context.Context, params GenerateSecretKeyParams) (*GeneratedSecretKey, error) {
	var resp struct {
		Key GeneratedSecretKey `json:"key"`
	}
	if err := s.client.do(ctx, "POST", "/keys/generate", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Key, nil
}

// Page lists safe secret key metadata.
func (s *KeysService) Page(ctx context.Context, params PageSecretKeysParams) (*SecretKeyPage, error) {
	var resp struct {
		Page SecretKeyPage `json:"page"`
	}
	if err := s.client.do(ctx, "POST", "/keys/page", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Page, nil
}

// Lookup retrieves safe metadata for a secret key by ID.
func (s *KeysService) Lookup(ctx context.Context, secretKeyID string) (*SecretKey, error) {
	return s.LookupWithParams(ctx, SecretKeyLookupParams{SecretKeyID: secretKeyID})
}

// LookupWithParams retrieves safe metadata using the canonical secret key ID field.
func (s *KeysService) LookupWithParams(ctx context.Context, params SecretKeyLookupParams) (*SecretKey, error) {
	var resp struct {
		Key SecretKey `json:"key"`
	}
	if err := s.client.do(ctx, "POST", "/keys/lookup", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Key, nil
}

// Update changes safe mutable metadata for a secret key.
func (s *KeysService) Update(ctx context.Context, params UpdateSecretKeyParams) (*SecretKey, error) {
	var resp struct {
		Key SecretKey `json:"key"`
	}
	if err := s.client.do(ctx, "POST", "/keys/update", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Key, nil
}

// Destroy revokes a secret key.
func (s *KeysService) Destroy(ctx context.Context, secretKeyID string) (*SecretKey, error) {
	return s.DestroyWithParams(ctx, DestroySecretKeyParams{SecretKeyID: secretKeyID})
}

// DestroyWithParams revokes a secret key using the canonical secret key ID field.
func (s *KeysService) DestroyWithParams(ctx context.Context, params DestroySecretKeyParams) (*SecretKey, error) {
	var resp struct {
		Key SecretKey `json:"key"`
	}
	if err := s.client.do(ctx, "POST", "/keys/destroy", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Key, nil
}

// Usage retrieves recent public authentication outcomes for a secret key.
func (s *KeysService) Usage(ctx context.Context, params SecretKeyUsageParams) (*SecretKeyUsageResponse, error) {
	var resp SecretKeyUsageResponse
	if err := s.client.do(ctx, "POST", "/keys/usage", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
