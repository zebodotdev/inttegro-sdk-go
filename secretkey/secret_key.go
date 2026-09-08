// Package secretkey provides API secret-key resources and operations.
package secretkey

import inttegro "github.com/zebodotdev/inttegro-sdk-go/v4"

type (
	Service        = inttegro.KeysService
	TokenType      = inttegro.SecretKeyTokenType
	Status         = inttegro.SecretKeyStatus
	AuthResult     = inttegro.SecretKeyAuthResult
	GenerateParams = inttegro.GenerateSecretKeyParams
	LookupParams   = inttegro.SecretKeyLookupParams
	UpdateParams   = inttegro.UpdateSecretKeyParams
	DestroyParams  = inttegro.DestroySecretKeyParams
	PageParams     = inttegro.PageSecretKeysParams
	UsageParams    = inttegro.SecretKeyUsageParams
	Generated      = inttegro.GeneratedSecretKey
	Resource       = inttegro.SecretKey
	Page           = inttegro.SecretKeyPage
	UsageRow       = inttegro.SecretKeyUsageRow
	UsagePage      = inttegro.SecretKeyUsagePage
	Usage          = inttegro.SecretKeyUsage
)

const (
	TokenTypeBearer = inttegro.SecretKeyTokenTypeBearer

	StatusActive  = inttegro.SecretKeyStatusActive
	StatusRevoked = inttegro.SecretKeyStatusRevoked
	StatusExpired = inttegro.SecretKeyStatusExpired

	AuthResultSucceeded = inttegro.SecretKeyAuthResultSucceeded
	AuthResultFailed    = inttegro.SecretKeyAuthResultFailed
)
