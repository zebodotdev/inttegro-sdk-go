package inttegro_test

import (
	"encoding/json"
	"testing"

	inttegro "github.com/zebodotdev/inttegro-sdk-go/v4"
	"github.com/zebodotdev/inttegro-sdk-go/v4/bankaccounts"
	"github.com/zebodotdev/inttegro-sdk-go/v4/paymentmethods"
	"github.com/zebodotdev/inttegro-sdk-go/v4/wallets"
)

func TestFinancialAccountVariantsUseFocusedPackages(t *testing.T) {
	params := inttegro.FinancialAccountCreateParams{
		Type: inttegro.FinancialAccountTypeWallet,
		Wallet: &wallets.Config{
			Type: wallets.TypeMobileMoney,
			MobileMoney: &wallets.MobileMoney{
				AccountNumber: "233200000000",
				Network:       paymentmethods.MobileMoneyNetworkMTN,
			},
		},
		BankAccount: &bankaccounts.Config{
			Type: bankaccounts.TypeGhanaBankAccount,
			GhanaBankAccount: &bankaccounts.GhanaBankAccount{
				Number: "0123456789",
			},
		},
	}

	payload, err := json.Marshal(params)
	if err != nil {
		t.Fatal(err)
	}
	if len(payload) == 0 {
		t.Fatal("expected a serialized financial account request")
	}
}
