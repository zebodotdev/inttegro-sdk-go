package inttegro

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"testing"
)

var openAPICapabilityURLPaths = map[string]bool{
	"/file_links/open":        true,
	"/upload_requests/upload": true,
}

func TestSDKPathsCoverOpenAPI(t *testing.T) {
	specPaths, err := readOpenAPIPaths(openAPISpecPath())
	if err != nil {
		t.Fatal(err)
	}
	sdkPaths := recordSDKPaths(t)

	var missing []string
	for _, path := range specPaths {
		if openAPICapabilityURLPaths[path] {
			continue
		}
		if !sdkPaths[path] {
			missing = append(missing, path)
		}
	}
	if len(missing) > 0 {
		t.Fatalf("Go SDK missing OpenAPI paths from %s:\n%s", openAPISpecPath(), strings.Join(missing, "\n"))
	}
}

func openAPISpecPath() string {
	if path := os.Getenv("INTTEGRO_OPENAPI_SPEC"); path != "" {
		return path
	}
	return filepath.Clean("../../openapi/commerce.yml")
}

func readOpenAPIPaths(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var paths []string
	inPaths := false
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		trimmed := strings.TrimSpace(line)
		if trimmed == "" || trimmed == "---" {
			continue
		}
		if !inPaths {
			if trimmed == "paths:" {
				inPaths = true
			}
			continue
		}
		if !strings.HasPrefix(line, " ") && !strings.HasPrefix(line, "\t") {
			break
		}
		if strings.HasPrefix(line, "    /") {
			idx := strings.Index(trimmed, ":")
			if idx <= 0 {
				return nil, fmt.Errorf("malformed OpenAPI path line: %q", line)
			}
			paths = append(paths, trimmed[:idx])
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	if !inPaths {
		return nil, fmt.Errorf("OpenAPI paths block not found in %s", path)
	}
	if len(paths) == 0 {
		return nil, fmt.Errorf("OpenAPI paths block in %s was empty", path)
	}
	return paths, nil
}

func recordSDKPaths(t *testing.T) map[string]bool {
	t.Helper()
	var paths []string
	client, close := newTestClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		paths = append(paths, r.URL.Path)
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(openAPICoverageResponse())
	}))
	if client == nil {
		return map[string]bool{}
	}
	defer close()

	ctx := context.Background()
	tempFile := filepath.Join(t.TempDir(), "upload.txt")
	if err := os.WriteFile(tempFile, []byte("upload"), 0o600); err != nil {
		t.Fatal(err)
	}

	check := func(err error) {
		t.Helper()
		if err != nil {
			t.Fatal(err)
		}
	}
	closeDownload := func(download *FileDownload, err error) {
		t.Helper()
		check(err)
		_, _ = io.Copy(io.Discard, download)
		check(download.Close())
	}

	_, err := client.Otp.Initiate(ctx, map[string]any{"recipient": "+233"})
	check(err)
	_, err = client.Otp.Verify(ctx, map[string]any{"transaction_id": "otp_1", "token": "123456"})
	check(err)
	_, err = client.Otp.Lookup(ctx, map[string]any{"transaction_id": "otp_1"})
	check(err)
	_, err = client.Otp.Cancel(ctx, map[string]any{"transaction_id": "otp_1"})
	check(err)

	_, err = client.Chimes.Send(ctx, SendChimeParams{FullMessage: "hello"})
	check(err)
	_, err = client.Chimes.Lookup(ctx, "ch_1")
	check(err)
	_, err = client.Chimes.Page(ctx, ChimePageParams{PageNumber: 1, PageSize: 20})
	check(err)
	_, err = client.Chimes.Schedule(ctx, ScheduleChimeParams{FullMessage: "later"})
	check(err)
	_, err = client.Chimes.Broadcast(ctx, BroadcastChimeParams{MessageTemplate: "hello"})
	check(err)
	_, err = client.Schedules.Lookup(ctx, "sch_1")
	check(err)
	_, err = client.Schedules.Cancel(ctx, "sch_1")
	check(err)
	_, err = client.Broadcasts.Lookup(ctx, "brc_1")
	check(err)
	_, err = client.Broadcasts.Cancel(ctx, "brc_1")
	check(err)

	_, err = client.MessageTemplates.Create(ctx, MessageTemplateCreateParams{
		Name:    "welcome_sms",
		Channel: "sms",
		Purpose: "marketing",
		SMS:     &MessageTemplateSMSContent{MessageTemplate: "Welcome {{name}}"},
	})
	check(err)
	_, err = client.MessageTemplates.Update(ctx, MessageTemplateUpdateParams{ID: "mtpl_1", Name: "welcome_sms"})
	check(err)
	_, err = client.MessageTemplates.Publish(ctx, "mtpl_1")
	check(err)
	_, err = client.MessageTemplates.Archive(ctx, "mtpl_1")
	check(err)
	_, err = client.MessageTemplates.Lookup(ctx, "mtpl_1")
	check(err)
	_, err = client.MessageTemplates.Page(ctx, MessageTemplatePageParams{Page: 1, Size: 20})
	check(err)
	_, err = client.MessageTemplates.RenderPreview(ctx, MessageTemplateRenderPreviewParams{
		MessageTemplate: MessageTemplateReference{TemplateID: "mtpl_1"},
	})
	check(err)

	_, err = client.Customers.Create(ctx, CreateCustomerParams{Name: "Jane Doe"})
	check(err)
	_, err = client.Customers.Lookup(ctx, "cu_1")
	check(err)
	_, err = client.Customers.Page(ctx, PageCustomersParams{PageNumber: 1, PageSize: 20})
	check(err)

	_, err = client.Orders.Create(ctx, OrderCreateParams{Number: "ORDER-1"})
	check(err)
	_, err = client.Orders.New(ctx, OrderCreateParams{Number: "ORDER-2"})
	check(err)
	_, err = client.Orders.Lookup(ctx, "or_1")
	check(err)
	_, err = client.Orders.Update(ctx, map[string]any{"order_id": "or_1", "number": "ORDER-1A"})
	check(err)
	_, err = client.Orders.Pay(ctx, OrderPayParams{OrderID: "or_1"})
	check(err)
	_, err = client.Orders.ConfirmPayment(ctx, OrderConfirmParams{OrderID: "or_1", Token: "123456"})
	check(err)
	check(client.Orders.RequestConfirmation(ctx, "or_1"))
	_, err = client.Orders.Cancel(ctx, "or_1")
	check(err)
	_, err = client.Orders.Finalize(ctx, "or_1")
	check(err)
	_, err = client.Orders.Complete(ctx, OrderCompleteParams{OrderID: "or_1"})
	check(err)
	_, err = client.Orders.SendInvoice(ctx, OrderSendInvoiceParams{OrderID: "or_1"})
	check(err)
	_, err = client.Orders.SendReceipt(ctx, OrderSendReceiptParams{OrderID: "or_1"})
	check(err)
	_, err = client.Orders.Page(ctx, OrderPageParams{PageNumber: 1, PageSize: 20})
	check(err)
	refundRequest := CreateRefundRequest{
		OrderID: "or_1",
		Reason:  RefundReasonRequestedByCustomer,
		LineItems: []CreateRefundLineItem{{
			OrderLineItemID: "oli_1",
			RefundAmount:    Money{Currency: "ghs", Value: 100},
		}},
	}
	_, err = client.Refunds.Create(ctx, refundRequest)
	check(err)
	_, err = client.Orders.Refund(ctx, refundRequest)
	check(err)
	_, err = client.Refunds.Cancel(ctx, CancelRefundRequest{RefundID: "rf_1"})
	check(err)
	_, err = client.Refunds.Lookup(ctx, LookupRefundRequest{RefundID: "rf_1"})
	check(err)
	_, err = client.Refunds.Page(ctx, PageRefundsRequest{PageNumber: 1, PageSize: 20})
	check(err)

	_, err = client.Apps.Create(ctx, map[string]any{"name": "App"})
	check(err)
	_, err = client.Apps.Lookup(ctx)
	check(err)
	_, err = client.Apps.Update(ctx, map[string]any{"alias": "app"})
	check(err)
	_, err = client.Keys.Generate(ctx, GenerateSecretKeyParams{Label: "Integration"})
	check(err)
	_, err = client.Keys.Page(ctx, PageSecretKeysParams{Page: 1, Size: 20})
	check(err)
	_, err = client.Keys.Lookup(ctx, "sk_1")
	check(err)
	_, err = client.Keys.Update(ctx, UpdateSecretKeyParams{SecretKeyID: "sk_1", Label: "Renamed"})
	check(err)
	_, err = client.Keys.Destroy(ctx, "sk_1")
	check(err)
	_, err = client.Keys.Usage(ctx, SecretKeyUsageParams{SecretKeyID: "sk_1", Page: 1, Size: 20})
	check(err)

	_, err = client.FinancialAccounts.Create(ctx, FinancialAccountCreateParams{Label: "Primary"})
	check(err)
	_, err = client.FinancialAccounts.Lookup(ctx, "fa_1")
	check(err)
	_, err = client.FinancialAccounts.Archive(ctx, map[string]any{"account_id": "fa_1"})
	check(err)
	_, err = client.FinancialAccounts.Page(ctx, PageFinancialAccountsParams{PageNumber: 1, PageSize: 20})
	check(err)
	_, err = client.FinancialAccounts.Verify(ctx, map[string]any{"account_id": "fa_1"})
	check(err)
	_, err = client.FinancialAccounts.Connect(ctx, FinancialAccountCreateParams{Label: "Primary"})
	check(err)
	_, err = client.FinancialAccounts.Update(ctx, map[string]any{"account_id": "fa_1", "label": "Updated"})
	check(err)
	_, err = client.FinancialAccounts.EnablePush(ctx, "fa_1")
	check(err)
	_, err = client.FinancialAccounts.DisablePush(ctx, "fa_1")
	check(err)
	_, err = client.FinancialAccounts.Disconnect(ctx, FinancialAccountDisconnectParams{AccountID: "fa_1"})
	check(err)
	_, err = client.FinancialAccounts.Reconnect(ctx, "fa_1")
	check(err)
	_, err = client.FinancialAccounts.EnablePull(ctx, "fa_1")
	check(err)
	_, err = client.FinancialAccounts.DisablePull(ctx, "fa_1")
	check(err)

	_, err = client.Balances.Get(ctx)
	check(err)
	_, err = client.BalanceTransactions.Lookup(ctx, "bt_1")
	check(err)
	_, err = client.BalanceTransactions.Page(ctx, BalanceTransactionPageParams{PageNumber: 1, PageSize: 20})
	check(err)

	_, err = client.Payouts.Schedule(ctx, SchedulePayoutParams{DestinationID: "fa_1", MaxAmount: 100, Reference: "PAYOUT-1"})
	check(err)
	_, err = client.Payouts.Lookup(ctx, "po_1")
	check(err)
	_, err = client.Payouts.SetDestinations(ctx, map[string]string{"ghs": "fa_1"})
	check(err)
	_, err = client.Payouts.Settings(ctx)
	check(err)
	_, err = client.Payouts.DisableAutomatic(ctx)
	check(err)
	_, err = client.Payouts.EnableAutomatic(ctx)
	check(err)
	_, err = client.Payouts.EnableFX(ctx)
	check(err)
	_, err = client.Payouts.DisableFX(ctx)
	check(err)
	_, err = client.Payouts.Page(ctx, PayoutPageParams{PageNumber: 1, PageSize: 20})
	check(err)
	_, err = client.Payouts.Cancel(ctx, "po_1")
	check(err)

	_, err = client.Files.Create(ctx, FileCreateParams{File: tempFile, Purpose: "identity"})
	check(err)
	_, err = client.Files.Lookup(ctx, "file_1")
	check(err)
	_, err = client.Files.Page(ctx, FilePageParams{PageNumber: 1, PageSize: 20})
	check(err)
	closeDownload(client.Files.Contents(ctx, FileContentsParams{FileID: "file_1"}))
	_, err = client.Files.Delete(ctx, "file_1")
	check(err)
	_, _, err = client.FileLinks.Create(ctx, FileLinkCreateParams{FileID: "file_1"})
	check(err)
	_, err = client.FileLinks.Lookup(ctx, "fl_1")
	check(err)
	_, err = client.FileLinks.Page(ctx, FileLinkPageParams{PageNumber: 1, PageSize: 20})
	check(err)
	_, err = client.FileLinks.Revoke(ctx, FileLinkRevokeParams{ID: "fl_1"})
	check(err)
	_, err = client.UploadRequests.Create(ctx, UploadRequestCreateParams{Purpose: "identity"})
	check(err)
	_, err = client.UploadRequests.Lookup(ctx, "ur_1")
	check(err)
	_, err = client.UploadRequests.Page(ctx, UploadRequestPageParams{PageNumber: 1, PageSize: 20})
	check(err)
	_, err = client.UploadRequests.Cancel(ctx, UploadRequestCancelParams{ID: "ur_1"})
	check(err)
	_, err = client.FileReferences.Reconcile(ctx, FileReferenceReconcileParams{ResourceType: "product", ResourceID: "prod_1"})
	check(err)

	_, err = client.PaymentMethods.Tokenize(ctx, TokenizePaymentMethodParams{CustomerID: "cu_1"})
	check(err)
	_, err = client.PaymentMethods.Verify(ctx, "pm_1")
	check(err)
	_, err = client.PaymentMethods.Lookup(ctx, "pm_1")
	check(err)
	_, err = client.PaymentMethods.Page(ctx, PaymentMethodPageParams{PageNumber: 1, PageSize: 20})
	check(err)
	_, err = client.PaymentMethods.Update(ctx, map[string]any{"payment_method_id": "pm_1", "active": true})
	check(err)
	_, err = client.PaymentMethods.Activate(ctx, "pm_1")
	check(err)
	_, err = client.PaymentMethods.Disactivate(ctx, "pm_1")
	check(err)
	_, err = client.PaymentMethods.Archive(ctx, "pm_1")
	check(err)
	_, err = client.PaymentMethods.Unarchive(ctx, "pm_1")
	check(err)
	_, err = client.PaymentMethods.Delete(ctx, "pm_1")
	check(err)
	_, err = client.PaymentMethods.Settings(ctx)
	check(err)

	_, err = client.Products.Create(ctx, CreateProductParams{Type: "physical", Name: "Product"})
	check(err)
	_, err = client.Products.AddPrice(ctx, AddProductPriceParams{
		ProductID: "prod_1",
		Amount:    ProductPriceAmount{Currency: "ghs", Value: 100},
	})
	check(err)
	_, err = client.Products.Lookup(ctx, "prod_1")
	check(err)
	_, err = client.Products.Update(ctx, UpdateProductParams{ProductID: "prod_1", Name: "Updated"})
	check(err)
	_, err = client.Products.Publish(ctx, "prod_1")
	check(err)
	_, err = client.Products.Unpublish(ctx, "prod_1")
	check(err)
	_, err = client.Products.Archive(ctx, "prod_1")
	check(err)
	_, err = client.Products.Page(ctx, PageProductsParams{PageNumber: 1, PageSize: 20})
	check(err)

	_, err = client.PurchaseIntents.Create(ctx, CreatePurchaseIntentParams{
		ProductID: "prod_1",
		PriceID:   "pr_1",
		Quantity:  PurchaseIntentQuantity{Min: 1},
	})
	check(err)
	_, err = client.PurchaseIntents.Update(ctx, UpdatePurchaseIntentParams{
		ID:       "sale_1",
		Quantity: &PurchaseIntentQuantity{Min: 1},
	})
	check(err)
	_, err = client.PurchaseIntents.Cancel(ctx, "sale_1")
	check(err)
	_, err = client.PurchaseIntents.Lookup(ctx, "sale_1")
	check(err)
	_, err = client.PurchaseIntents.Page(ctx, PagePurchaseIntentsParams{PageNumber: 1, PageSize: 20})
	check(err)

	_, err = client.Prices.Create(ctx, CreatePriceParams{Currency: "ghs", Amount: 100})
	check(err)
	_, err = client.Prices.Lookup(ctx, "pr_1")
	check(err)
	_, err = client.Prices.Page(ctx, PricePageParams{PageNumber: 1, PageSize: 20})
	check(err)
	_, err = client.Prices.Update(ctx, UpdatePriceParams{PriceID: "pr_1", Label: "Updated"})
	check(err)
	_, err = client.Prices.Activate(ctx, "pr_1")
	check(err)
	_, err = client.Prices.Deactivate(ctx, "pr_1")
	check(err)
	_, err = client.Spec.Countries(ctx)
	check(err)

	covered := make(map[string]bool, len(paths))
	for _, path := range paths {
		covered[path] = true
	}
	return covered
}

