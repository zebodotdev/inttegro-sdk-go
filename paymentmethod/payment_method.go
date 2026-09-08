// Package paymentmethod provides paymentmethod resources and operations.
package paymentmethod

import (
	"context"

	"github.com/zebodotdev/inttegro-sdk-go/v5/bankaccount"
	"github.com/zebodotdev/inttegro-sdk-go/v5/internal/transport"
	"github.com/zebodotdev/inttegro-sdk-go/v5/request"
)

type Type string

const (
	TypeMobileMoney Type = "mobile_money"
)

const (
	TypeBankAccount Type = "bank_account"
)

const (
	TypeCard Type = "card"
)

const (
	TypeMotito Type = "motito"
)

// PaymentMethodsService manages payment method tokenization, verification, and deletion.
//
// Payment methods are saved customer payment instruments that can be charged
// repeatedly. Use this service to:
//
//   - Tokenize payment methods for repeat customers
//   - Verify ownership via OTP
//   - Look up payment method details
//   - Delete payment methods
//   - View payment method acceptance settings
//
// Example:
//
//	// Tokenize and verify a mobile money wallet
//	pm, err := client.PaymentMethods.Tokenize(ctx, paymentmethod.TokenizeParams{
//	    CustomerID: "cu_abc123",
//	    PaymentMethodData: paymentmethod.Data{
//	        Type: paymentmethod.TypeMobileMoney,
//	        MobileMoney: &paymentmethod.MobileMoneyParams{
//	            Network: "mtn",
//	            AccountNumber: "+233244123456",
//	        },
//	    },
//	    VerifyImmediately: inttegro.Bool(true),
//	})
//
// Learn more: https://studio.inttegro.com/save-payment-methods
type Service struct {
	client transport.Client
}

type PageParams struct {
	CustomerID string `json:"customer_id,omitempty"`
	PageNumber int    `json:"page_number,omitempty"`
	PageSize   int    `json:"page_size,omitempty"`
}

type Page struct {
	Number         int        `json:"number,omitempty"`
	Size           int        `json:"size,omitempty"`
	PaymentMethods []Resource `json:"payment_methods,omitempty"`
}

type ActionParams struct {
	PaymentMethodID string `json:"payment_method_id"`
}

// Tokenize saves a payment method for future use. Optionally verifies immediately.
//
// Tokenized payment methods can be charged repeatedly without re-entering details.
// The customer owns the payment method—only they can delete it.
//
// Learn more: https://studio.inttegro.com/tokenize-payment-methods
func (s *Service) Tokenize(ctx context.Context, params TokenizeParams) (*Resource, error) {
	var resp struct {
		PaymentMethod Resource `json:"payment_method"`
	}
	if err := s.client.Do(ctx, "POST", "/payment_methods/tokenize", params, &resp); err != nil {
		return nil, err
	}
	return &resp.PaymentMethod, nil
}

// Verify sends an OTP to confirm the customer owns the payment method.
//
// Required before first use (unless confirms_use is false). Returns verification
// status and OTP delivery details.
func (s *Service) Verify(ctx context.Context, paymentMethodID string) (*VerificationSession, error) {
	return s.VerifyWithParams(ctx, VerifyParams{
		PaymentMethodID: paymentMethodID,
		RequestMeta:     stablePaymentMethodRequestMeta("verify", paymentMethodID),
	})
}

