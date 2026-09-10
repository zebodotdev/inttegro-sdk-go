package paymentmethod

import (
	"time"

	"github.com/zebodotdev/inttegro-sdk-go/v6/bankaccount"
)

// MobileMoneyNetwork identifies a supported mobile money network.
type MobileMoneyNetwork string

const (
	MobileMoneyNetworkAirtel   MobileMoneyNetwork = "airtel"
	MobileMoneyNetworkMTN      MobileMoneyNetwork = "mtn"
	MobileMoneyNetworkTelecel  MobileMoneyNetwork = "telecel"
	MobileMoneyNetworkVodafone MobileMoneyNetwork = "vodafone"
)

type MobileMoney struct {
	Network       MobileMoneyNetwork `json:"network"`
	AccountNumber string             `json:"account_number"`
	Last4         string             `json:"last4"`
}

type BankAccount struct {
	Type             bankaccount.Type  `json:"type"`
	GhanaBankAccount *GhanaBankAccount `json:"ghana_bank_account,omitempty"`
}

type GhanaBankAccount struct {
	AccountNumber string `json:"account_number"`
	Branch        string `json:"branch,omitempty"`
	Name          string `json:"name,omitempty"`
	SortCode      string `json:"sort_code,omitempty"`
	SwiftCode     string `json:"swift_code,omitempty"`
}

// Card is an explicit marker. Card credentials are never returned.
type Card struct{}

type Owner struct {
	Name    string   `json:"name"`
	Address *Address `json:"address,omitempty"`
}

type Address struct {
	City        string `json:"city,omitempty"`
	Country     string `json:"country"`
	Line1       string `json:"line_1,omitempty"`
	Line2       string `json:"line_2,omitempty"`
	Name        string `json:"name,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
	PostCode    string `json:"post_code,omitempty"`
	Region      string `json:"region,omitempty"`
}

type Verification struct {
	CompletedAt *time.Time `json:"completed_at,omitempty"`
	InitiatedAt time.Time  `json:"initiated_at"`
	Mechanism   string     `json:"mechanism,omitempty"`
	RequestID   string     `json:"request_id"`
	Type        string     `json:"type"`
}

type Supplied struct {
	AttemptID    string    `json:"attempt_id,omitempty"`
	By           string    `json:"by"`
	Channel      string    `json:"channel,omitempty"`
	ResourceID   string    `json:"resource_id,omitempty"`
	ResourceType string    `json:"resource_type,omitempty"`
	SuppliedAt   time.Time `json:"supplied_at"`
}
