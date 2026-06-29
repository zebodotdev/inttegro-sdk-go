package commerce

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"
)

func TestProductsEndpointsMatchSpec(t *testing.T) {
	var paths []string
	client, close := newTestClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		paths = append(paths, r.URL.Path)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{
			"product": map[string]any{"id": "prod_123"},
			"price":   map[string]any{"id": "pr_123"},
			"page":    map[string]any{"number": 1, "size": 1, "products": []any{}},
		})
	}))
	if client == nil {
		return
	}
	defer close()

	ctx := context.Background()
	if _, err := client.Products.Create(ctx, CreateProductParams{Type: "physical", Name: "T-Shirt"}); err != nil {
		t.Fatal(err)
	}
	if _, err := client.Products.AddPrice(ctx, AddProductPriceParams{
		ProductID:    "prod_123",
		Amount:       ProductPriceAmount{Currency: "ghs", Value: 5000},
		SetAsDefault: true,
	}); err != nil {
		t.Fatal(err)
	}
	if _, err := client.Products.SetDefaultUnitPrice(ctx, SetDefaultUnitPriceParams{ProductID: "prod_123", PriceID: "pr_123"}); err != nil {
		t.Fatal(err)
	}
	if _, err := client.Products.Lookup(ctx, "prod_123"); err != nil {
		t.Fatal(err)
	}
	if _, err := client.Products.Update(ctx, UpdateProductParams{ProductID: "prod_123", Name: "Updated"}); err != nil {
		t.Fatal(err)
	}
	if _, err := client.Products.Publish(ctx, "prod_123"); err != nil {
		t.Fatal(err)
	}
	if _, err := client.Products.Unpublish(ctx, "prod_123"); err != nil {
		t.Fatal(err)
	}
	if _, err := client.Products.Archive(ctx, "prod_123"); err != nil {
		t.Fatal(err)
	}
	if _, err := client.Products.Page(ctx, PageProductsParams{PageNumber: 1, PageSize: 20}); err != nil {
		t.Fatal(err)
	}

	want := []string{
		"/products/create",
		"/products/add_price",
		"/products/set_default_unit_price",
		"/products/lookup",
		"/products/update",
		"/products/publish",
		"/products/unpublish",
		"/products/archive",
		"/products/page",
	}
	if len(paths) != len(want) {
		t.Fatalf("got paths %v, want %v", paths, want)
	}
	for i := range want {
		if paths[i] != want[i] {
			t.Fatalf("path %d: got %q, want %q", i, paths[i], want[i])
		}
	}
}

func TestPlatformEndpointsMatchSpec(t *testing.T) {
	var paths []string
	client, close := newTestClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		paths = append(paths, r.URL.Path)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{"ok": true})
	}))
	if client == nil {
		return
	}
	defer close()

	ctx := context.Background()
	if _, err := client.Platform.CreateApp(ctx, map[string]any{"name": "App"}); err != nil {
		t.Fatal(err)
	}
	if _, err := client.Platform.GenerateKey(ctx, map[string]any{"app_id": "app_123"}); err != nil {
		t.Fatal(err)
	}
	if _, err := client.Platform.NewSession(ctx, map[string]any{"app_id": "app_123"}); err != nil {
		t.Fatal(err)
	}

	want := []string{"/apps/create", "/keys/generate", "/sessions/new"}
	if len(paths) != len(want) {
		t.Fatalf("got paths %v, want %v", paths, want)
	}
	for i := range want {
		if paths[i] != want[i] {
			t.Fatalf("path %d: got %q, want %q", i, paths[i], want[i])
		}
	}
}
