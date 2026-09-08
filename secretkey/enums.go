package secretkey

type TokenType string

const TokenTypeBearer TokenType = "bearer"

type Status string

const (
	StatusActive  Status = "active"
	StatusRevoked Status = "revoked"
	StatusExpired Status = "expired"
)

type AuthResult string

const (
	AuthResultSucceeded AuthResult = "succeeded"
	AuthResultFailed    AuthResult = "failed"
)
