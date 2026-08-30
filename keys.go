package commerce

import "context"

// KeysService manages secret keys for the authenticated application.
type KeysService struct {
	client *Client
}

type GenerateSecretKeyParams struct {
	Label      string `json:"label,omitempty"`
	UserAgent  string `json:"user_agent,omitempty"`
	RemoteAddr string `json:"remote_addr,omitempty"`
}

type SecretKeyLookupParams struct {
	SecretKeyID string `json:"secret_key_id,omitempty"`
	KeyID       string `json:"key_id,omitempty"`
	ID          string `json:"id,omitempty"`
}

type UpdateSecretKeyParams struct {
	SecretKeyID string `json:"secret_key_id,omitempty"`
	KeyID       string `json:"key_id,omitempty"`
	ID          string `json:"id,omitempty"`
	Label       string `json:"label"`
	UserAgent   string `json:"user_agent,omitempty"`
	RemoteAddr  string `json:"remote_addr,omitempty"`
}

type DestroySecretKeyParams struct {
	SecretKeyID string `json:"secret_key_id,omitempty"`
	KeyID       string `json:"key_id,omitempty"`
	ID          string `json:"id,omitempty"`
	UserAgent   string `json:"user_agent,omitempty"`
	RemoteAddr  string `json:"remote_addr,omitempty"`
}

type PageSecretKeysParams struct {
	Page   int `json:"page,omitempty"`
	Number int `json:"number,omitempty"`
	Size   int `json:"size,omitempty"`
}

type SecretKeyUsageParams struct {
	SecretKeyID string `json:"secret_key_id,omitempty"`
	KeyID       string `json:"key_id,omitempty"`
	ID          string `json:"id,omitempty"`
	Page        int    `json:"page,omitempty"`
	Number      int    `json:"number,omitempty"`
	Size        int    `json:"size,omitempty"`
}

type GeneratedSecretKey struct {
	ID        string `json:"id"`
	Label     string `json:"label,omitempty"`
	TokenType string `json:"token_type"`
	IssuedAt  string `json:"issued_at"`
	Token     string `json:"token"`
}

type SecretKey struct {
	ID                      string `json:"id"`
	Label                   string `json:"label,omitempty"`
	TokenType               string `json:"token_type"`
	IssuedAt                string `json:"issued_at"`
	UpdatedAt               string `json:"updated_at,omitempty"`
	UpdatedBy               string `json:"updated_by,omitempty"`
	ExpiresAt               string `json:"expires_at,omitempty"`
	Status                  string `json:"status"`
	Active                  bool   `json:"active"`
	RevokedAt               string `json:"revoked_at,omitempty"`
	RevokedBy               string `json:"revoked_by,omitempty"`
	RequestID               string `json:"request_id,omitempty"`
	IPAddress               string `json:"ip_address,omitempty"`
	UserAgent               string `json:"user_agent,omitempty"`
	KeyGen                  string `json:"key_gen,omitempty"`
	GeneratedByService      string `json:"generated_by_service,omitempty"`
	GeneratedByUserID       string `json:"generated_by_user_id,omitempty"`
	GeneratedByTeamMemberID string `json:"generated_by_team_member_id,omitempty"`
	GeneratedByEmail        string `json:"generated_by_email,omitempty"`
	GeneratedByName         string `json:"generated_by_name,omitempty"`
	RevocationRequestID     string `json:"revocation_request_id,omitempty"`
	RevocationIPAddress     string `json:"revocation_ip_address,omitempty"`
	RevocationUserAgent     string `json:"revocation_user_agent,omitempty"`
	RevokedByService        string `json:"revoked_by_service,omitempty"`
	RevokedByUserID         string `json:"revoked_by_user_id,omitempty"`
	RevokedByTeamMemberID   string `json:"revoked_by_team_member_id,omitempty"`
	RevokedByEmail          string `json:"revoked_by_email,omitempty"`
	RevokedByName           string `json:"revoked_by_name,omitempty"`
	CipherTextPrefix        string `json:"cipher_text_prefix,omitempty"`
	CipherTextLength        int    `json:"cipher_text_length,omitempty"`
	LastUsedAt              string `json:"last_used_at,omitempty"`
	UsageCount              int    `json:"usage_count"`
	UsageMetricsAvailable   bool   `json:"usage_metrics_available"`
}

type SecretKeyPage struct {
	Number                int         `json:"number"`
	Size                  int         `json:"size"`
	Count                 int         `json:"count"`
	Total                 int         `json:"total"`
	HasMore               bool        `json:"has_more"`
	UsageMetricsAvailable bool        `json:"usage_metrics_available"`
	Keys                  []SecretKey `json:"keys"`
}

type SecretKeyUsageRow struct {
	ID             string `json:"id"`
	Type           string `json:"type"`
	SecretKeyID    string `json:"secret_key_id"`
	SessionID      string `json:"session_id,omitempty"`
	VerificationID string `json:"verification_id,omitempty"`
	RequestID      string `json:"request_id,omitempty"`
	OccurredAt     string `json:"occurred_at"`
	CreatedAt      string `json:"created_at,omitempty"`
	InitiatedAt    string `json:"initiated_at,omitempty"`
	ExpiresAt      string `json:"expires_at,omitempty"`
	IPAddress      string `json:"ip_address,omitempty"`
	UserAgent      string `json:"user_agent,omitempty"`
	Verified       bool   `json:"verified"`
	AuthResult     string `json:"auth_result"`
	MultiUse       bool   `json:"multi_use,omitempty"`
}

type SecretKeyUsagePage struct {
	Number                        int                 `json:"number"`
	Size                          int                 `json:"size"`
	Count                         int                 `json:"count"`
	Total                         int                 `json:"total"`
	HasMore                       bool                `json:"has_more"`
	VerificationAttemptsAvailable bool                `json:"verification_attempts_available"`
	Rows                          []SecretKeyUsageRow `json:"rows"`
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

// LookupWithParams retrieves safe metadata using any supported ID alias.
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

// DestroyWithParams revokes a secret key using any supported ID alias.
func (s *KeysService) DestroyWithParams(ctx context.Context, params DestroySecretKeyParams) (*SecretKey, error) {
	var resp struct {
		Key SecretKey `json:"key"`
	}
	if err := s.client.do(ctx, "POST", "/keys/destroy", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Key, nil
}

// Usage retrieves successful session usage and attributable failed verification attempts.
func (s *KeysService) Usage(ctx context.Context, params SecretKeyUsageParams) (*SecretKeyUsageResponse, error) {
	var resp SecretKeyUsageResponse
	if err := s.client.do(ctx, "POST", "/keys/usage", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
