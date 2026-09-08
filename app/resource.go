package app

import (
	"github.com/zebodotdev/inttegro-sdk-go/v5/secretkey"
)

// AppSecretKey is the initial secret key returned when an app is created.
// The token is returned only once and is not included in later app lookups.
type SecretKey struct {
	ID        string              `json:"id,omitempty"`
	TokenType secretkey.TokenType `json:"token_type,omitempty"`
	IssuedAt  string              `json:"issued_at,omitempty"`
	Token     string              `json:"token,omitempty"`
}

// App represents an Inttegro account.
type Resource struct {
	ID           string        `json:"id"`
	Name         string        `json:"name"`
	Alias        string        `json:"alias,omitempty"`
	Description  string        `json:"description,omitempty"`
	CreatedAt    string        `json:"created_at"`
	UpdatedAt    string        `json:"updated_at,omitempty"`
	ArchivedAt   string        `json:"archived_at,omitempty"`
	SecretKey    *SecretKey    `json:"secret_key,omitempty"`
	Relationship *Relationship `json:"relationship,omitempty"`
}
