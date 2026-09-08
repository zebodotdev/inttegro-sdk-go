package inttegro_test

import (
	"bytes"
	"encoding/json"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	inttegro "github.com/zebodotdev/inttegro-sdk-go/v6"
	"github.com/zebodotdev/inttegro-sdk-go/v6/payment"
	"github.com/zebodotdev/inttegro-sdk-go/v6/product"
	"github.com/zebodotdev/inttegro-sdk-go/v6/purchaseintent"
	"github.com/zebodotdev/inttegro-sdk-go/v6/refund"
)

const modulePath = "github.com/zebodotdev/inttegro-sdk-go/v6"

func TestClientUsesResourceOwnedServices(t *testing.T) {
	client := inttegro.NewClient("sk_test_resource_packages")
	services := []any{
		client.Apps, client.Balances, client.BalanceTransactions,
		client.Broadcasts, client.Chimes, client.Customers, client.Files,
		client.FileLinks, client.FileReferences, client.FinancialAccounts,
		client.MessageTemplates, client.Orders, client.Otp,
		client.PaymentMethods, client.Payouts, client.Prices, client.Products,
		client.PurchaseIntents, client.Refunds, client.Schedules, client.Keys,
		client.Spec, client.UploadRequests,
	}
	for index, service := range services {
		if service == nil {
			t.Fatalf("service %d is nil", index)
		}
		if reflect.TypeOf(service).Elem().PkgPath() == modulePath {
			t.Fatalf("service %T is still owned by the root package", service)
		}
	}
}

func TestPrimaryResourceTypesHavePackageIdentity(t *testing.T) {
	types := []struct {
		value any
		pkg   string
	}{
		{product.Product{}, modulePath + "/product"},
		{payment.Payment{}, modulePath + "/payment"},
		{purchaseintent.PurchaseIntent{}, modulePath + "/purchaseintent"},
		{refund.Refund{}, modulePath + "/refund"},
		{product.Type(""), modulePath + "/product"},
		{payment.Status(""), modulePath + "/payment"},
		{purchaseintent.Status(""), modulePath + "/purchaseintent"},
		{refund.Status(""), modulePath + "/refund"},
	}
	for _, item := range types {
		if got := reflect.TypeOf(item.value).PkgPath(); got != item.pkg {
			t.Errorf("%T package = %q, want %q", item.value, got, item.pkg)
		}
	}
}

func TestResourcePackagesNamePrimaryTypesAfterPackages(t *testing.T) {
	expected := map[string]string{
		"app":                "App",
		"balance":            "Balance",
		"balancetransaction": "BalanceTransaction",
		"broadcast":          "Broadcast",
		"chime":              "Chime",
		"customer":           "Customer",
		"file":               "File",
		"filelink":           "FileLink",
		"filereference":      "FileReference",
		"financialaccount":   "FinancialAccount",
		"invoice":            "Invoice",
		"messagetemplate":    "MessageTemplate",
		"order":              "Order",
		"payment":            "Payment",
		"paymentmethod":      "PaymentMethod",
		"payout":             "Payout",
		"price":              "Price",
		"product":            "Product",
		"purchaseintent":     "PurchaseIntent",
		"refund":             "Refund",
		"schedule":           "Schedule",
		"secretkey":          "SecretKey",
		"spec":               "Spec",
		"uploadrequest":      "UploadRequest",
	}

	for packageName, typeName := range expected {
		paths, err := filepath.Glob(filepath.Join(packageName, "*.go"))
		if err != nil {
			t.Fatal(err)
		}
		found := false
		for _, path := range paths {
			if strings.HasSuffix(path, "_test.go") {
				continue
			}
			file, err := parser.ParseFile(token.NewFileSet(), path, nil, 0)
			if err != nil {
				t.Fatal(err)
			}
			for _, declaration := range file.Decls {
				general, ok := declaration.(*ast.GenDecl)
				if !ok || general.Tok != token.TYPE {
					continue
				}
				for _, specification := range general.Specs {
					declaredName := specification.(*ast.TypeSpec).Name.Name
					if declaredName == "Resource" {
						t.Errorf("package %s exports generic type Resource", packageName)
					}
					if declaredName == typeName {
						found = true
					}
				}
			}
		}
		if !found {
			t.Errorf("package %s does not export primary type %s", packageName, typeName)
		}
	}
}

func TestResourceValuesPreserveWireValues(t *testing.T) {
	value := struct {
		ProductType          product.Type          `json:"product_type"`
		PaymentStatus        payment.Status        `json:"payment_status"`
		PurchaseIntentStatus purchaseintent.Status `json:"purchase_intent_status"`
		RefundStatus         refund.Status         `json:"refund_status"`
	}{
		ProductType:          product.TypeDigital,
		PaymentStatus:        payment.StatusPaid,
		PurchaseIntentStatus: purchaseintent.StatusActive,
		RefundStatus:         refund.StatusSucceeded,
	}
	encoded, err := json.Marshal(value)
	if err != nil {
		t.Fatal(err)
	}
	want := `{"product_type":"digital","payment_status":"paid","purchase_intent_status":"active","refund_status":"succeeded"}`
	if string(encoded) != want {
		t.Fatalf("wire values = %s, want %s", encoded, want)
	}
}

