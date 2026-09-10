package secretkey

import "time"

type Generated struct {
	ID        string    `json:"id"`
	Label     string    `json:"label,omitempty"`
	TokenType TokenType `json:"token_type"`
	IssuedAt  time.Time `json:"issued_at"`
	Token     string    `json:"token"`
}

// SecretKey represents an application secret key.
type SecretKey struct {
	ID         string     `json:"id"`
	Label      string     `json:"label,omitempty"`
	TokenType  TokenType  `json:"token_type"`
	IssuedAt   time.Time  `json:"issued_at"`
	UpdatedAt  *time.Time `json:"updated_at,omitempty"`
	ExpiresAt  *time.Time `json:"expires_at,omitempty"`
	Status     Status     `json:"status"`
	Active     bool       `json:"active"`
	RevokedAt  *time.Time `json:"revoked_at,omitempty"`
	LastUsedAt *time.Time `json:"last_used_at,omitempty"`
	UsageCount int        `json:"usage_count,omitempty"`
}

type Page struct {
	Number  int         `json:"number"`
	Size    int         `json:"size"`
	Count   int         `json:"count"`
	Total   int         `json:"total"`
	HasMore bool        `json:"has_more"`
	Keys    []SecretKey `json:"keys"`
}
