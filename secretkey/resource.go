package secretkey

type Generated struct {
	ID        string    `json:"id"`
	Label     string    `json:"label,omitempty"`
	TokenType TokenType `json:"token_type"`
	IssuedAt  string    `json:"issued_at"`
	Token     string    `json:"token"`
}

// SecretKey represents an application secret key.
type SecretKey struct {
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
	Number  int         `json:"number"`
	Size    int         `json:"size"`
	Count   int         `json:"count"`
	Total   int         `json:"total"`
	HasMore bool        `json:"has_more"`
	Keys    []SecretKey `json:"keys"`
}
