// Package money contains currency and amount primitives used by the Inttegro API.
package money

// Currency is a lowercase currency code on the wire.
type Currency string

// Currency constants use their conventional uppercase identifiers while
// retaining the lowercase values required by the API.
const (
	GHS Currency = "ghs"
	USD Currency = "usd"
	GBP Currency = "gbp"
	EUR Currency = "eur"
	CNY Currency = "cny"
)

// AmountParams is an amount supplied in a request. Value is expressed in the
// currency's smallest unit.
type AmountParams struct {
	Currency Currency `json:"currency"`
	Value    int64    `json:"value"`
}

// Amount is an amount returned by the API. Value is expressed in the
// currency's smallest unit.
type Amount struct {
	Currency Currency `json:"currency"`
	Value    int64    `json:"value"`
}
