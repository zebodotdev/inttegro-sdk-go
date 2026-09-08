package inttegro_test

import (
	"encoding/json"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strings"
	"testing"

	inttegro "github.com/zebodotdev/inttegro-sdk-go/v4"
	"github.com/zebodotdev/inttegro-sdk-go/v4/app"
	"github.com/zebodotdev/inttegro-sdk-go/v4/balance"
	"github.com/zebodotdev/inttegro-sdk-go/v4/balancetransaction"
	"github.com/zebodotdev/inttegro-sdk-go/v4/broadcast"
	"github.com/zebodotdev/inttegro-sdk-go/v4/chime"
	"github.com/zebodotdev/inttegro-sdk-go/v4/customer"
	sdkfile "github.com/zebodotdev/inttegro-sdk-go/v4/file"
	"github.com/zebodotdev/inttegro-sdk-go/v4/filelink"
	"github.com/zebodotdev/inttegro-sdk-go/v4/filereference"
	"github.com/zebodotdev/inttegro-sdk-go/v4/financialaccount"
	"github.com/zebodotdev/inttegro-sdk-go/v4/messagetemplate"
	"github.com/zebodotdev/inttegro-sdk-go/v4/order"
	"github.com/zebodotdev/inttegro-sdk-go/v4/otp"
	"github.com/zebodotdev/inttegro-sdk-go/v4/paymentmethod"
	"github.com/zebodotdev/inttegro-sdk-go/v4/payout"
	"github.com/zebodotdev/inttegro-sdk-go/v4/price"
	"github.com/zebodotdev/inttegro-sdk-go/v4/product"
	"github.com/zebodotdev/inttegro-sdk-go/v4/purchaseintent"
	"github.com/zebodotdev/inttegro-sdk-go/v4/refund"
	"github.com/zebodotdev/inttegro-sdk-go/v4/schedule"
	"github.com/zebodotdev/inttegro-sdk-go/v4/secretkey"
	"github.com/zebodotdev/inttegro-sdk-go/v4/spec"
	"github.com/zebodotdev/inttegro-sdk-go/v4/uploadrequest"
)

func TestResourcePackageServicesMatchClient(t *testing.T) {
	client := inttegro.NewClient("sk_test_resource_packages")
	services := []struct {
		name    string
		service any
	}{
		{"app", (*app.Service)(client.Apps)},
		{"balance", (*balance.Service)(client.Balances)},
		{"balance transaction", (*balancetransaction.Service)(client.BalanceTransactions)},
		{"broadcast", (*broadcast.Service)(client.Broadcasts)},
		{"chime", (*chime.Service)(client.Chimes)},
		{"customer", (*customer.Service)(client.Customers)},
		{"file", (*sdkfile.Service)(client.Files)},
		{"file link", (*filelink.Service)(client.FileLinks)},
		{"file reference", (*filereference.Service)(client.FileReferences)},
		{"financial account", (*financialaccount.Service)(client.FinancialAccounts)},
		{"message template", (*messagetemplate.Service)(client.MessageTemplates)},
		{"order", (*order.Service)(client.Orders)},
		{"otp", (*otp.Service)(client.Otp)},
		{"payment method", (*paymentmethod.Service)(client.PaymentMethods)},
		{"payout", (*payout.Service)(client.Payouts)},
		{"price", (*price.Service)(client.Prices)},
		{"product", (*product.Service)(client.Products)},
		{"purchase intent", (*purchaseintent.Service)(client.PurchaseIntents)},
		{"refund", (*refund.Service)(client.Refunds)},
		{"schedule", (*schedule.Service)(client.Schedules)},
		{"secret key", (*secretkey.Service)(client.Keys)},
		{"spec", (*spec.Service)(client.Spec)},
		{"upload request", (*uploadrequest.Service)(client.UploadRequests)},
	}
	for _, service := range services {
		if service.service == nil {
			t.Errorf("%s service is nil", service.name)
		}
	}
}

func TestResourcePackageNamesPreserveWireValues(t *testing.T) {
	value := struct {
		ProductType          product.Type          `json:"product_type"`
		PurchaseIntentStatus purchaseintent.Status `json:"purchase_intent_status"`
		RefundStatus         refund.Status         `json:"refund_status"`
		OrderStatus          order.Status          `json:"order_status"`
	}{
		ProductType:          product.TypeDigital,
		PurchaseIntentStatus: purchaseintent.StatusActive,
		RefundStatus:         refund.StatusSucceeded,
		OrderStatus:          order.StatusCompleted,
	}
	got, err := json.Marshal(value)
	if err != nil {
		t.Fatal(err)
	}
	want := `{"product_type":"digital","purchase_intent_status":"active","refund_status":"succeeded","order_status":"completed"}`
	if string(got) != want {
		t.Fatalf("resource package wire values = %s, want %s", got, want)
	}
}