func (s *Service) VerifyWithParams(ctx context.Context, params VerifyParams) (*VerificationSession, error) {
	var resp struct {
		Verification VerificationSession `json:"verification"`
	}
	if err := s.client.Do(ctx, "POST", "/payment_methods/verify", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Verification, nil
}

// ConfirmVerification submits the OTP to complete verification.
//
// Call this after Verify() once the customer provides their OTP.
// Returns the verified payment method.
func (s *Service) ConfirmVerification(ctx context.Context, params ConfirmVerificationParams) (*Resource, error) {
	var resp struct {
		PaymentMethod Resource `json:"payment_method"`
	}
	if err := s.client.Do(ctx, "POST", "/payment_methods/confirm_verification", params, &resp); err != nil {
		return nil, err
	}
	return &resp.PaymentMethod, nil
}

// Lookup retrieves payment method details by ID.
//
// Returns masked payment details, verification status, and enabled state.
func (s *Service) Lookup(ctx context.Context, paymentMethodID string) (*Resource, error) {
	var resp struct {
		PaymentMethod Resource `json:"payment_method"`
	}
	if err := s.client.Do(ctx, "POST", "/payment_methods/lookup", LookupParams{PaymentMethodID: paymentMethodID}, &resp); err != nil {
		return nil, err
	}
	return &resp.PaymentMethod, nil
}

// Page retrieves a paginated list of payment methods.
func (s *Service) Page(ctx context.Context, params PageParams) (*Page, error) {
	var resp struct {
		Page Page `json:"page"`
	}
	if err := s.client.Do(ctx, "POST", "/payment_methods/page", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Page, nil
}

// Update modifies mutable payment method fields.
func (s *Service) Update(ctx context.Context, payload any) (*Resource, error) {
	var resp struct {
		PaymentMethod Resource `json:"payment_method"`
	}
	if err := s.client.Do(ctx, "POST", "/payment_methods/update", payload, &resp); err != nil {
		return nil, err
	}
	return &resp.PaymentMethod, nil
}

// Activate marks a payment method active.
func (s *Service) Activate(ctx context.Context, paymentMethodID string) (*Resource, error) {
	return s.paymentMethodAction(ctx, "/payment_methods/activate", paymentMethodID)
}

// Disactivate marks a payment method inactive.
func (s *Service) Disactivate(ctx context.Context, paymentMethodID string) (*Resource, error) {
	return s.paymentMethodAction(ctx, "/payment_methods/disactivate", paymentMethodID)
}

// Deactivate is an alias for Disactivate.
func (s *Service) Deactivate(ctx context.Context, paymentMethodID string) (*Resource, error) {
	return s.Disactivate(ctx, paymentMethodID)
}

// Archive archives a payment method.
func (s *Service) Archive(ctx context.Context, paymentMethodID string) (*Resource, error) {
	return s.paymentMethodAction(ctx, "/payment_methods/archive", paymentMethodID)
}

// Unarchive unarchives a payment method.
func (s *Service) Unarchive(ctx context.Context, paymentMethodID string) (*Resource, error) {
	return s.paymentMethodAction(ctx, "/payment_methods/unarchive", paymentMethodID)
}

func (s *Service) paymentMethodAction(ctx context.Context, path, paymentMethodID string) (*Resource, error) {
	var resp struct {
		PaymentMethod Resource `json:"payment_method"`
	}
	if err := s.client.Do(ctx, "POST", path, ActionParams{PaymentMethodID: paymentMethodID}, &resp); err != nil {
		return nil, err
	}
	return &resp.PaymentMethod, nil
}

// Delete permanently removes a payment method.
//
// Deleted payment methods cannot be restored. Customers can delete their
// own payment methods—merchants cannot prevent this or restore them.
func (s *Service) Delete(ctx context.Context, paymentMethodID string) (*Deletion, error) {
	return s.DeleteWithParams(ctx, DeleteParams{
		PaymentMethodID: paymentMethodID,
		RequestMeta:     stablePaymentMethodRequestMeta("delete", paymentMethodID),
	})
}

func (s *Service) DeleteWithParams(ctx context.Context, params DeleteParams) (*Deletion, error) {
	var resp Deletion
	if err := s.client.Do(ctx, "POST", "/payment_methods/delete", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func stablePaymentMethodRequestMeta(action, paymentMethodID string) *request.Meta {
	return &request.Meta{IdempotencyKey: "payment_methods_" + action + "_" + paymentMethodID}
}

// Settings retrieves payment method acceptance configuration.
//
// Shows which payment types are enabled and whether OTP confirmation is required.
func (s *Service) Settings(ctx context.Context) (*Settings, error) {
	var resp struct {
		Settings Settings `json:"settings"`
	}
	if err := s.client.Do(ctx, "POST", "/payment_methods/settings", map[string]any{}, &resp); err != nil {
		return nil, err
	}
	return &resp.Settings, nil
}

// MobileMoneyParams describes a mobile money wallet.
//
// Used when creating orders or tokenizing payment methods with inline
// mobile money data instead of referencing a saved payment method.
type MobileMoneyParams struct {
	// Network is the mobile money network code.
	// Examples: "mtn", "vodafone", "airteltigo", "airtel", "telecel".
	Network MobileMoneyNetwork `json:"network"`

	// AccountNumber is the mobile money account phone number.
	// Must include country code. Example: "+233244123456"
	// Used as the payment source and for sending OTP verification codes.
	AccountNumber string `json:"account_number"`
}

// PaymentMethodData represents inline payment method data for one-time use.
//
// Use this to charge a payment method without saving it for future use.
// For repeat customers, tokenize the payment method first using
// PaymentMethods.Tokenize, then reference it by ID.
//
// Example (mobile money):
//
//	paymentData := &paymentmethod.Data{
//	    Type: paymentmethod.TypeMobileMoney,
//	    MobileMoney: &paymentmethod.MobileMoneyParams{
//	        Network: "mtn",
//	        AccountNumber: "+233244123456",
//	    },
//	}
type Data struct {
	// Type specifies the payment method category.
	Type Type `json:"type"`

	// MobileMoney provides mobile money wallet details.
	// Required when Type is PaymentMethodTypeMobileMoney.
	MobileMoney *MobileMoneyParams `json:"mobile_money,omitempty"`
}

// PaymentMethod represents a tokenized payment method.
//
// Payment methods are saved customer payment instruments that can be
// charged repeatedly without re-entering payment details.
//
// Payment methods must be verified before use (except when confirms_use
// is false). Verification sends an OTP to confirm the customer owns the
// payment instrument.
type Resource struct {
	// ID is the unique payment method identifier.
	// Starts with "pm_". Example: "pm_abc123def456"
	ID string `json:"id"`

	// CustomerID is the owning customer's ID.
	CustomerID string `json:"customer_id"`

	// Type is the payment method category.
	// Values: "mobile_money", "bank_account", "card", "motito"
	Type Type `json:"type"`

	// MobileMoney contains wallet details when Type is "mobile_money".
	MobileMoney *struct {
		AccountNumber string             `json:"account_number,omitempty"`
		Network       MobileMoneyNetwork `json:"network,omitempty"`
	} `json:"mobile_money,omitempty"`

	// BankAccount contains bank details when Type is "bank_account".
	BankAccount *struct {
		GhanaBankAccount *struct {
			Branch        string `json:"branch,omitempty"`
			Name          string `json:"name,omitempty"`
			AccountNumber string `json:"account_number,omitempty"`
			SortCode      string `json:"sort_code,omitempty"`
			SwiftCode     string `json:"swift_code,omitempty"`
		} `json:"ghana_bank_account,omitempty"`
		Type bankaccount.Type `json:"type,omitempty"`
	} `json:"bank_account,omitempty"`

	// Card contains card details when Type is "card".
	Card *struct {
		Brand     string  `json:"brand,omitempty"`
		ExpiresOn *string `json:"expires_on,omitempty"`
		Issuer    *struct {
			EmailAddress string `json:"email_address,omitempty"`
			Name         string `json:"name,omitempty"`
			PhoneNumber  string `json:"phone_number,omitempty"`
			Type         string `json:"type,omitempty"`
		} `json:"issuer,omitempty"`
		Owner *struct {
			EmailAddress string `json:"email_address,omitempty"`
			Name         string `json:"name,omitempty"`
			PhoneNumber  string `json:"phone_number,omitempty"`
		} `json:"owner,omitempty"`
		Type string `json:"type,omitempty"`
	} `json:"card,omitempty"`

	// Verification contains verification metadata when available.
	Verification *struct {
		CompletedAt *string `json:"completed_at,omitempty"`
		InitiatedAt *string `json:"initiated_at,omitempty"`
		Mechanism   string  `json:"mechanism,omitempty"`
		RequestID   string  `json:"request_id,omitempty"`
		Type        string  `json:"type,omitempty"`
	} `json:"verification,omitempty"`

	// CustomData holds arbitrary metadata attached to the payment method.
	CustomData map[string]string `json:"custom_data,omitempty"`

	// ExpiresOn is when this payment method expires (ISO 8601), if applicable.
	ExpiresOn *string `json:"expires_on,omitempty"`

	// CreatedAt is the tokenization timestamp (ISO 8601).
	CreatedAt string `json:"created_at"`

	// Verified indicates whether the payment method has been verified.
	// Unverified payment methods cannot be charged (unless confirms_use is false).
	Verified bool `json:"verified"`

	// VerifiedAt is the verification completion timestamp (ISO 8601).
	// Nil if not yet verified.
	VerifiedAt *string `json:"verified_at,omitempty"`
}

// TokenizePaymentMethodParams saves a payment method for future use.
//
// Tokenization stores customer payment details securely for repeat charges.
// The customer owns the payment method—only they can delete it.
//
// Example:
//
//	params := paymentmethod.TokenizeParams{
//	    CustomerID: "cu_abc123",
//	    PaymentMethodData: paymentmethod.Data{
//	        Type: paymentmethod.TypeMobileMoney,
//	        MobileMoney: &paymentmethod.MobileMoneyParams{
//	            Network: "mtn",
//	            AccountNumber: "+233244123456",
//	        },
//	    },
//	    VerifyImmediately: inttegro.Bool(true),
//	}
type TokenizeParams struct {
	// RequestMeta carries per-request controls such as idempotency.
	RequestMeta *request.Meta `json:"request_meta,omitempty"`

	// CustomerID is who owns this payment method (required).
	CustomerID string `json:"customer_id"`

	// PaymentMethodData contains the payment details to save (required).
	PaymentMethodData Data `json:"payment_method_data"`

	// VerifyImmediately triggers verification right after tokenization (optional).
	// If true, sends OTP immediately for customer to confirm ownership.
	// If false (default), must call Verify() separately before first use.
	VerifyImmediately *bool `json:"verify_immediately,omitempty"`
}

// VerifyPaymentMethodParams starts verification for a payment method.
//
// Sends an OTP to the payment method (phone number for mobile money)
// to confirm the customer owns it. Required before first use.
type VerifyParams struct {
	// RequestMeta carries per-request controls such as idempotency.
	RequestMeta *request.Meta `json:"request_meta,omitempty"`

	// PaymentMethodID is the payment method to verify (required).
	PaymentMethodID string `json:"payment_method_id"`
}

// PaymentMethodVerificationSession contains verification state and delivery details.
type VerificationSession struct {
	PaymentMethodID string                `json:"payment_method_id,omitempty"`
	Status          string                `json:"status,omitempty"`
	TokenSentAt     *string               `json:"token_sent_at,omitempty"`
	ExpiresAt       *string               `json:"expires_at,omitempty"`
	Delivery        *VerificationDelivery `json:"delivery,omitempty"`
}

// PaymentMethodVerificationDelivery describes where a verification token was sent.
type VerificationDelivery struct {
	Recipient string `json:"recipient,omitempty"`
	Channel   string `json:"channel,omitempty"`
	SenderID  string `json:"sender_id,omitempty"`
}

// ConfirmPaymentMethodVerificationParams submits the OTP to complete verification.
type ConfirmVerificationParams struct {
	// PaymentMethodID is the payment method being verified (required).
	PaymentMethodID string `json:"payment_method_id"`

	// Token is the OTP from the customer (required).
	// Typically 4-6 digits sent via SMS.
	Token string `json:"token"`
}

// LookupPaymentMethodParams specifies which payment method to retrieve.
type LookupParams struct {
	// PaymentMethodID is the payment method identifier (required).
	// Starts with "pm_". Example: "pm_abc123def456"
	PaymentMethodID string `json:"payment_method_id"`
}

// DeletePaymentMethodParams permanently removes a payment method.
//
// Deleted payment methods cannot be restored. Customers can delete
// their own payment methods; merchants cannot force re-enablement.
type DeleteParams struct {
	// RequestMeta carries per-request controls such as idempotency.
	RequestMeta *request.Meta `json:"request_meta,omitempty"`

	// PaymentMethodID is the payment method to delete (required).
	PaymentMethodID string `json:"payment_method_id"`
}

// PaymentMethodDeletion confirms deletion.
type Deletion struct {
	// Deleted indicates whether deletion succeeded.
	Deleted bool `json:"deleted"`

	// PaymentMethodID is the deleted payment method's ID.
	PaymentMethodID string `json:"payment_method_id,omitempty"`
}

// NewService constructs the resource service used by inttegro.Client.
func NewService(client transport.Client) *Service { return &Service{client: client} }