func TestProductionPackagesContainNoTypeAliases(t *testing.T) {
	err := filepath.WalkDir(".", func(path string, entry os.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if entry.IsDir() {
			if strings.HasPrefix(entry.Name(), ".") || entry.Name() == "_tools" {
				return filepath.SkipDir
			}
			return nil
		}
		if !strings.HasSuffix(path, ".go") || strings.HasSuffix(path, "_test.go") {
			return nil
		}
		file, err := parser.ParseFile(token.NewFileSet(), path, nil, 0)
		if err != nil {
			return err
		}
		for _, declaration := range file.Decls {
			general, ok := declaration.(*ast.GenDecl)
			if !ok || general.Tok != token.TYPE {
				continue
			}
			for _, specification := range general.Specs {
				typeSpec := specification.(*ast.TypeSpec)
				if typeSpec.Assign.IsValid() {
					t.Errorf("production type alias %s found in %s", typeSpec.Name.Name, path)
				}
			}
		}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestProductionEnumConstantsAreGroupedByType(t *testing.T) {
	type declarationLocation struct {
		path string
		line int
	}

	declarations := make(map[string]declarationLocation)
	err := filepath.WalkDir(".", func(path string, entry os.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if entry.IsDir() {
			if strings.HasPrefix(entry.Name(), ".") || entry.Name() == "_tools" {
				return filepath.SkipDir
			}
			return nil
		}
		if !strings.HasSuffix(path, ".go") || strings.HasSuffix(path, "_test.go") {
			return nil
		}

		fileSet := token.NewFileSet()
		file, err := parser.ParseFile(fileSet, path, nil, 0)
		if err != nil {
			return err
		}
		for _, declaration := range file.Decls {
			general, ok := declaration.(*ast.GenDecl)
			if !ok || general.Tok != token.CONST {
				continue
			}
			if general.Lparen.IsValid() && len(general.Specs) == 1 {
				position := fileSet.Position(general.Pos())
				t.Errorf("singleton parenthesized const block in %s:%d", path, position.Line)
			}

			location := declarationLocation{path: path, line: fileSet.Position(general.Pos()).Line}
			for _, specification := range general.Specs {
				valueSpec := specification.(*ast.ValueSpec)
				typeName, ok := valueSpec.Type.(*ast.Ident)
				if !ok || !ast.IsExported(typeName.Name) {
					continue
				}
				key := filepath.Dir(path) + ":" + typeName.Name
				if previous, found := declarations[key]; found && previous != location {
					t.Errorf(
						"constants of type %s are split between %s:%d and %s:%d",
						typeName.Name, previous.path, previous.line, path, location.line,
					)
					continue
				}
				declarations[key] = location
			}
		}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestProductionPackagesAvoidMonolithicSourceFiles(t *testing.T) {
	entries, err := os.ReadDir(".")
	if err != nil {
		t.Fatal(err)
	}
	for _, entry := range entries {
		if !entry.IsDir() || strings.HasPrefix(entry.Name(), ".") || entry.Name() == "internal" {
			continue
		}
		paths, err := filepath.Glob(filepath.Join(entry.Name(), "*.go"))
		if err != nil {
			t.Fatal(err)
		}
		fileCount := 0
		totalLines := 0
		for _, path := range paths {
			if strings.HasSuffix(path, "_test.go") {
				continue
			}
			contents, err := os.ReadFile(path)
			if err != nil {
				t.Fatal(err)
			}
			lines := bytes.Count(contents, []byte{'\n'})
			fileCount++
			totalLines += lines
			if lines > 300 {
				t.Errorf("%s has %d lines; split files that exceed 300 lines by concern", path, lines)
			}
		}
		if totalLines >= 100 && fileCount < 2 {
			t.Errorf(
				"package %s has %d production lines in one file; split it by concern",
				entry.Name(), totalLines,
			)
		}
	}
}

func TestResourcePackagesDoNotImportRoot(t *testing.T) {
	entries, err := os.ReadDir(".")
	if err != nil {
		t.Fatal(err)
	}
	for _, entry := range entries {
		if !entry.IsDir() || strings.HasPrefix(entry.Name(), ".") || entry.Name() == "internal" {
			continue
		}
		files, err := filepath.Glob(filepath.Join(entry.Name(), "*.go"))
		if err != nil {
			t.Fatal(err)
		}
		for _, path := range files {
			if strings.HasSuffix(path, "_test.go") {
				continue
			}
			file, err := parser.ParseFile(token.NewFileSet(), path, nil, parser.ImportsOnly)
			if err != nil {
				t.Fatal(err)
			}
			for _, importSpec := range file.Imports {
				if strings.Trim(importSpec.Path.Value, `"`) == modulePath {
					t.Errorf("%s imports the root package", path)
				}
			}
		}
	}
}
