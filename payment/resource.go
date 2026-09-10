package payment

import (
	"time"

	"github.com/zebodotdev/inttegro-sdk-go/v6/balancetransaction"
	"github.com/zebodotdev/inttegro-sdk-go/v6/money"
	"github.com/zebodotdev/inttegro-sdk-go/v6/payout"
)

// Payment is the payment projection returned inside an order.
type Payment struct {
	ID                  string                                 `json:"id"`
	StatementDescriptor string                                 `json:"statement_descriptor"`
	PaymentMethodTypes  []string                               `json:"payment_method_types,omitempty"`
	PaymentMethod       *PaymentMethod                         `json:"payment_method,omitempty"`
	BillingDetails      *BillingDetails                        `json:"billing_details,omitempty"`
	Customer            *Customer                              `json:"customer,omitempty"`
	LatestAttempt       *Attempt                               `json:"latest_attempt,omitempty"`
	Amount              money.Amount                           `json:"amount"`
	NextAction          *NextAction                            `json:"next_action,omitempty"`
	LatestError         *Error                                 `json:"latest_error,omitempty"`
	BalanceTransaction  *balancetransaction.BalanceTransaction `json:"balance_transaction,omitempty"`
	PayoutConfiguration *payout.Configuration                  `json:"payout_configuration,omitempty"`
	Status              Status                                 `json:"status"`
	InitiatedAt         time.Time                              `json:"initiated_at"`
	ExecutedAt          *time.Time                             `json:"executed_at,omitempty"`
	DueAt               *time.Time                             `json:"due_at,omitempty"`
	CanceledAt          *time.Time                             `json:"canceled_at,omitempty"`
	ExpiredAt           *time.Time                             `json:"expired_at,omitempty"`
	PaidAt              *time.Time                             `json:"paid_at,omitempty"`
	PaidOffline         *bool                                  `json:"paid_offline,omitempty"`
	FailedAt            *time.Time                             `json:"failed_at,omitempty"`
}