func openAPICoverageResponse() map[string]any {
	return map[string]any{
		"ok":    true,
		"order": map[string]any{"id": "or_1"},
		"refund": map[string]any{
			"id":         "rf_1",
			"order_id":   "or_1",
			"status":     "pending",
			"total":      map[string]any{"currency": "ghs", "value": 100},
			"line_items": []any{},
			"reason":     "requested_by_customer",
			"created_at": "2026-01-01T00:00:00Z",
		},
		"delivery":          map[string]any{"document_kind": "invoice", "document_url": "https://pages.inttegro.com/invoices/or_1"},
		"payment_method":    map[string]any{"id": "pm_1", "customer_id": "cu_1", "type": "mobile_money", "created_at": "2026-01-01T00:00:00Z"},
		"settings":          map[string]any{},
		"transaction":       map[string]any{"id": "bt_1"},
		"account":           map[string]any{"id": "fa_1"},
		"financial_account": map[string]any{"id": "fa_1"},
		"balances":          map[string]any{},
		"customer":          map[string]any{"id": "cu_1"},
		"product":           map[string]any{"id": "prod_1"},
		"price":             map[string]any{"id": "pr_1"},
		"chime":             map[string]any{"id": "ch_1"},
		"scheduled_chime":   map[string]any{"id": "sch_1"},
		"broadcast":         map[string]any{"id": "brc_1"},
		"message_template": map[string]any{
			"id":                      "mtpl_1",
			"channel":                 "sms",
			"created_at":              "2026-01-01T00:00:00Z",
			"draft_version":           1,
			"has_unpublished_changes": false,
			"locale":                  "en",
			"name":                    "welcome_sms",
			"purpose":                 "marketing",
			"status":                  "draft",
			"updated_at":              "2026-01-01T00:00:00Z",
			"version":                 1,
		},
		"rendered":        map[string]any{"channel": "sms"},
		"app":             map[string]any{"id": "app_1"},
		"key":             map[string]any{"id": "sk_1", "token_type": "bearer", "issued_at": "2026-01-01T00:00:00Z", "token": "sk_test_1", "status": "active", "active": true},
		"usage":           map[string]any{"number": 1, "size": 0, "count": 0, "total": 0, "has_more": false, "rows": []any{}},
		"purchase_intent": map[string]any{"id": "sale_1", "application_id": "app_1", "product_id": "prod_1", "price_id": "pr_1", "quantity": map[string]any{"min": 1, "max": 5}, "adjustable_quantity": true, "allow_variants": false, "status": "active", "created_at": "2026-01-01T00:00:00Z"},
		"file":            map[string]any{"id": "file_1", "purpose": "identity", "status": "available"},
		"file_link":       map[string]any{"id": "fl_1", "file_id": "file_1", "status": "active"},
		"url":             "https://files.inttegro.com/open/fl_1",
		"upload_request":  map[string]any{"id": "ur_1", "purpose": "identity", "status": "active"},
		"reconciled":      true,
		"countries":       map[string]any{},
		"page": map[string]any{
			"number":             1,
			"size":               0,
			"count":              0,
			"total":              0,
			"has_more":           false,
			"orders":             []any{},
			"refunds":            []any{},
			"chimes":             []any{},
			"customers":          []any{},
			"keys":               []any{},
			"accounts":           []any{},
			"financial_accounts": []any{},
			"transactions":       []any{},
			"payouts":            []any{},
			"files":              []any{},
			"file_links":         []any{},
			"upload_requests":    []any{},
			"payment_methods":    []any{},
			"products":           []any{},
			"purchase_intents":   []any{},
			"prices":             []any{},
			"message_templates":  []any{},
		},
	}
}

func TestOpenAPICapabilityURLPathExceptionsExist(t *testing.T) {
	specPaths, err := readOpenAPIPaths(openAPISpecPath())
	if err != nil {
		t.Fatal(err)
	}
	specSet := make(map[string]bool, len(specPaths))
	for _, path := range specPaths {
		specSet[path] = true
	}
	var missing []string
	for path := range openAPICapabilityURLPaths {
		if !specSet[path] {
			missing = append(missing, path)
		}
	}
	sort.Strings(missing)
	if len(missing) > 0 {
		t.Fatalf("Go SDK OpenAPI coverage exception no longer exists in spec: %s", strings.Join(missing, ", "))
	}
}