func TestResourcePackagesPreserveV4TypeIdentity(t *testing.T) {
	pairs := []struct {
		name   string
		root   reflect.Type
		scoped reflect.Type
	}{
		{"product", reflect.TypeOf(inttegro.Product{}), reflect.TypeOf(product.Resource{})},
		{"purchase intent", reflect.TypeOf(inttegro.PurchaseIntent{}), reflect.TypeOf(purchaseintent.Resource{})},
		{"refund", reflect.TypeOf(inttegro.Refund{}), reflect.TypeOf(refund.Resource{})},
		{"order", reflect.TypeOf(inttegro.Order{}), reflect.TypeOf(order.Resource{})},
	}
	for _, pair := range pairs {
		if pair.root != pair.scoped {
			t.Errorf("%s package introduced a distinct v4 type", pair.name)
		}
	}
}

func TestEveryRootResourceTypeHasPackageScopedName(t *testing.T) {
	rootTypes := exportedTypeNames(t, ".")
	packageAliases := rootSelectorsInResourcePackages(t)
	crossCutting := map[string]bool{
		"APIError": true, "APIErrorReportContext": true,
		"BankAccountConfig": true, "BankAccountOwner": true,
		"BankAccountOwnerAddress": true, "BankAccountType": true,
		"Client": true, "ClientOption": true,
		"ErrorReport": true, "ErrorReporter": true, "ErrorReportingPolicy": true,
		"GhanaBankAccount": true, "HTTPReportContext": true,
		"RequestMeta": true, "RequestOption": true,
		"SDKReportContext": true, "TraceReportContext": true,
		"WalletConfig": true, "WalletMobileMoney": true, "WalletType": true,
	}
	var missing []string
	for name := range rootTypes {
		if !crossCutting[name] && !packageAliases[name] {
			missing = append(missing, name)
		}
	}
	sort.Strings(missing)
	if len(missing) != 0 {
		t.Fatalf("root resource types missing package-scoped names: %s", strings.Join(missing, ", "))
	}
}

func TestEveryRootEnumConstantHasPackageScopedName(t *testing.T) {
	enumConstants := exportedConstNames(t, "enums.go")
	packageAliases := rootSelectorsInResourcePackages(t)
	var missing []string
	for name := range enumConstants {
		if !packageAliases[name] {
			missing = append(missing, name)
		}
	}
	sort.Strings(missing)
	if len(missing) != 0 {
		t.Fatalf("root enum constants missing package-scoped names: %s", strings.Join(missing, ", "))
	}
}

func exportedTypeNames(t *testing.T, dir string) map[string]bool {
	t.Helper()
	names := make(map[string]bool)
	for _, file := range goFiles(t, dir) {
		parsed, err := parser.ParseFile(token.NewFileSet(), file, nil, 0)
		if err != nil {
			t.Fatal(err)
		}
		for _, declaration := range parsed.Decls {
			general, ok := declaration.(*ast.GenDecl)
			if !ok || general.Tok != token.TYPE {
				continue
			}
			for _, specification := range general.Specs {
				typeSpec := specification.(*ast.TypeSpec)
				if typeSpec.Name.IsExported() {
					names[typeSpec.Name.Name] = true
				}
			}
		}
	}
	return names
}

func exportedConstNames(t *testing.T, file string) map[string]bool {
	t.Helper()
	parsed, err := parser.ParseFile(token.NewFileSet(), file, nil, 0)
	if err != nil {
		t.Fatal(err)
	}
	names := make(map[string]bool)
	for _, declaration := range parsed.Decls {
		general, ok := declaration.(*ast.GenDecl)
		if !ok || general.Tok != token.CONST {
			continue
		}
		for _, specification := range general.Specs {
			valueSpec := specification.(*ast.ValueSpec)
			for _, name := range valueSpec.Names {
				if name.IsExported() {
					names[name.Name] = true
				}
			}
		}
	}
	return names
}

func rootSelectorsInResourcePackages(t *testing.T) map[string]bool {
	t.Helper()
	selectors := make(map[string]bool)
	resourcePackages := []string{
		"app", "balance", "balancetransaction", "bankaccount", "broadcast",
		"checkout", "chime", "customer", "file", "filelink", "filereference",
		"financialaccount", "invoice", "messagetemplate", "order", "otp",
		"payment", "paymentmethod", "payout", "price", "product",
		"purchaseintent", "refund", "schedule", "secretkey", "spec",
		"uploadrequest", "wallet",
	}
	for _, resourcePackage := range resourcePackages {
		for _, file := range goFiles(t, resourcePackage) {
			parsed, err := parser.ParseFile(token.NewFileSet(), file, nil, 0)
			if err != nil {
				t.Fatal(err)
			}
			ast.Inspect(parsed, func(node ast.Node) bool {
				selector, ok := node.(*ast.SelectorExpr)
				if !ok {
					return true
				}
				packageName, ok := selector.X.(*ast.Ident)
				if ok && packageName.Name == "inttegro" {
					selectors[selector.Sel.Name] = true
				}
				return true
			})
		}
	}
	return selectors
}

func goFiles(t *testing.T, dir string) []string {
	t.Helper()
	entries, err := os.ReadDir(dir)
	if err != nil {
		t.Fatal(err)
	}
	var files []string
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".go") ||
			strings.HasSuffix(entry.Name(), "_test.go") {
			continue
		}
		files = append(files, filepath.Join(dir, entry.Name()))
	}
	return files
}
