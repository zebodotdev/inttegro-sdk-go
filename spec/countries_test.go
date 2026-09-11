package spec

import (
	"encoding/json"
	"testing"
)

func TestCountrySpecificationsDecodeContractShape(t *testing.T) {
	var specifications CountrySpecifications
	if err := json.Unmarshal([]byte(`{
		"gh": {
			"country_code": "gh",
			"legal_entity_types": ["company"],
			"financial_account_types": ["bank_account"],
			"id_document_types": ["passport"]
		}
	}`), &specifications); err != nil {
		t.Fatalf("decode country specifications: %v", err)
	}

	ghana := specifications["gh"]
	if got := ghana.LegalEntityTypes[0]; got != "company" {
		t.Fatalf("legal entity type = %q, want company", got)
	}
	if got := ghana.FinancialAccountTypes[0]; got != "bank_account" {
		t.Fatalf("financial account type = %q, want bank_account", got)
	}
	if got := ghana.IDDocumentTypes[0]; got != "passport" {
		t.Fatalf("ID document type = %q, want passport", got)
	}
}

func TestSpecCompatibilityAlias(t *testing.T) {
	var legacy Spec = CountrySpecification{CountryCode: "gh"}
	if legacy.CountryCode != "gh" {
		t.Fatalf("country code = %q, want gh", legacy.CountryCode)
	}
}
