package app

import (
	"time"

	"github.com/zebodotdev/inttegro-sdk-go/v6/secretkey"
)

// AppSecretKey is the initial secret key returned when an app is created.
// The token is returned only once and is not included in later app lookups.
type SecretKey struct {
	ID        string              `json:"id,omitempty"`
	TokenType secretkey.TokenType `json:"token_type,omitempty"`
	IssuedAt  *time.Time          `json:"issued_at,omitempty"`
	Token     string              `json:"token,omitempty"`
}

// App represents an Inttegro account.
type App struct {
	ID           string        `json:"id"`
	Name         string        `json:"name"`
	Alias        string        `json:"alias,omitempty"`
	Description  string        `json:"description,omitempty"`
	CreatedAt    time.Time     `json:"created_at"`
	UpdatedAt    *time.Time    `json:"updated_at,omitempty"`
	ArchivedAt   *time.Time    `json:"archived_at,omitempty"`
	SecretKey    *SecretKey    `json:"secret_key,omitempty"`
	Relationship *Relationship `json:"relationship,omitempty"`
}
